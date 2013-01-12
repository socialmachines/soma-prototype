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

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the BSD-LICENSE file.

package scan

type Token int

const (
	UNKN Token = iota
	ENDF

	DEFN
	IDEN
	NAME
	PLUS
)

var tokens = [...]string{
	UNKN: "UNKNOWN",
	ENDF: "ENDF",

	DEFN: "=>",
	IDEN: "IDEN",
	NAME: "NAME",
	PLUS: "PLUS",
}