// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package geometry
import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Point struct{ x, y float64 }

func (p Point) X () float64 {
	return p.x
}

func (p Point) Y () float64 {
	return p.y
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X() - p.X(), q.Y() - p.Y())
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X() - p.X(), q.Y() - p.Y())
}

func (p Point) checkSegment (q Point, r Point) bool{
	vec_1 := Point{q.X() - p.X(), q.Y() - p.Y()}
	vec_2 := Point{r.X() - p.X(), r.Y() - p.Y()}
	cross := vec_1.X() * vec_2.Y() - vec_2.X() * vec_1.Y()
	if cross < 0.1{
		return true
	}else{
		return false
	}
}

func doIntersect (path Path, lastIndex int) bool {
	itIntersect := false
	for i := lastIndex ; i >= 2; i--{
		p := path[i-2]
		q := path[i-1]
		r := path[i]
		if p.checkSegment(q, r) {
			itIntersect = true
			break
		}
		if q.checkSegment(p, r) {
			itIntersect = true
			break
		}
	}
	return itIntersect
}

func (p Point) onSegment (q Point, r Point) bool {
	if q.X() <= math.Max(p.X(), r.X()) && q.X() >= math.Min(p.X(), r.X()) &&
		q.Y() <= math.Max(p.Y(), r.Y()) && q.Y() >= math.Min(p.Y(), r.Y()){
		return true
	}
	return false
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (path Path) printVerticles () {
	fmt.Println("Figure's vertices")
	for _, pt := range path {
		fmt.Printf("(%f, %f)\n" ,pt.X(), pt.Y())
	}
}


func main (){
	var sides, err = strconv.Atoi(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Generating a [\"%s\"] sides figure \n", os.Args[1])

		path := make(Path, sides)
		for i := 0; i < sides; i++{
			x := -100 + rand.Float64() * (100)
			y := -100 + rand.Float64() * (100)
			path[i] = Point{x, y}
			if i >= 2 {
				if doIntersect(path, i){
					i--
				}
			}
		}
		path.printVerticles()
		fmt.Println("Figure's Perimeter")
		fmt.Print(path.Distance())
	}
}


