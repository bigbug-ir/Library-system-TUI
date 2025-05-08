package main

import (
	"bufio"

	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bigbug-ir/Library-system-TUI/models"
	"github.com/bigbug-ir/Library-system-TUI/services"
)

func main() {
	lib := models.Library{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== library Menu ===")
		fmt.Println("1.Manage Books->")
		fmt.Println("2.Manage Members->")
		fmt.Println("3.Manage Loan->")
		fmt.Println("4.<Exit>")

		fmt.Print("Slect Item : ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Please enter a Number")
			continue
		}
		switch choice {
		case 1:
			services.ManageBooks(&lib, reader)
		case 2:
			services.ManageMember(&lib, reader)
		case 3:
			services.ListMembers(&lib)
		case 4:
			fmt.Println("Exit Program")
			return
		default:
			fmt.Println("Item is not available!")
		}
	}
}
