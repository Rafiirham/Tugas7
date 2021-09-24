  
package main

import (
	"fmt"
	"math"
)

var n, i, count float64

func suma(n float64) float64 {
	if n == 1 {
		i = 6
		fmt.Print(i, "+", n)
		return i + n
	} else {
		count = math.Pow(6, n)
		fmt.Print(count, "+")
		return count + suma(n-1)
	}
}

func sumb(n float64) float64 {
	if n == 1 {
		i = 5
		fmt.Println(i)
		return i
	} else {
		count = 5 * (math.Pow(10, n-1))
		fmt.Print(count, "+")
		return count + sumb(n-1)
	}
}

func sumc(n float64) float64 {
	if n == 1 {
		i = 7
		fmt.Print(i, "+", n)
		return i + n
	} else {
		count = math.Pow(7, n)
		fmt.Print(count, "+")
		return count + sumc(n-1)
	}
}

func sumd(n float64) float64 {
	if n == 1 {
		i = 1
		fmt.Println(i)
		return i
	} else {
		count = n * (math.Pow(10, n-1))
		fmt.Print(count, "+")
		return count + sumd(n-1)
	}
}

func sume(n float64) float64 {
	if n == 1 {
		i = 5
		fmt.Println(i)
		return i
	} else {
		count = 5 * (math.Pow(10, n-1))
		fmt.Print(count, "+")
		return count + sume(n-1)
	}
}

func main() {
	//driver code
	n = 5
	fmt.Println("Result series a : ", suma(n))
	fmt.Println("Result series b : ", sumb(n))
	fmt.Println("Result series c : ", sumc(n))
	fmt.Println("Result series d : ", sumd(n))
	fmt.Println("Result series e : ", sume(n))
}