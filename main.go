package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const lowTempFWarning float64 = 0
const highTempFWarning float64 = 100
const maxLoop int = 5

func main() {
	var f, c, wc, wcel float64 = 0, 0, 0, 0

	var err error
	var fText, wcText string

	for i := 0; i < maxLoop; i++ {
		fmt.Println("\nEnter a temperature in Fahrenheit: ")
		fmt.Scanln(&fText)

		fmt.Println("\nEnter the windspeed: ")
		fmt.Scanln(&wcText)

		fText = strings.Trim(fText, " \n\r")
		if f, err = strconv.ParseFloat(fText, 64); err != nil {
			fmt.Println("\nNot a valid number")
			os.Exit(-1)
		}

		wcText = strings.Trim(wcText, " \n\r")
		var ws float64
		if ws, err = strconv.ParseFloat(wcText, 64); err != nil {
			fmt.Println("\nNot a valid number")
			os.Exit(-1)
		}

		wc = 35.74 + 0.6215*f - 35.75*math.Pow(ws, 0.16) + 0.4275*f*math.Pow(ws, 0.16)

		c = (f - 32) * 5 / 9
		wcel = (wc - 32) * 5 / 9

		fmt.Printf("\nThe temperature in Celsius is: %0.2f but feels like: %0.2f in Farenheit (%0.2f Celsius)", c, wc, wcel)

		if f > 100 || wc > 100 {
			fmt.Println("\nRemember to hydrate")
		}
		if f < 0 || wc < 0 {
			fmt.Println("\nRemember to pack Long underwear")
		}
	}
	os.Exit(-1)
}
