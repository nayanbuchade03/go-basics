package main
import "fmt"

func forFunc(){

	i:=1					//works like while loop
	for i<=3{
		fmt.Println(i)
		i++
	}

	for i:=0;i<=3;i++{		//normal for loop
		fmt.Println(i)
	}

	for i:= range 3{			//using range
		fmt.Println("range",i)
	}
}