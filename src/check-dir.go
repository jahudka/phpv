package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"

	"github.com/Masterminds/semver/v3"
)

type Composer struct {
	Require map[string]string `json:"require"`
}

func main() {
	if os.Getenv("PHPV_BREW_PREFIX") == "" {
        return
    }

	versionStrings := strings.Split(strings.TrimSpace(os.Getenv("PHPV_AVAILABLE_VERSIONS")), "\n")

	rangeStr := readPHPConstraint("composer.json")
	if rangeStr == "" {
		return
	}

	constraint, err := semver.NewConstraint(rangeStr)
	if err != nil {
		return
	}

	for _, vs := range versionStrings {
		v, err := semver.NewVersion(vs)
		if err != nil {
			continue
		}

		if !constraint.Check(v) {
		    continue
		}

        origPaths := strings.Split(os.Getenv("PATH"), ":")
        paths := origPaths[:0]
        prefix := fmt.Sprintf("%s/php@", os.Getenv("PHPV_BREW_PREFIX"))

        for _, path := range origPaths {
            if !strings.HasPrefix(path, prefix) {
                paths = append(paths, path)
            }
        }

        fmt.Printf("export PHPV_CURRENT_VERSION=\"%s\"\n", v.Original())
        fmt.Printf("export PATH=\"%s%s/bin:%s\"\n", prefix, v.Original(), strings.Join(paths, ":"))
        fmt.Printf("echo 'Using PHP %s'\n", v.Original())
        return
	}
}

func readPHPConstraint(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	var composer Composer
	if err := json.Unmarshal(data, &composer); err != nil {
		return ""
	}

	if composer.Require == nil {
		return ""
	}

	rangeStr, ok := composer.Require["php"]
	if !ok {
		return ""
	}

	if rangeStr == "" {
		return ""
	}

	return rangeStr
}
