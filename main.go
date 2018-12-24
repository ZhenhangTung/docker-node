package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	args := []string{
		"poster.js",
	}
	cmd := exec.Command("node", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed. err: %v", err)
		return
	}
	fmt.Println("ok")
}
