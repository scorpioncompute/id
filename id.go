// Copyright 2016, 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package id

import (
	"errors"
)


func Encode(shardBitSize uint, typeBitSize uint, localBitSize uint) (func(uint64, uint64, uint64) uint64, error) {
	if shardBitSize + typeBitSize + localBitSize > 64 {
		return nil, errors.New("total number of bits cannot exceed 64")
	}

	shardShift := typeBitSize + localBitSize
	typeShift := localBitSize

	return func(shardID uint64, typeID uint64, localID uint64) uint64 {
		return (shardID << shardShift) | (typeID << typeShift) | (localID)
	}, nil
}

func Decode(shardBitSize uint, typeBitSize uint, localBitSize uint) (func(uint64) (uint64, uint64, uint64), error) {
	if shardBitSize + typeBitSize + localBitSize > 64 {
		return nil, errors.New("total number of bits cannot exceed 64")
	}

	shardShift := typeBitSize + localBitSize
	typeShift := localBitSize

	return func(ID uint64) (uint64, uint64, uint64) {
		return (ID >> shardShift) & 0xFFFF, (ID >> typeShift) & 0x3FF, ID & 0xFFFFFFFFF
	}, nil
}