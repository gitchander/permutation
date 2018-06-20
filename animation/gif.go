package main

import (
	"bytes"
	"encoding/hex"
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"log"
	"math"

	//prmt "github.com/gitchander/permutation"
)

func makeGIF(filename string, render func(*gif.GIF) error) error {
	anim := new(gif.GIF)
	err := render(anim)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = gif.EncodeAll(&buf, anim)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buf.Bytes(), 0666)
	return err
}

func testGIF() {

	err := makeGIF("result.gif", render)
	checkError(err)
}

var sortPalette = []color.Color{
	rgb(0xFF, 0xFF, 0xFF), // background
	rgb(0xFF, 0xFF, 0xFF), // digits
	//	colorHex("#3c582f"),
	//	colorHex("#545284"),
	//	colorHex("#2e5a7f"),
	//	colorHex("#355752"),
	//	colorHex("#5f2535"),
	//	colorHex("#5a472d"),
	//	colorHex("#8b475a"),
	//	colorHex("#775073"),
	//	colorHex("#7c4835"),
	//	colorHex("#473156"),

	colorHex("#7a80cb"),
	colorHex("#50a77b"),
	colorHex("#908542"),
	colorHex("#ca703e"),
	colorHex("#c9687d"),
	colorHex("#cd9f36"),
	colorHex("#d157a8"),
	colorHex("#a54ddc"),
	colorHex("#6daf39"),
	colorHex("#dc3e41"),
}

func rgb(r, g, b byte) color.Color {
	return color.RGBA{r, g, b, 0xFF}
}

func colorHex(s string) color.Color {

	bs, err := hex.DecodeString(s[1:])
	checkError(err)

	return color.RGBA{bs[0], bs[1], bs[2], 0xFF}
}

func render(anim *gif.GIF) error {

	imageSize := image.Point{X: 256, Y: 256}

	//as := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	as := []int{1, 2, 3, 4}
	bs := make([]int, len(as))

	delay := 50
	pal := sortPalette

	//p := prmt.New(prmt.IntSlice(as))

	first := true
	for i := 0; i < 1000; i++ {

		if first {
			first = false
		} else {

			n := flipSize(bs)
			if n > len(as) {

				n = len(as)

				drawFlipElements(anim, imageSize, pal, as, n)
				flipInts(as, n)

				break
			}

			drawFlipElements(anim, imageSize, pal, as, n)
			flipInts(as, n)
		}

		r := image.Rectangle{Max: imageSize}
		ip := image.NewPaletted(r, pal)
		renderPal(ip, as)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, ip)
	}

	return nil
}

func flipSize(b []int) int {
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			return i + 2
		}
		b[i] = 0
	}
	return len(b) + 1
}

func flipInts(a []int, n int) {
	i, j := 0, n-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}

func drawElement(ip *image.Paletted, val int, pos image.Point, width int) {
	ssb, ok := mdig[val]
	if !ok {
		return
	}
	for y, sb := range ssb {
		for x, b := range sb {
			resIndex := uint8(val + 2)
			if b {
				resIndex = 1
			}

			p := pos.Add(image.Pt(x, y).Mul(width))

			drawUnit(ip, p, width, resIndex)
		}
	}
}

func drawUnit(ip *image.Paletted, pos image.Point, width int, index uint8) {
	var (
		x0 = pos.X
		x1 = x0 + width

		y0 = pos.Y
		y1 = y0 + width
	)
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			ip.SetColorIndex(x, y, index)
		}
	}
}

type Point2f struct {
	X, Y float64
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func renderPal(ip *image.Paletted, as []int) {

	bs := ip.Bounds()

	w := unitWidth * elementSize

	p0 := image.Point{
		X: (bs.Min.X + bs.Max.X) - w - ((len(as) - 1) * elementsDistatce),
		Y: (bs.Min.Y + bs.Max.Y) - w,
	}.Div(2)

	for i := range as {

		p := p0.Add(
			image.Point{
				X: elementsDistatce * i,
				Y: 0,
			})

		drawElement(ip, as[i], p, unitWidth)
	}
}

func renderAnimation(ip *image.Paletted, as []int, n int, t float64) {

	bs := ip.Bounds()

	w := unitWidth * elementSize

	p0 := image.Point{
		X: (bs.Min.X + bs.Max.X) - w - ((len(as) - 1) * elementsDistatce),
		Y: (bs.Min.Y + bs.Max.Y) - w,
	}.Div(2)

	p1 := p0.Add(
		image.Point{
			X: elementsDistatce * (n - 1),
			Y: 0,
		})

	var (
		x0 = float64(p0.X+p1.X) / 2
		y0 = float64(p0.Y+p1.Y) / 2
	)

	for i := 0; i < n; i++ {

		p := p0.Add(
			image.Point{
				X: elementsDistatce * i,
				Y: 0,
			})

		x1 := float64(p.X)
		y1 := float64(p.Y)

		x, y := rotate(x0, y0, x1, y1, t*math.Pi)
		pp := image.Point{
			X: round(x),
			Y: round(y),
		}

		drawElement(ip, as[i], pp, unitWidth)
		p.X += elementsDistatce
	}

	for i := n; i < len(as); i++ {
		p := p0.Add(
			image.Point{
				X: elementsDistatce * i,
				Y: 0,
			})

		drawElement(ip, as[i], p, unitWidth)
	}
}

// x0, y0 - center
func rotate(x0, y0, x1, y1 float64, angle float64) (x float64, y float64) {

	sin, cos := math.Sincos(angle)

	x = x0 + (x1-x0)*cos - (y1-y0)*sin
	y = y0 + (y1-y0)*cos + (x1-x0)*sin

	return
}

func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func drawFlipElements(anim *gif.GIF, imageSize image.Point,
	pal []color.Color, as []int, n int) {

	delay := 4

	t := 0.0
	for t < 1 {

		r := image.Rectangle{Max: imageSize}
		ip := image.NewPaletted(r, pal)

		renderAnimation(ip, as, n, t)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, ip)

		t += 0.1
	}
}
