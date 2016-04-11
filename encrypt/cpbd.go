package encrypt

import (
	"./sh"
	"fmt"
)

// ClipboardIt Copy the data given to the clipboard by shell
func ClipboardIt(data string) {
	commands := [2]string{fmt.Sprintf("%s", data), "pbcopy"}
	sh.Command(commands[len(commands)-1]).SetInput(commands[0]).Output()
}
