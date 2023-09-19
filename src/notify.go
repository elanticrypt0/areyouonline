package src

import "fmt"

// Notifica por email a uno o varios destinarios.

type Notify struct {
	DestinyAddresses string
	Message          string
}

func NewNotify(destinyAddresses, msg string) *Notify {
	return &Notify{
		DestinyAddresses: destinyAddresses,
		Message:          msg,
	}
}

func (n *Notify) notify(msg string) {
	n.showMsgInConsole(msg)
	n.sendEmail(msg)

}

func (n *Notify) sendEmail(msg string) {
	// TODO
}

func (n *Notify) showMsgInConsole(msg string) {
	fmt.Println(msg)
}
