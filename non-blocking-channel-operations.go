package main 

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("revieved message:", msg)
	default:
		fmt.Println("no message recieved")
	}

	msg := "hi"
	
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("recevied signals", sig)
	default:
		fmt.Println("no activity")
	}
}