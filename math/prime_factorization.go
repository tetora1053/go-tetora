package main

import (
	"fmt"
	"math"
)

func PrimeFact(n int) (pfs []int) {
	fmt.Println("test")
	for n % 2 == 0 {
		pfs = append(pfs, 2)
		n /= 2
	}
	for i:=3; i<int(math.Sqrt(float64(n))); i+=2 {
		if n % i == 0 {
			pfs = append(pfs, i)
			n /= i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}
	return 
}

func main() {
	ans := PrimeFact(124)
	fmt.Println(ans)
}
