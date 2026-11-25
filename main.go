// Command gobasic is a tiny Go fixture: one dependency, one call into it.
package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	var v map[string]string
	if err := yaml.Unmarshal([]byte("greeting: hello"), &v); err != nil {
		panic(err)
	}
	fmt.Println(v["greeting"])
}
