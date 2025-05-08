package services

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/bigbug-ir/Library-system-TUI/models"
)

func ManageBooks(lib *models.Library, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== library Menu > Manage Books ===")
		fmt.Println("0.<-Back")
		fmt.Println("1.Add Book->")
		fmt.Println("2.List Books->")
		fmt.Println("3.Filter Books->")
		fmt.Println("4.Find Books->")
		fmt.Println("5.Edit Book->")
		fmt.Println("6.Delete Book->")

		fmt.Print("Slect Item : ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Please enter a Number")
			continue
		}
		switch choice {
		case 0:
			fmt.Println("<-Back Page")
			return
		case 1:
			AddBook(lib, reader)
		case 2:
			ListBooks(lib)
		case 3:
			FilterBooks(lib, reader)
		case 4:
			FindBooks(lib, reader)
		case 5:
			EditBook(lib, reader)
		case 6:
			DeleteBook(lib, reader)
		default:
			fmt.Println("❌Item is not available!")
		}
	}
}
func AddBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Add Book ===")
	fmt.Print("Book Title :")
	title, _ := reader.ReadString('\n')
	fmt.Print("Book Author :")
	author, _ := reader.ReadString('\n')
	fmt.Print("Book ISBN : ")
	isbn, _ := reader.ReadString('\n')
	book := models.Book{
		Title:     strings.TrimSpace(title),
		Author:    strings.TrimSpace(author),
		ISBN:      strings.TrimSpace(isbn),
		Available: true,
		ID:        len(lib.Books) + 1,
	}
	lib.Books = append(lib.Books, book)
	fmt.Println("✅Successfully added book")
}
func ListBooks(lib *models.Library) {
	fmt.Println("\n=== library Menu > Manage Books > Books List ===")
	if len(lib.Books) == 0 {
		fmt.Println("📭 No books found.")
		return
	}
	for _, book := range lib.Books {
		fmt.Printf("- %d: %s - %s  -- %s -- %v \n", book.ID, book.Title, book.Author, book.ISBN, book.Available)
	}
}
func FilterBooks(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Filter Books ===")
	if len(lib.Books) == 0 {
		fmt.Println("📭 No books found.")
		return
	}
	for {
		fmt.Println("0.<-back")
		fmt.Println("1.Filter By Title")
		fmt.Println("2.Filter By Author")
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
			FilterBooksByTitle(lib, reader)
		case 2:
			FilterBooksByAuthor(lib, reader)
		default:
			fmt.Println("❌ Invalid option.")
		}
	}
}
func FilterBooksByTitle(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Filter Books > Filter By Title ===")
	fmt.Print("Enter Book Title to filter: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	var result []models.Book
	for _, book := range lib.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			result = append(result, book)
		}
	}
	fmt.Println("*** Filter Result *** ")
	if len(result) == 0 {
		fmt.Println("❌ Book not found!")
		return
	}
	for _, r := range result {
		fmt.Printf("- %d: %s - %s  -- %s -- %v \n", r.ID, r.Title, r.Author, r.ISBN, r.Available)
	}
}
func FilterBooksByAuthor(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Filter Books > Filter By Author ===")
	fmt.Print("Enter Book Author to filter: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	var result []models.Book
	for _, book := range lib.Books {
		if strings.Contains(strings.ToLower(book.Author), strings.ToLower(author)) {
			result = append(result, book)
		}
	}
	fmt.Println("*** Filter Result ***")
	if len(result) == 0 {
		fmt.Println("❌ Book not found!")
		return
	}
	for _, r := range result {
		fmt.Printf("- %d: %s - %s  -- %s -- %v \n", r.ID, r.Title, r.Author, r.ISBN, r.Available)
	}
}
func FindBooks(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Find Book ===")
	if len(lib.Books) == 0 {
		fmt.Println("📭 No books found.")
		return
	}
	for {
		fmt.Println("0.<-back")
		fmt.Println("1.Find Book By ID")
		fmt.Println("2.find Book By ISBN")

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
			FindBooksByID(lib, reader)
		case 2:
			FindBooksByISBN(lib, reader)
		default:
			fmt.Println("❌ Invalid option.")
		}
	}
}
func FindBooksByID(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Find Book > Find Book By ID ===")
	fmt.Print("Enter Book ID to find: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("❌ Please enter a valid number ID.")
		return
	}
	fmt.Println("*** Search Result ***")
	for _, book := range lib.Books {
		if book.ID == ID {
			fmt.Printf("- %d: %s - %s  -- %s -- %v \n", book.ID, book.Title, book.Author, book.ISBN, book.Available)
			return
		}
	}
	fmt.Println("❌book not found!")
}
func FindBooksByISBN(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Find Book > Find Book By ISBN ===")
	fmt.Print("Enter Edit Book ISBN to Filter: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)
	fmt.Println("*** Search Reasult ***")
	for _, book := range lib.Books {
		if book.ISBN == isbn {
			fmt.Printf("- %d: %s - %s  -- %s -- %v \n", book.ID, book.Title, book.Author, book.ISBN, book.Available)
			return
		}
	}
	fmt.Println("❌book not found!")
}
func EditBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Edit Book ===")
	if len(lib.Books) == 0 {
		fmt.Println("📭 No books found.")
		return
	}
	fmt.Print("Enter Book ISBN to edit: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)
	for i, book := range lib.Books {
		if book.ISBN == isbn {
			fmt.Println("*** Your book before editing ***")
			fmt.Println("*** for no changing press enter ***")
			fmt.Printf("- %d: %s - %s  -- %s -- %v \n", book.ID, book.Title, book.Author, book.ISBN, book.Available)
			fmt.Printf("New Author (old: %s) :", lib.Books[i].Author)
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			if author != "" {
				lib.Books[i].Author = author
			}
			fmt.Printf("New Title (old: %s) : ", book.Title)
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title != "" {
				lib.Books[i].Title = title
			}
			fmt.Printf("New ISBN (old: %s)", book.ISBN)
			newIsbn, _ := reader.ReadString('\n')
			newIsbn = strings.TrimSpace(newIsbn)
			if newIsbn != "" {
				lib.Books[i].ISBN = newIsbn
			}
			fmt.Println("✅ Book updated successfully.")
			return
		}
	}
	fmt.Println("❌Book not found!")
}
func DeleteBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Println("\n=== library Menu > Manage Books > Delete Book ===")
	if len(lib.Books) == 0 {
		fmt.Println("📭 No books found.")
		return
	}
	fmt.Print("Enter Book ID to delete : ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("❌ Please enter a valid number ID.")
		return
	}
	for i, book := range lib.Books {
		if book.ID == ID {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Println("✅ Book deleted successfully.")
			return
		}
	}
	fmt.Println("❌book not found!")
}
