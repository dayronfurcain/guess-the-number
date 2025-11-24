package main

import (
	"fmt"
	"math/rand"
)

func main() {
	value := rand.Intn(100) + 1
	attemps := 10
	guessed := false

	fmt.Println("Adivina el numero de entre 1 y 100.")
	fmt.Printf("Tiene %d intentos para adivinar el numero.\n", attemps)

	var number int

	for attemps > 0 {
		fmt.Println("\nIntroduzca el numero:")
		fmt.Scanln(&number)
		if number > value {
			attemps--
			fmt.Printf("Demasiado alto. Intento numero %d.", 10-attemps)
		} else if number < value {
			attemps--
			fmt.Printf("Demasiado bajo. Intento numero %d.", 10-attemps)
		} else {
			guessed = true
			fmt.Printf("Adivino el numero %d en %d intentos", value, 10-attemps)

		}

	}

	if !guessed {
		fmt.Printf("\nEl numero era %d.", value)
	}

}
