package nlp

import "testing"
import "os"

//Test Instantiating an NLP
func TestCreate(t *testing.T) {
	_ = NewNLP(`c:\data.gob`)
} //End TestCreate()

//Test Instantiating an NLP
func TestSave(t *testing.T) {
	n := NewNLP(``)
	n.Save(`c:\Users\mshannon\Desktop\data.gob`)

	nn := NewNLP(`c:\Users\mshannon\Desktop\data.gob`)
	if len(n.Nouns) != len(nn.Nouns) ||
		len(n.Adjectives) != len(nn.Adjectives) ||
		len(n.Adverbs) != len(nn.Adverbs) ||
		len(n.Conjunctions) != len(nn.Conjunctions) ||
		len(n.Verbs) != len(nn.Verbs) ||
		len(n.Prepositions) != len(nn.Prepositions) {
		t.Error("Load/Save failed using default library")
	}

	os.Remove(`c:\Users\mshannon\Desktop\data.gob`)
} //End TestCreate()
