package main

import (
"fmt"
//"time"
"reflect"
)

var (
a=10
b=0
)

func main() {
	
//	defer adddd()
	var x int=1
	fmt.Println("hello x,a,b", x,a,b)
	const ca string=" i am constant"
	fmt.Println(ca)

	var num1,num2 = 10,20
	fmt.Println(num1>num2)
	fmt.Println(add(num1,num2))
	//time.Sleep (10*time.Second)
	test("abc","ddd","dddde","ffe")
	test1(111,"red",true,10.045,[]string{"foo","bar","baz"},
		   map[string]int{"apple":23,"tomato":13})
	fmt.Println("enter action selection ( 1/2):")
	var action int
	fmt.Scanln(&action)

	switch action {
		case 1:
			fmt.Println("Action 1")
		case 2:
			fmt.Println("Action 2")
		default:
			fmt.Println("Other action")		
	}

	fmt.Println("enter country selection ( 1/2):")
	fmt.Scanln(&action)

	switch action{
		case 1:
			fmt.Println("country 1")
		case 2:
			fmt.Println("country 2")
		default:
			panic(fmt.Sprintln("Other country %d",action))		
	}
    defer func() {
        action := recover()     
        fmt.Println("hello",action)         
    }()
}


func add (x int,y int) int {

	total :=10
	total +=x; total +=y
	return total
}

func adddd(){
	fmt.Println("this is defer fucntion")
}

func test (s ... string){
	fmt.Println(len(s),s)
	//var m int
	//var n string
	for m,n := range s{
		fmt.Println(m,n)
	}

}

func test1(s ... interface{}) {
for _,v :=range s{
	fmt.Println(v,"--",reflect.ValueOf(v).Kind())
}

}