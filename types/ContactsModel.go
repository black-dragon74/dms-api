package types

type ContactsModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Image       string `json:"image"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
}
