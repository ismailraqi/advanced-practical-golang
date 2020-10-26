package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func addition(num ...float64) float64 {
	var result float64 = 0
	for _, e := range num {
		result += e
	}
	return result
}

func calc(transType string, amount ...float64) (float64, error) {
	var totalAmount float64 = 0
	if len(strings.TrimSpace(transType)) == 0 {
		return 0, errors.New("transaction type is required")
	}
	for n, a := range amount {
		switch strings.TrimSpace(transType) {
		case "addition":
			totalAmount += a
		case "substract":
			if n > 0 {
				totalAmount -= a
			} else if n == 0 {
				totalAmount = a
			}
		}
	}
	return totalAmount, nil
}

func strVariadic(strVal ...string) {
	for n, s := range strVal {
		fmt.Println("n: ", n, " s: ", s)
	}
}

type user struct {
	username, firstname, lastname string
}

func strVariadicInterface(strVal ...interface{}) {
	for _, i := range strVal {
		fmt.Println(i, ": ", reflect.ValueOf(i).Kind())
	}
}

func main() {
	fmt.Println("Hello Variadic Function")
	add := addition(2, 2, 4)
	fmt.Println("add: ", add)

	// addition variadic function
	amount, err := calc("addition", 10, 12, 3)
	if err != nil {
		fmt.Println("Oops!, got some errors: ", err)
	}
	fmt.Println("addition: ", amount)

	// string variadic
	strVariadic("a", "b", "c")

	strVariadicInterface(user{username:  "jr",firstname: "john",lastname:  "rambo",}, 10.66, "hello")
}
