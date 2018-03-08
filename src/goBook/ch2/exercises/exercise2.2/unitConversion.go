package main

import (
	"fmt"
	wc "goBook/ch2/exercises/exercise2.2/weightConv"
	lc "goBook/ch2/exercises/exercise2.2/lengthConv"
	tc "goBook/ch2/exercises/exercise2.2/tempconv"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printChoices()
		scanner.Scan()
		c := scanner.Text()
		choice, _ := strconv.Atoi(c)
		fmt.Println("Enter number to convert.")
		scanner.Scan()
		toConvert := scanner.Text()
		temp, _ := strconv.ParseFloat(toConvert, 64)
		switch {
		case choice == 1:
			k := wc.Kilogram(temp)
			fmt.Printf("%s == %s\n", k, wc.KgToLb(k))
		case choice == 2:
			lb := wc.Pound(temp)
			fmt.Printf("%s == %s\n", lb, wc.LbToKG(lb))
		case choice == 3:
			f := lc.Feet(temp)
			fmt.Printf("%s == %s\n", f, lc.FToM(f))
		case choice == 4:
			m := lc.Meter(temp)
			fmt.Printf("%s == %s\n", m, lc.MToF(m))
		case choice == 5:
			c := tc.Celsius(temp)
			fmt.Printf("%s == %s\n", c, tc.CToF(c))
		case choice == 6:
			c := tc.Celsius(temp)
			fmt.Printf("%s == %s\n", c, tc.CToK(c))
		case choice == 7:
			f := tc.Fahrenheit(temp)
			fmt.Printf("%s == %s\n", f, tc.FToC(f))
		case choice == 8:
			f := tc.Fahrenheit(temp)
			fmt.Printf("%s == %s\n", f, tc.FToK(f))
		case choice == 9:
			k := tc.Kelvin(temp)
			fmt.Printf("%s == %s\n", k, tc.KToC(k))
		case choice == 10:
			k := tc.Kelvin(temp)
			fmt.Printf("%s == %s\n", k, tc.KToF(k))
		}
		fmt.Println("Choose again? Y/N")
		scanner.Scan()
		res := scanner.Text()
		if res == "Y" || res == "y" {
			continue
		} else {
			break
		}

	}
}

func printChoices() {
	fmt.Println("1.  Convert from KG to LBs")
	fmt.Println("2.  Convert from LBs to KG")
	fmt.Println("3.  Convert from Feet to Meters")
	fmt.Println("4.  Convert from Meters to Feet")
	fmt.Println("5.  Convert from C to F")
	fmt.Println("6.  Convert from C to K")
	fmt.Println("7.  Convert from F to C")
	fmt.Println("8.  Convert from F to K")
	fmt.Println("9.  Convert from K to C")
	fmt.Println("10. Convert from K to F")

}
