package exec

import (
	"fmt"
	"os/exec"
)

func Run(cn string, error string) {
	cmd, err := exec.Command("./exec.sh", cn, error).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(cmd))
	}
}
