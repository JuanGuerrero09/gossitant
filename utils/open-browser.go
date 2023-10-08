package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Open(urls []string) error {
  var osConfig string
	switch runtime.GOOS{
	case "linux":
    osConfig = "sensible-browser"
		cmd := exec.Command(osConfig)
    cmd.Args = append(cmd.Args, "--new-window")
		startProcess(cmd, urls...)
		fmt.Println(cmd)
  default:
    return fmt.Errorf("unsupported platform")
  }
	return nil
}

func startProcess(cmd *exec.Cmd, urls ...string) {
	cmd.Args = append(cmd.Args, urls...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
