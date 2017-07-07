/**
实现如下功能：
	1：音乐库功能，用户可以查看、增加和删除曲目
	2：播放音乐
	3：支持mp3和wav，支持其他类型的扩展
	4：命令行退出(Q)
**/

package main

import (
	"av"
	"bufio"
	"fmt"
	"mlib"
	"os"
	"strconv"
	"strings"
)

var id int = 1
var lib *library.MusicManager

var ctrl, signal chan int

func handleCommandLibs(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(e)
			fmt.Println(i+1, ": ", e.Name, e.Artist, e.Source, e.Type)
		}

	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}

	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}

	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handleCommandPlay(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)

	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter Command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "Q" || line == "q" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleCommandLibs(tokens)
		} else if tokens[0] == "play" {
			handleCommandPlay(tokens)
		} else {
			fmt.Println("Unrecoginzed command ", tokens[0])
		}
	}

}
