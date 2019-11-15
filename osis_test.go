// Copyright 2019, Honza Pokorny <me@honza.ca>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package osis

import (
	"testing"
)

func TestBasic(t *testing.T) {
	pairs := []struct{ in, out string }{
		{"Ps", "Psalm"},
		{"Rom.8", "Romans 8"},
		{"John.3.16", "John 3:16"},
		{"Jude.1.1", "Jude 1"},
		{"Gen.1-Gen.2", "Genesis 1-2"},
		{"Gen.1.1-Gen.1.2", "Genesis 1:1-1:2"},
		{"Ps.1.1", "Psalm 1:1"},
		{"Ps.1-Ps.2", "Psalms 1-2"},
		{"John.7.53-John.8.11", "John 7:53-8:11"},
		{"Luke.12-Acts.1", "Luke 12-Acts 1"},
		{"Luke.12.1-Acts.1.1", "Luke 12:1-Acts 1:1"},
	}

	for _, p := range pairs {
		out, err := Format(p.in)
		if err != nil {
			t.Fail()
		}

		if p.out != out {
			t.Errorf("Expected: '%s', got: '%s'", p.out, out)
		}

	}

}

func TestGarbage(t *testing.T) {
	_, err := Format("haha.1.1")

	if err == nil {
		t.Fail()
	}

	_, err = Format("haha")
	if err == nil {
		t.Fail()
	}
}

func TestMany(t *testing.T) {
	pairs := []struct {
		in  string
		out []string
	}{
		{"John.3.16,Ps.1.1", []string{"John 3:16", "Psalm 1:1"}},
	}

	for _, p := range pairs {
		out, err := FormatMany(p.in)
		if err != nil {
			t.Fail()
		}

		if p.out[0] != out[0] {
			t.Errorf("Expected: '%s', got: '%s'", p.out[0], out[0])
		}

		if p.out[1] != out[1] {
			t.Errorf("Expected: '%s', got: '%s'", p.out[1], out[1])
		}

	}

}
