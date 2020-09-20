package terraform

import (
	"errors"
)

// Init invokes Terraform's "init" command.
func (terraformer *Terraformer) Init() error {
	success, err := terraformer.RunStreamed("init",
		"-input=false", // non-interactive
		"-no-color",
		terraformer.LockFlag,
	)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("failed to execute 'terraform init'")
	}
	return nil
}
