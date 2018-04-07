# GoLangRegExp
A program in the Go programming language that can build a non-deterministic ﬁnite automaton (NFA) from a regular expression.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “ ∗” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1 ∗ means any number of 1’s.

### Method
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.

### Running the program
1. git clone https://github.com/Jawneck/GoLangRegExp.git
2. Navigate to the local repository
3. go run rega.go

### References
GMIT. Quality assurance framework.
https://www.gmit.ie/general/quality- assurance- framework.

Google. The go programming language.
https://golang.org/.

Video Streams
https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d


