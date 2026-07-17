package main
import "fmt"

func arraysFunc(){
	var a [5]int
	fmt.Println(a)

	a[4]=400
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a))

	b:= [4]int{1,2,4,5}
	fmt.Println(b)

	c:= [...]int{34,5,6,7,78,1}
	fmt.Println(c)

	d:=[...]int{45,400, 4: 6,23}		//here 4: means till 3rd index verything will be zero but starting elements will be as it is
	fmt.Println(d)

	var twoD [3][2]int
	for i:= range 3{
		for j:= range 2{
			twoD[i][j]=i*j
		}
	}
	fmt.Println(twoD)
}