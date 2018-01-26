// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ip2location

import (
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	s := NewMemoryStore(1024*1024*1024, time.Hour*24*365*10)
	s.SetLocation("8.8.8.8", &Location{"中国", "湖北", "武汉", "电信"})

	loc := s.GetLocation("8.8.8.8")
	t.Log(loc)
}
