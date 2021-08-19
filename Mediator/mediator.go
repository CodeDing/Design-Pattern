package Mediator

import (
    `fmt`
    `time`
)

type User struct {
    name string
}

func (u User) GetName() string {
    return u.name
}

func (u *User) SetName(n string) {
    u.name = n
}

func (u User) SendMessage(m string) {
    DefaultChatRoom.ShowMessage(&u, m)
}

var DefaultChatRoom = chatroom{"000"}

type chatroom struct {
    id string
}

func (c chatroom) ShowMessage(u *User, m string) {
    fmt.Printf("[%s] name:%s, msg=>%s\n", time.Now().Format(time.RFC3339), u.GetName(), m)
}
