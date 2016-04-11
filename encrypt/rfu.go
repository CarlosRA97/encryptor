package encrypt

import "fmt"

func readFromUser(question string) string {
	var answer string
	str1 := question
	fmt.Println(str1)
	fmt.Scanf("%s", &answer)
	// fmt.Println("Hello", answer)
	return answer
}
