# winch

![Build](https://github.com/winchci/winch/workflows/Build/badge.svg)

Winch is a universal build and release tool.

## Configuration File

Winch reads from a `winch.yml` in the current directory (optionally from a specified using the `-f` option on the command line).

```yaml
name: "string"                          # the name of the project
description: "string"                   # the description of the project
repository: "string"                    # URL to the GitHub repository
local: false                            # set to true to use only the local repository
verbose: false                          # set to true to get verbose output
quiet: false                            # set to true to get quiet output
language: "go"                          # "go" and "node" are the only supported languages currently
toolchain: "string"                     # node projects can choose "npm" or "yarn"
ci:
  enabled: true                         # set to false to disable CI
  branches:
    ignore: "string/regexp"             # execute CI unless the current branch matches the given pattern
    only: "string/regexp"               # execute CI only if the current branch matches the given pattern
  tags:
    ignore: "string/regexp"             # execute CI unless the current tag matches the given pattern
    only: "string/regexp"               # execute CI unless the current tag matches the given pattern
before_install:                         # executed before installation (see commands/"string")
install:                                # executed to install (see commands/"string")
after_install:                          # executed after installation (see commands/"string")
before_build:                           # executed before building (see commands/"string")
build:                                  # executed to build (see commands/"string")
after_build:                            # executed after building (see commands/"string")
before_test:                            # executed before testing (see commands/"string")
test:                                   # executed to test (see commands/"string")
after_test:                             # executed after testing (see commands/"string")
assets:
- enabled: true                         # set to false to disable asset generation
  filename: "string"                    # the filename to generate
  directory: "string"                   # the directory containing the assets
  package: "string"                     # the Go package containing the assets
  variable: "string"                    # the Assets variable to generate
  tag: "string"                         # the build tag to add to the file
  only:
    - "string"                          # only include files matching the given pattern
  except:
    - "string"                          # only include files except those matching the given pattern
  branches:
    ignore: "string/regexp"             # generate assets unless the current branch matches the given pattern
    only: "string/regexp"               # generate assets only if the current branch matches the given pattern
  tags:
    ignore: "string/regexp"             # generate assets unless the current tag matches the given pattern
    only: "string/regexp"               # generate assets unless the current tag matches the given pattern
changelog:
  enabled: true                         # set to false to disable changelog generation
  template: "string"                    # name of a file containing the CHANGELOG template
  file: "string"                        # name of the output file
version:
  enabled: true                         # set to false to disable version generation
  template: "string"                    # name of a file containing the version file template
  file: "string"                        # name of the output file
transom:
  enabled: true                         # set to false to disable transom publication
  server: "string"                      # hostname of the transom server (defaults to transom.b10s.io)
  application: "string"                 # application name defined in transom
  token: "string"                       # authentication token to use
  username: "string"                    # username if token not provided
  password: "string"                    # password if token not provided
  directory: "string"                   # the directory containing assets
  artifacts:
  - "string"                            # artifacts to upload to the Transom bundle
  - "string"
  branches:
    ignore: "string/regexp"             # publish to Transom unless the current branch matches the given pattern
    only: "string/regexp"               # publish to Transom only if the current branch matches the given pattern
  tags:
    ignore: "string/regexp"             # publish to Transom unless the current tag matches the given pattern
    only: "string/regexp"               # publish to Transom unless the current tag matches the given pattern
dockerfile:
  enabled: true                         # set to false to disable Dockerfile generation
  template: "string"                    # name of a file containing the Dockerfile template
  file: "string"                        # name of the output file
docker:
  enabled: true                         # set to false to disable docker image pushing
  organization: "string"                # DockerHub organization
  repository: "string"                  # DockerHub repository
  username: "string"                    # username to login to DockerHub
  password: "string"                    # password to login to DockerHub
  dockerfile: "string"                  # Dockerfile to process (defaults to "Dockerfile")
  context: "string"                     # context directory (defaults to ".")
  tag: "string"                         # tag to use in addition to version (defaults to "latest")
  buildargs:
    "string": value                     # build args to add to the build
  branches:
    ignore: "string/regexp"             # publish to DockerHub unless the current branch matches the given pattern
    only: "string/regexp"               # publish to DockerHub only if the current branch matches the given pattern
  tags:
    ignore: "string/regexp"             # publish to DockerHub unless the current tag matches the given pattern
    only: "string/regexp"               # publish to DockerHub unless the current tag matches the given pattern
dockers:
- @docker                               # same as the docker configuration above
release:
  enabled: true                         # set to false to disable creating GitHub releases
  artifacts:
  - "string"                            # artifacts to add to the GitHub release
  - "string"
  branches:
    ignore: "string/regexp"             # create release unless the current branch matches the given pattern
    only: "string/regexp"               # create release only if the current branch matches the given pattern
  tags:
    ignore: "string/regexp"             # create release unless the current tag matches the given pattern
    only: "string/regexp"               # create release unless the current tag matches the given pattern
environment:
  "string": value                       # environment variables to use for all commands
```

## Global options

```
Flags:
  -h, --help   help for build

Global Flags:
  -f, --file string   configuration file (default "winch.yml")
  -q, --quiet         quiet output
  -v, --verbose       verbose output
```

## Build

Execute the build command sequence.

```
$ winch build --help
Build

Usage:
  winch build [flags]
```

## CI

Execute a CI build command sequence:

1. Create version file
2. Create CHANGELOG file
3. Install
4. Build
5. Test
6. Release
7. Publish to Transom
8. Publish to Docker

```
# winch ci --help
Execute a CI build

Usage:
  winch ci [flags]
```

## Docker Build

Build a Docker image.

```
$ winch docker build --help
Build docker image

Usage:
  winch docker build [flags]
```

## Docker Publish

Push an image to DockerHub.

```
$ winch docker publish --help
Publish container to DockerHub

Usage:
  winch docker publish [flags]
```

## Generate Changelog

Generate a CHANGELOG file.

```
$ winch generate changelog --help
Generate a changelog

Usage:
  winch generate changelog [flags]

Flags:
  -o, --output string   output file
```

## Generate Version

Generate a version file (or update version information).

```
$ winch generate version --help
Generate the version file

Usage:
  winch generate version [flags]

Flags:
  -o, --output string   output file
```

## Init

Initialize the current directory to use winch, creating a `winch.yml`.

```
$ winch init --help
Initialize a configuration file

Usage:
  winch init [flags]
```

## Install

Run the install command sequence.

```
$ winch install --help
Install

Usage:
  winch install [flags]
```

## Name

Generate a new release name.

```
$ winch name --help
Generate a release name

Usage:
  winch name [flags]
```

## Release

Create a GitHub release and upload artifacts. The GitHub release description includes a description of the changes from the previous release.

```
$ winch release --help
Release changes

Usage:
  winch release [flags]
```

## Run

Run the specified command.

```
$ winch run --help
Run a command

Usage:
  winch run COMMAND [flags]
```

## Test

Run the test command sequence.

```
$ winch test --help
Test

Usage:
  winch test [flags]
```

## Transom Publish

Publish the artifacts to transom.

```
$ winch transom publish --help
Publish artifacts to Transom

Usage:
  winch transom publish [flags]
```

## Version

Output the version.

```
$ winch version --help
Show the version

Usage:
  winch version [flags]
```

## LICENSE

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
