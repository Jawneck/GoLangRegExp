package main

import (
	"fmt"
)

//A function used to convert infix regular expressions to postfix regular expressions.
func intopost(infix string) string{
	//A map which keeps track of the precedence of the special characters.
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	pofix := []rune{} //An empty array of runes.
	s := []rune{}

	//A loop over the infix string a character at a time.
	for _, r := range infix{
		switch{
		case r == '(':
			s = append(s, r)//Appends '(' onto the end of the stack	.
		case r == ')':
			//While the last character on the stack doesn't equal an open bracket.
			for s[len(s)-1] != '('{
				pofix = append(pofix, s[len(s)-1])//Appends the top element of s onto pofix.
				s = s[:len(s)-1]//s = everything in s except for the last character.
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			//While the stack still has elements on it AND the precedence of the current character that you're reading is
			//less than the precendence of the top element of the stack.
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]]{
				pofix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)//Appends character r onto the end of the pofix array.
		}
	}

	//If there is any element left on top of the stack , append it to the output.
	for len(s) > 0{
		pofix = append(pofix, s[len(s)-1])
		s = s[:len(s)-1]
	}

	return string(pofix)
}

func main(){

	//Four regular expressions written in infix notation.

	//Answer: ab.c*.
	fmt.Println("Infix    ",  "a.b.c*")//An A followed by a B, followed by 0 or more C's.
	fmt.Println("Postfix  ",  intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix    ",  "(a.(b|d))*")//Zero or more of the following and A followed by a B or a D.
	fmt.Println("Postfix  ",  intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix    ",  "a.(b|d).c*")
	fmt.Println("Postfix  ",  intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix    ",  "a.(b.b)+.c")
	fmt.Println("Postfix  ",  intopost("a.(b.b)+.c"))
}