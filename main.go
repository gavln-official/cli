package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"strings"

	"github.com/Gvinaxu/cli/handler"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	Version = "v1.0.0"
	Website = "https://gavln.com"
	Banner  = `
     ___          ___                      ___   ___     
    /  /\        /  /\        ___         /  /\ /  /\    
   /  /::\      /  /::\      /  /\       /  /://  /::|   
  /  /:/\:\    /  /:/\:\    /  /:/      /  /://  /:|:|   
 /  /:/  \:\  /  /::\ \:\  /  /:/      /  /://  /:/|:|__ 
/__/:/_\_ \:\/__/:/\:\_\:\/__/:/  ___ /__/://__/:/ |:| /\
\  \:\__/\_\/\__\/  \:\/:/|  |:| /  /\\  \:\\__\/  |:|/:/
 \  \:\ \:\       \__\::/ |  |:|/  /:/ \  \:\   |  |:/:/ 
  \  \:\/:/       /  /:/  |__|:|__/:/   \  \:\  |__|::/  
   \  \::/       /__/:/    \__\::::/     \  \:\ /__/:/   
    \__\/        \__\/         ~~~~       \__\/ \__\/        	%s
	
Know its white, keep its black
%s
______________________________
					
`
)

var (
	h *handler.Handler
)

func main() {
	checkUser()
	fmt.Printf(Banner, Version, Website)

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		line, err := f.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command, args, err := getCommandAndArgs(line)
		if err != nil {
			continue
		}
		if command == "quit" {
			fmt.Println("bye!")
			break
		}
		h.InvokeCmd(command, args)
	}
}

func init() {
	f := &handler.FileReq{}
	h = handler.NewHandler(f)
}

func getCommandAndArgs(line string) (command string, args []string, err error) {
	line = strings.TrimSpace(line)
	all := strings.Split(line, " ")
	if len(all) == 0 {
		return "", nil, errors.New("input is nil")
	}
	args = make([]string, 0)
	for i, v := range all {
		if i == 0 {
			continue
		}
		if v == "" {
			continue
		}
		args = append(args, v)
	}
	return all[0], args, nil
}

func checkUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Gavln User Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter Password: ")
	bytes, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}
	password := string(bytes)
	name = strings.ReplaceAll(name, "\n", "")
	account := handler.NewAccount(name, password)
	_, err = account.Login()
	if err != nil {
		panic(err)
	}
}
