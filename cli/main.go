package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/prazd/go-cli/completer"
	"github.com/prazd/go-cli/executor"
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
