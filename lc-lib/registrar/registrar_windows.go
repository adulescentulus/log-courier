/*
 * Copyright 2014-2015 Jason Woods.
 *
 * This file is a modification of code from Logstash Forwarder.
 * Copyright 2012-2013 Jordan Sissel and contributors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package registrar

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func (r *Registrar) writeRegistry() error {
	fname := path.Join(r.persistdir, r.statefile)
	tname := fname + ".new"
	file, err := os.Create(tname)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(r.toCanonical())
	file.Close()

	var d_err error
	if _, err = os.Stat(fname); err == nil || !os.IsNotExist(err) {
		d_err = os.Remove(fname)
	}

	err = os.Rename(tname, fname)
	if err != nil {
		return fmt.Errorf("%s -> %s", d_err, err)
	}

	return nil
}
