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

func myFunc() func() int{
	i:=0
	return func () int{
		i++
		return i
	}
}

func testFunc(testfunc func(int) int){
	fmt.Println(testfunc(7))
}

//closure
func returnFunc(x string) func(){
	return func(){
		fmt.Println(x)			//here when we call returnFunc("hello") x gets that "hello" and returns not printed yet, still is not disappears even its returned because its job is not over i.e to print "hello" and this is called closure
	}
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

	// sum3(4,3,2)
	// sum3(4,3)

	// nums:=[]int {3,6,7}
	// sum3(nums...)

	// v:=myFunc()
	// fmt.Println(v())
	// fmt.Println(v())
	// fmt.Println(v())

	// v2:=myFunc()
	// fmt.Println(v2())

	// test:= func (x int) int {			//anonymous function
	// 	return x+9
	// }
	// testFunc(test)

	// func (x int){
	// 	fmt.Println(x)
	// }(9)

	returnFunc("hello")()			// it is same like temp:=returnFunc("hello")   and then temp()
}