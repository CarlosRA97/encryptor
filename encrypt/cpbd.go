package encrypt

import (
	"./sh"
	"fmt"
)

func clipboardIt(data string) {
	commands := [2]string{fmt.Sprintf("%s", data), "pbcopy"}
	sh.Command(commands[len(commands)-1]).SetInput(commands[0]).Output()
}
