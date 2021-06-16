/*
Copyright © 2021 Mingfei Huang <himax1023@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	for _, tc := range []struct {
		name string

		url        string
		wantScheme string
		wantHost   string
		wantPort   string
		wantQuery  string
		wantErr    string
	}{
		{
			url:        "https://foo.com:80/bar/baz?a=b#yolo",
			wantScheme: "https",
			wantHost:   "foo.com",
			wantPort:   "80",
			wantQuery:  "/bar/baz?a=b#yolo",
		},
		{
			url:        "http://foo.com:80/bar/baz?a=b#yolo",
			wantScheme: "http",
			wantHost:   "foo.com",
			wantPort:   "80",
			wantQuery:  "/bar/baz?a=b#yolo",
		},
		{
			url:        "foo.com:80/bar/baz?a=b#yolo",
			wantScheme: "",
			wantHost:   "foo.com",
			wantPort:   "80",
			wantQuery:  "/bar/baz?a=b#yolo",
		},
		{
			url:        "foo.com/bar/baz?a=b#yolo",
			wantScheme: "",
			wantHost:   "foo.com",
			wantPort:   "",
			wantQuery:  "/bar/baz?a=b#yolo",
		},
		{
			url:        "foo.com",
			wantScheme: "",
			wantHost:   "foo.com",
			wantPort:   "",
			wantQuery:  "",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			gotScheme, gotHost, gotPort, gotQuery, gotErr := Parse(tc.url)
			if gotErr != nil {
				if tc.wantErr == "" {
					// got err, don't expect error
					t.Errorf("Unexpected error %v", gotErr)
				} else {
					// got err, want err, check err matches
					assert.Contains(gotErr.Error(), tc.wantErr)
				}
				return
			}
			if tc.wantErr != "" {
				t.Errorf("Expecting error %v, got %v", tc.wantErr, map[string]string{"scheme": gotScheme, "host": gotHost, "port": gotPort, "query": gotQuery})
				return
			}

			assert.Equal(tc.wantScheme, gotScheme)
			assert.Equal(tc.wantHost, gotHost)
			assert.Equal(tc.wantPort, gotPort)
			assert.Equal(tc.wantQuery, gotQuery)
		})
	}
}
