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
	"io"
)

var defaultPool = New()

// Get a bufio.Reader for given io.Reader.
// Reads a bufio.Reader from the pool and resets it.
func Get(reader io.Reader) (r *bufio.Reader) {
	return defaultPool.Get(reader)
}

// Put a bufio.Reader to the pool if possible.
func Put(r *bufio.Reader) {
	defaultPool.Put(r)
}
