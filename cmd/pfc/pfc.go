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
	"os/signal"
	"regexp"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/utils/exec"

	"github.com/max0ne/kubectl-pfc/pkg/curlflags"
	"github.com/max0ne/kubectl-pfc/pkg/url"
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
func runPortForward(args []string) (exec.Cmd, string, string, error) {
	cmd := exec.New().Command("kubectl", append([]string{"port-forward"}, args...)...)
	log.Debugf("Executing port-forward with args %v", args)

	cmd.SetStderr(os.Stderr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, "", "", err
	}
	if err := cmd.Start(); err != nil {
		return nil, "", "", err
	}

	reader := bufio.NewReader(stdout)
	regex := regexp.MustCompile("^Forwarding from (.+?):(\\d+) ->")
	for {
		lineBytes, err := reader.ReadBytes('\n')
		if err != nil {
			return nil, "", "", err
		}
		matches := regex.FindSubmatch(lineBytes)
		if len(matches) > 2 {
			return cmd, string(matches[1]), string(matches[2]), nil
		}
	}
}

func runCurl(command string, args []string, localHost, localPort string) error {
	urlFlagIndex := curlflags.IndexOfURLArg(args)
	if urlFlagIndex == -1 {
		return fmt.Errorf("unable to find url flag in curl args %v", args)
	}

	originalURL := args[urlFlagIndex]
	scheme, host, port, query, err := url.Parse(originalURL)
	if err != nil {
		return errors.Wrapf(err, "unable to parse url %s", originalURL)
	}
	log.Debugf("Original URL %s parsed to %+v", originalURL, []string{scheme, host, port, query})
	if len(port) > 0 {
		return fmt.Errorf("curl url does not allow specifying port, url: %s, port: %s", originalURL, port)
	}

	targetHost := fmt.Sprintf("%s%s:%s%s", scheme, host, localPort, query)
	resolveArg := fmt.Sprintf("%s:%s:%s", host, localPort, localHost)

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

	portForwardCmd, localHost, localPort, err := runPortForward(portForwardArgs)
	if err != nil {
		return err
	}

	// Cleanup port-forward when main process finishes
	defer portForwardCmd.Stop()

	// Cleanup port-forward when the main process is terminated
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		portForwardCmd.Stop()
	}()

	return runCurl(curlCommand, curlArgs, localHost, localPort)
}
