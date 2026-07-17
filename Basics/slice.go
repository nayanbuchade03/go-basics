package main
import (
	"fmt"
	"slices"
)

func sliceFunc(){

	// these are two ways of initializing slice
	var s []int
	fmt.Println(s, s==nil, len(s)==0)

	s=make([]int, 3)
	fmt.Println(s,len(s),cap(s))

	str:=[]string{"s","r"}
	fmt.Println(str)

	s[1]=5
	fmt.Println(s,len(s),cap(s))

	s = append(s, 8)
	s = append(s, 5)
	s = append(s, 34)
	s = append(s, 23)
	fmt.Println(s,len(s),cap(s))

	c:=make([]int, len(s))
	copy(c,s)
	fmt.Println(c)

	l:=s[2:6]			//here 2nd index is included but not only till 5th not 6th
	fmt.Println(l)

	x:=s[3:]
	fmt.Println(x)

	y:=s[:3]
	fmt.Println(y)

	str2:=make([]string,len(str))
	copy(str2,str)

	if slices.Equal(str,str2){
		fmt.Println("equal")
	}
}