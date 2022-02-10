/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func shellExec(cmd string) {
	shell := exec.Command("bash", "-c", cmd)

	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	shell.Run()
}

func getAllIps() []string {
	content, err := ioutil.ReadFile("ips")

	if err != nil {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	var ips []string = nil

	for _, line := range lines {
		ip := strings.Split(line, "\t")

		if (len(ip) > 1) {
			ips = append(ips, ip[1])
		}
	}

	return ips
}
