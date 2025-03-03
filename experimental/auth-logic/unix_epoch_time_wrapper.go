// Copyright 2022 The Project Oak Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package authlogic contains logic and tests for interfacing with the
// authorization logic compiler
package authlogic

import (
	"fmt"
	"time"
)

// This file contains a wrapper that produces the current
// time as the number of seconds since [the unix
// epoch](https://en.wikipedia.org/wiki/Unix_time).

type UnixEpochTime struct{}

func (timeWrapper UnixEpochTime) EmitStatement() UnattributedStatement {
	epochTime := time.Now().Unix()
	return UnattributedStatement{Contents: fmt.Sprintf("RealTimeIs(%v).", epochTime)}
}
