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
        Placemarks []Placemark `xml:"Placemark"`
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

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
    
        var la1, lo1, la2, lo2, r float64
        la1 = lat1 * math.Pi / 180
        lo1 = lon1 * math.Pi / 180
        la2 = lat2 * math.Pi / 180
        lo2 = lon2 * math.Pi / 180

        r = 6378100 // Earth radius in METERS

        // calculate
        h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

        m:=2 * r * math.Asin(math.Sqrt(h))

        //km:=m/1000.0
        return m
}
 func splitLink(s string) []string {

    s=strings.Replace(s, "0 ", "", -1)
    k:=strings.TrimSpace(s)
    x := strings.Split(k,",")
    return x[:len(x)-1]
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

         XMLdata, _ := ioutil.ReadAll(xmlFile)

         var k kml
         xml.Unmarshal(XMLdata, &k)

        co:=k.Document.Placemarks[1].LineString.Coordinates
        co2:=k.Document.Placemarks[2].LineString.Coordinates
        co3:=k.Document.Placemarks[3].LineString.Coordinates

         coArray:=splitLink(co)
         coArray2:=splitLink(co2)
         coArray3:=splitLink(co3)
         temp:=append(coArray2,coArray[38:]...)
         final:=append(temp,coArray3[32:]...)
         var la1,lo1,la2,lo2,d float64
         //var i,j int
         //i=1
         //j=i
         lo1, err = strconv.ParseFloat(final[0], 64)
         la1, err = strconv.ParseFloat(final[1], 64)
         lo2, err = strconv.ParseFloat(final[2], 64)
         la2, err = strconv.ParseFloat(final[3], 64) 
         //j=i+1
         d=Distance(la1,lo1,la2,lo2)
         total:=d
         //fmt.Println("point",i,"to",j,"distance is ",":",d,"meters")
         for c := 4; c <len(final); c=c+2 {
            
                lo1, err = strconv.ParseFloat(final[c], 64)
                la1, err = strconv.ParseFloat(final[c+1], 64)
                d=Distance(la2,lo2,la1,lo1)
                total=total+d
                la2=la1
                lo2=lo1
                     //i++
                     //j=i+1
                //fmt.Println("point",i,"to",j,"distance is ",":",d,"meters")
        }
        km:=total/1000.0
                fmt.Println("total distance from A4 to A8 is",km,"KMS")
 }