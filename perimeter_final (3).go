package main

import (
    "fmt"
    "math"
)

type Shape interface {
    perimeter() float64
    
}

type Circle struct {
    radii float64
}

type Rectangle struct {
    length, width float64
}

func (c Circle) perimeter() float64 {
    return 2.0 * math.Pi * c.radii
}


func (r Rectangle) perimeter() float64 {
    return 2.0 * (r.width + r.length)
}


func finalCall(s Shape) {
    fmt.Printf("perimeter is %f ", s.perimeter())
}

func main() {
    finalCall(Rectangle{5.5, 20.3})
    finalCall(Circle{10})
}