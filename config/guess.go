/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package config

import "os"

func GuessLanguageAndToolchain() (string, string) {
	if fileExists("package.json") {
		return "node", guessNodeToolchain()
	}
	if fileExists("go.mod") {
		return "go", ""
	}
	if fileExists("pom.xml") {
		return "java", "mvn"
	}
	if fileExists("build.sbt") {
		return "scala", "sbt"
	}
	if fileExists("setup.py") {
		return "python", ""
	}
	if fileExists("Dockerfile") {
		return "docker", ""
	}
	return "go", ""
}

func guessNodeToolchain() string {
	if fileExists(".yarnrc") || fileExists("yarn.lock") {
		return "yarn"
	}

	return "npm"
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
