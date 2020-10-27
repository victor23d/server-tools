package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var processedEnv string = process(".env-dev")
	fmt.Printf(processedEnv)
	// trimDuplicate(envs)

}

func process(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)

	// Comments at top, then alphabetical order
	var envs []string
	envfile := string(data)
	envs = strings.Split(envfile, "\n")

	// trimEmpty(envs)

	sort.Strings(envs)

	var envSorted strings.Builder
	for i, _ := range envs {
		envSorted.WriteString(envs[i])
		envSorted.WriteString("\n")
	}

	return (envSorted.String())
}

func trimEmpty(envs []string) []string {
	for i := 0; i < len(envs); i++ {
		if envs[i] == "" {
			// Delete without preserving order
			envs[i] = envs[len(envs)-1]
			envs = envs[:len(envs)-1]
		}
	}
	return envs
}

func trimDuplicate() {

}

func trimK8S() {

}
