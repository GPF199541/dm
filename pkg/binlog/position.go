// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package binlog

import (
	"fmt"
	"strconv"
	"strings"

	gmysql "github.com/siddontang/go-mysql/mysql"
	"go.uber.org/zap"

	"github.com/pingcap/dm/pkg/gtid"
	"github.com/pingcap/dm/pkg/log"
	"github.com/pingcap/dm/pkg/terror"
	"github.com/pingcap/dm/pkg/utils"
)

const (
	// in order to differ binlog pos from multi (switched) masters, we added a UUID-suffix field into binlogPos.Name
	// and we also need support: with UUIDSuffix's pos should always > without UUIDSuffix's pos, so we can update from @without to @with automatically
	// conversion: originalPos.NamePrefix + posUUIDSuffixSeparator + UUIDSuffix + binlogFilenameSep + originalPos.NameSuffix => convertedPos.Name
	// UUIDSuffix is the suffix of sub relay directory name, and when new sub directory created, UUIDSuffix is incremented
	// eg. mysql-bin.000003 in c6ae5afe-c7a3-11e8-a19d-0242ac130006.000002 => mysql-bin|000002.000003
	// where `000002` in `c6ae5afe-c7a3-11e8-a19d-0242ac130006.000002` is the UUIDSuffix
	posUUIDSuffixSeparator = "|"
)

var (
	// MinPosition is the min binlog position
	MinPosition = gmysql.Position{Pos: 4}
)

// PositionFromStr constructs a mysql.Position from a string representation like `mysql-bin.000001:2345`
func PositionFromStr(s string) (gmysql.Position, error) {
	parsed := strings.Split(s, ":")
	if len(parsed) != 2 {
		return gmysql.Position{}, terror.ErrBinlogParsePosFromStr.Generatef("the format should be filename:pos, position string %s", s)
	}
	pos, err := strconv.ParseUint(parsed[1], 10, 32)
	if err != nil {
		return gmysql.Position{}, terror.ErrBinlogParsePosFromStr.Generatef("the pos should be digital, position string %s", s)
	}

	return gmysql.Position{
		Name: parsed[0],
		Pos:  uint32(pos),
	}, nil
}

// RealMySQLPos parses a relay position and returns a mysql position and whether error occurs
// if parsed successfully and `UUIDSuffix` exists, sets position Name to
// `originalPos.NamePrefix + binlogFilenameSep + originalPos.NameSuffix`.
// if parsed failed returns the given position and the traced error.
func RealMySQLPos(pos gmysql.Position) (gmysql.Position, error) {
	parsed, err := ParseFilename(pos.Name)
	if err != nil {
		return pos, err
	}

	sepIdx := strings.LastIndex(parsed.BaseName, posUUIDSuffixSeparator)
	if sepIdx > 0 && sepIdx+len(posUUIDSuffixSeparator) < len(parsed.BaseName) {
		if !verifyUUIDSuffix(parsed.BaseName[sepIdx+len(posUUIDSuffixSeparator):]) {
			// NOTE: still can't handle the case where `log-bin` has the format of `mysql-bin|666888`.
			return pos, nil // pos is just the real pos
		}
		return gmysql.Position{
			Name: ConstructFilename(parsed.BaseName[:sepIdx], parsed.Seq),
			Pos:  pos.Pos,
		}, nil
	}

	return pos, nil
}

// ExtractPos extracts (uuidWithSuffix, uuidSuffix, originalPos) from input pos (originalPos or convertedPos)
func ExtractPos(pos gmysql.Position, uuids []string) (uuidWithSuffix string, uuidSuffix string, realPos gmysql.Position, err error) {
	if len(uuids) == 0 {
		err = terror.ErrBinlogExtractPosition.New("empty UUIDs not valid")
		return
	}

	parsed, err := ParseFilename(pos.Name)
	if err != nil {
		return
	}
	sepIdx := strings.LastIndex(parsed.BaseName, posUUIDSuffixSeparator)
	if sepIdx > 0 && sepIdx+len(posUUIDSuffixSeparator) < len(parsed.BaseName) {
		realBaseName, masterUUIDSuffix := parsed.BaseName[:sepIdx], parsed.BaseName[sepIdx+len(posUUIDSuffixSeparator):]
		if !verifyUUIDSuffix(masterUUIDSuffix) {
			err = terror.ErrBinlogExtractPosition.Generatef("invalid UUID suffix %s", masterUUIDSuffix)
			return
		}

		// NOTE: still can't handle the case where `log-bin` has the format of `mysql-bin|666888` and UUID suffix `666888` exists.
		uuid := utils.GetUUIDBySuffix(uuids, masterUUIDSuffix)

		if len(uuid) > 0 {
			// valid UUID found
			uuidWithSuffix = uuid
			uuidSuffix = masterUUIDSuffix
			realPos = gmysql.Position{
				Name: ConstructFilename(realBaseName, parsed.Seq),
				Pos:  pos.Pos,
			}
		} else {
			err = terror.ErrBinlogExtractPosition.Generatef("UUID suffix %s with UUIDs %v not found", masterUUIDSuffix, uuids)
		}
		return
	}

	// use the latest
	var suffixInt int
	uuid := uuids[len(uuids)-1]
	_, suffixInt, err = utils.ParseSuffixForUUID(uuid)
	if err != nil {
		return
	}
	uuidWithSuffix = uuid
	uuidSuffix = utils.SuffixIntToStr(suffixInt)
	realPos = pos // pos is realPos
	return
}

// verifyUUIDSuffix verifies suffix whether is a valid UUID suffix.
func verifyUUIDSuffix(suffix string) bool {
	v, err := strconv.ParseInt(suffix, 10, 64)
	if err != nil || v <= 0 {
		return false
	}
	return true
}

// AdjustPosition adjusts the filename with uuid suffix in mysql position
// for example: mysql-bin|000001.000002 -> mysql-bin.000002
func AdjustPosition(pos gmysql.Position) gmysql.Position {
	realPos, err := RealMySQLPos(pos)
	if err != nil {
		// just return the origin pos
		return pos
	}

	return realPos
}

// ComparePosition returns:
//   1 if pos1 is bigger than pos2
//   0 if pos1 is equal to pos2
//   -1 if pos1 is less than pos2
func ComparePosition(pos1, pos2 gmysql.Position) int {
	adjustedPos1 := AdjustPosition(pos1)
	adjustedPos2 := AdjustPosition(pos2)

	// means both pos1 and pos2 have uuid in name, so need also compare the uuid
	if adjustedPos1.Name != pos1.Name && adjustedPos2.Name != pos2.Name {
		return pos1.Compare(pos2)
	}

	return adjustedPos1.Compare(adjustedPos2)
}

// Location is used for save binlog's position and gtid
type Location struct {
	Position gmysql.Position

	GTIDSet gtid.Set
}

// NewLocation returns a new Location
func NewLocation(flavor string) Location {
	return Location{
		Position: MinPosition,
		GTIDSet:  gtid.MinGTIDSet(flavor),
	}
}

func (l Location) String() string {
	return fmt.Sprintf("position: %v, gtid-set: %s", l.Position, l.GTIDSetStr())
}

// GTIDSetStr returns gtid set's string
func (l Location) GTIDSetStr() string {
	gsetStr := ""
	if l.GTIDSet != nil {
		gsetStr = l.GTIDSet.String()
	}

	return gsetStr
}

// Clone clones a same Location
func (l Location) Clone() Location {
	return l.CloneWithFlavor("")
}

// CloneWithFlavor clones the location, and if the GTIDSet is nil, will create a GTIDSet with specified flavor.
func (l Location) CloneWithFlavor(flavor string) Location {
	var newGTIDSet gtid.Set
	if l.GTIDSet != nil {
		newGTIDSet = l.GTIDSet.Clone()
	} else if len(flavor) != 0 {
		newGTIDSet = gtid.MinGTIDSet(flavor)
	}

	return Location{
		Position: gmysql.Position{
			Name: l.Position.Name,
			Pos:  l.Position.Pos,
		},
		GTIDSet: newGTIDSet,
	}
}

// CompareLocation returns:
//   1 if point1 is bigger than point2
//   0 if point1 is equal to point2
//   -1 if point1 is less than point2
func CompareLocation(location1, location2 Location) int {
	if location1.GTIDSet != nil && len(location1.GTIDSet.String()) != 0 &&
		location2.GTIDSet != nil && len(location2.GTIDSet.String()) != 0 {
		contain1 := location1.GTIDSet.Contain(location2.GTIDSet)
		contain2 := location2.GTIDSet.Contain(location1.GTIDSet)
		if contain1 && contain2 {
			// gtidSet1 contains gtidSet2 and gtidSet2 contains gtidSet1 means gtidSet1 equals to gtidSet2,
			// then need to compare by position.
		} else {
			if contain1 {
				return 1
			} else if contain2 {
				return -1
			}

			// can't compare location by gtid, and will compare by position
			log.L().Warn("gtidSet can't be compared", zap.Stringer("location1", location1), zap.Stringer("location2", location2))
		}
	}

	return ComparePosition(location1.Position, location2.Position)
}
