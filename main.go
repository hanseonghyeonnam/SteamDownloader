package main

import (
	"github.com/pkg/browser"
	"os"
	"fmt"
)

func main() {
	fmt.Println("[This is a security-awareness demo, not a game unlocker.]")
	browser.OpenURL("https://pastebin.com/eQETzdvW")
	os.Exit(1)
}