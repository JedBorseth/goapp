package initializers

import (
	"fmt"
	"os/exec"
)
func TailwindCompile() {
	cmd := exec.Command("./tailwindcss", "-i", "styles/input.css", "-o", "styles/output.css", "--minify")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}