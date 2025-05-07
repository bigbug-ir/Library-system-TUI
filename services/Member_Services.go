package services

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/bigbug-ir/Library-system-TUI/models"
)

func ManageMember(lib *models.Library, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== library Menu > Manage Members ===")
		fmt.Println("0.<-Back")
		fmt.Println("1.Add Member->")
		fmt.Println("2.List Members->")
		fmt.Println("3.Filter Members->")
		fmt.Println("4.Find Member->")
		fmt.Println("5.Edit Member->")
		fmt.Println("6.Delete Member->")

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
			AddMember(lib, reader)
		case 2:
			ListMembers(lib)
		case 0:
			fmt.Println("Back Page")
			return
		default:
			fmt.Println("Item is not available!")
		}
	}
}

func AddMember(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Members > Add Member ===")
	fmt.Print("Member Name : ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Member Phone : ")
	phone, _ := reader.ReadString('\n')
	member := models.Member{
		ID:    len(lib.Members) + 1,
		Name:  strings.TrimSpace(name),
		Phone: strings.TrimSpace(phone),
	}
	lib.Members = append(lib.Members, member)
	fmt.Println("Successfully added member✅")
}

func ListMembers(lib *models.Library) {
	fmt.Println("\n=== library Menu > Manage Members > Members List ===")
	for _, member := range lib.Members {
		fmt.Printf("- %d : %s - %s", member.ID, member.Name, member.Phone)
	}
}
