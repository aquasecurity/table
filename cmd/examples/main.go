package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(readme), "<!--eg-->")
	before := parts[0]
	parts = strings.Split(string(readme), "<!--/eg-->")
	after := parts[len(parts)-1]

	examples, err := os.ReadDir("_examples")
	if err != nil {
		panic(err)
	}

	var output string
	for _, example := range examples {
		if !example.IsDir() {
			continue
		}
		source, err := ioutil.ReadFile(filepath.Join("_examples", example.Name(), "main.go"))
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(source), "noreadme") {
			continue
		}

		name := strings.Title(strings.ReplaceAll(example.Name(), "-", " ")[3:])

		buffer := bytes.NewBuffer([]byte{})
		cmd := exec.Command("go", "run", "."+string(filepath.Separator)+filepath.Join("_examples", example.Name()))
		cmd.Stdout = buffer
		cmd.Stderr = os.Stderr
		cwd, _ := os.Getwd()
		cmd.Dir = cwd
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		output += fmt.Sprintf("\n### Example: %s\n```go\n%s\n```\n\n#### Output\n```\n%s\n```\n", name, string(source), buffer.String())
	}

	output = before + "<!--eg-->" + output + "<!--/eg-->" + after
	if err := ioutil.WriteFile("README.md", []byte(output), 0600); err != nil {
		panic(err)
	}
}
