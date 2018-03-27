package main

import (
	"fmt"
)

//A function used to convert infix regular expressions to postfix regular expressions.
func intopost(infix string) string{
	postfix := ""
	return postfix
}

func main(){

	//Four regular expressions written in infix notation.

	//Answer: ab.c*.
	fmt.Println("Infix    ",  "a.b.c")
	fmt.Println("Postfix  ",  intopost("a.b.c"))

	//Answer: abd|.*
	fmt.Println("Infix    ",  "(a.(b|d))*")
	fmt.Println("Postfix  ",  intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix    ",  "a.(b|d).c*")
	fmt.Println("Postfix  ",  intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix    ",  "a.(b.b)+.c")
	fmt.Println("Postfix  ",  intopost("a.(b.b)+.c"))


}