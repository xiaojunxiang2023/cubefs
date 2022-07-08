// Copyright 2022 The CubeFS Authors.
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

package normaldb

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/cubefs/cubefs/blobstore/common/kvstore"

	"github.com/stretchr/testify/assert"
)

func TestScopeTbl(t *testing.T) {
	tmpDBPath := "/tmp/tmpnormaldb" + strconv.Itoa(rand.Intn(1000000000))
	defer os.RemoveAll(tmpDBPath)

	db, err := OpenNormalDB(tmpDBPath, false)
	assert.NoError(t, err)
	defer db.Close()

	scopeTbl, err := OpenScopeTable(db)
	assert.NoError(t, err)

	key1 := "testkey1"
	key2 := "testkey2"

	current, err := scopeTbl.Get("testkey")
	assert.Equal(t, kvstore.ErrNotFound, err)
	assert.Equal(t, uint64(0), current)

	err = scopeTbl.Put(key1, 1)
	assert.NoError(t, err)

	current, err = scopeTbl.Get(key1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), current)

	err = scopeTbl.Put(key1, 2)
	assert.NoError(t, err)

	err = scopeTbl.Put(key2, 1)
	assert.NoError(t, err)

	scopes, err := scopeTbl.Load()
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), scopes[key1])
	assert.Equal(t, uint64(1), scopes[key2])
}
