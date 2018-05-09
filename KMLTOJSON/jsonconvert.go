package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)
type kml struct {
    
        Document Document `xml:"Document"`
 }
 type Document struct {
        
        Name string `xml:"name"`
        Open string `xml:"open"`
        Style[] Style `xml:"Style"`
        StyleMap[] StyleMap `xml:"StyleMap"`
        Folder[] Folder `xml:"Folder"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Style struct {
      
 		Id string `xml:"id,attr"`
        IconStyle IconStyle `xml:"IconStyle"`
        LineStyle LineStyle `xml:"LineStyle"`
 }
 type IconStyle struct {
     
        Scale string `xml:"scale"`
        Icon Icon `xml:"Icon"`
        HotSpot HotSpot `xml:"hotSpot"`
 }

 type HotSpot struct {
 	
 		X string `xml:"x,attr"`
 		Y string `xml:"y,attr"`
 		Xunits string `xml:"xunits,attr"`
 		Yunits string `xml:"yunits,attr"`
 }
 type Icon struct {
        
        Href string `xml:"href"`
 }
  type LineStyle struct {
     
        Color string `xml:"color"`
 }
 type Folder struct {
        
        Name string `xml:"name"`
        Description string `xml:"description"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Placemark struct {
       
        Name string `xml:"name"`
        Description string `xml:"description"`
        LookAt LookAt `xml:"LookAt"`
        StyleUrl string `xml:"styleUrl"`
        LineString LineString `xml:"LineString"`
        Point Point `xml:"Point"`
 }
 type LookAt struct {
  
        Longitude string `xml:"longitude"`
        Latitude string `xml:"latitude"`
        Altitude string `xml:"altitude"`
        Heading string `xml:"heading"`
        Tilt string `xml:"tilt"`
        Range string `xml:"range"`
        GxaltitudeMode string `xml:"http://www.google.com/kml/ext/2.2 altitudeMode"`
 }
type Point struct {
    
        GxdrawOrder string `xml:"http://www.google.com/kml/ext/2.2 drawOrder"`
        Coordinates string `xml:"coordinates"`
 }
 type LineString struct {
 
        Tessellate string `xml:"tessellate"`
        Coordinates string `xml:"coordinates"`
 }
 type StyleMap struct{
 	   
 		Id string `xml:"id,attr"`
 	    Pair[] Pair `xml:"Pair"`
 }
 type Pair struct{
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