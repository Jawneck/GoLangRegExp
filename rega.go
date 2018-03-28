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
	for _, r := range pofix{
		switch r{
		case '.'://POPS two elements off the NFA stack.
			frag2 := nfaStack[len(nfaStack)-1]//The last element on the stack.
			nfaStack := nfaStack[:len(nfaStack)-1]//Every element in stack up to but not including the last element.
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//Join the two fragments together
			//Fragment 1s accept states edge1 points at fragment 2s initial state.
			frag1.accept.edge1 = frag2.initial
			
			//APPENDS a pointer to an nfa that represents new fragment(frag1 + frag2).
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|'://POPS two elements off the NFA stack.
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack := nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			//New initial state where edge1 points to frag1 initial state, edge2 points to frag2 initial state.
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//Fragment 1 accept states edge points to the new accept state, and likewise for fragment 2.
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//APPENDS a pointer to an nfa but the initial and accept states are new.
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*'://POPS one elements off the NFA stack.
			frag := nfaStack[len(nfaStack)-1]
			nfaStack := nfaStack[:len(nfaStack)-1]

			accept := state{}
			//A new initial state that has its edge1 point to the initial state of the fragment popped,
			//edge2 points to the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//Joins the edge 1 accept state of the fragment to the initial state of that fragment.
			//Joins the edge 2 accept state to the new accept state.
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//Any character that is not one of the three special characters.
		default://Push something to the stack
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	if len(nfaStack) != 1{
		fmt.Println("Uh oh:", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

//Function which takes a list of pointers to states, takes a single state, and the accept state.
//Returns a list of pointers to states.
func addState(l []*state, s *state, a *state) []*state{
	l = append(l, s)//Append state that is passed in

	//Any state that has its symbol as 0, means it has edge arrows coming from it.
	if s != a && s.symbol == 0{
		l = addState(l, s.edge1, a)
		if s.edge2 != nil{
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

			//Does string po match string s.
func pomatch(po string, s string) bool{
	ismatch := false//Default position false.
	//Creating a NFA from the regular expression.
	ponfa := poregtonfa(po)
	
	//Keeping track of the current states of the NFA.
	current := []*state{}
	//Any state that you can get to from current.
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	//Loops through s a character at a time.
	for _, r := range s{
		//Loops through current array.
		for _, c := range current{
			if c.symbol == r{
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//Swapping current and next arrays after getting what all current states are, and setting next to an empty array.
		current, next = next, []*state{}
	}

	//Loops through the current array containing the states that we end up in.
	for _, c := range current{
		//if the state in the current array is equal to the accept state of ponfa.
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}

	return ismatch
}

func main(){
	fmt.Println(pomatch("ab.c*|" , "ccccccccccccc"))
}