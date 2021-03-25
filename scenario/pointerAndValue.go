package scenario

import "fmt"

func PointerRunner() {
	original := 1
	cloner := &original

	*cloner += 10
	fmt.Printf("original : %d\n", original) // -- should be 1 + 10

	original += 20
	fmt.Printf("original : %d\n", original) // -- 1 + 10 + 20
	fmt.Printf("cloner : %d\n", *cloner)
}

func ValueRunner() {
	original := 1
	cloner := original
	cloner += 10
	original += 20

	fmt.Printf("original : %d\n", original) // -- 1 + 20
	fmt.Printf("cloner : %d\n", cloner)     // 1 + 10
}
