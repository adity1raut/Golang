package main

import (
	"fmt"
)

type Notifier interface {
	Notify() (string, error)
}

type Email struct {
	email string
}

type SMS struct {
	num int
}

type Voice struct {
	voice string
}

func (e Email) Notify() (string, error) {
	if e.email == "" {
		return "", fmt.Errorf("email is empty")
	}
	return fmt.Sprintf("Sending Email to %s", e.email), nil
}

func (s SMS) Notify() (string, error) {
	if s.num == 0 {
		return "", fmt.Errorf("phone number is empty")
	}
	return fmt.Sprintf("You have Message from %d", s.num), nil
}

func (v Voice) Notify() (string, error) {
	if v.voice == "" {
		return "", fmt.Errorf("voice message is empty")
	}
	return fmt.Sprintf("You have Voice Message: %s", v.voice), nil
}

func main() {
	var n Notifier

	n = Email{""}
	result, err := n.Notify()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
