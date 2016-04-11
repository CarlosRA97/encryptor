package encrypt

import "fmt"

// ReadFromUser Reads the input from user command line
func ReadFromUser(question string) string {
	var answer string
	str1 := question
	fmt.Println(str1)
	fmt.Scanf("%s", &answer)
	return answer
}
