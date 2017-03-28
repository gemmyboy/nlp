package nlp

/*
	sentence.go
		By Marcus Shannon

	Container(s) for passing on interpretted content detailed information
	about sentence processed
*/

//Sentence -: data structure used to hole entire input
type Sentence struct {
	sType string //Sentence type
	raw   string //Raw Sentence
	words []*Word
} //End Sentence

//Word -: container for grammar data
type Word struct {
	rawWord string //Raw word

	//Semantical Definition
	tense   int  //unknown: 0, past: 1, present: 2, future: 3
	self    bool //Does this word reference self?
	unknown bool //Does nlp know this word?
	pov     int  //unknown: 0, 1st: 1, 2nd: 2, 3rd: 3

	//Words can be ambigious. Context will clarify this later
	//Syntactical Identification
	noun         bool
	verb         bool
	adjective    bool
	adverb       bool
	conjunction  bool
	prepositions bool
	syntaxID     int
} //End Word
