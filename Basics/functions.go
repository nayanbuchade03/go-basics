package main
import "fmt"

func sum1(a int, b int)int{
	return a+b
}

func sum2(a,b,c int) int {
	return a+b+c
}

func vals()(int , int){
	return 6,7
}

func sum3(arrs ...int){
	fmt.Println(arrs," ")
	total:=0

	for _, num := range(arrs){
		total+=num
	}
	fmt.Println(total)
}

func functionFunc(){
	// r1:=sum1(2,3)
	// fmt.Println(r1)

	// r2:=sum2(5,6,7)
	// fmt.Println(r2)

	// v1,v2:=vals()
	// fmt.Println(v1,v2)

	// _,v3:=vals()				//here _ suggests that the value return is not stored or used anywhere
	// fmt.Println(v3)

	sum3(4,3,2)
	sum3(4,3)

	nums:=[]int {3,6,7}
	sum3(nums...)

}