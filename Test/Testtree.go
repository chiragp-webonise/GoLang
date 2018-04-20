package main 

import (
		"fmt"
		"github.com/dhconnelly/rtreego"
	)
	type Point []float64
	type Rect struct {
		p, q Point // Enforced by NewRect: p[i] <= q[i] for all i.
	}	
	type Spatial interface {
 		 Bounds() *Rect
	}
	var tol = 0.01

	type Somewhere struct {
	  location rtreego.Point
	  name string
	  wormhole chan int
	}

	func (s *Somewhere) Bounds() *rtreego.Rect {
	  // define the bounds of s to be a rectangle centered at s.location
	  // with side lengths 2 * tol:
	  return s.location.ToRect(tol)
	}
	// type Thing struct {
 //  		where *rtreego.Rect
 //  		name string
	// }	
	// func (t *Thing) Bounds() *rtreego.Rect {
 //  		return t.where
	// }
	func main() {

		rt := rtreego.NewTree(2, 25, 50)

		// p1 := rtreego.Point{85.28285559459999,25.0895488165}
		// p2 := rtreego.Point{85.27981097430001,25.0937297824}

		// r1, _ := rtreego.NewRect(p1, []float64{1, 2})
		// r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})



		// rt.Insert(&Thing{r1, "foo"})
		// rt.Insert(&Thing{r2, "bar"})
		
		rt.Insert(&Somewhere{rtreego.Point{85.28284799367495,25.08959000813778}, "Someplace2", nil})
		rt.Insert(&Somewhere{rtreego.Point{85.27981097430001,25.0937297824}, "Someplace", nil})
		size := rt.Size()
		q := rtreego.Point{85.28285559459999,25.0895488165}
		//k := 1

		// Get a slice of the k objects in rt closest to q:
		results:= rt.NearestNeighbors(3, q)
		fmt.Println(results[1].Bounds())
		fmt.Println(size)
	}