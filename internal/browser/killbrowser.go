package browser

import (
	"fmt"
	"os/exec"
)

func KillEdgeProcess() error {
	return killProcess("msedge.exe")
}

func KillChromeProcess() error {
	return killProcess("chrome.exe")
}

func KillBraveProcess() error {
	return killProcess("brave.exe")
}

func KillFirefoxProcess() error {
	return killProcess("firefox.exe")
}

func killProcess(processName string) error {
	cmd := exec.Command("taskkill", "/IM", processName, "/F")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[-] error killing %s process: %v", processName, err)
	}
	return nil
}
