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

package examples

import "github.com/uber-go/dosa"

// SinglePrimaryKey is used for testing and examples
type SinglePrimaryKey struct {
	dosa.Entity `dosa:"primaryKey=(PrimaryKey)"`
	PrimaryKey  int64
	Data        string
}

// SinglePrimaryKeyNoParen is used for testing and examples
type SinglePrimaryKeyNoParen struct {
	dosa.Entity `dosa:"primaryKey=PrimaryKey"`
	PrimaryKey  int64
	Data        string
}

// SinglePartitionKey is used for testing and examples
type SinglePartitionKey struct {
	dosa.Entity `dosa:"primaryKey=PrimaryKey"`
	PrimaryKey  int64
	data        string
}

// PrimaryKeyWithSecondaryRange is used for testing and examples
type PrimaryKeyWithSecondaryRange struct {
	dosa.Entity    `dosa:"primaryKey=(PartKey,PrimaryKey)"`
	PartKey        int64
	PrimaryKey     int64
	data, moredata string
}

// PrimaryKeyWithDescendingRange is used for testing and examples
type PrimaryKeyWithDescendingRange struct {
	dosa.Entity `dosa:"primaryKey=(PartKey,PrimaryKey desc)"`
	PartKey     int64
	PrimaryKey  int64
	data        string
}

// MultiComponentPrimaryKey is used for testing and examples
type MultiComponentPrimaryKey struct {
	dosa.Entity    `dosa:"primaryKey=((PartKey, AnotherPartKey))"`
	PartKey        int64
	AnotherPartKey int64
	data           string
}
