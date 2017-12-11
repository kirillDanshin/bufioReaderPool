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

package bufioReaderPool

import (
	"bufio"
	"net"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("New() must not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fail()
			}
		}()
		New()
	})
	t.Run("New() must not return nil", func(t *testing.T) {
		if New() == nil {
			t.Fail()
		}
	})
}

func TestNewSize(t *testing.T) {
	t.Run("NewSize() must not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fail()
			}
		}()
		NewSize(65536)
	})
	t.Run("NewSize() must not return nil", func(t *testing.T) {
		if NewSize(65536) == nil {
			t.Fail()
		}
	})
}

func TestPool_Get(t *testing.T) {
	brp := New()
	r := &net.TCPConn{}
	t.Run("Get() must not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fail()
			}
		}()
		brp.Get(r)
	})
	t.Run("Get() must not return nil", func(t *testing.T) {
		if brp.Get(r) == nil {
			t.Fail()
		}
	})
}

func BenchmarkPool(b *testing.B) {
	p := New()
	conn := &net.TCPConn{}
	r := bufio.NewReader(conn)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(r)
			p.Get(conn)
		}
	})
}

func BenchmarkNoPool(b *testing.B) {
	conn := &net.TCPConn{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := bufio.NewReader(conn)
			_ = r
		}
	})
}
