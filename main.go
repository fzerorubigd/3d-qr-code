package main

import (
	"fmt"
	"log"

	"github.com/ogier/pflag"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	var (
		content            string
		depth, box, margin int
	)
	pflag.StringVarP(&content, "content", "c", "", "Content of the qr-code")
	pflag.IntVarP(&depth, "depth", "d", 7, "Depth of the layer, in millimeter")
	pflag.IntVarP(&box, "box", "b", 3, "Width of each small square in the qr-code")
	pflag.IntVarP(&margin, "margin", "m", 0, "Extra margin to add")

	pflag.Parse()

	qq, err := qrcode.New(content, qrcode.High)
	if err != nil {
		log.Fatal(err)
	}

	bmp := qq.Bitmap()
	fmt.Printf(`
depth = %d;
box = %d;
size = %d;
margin = %d;

width = size * box + margin;
module dot(x,y,depth, margin) {
	translate([x*box+margin,y*box+margin,depth]) {
		cube([box,box, depth]);
	}
}

cube([width, width,depth]);	
`, depth, box, len(bmp), margin)
	for i := range bmp {
		for j := range bmp[i] {
			if bmp[i][j] {
				fmt.Printf("dot(%d,%d,depth,margin);\n", i, j)
			}
		}
	}
}
