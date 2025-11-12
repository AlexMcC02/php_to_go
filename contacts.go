package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"slices"
	"strconv"
)

type JsonData struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Telephone string `json:"telephone"`
}

func loadJson() []Contact {
	const path = "./files/contacts.json"
	jsonFile, err := os.Open(path)
	
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	var data JsonData
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)

	return data.Contacts
}

func prettyPrint(contact Contact) {
	fmt.Printf("\nName: %s\nEmail: %s\nTelephone Num: %s\n", contact.Name, contact.Email, contact.Telephone)
}

func getContactIndexFromContactName(name string, contacts *[]Contact) int {
	var index = -1

	for i := 0; i < len(*contacts); i++ {
		if (*contacts)[i].Name == name {
			index = i
			break
		}
	}

	return index
}

func createContact(contacts *[]Contact) {
	var name, email, telephone string
	var reader = bufio.NewReader(os.Stdin)
	
	fmt.Print("\nEnter the name for the new contact:\n\n> ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("\nEnter the email for the new contact:\n\n> ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	fmt.Print("\nEnter the telephone number for the new contact:\n\n> ")
	telephone, _ = reader.ReadString('\n')
	telephone = strings.TrimSpace(telephone)

	newContact := Contact{Name: name, Email: email, Telephone: telephone}

	*contacts = append(*contacts, newContact)

	if getContactIndexFromContactName(name, contacts) == -1 {
		fmt.Print("\nFailed to create new contact.\n")
	} else {
		fmt.Print("\nNew contact created successfully.\n")
		prettyPrint(newContact)
	}
}

func deleteContactByName(contacts *[]Contact) {
	var nameToDelete string
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the name of the contact you wish to delete:\n\n> ")
	nameToDelete, _ = reader.ReadString('\n')
	nameToDelete = strings.TrimSpace(nameToDelete)
	contactId := getContactIndexFromContactName(nameToDelete, contacts)

	if contactId == -1 {
		fmt.Print("\nThe contact was not found.\n")
		return
	}

	fmt.Printf("\n%s was found!\n", nameToDelete)
	prettyPrint((*contacts)[contactId])

	var decision int
	fmt.Printf("\nAre you sure you want to delete %s?\n\n1. Yes.\n2. No.\n\n> ", nameToDelete)
	decisionString, _ := reader.ReadString('\n')
	decision, _ = strconv.Atoi(decisionString)

	if decision == 2 { return }

	*contacts = slices.Delete(*contacts, contactId, contactId + 1)
}

func commitChanges(contacts *[]Contact) {
	const path = "./files/contacts.json"

	var data = JsonData{Contacts: *contacts}
	var jsonData, _ = json.MarshalIndent(data, " ", "  ")
	var jsonFile, _ = os.Create(path)

	defer jsonFile.Close()

	_, err := jsonFile.Write(jsonData)

	if err != nil {
		panic(err)
	}

	err = jsonFile.Sync()
	if err != nil {
		panic(err)
	}
}

func displayMenu() {
	menu := "\nPlease select a course of action:\n\n" +
			"1. Display contacts.\n" +
			"2. Add contact.\n" +
			"3. Remove contact.\n" +
			"4. Commit changes.\n" +
			"5. Quit.\n\n" +
			"> "
    fmt.Print(menu)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func contactsChoicer(running *bool, contacts *[]Contact) {
	var choice int
	displayMenu()
	fmt.Scan(&choice);

	switch choice {
	case 1:
		for _, contact := range *contacts {
			prettyPrint(contact)
		}
	case 2:
		createContact(contacts)
	case 3:
		deleteContactByName(contacts)
	case 4:
		commitChanges(contacts)
	case 5:
		*running = false
	}
}

func main() {
	var contacts = loadJson()
	var running = true
	clearScreen()

	for running {
		contactsChoicer(&running, &contacts)
	}
}