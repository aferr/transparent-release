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
	"io/ioutil"
	"regexp"
	"strconv"
	"testing"
)

const (
	// This is March 24, 2022
	pastDate = 1648146779
	// This is January 1, 3022
	futureDate = 33197947200
)

func (time UnixEpochTime) Identify() Principal {
	return Principal{Contents: "UnixEpochTime"}
}

func TestUnixEpochTimeWrapper(t *testing.T) {
	handleErr := func(err error) {
		if err != nil {
			t.Fatalf("test generated error %v", err)
			panic(err)
		}
	}

	// Write a statement from the time wrapper to a file
	writeErr := EmitWrapperStatement(UnixEpochTime{}, "wrapped_time.auth_logic")
	handleErr(writeErr)

	// Read the contents of the file
	fileReadBytes, readErr := ioutil.ReadFile("wrapped_time.auth_logic")
	handleErr(readErr)
	fileReadString := string(fileReadBytes)

	timeTestRegex := regexp.MustCompile("UnixEpochTime says {\nRealTimeIs\\(([0-9]+)\\).\n}")
	match := timeTestRegex.FindStringSubmatch(fileReadString)
	if len(match) != 2 {
		t.Errorf("Result of time wrapper did not have valid format. Got: %v.",
			fileReadString)
	}

	timeValue, conversionErr := strconv.Atoi(match[1])
	handleErr(conversionErr)

	if timeValue < pastDate {
		t.Errorf("The emitted current time %v, already happened", timeValue)
	}

	if timeValue > futureDate {
		t.Errorf("The emitted current time %v, is far into the future", timeValue)
	}

}
