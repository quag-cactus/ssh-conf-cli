package proc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

func GetHostPtnList(confPath string, recursive bool) ([][]string, error) {

	hostPtnList := [][]string{}

	// Read config file
	inputFs, err := os.Open(confPath)
	if err != nil {
		fmt.Printf("Failed to open ssh config file: %s", confPath)
		return hostPtnList, err
	}
	inputFs.Seek(0, 0)
	cfg, err := ssh_config.Decode(inputFs)
	if err != nil {
		fmt.Printf("Failed to decode ssh config file: %s", confPath)
		return hostPtnList, err
	}
	defer inputFs.Close()

	for _, host := range cfg.Hosts {

		// get host patterns RECURSIVELY (optional)
		if recursive {
			for _, node := range host.Nodes {
				_, ok := node.(*ssh_config.KV)
				if !ok {
					if strings.HasPrefix(node.String(), "Include") {

						// If includePath is relative path, convert to absolute path
						includePath := strings.Split(node.String(), " ")[1]
						if !filepath.IsAbs(includePath) {
							includePath = filepath.Join(filepath.Dir(confPath), includePath)
						}

						// get matched files
						matchedPathList, err := filepath.Glob(includePath)
						if err != nil {
							return hostPtnList, err
						}

						// recursively get host patterns
						for _, path := range matchedPathList {
							ptnList, err := GetHostPtnList(path, recursive)
							if err != nil {
								return hostPtnList, err
							}
							hostPtnList = append(hostPtnList, ptnList...)
						}
					}
				}
			}
		}

		if host.Patterns[0].String() == "*" && len(host.Patterns) == 1 {
			continue
		}

		// get host patterns in config file
		ptnList := []string{}
		for _, pattern := range host.Patterns {
			ptnList = append(ptnList, pattern.String())
		}
		hostPtnList = append(hostPtnList, []string{strings.Join(ptnList, ", "), confPath})

	}

	return hostPtnList, nil
}
