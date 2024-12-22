package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Form struct {
	prompt string
	name   string
}

func createForm(forms []Form) map[string]string {
	reader := bufio.NewReader(os.Stdin)
	body := map[string]string{}
	for _, form := range forms {
		fmt.Print(form.prompt)
		nominal, _ := reader.ReadString('\n')
		body[form.name] = strings.Trim(nominal, "\n\r")
	}

	return body
}
