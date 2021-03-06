package main

import "os/exec"
import "syscall"
import "os"

func main()  {
	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
