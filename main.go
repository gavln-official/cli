package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"

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
		fmt.Println(line)
	}
}

func getCommandAndArgs(line string) (command string, args []string, err error) {
	return "", nil, nil
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
	account := NewAccount(name, password)
	_, err = account.Login()
	if err != nil {
		panic(err)
	}
}
