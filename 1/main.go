package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func fuelRequired(mass int64) int64 {
	return int64(math.Floor(float64(mass/3) - 2))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			a := fuelRequired(n)
			fmt.Printf("Fuel Required: %d\n", a)
			total += a
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %d\n", total)
}
