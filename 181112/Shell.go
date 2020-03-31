package main

import (
	"log"
	"os"
	"os/exec"
)func main() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	log.Printf("Running command and waiting for it to finish...")
	cmd.Run()
	//log.Printf("Command finished with error: %v", err)
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	log.Printf("Running command and waiting for it to finish...")
	cmd.Run()
	//log.Printf("Command finished with error: %v", err)
}

// Platform: runtime.GOOS
