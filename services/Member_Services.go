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
			fmt.Println("Please enter a Number!")
			continue
		}
		switch choice {
		case 0:
			fmt.Println("<-Back Page")
			return
		case 1:
			AddMember(lib, reader)
		case 2:
			ListMembers(lib)
		case 3:
			FilterMembers(lib, reader)
		case 4:
			FindMember(lib, reader)
		case 5:
			EditMember(lib, reader)
		case 6:
			DeleteMember(lib, reader)
		default:
			fmt.Println("❌Item is not available!")
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
	fmt.Println("✅Successfully added member")
}

func ListMembers(lib *models.Library) {
	fmt.Println("\n=== library Menu > Manage Members > Members List ===")
	for _, member := range lib.Members {
		fmt.Printf("- %d : %s - %s", member.ID, member.Name, member.Phone)
	}
}

func FilterMembers(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Members > Filter Members ===")
	fmt.Print("Member Name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	var result []models.Member
	for _, member := range lib.Members {
		if strings.Contains(strings.ToLower(member.Name), strings.ToLower(name)) {
			result = append(result, member)
		}
	}
	fmt.Println("*** Filter Result ***")
	if len(result) == 0 {
		fmt.Println("❌Member not found")
		return
	}
	for _, r := range result {
		fmt.Printf("-%d - %s - %s \n", r.ID, r.Name, r.Phone)
	}
}

func FindMember(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Members > Find Member ===")
	fmt.Print("Member Phone : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	fmt.Println("*** Search Result ***")
	for _, p := range lib.Members {
		if p.Phone == phone {
			fmt.Printf("-%d - %s - %s \n", p.ID, p.Name, p.Phone)
			return
		}
	}
	fmt.Println("❌Member not found")
}

func EditMember(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Members > Edit Member ===")
	fmt.Print("Member Phone : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	for i, m := range lib.Members {
		if m.Phone == phone {
			fmt.Println("*** Your member before editing ***")
			fmt.Println("*** for no changing press enter ***")
			fmt.Printf("-%d - %s - %s \n", m.ID, m.Name, m.Phone)
			fmt.Printf("New Name (old: %s)", m.Name)
			nName, _ := reader.ReadString('\n')
			nName = strings.TrimSpace(nName)
			if nName != "" {
				m.Name = nName
			}
			fmt.Printf("New Phone (old: %s)", m.Phone)
			nPhone, _ := reader.ReadString('\n')
			nPhone = strings.TrimSpace(nPhone)
			if nPhone != "" {
				m.Phone = nPhone
			}
			lib.Members[i] = m
			fmt.Println("✅ Member updated successfully.")
			return
		}
	}
	fmt.Println("❌Member not found!")
}

func DeleteMember(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Delete Member ===")
	fmt.Print("Member ID : ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please enter a integer ID!")
		return
	}
	for i, m := range lib.Members {
		if m.ID == ID {
			lib.Members = append(lib.Members[:i], lib.Members[i+1:]...)
			fmt.Println("✅ Member deleted successfully.")
			return
		}
	}
	fmt.Println("❌Member not found!")
}
