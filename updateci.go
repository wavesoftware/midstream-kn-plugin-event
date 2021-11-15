//go:build mage
// +build mage

package main

import (
	"log"
	"os"
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/pkg/files"
)

// UpdateCI configuration in the $OPENSHIFT directory.
// NOTE: override $OPENSHIFT env variable if guessed value is not correct.
// NOTE: Makes changes outside this repository.
// nolint:deadcode
func UpdateCI() error {
	guess := path.Join(files.ProjectDir(), "../../github.com/openshift/release")
	openshift, ok := os.LookupEnv("OPENSHIFT")
	if !ok {
		openshift = guess
		log.Println("WARN: Using guessed dir for openshift/release: ", guess)
	}
	return sh.RunV("sh", "./openshift/ci-operator/update-ci.sh", openshift) // nolint:wrapcheck
}
