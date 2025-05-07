package models

type Loan struct {
	LoanID   int
	BookID   int
	MemberID int
	DateOut  string
	DateIn   string
}
