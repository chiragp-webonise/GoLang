package main

 import (
         "fmt"
         "io/ioutil"
         "os"
         "encoding/xml"
         "strings"
         "github.com/golang-geo/geo"
         _"database/sql"
         "strconv"
         )


 type kml struct {
        XMLName xml.Name `xml: "kml"`
        Document Document `xml:"Document"`
        
 }

 type Document struct {
        XMLName xml.Name `xml:"Document"`
        Folder[] Folder `xml:"Folder"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Folder struct {
        XMLName xml.Name `xml:"Folder"`
        //LineString LineString `xml:"LineString"`
        Placemark[] Placemark `xml:"Placemark"`
 }
 type Placemark struct {
        XMLName xml.Name `xml:"Placemark"`
        LineString LineString `xml:"LineString"`
        Point Point `xml:"Point"`
 }
type Point struct {
        XMLName xml.Name `xml:"Point"`
        Coordinates string `xml:"coordinates"`
 }
 type LineString struct {
        XMLName xml.Name `xml:"LineString"`
        Coordinates string `xml:"coordinates"`
 }
 // func (l LineString) String() string {
 //         return fmt.Sprintf(l.Coordinates)
 // }
func splitLink(s string) []string {

    s=strings.Replace(s, "0 ", "", -1)
    s=strings.TrimSpace(s)
    x := strings.Split(s,",")
    return x[:len(x)-1]
}

func FlushTestDB(s *geo.SQLMapper) {
    s.SqlConn.Exec("DELETE FROM points;")
}
func RoundFloat(x float64, prec int) float64 {
    frep := strconv.FormatFloat(x, 'g', prec, 64)
    f, _ := strconv.ParseFloat(frep, 64)
    return f
}   
func main() {

        // var la,ln float64
        var pm int64
        i:=0
        j:=0
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
         for i = 0; i < len(k.Document.Folder);  i++ {
            
            for j = 0; j < len(k.Document.Folder[i].Placemark); j++ {
                // k.Document.Folder[i].Placemark[j].Point.Coordinates   
               
                    pm++
            }
        }
         for i = 0; i < len(k.Document.Placemark);  i++ {
                pm++
            }
         // co2:=k.Document.Placemark[2].LineString.Coordinates
        fmt.Println(pm)
}
         