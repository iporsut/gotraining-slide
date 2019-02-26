package main

import "fmt"

type Sender interface {
	Send()
}

type Mail struct {
	FromAddr string
	ToAddr   string
	Msg      string
}

func (m *Mail) Send() {
	fmt.Println("Sent", m.Msg, "from", m.FromAddr, "to", m.ToAddr)
}

func main() {
	var sender Sender // default value of interface is nil
	mail := &Mail{
		FromAddr: "my@mail.com",
		ToAddr:   "you@mail.com",
		Msg:      "Hello",
	}
	sender = mail
	sender.Send()
}
