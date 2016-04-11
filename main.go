package main

import (
	"./encrypt"
)

func main() {
	const question = "Write what you want to hash:"
	encrypt.HashIt(encrypt.readFromUser(question))
}
