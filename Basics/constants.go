package main
import (
	"fmt"
	"math"
)

func constantsFunc(){
	const a string="constant"
	// const a="Variable"     this gives error
	fmt.Println(a)

	const n=500000000
	fmt.Println(math.Sin(n))

	const d=3e20/n;
	fmt.Println(d);
}