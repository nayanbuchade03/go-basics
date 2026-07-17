package main
import (
	"fmt"
	"time"
)

func switchcaseFunc(){
	i:=2

	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	x:=time.Now()
	switch {
	case x.Hour()<12:
		fmt.Println("Before Noon")
	default:
		fmt.Println("After Noon")
	}

	z:=func(f interface{}){
		switch f.(type){
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("char")
		}
	}

	z(true)
	z('c')
	z("sddf")
	z(8)
}