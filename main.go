package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

func clearScreen() {
	fmt.Print("\033[2J\033[1;1H")
}

func main() {
	const cursor string = "█"
	// getting the width and height of the terminal
	w, h, _ := term.GetSize(int(os.Stdout.Fd()))

	// initial placement of the player
	x, y := 3, 0

	// saving
	screen := make([][]rune, h)
	for i := range screen {
		screen[i] = make([]rune, w)
		for j := range screen[i] {
			screen[i][j] = ' '
		}
	}

	// opening keyboard
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error while opennig keyboard:", err)
		return
	}
	defer keyboard.Close()

	fmt.Print("\033[?25l") // hiding cursor
	fmt.Print("\033[2J")   // clearing the screen
	// line numbering
	for i := 0; i < h-1; i++ {
		screen[i][0] = rune(((i+1)/10)%10 + 48)
		screen[i][1] = rune((i+1)%10 + 48)
	}
	defer fmt.Print("\033[?25h")

	// initilize the screen
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Printf("\033[%d;%dH%c", i+1, j+1, screen[i][j])
		}
	}

	// showing player in initail position
	fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)

	// main loop
	for {
		// reading keyboard
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error while reading the key:", err)
			return
		}

		// movement with ArrowKey
		switch key {
		case keyboard.KeyArrowUp:
			if y > 0 {
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, screen[y][x])
				y--
				fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)
			}
		case keyboard.KeyArrowDown:
			if y < h-1 {
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, screen[y][x])
				y++
				fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)
			}
		case keyboard.KeyArrowLeft:
			if x > 3 {
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, screen[y][x])
				x--
				fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)
			}
		case keyboard.KeyArrowRight:
			if x < w-1 {
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, screen[y][x])
				x++
				fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)
			}
		case keyboard.KeyBackspace:
			if x > 3 {
				screen[y][x-1] = ' '
				fmt.Printf("\033[%d;%dH%c", y+1, x+1, screen[y][x])
				x--
				fmt.Printf("\033[%d;%dH%s", y+1, x+1, cursor)
			}
		case keyboard.KeyEsc:
			clearScreen()
			return
		}

		// showing typed chars
		if char != 0 && key == 0 {
			screen[y][x] = char
			fmt.Printf("\033[%d;%dH%c", y+1, x+1, char)
			fmt.Printf("\033[%d;%dH%s", y+1, x+2, cursor)
			x++
		}
	}

}
