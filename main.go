package main

import "plugin"

type pluginFunc func(string)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Print")
	if err != nil {
		panic(err)
	}

	f.(pluginFunc)("Hello, World")
}
