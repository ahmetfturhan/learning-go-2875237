package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Value 1: ")
	input1 := getInput()
	float1, err := strconv.ParseFloat(input1, 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2 := getInput()
	float2, err := strconv.ParseFloat(input2, 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	fmt.Print("Select operation (*, -, +, /): ")
	opStr := getInput()

	var result float64
	switch opStr {
	case "*":
		result = multiply(float1, float2)
		result = math.Round(result*100) / 100
		fmt.Printf("The multiplication of %v and %v is %v\n\n", float1, float2, result)
	case "+":
		result = add(float1, float2)
		result = math.Round(result*100) / 100
		fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, result)
	case "-":
		result = subtract(float1, float2)
		result = math.Round(result*100) / 100
		fmt.Printf("The subtraction of %v and %v is %v\n\n", float1, float2, result)
	case "/":
		result = divide(float1, float2)
		result = math.Round(result*100) / 100
		fmt.Printf("The division of %v and %v is %v\n\n", float1, float2, result)
	default:
		fmt.Println("Error")
	}


}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	theInput := strings.TrimSpace(input1)

	return theInput
}

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}
func subtract(num1 float64, num2 float64) float64 {
	return num1 - num2
}
func multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}
func divide(num1 float64, num2 float64) float64 {
	return num1 / num2
}
