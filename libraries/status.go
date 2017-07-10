/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package libraries

import "github.com/pmylund/sortutil"
import "fmt"

// StatusContext keeps the current status of the libraries in the system
// (the list of libraries, revisions, installed paths, etc.)
type StatusContext struct {
	Libraries map[string]*Library `json:"libraries"`
}

// AddLibrary adds an indexRelease to the status context
func (l *StatusContext) AddLibrary(indexLib *indexRelease) {
	name := indexLib.Name
	if l.Libraries[name] == nil {
		l.Libraries[name] = indexLib.extractLibrary()
	} else {
		release := indexLib.extractRelease()
		lib := l.Libraries[name]
		lib.Releases[fmt.Sprint(release.Version)] = release
	}
}

// Names returns an array with all the names of the registered libraries.
func (l *StatusContext) Names() []string {
	res := make([]string, len(l.Libraries))
	i := 0
	for n := range l.Libraries {
		res[i] = n
		i++
	}
	sortutil.CiAsc(res)
	return res
}

// CreateStatusContext creates a status context from index data.
func (index *Index) CreateStatusContext() (*StatusContext, error) {
	// Start with an empty status context
	libraries := StatusContext{
		Libraries: map[string]*Library{},
	}
	for _, lib := range index.Libraries {
		// Add all indexed libraries in the status context
		libraries.AddLibrary(&lib)
	}
	return &libraries, nil
}
