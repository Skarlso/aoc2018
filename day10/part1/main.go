package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

type star struct {
	x  int
	y  int
	vx int
	vy int
}

func (s *star) add() {
	s.x += s.vx
	s.y += s.vy
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	maxX := 0
	maxY := 0
	minX := 10000000
	minY := 10000000

	for _, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d,  %d> velocity=<%d,  %d>", &x, &y, &vx, &vy)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if abs(x) < minX {
			minX = x
		}
		if abs(y) < minY {
			minY = y
		}
	}
	fmt.Println(maxX, maxY, minX, minY)

	startX := maxX
	startY := maxY

	stars := make([]*star, 0)
	for _, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d,  %d> velocity=<%d,  %d>", &x, &y, &vx, &vy)
		s := star{
			x:  startX + x,
			y:  startY + y,
			vx: vx,
			vy: vy,
		}
		stars = append(stars, &s)
	}

	seconds := 0
	for {
		// sky := make(map[coord]bool)
		for _, s := range stars {
			s.add()
			if s.x > maxX*2 || s.y > maxY*2 || s.x < minX || s.y < minY {
				fmt.Println("reached the end")
				return
			}
			// sky[coord{x: s.x, y: s.y}] = true
		}
		filename := fmt.Sprint("name", seconds)
		saveImage(filename, minX, maxX, minY, maxY, stars)
		// fmt.Println(draw(minX, maxX, minY, maxY, stars))
		// for i := 0; i < maxY*2; i++ {
		// 	for j := 0; j < maxX*2; j++ {
		// 		if _, ok := sky[coord{x: j, y: i}]; ok {
		// 			fmt.Print("#")
		// 		} else {
		// 			fmt.Print(" ")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		seconds++
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func writeSvg(minx, maxx, miny, maxy int, coords []*star) {
	header := fmt.Sprintf(`<svg width="%s" height="%s" xmlns="http://www.w3.org/2000/svg">`, maxx, maxy)
	f, err := os.Create("sky.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(header)
	for _, c := range coords {
		img.Set(c.x, c.y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}
}

func saveImage(fileName string, minx, maxx, miny, maxy int, coords []*star) {
	imgRect := image.Rect(minx, miny, maxx*2, maxy*2)
	img := image.NewRGBA(imgRect)
	for _, c := range coords {
		img.Set(c.x, c.y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}
	f, err := os.Create(fmt.Sprintf("%s.png", fileName))
	if err != nil {
		log.Fatal(err)
	}
	enc := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	if err := enc.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
