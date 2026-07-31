package phonebook

import (
	"encoding/json"
	"fmt"
	"phoneBook/files"
	"strings"
)

type Phonebook struct {
	Contacts []Contact `json:"contacts"`
}

func NewPhonebook() *Phonebook {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Phonebook{
			Contacts: make([]Contact, 0),
		}
	}
	var pb Phonebook
	err = json.Unmarshal(file, &pb)
	if err != nil {
		fmt.Println(err)
		return &Phonebook{
			Contacts: make([]Contact, 0),
		}
	}
	return &pb
}

func (p *Phonebook) AddContact(contact Contact) {
	p.Contacts = append(p.Contacts, contact)
	p.save()
}

func (p *Phonebook) FindContactsByName(name string) []Contact {
	var contacts []Contact
	for _, contact := range p.Contacts {
		if strings.Contains(contact.Name, name) {
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

func (p *Phonebook) DeleteContactByName(name string) bool {
	var updatedContacts []Contact
	var isDeleted bool
	for _, contact := range p.Contacts {
		if strings.Contains(contact.Name, name) {
			isDeleted = true
			continue
		}
		updatedContacts = append(updatedContacts, contact)
	}
	p.Contacts = updatedContacts
	p.save()
	return isDeleted
}

func (p *Phonebook) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Phonebook) save() {
	data, err := p.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.WriteFile(data, "data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
}
