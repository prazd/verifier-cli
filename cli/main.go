package main

import (
	"github.com/c-bata/go-prompt"
	"../completer"
	"../executor"
	"fmt"
)


func main() {
	fmt.Println("------------Plasma Verifier----------")
	p := prompt.New(
		executor.Executor,
		completer.Completer,
		prompt.OptionPrefix("--> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
