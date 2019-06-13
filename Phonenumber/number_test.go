package phonenumber

import (
	"testing"
	"regexp"
	"strings"
	"errors"
)


func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nNumber(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
			if actual != test.number {
				t.Errorf("FAIL: %s\nNumber(%q): expected [%s], actual: [%s]", test.description, test.input, test.number, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nNumber(%q): expected an error, but error is nil", test.description, test.input)
		}
	
	}
}
func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
		}
	}
}
func TestAreaCode(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := AreaCode(test.input)
		if !test.expectErr {
			if actual != test.areaCode {
				t.Errorf("FAIL: %s\nAreaCode(%q): expected [%s], actual: [%s]", test.description, test.input, test.areaCode, actual)
			}
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nAreaCode(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}
func BenchmarkAreaCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			AreaCode(test.input)
		}
	}
}
func TestFormat(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Format(test.input)
		if !test.expectErr {
			if actualErr != nil {
				// if we don't expect an error and there is one
				var _ error = actualErr
				t.Errorf("FAIL: %s\nFormat(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
			}
			if actual != test.formatted {
				t.Errorf("FAIL: %s\nFormat(%q): expected [%s], actual: [%s]", test.description, test.input, test.formatted, actual)
			}
		} else if actualErr == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nFormat(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}
func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Format(test.input)
		}
	}
}

func Number(input string) (string,error){

	arr:=strings.Fields(input)
    if arr[0]=="+1"{
    	arr[0]=""
    }
    		
    input=strings.Join(arr,"")
    for _, r := range input {
     	if !(r < 'a' || r > 'z') || !(r < 'A' || r > 'Z') {
     	 	      	return input,errors.New("number shouldn't contain alphabets")
        	}
    	}
    re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{0,}`)
    final := re_leadclose_whtsp.ReplaceAllString(input, "")
    final = re_inside_whtsp.ReplaceAllString(input, "")

    reg := regexp.MustCompile("[^0-9]+")

    replaceStr := reg.ReplaceAllString(final, "")
    area:=replaceStr[:3]
	
	if len(replaceStr)==9{
		 return replaceStr,errors.New("invalid 9 digits")
	}else if len(replaceStr)==11 && replaceStr[:1]!="1"{
		 return replaceStr,errors.New("length is 11 but 1st character must be 1")
	}else if len(replaceStr)>11{
		 return replaceStr,errors.New("number shouldn't be more then 11")
	}else if len(replaceStr)==11 && replaceStr[:1]=="1"{	 
		    return replaceStr[1:],nil
    }else if area[:1]=="0"{
		 return replaceStr,errors.New("area code shouldn't be start with zero")
	}else if area[:1]=="1"{
		 return replaceStr,errors.New("area code shouldn't be start with one")
	}else if replaceStr[3:4]=="0"{
		 return replaceStr,errors.New("exchange code shouldn't be start with zero")
	}else if replaceStr[3:4]=="1"{
		 return replaceStr,errors.New("exchange code shouldn't be start with one")
	}else if strings.ContainsAny(replaceStr,"-@:!-"){	 
		   return replaceStr,errors.New("contains punctuations")
    }
	return replaceStr,nil
}
func AreaCode(input string) (string,error){
	arr:=strings.Fields(input)
    if arr[0]=="+1"{
    	arr[0]=""
    }
    		
    input=strings.Join(arr,"")
    for _, r := range input {
     	if !(r < 'a' || r > 'z') || !(r < 'A' || r > 'Z') {
     	 	      	return input,errors.New("number shouldn't contain alphabets")
        	}
    	}
    re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{0,}`)
    final := re_leadclose_whtsp.ReplaceAllString(input, "")
    final = re_inside_whtsp.ReplaceAllString(input, "")

    reg := regexp.MustCompile("[^0-9]+")
    replaceStr := reg.ReplaceAllString(final, "")
	area:=replaceStr[:3]
	if len(replaceStr)==9{
		 return area,errors.New("invalid 9 digits")
	}else if len(replaceStr)==11 && replaceStr[:1]!="1"{
		 return area,errors.New("length is 11 but 1st character must be 1")
	}else if len(replaceStr)>11{
		 return area,errors.New("number shouldn't be more then 11")
	}else if len(replaceStr)==11 && replaceStr[:1]=="1"{	 
		    return replaceStr[1:4],nil
    }else if area[:1]=="0"{
		 return area,errors.New("area code shouldn't be start with zero")
	}else if area[:1]=="1"{
		 return area,errors.New("area code shouldn't be start with one")
	}else if replaceStr[3:4]=="0"{
		 return area,errors.New("exchange code shouldn't be start with zero")
	}else if replaceStr[3:4]=="1"{
		 return area,errors.New("exchange code shouldn't be start with one")
	}else if strings.ContainsAny(replaceStr,"-@:!-"){	 
		   return area,errors.New("contains punctuations")
    }
	return area,nil
}
func Format(input string) (string,error){
	arr:=strings.Fields(input)
    if arr[0]=="+1"{
    	arr[0]=""
    }
    		
    input=strings.Join(arr,"")
    for _, r := range input {
     	if !(r < 'a' || r > 'z') || !(r < 'A' || r > 'Z') {
     	 	      	return input,errors.New("number shouldn't contain alphabets")
        	}
    	}
   	
    re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{0,}`)
    final := re_leadclose_whtsp.ReplaceAllString(input, "")
    final = re_inside_whtsp.ReplaceAllString(input, "")

    reg := regexp.MustCompile("[^0-9]+")
 
    replaceStr := reg.ReplaceAllString(final, "")
	
	var replaceStr1 string
	if len(replaceStr)==11 && replaceStr[:1]=="1"{	 
		    replaceStr1=replaceStr[1:]
		    replaceStr1="("+replaceStr1[0:3]+")"+" "+replaceStr1[3:6]+"-"+replaceStr1[6:]
    }else{
    	replaceStr1="("+replaceStr[0:3]+")"+" "+replaceStr[3:6]+"-"+replaceStr[6:]	
    }
	area:=replaceStr[:3]
	

	if len(replaceStr)==9{
		 return replaceStr1,errors.New("invalid 9 digits")
	}else if len(replaceStr)==11 && replaceStr[:1]!="1"{
		 return replaceStr1,errors.New("length is 11 but 1st character must be 1")
	}else if len(replaceStr)>11{
		 return replaceStr1,errors.New("number shouldn't be more then 11")
	}else if len(replaceStr)==11 && replaceStr[:1]=="1"{	 
		    return replaceStr1,nil
    }else if area[:1]=="0"{
		 return replaceStr1,errors.New("area code shouldn't be start with zero")
	}else if area[:1]=="1"{
		 return replaceStr1,errors.New("area code shouldn't be start with one")
	}else if replaceStr[3:4]=="0"{
		 return replaceStr1,errors.New("exchange code shouldn't be start with zero")
	}else if replaceStr[3:4]=="1"{
		 return replaceStr1,errors.New("exchange code shouldn't be start with one")
	}else if strings.ContainsAny(replaceStr,"-@:!-"){	 
		   return replaceStr1,errors.New("contains punctuations")
    }
	return replaceStr1,nil	
}
