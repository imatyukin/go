// Упражнение 1.12.
// Измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из URL. Например, URL вида
// http://localhost:8000/?cycles=20 устанавливает количество циклов равным 20 вместо значения по умолчанию, равного 5.
// Используйте функцию strconv.Atoi для преобразования строкового параметра в целое число. Просмотреть документацию по
// данной функции можно с помощью команды go doc strconv.Atoi.

package main

import (
	"log"
	"net/http"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"image/color"
	"strconv"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles := 5
	if cyclesStr := r.FormValue("cycles"); cyclesStr != "" {
		if value, err := strconv.Atoi(cyclesStr); err == nil {
			cycles = value
		}
	}
	lissajous(w, cycles)
}
//!-handler

//!+lissajous
// Lissajous generates GIF animations of random Lissajous figures.
func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
//!-lissajous

//!-main
