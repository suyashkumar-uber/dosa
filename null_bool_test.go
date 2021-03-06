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

package dosa

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestNullBool(t *testing.T) {
	v := NewNullBool(true)
	actual, err := v.Get()
	assert.NoError(t, err)
	assert.True(t, actual)

	v.Set(false)
	actual, err = v.Get()
	assert.NoError(t, err)
	assert.False(t, actual)

	v.Nullify()
	actual, err = v.Get()
	assert.Error(t, err, "Value is null")
}

func TestNullBool_MarshalText(t *testing.T) {
	v := NewNullBool(true)
	bytes, err := v.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "true", string(bytes))

	v = NewNullBool(false)
	bytes, err = v.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "false", string(bytes))

	v.Nullify()
	bytes, err = v.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "", string(bytes))
}

func TestNullBool_UnmarshalText(t *testing.T) {
	var v NullBool
	err := v.UnmarshalText([]byte("true"))
	assert.NoError(t, err)

	value, err := v.Get()
	assert.NoError(t, err)
	assert.True(t, value)

	err = v.UnmarshalText([]byte("false"))
	assert.NoError(t, err)

	value, err = v.Get()
	assert.NoError(t, err)
	assert.False(t, value)

	err = v.UnmarshalText([]byte(""))
	assert.NoError(t, err)
	_, err = v.Get()
	assert.Equal(t, ErrNullValue, err)

	err = v.UnmarshalText([]byte("null"))
	assert.NoError(t, err)
	_, err = v.Get()
	assert.Equal(t, ErrNullValue, err)
}

func TestNullBool_MarshalJSON(t *testing.T) {
	v := NewNullBool(true)
	bytes, err := v.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "true", string(bytes))

	v.Set(false)
	bytes, err = v.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "false", string(bytes))

	v.Nullify()
	bytes, err = v.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "null", string(bytes))
}

func TestNullBool_UnmarshalJSON(t *testing.T) {
	var v NullBool
	err := v.UnmarshalJSON([]byte("true"))
	assert.NoError(t, err)

	value, err := v.Get()
	assert.NoError(t, err)
	assert.True(t, value)

	err = v.UnmarshalJSON([]byte("null"))
	assert.NoError(t, err)
	value, err = v.Get()
	assert.EqualError(t, err, "Value is null")

	// Invalid case
	badData, err := json.Marshal(float64(3.14))
	err = v.UnmarshalJSON(badData)
	assert.Error(t, err)
}
