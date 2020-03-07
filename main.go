package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("Please enter user ID: ")
		if id, err := getUserIdFromConsole(); err != nil {
			panic(err)
		} else {
			Sign(id)
			printUserCountInConsole()
		}
	}
}

func getUserIdFromConsole() (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	id, err := strconv.ParseInt(strings.Replace(text, "\n", "", -1), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func printUserCountInConsole() {
	count, err := GetActiveUserCount()
	if err != nil {
		fmt.Println("Cannot get active user count")
	}
	fmt.Println(fmt.Sprintf("The active user is: %d\n", count))
}
