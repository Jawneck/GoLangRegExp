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
func poregtonfa(pofix string){

}

func main(){
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}