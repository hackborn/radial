# radial

Generate a shaded circle. By default, it will create an image 128x128 with transparency outside the circle, and slight shading on the circle so it's white in the center and darker on the edges.

![Example](https://github.com/hackborn/radial/blob/master/readme.png)

## Compiling

* Download Go (https://golang.org/). I'm on 1.8.1.
* Install in the correct folder according to Go rules (which would be your Go folder, *src/github.com/hackborn*.
* Enter this folder.
* From a command line, type "go build".
* Run the resulting executable.

## Modifying
There are no command line or configuration options. If you want it to do something different, you'll need to modify the source. On the plus size, the source that does anything is only a couple lines long and should be fairly obvious.
