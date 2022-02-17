package main

import (
	"errors"
	"os"
	"path/filepath"
	"plugin"
)

const (
	pluginFileName = "plugin.so"
	pluginDirName  = "so"
)

func main() {
	pp, err := getPluginFile()
	if err != nil {
		panic(err)
	}

	f, err := pp.Lookup("Print")
	if err != nil {
		panic(err)
	}

	m := map[string]string{
		"key": "value",
	}
	f.(func(map[string]string))(m)
}

func getPluginFile() (*plugin.Plugin, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	p, err := plugin.Open(filepath.Join(pwd, pluginFileName))
	if err == nil {
		return p, nil
	}

	p, err = plugin.Open(filepath.Join(pwd, pluginDirName, pluginFileName))
	if err == nil {
		return p, nil
	}

	return nil, errors.New("plugin.Open failed")
}
