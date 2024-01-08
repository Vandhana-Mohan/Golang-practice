package main

import "fmt"

type Point struct {
	x int
	y int
}

type Circle struct {
	Point
	radius int
}

type Cake struct {
	typeOfCake string
	weight int // Weight in grams
}

type Triangle struct {
	height float32
	base float32
}

func (triangle Triangle) area() float32{
    return 0.5 * (triangle.base * triangle.height)
}

func (triangle *Triangle) updateBase(newBase float32){
  triangle.base = newBase
}

func main() {
	circle := Circle{Point{4,5}, 2}
	
  fmt.Println(circle.x)
  fmt.Println(circle)

	
  cakes := []Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}}
	
	fmt.Println(cakes[0].weight)
	
	cakes[1].weight = 1500
	fmt.Println(cakes)

	triangle := Triangle{10, 4}
  fmt.Println(triangle)

	fmt.Println(triangle.area())


  fmt.Println(triangle)

  triangle.updateBase(13)
  fmt.Println(triangle)
}
