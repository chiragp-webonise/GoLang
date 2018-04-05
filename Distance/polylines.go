 package main

 import (
         "fmt"
         "io/ioutil"
         "os"
         "encoding/xml"
         "strings"
         "math"
         "strconv"
 )


 type kml struct {
        XMLName xml.Name `xml:"kml"`
        Document Document `xml:"Document"`
        
 }

 type Document struct {
        XMLName xml.Name `xml:"Document"`
        Placemark Placemark `xml:"Placemark"`
 }
 type Placemark struct {
        XMLName xml.Name `xml:"Placemark"`
        LineString LineString `xml:"LineString"`
 }
type LineString struct {
        XMLName xml.Name `xml:"LineString"`
        Coordinates string `xml:"coordinates"`
 }
 func (l LineString) String() string {
         return fmt.Sprintf(l.Coordinates)
 }

 func hsin(theta float64) float64 {
        return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, al1, lat2, lon2, al2 float64) float64 {
    
        var la1, lo1, la2, lo2, r float64
        la1 = lat1 * math.Pi / 180
        lo1 = lon1 * math.Pi / 180
        la2 = lat2 * math.Pi / 180
        lo2 = lon2 * math.Pi / 180
        r = 6378100 // Earth radius in METERS

        // calculate
        h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
        ald:=al1-al2

        hp:=math.Pow(h,2)

        aldp:=math.Pow(ald,2)

        D:=math.Sqrt(hp+aldp)

        m:=2 * r * math.Asin(math.Sqrt(D))

        km:=m/1000.0
        return km
}
func splitLink(s string) (string, string, string, string,string,string) {

    x := strings.Split(s,",")
    return x[0], x[1], x[2], x[3], x[4], x[5]
}
func main() {
         xmlFile, err := os.Open("polylines.kml")
         if err != nil {
                 fmt.Println("Error opening file:", err)
                 return
         }
         defer xmlFile.Close()

         XMLdata, _ := ioutil.ReadAll(xmlFile)

         var k kml
         xml.Unmarshal(XMLdata, &k)

         co:=k.Document.Placemark.LineString.Coordinates
         f1,f2,f3,f4,f5,f6:=splitLink(co)
         var la1,lo1,al1,la2,lo2,al2 float64
         la1, err = strconv.ParseFloat(f1, 64)
         lo1, err = strconv.ParseFloat(f2, 64)
         al1, err = strconv.ParseFloat(f3, 64)
         la2, err = strconv.ParseFloat(f4, 64)
         lo2, err = strconv.ParseFloat(f5, 64)
         al2, err = strconv.ParseFloat(f6, 64)
         d:=Distance(la1,lo1,al1,la2,lo2,al2)
         fmt.Println(d)
 }