package main

import (
	"fmt"
	"phoneBook/phonebook"
)

func main() {
	pb := phonebook.NewPhonebook()
	fmt.Println("--- Телефонная Книга ---")
Menu:
	for {
		switch ch := getMenu(); ch {
		case 1:
			createContact(pb)
		case 2:
			findContact(pb)
		case 3:
			deleteContact(pb)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	fmt.Println(`1. Добавить контакт
2. Найти контакт
3. Удалить контакт
4. Выход`)
	var input int
	fmt.Print("Enter: ")
	fmt.Scan(&input)
	return input
}

func promptData(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}

func createContact(pb *phonebook.Phonebook) {
	name := promptData("Name")
	phone := promptData("Phone")
	contact, err := phonebook.NewContact(name, phone)
	if err != nil {
		fmt.Println(err)
		return
	}
	pb.AddContact(*contact)
	fmt.Println("Contact successfully added!")
}

func findContact(pb *phonebook.Phonebook) {
	name := promptData("Name for find")
	contacts := pb.FindContactsByName(name)
	if len(contacts) == 0 {
		fmt.Println("Contacts not find")
		return
	}
	for _, v := range contacts {
		v.Print()
	}
}

func deleteContact(pb *phonebook.Phonebook) {
	name := promptData("Name for delete")
	isDeleted := pb.DeleteContactByName(name)
	if isDeleted {
		fmt.Println("Contact deleted")
		return
	}
	fmt.Println("Contact not found")
}
