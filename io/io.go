//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package io

import (
	"io/ioutil"
	"os"
	"path"
)

const Perm = 0777

func WriteFile(fpath string, data []byte) {
	dir, _ := path.Split(fpath)
	os.MkdirAll(dir, Perm)
	ioutil.WriteFile(fpath, data, Perm)
}

func WriteFileString(fpath string, data string) {
	WriteFile(fpath, []byte(data))
}
