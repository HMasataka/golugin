package main

import (
	"errors"
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

	f.(func(string))("Hello, World")
}

func getPluginFile() (*plugin.Plugin, error) {
	p, err := plugin.Open(pluginFileName)
	if err == nil {
		return p, nil
	}

	p, err = plugin.Open(filepath.Join(pluginDirName, pluginFileName))
	if err == nil {
		return p, nil
	}

	return nil, errors.New("plugin.Open failed")
}
