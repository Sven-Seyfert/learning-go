package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// anonymous function (no name)
	// defer waits until the end of the function before execute the function body
	defer func() {
		_ = keyboard.Close()
		fmt.Println("End")
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	menu := `Menu
----
1 - Cappucino
2 - Latte
3 - Americano
4 - Mocha
5 - Macchiato
6 - Espresso
Q - Quit the program
`

	fmt.Println(menu)

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		text := fmt.Sprintf("You chose: %s", coffees[i])

		fmt.Println(text)
	}

	fmt.Println("Program exiting...")
}
