package main

//we need to import "strconv" for converting string.

import (
	"fmt"
	"strconv"
)

//we can declare multiple variable just like this:
//if we declare variable at the same level, for example at main level, that will be an error.
//But, if we declare at different level, for example variable i declared at package level and main variable, it won't error.
//But the value will be overwrite.



func main(){

	//%v is for printing the value of variables, %T is for print type of the variable.

	var name string = "Made Raja Adi"
	fmt.Printf("%v, %T \n",name,name)

	//converting variables example
	var i int = 42
	fmt.Printf("%v , %T \n", i, i)

	//converting into float32
	var j float32
	j = float32(i)
	fmt.Printf("%v , %T \n", j, j)

	//converting to string
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v , %T \n", k, k)

	primitiveType()
	arithmeticOperation()

}

func primitiveType() {
	//boolean data
	//the default value in primitive type is zero values. in this case, for boolean is false.
	var n bool = true
	var m bool
	fmt.Printf("%v,%T \n", n, n)
	fmt.Printf("%v,%T \n", m, m)

	//how to check item if it equals or not
	a := 1 == 1
	b := 1 == 2

	//a will return true because the values are same
	fmt.Printf("%v \n", a)
	//b will return false because the value are not same
	fmt.Printf("%v \n", b)

	var i uint16 = 42
	fmt.Printf("%v,%T \n", i, i)

}

func arithmeticOperation(){
	a := 11
	b := 5
	fmt.Printf("a + b = %v \n", a+b)
	fmt.Printf("a - b = %v \n", a-b)
	fmt.Printf("a * b = %v \n", a*b)
	fmt.Printf("a / b = %v \n", a/b)
	fmt.Printf("a modulus b is: ", a%b)
}

