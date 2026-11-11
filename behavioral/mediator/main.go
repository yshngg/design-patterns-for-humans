package main

import (
	"fmt"
	"time"
)

type User struct {
	name         string
	chatMediator ChatRoomMediator
}

func NewUser(name string, mediator ChatRoomMediator) *User {
	return &User{
		name:         name,
		chatMediator: mediator,
	}
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Send(message string) {
	u.chatMediator.ShowMessage(u, message)
}

type ChatRoomMediator interface {
	ShowMessage(user *User, message string)
}

// 中介者 (Mediator)
type ChatRoom struct{}

func (r *ChatRoom) ShowMessage(user *User, message string) {
	now := time.Now()
	sender := user.Name()
	fmt.Printf("%s [%s]: %s\n", now, sender, message)
}

func NewChatRoom() ChatRoomMediator {
	return &ChatRoom{}
}

func main() {
	mediator := NewChatRoom()
	john := NewUser("John Doe", mediator)
	jane := NewUser("Jane Doe", mediator)

	john.Send("Hi there!")
	jane.Send("Hey!")

	// 输出结果
	// 2026-11-11 17:15:38.520103 +0800 CST m=+0.000132810 [John Doe]: Hi there!
	// 2026-11-11 17:15:38.535787 +0800 CST m=+0.015816074 [Jane Doe]: Hey!
}
