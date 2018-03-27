package main

import (
	"fmt"
)

//struct called state with two possible edges that point to other states.
type state struct{
	symbol rune//Represents a state with one or two arrows labelled by Epsilon.
	edge1 *state
	edge2 *state
}

//struct which keeps track of the initial state and accept state of a fragment of the NFA.
type nfa struct{
	initial *state
	accept *state
}

//Postfix RegExp to NFA function.
func poregtonfa(pofix string) *nfa{
	nfaStack := []*nfa{}//This stack is an array of pointers to NFA's that is empty. A stack is LIFO.

	//Loops through postfix regular expression a character/rune at a time.
	for _, r := pxfix{
		switch r{
		case '.'://Pops two elements off the NFA stack.
			frag2 := nfaStack[len(nfaStack)-1]//The last element on the stack.
			nfaStack := nfaStack[:len(nfaStack)-1]//Every element in stack up to but not including the last element.
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack := nfaStack[:len(nfaStack)-1]

			//Join the two fragments together
			//Fragment 1s accept states edge1 points at fragment 2s initial state.
			frag1.accept.edge1 = frag2.initial
			
			//Appends a new pointer to an nfa struct that represents new fragment(frag1 + frag2)
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|'://Pops two elements off the NFA stack.
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack := nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack := nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{}
			
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '*':

		//Any character that is not one of the three special characters.
		default:
		}
	}

	return nfaStack[0]
}

func main(){
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}