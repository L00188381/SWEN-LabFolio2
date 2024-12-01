package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNoOfOnlineContacts(t *testing.T) {
	contacts := NewOnlineContacts()
	onlineContacts := contacts.GetNoOfOnlineContacts()

	assert.Equal(t, onlineContacts, 0)
}

func TestSetNoOfOnlineContacts(t *testing.T) {
	contacts := NewOnlineContacts()
	contacts.SetNoOfOnlineContacts(5)
	onlineContacts := contacts.GetNoOfOnlineContacts()

	assert.Equal(t, onlineContacts, 5)
}

func TestSetNegativeNoOfOnlineContacts(t *testing.T) {
	contacts := NewOnlineContacts()
	contacts.SetNoOfOnlineContacts(-5)
	onlineContacts := contacts.GetNoOfOnlineContacts()

	assert.Equal(t, onlineContacts, 0)
}

func TestSetMaxNoOfOnlineContacts(t *testing.T) {
	contacts := NewOnlineContacts()
	contacts.SetNoOfOnlineContacts(20)
	onlineContacts := contacts.GetNoOfOnlineContacts()

	assert.Equal(t, onlineContacts, 10)
}
