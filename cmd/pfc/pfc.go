/*
Copyright © 2021 Mingfei Huang <himax1023@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pfc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
	"k8s.io/utils/exec"

	"github.com/max0ne/kubectl-pfc/pkg/curlflags"
)

func parseArgs(args []string) (portForwardArgs []string, curlCommand string, curlArgs []string, err error) {
	for idx, arg := range args {
		if arg == "--" {
			portForwardArgs = args[:idx]
			if idx >= len(args)-2 {
				return nil, "", nil, fmt.Errorf("require curl command and arguments after `--`")
			}
			curlCommand = args[idx+1]
			curlArgs = args[idx+2:]
			return
		}
	}
	return nil, "", nil, fmt.Errorf("unable to find arg separator `--`")
}

// runPortForward runs port-forward with given args, returns the local host and port it's using
func runPortForward(args []string) (string, string, error) {
	cmd := exec.New().Command("kubectl", append([]string{"port-forward"}, args...)...)
	log.Debugf("Executing port-forward with args %v", args)

	cmd.SetStderr(os.Stderr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}
	go cmd.Run()

	reader := bufio.NewReader(stdout)
	for {
		lineBytes, err := reader.ReadBytes('\n')
		if err != nil {
			return "", "", err
		}

		regex := regexp.MustCompile("^Forwarding from (.+?):(\\d+) ->")
		matches := regex.FindSubmatch(lineBytes)
		if len(matches) > 2 {
			return string(matches[1]), string(matches[2]), nil
		}
	}
}

func runCurl(command string, args []string, localHost, localPort string) error {
	urlFlagIndex := curlflags.IndexOfURLArg(args)
	if urlFlagIndex == -1 {
		return fmt.Errorf("unable to find url flag in curl args %v", args)
	}

	originalHost := args[urlFlagIndex]
	targetHost := fmt.Sprintf("%s:%s", originalHost, localPort)
	resolveArg := fmt.Sprintf("%s:%s:%s", originalHost, localPort, localHost)

	curlArgs := args[:]
	curlArgs[urlFlagIndex] = targetHost
	curlArgs = append(curlArgs, "--resolve", resolveArg)

	cmd := exec.New().Command(command, curlArgs...)
	cmd.SetStdout(os.Stdout)
	cmd.SetStderr(os.Stderr)
	log.Debugf("Executing curl with args %v", curlArgs)

	return cmd.Run()
}

func Run() error {
	portForwardArgs, curlCommand, curlArgs, err := parseArgs(os.Args[1:])
	if err != nil {
		return err
	}

	localHost, localPort, err := runPortForward(portForwardArgs)
	if err != nil {
		return err
	}

	err = runCurl(curlCommand, curlArgs, localHost, localPort)
	return err
}
