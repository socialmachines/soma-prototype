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

package file

import (
	"fmt"
	"os"
	"path"
)

func CreateProjectDir(name string, pwd string) {
	fs := createProjectFS(pwd, name)

	createProjectDir(fs, name, "doc")
	createProjectDir(fs, name, "lib")
	createProjectDir(fs, name, "src")
	
	pfile := fmt.Sprintf("%s/%s.disco", "src", name) 
	createProjectFile(fs, name, pfile) 
}

func createProjectFS(pwd string, name string) (dir string) {
	dir = path.Join(pwd, name)

	err := os.Mkdir(dir, 0700)
	if err != nil {
		fmt.Printf("disco create: error creating project directory: %s\n", err)
		os.Exit(0)
	}
	
	fmt.Printf("    created %s/\n", name)
	createProjectDir(dir, name, ".disco")
	
	return
}

func createProjectDir(pdir string, pname string, dname string) {
	dir := path.Join(pdir, dname)
	err := os.Mkdir(dir, 0700)
	
	if err != nil {
		fmt.Printf("disco create: error creating project directory '%s': %s\n", dname, err)
		os.Exit(0)
	}
	
	fmt.Printf("    created %s/%s\n", pname, dname)
}

func createProjectFile(pdir string, pname string, fname string) {
	f := path.Join(pdir, fname)
	file, err := os.OpenFile(f, os.O_CREATE, 0700)

	defer file.Close()

	if err != nil {
		fmt.Printf("disco create: error creating project file '%s': %s\n", fname, err)
		os.Exit(0)
	}
	
	fmt.Printf("    created %s/%s\n", pname, fname)
}