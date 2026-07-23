package main
import "fmt"

func zeroVal(ival int){
	ival=0
}

func zeroPtr(ival *int){
	*ival++
}

func pointersFunc(){
	i:=1
	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)

	zeroPtr(&i)
	fmt.Println(i)
	fmt.Println(&i)

	p:=new(89)
	fmt.Println(*p)
	zeroPtr(p)
	fmt.Println(*p)

}