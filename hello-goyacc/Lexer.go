package main

import (
	"fmt"
	"time"
)

type myLexer struct {
	// 待解析的指令先暂存在input里面 用byte数组方便逐字符操作
	input []byte
	// 当前字符索引
	index int
	// 存储以及扫描到的内容 内容为hello时 返回hello
	buffer string
}

func (m *myLexer) Lex(lval *helloSymType) int {
	if len(m.input) == 0 {
		return 0
	}
	for i := 0; i < len(m.input); i++ {
		char := m.input[i]
		switch char {
		case 'h', 'e', 'l', 'o':
			m.buffer += string(char)
		}
		if m.buffer == "hello" {
			return HELLO
		}
	}
	return 0
}

func (m *myLexer) Error(s string) {
	fmt.Println(s)
}

func main() {
	helloParse(&myLexer{
		input:  []byte("hello"),
		index:  0,
		buffer: "",
	})
	time.Sleep(time.Second)
	helloParse(&myLexer{
		input:  []byte("hhhhhh"),
		index:  0,
		buffer: "",
	})
}
