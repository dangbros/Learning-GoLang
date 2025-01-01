package main

import (
	"fmt"
	"math"
	"strconv"
)

func FloattoNum() {
	var floatNum float64
	fmt.Println("enter the floating point number: ")
	fmt.Scan(&floatNum)
	intNum := int(floatNum)
	fmt.Println("floating point number: ", floatNum, "type-casted int number: ", intNum)
}

func StringToIntegerAndViceVersa() {
	var str string
	fmt.Println("Enter the numerical string: ")
	fmt.Scan(&str)
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	num = num + 10
	fmt.Println("adding 10 to the num", num)
	s := strconv.Itoa(num)
	concatStr := str + s
	fmt.Println("concatenated string: ", concatStr)
}

func ConversionsAndRanges() {
	num := 12383982737.45678289732
	num32 := int32(num)
	num8 := int8(num)
	fmt.Printf("float num: %v\nnum32: %v\nnum8: %v\n", num, num32, num8)
}

func TypeConversionInExpressions() {
	var num1 float32
	var num2 int
	fmt.Printf("Enter the numbers: ")
	fmt.Scan(&num1, &num2)
	num2float := num1 * float32(num2)
	fmt.Printf("converting num2 to float: %v\n", num2)
	fmt.Printf("%v x %v = %v\n", num1, num2, num2float)
}
func SquareRoot() {
	var num int
	fmt.Printf("Enter the number: ")
	fmt.Scan(&num)
	sqRoot := math.Sqrt(float64(num))
	fmt.Printf("Square Root of %v is %f\n", num, sqRoot)
}
func main() {
	// FloattoNum()
	//StringToIntegerAndViceVersa()
	//ConversionsAndRanges()
	//TypeConversionInExpressions()
	//SquareRoot()
}
