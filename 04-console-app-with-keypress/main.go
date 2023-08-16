package main

import (
	"fmt"
	"log"

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

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("1 You pressed:", char, key)

		} else {
			fmt.Println("2 You pressed:", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program exiting...")
}
