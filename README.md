# image-resizer
Commandline tool for resizing images. Exercise from the [Gophers Aachen](https://github.com/gophersaachen) workshop.



## usage

Read all jpg files from the input directory, resize them with the given width and height and write the resized images to the output directory.
```
go run main.go input-dir output-dir width height
```

Sample usage
```
go run main.go "C:\Users\philipilihp\images" "C:\Users\philipilihp\resized-images" "800" "600"
```