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

package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/uber-go/dosa"
)

const (
	version = 1
	status  = "APPLIED"
)

// Connector maintains an in-memory map of DOSA entities. Entities are indexed
// by struct name. Any scan, range or multi* query will return the entire list
// for simplicity. If more fine-grained functionality is needed, consider
// using the connectors in the mocks package.
type Connector struct {
	lock sync.RWMutex
	idx  map[string][]map[string]dosa.FieldValue
}

// NewConnector returns a dosa.Connector implementation that does all
// operations in memory which can be used for testing.
func NewConnector() dosa.Connector {
	return &Connector{
		idx: make(map[string][]map[string]dosa.FieldValue),
	}
}

// CreateIfNotExists always adds the given entity to the connector's internal index.
func (c *Connector) CreateIfNotExists(ctx context.Context, ei *dosa.EntityInfo, values map[string]dosa.FieldValue) error {
	if ei.Ref == nil || ei.Ref.EntityName == "" {
		return errors.New("invalid entity info")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := c.idx[ei.Ref.EntityName]
	if !ok {
		c.idx[ei.Ref.EntityName] = []map[string]dosa.FieldValue{values}
	}
	c.idx[ei.Ref.EntityName] = append(v, values)
	return nil
}

// Read returns the first entity matching the given entity info from the connector's internal index.
func (c *Connector) Read(ctx context.Context, ei *dosa.EntityInfo, keys map[string]dosa.FieldValue, fieldsToRead []string) (map[string]dosa.FieldValue, error) {
	if ei.Ref == nil || ei.Ref.EntityName == "" {
		return nil, errors.New("invalid entity info")
	}
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.idx[ei.Ref.EntityName][0], nil
}

// MultiRead is not implemented and should not be used yet.
func (c *Connector) MultiRead(ctx context.Context, ei *dosa.EntityInfo, keys []map[string]dosa.FieldValue, fieldsToRead []string) ([]*dosa.FieldValuesOrError, error) {
	return make([]*dosa.FieldValuesOrError, 1), errors.New("not implemented")
}

// Upsert creates or updates an entity in the connector's internal index.
func (c *Connector) Upsert(ctx context.Context, ei *dosa.EntityInfo, values map[string]dosa.FieldValue) error {
	return errors.New("not implemented")
}

// MultiUpsert is not implemented and should not be used yet.
func (c *Connector) MultiUpsert(ctx context.Context, ei *dosa.EntityInfo, multiValues []map[string]dosa.FieldValue) ([]error, error) {
	return []error{errors.New("not implemented")}, errors.New("not implemented")
}

// Remove removes an entity from the connector's internal index.
func (c *Connector) Remove(ctx context.Context, ei *dosa.EntityInfo, keys map[string]dosa.FieldValue) error {
	return errors.New("not implemented")
}

// MultiRemove is not implemented and should not be used yet.
func (c *Connector) MultiRemove(ctx context.Context, ei *dosa.EntityInfo, multiKeys []map[string]dosa.FieldValue) ([]error, error) {
	return []error{errors.New("not implemented")}, errors.New("not implemented")
}

// Range returns all entities matching the given entity info. Ordering will match the order in which entities were created.
func (c *Connector) Range(ctx context.Context, ei *dosa.EntityInfo, columnConditions map[string][]*dosa.Condition, fieldsToRead []string, token string, limit int) ([]map[string]dosa.FieldValue, string, error) {
	return nil, "", errors.New("not implemented")
}

// Search is not implemented and should not be used yet.
func (c *Connector) Search(ctx context.Context, ei *dosa.EntityInfo, fieldPairs dosa.FieldNameValuePair, fieldsToRead []string, token string, limit int) ([]map[string]dosa.FieldValue, string, error) {
	return nil, "", errors.New("not implemented")
}

// Scan is not implemented and should not be used yet.
func (c *Connector) Scan(ctx context.Context, ei *dosa.EntityInfo, fieldsToRead []string, token string, limit int) ([]map[string]dosa.FieldValue, string, error) {
	return nil, "", errors.New("not implemented")
}

// CheckSchema returns a constant version.
func (c *Connector) CheckSchema(ctx context.Context, scope string, namePrefix string, ed []*dosa.EntityDefinition) (int32, error) {
	return version, nil
}

// UpsertSchema returns a SchemaStatus with constant version and status.
func (c *Connector) UpsertSchema(ctx context.Context, scope string, namePrefix string, ed []*dosa.EntityDefinition) (*dosa.SchemaStatus, error) {
	return &dosa.SchemaStatus{
		Status:  status,
		Version: version,
	}, nil
}

// CheckSchemaStatus returns a SchemaStatus with constant version and status.
func (c *Connector) CheckSchemaStatus(ctx context.Context, scope string, namePrefix string, version int32) (*dosa.SchemaStatus, error) {
	return &dosa.SchemaStatus{
		Status:  status,
		Version: version,
	}, nil
}

// CreateScope always returns nil
func (c *Connector) CreateScope(ctx context.Context, scope string) error {
	return nil
}

// TruncateScope always returns nil
func (c *Connector) TruncateScope(ctx context.Context, scope string) error {
	return nil
}

// DropScope always returns nil
func (c *Connector) DropScope(ctx context.Context, scope string) error {
	return nil
}

// ScopeExists always returns true
func (c *Connector) ScopeExists(ctx context.Context, scope string) (bool, error) {
	return true, nil
}

// Shutdown is a no-op
func (c *Connector) Shutdown() error {
	return nil
}
