package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"studytimer"
)

func main() {
	now := time.Now().Unix()
	fn := fmt.Sprintf("/tmp/st-%05d.txt", os.Getuid())
	// already started
	if studytimer.IsExist(fn) {
		if !prompt("timer file found. finished study? (y/n):") {
			fmt.Println("continue your study")
			return
		}

		start, err := studytimer.ReadTime(fn)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		timeDiff := now - start
		fmt.Printf("you spent %02d:%02d for this study\n", timeDiff/3600, (timeDiff%3600)/60)
		err = os.Remove(fn)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("timer file delete suceeded %s\n", fn)
		return
	}

	if !prompt("timer file not found. start study? (y/n):") {
		fmt.Println("see you")
		return
	}
	err := studytimer.WriteTime(fn, now)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("timer file create suceeded %s\n", fn)
	return
}

func prompt(text string) bool {
	fmt.Printf(text)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input != "y" || scanner.Err() != nil {
		return false
	}
	return true
}
