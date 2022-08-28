package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
)

/*
	Взаимодействие с ОС
	Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

	- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
	- pwd - показать путь до текущего каталога
	- echo <args> - вывод аргумента в STDOUT
	- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
	- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


	Так же требуется поддерживать функционал fork/exec-команд
	Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

	*Шелл — это обычная консольная программа, которая будучи запущенной,
	в интерактивном сеансе выводит некое приглашение
	в STDOUT и ожидает ввода пользователя через STDIN.
	Дождавшись ввода, обрабатывает команду согласно своей логике
	и при необходимости выводит результат на экран.
	Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	// запускаем считывание с stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		curDir, _ := os.Getwd()
		fmt.Printf("%s$ ", curDir)
		if scanner.Scan() {
			if cmd := scanner.Text(); cmd != "" {
				err := parseCmds(cmd)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// parseCmds проверяем что за команда нам пришла и выполняем ее
func parseCmds(cmds string) error {
	cmd := strings.Split(cmds, " ")
	switch cmd[0] {
	case "cd":
		err := cd(cmd)
		if err != nil {
			return err
		}
	case "pwd":
		curDir, _ := os.Getwd()
		fmt.Println(curDir)
	case "echo":
		fmt.Println(strings.Join(cmd[1:], " "))
	case "kill":
		if len(cmd) > 1 {
			err := kill(cmd[1])
			if err != nil {
				return err
			}
		} else {
			return errors.New("Enter pid process to kill")
		}
	case "\\quit":
		os.Exit(0)
	default:
		err := execCmd(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// cd команда смены директории
func cd(arg []string) error {
	if len(arg) == 1 {
		os.Chdir(os.Getenv("HOME"))
		return nil
	}
	if len(arg) == 2 {
		err := os.Chdir(arg[1])
		if err != nil {
			return errors.New("cd: No such file or directory")
		}
	} else {
		return errors.New("cd: too many arguments")
	}
	return nil
}

// kill команда "убить процесс"
func kill(arg string) error {
	pid, err := strconv.Atoi(arg)
	if err != nil {
		return errors.New("Wrong PID")
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = process.Kill()
	if err != nil {
		return err
	}
	return nil
}

// execCmd выполняет exec команду и ждет ее завершения
// если не находит такую команду возвращает ошибку
func execCmd(cmd []string) error {

	var execCmd *exec.Cmd
	if len(cmd) > 1 {
		execCmd = exec.Command(cmd[0], cmd[1:]...)
	} else {
		execCmd = exec.Command(cmd[0])
	}
	execCmd.Stdout = os.Stdout
	execCmd.Stdin = os.Stdin
	execCmd.Stderr = os.Stderr
	err := execCmd.Start()
	if err != nil {
		return err
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	select {
	case <-sigCh:
		return nil
	default:
		execCmd.Wait()
	}
	return nil
}
