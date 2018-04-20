package main

import (
	"fmt"

	"github.com/kelvins/geocoder"
)

func main() {
	// This example should work without need of an API Key,
	// but when you will use the API in your project you need
	// to set your API KEY as explained here:
	// https://developers.google.com/maps/documentation/geocoding/get-api-key
	// geocoder.ApiKey = "YOUR_API_KEY"

	// See all Address fields in the documentation
	address := geocoder.Address{
		Street:  "Central Park West",
		Number:  115,
		City:    "New York",
		State:   "New York",
		Country: "United States",
	}

	// Convert address to location (latitude, longitude)
	location, err := geocoder.Geocoding(address)

	if err != nil {
		fmt.Println("Could not get the location: ", err)
	} else {
		fmt.Println("Latitude: ", location.Latitude)
		fmt.Println("Longitude: ", location.Longitude)
	}

	// Set the latitude and longitude
	location = geocoder.Location{
		Latitude:  /*40.775807,*/25.0935900198,
		Longitude: /*-73.97632,*/85.29507335629999,
	}

	// Convert location (latitude, longitude) to a slice of addresses
	addresses, err := geocoder.GeocodingReverse(location)

	if err != nil {
		fmt.Println("Could not get the addresses: ", err)
	} else {
		// Usually, the first address returned from the API
		// is more detailed, so let's work with it
		address = addresses[0]

		// Print the address formatted by the geocoder package
		fmt.Println(address.FormatAddress())
		// Print the formatted address from the API
		fmt.Println(address.FormattedAddress)
		// Print the type of the address
		fmt.Println(address.Types)
	}
}