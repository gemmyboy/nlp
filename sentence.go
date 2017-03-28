package nlp

/*
	sentence.go
		By Marcus Shannon

	Container(s) for passing on interpretted content detailed information
	about sentence processed
*/

//Sentence -: data structure used to hole entire input
type Sentence struct {
	raw   string //Raw Sentence
	words map[string]*Word
} //End Sentence

//Word -: container for grammar data
type Word struct {
	rawWord    string   //Raw word
	identifier []string //Array of grammar identifiers IE noun, verb, etc...
	self       bool     //Does this word reference self?
	unknown    bool     //Does nlp know this word?
} //End Word
