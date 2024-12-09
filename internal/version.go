package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/mod/semver"
)

const Version = "v0.7.0"

func ValidateLatestVersion() {
	cmd := exec.Command(goCLI, "list", "-m", "github.com/renanbastos93/boneless@latest")
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	if cmd.Run() != nil {
		return
	}

	if result := stdOut.String(); result != "" {
		_, v, _ := strings.Cut(result, " ")
		v = v[:len(v)-1]

		// Check if the latest version is greater than the current version
		if semver.Compare(Version, v) > 0 || v == Version {
			return
		}

		msg := "\033[31mNew version available! Check out our release and update the Boneless!\n"
		msg += "\033[0m\033[1mMore details: \033[0mhttps://github.com/renanbastos93/boneless/releases/tag/%s\n---\n"
		fmt.Printf(msg, v)
	}
}
