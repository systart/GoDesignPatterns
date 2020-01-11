/*
I think this pattern is the essence of AOP (Aspect Oriented Programming)
and SOLID's "Single Responsibiliy" rule.
In a nutshell; a function takes another function as a parameter.
Main benefit of this pattern can be achieve while logging, caching,
error handling, transaction management etc. (See also: Crosscutting concerns)
*** NOTE: I'm a newbie at golang, maybe this implementation was wrong.
*** I would be very grateful to you if you warn me about anything wrong :)
*/

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func main() {
	calc := &Calculator{}
	logger := &Logger{}
	result1 := logger.logDecorator(calc.sum, 2, 4)
	fmt.Println("Result: ", result1)
	result2 := logger.logDecorator(calc.multiply, 2, 4)
	fmt.Println("Result: ", result2)
	logger.printAllLogs()
}

type Decoratable func(x, y int) int

type Calculator struct{}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (calculator *Calculator) sum(x, y int) int {
	return x + y
}

func (calculator *Calculator) multiply(x, y int) int {
	return x * y
}

type Logger struct {
	logs []string
}

func (logger *Logger) log(message string) {
	logger.logs = append(logger.logs, message)
	fmt.Println("Log added succesfully.")
}

func (logger *Logger) logDecorator(fn Decoratable, x, y int) int {
	result := fn(x, y)
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	strResult := strconv.Itoa(result)
	logger.log("Function Name: " + GetFunctionName(fn) + " *** X: " + strX + " *** Y: " + strY + " *** Result: " + strResult)
	return result
}

func (logger *Logger) printAllLogs() {
	for i, log := range logger.logs {
		fmt.Println("Log ", i, ": ", log)
	}
}
