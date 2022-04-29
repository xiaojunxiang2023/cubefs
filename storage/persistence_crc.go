// Copyright 2018 The Chubao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package storage

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cubefs/cubefs/util"
)

type BlockCrc struct {
	BlockNo int
	Crc     uint32
}
type BlockCrcArr []*BlockCrc

const (
	BaseExtentIDOffset = 0
)

func (arr BlockCrcArr) Len() int           { return len(arr) }
func (arr BlockCrcArr) Less(i, j int) bool { return arr[i].BlockNo < arr[j].BlockNo }
func (arr BlockCrcArr) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }

type UpdateCrcFunc func(e *Extent, blockNo int, crc uint32, data []byte) (err error)
type GetExtentCrcFunc func(extentID uint64) (crc uint32, err error)

// If data is not nil, write the whole extent header to file.
// Otherwise, just update the extent header in memory.
func (s *ExtentStore) PersistenceBlockCrc(e *Extent, blockNo int, blockCrc uint32, data []byte) (err error) {
	if data != nil {
		if len(data) > util.BlockHeaderSize {
			return fmt.Errorf("PersistenceBlockCrc: error! data length(%v) which islarger than (%v)", len(data), util.BlockHeaderSize)
		}
		verifyStart := int(util.BlockHeaderSize * e.extentID)
		if _, err = s.verifyExtentFp.WriteAt(data, int64(verifyStart)); err != nil {
			return
		}
	}

	startIdx := blockNo * util.PerBlockCrcSize
	endIdx := startIdx + util.PerBlockCrcSize
	binary.BigEndian.PutUint32(e.header[startIdx:endIdx], blockCrc)
	return
}

func (s *ExtentStore) PunchBlockCRC(from, count int) error {
	return fallocate(int(s.verifyExtentFp.Fd()), FallocFLPunchHole|FallocFLKeepSize,
		int64(util.BlockHeaderSize*from), int64(count*util.BlockHeaderSize))
}

func (s *ExtentStore) GetHasDeleteExtent() (extentDes []ExtentDeleted, err error) {
	data := make([]byte, 8)
	offset := int64(0)
	for {
		_, err = s.normalExtentDeleteFp.ReadAt(data, offset)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		extent := ExtentDeleted{}
		extent.ExtentID = binary.BigEndian.Uint64(data)
		extentDes = append(extentDes, extent)
		offset += 8
	}

	return
}
