package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"golang.org/x/mod/modfile"
)

var opts struct {
	path *string
}

func setupFlags() {
	opts.path = pflag.String("path", "", "Path to the go.mod file")
}

func errorExit(e error) {
	fmt.Fprintln(os.Stderr, e.Error())
	os.Exit(1)
}

func main() {

	setupFlags()
	pflag.Parse()

	expanded, err := homedir.Expand(*opts.path)
	if err != nil {
		errorExit(err)
	}

	abs, err := filepath.Abs(expanded)
	if err != nil {
		errorExit(err)
	}

	stat, statErr := os.Stat(abs)
	if statErr != nil {
		errorExit(errors.Wrap(statErr, fmt.Sprintf("could not stat file at: %s", abs)))
	}
	if stat.IsDir() {
		errorExit(fmt.Errorf("expected file at path: %s", abs))
	}

	file_bytes, err := ioutil.ReadFile(abs)
	if err != nil {
		errorExit(errors.Wrap(err, fmt.Sprintf("failed reading file at path: %s", abs)))
	}

	parsedFile, err := modfile.Parse(abs, file_bytes, nil)
	if err != nil {
		errorExit(errors.Wrap(err, fmt.Sprintf("failed reading file at module file: %s", abs)))
	}

	rawBytes, jsonErr := json.Marshal(parsedFile)
	if err != nil {
		errorExit(errors.Wrap(jsonErr, fmt.Sprintf("failed converting module file to JSON: %s", abs)))
	}

	fmt.Println(string(rawBytes))

}
