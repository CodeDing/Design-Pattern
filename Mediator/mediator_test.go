package Mediator

import `testing`

func TestMediator(t *testing.T) {
    robert := &User{}
    robert.SetName("Robert")
    john := &User{}
    john.SetName("John")

    robert.SendMessage("Hi, John!")
    john.SendMessage("Hello, Robert!")
}
