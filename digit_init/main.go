package digit_init

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Test() {
	fmt.Println("=== DIGIT Authentication ===")
	for {
		choice := readInput("Do you want to (1) Register or (2) Login? ")
		if choice == "1" {
			fmt.Println("1.")
		} else if choice == "2" {
			fmt.Println("2.")
		}
		fmt.Println("Invalid choice. Try again.")
	}
}
