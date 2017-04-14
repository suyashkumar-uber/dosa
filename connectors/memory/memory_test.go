// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package memory_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/dosa"
	"github.com/uber-go/dosa/connectors/memory"
)

type TestEntity struct {
	dosa.Entity `dosa:"primaryKey=UUID"`
	UUID        dosa.UUID
}

var (
	conn         = memory.NewConnector()
	testTable, _ = dosa.TableFromInstance((*TestEntity)(nil))
	testInfo     = &dosa.EntityInfo{
		Def: &testTable.EntityDefinition,
		Ref: &dosa.SchemaRef{
			Scope:      "testScope",
			NamePrefix: "testPrefix",
			EntityName: "testEntityName",
		},
	}
	testPairs       = dosa.FieldNameValuePair{}
	testValues      = make(map[string]dosa.FieldValue)
	testMultiValues = make([]map[string]dosa.FieldValue, 50)
	fieldsToRead    = []string{"uuid"}
	ctx             = context.Background()
)

func TestMemory_CreateIfNotExists(t *testing.T) {
	assert.NoError(t, conn.CreateIfNotExists(ctx, testInfo, testValues))
}

func TestMemory_Read(t *testing.T) {
	val, err := conn.Read(ctx, testInfo, testValues, fieldsToRead)
	assert.NoError(t, err)
	assert.NotNil(t, val)
	for _, field := range fieldsToRead {
		assert.NotNil(t, val[field])
	}
}

func TestMemory_MultiRead(t *testing.T) {
	v, e := conn.MultiRead(ctx, testInfo, testMultiValues, fieldsToRead)
	assert.NotNil(t, v)
	assert.Error(t, e)
}

func TestMemory_Upsert(t *testing.T) {
	assert.NoError(t, conn.Upsert(ctx, testInfo, testValues))
}

func TestMemory_MultiUpsert(t *testing.T) {
	errs, err := conn.MultiUpsert(ctx, testInfo, testMultiValues)
	assert.NotNil(t, errs)
	assert.Error(t, err)
}

func TestMemory_Remove(t *testing.T) {
	assert.NoError(t, conn.Remove(ctx, testInfo, testValues))
}

func TestMemory_MultiRemove(t *testing.T) {
	errs, err := conn.MultiRemove(ctx, testInfo, testMultiValues)
	assert.NotNil(t, errs)
	assert.Error(t, err)
}

func TestMemory_Range(t *testing.T) {
	conditions := make(map[string][]*dosa.Condition)
	vals, _, err := conn.Range(ctx, testInfo, conditions, fieldsToRead, "", 32)
	assert.NotNil(t, vals)
	assert.NoError(t, err)
}

func TestMemory_Search(t *testing.T) {
	vals, _, err := conn.Search(ctx, testInfo, testPairs, fieldsToRead, "", 32)
	assert.NotNil(t, vals)
	assert.NoError(t, err)
}

func TestMemory_Scan(t *testing.T) {
	vals, _, err := conn.Scan(ctx, testInfo, fieldsToRead, "", 32)
	assert.NotNil(t, vals)
	assert.NoError(t, err)
}

func TestMemory_CheckSchema(t *testing.T) {
	defs := make([]*dosa.EntityDefinition, 4)
	version, err := conn.CheckSchema(ctx, "testScope", "testPrefix", defs)
	assert.NotNil(t, version)
	assert.NoError(t, err)
}

func TestMemory_UpsertSchema(t *testing.T) {
	defs := make([]*dosa.EntityDefinition, 4)
	status, err := conn.UpsertSchema(ctx, "testScope", "testPrefix", defs)
	assert.NotNil(t, status)
	assert.NoError(t, err)
}

func TestMemory_CheckSchemaStatus(t *testing.T) {
	status, err := conn.CheckSchemaStatus(ctx, "testScope", "testPrefix", 1)
	assert.NotNil(t, status)
	assert.NoError(t, err)
}

func TestMemory_CreateScope(t *testing.T) {
	assert.NoError(t, conn.CreateScope(ctx, ""))
}

func TestMemory_TruncateScope(t *testing.T) {
	assert.NoError(t, conn.TruncateScope(ctx, ""))
}

func TestMemory_DropScope(t *testing.T) {
	assert.NoError(t, conn.DropScope(ctx, ""))
}

func TestMemory_ScopeExists(t *testing.T) {
	exists, err := conn.ScopeExists(ctx, "")
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestMemory_Shutdown(t *testing.T) {
	assert.NoError(t, conn.Shutdown())
}
