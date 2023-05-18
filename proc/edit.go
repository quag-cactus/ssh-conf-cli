package proc

import (
	"fmt"
	"strings"

	"github.com/kevinburke/ssh_config"
)

type EditResult struct {
	previousValue string
	currentValue  string
	lineno        int
}

func TestEdit() ([]string, bool, error) {
	fmt.Println("testedit called")

	returnList := []string{"a", "b", "c"}

	return returnList, true, nil
}

func RewriteConfigValue(cfg *ssh_config.Config, targetPtn string, targetKeyName string, inputValue string) ([]EditResult, error) {

	completedHostList := []EditResult{}

	// Match HostName

	for _, host := range cfg.Hosts {

		isContainedWildCard := false

		if host.Matches(targetPtn) {
			// A wildCard is not supported
			for _, pattern := range host.Patterns {
				if strings.Contains(pattern.String(), "*") {
					isContainedWildCard = true
					break
				}
			}

			// rewriting
			if !isContainedWildCard {
				for _, node := range host.Nodes {

					kv, ok := node.(*ssh_config.KV)
					if ok && kv.Key == targetKeyName {
						previousHostName := kv.Value
						kv.Value = inputValue
						completedHostList = append(completedHostList, EditResult{previousHostName, kv.Value, kv.Pos().Line})
						kv.Comment = "This value was rewritten by ssh-conf-cli"
						// fmt.Printf("Hostname rewrited: %s -> %s (ln: %d)\n", previousHostName, kv.Value, kv.Pos().Line)

						break
					}
				}
			}
		}
	}
	// flag, rewrited-string, line-number, error
	return completedHostList, nil
}
