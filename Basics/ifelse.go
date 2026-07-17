package main

import "fmt"

func ifelseFunc() {
	if 75%2 == 0 {
		fmt.Println("even")
	} else{							//ekse should be in same line of closing braces of if
		fmt.Println("odd")
	}

	if 8%4==0{
		fmt.Println("divisible")
	}

	if x:=12; x<10{
		fmt.Println("Small")
	}else if x<50{
		fmt.Println("Medium")
	}else{
		fmt.Println("Big")
	}

	//no ternary operation
}