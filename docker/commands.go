package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func run(cmdname string, args string, showoutput bool) error {
	if strings.HasPrefix(strings.ToLower(cmdname), "docker") {
		return errors.New("Don't start with `docker`, start with the command, usually `run`")
	}
	cmdStr := fmt.Sprintf("docker %s %s", cmdname, args)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	if showoutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}

//StopContainer takes the name of a docker container and stops it, or returns an error.
func StopContainer(name string, showoutput bool) error {
	return run("stop", name, showoutput)
}

//InspectContainer takes in the name of a container and returns a ContainerInfo instance with data about the container.
func InspectContainer(name string) (ContainerInfo, error) {
	cmdStr := fmt.Sprintf("docker inspect %s", name)
	var info []ContainerInfo
	data, errExec := exec.Command("/bin/sh", "-c", cmdStr).CombinedOutput()
	if strings.Contains(string(data), fmt.Sprintf("Error: No such image or container: %s", name)) {
		return ContainerInfo{}, NewContainerNotFound(name)
	}
	if errExec != nil {
		return ContainerInfo{}, errExec
	}
	err := json.Unmarshal(data, &info)
	if err != nil {
		return ContainerInfo{}, err
	}
	return info[0], nil
}

//RemoveContainer takes the name of a docker container and stop it or returns an error.
func RemoveContainer(name string, showoutput bool) error {
	return run("rm", name, showoutput)
}

//Run accepts argurments for docker run, runs the command and returns the first line of stdout, or an error
func Run(runargs string, showoutput bool) error {
	return run("run", runargs, showoutput)
}

//GetHostIP pass in the docker-machine name, get back the ip address or an error
//This would be used to help attach to running docker containers.
func GetHostIP(name string) (string, error) {
	var out []byte
	out, err := exec.Command("docker-machine", "ip", name).Output()
	return string(out), err
}
