package ansible

import (
	ansibler "github.com/febrianrendak/go-ansible"
)

// Run takes two parameters, a playbook file location and the user in whos
// space the play runs. The playbook is only executed on the local machine.
// And as such there is no need to setup ssh keys for remote nodes.
// btw this function call returns an error.
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
