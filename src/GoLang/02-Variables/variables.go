package main

import "fmt"

const LoginToken string = "abcd" //public

func main() {
	var userName string = "Piyush"
	fmt.Println(userName)

	fmt.Printf("Variable is of type:%T \n", userName)

	var isLogin bool = false
	fmt.Println(isLogin)
	fmt.Printf("Variable is of type:%T \n", isLogin)

	var uInt uint8 = 255
	fmt.Println(uInt)
	fmt.Printf("Variable is of type:%T \n", uInt)

	var floatNo float64 = 25.34567
	fmt.Println(floatNo)
	fmt.Printf("Variable is of type:%T \n", floatNo)

	//default value and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:%T \n", anotherVariable)

	//no var style
	noOfuser := 3000
	fmt.Println(noOfuser)

	fmt.Println(LoginToken)

}
