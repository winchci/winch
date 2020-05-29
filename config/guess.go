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
