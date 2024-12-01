package main

import "fmt"

// OnlineContacts struct to represent the contacts
type OnlineContacts struct {
	NoOfOnlineContacts int
	NoOfContacts       int
}

// NewOnlineContacts initializes a new instance of OnlineContacts
func NewOnlineContacts() *OnlineContacts {
	return &OnlineContacts{
		NoOfOnlineContacts: 0,
		NoOfContacts:       10,
	}
}

// GetNoOfOnlineContacts returns the number of online contacts
func (oc *OnlineContacts) GetNoOfOnlineContacts() int {
	return oc.NoOfOnlineContacts
}

// SetNoOfOnlineContacts sets the number of online contacts with validation
func (oc *OnlineContacts) SetNoOfOnlineContacts(noOfOnlineContacts int) {
	if noOfOnlineContacts < 0 {
		oc.NoOfOnlineContacts = 0
	} else if noOfOnlineContacts > oc.NoOfContacts {
		oc.NoOfOnlineContacts = oc.NoOfContacts
	} else {
		oc.NoOfOnlineContacts = noOfOnlineContacts
	}
}

// Main function for testing
func main() {
	// Example usage
	contacts := NewOnlineContacts()

	fmt.Println("Initial online contacts:", contacts.GetNoOfOnlineContacts())

	// Set valid number of online contacts
	contacts.SetNoOfOnlineContacts(5)
	fmt.Println("Online contacts after setting to 5:", contacts.GetNoOfOnlineContacts())

	// Set invalid (negative) number of online contacts
	contacts.SetNoOfOnlineContacts(-3)
	fmt.Println("Online contacts after setting to -3:", contacts.GetNoOfOnlineContacts())

	// Set number of online contacts exceeding the total
	contacts.SetNoOfOnlineContacts(12)
	fmt.Println("Online contacts after setting to 12:", contacts.GetNoOfOnlineContacts())
}
