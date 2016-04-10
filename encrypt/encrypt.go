package encrypt

import (
	"crypto/sha1"
	"fmt"
	"os/exec"
)

// HashIt This funcion hash the string given
// and send it directly to the clipboard
func HashIt(text string) {
	data := []byte(text)
	hash := fmt.Sprintf("%x\n", sha1.Sum(data))
	// fmt.Println(hash)
	// clipboardIt(hash)
	return hash
}

func readFromUser() string {
	var myname string
	str1 := "What you want to hash?"
	fmt.Println(str1)
	fmt.Scanf("%s", &myname)
	// fmt.Println("Hello", myname)
	return myname
}

func clipboardIt(data string) {
	// commands := [4]string{"echo ", data, " | ", "pbcopy"}
	// fmt.Println(co)
	command := fmt.Sprintf("echo '%s' | pbcopy", data)
	fmt.Println(command)
	exec.Command(command).Start()
}
