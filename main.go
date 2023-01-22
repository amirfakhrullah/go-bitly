package main

import (
	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/server"
)

func main() {
	model.SetupDB()
	server.SetupAndListen()
}
