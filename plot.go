package main

import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
)

type Plot struct {
	plot *plot.Plot
}

func (pl *Plot) initialize() {
	p, err := plot.New()
	if err != nil {
	}

	pl.plot = p

	pl.plot.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{0, "0"}, {25, ""}, {50, "50"}, {75, ""}, {100, "100"},
	})
	pl.plot.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{0, "0"}, {25, ""}, {50, "50"}, {75, ""}, {100, "100"},
	})
}

func (pl *Plot) drawNode(nodes map[int]Node, result []int) {
	pts := make(plotter.XYs, len(nodes))

	for _, n := range nodes {
		for i, r := range result {
			if r == n.n_id {
				pts[i].X = n.x
				pts[i].Y = n.y
			} else {
			}
		}
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
	}

	pl.plot.Add(line, scatter)
}

func (pl *Plot) export(fileName string) {
	err := pl.plot.Save(1000, 1000, fileName+".png")
	if err != nil {
	}

	/*
		img := ImageRead(fileName)
		Formatgif(img, fileName)
	*/
}

/*
func ImageRead(ImageFile string) (image image.Image) {
	file, err := os.Open(ImageFile + ".png")
	if err != nil {
	}

	img, err := png.Decode(file)
	if err != nil {
	}
	file.Close()

	return img
}

func Formatgif(img image.Image, fileName string) {
	out, err := os.Create(fileName + ".gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt gif.Options
	opt.NumColors = 256
	err = gif.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (pl *Plot) CreateGIF(prefix string, files []string) {
	outGif := &gif.GIF{}
	for _, name := range files {
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	// out.gif に保存する
	f, _ := os.OpenFile(prefix+"out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
*/
