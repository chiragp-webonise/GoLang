package main

 import (
     "fmt"

     "github.com/umahmood/haversine"
 )

 func main() {
    oxford := haversine.Coord{Lat: 25.0895488165, Lon: 85.28285559459999}  // Oxford, UK
     turin  := haversine.Coord{Lat: 25.0935900198, Lon: 85.29507335629999}  // Turin, Italy
     mi, km := haversine.Distance(oxford, turin)
     fmt.Println("Miles:", mi, "Kilometers:", km)
 }
// package main

// import("github.com/golang-geo"
//        "fmt")

// func main() {
//      // Make a few points
//      p := geo.NewPoint(25.0895488165, 85.28285559459999)
//      p2 := geo.NewPoint(25.0935900198, 85.29507335629999)
     
//      // find the great circle distance between them
//      dist := p.GreatCircleDistance(p2)
//      fmt.Printf("great circle distance: %d\n", dist)
// }