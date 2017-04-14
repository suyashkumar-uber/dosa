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

package cassandra

import (
	"testing"
	"net"
	"github.com/uber-go/dosa"
	"github.com/stretchr/testify/assert"
	"context"
)

var skipTests = false
var sut dosa.Connector

var ei *dosa.EntityInfo
var ed *dosa.EntityDefinition

func isLocalCassandraRunning() bool {
	l, err := net.Listen("tcp", "localhost:9042")
	if err == nil {
		l.Close()
		return false
	}
	return true
}

func TestConnector_SchemaOps(t *testing.T) {
	if skipTests {
		t.Skip("No local cassandra database to test against")
	}
	ss, err := sut.UpsertSchema(context.TODO(), "test", "test", []*dosa.EntityDefinition{ed})
	assert.Equal(t, int32(1), ss.Version)
	assert.NoError(t, err)
}

func TestConnector_UpsertAndRead(t *testing.T) {
	if skipTests {
		t.Skip("No local cassandra database to test against")
	}
	assert.NotNil(t, sut)

	err := sut.Upsert(context.TODO(), ei, map[string]dosa.FieldValue{
		"booltype": false,
		"int32type": int32(111),
	})
	assert.NoError(t, err)
	vals, err := sut.Read(context.TODO(), ei, map[string]dosa.FieldValue{
		"booltype": false,
	}, []string{"int32type"})
	assert.NoError(t, err)
	assert.NotNil(t, vals)
	assert.Equal(t, dosa.FieldValue(int32(111)), vals["int32type"])
}

func TestUnimplemented(t *testing.T) {
	assert.Panics(t, func() {
		sut.MultiRead(context.TODO(), ei, nil, nil)
	})
	assert.Panics(t, func() {
		sut.MultiRemove(context.TODO(), ei, nil)
	})
	assert.Panics(t, func() {
		sut.MultiUpsert(context.TODO(), ei, nil)
	})
}

func init() {
	if !isLocalCassandraRunning() {
		skipTests = true
		return
	}
	conn, err := dosa.GetConnector("cassandra", map[string]interface{}{})
	if err != nil {
		panic(err)
	}
	sut = conn
	t, err := dosa.TableFromInstance(&AllTypes{})
	ed = &t.EntityDefinition
	ei = &dosa.EntityInfo{
		Ref: &dosa.SchemaRef{NamePrefix: "test", Scope: "test"},
		Def: ed,
	}
}