package main

import (
	"fmt"
	"os/exec"
	"bytes"
)

func main()  {
	args := []string{
		"index.js",
	}
	cmd := exec.Command("node", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed. err: %v", stderr.String())
		return
	}
	fmt.Println("ok")
}
