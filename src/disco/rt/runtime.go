// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package rt

import (
	"disco/ast"
	"math/rand"
	"runtime"
)

var RT *Runtime

type Runtime struct {
	Global *Scope

	ID    uint32
	Procs int
}

func (r *Runtime) Init() *Runtime {
	procs := runtime.NumCPU()
	runtime.GOMAXPROCS(procs)

	return &Runtime{NewScope(nil, nil), rand.Uint32(), procs}
}

func (r *Runtime) StartDefine(def ast.Define) {
}

func init() {
	RT = RT.Init()
}
