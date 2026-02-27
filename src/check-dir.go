package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
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

	cwd, err := os.Getwd()
	if err != nil {
	    return
	}

	rangeStr := resolveRequiredVersion(cwd)
	if rangeStr == "" {
		return
	}

	constraint, err := semver.NewConstraint(rangeStr)
	if err != nil {
		return
	}

    versionStrings := strings.Split(strings.TrimSpace(os.Getenv("PHPV_AVAILABLE_VERSIONS")), "\n")
    currentVersionStr := os.Getenv("PHPV_CURRENT_VERSION")

	for _, vs := range versionStrings {
		v, err := semver.NewVersion(vs)
		if err != nil {
			continue
		}

		if !constraint.Check(v) {
		    continue
		}

        if vs == currentVersionStr {
            return
        }

        setActiveVersion(vs)
        return
	}
}

func resolveRequiredVersion(dir string) string {
    for {
        data, err := os.ReadFile(fmt.Sprintf("%s/composer.json", dir))

        if err == nil {
            return readComposerJson(data)
        }

        parent := filepath.Dir(dir)

        if parent == dir {
            return ""
        }

        dir = parent
    }
}

func readComposerJson(data []byte) string {
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

func setActiveVersion(version string) {
    prefix := fmt.Sprintf("%s/php@", os.Getenv("PHPV_BREW_PREFIX"))
    origPaths := strings.Split(os.Getenv("PATH"), ":")
    paths := origPaths[:0]

    for _, path := range origPaths {
        if !strings.HasPrefix(path, prefix) {
            paths = append(paths, path)
        }
    }

    fmt.Printf("export PHPV_CURRENT_VERSION=\"%s\"\n", version)
    fmt.Printf("export PATH=\"%s%s/bin:%s%s/sbin:%s\"\n", prefix, version, prefix, version, strings.Join(paths, ":"))
    fmt.Printf("echo 'Using PHP %s'\n", version)
}
