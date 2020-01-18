package models

type Message struct {
	ID         string
	Text       string
	ChatroomID string 
	Chatroom   Chatroom
}
