package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	//Abrindo o buffer para ler o input
	reader := bufio.NewReader(os.Stdin)
	//Pegando o hostanme
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	//Pegando o username
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	for {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		//Printa o usuario
		fmt.Print(user.Username + "@" + hostname + " " + path + " $ ")
		input, err := reader.ReadString('\n')

		//Checando por erro
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		//Usando a minha função para executar o comando
		if err = execCmd(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func execCmd(cmd string) error {
	//Removendo o \n da string
	cmd = strings.TrimSuffix(cmd, "\n")
	args := strings.Split(cmd, " ")
	command := exec.Command(args[0], args[1:]...)
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	//Syncando com o Standart Output e Standart error
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	//Retornando um possível erro
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir("/home/" + user.Username)
		} else {
			return os.Chdir(args[1])
		}
	case "exit":
		os.Exit(0)
	}

	return command.Run()
}
