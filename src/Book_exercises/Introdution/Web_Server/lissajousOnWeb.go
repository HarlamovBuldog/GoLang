//Lissajous generates animated GIF
//from random Lissajou's figures
//Task 1.12 realization from book page 45
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
	"os"
	"strconv"
	"time"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		regulator(r)
		lissajous(w)
	}
	http.HandleFunc("/lissajous", handler) //each request calls handler

	//next lines are same as previous two
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		})
	*/
	//http.HandleFunc("/", regulator)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

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

var cycles = 5.0 //Количество полных колебаний х
var res = 0.001  //Угловое разрешение
var size = 100   //Канва изображения охватывает [size..+size]
var nframes = 64 //Количество кадров анимации
var delay = 8    //Задержка между кадрами (единица - 10 мс)

func lissajous(out io.Writer) {
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
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), randomColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func regulator(r *http.Request) {

	r.ParseForm()
	keyValuePair := r.Form
	for key, value1 := range keyValuePair {
		if len(value1) < 1 {
			continue
		}
		fmt.Fprintf(os.Stdout, key+" = "+value1[0]+"\n")
	}

	for key, value1 := range keyValuePair {
		if len(value1) < 1 {
			continue
		}
		switch key {
		case "cycles":
			f64, err := strconv.ParseFloat(value1[0], 64)
			if err == nil && cycles != f64 {
				cycles = f64
				fmt.Fprintf(os.Stdout, "Success! Value of "+key+
					" changed to "+value1[0]+"\n")
			}
		case "size":
			i, err := strconv.ParseInt(value1[0], 10, 64)
			if err == nil && size != int(i) {
				size = int(i)
				fmt.Fprintf(os.Stdout, "Success! Value of "+key+
					" changed to "+value1[0]+"\n")
			}
		case "nframes":
			i, err := strconv.ParseInt(value1[0], 10, 64)
			if err == nil && nframes != int(i) {
				nframes = int(i)
				fmt.Fprintf(os.Stdout, "Success! Value of "+key+
					" changed to "+value1[0]+"\n")
			}
		default:
			fmt.Fprintf(os.Stdout, "Wrong key! "+key+"\n")
		}
	}
}
