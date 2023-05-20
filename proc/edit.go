package proc

import (
	"fmt"
	"strings"

	"github.com/kevinburke/ssh_config"
)

type EditResult struct {
	HostPatterns  []string
	PreviousValue string
	CurrentValue  string
	LineNo        int
}

func TestEdit() ([]string, bool, error) {
	fmt.Println("testedit called")

	returnList := []string{"a", "b", "c"}

	return returnList, true, nil
}

func RewriteConfigValue(cfg *ssh_config.Config, targetPtn string, targetKeyName string, inputValue string) ([]EditResult, error) {

	completedHostList := []EditResult{}

	// HostName Matching
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

			if !isContainedWildCard {
				for _, node := range host.Nodes {

					kv, ok := node.(*ssh_config.KV)
					if ok && kv.Key == targetKeyName {

						// rewriting
						previousHostName := kv.Value
						kv.Value = inputValue
						kv.Comment = "This value was rewritten by ssh-conf-cli"

						// add result-info to list
						editResult := EditResult{[]string{}, previousHostName, kv.Value, kv.Pos().Line}
						for _, ptn := range host.Patterns {
							editResult.HostPatterns = append(editResult.HostPatterns, ptn.String())
						}
						completedHostList = append(completedHostList, editResult)

						break
					}
				}
			}
		}
	}
	// flag, rewrited-string, line-number, error
	return completedHostList, nil
}
