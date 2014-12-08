package heatmap

import (
	"github.com/dustin/go-heatmap"
	"github.com/dustin/go-heatmap/schemes"
	"image"
	"image/color"
	"image/png"
	"os"
)

type HeatmapPoint struct {
	X float64
	Y float64
}

type Heatmap struct {
	scheme []color.Color
}

func NewHeatmap(scheme_path string) Heatmap {
	var heatmap Heatmap
	heatmap.scheme, _ = schemes.FromImage(scheme_path)
	return heatmap
}

func (this *Heatmap) DrawHeatmap(points []HeatmapPoint, width int, height int, img_path string, dotwith int, dotheight uint8) {
	pointlist := []heatmap.DataPoint{}
	for _, value := range points {
		pointlist = append(pointlist, heatmap.P(value.X-float64(dotwith)/float64(2), value.Y-float64(dotheight)/float64(2)))
	}

	imgfile, _ := os.Create(img_path)
	defer imgfile.Close()
	img := heatmap.Heatmap(image.Rect(0, 0, width, height),
		pointlist, dotwith, dotheight, this.scheme)
	png.Encode(imgfile, img)
}
