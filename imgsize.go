// get image width/height size
package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// get image width/height value
func getSize(filePath string) (int, int) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	bounds := img.Bounds()

	return bounds.Max.X, bounds.Max.Y
}

func main() {
	var isVerbose bool

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
  %s [OPTIONS] FILE

Options
  -h: Show this message
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&isVerbose, "v", false, "output with file name")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, filePath := range flag.Args() {
		var (
			w int
			h int
			f string = ""
		)
		w, h = getSize(filePath)

		if isVerbose {
			f = fmt.Sprintf(",%s", filePath)
		}

		fmt.Printf("%d,%d%s\n", w, h, f)
	}
}
