package main

import "fmt"

func main() {
	fmt.Printf("hi")
func GachaponRunner() {
	fmt.Println("Gachapox generator")
	fmt.Println("Item should be empty after 10 roll")
	gachaponer := scenario.GenerateGachapox()
	for i := 1; i <= 15; i++ {
		item := scenario.RNG(gachaponer)
		fmt.Printf("Gachapon random %d got %+v\n", i, item)
	}

	fmt.Println("Gachapin generator")
	fmt.Println("Alway get item")
	gachaponer = scenario.GenerateGachapin()
	for i := 1; i <= 15; i++ {
		item := scenario.RNG(gachaponer)
		fmt.Printf("Gachapin random %d got %+v\n", i, item)
	}
}
