package main 
import (
	"github.com/jasonwinn/geocoder"
	"fmt"
	)
 

func main() {

	query := "Piplod,Surat, India"
  lat, lng, err := geocoder.Geocode(query)
  if err != nil {
    panic("THERE WAS SOME ERROR!!!!!")
  }

  	fmt.Println(lat,lng)
  	address, err := geocoder.ReverseGeocode(lat,lng)
  if err != nil {
    panic("THERE WAS SOME ERROR!!!!!")
  }

  fmt.Println(address.Street) 	        // 542 Marion St   
  fmt.Println(address.City) 		        // Seattle
  fmt.Println(address.State) 	        // WA
  fmt.Println(address.PostalCode) 	    // 98104 
  fmt.Println(address.County) 	        // King
  fmt.Println(address.CountryCode)       // US 

  directions := geocoder.NewDirections(/*"Sohjana,Bihar"*/"Laranpur Hospital Road, India", []string{"Sunri, India"})
  results, err := directions.Get()
  if err != nil {
    panic("THERE WAS SOME ERROR!!!!!")
  }

  route := results.Route
  //time := route.Time
  //legs := route.Legs
  distance := route.Distance

  // or get distance with this shortcut
  //
  // use "k" to return result in km
  // use "m" to return result in miles
  distance, err = directions.Distance("k")
  if err != nil {
    panic("THERE WAS SOME ERROR!!!!!")
  }
  fmt.Println(distance)       // US 
}
 
