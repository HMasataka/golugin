package main

import "plugin"

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Print")
	if err != nil {
		panic(err)
	}

	f.(func(string))("Hello, World")
}
