package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
输入类似
88 8  8A i i OOI              IIIaa
A
*/
func getCharCount() {
	args := [2]string{}
	input := bufio.NewScanner(os.Stdin)
	for i := range args {
		input.Scan() //读取一行内容
		args[i] = input.Text()
	}

	strs := strings.ToLower(args[0])
	target := strings.ToLower(args[1])

	cnt := 0
	for i := range strs {
		ch := string(strs[i])
		if ch == target {
			cnt++
		}
	}
	fmt.Printf("%d\n", cnt)
}

//
