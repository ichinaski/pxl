# pxl

pxl is a little hack to display images in the terminal.

## Installation

You will need to have [Go](https://golang.org) installed and configured in your path.
Your terminal must have xterm-256color mode enabled.

### For Go1.16 or later:

`go install github.com/ichinaski/pxl/cmd/pxl@latest`

will fetch, build, and install the lastest version to wherever Go normally installs binaries
(see [How to Write Go Code](https://golang.org/doc/code#Command) for details).
If you wish to install somewhere else, you could do:

`env GOBIN=/tmp go install github.com/ichinaski/pxl/cmd/pxl@latest`

Where GOBIN is set to some directory where you want it installed.

You could also specify a specific version other than `@latest` if desired.

#### For Go1.15 or lower:

`go get github.com/ichinaski/pxl/cmd/pxl`

## Usage

`pxl filename …`

Move to the next image or exit with `<ESC>` or `q`.

## Disclaimer

You may want to squint your eyes or take a few steps backwards when looking at the output.

## Package

There is also a simple Go package available.
Note, requires clients use the [`github.com/nsf/termbox-go`](https://pkg.go.dev/github.com/nsf/termbox-go) package.

[![Go Reference](https://pkg.go.dev/badge/github.com/ichinaski/pxl.svg)](https://pkg.go.dev/github.com/ichinaski/pxl)
Online package documentation is available via
[pkg.go.dev](https://pkg.go.dev/github.com/ichinaski/pxl).

## Examples


|  image  | pxl   |
|:--:|:--:|
|<img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/gh.png" height="250"> | <img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/gh.pxl.png" height="250"> |
|<img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/batman.jpg" height="250"> | <img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/batman.pxl.png" height="250"> |
|<img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/elvis.jpg" height="250"> | <img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/elvis.pxl.png" height="250"> |
|<img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/gd.jpg" height="250"> | <img src="https://raw.githubusercontent.com/ichinaski/pxl/master/img/gd.pxl.png" height="250"> |

## License
This software is distributed under the BSD-style license found in the LICENSE file.
