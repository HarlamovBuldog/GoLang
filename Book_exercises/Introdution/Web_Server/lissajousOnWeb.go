//Lissajous generates animated GIF
//from random Lissajou's figures

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 //1st palette color
	blackIndex = 1 //Next palette color
)

func main() {

	//lissajous(os.Stdout)
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler) //each request calls handler
	http.HandleFunc("/?", regulator)

	//next lines are same as previous
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	http.HandleFunc("/?", regulator)
	log.Fatal(http.ListenAndServe(":8000", nil))
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
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func regulator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count = %d\n")
}
