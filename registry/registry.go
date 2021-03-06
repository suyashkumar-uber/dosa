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

package registry

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/uber-go/dosa"
	"github.com/uber-go/dosa/config"
)

type registrar struct {
	scope   string
	prefix  string
	baseFQN dosa.FQN
	idx     map[dosa.FQN]*dosa.RegisteredEntity
}

// Scope returns the registrar's scope.
func (r *registrar) Scope() string {
	return r.scope
}

// NamePrefix returns the registrar's prefix.
func (r *registrar) NamePrefix() string {
	return r.prefix
}

// Find looks at its internal index to find a registration that matches the
// entity instance provided. Return an error when not found.
func (r *registrar) Find(entity dosa.DomainObject) (*dosa.RegisteredEntity, error) {
	name := reflect.TypeOf(entity).Name()
	fqn, _ := r.baseFQN.Child(name)
	re, ok := r.idx[fqn]
	if !ok {
		return nil, errors.Errorf("failed to find registration for entity %s", name)
	}
	return re, nil
}

// FindAll returns all registered entities from its internal index.
func (r *registrar) FindAll() ([]*dosa.RegisteredEntity, error) {
	res := []*dosa.RegisteredEntity{}
	for _, re := range r.idx {
		res = append(res, re)
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("registry.FindAll returned empty")
	}
	return res, nil
}

// NewRegistrar returns a new Registrar for the configuration provided.
func NewRegistrar(cfg *config.Config) (dosa.Registrar, error) {
	idx := make(map[dosa.FQN]*dosa.RegisteredEntity)
	baseFQN, err := dosa.ToFQN(cfg.NamePrefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct Registrar")
	}

	// "warnings" mean entity was found but contained invalid annotations
	eds, warns, err := dosa.FindEntities(cfg.EntityPaths, cfg.Excludes)
	if len(warns) > 0 {
		return nil, dosa.NewEntityErrors(warns)
	}
	// I/O and AST parsing errors
	if err != nil {
		return nil, err
	}

	// index entity definitions (aka "tables")
	for _, ed := range eds {
		re := dosa.NewRegisteredEntity(cfg.Scope, cfg.NamePrefix, ed)

		// index by prefix + normalized name (aka "FQN")
		fqn, _ := baseFQN.Child(ed.StructName)
		idx[fqn] = re
	}

	return &registrar{
		scope:   cfg.Scope,
		prefix:  cfg.NamePrefix,
		baseFQN: baseFQN,
		idx:     idx,
	}, nil
}
