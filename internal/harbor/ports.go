package harbor

import (
	"log"
	"os/exec"
	"strings"
)

type Port struct {
	Process  string
	PID      string
	User     string
	Protocol string
	Address  string
}

func FindPorts() []*Port {
	out, err := exec.Command("lsof", "-i", "-P", "-n").Output()
	if err != nil {
		log.Fatal("portman: ", err)
	}
	var ports []*Port
	lines := strings.Split(string(out), "\n")

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		ports = append(ports, &Port{
			Process: fields[0],
			PID: fields[1],
			User: fields[2],
			Protocol: fields[7],
			Address: fields[8],
		})
	}
	return ports
}

func (p *Port) Kill() error {
	return exec.Command("kill", p.PID).Err
}
