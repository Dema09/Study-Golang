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

	var c string = "hello"
	fmt.Printf("%v , %T \n", c,c)

	primitiveType()
	arithmeticOperation()
	bitOperator()

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
	fmt.Printf("a modulus b is: \n", a%b)
}

func bitOperator(){
	a:= 10 //the bit operator is 1010
	b:= 3 // the bit operator is 0011

	//this system will check a pair of bit operator

	fmt.Println(a & b) // AND OPERATOR. This will return 0 if ONE or more of the pairing operator equals 0. so the result will be: 0010 = 2
	fmt.Println(a | b) // OR OPERATOR. This will return 1 if ONE or more of the pairing operator equals 1. so the result will be: 1011 = 11
	fmt.Println(a ^ b) // XOR OPERATOR. This will return 1 if ONE of the operator equals 1, but NOT BOTH equals 1. It will return 0, so the result will be: 1001 = 9
	fmt.Println(a &^ b) // NOR OPERATOR. This will return 1 if BOTH of operator equals 0, otherwise is 0. This will be: 0100 = 8
}

