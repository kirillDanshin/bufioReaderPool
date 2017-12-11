// Copyright 2017-present Kirill Danshin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package bufioReaderPool provides bufio.Reader as a sync.Pool
// in an resource-effective way.
//
//
//	brp := bufioReaderPool.New()
//	reader := brp.Get()
//	defer brp.Put(reader)
//	// work with reader
package bufioReaderPool

import (
	"bufio"
	"io"
	"sync"

	"github.com/gramework/runtimer"
	"github.com/gramework/utils/nocopy"
)

// Pool describes bufioReader pool.
type Pool struct {
	pool *sync.Pool

	nocopy.NoCopy
}

// New allocates an empty bufioReader pool.
func New() *Pool {
	return &Pool{
		pool: &sync.Pool{
			New: func() interface{} {
				return bufio.NewReader(nil)
			},
		},
	}
}

// NewSize allocates an empty bufioReader pool with given size.
func NewSize(size int) *Pool {
	return &Pool{
		pool: &sync.Pool{
			New: func() interface{} {
				return bufio.NewReaderSize(nil, size)
			},
		},
	}
}

// Get a bufio.Reader for given io.Reader.
// Reads a bufio.Reader from the pool and resets it.
func (p *Pool) Get(reader io.Reader) (r *bufio.Reader) {
	r = (*bufio.Reader)(runtimer.GetEfaceDataPtr(p.pool.Get()))
	r.Reset(reader)
	return
}

// Put a bufio.Reader to the pool if possible.
func (p *Pool) Put(r *bufio.Reader) {
	if r != nil {
		p.pool.Put(r)
	}
}
