package services

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/bigbug-ir/Library-system-TUI/models"
)

func ManageLoans(lib *models.Library, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== library Menu > Manage Loans ===")
		fmt.Println("0.<-Back")
		fmt.Println("1.Add Loan->")
		fmt.Println("2.List Loans->")
		fmt.Println("3.Filter Loans->")
		fmt.Println("4.Find Loan->")
		fmt.Println("5.Edit Loan->")
		fmt.Println("6.Delete Loan->")
		fmt.Print("Slect Item : ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("❌ Please enter a valid number.")
			continue
		}
		switch choice {
		case 0:
			return
		case 1:
			AddLoan(lib, reader)
		case 2:
			ListLoans(lib)
		case 3:
			FilterLoans(lib, reader)
		case 4:
			FindLoan(lib, reader)
		case 5:
			EditLoan(lib, reader)
		case 6:
			DeleteLoan(lib, reader)
		default:
			fmt.Println("❌ Invalid option.")
		}
	}
}

func AddLoan(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Add Loan ===")
	fmt.Print("Book ID : ")
	bookIdstr, _ := reader.ReadString('\n')
	bookIdstr = strings.TrimSpace(bookIdstr)
	bookId, err := strconv.Atoi(bookIdstr)
	if err != nil {
		fmt.Println("❌ Invalid Book ID")
		return
	}
	var book *models.Book
	for i := range lib.Books {
		if lib.Books[i].ID == bookId {
			book = &lib.Books[i]
			break
		}
	}
	if book == nil {
		fmt.Println("❌ Book not found")
		return
	}
	if !book.Available {
		fmt.Println("❌ Book is currently not available")
		return
	}
	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("❌ Invalid Member ID")
		return
	}
	var memberExists bool
	for _, m := range lib.Members {
		if m.ID == memberID {
			memberExists = true
			break
		}
	}
	if !memberExists {
		fmt.Println("❌ Member not found")
		return
	}
	fmt.Print("Enter Loan Date (e.g. 2025-05-09): ")
	dateOut, _ := reader.ReadString('\n')
	dateOut = strings.TrimSpace(dateOut)
	loan := models.Loan{
		BookID:   bookId,
		MemberID: memberID,
		DateOut:  dateOut,
		DateIn:   "",
	}
	lib.Loans = append(lib.Loans, loan)
	book.Available = false
	fmt.Println("✅ Loan successfully added.")
}
func ListLoans(lib *models.Library) {
	fmt.Println("\n=== library Menu > Manage Loans > List Loans ===")
	if len(lib.Loans) == 0 {
		fmt.Println("📭 No loans found.")
		return
	}
	for i, loan := range lib.Loans {
		var bookTitle, memberName string

		for _, book := range lib.Books {
			if book.ID == loan.BookID {
				bookTitle = book.Title
				break
			}
		}
		for _, member := range lib.Members {
			if member.ID == loan.MemberID {
				memberName = member.Name
				break
			}
		}
		fmt.Printf("\n%d. Book: %s (ID: %d)\n", i+1, bookTitle, loan.BookID)
		fmt.Printf("   Borrower: %s (ID: %d)\n", memberName, loan.MemberID)
		fmt.Printf("   Loan Date: %s\n", loan.DateOut)
		if loan.DateIn != "" {
			fmt.Printf("   Returned: %s\n", loan.DateIn)
		} else {
			fmt.Printf("   Returned: ❌ Not yet\n")
		}
	}
}
func FilterLoans(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Filter Loans ===")
	if len(lib.Loans) == 0 {
		fmt.Println("📭 No loans found.")
		return
	}
	for {
		fmt.Println("\n=== library Menu > Manage Loans > Filter Loans ===")
		fmt.Println("0.<-Back")
		fmt.Println("1.Filter Loans By Member")
		fmt.Println("2.Filter Loans By Book")
		fmt.Println("3.Filter Loans By Return Status")
		fmt.Println("4.Filter Loans By Date Out")
		fmt.Print("Select filter option: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("❌ Please enter a valid number.")
			continue
		}

		switch choice {
		case 0:
			return
		case 1:
			FilterLoansByMember(lib, reader)
		case 2:
			FilterLoansByBook(lib, reader)
		case 3:
			FilterLoansByReturnStatus(lib)
		case 4:
			FilterLoansByDateOut(lib, reader)
		default:
			fmt.Println("❌ Invalid option.")
		}
	}
}

func FilterLoansByMember(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Filter Loans > Filter Loans By Member ===")
	fmt.Print("Enter Member ID or Name : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	var matchedMembers []models.Member
	if id, err := strconv.Atoi(input); err == nil {
		for _, m := range lib.Members {
			if m.ID == id {
				matchedMembers = append(matchedMembers, m)
				break
			}
		}
	} else {
		for _, m := range lib.Members {
			if strings.Contains(strings.ToLower(m.Name), strings.ToLower(input)) {
				matchedMembers = append(matchedMembers, m)
			}
		}
	}
	if len(matchedMembers) == 0 {
		fmt.Println("❌ Member not found.")
		return
	}
	for _, member := range matchedMembers {
		fmt.Printf("\nLoans for Member: %s (ID: %d)\n", member.Name, member.ID)
		found := false
		for _, loan := range lib.Loans {
			if loan.MemberID == member.ID {
				bookTitle := "Unknown"
				for _, book := range lib.Books {
					if book.ID == loan.BookID {
						bookTitle = book.Title
						break
					}
				}
				fmt.Printf("- Book: %s | Date Out: %s | Date In: %s\n", bookTitle, loan.DateOut, loan.DateIn)
				found = true
			}
		}
		if !found {
			fmt.Println("No loans found.")
		}
	}
}
func FilterLoansByBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Filter Loans > Filter Loans By Book ===")
	fmt.Print("Enter Book Title or ISBN: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var matchedBooks []models.Book
	for _, book := range lib.Books {
		if strings.EqualFold(book.ISBN, input) || strings.Contains(strings.ToLower(book.Title), strings.ToLower(input)) {
			matchedBooks = append(matchedBooks, book)
		}
	}

	if len(matchedBooks) == 0 {
		fmt.Println("❌ Book not found.")
		return
	}

	for _, book := range matchedBooks {
		fmt.Printf("\nLoans for Book: %s (ID: %d)\n", book.Title, book.ID)
		found := false
		for _, loan := range lib.Loans {
			if loan.BookID == book.ID {
				memberName := "Unknown"
				for _, member := range lib.Members {
					if member.ID == loan.MemberID {
						memberName = member.Name
						break
					}
				}
				fmt.Printf("- Member: %s | Date Out: %s | Date In: %s\n", memberName, loan.DateOut, loan.DateIn)
				found = true
			}
		}
		if !found {
			fmt.Println("No loans found.")
		}
	}

}
func FilterLoansByReturnStatus(lib *models.Library) {
	fmt.Println("\n=== library Menu > Manage Loans > Filter Loans > Filter Loans By Return Status ===")
	found := false
	for _, loan := range lib.Loans {
		if strings.TrimSpace(loan.DateIn) == "" {
			memberName := "Unknown"
			bookTitle := "Unknown"
			for _, member := range lib.Members {
				if member.ID == loan.MemberID {
					memberName = member.Name
					break
				}
			}
			for _, book := range lib.Books {
				if book.ID == loan.BookID {
					bookTitle = book.Title
					break
				}
			}
			fmt.Printf("- Book: %s | Member: %s | Date Out: %s | Status: ❌ Not Returned\n", bookTitle, memberName, loan.DateOut)
			found = true
		}
	}
	if !found {
		fmt.Println("✅ All books are returned.")
	}
}
func FilterLoansByDateOut(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Filter Loans > Filter Loans By Date Out ===")
	fmt.Print("Enter Date Out (e.g., 2025-05-09): ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)
	found := false
	for _, loan := range lib.Loans {
		if loan.DateOut == date {
			memberName := "Unknown"
			bookTitle := "Unknown"
			for _, member := range lib.Members {
				if member.ID == loan.MemberID {
					memberName = member.Name
					break
				}
			}
			for _, book := range lib.Books {
				if book.ID == loan.BookID {
					bookTitle = book.Title
					break
				}
			}
			fmt.Printf("- Book: %s | Member: %s | Date In: %s\n", bookTitle, memberName, loan.DateIn)
			found = true
		}
	}
	if !found {
		fmt.Println("❌ No loans found for the given date.")
	}
}
func FindLoan(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Find Loan ===")
	if len(lib.Loans) == 0 {
		fmt.Println("📭 No loans found.")
		return
	}
	fmt.Print("Enter Loan ID to find: ")
	loanIDStr, _ := reader.ReadString('\n')
	loanIDStr = strings.TrimSpace(loanIDStr)
	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		fmt.Println("❌ Please enter a valid Loan ID!")
		return
	}
	found := false
	for _, loan := range lib.Loans {
		if loan.BookID == loanID || loan.MemberID == loanID {
			fmt.Printf("Found Loan: Book ID: %d, Member ID: %d, Date Out: %s, Date In: %s\n", loan.BookID, loan.MemberID, loan.DateOut, loan.DateIn)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("❌ Loan not found!")
	}

}
func EditLoan(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Edit Loan ===")
	if len(lib.Loans) == 0 {
		fmt.Println("📭 No loans found.")
		return
	}
	fmt.Print("Enter Loan ID to edit: ")
	loanIDStr, _ := reader.ReadString('\n')
	loanIDStr = strings.TrimSpace(loanIDStr)
	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		fmt.Println("❌ Please enter a valid Loan ID!")
		return
	}

	loanExist := false
	for i, loan := range lib.Loans {
		if loan.BookID == loanID || loan.MemberID == loanID {
			loanExist = true
			fmt.Println("*** Current Loan Info ***")
			fmt.Printf("Book ID: %d, Member ID: %d, Date Out: %s, Date In: %s\n", loan.BookID, loan.MemberID, loan.DateOut, loan.DateIn)
			fmt.Print("New Date In (leave empty for no change): ")
			newDateIn, _ := reader.ReadString('\n')
			newDateIn = strings.TrimSpace(newDateIn)
			if newDateIn != "" {
				lib.Loans[i].DateIn = newDateIn
			}
			fmt.Println("✅ Loan updated successfully.")
			break
		}
	}
	if !loanExist {
		fmt.Println("❌ Loan not found!")
	}
}
func DeleteLoan(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Loans > Delete Loan ===")
	if len(lib.Loans) == 0 {
		fmt.Println("📭 No loans found.")
		return
	}
	fmt.Print("Enter Loan ID to delete: ")
	loanIDStr, _ := reader.ReadString('\n')
	loanIDStr = strings.TrimSpace(loanIDStr)
	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		fmt.Println("❌ Please enter a valid Loan ID!")
		return
	}

	loanExist := false
	for i, loan := range lib.Loans {
		if loan.BookID == loanID || loan.MemberID == loanID {
			// Delete the loan by removing it from the slice
			lib.Loans = append(lib.Loans[:i], lib.Loans[i+1:]...)
			loanExist = true
			fmt.Println("✅ Loan deleted successfully.")
			break
		}
	}

	if !loanExist {
		fmt.Println("❌ Loan not found!")
	}
}
