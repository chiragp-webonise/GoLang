package main 

import (
		"encoding/csv"
		"fmt"
		"flag"
		"os"
		"strings"
	)

func main(){

	csvFile:= flag.String("csv","file.csv","comma seperated file")

	flag.Parse()

	file,err :=os.Open(*csvFile)

	
	if err != nil{
		exit(fmt.Sprintf("Failed to open the CSV file:%s\n",*csvFile))
	}

	r:= csv.NewReader(file)

	lines, err:= r.ReadAll()
	if err!= nil {
		exit("Failed to parse the provided csv file")
	}
	problems:=parseLines(lines)
	fmt.Println(problems)
	correct:=0
	for i, p:= range problems{
		fmt.Printf("Problem #%d: %s = \n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n",correct,len(problems))
}
func parseLines(lines [][]string) []problem{
	ret := make([]problem,len(lines))
	for i,line := range lines{
		ret[i] = problem {
			q: line[0],
			a:strings.TrimSpace(line[1]),
		}
	}
	return ret
}
type problem struct {
	q string
	a string
}
func exit(msg string) {
		fmt.Println(msg)
		os.Exit(1)
	}