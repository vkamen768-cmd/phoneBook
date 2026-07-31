package phonebook

import (
	"errors"
	"fmt"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func NewContact(name, phone string) (*Contact, error) {
	if name == "" || phone == "" {
		return nil, errors.New("empty name or phone number")
	}
	return &Contact{
		Name:  name,
		Phone: phone,
	}, nil
}

func (c Contact) Print() {
	fmt.Printf("Name: %s | Phone: %s\n", c.Name, c.Phone)
}
