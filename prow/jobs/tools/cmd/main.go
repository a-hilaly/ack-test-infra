package main

import (
	"fmt"
)

// first step get the Go version from ecrpublic
// second step read the build-config.yaml
// third step compare the latest Go version with the one in build-config
// if gt => icr patch in image_config.yaml,
//       => update build-config.yaml (go verison)
// open PR

func listGoVersions(repositoryName string) ([]string, error) {
	return nil, nil
}

func findHighestGoVersion(versions []string) (string, error) {
	return "", nil
}

func readCurrentBuildGoVersion(filepath string) (string, error) {
	return "", nil
}

func patchImageConfig(filepath string, newMap map[string]string) error {
	return nil
}

func patchBuildGoVersion(filepath string, goVersion string) error {
	return nil
}

func commitAndSendPR(prRepository string, issueRepository string) error {
	return nil
}

func isGreaterVersion(v1, v2 string) bool {
	return true
}

func runCobraJob() error {
	var (
		ecrRepositoryName string
		buildConfigFile   string
	)

	goVersions, err := listGoVersions(ecrRepositoryName)
	if err != nil {
		return fmt.Errorf("couldn't list the go verison from ECR: %v", err)
	}

	latestGoVersion, err := findHighestGoVersion(goVersions)
	if err != nil {
		return fmt.Errorf("couldn't find the hightest go version: %v", err)
	}

	currentGoVersion, err := readCurrentBuildGoVersion(buildConfigFile)
	if err != nil {
		return fmt.Errorf("couldn't read the ACK build config: %v", err)
	}

	if latestGoVersion == currentGoVersion {
		return nil
	}

	if !isGreaterVersion(latestGoVersion, currentGoVersion) {
		return fmt.Errorf("big problem")
	}

	err = patchBuildGoVersion(buildConfigFile, latestGoVersion)
	if err != nil {
		return fmt.Errorf("couldn't : %v", err)
	}

	// build the new image config map

	err = patchImageConfig(buildConfigFile, struct{}{})
	if err != nil {
		return fmt.Errorf("couldn't read the ACK build config: %v", err)
	}

	err = commitAndSendPR("test-infra", "community")
	if err != nil {
		return fmt.Errorf("couldn't read the ACK build config: %v", err)
	}

	return nil
}

func main() {}
