package main 

import "fmt"

func main(){
	s:=make([]string,3)
	fmt.Println("emp:",s)

	s=append(s,"c","h","i","r","a","g")
	fmt.Println("s:",s)

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println	("c:",c)

	l:=s[2:5]
	fmt.Println("l:",l)

	l=s[:5]
	fmt.Println("l:",l)

	l=s[2:]
	fmt.Println("l:",l)

	twoD:=make([][]int,3)
	for i:=0; i<3; i++{
		innerlen:=i+1
		twoD[i]=make([]int,innerlen)
		for j:=0; j<innerlen;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2D:",twoD)
}