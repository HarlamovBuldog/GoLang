//Lissajous generates animated GIF
//from random Lissajou's figures
//Note: go run lissajous.go > out.gif to make .gif file
//Task 1.5 and 1.6 realization from book page 37
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 1}, color.White,
	color.RGBA{255, 0, 0, 1}, color.RGBA{0, 0, 255, 1}, color.RGBA{255, 255, 0, 1}}

const (
	blackIndex  = 0 //1st palette color
	greenIndex  = 1 //Next palette color
	whiteIndex  = 2
	redIndex    = 3
	blueIndex   = 4
	yellowIndex = 5
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //Количество полных колебаний х
		res     = 0.001 //Угловое разрешение
		size    = 100   //Канва изображения охватывает [size..+size]
		nframes = 64    //Количество кадров анимации
		delay   = 8     //Задержка между кадрами (единица - 10 мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 //Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		randomColor := uint8(rand.Intn(5) + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
