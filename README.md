# pxl

pxl is a little hack to display images in the terminal.

### Installation

You will need to have [Go](https://golang.org) installed and configured in your path. Your terminal must have xterm-256color mode enabled.

`go get github.com/ichinaski/pxl`

### Usage

`pxl [-r whratio] filename`

where the optional flag `-r` sets the cursor's width/height ratio. This flag has a default value of 2.35, make sure you update it to match your terminal environment.

### Disclaimer

You may want to squint your eyes or take a few steps backwards when looking at the output.

### Examples

|  image  | pxl   |
|:--:|:--:|
|![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/gh.png =x250) | ![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/gh.pxl.png =x250) |
|
|![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/batman.jpg =250x) | ![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/batman.pxl.png =250x) |
|
|![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/elvis.jpg =250x) | ![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/elvis.pxl.png =250x) |
|
|![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/gd.jpg =250x) | ![image](file:////Users/ichinaski/projects/go/src/github.com/ichinaski/pxl/img/gd.pxl.png =250x) |


