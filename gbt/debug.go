package gbt

import (
	"fmt"
	"reflect"
)

// Debug print crap out
func Debug(p Renderer) {
	d(p, 0)
}

func d(p Renderer, depth int) {
	pd(depth)
	typeof := reflect.TypeOf(p).String()

	switch typeof {
	case "*gbt.Panel":
		fmt.Print("| ", "Panel", p.(*Panel).ID)
		fmt.Println("  html:", p.(*Panel).HTML)
		pd(depth)
		fmt.Println("  style:", p.(*Panel).Styles)
		for _, i := range p.(*Panel).Items {
			d(i, depth+1)
		}
	case "*gbt.Innerhtml":
		fmt.Print("| ", "Innerhtml", p.(*Innerhtml).ID)
		fmt.Println("  html:", p.(*Innerhtml).HTML)
	case "*gbt.Layout":
		fmt.Print("| ", "Layout", p.(*Layout).ID)
		fmt.Println(":", p.(*Layout).Type)
		for _, i := range p.(*Layout).Items {
			d(i, depth+1)
		}
	// case "*gbt.Body":
	// 	fmt.Print("| ", "Body", p.(*DivContainer).ID)
	// 	fmt.Println("")
	// 	for _, i := range p.(*DivContainer).Items {
	// 		d(i, depth+1)
	// 	}
	case "*gbt.Header":
		fmt.Print("| ", "Header", p.(*Header).ID)
		fmt.Println("::", p.(*Header).Title)
	default:
		fmt.Print("| ?", typeof)
		fmt.Println(" ?")
	}
}

// Print depth spaces
func pd(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
}
