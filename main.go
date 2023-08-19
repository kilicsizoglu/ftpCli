package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome Ftp Client")

	reader := bufio.NewReader(os.Stdin)
	var client ftpConnect

	for {

		fmt.Println("Login Ftp")

		for {
			fmt.Println("url :")
			urlbytes, _, err2 := reader.ReadLine()
			if err2 != nil {
				continue
			}
			url := string(urlbytes)

			fmt.Println("username :")
			usernamebytes, _, err3 := reader.ReadLine()
			if err3 != nil {
				continue
			}
			username := string(usernamebytes)

			fmt.Println("password :")
			passwordbytes, _, err4 := reader.ReadLine()
			if err4 != nil {
				continue
			}
			password := string(passwordbytes)

			err1 := client.login(url, username, password)
			if err1 != nil {
				continue
			} else {
				break
			}
		}

		fmt.Println("Command List")

		fmt.Println("ls => List File and Folder")

		fmt.Println("mkdir => Create Folder")
		fmt.Println("rmdir => Delete Folder")

		fmt.Println("mk => Create File")
		fmt.Println("read => Read File")
		fmt.Println("edit => Edit File")
		fmt.Println("rm => Delete File")

		fmt.Println("Enter Command")

		line, _, err := reader.ReadLine()
		if err != nil {
			continue
		}

		command := string(line)

		commands := strings.Split(command, " ")

		if commands[0] == "exit" {
			break
		}

		switch commands[0] {

		case "ls":
			dir, err := client.ReadDir(commands[1])
			if err != nil {

				continue

			}
			for i := 0; i == len(dir); i++ {

				fmt.Println(dir[i].Name())

			}
			break
		case "mkdir":
			err := client.CreateDir(commands[1])
			if err != nil {

				continue

			}
			fmt.Println("ok...")
			break
		case "rmdir":
			err := client.DeleteDir(commands[1])
			if err != nil {

				continue

			}
			fmt.Println("ok...")
			break
		case "mk":
			err := client.CreateFile(commands[1])
			if err != nil {

				continue

			}
			fmt.Println("ok...")
			break
		case "read":
			res, err := client.ReadFile(commands[1])
			if err != nil {

				continue

			}
			var sb strings.Builder
			for {
				var b []byte
				number, err6 := res.Read(b)
				if err6 != nil {
					break
				}
				sb.Write([]byte(strconv.Itoa(number)))
			}
			fmt.Println(sb.String())
			break
		case "edit":
			var writer io.Writer
			for {
				readLine, _, err := reader.ReadLine()
				if err != nil {
					return
				}
				if string(readLine) != "exit" {
					_, err := writer.Write(readLine)
					if err != nil {
						break
					}
				} else {
					break
				}
			}
			err := client.WriteFile(commands[1], writer)
			if err != nil {

				continue

			}
			fmt.Println("")
			break
		case "rm":
			err := client.DeleteFile(commands[1])
			if err != nil {

				continue

			}
			fmt.Println("ok ...")
			break

		}
	}

}
