package main

/*
 * Driver implementation (Terraform executor integration)
 * ------------------------------------------------------
 */

import (
	"github.com/krzysztof-miemiec/docker-machine-driver-terraform/terraform"
)

// Ensure that the Terraform executor has been successfully configured.
func (driver *Driver) validateTerraformer() error {
	_, err := driver.getTerraformer()

	return err
}

// Get the Terraform executor for the driver's current configuration.
func (driver *Driver) getTerraformer() (*terraform.Terraformer, error) {
	if driver.terraformer == nil {
		err := driver.ensureConfigDirIsResolved()
		if err != nil {
			return nil, err
		}

		if driver.TerraformExecutablePath != "" {
			driver.terraformer, err = terraform.NewWithExecutable(driver.TerraformExecutablePath, driver.ConfigDir)
		} else {
			driver.terraformer, err = terraform.New(driver.ConfigDir)
		}
		if err != nil {
			return nil, err
		}
	}
	if driver.LockTimeout != "0" {
		driver.terraformer.LockFlag = "-lock-timeout=" + driver.LockTimeout
	} else {
		driver.terraformer.LockFlag = "-lock=false"
	}
	return driver.terraformer, nil
}
