package process

import "github.com/shirou/gopsutil/process"

func ProcessExists(processName string) (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}

	for _, process := range processes {
		name, _ := process.Name()
		if name == processName {
			return true, nil
		}
	}

	return false, nil
}
