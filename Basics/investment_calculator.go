package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years : ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
