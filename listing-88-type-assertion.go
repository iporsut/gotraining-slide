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
	var v interface{}
	v = "Hello"
	s := v.(string)
	fmt.Println(s)

	// i := v.(int) // panic
	// fmt.Println(i)

	var sender Sender
	sender = &Mail{
		FromAddr: "my@mail.com",
		ToAddr:   "you@mail.com",
		Msg:      "Hello",
	}
	m := sender.(*Mail)
	fmt.Println(m.FromAddr)
}
