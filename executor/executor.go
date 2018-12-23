package executor

import (
	"strings"
	"fmt"
	"os"
)

func Executor(userText string){
	if  userText == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	args := strings.Split(userText," ")
	if len(args) > 2 {
		switch args[0] {
		case "get":
			switch args[1] {
			case "smartContractAddress", "sca":
				if len(args) == 3 {
					fmt.Println("Smart contract address: " + args[2])
				}
			case "plasmaBalance", "pb":
				if len(args) == 3 {
					fmt.Println("Plasma balance:" + args[2])
				}
			case "smartContractBalance", "scb":
				if len(args) == 3 {
					fmt.Println("Smart contract balance:" + args[2])
				}
			case "events", "e":
				if len(args) == 3 {
					fmt.Println("Events ...")
				}
			}
		}
	}else{
		fmt.Println("Bad args!")
	}
}
