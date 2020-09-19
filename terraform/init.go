package terraform

import (
	"errors"
)

// Init invokes Terraform's "init" command.
func (terraformer *Terraformer) Init() (success bool, err error) {
	success, err = terraformer.RunStreamed("init",
		"-input=false", // non-interactive
		"-no-color",
	)
	if err != nil {
		return
	}
	if !success {
		err = errors.New("failed to execute 'terraform init'")
	}

	return
}
