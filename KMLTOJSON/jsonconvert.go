package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)
type kml struct {
        //XMLName xml.Name `xml: "kml"`
        Document Document `xml:"Document"`
 }
 type Document struct {
        //XMLName xml.Name `xml:"Document"`
        Name string `xml:"name"`
        Open string `xml:"open"`
        Style[] Style `xml:"Style"`
        StyleMap[] StyleMap `xml:"StyleMap"`
        Folder[] Folder `xml:"Folder"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Style struct {
        //XMLName xml.Name `xml:"Style"`
 		Id string `xml:"id,attr"`
        IconStyle IconStyle `xml:"IconStyle"`
        LineStyle LineStyle `xml:"LineStyle"`
 }
 type IconStyle struct {
        //XMLName xml.Name `xml:"IconStyle"`
        Scale string `xml:"scale"`
        Icon Icon `xml:"Icon"`
        HotSpot HotSpot `xml:"hotSpot"`
 }

 type HotSpot struct {
 		//XMLName xml.Name `xml:"hotSpot"`
 		X string `xml:"x,attr"`
 		Y string `xml:"y,attr"`
 		Xunits string `xml:"xunits,attr"`
 		Yunits string `xml:"yunits,attr"`
 }
 type Icon struct {
        //XMLName xml.Name `xml:"Icon"`
        Href string `xml:"href"`
 }
  type LineStyle struct {
        //XMLName xml.Name `xml:"LineStyle"`
        Color string `xml:"color"`
 }
 type Folder struct {
        //XMLName xml.Name `xml:"Folder"`
        //LineString LineString `xml:"LineString"`
        Name string `xml:"name"`
        Description string `xml:"description"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Placemark struct {
        //XMLName xml.Name `xml:"Placemark"`
        Name string `xml:"name"`
        Description string `xml:"description"`
        LookAt LookAt `xml:"LookAt"`
        StyleUrl string `xml:"styleUrl"`
        LineString LineString `xml:"LineString"`
        Point Point `xml:"Point"`
 }
 type LookAt struct {
        //XMLName xml.Name `xml:"LookAt"`
        Longitude string `xml:"longitude"`
        Latitude string `xml:"latitude"`
        Altitude string `xml:"altitude"`
        Heading string `xml:"heading"`
        Tilt string `xml:"tilt"`
        Range string `xml:"range"`
        GxaltitudeMode string `xml:"http://www.google.com/kml/ext/2.2 altitudeMode"`
 }
type Point struct {
        //XMLName xml.Name `xml:"Point"`
        GxdrawOrder string `xml:"http://www.google.com/kml/ext/2.2 drawOrder"`
        Coordinates string `xml:"coordinates"`
 }
 type LineString struct {
        //XMLName xml.Name `xml:"LineString"`
        Tessellate string `xml:"tessellate"`
        Coordinates string `xml:"coordinates"`
 }
 type StyleMap struct{
 	   // XMLName xml.Name `xml:"StyleMap"`
 		Id string `xml:"id,attr"`
 	    Pair[] Pair `xml:"Pair"`
 }
 type Pair struct{
 		//XMLName xml.Name `xml:"Pair"`
 		Key string `xml:"key"`
 		StyleUrl string `xml:"styleUrl"`
 }
func main() {
	if len(os.Args) < 2 {
            fmt.Println("Missing parameter, provide file name!")
            return
        }   
         xmlFile, err := os.Open(os.Args[1])
         if err != nil {
                 fmt.Println("Error opening file:", err)
                 return
         }

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var k kml
	xml.Unmarshal(xmlData, &k)



	jsonData, err := json.Marshal(k)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	// Write to JSON file
	jsonFile, err := os.Create("./TEST.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}