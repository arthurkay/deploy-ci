package ansible

import (
	ansibler "github.com/febrianrendak/go-ansible"
)

func Run(play string, user string) error {
	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
		User:       user,
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          play,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "Deploy CI",
	}

	err := playbook.Run()
	if err != nil {
		return err
	}
	return nil
}
