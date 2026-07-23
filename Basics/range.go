package main
import "fmt"

func rangeFunc(){
	nums:=[]int {3,5,6}

	sum:=0

	for _, num:= range nums{				// here _, num represent the index and value in the slice or array eg. 0->3, 1->5, 2->6
		sum+=num
	}
	fmt.Println(sum)

	for i, num:=range nums{
		if num==5{						//5 is present at index 1 so, 1 will be printed
			fmt.Println(i)
		}
	}

	mp := map[string]string{"a":"apple", "b":"banana"}
	for k,v:=range mp{
		fmt.Println(k ,"->", v)
	}

	for i,c:= range "ac"{			//here i represents index and c represents ascii values of characters
		fmt.Println(i,c)
	}
}