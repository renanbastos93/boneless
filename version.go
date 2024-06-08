package boneless

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "embed"

	"golang.org/x/mod/semver"
)

type version struct {
	Name  string `json:"name"`
	Body  string `json:"body"`
	Draft bool   `json:"draft"`
}

const latestTagURL = "https://api.github.com/repos/renanbastos93/boneless/releases/latest"

// curl https://api.github.com/repos/renanbastos93/weaver/contents/runtime/version/version.go -H "Accept: application/vnd.github+json"

//go:embed VERSION
var Version string

func getLastestVersion() (v version) {
	var (
		client   *http.Client
		response *http.Response
		body     []byte
		err      error
	)

	defer func() {
		debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
		if debug && err != nil {
			println("\033[31mFailed to fetch the latest package version: \033[0m\033[1m" + err.Error() + "\033")
		}
	}()

	client = &http.Client{
		Timeout: time.Second,
	}

	response, err = client.Get(latestTagURL)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("because status code is %s", response.Status)
		return
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &v)
	return v
}

func ValidateLatestVersion() {
	v := getLastestVersion()
	if semver.Max(v.Name, Version) != Version {
		msg := "\033[31mNew version available! Check out our release and update the Boneless!\n"
		value := v.Name
		if withDetails, _ := strconv.ParseBool(os.Getenv("VERSION_DETAILS")); withDetails {
			msg += "\033[0m\033[1mMore details: \033\n---\n"
			msg += "===\n%s\n===\n"
			value = v.Body
		} else {
			msg += "\033[0m\033[1mMore details: \033[0mhttps://github.com/renanbastos93/boneless/releases/tag/%s\n---\n"
		}
		fmt.Printf(msg, value)
	}
}
