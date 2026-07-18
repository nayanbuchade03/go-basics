package main
import "fmt"

func mapsFunc(){
	m:=make(map[int]string)

	m[2]="d"
	m[4]="df"

	fmt.Println(m)

	v1:=m[3]				//zero value for string type is nothing(blank space)
	fmt.Println(v1)

	mp:=make(map[string]int)

	mp["tg"]=7
	mp["gh"]=90

	fmt.Println(mp)

	v2:=mp["l"]				//zero value for int type is 0
	fmt.Println(v2)

	fmt.Println(len(mp))

	delete(mp,"gh")
	fmt.Println(mp)
	fmt.Println(len(mp))

	clear(m)
	fmt.Println(m)
	fmt.Println(len(m))

	_,v3:=mp["tg"]			// _ tells that return boolean value if key is present or not
	fmt.Println(v3)

}