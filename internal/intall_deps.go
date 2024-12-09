package internal

import (
	"os/exec"
)

type packages struct {
	name, pkg string
}

var PackagesForInstall = []packages{
	{"golang-migrate", "github.com/golang-migrate/migrate/v4/cmd/migrate@latest"},
	{"sqlc", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest"},
	{"weaver", "github.com/ServiceWeaver/weaver/cmd/weaver@latest"},
}

func InstallDeps(name string) {
	if name != "" {
		for _, p := range PackagesForInstall {
			if p.name == name {
				installDeps(p)
			}
		}
	} else {
		installDeps(PackagesForInstall...)
	}
}

func installDeps(packages ...packages) {
	for _, p := range packages {
		if IsInstalled(p.pkg) {
			println(p.name, "already installed!")
			continue
		}
		GoInstall(p.name, p.pkg)
	}
}

func IsInstalled(packageName string) bool {
	pathBin, err := exec.LookPath(packageName)
	return err != nil || pathBin == ""
}

func GoInstall(packageName string, args ...string) {
	cmdArgs := append([]string{installCmd}, args...)
	err := runCmd(goCLI, cmdArgs...)
	if err != nil {
		panic(err)
	}
	println(packageName, "installed!")
}

func UpdateDeps(name string) {
	if name != "" {
		for _, p := range PackagesForInstall {
			if p.name == name {
				updateDeps(p)
			}
		}
	} else {
		updateDeps(PackagesForInstall...)
	}
}

func updateDeps(packages ...packages) {
	for _, p := range packages {
		GoInstall(p.name, p.pkg)
	}
}
