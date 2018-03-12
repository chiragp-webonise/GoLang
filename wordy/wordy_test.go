package wordy

import (
		"testing"
		"strconv"
		"regexp"
		"strings"
		)

func TestAnswer(t *testing.T) {
	for _, test := range tests {
		switch answer, ok := Answer(test.question,test.ok,test.answer); {
		case !ok:
			if test.ok {
				t.Fatalf("FAIL: %s\nAnswer(%q)\nreturned ok = false, expecting true.", test.description, test.question)
			}
		case !test.ok:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, %t, expecting ok = false.", test.description, test.question, answer, ok)
		case answer != test.answer:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, expected %d.", test.description, test.question, answer, test.answer)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func Answer(que string,check bool,ans int) (int,bool){
	
	re := regexp.MustCompile("[0-9]+")
	arr:=re.FindAllString(que, -1)
	val1,_:=strconv.Atoi(arr[0])
	val2,_:=strconv.Atoi(arr[1])
	result:=0	

		if len(arr)==2{

		
			
			if strings.Contains(que, "minus"){
				result=val1 - val2
			}else if strings.Contains(que, "multiplied"){
				result=val1 * val2			
			}else if strings.Contains(que, "divided"){
				result=val1 / val2
			}else if strings.Contains(que, "plus") {
				result=val1 +val2
			}else {
				if !check{
				return 0,false
				}
			}

			if result==ans && check==true{
				return result,true 
			}
		}

		if len(arr)==3{

			
			val3,_:=strconv.Atoi(arr[2])
			result:=0
			if strings.Contains(que, "multiplied "){
				que = strings.Replace(que,"multiplied","",1)
				result=val1 * val2			
			}else if strings.Contains(que, "minus"){
				que = strings.Replace(que,"minus","",1)
				result=val1 - val2
			}else if strings.Contains(que, "divided"){
				que = strings.Replace(que,"divided","",1)
				result=val1 / val2
			}else if strings.Contains(que, "plus") {
				que = strings.Replace(que,"plus","",1)
				result=val1 + val2
			}else {
				if !check{
				return 0,false
				}
			}

			if strings.Contains(que, "multiplied "){
				result=result * val3			
			}else if strings.Contains(que, "minus"){
				result=result - val3			
			}else if strings.Contains(que, "divided"){
				result=result / val3			
			}else if strings.Contains(que, "plus") {
				result=result + val3			
			}else {
				if !check{
				return 0,false
				}
			}
			if result==ans && check==true{
				return result,true 
			}
			
		}
		return result,false
	
}
//Benchmark combined time to answer all questions.
func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Answer(test.question,test.ok,test.answer)
		}
	}
}