package nlp

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

/*
	nlp.go
		by Marcus Shannon

	Natural Language Processing

	Natural Language Processing is the concept of taking regular human language (english in this instance)
	and converting into something a computer can understand, use, and interact with.

	Ambiguity is a significant issue in NLP. There's ambiguity in what any-one word would mean in a sentence. Improper
	usage of tense in a sentence, incomplete sentences, contextual situations, etc... All these issues come up with up
	when discussing NLP.

	Hence, I am of the personal opinion that creating a formal grammar for a language is generally, useless. Human beings
	do not learn english using a formal grammar, we learn by first defining noun(Entities) and then verbs(Actions/Functions).
	Eventually general categories of sentences begin to appear such as:

	:-- "I'm hungry" - "I am hungry"

	Generally, this is just a statement, a declaration which references Self or the current state of Self. Were I to tell this
	to an AI, this would be description of my current state. (Assuming the AI knows what hungry even means).

	Another common sentence structure formulated are Imperative sentences or generally just informing an entity to perform an
	action of some sort such as:

	:-- "Produce food" - "Make me food" - "Make me some food"

	These are extremely command like and demand an action be taken immediately.

	The third and probably most important sentence type which arises is the Question or Interrogative sentence. This will be used
	to request information about an entity, action, context, etc... The more information known about something, the more questions
	to be asked/not asked.

	:-- "What kind of food do you have?" - "Where is the food?" - "Do you want food?" - "Are you hungry?"

	The fourth sentence which arises are conditional statements. This is pre-emptively, a pre-cursor to thinking logically in terms
	of an ambigious language.

	:-- "If you are hungry then eat some food." - "If you are not hungry then don't eat any food." -
	    "You are happy if you are not hungry." - "You will be hungry when there's no food in your stomache."

	------------------------------------------

	A combination of conditional statements and declarative statements are useful for building out the knowledge-base for an AI.
	This would also be useful for building out the personality in which an AI would use to help govern it's world. Due to the
	complexity in which things can become. Generally, using this NLP module in common with a specialized AI is most useful for
	when delving into a particular topic/field.

	------------------------------------------

	For the scope of this nlp program, it will generalize all incoming sentences into the follow 4 categories listed above
	or atleast attempt to anyways. This identification will only happen AFTER parsing through the sentence and assigning
	general grammar terms to each word. Like identifying the entity(noun) & action(verb) being referenced.

	Algorithm:
		- Check for number of sentences in input
		- Loop through sentences in order:
		-	Identify Noun(s)
		-	Identify Verb(s)
		-	Identify Adjective(s) - (Attributes about the noun)
		-	Identify Adverb(s) - (Attributes about the verb, adjective, or adverb)
		-	Identify Conjunction(s) - (Used to separate sub-sentences)
		-	Identify Preposition(s) - (Used to define relationship)
		-	Identify Instance(s) of Self in given context
		-
		-	Reading each word from left to right, create generic
		-		relationships between entities, actions, and attributes which
		-		describe them.  This part should take grammar into account such
		-		as a listing of words, time, etc...
		-
		-	Determine sentence type using typical identifiers.
		- 		(ex: question -> ?, command -> sentence starts with verb)
		-
		-	Using the relationships created, formalize the sentence into something
		-	similiar to an Abstract-Syntax-Tree but essentially a data-structure.



*/

//NLP -: Natural Language Processing
type NLP struct {
	Nouns        map[string]struct{} //List of known nouns
	Verbs        map[string]struct{} //List of known verbs
	Adjectives   map[string]struct{} //List of known adjectives
	Adverbs      map[string]struct{} //List of known adverbs
	Conjunctions map[string]struct{} //List of known conjunctions
	Prepositions map[string]struct{} //List of known prepositions
} //End NLP

//NewNLP -: Instantiate a new NLP instance
func NewNLP(path string) *NLP {
	nlp := new(NLP)

	//Check to see if an nlp instance exists
	if _, err := os.Stat(path); err != nil {
		nlp = newNLPInstance(nlp)
	} else {
		//Load NLP Instance
		file, er := os.Open(path)
		fi, errr := file.Stat()
		if er != nil && errr != nil {
			log.Fatal("Failed to open NLP instance;", er, errr)
		}

		//Gobble nlp instance
		buf := make([]byte, int(fi.Size()))
		c, errrr := file.Read(buf)
		if errrr != nil || c != int(fi.Size()) {
			log.Fatal(errrr)
		}

		decoder := gob.NewDecoder(bytes.NewBuffer(buf))
		er = decoder.Decode(&nlp)
		if er != nil {
			log.Fatal(er)
		}
		file.Close()
	}

	return nlp
} //End NewNLP()

//newNLPInstance -: Loads default values for list of known words
func newNLPInstance(nlp *NLP) *NLP {
	v := struct{}{}

	nlp.Nouns = map[string]struct{}{
		"object": v, "time": v, "word": v,
		"place": v, "question": v, "number": v,
		"fact": v, "thing": v, "entity": v,
		"yes": v, "no": v,
	} //End nouns

	nlp.Verbs = map[string]struct{}{
		"be": v, "were": v, "been": v, "have": v,
		"had": v, "do": v, "did": v, "done": v,
		"say": v, "said": v, "go": v, "went": v,
		"gone": v, "get": v, "got": v, "gotten": v,
		"make": v, "made": v, "know": v, "knew": v,
		"known": v, "think": v, "thought": v, "take": v,
		"took": v, "taken": v, "come": v, "came": v,
		"want": v, "wanted": v, "use": v, "used": v,
		"find": v, "found": v, "give": v, "gave": v,
		"given": v, "tell": v, "told": v, "work": v,
		"worked": v, "call": v, "called": v, "try": v,
		"tried": v, "ask": v, "asked": v, "need": v,
		"needed": v, "feel": v, "felt": v, "become": v,
		"became": v, "leave": v, "left": v, "put": v,
		"mean": v, "meant": v, "keep": v, "kept": v,
		"let": v, "begin": v, "begun": v, "began": v,
		"seem": v, "seemed": v, "help": v, "helped": v,
		"show": v, "showed": v, "shown": v, "hear": v,
		"heard": v, "see": v, "saw": v, "run": v,
		"ran": v, "move": v, "moved": v, "live": v,
		"lived": v, "believe": v, "believed": v, "bring": v,
		"brought": v, "happen": v, "happened": v, "write": v,
		"written": v, "wrote": v, "read": v, "sit": v,
		"sat": v, "stand": v, "stood": v, "lose": v,
		"lost": v, "pay": v, "paid": v, "meet": v,
		"met": v, "include": v, "continue": v, "stop": v,
		"join": v, "joined": v, "set": v, "learn": v,
		"learned": v, "lead": v, "led": v, "follow": v,
		"followed": v, "understand": v, "understood": v,
		"watch": v, "watched": v, "create": v, "created": v,
		"destroy": v, "destroyed": v, "spend": v, "spent": v,
		"consider": v, "considered": v, "offer": v, "offered": v,
		"serve": v, "served": v, "die": v, "died": v, "send": v,
		"sent": v, "receive": v, "received": v, "who": v,
		"what": v, "when": v, "where": v, "why": v, "how": v,
		"does": v,
	} //End verbs

	nlp.Adjectives = map[string]struct{}{
		"different": v, "other": v, "new": v, "old": v,
		"good": v, "bad": v, "big": v, "small": v, "little": v,
		"large": v, "long": v, "short": v, "black": v, "white": v,
		"early": v, "late": v, "hard": v, "soft": v, "major": v,
		"minor": v, "better": v, "worse": v, "full": v, "empty": v,
		"local": v, "regional": v, "recent": v, "clear": v, "national": v,
		"easy": v, "available": v, "likely": v, "unlikely": v,
		"single": v, "wrong": v, "right": v, "private": v, "public": v,
		"past": v, "rich": v, "ready": v, "simple": v, "complex": v,
		"general": v, "precise": v, "accurate": v, "rude": v,
		"physical": v, "abstract": v, "nice": v, "mean": v,
		"final": v, "popular": v, "similiar": v, "dark": v,
		"dead": v, "alive": v, "hot": v, "cold": v, "light": v,
	} //End adjectives

	nlp.Adverbs = map[string]struct{}{
		"accidentally": v, "afterwards": v, "almost": v, "always": v,
		"angrily": v, "annually": v, "anxiously": v, "awkwardly": v,
		"badly": v, "blindly": v, "boastfully": v, "boldly": v,
		"bravely": v, "briefly": v, "brightly ": v, "busily": v,
		"calmly": v, "carefully": v, "carelessly": v, "cautiously": v,
		"cheerfully": v, "clearly": v, "correctly": v, "courageously ": v,
		"crossly": v, "cruelly": v, "daily": v, "defiantly": v,
		"deliberately": v, "doubtfully": v, "easily": v, "elegantly": v,
		"enormously": v, "enthusiastically": v, "equally": v, "even": v,
		"eventually": v, "exactly": v, "faithfully": v, "far": v,
		"fast": v, "fatally": v, "fiercely": v, "fondly": v,
		"foolishly": v, "fortunately": v, "frantically": v, "gently": v,
		"gladly": v, "gracefully": v, "greedily": v, "happily": v,
		"hastily": v, "honestly": v, "hourly": v, "hungrily": v,
		"innocently": v, "inquisitively": v, "irritably": v, "joyously": v,
		"justly": v, "kindly": v, "lazily": v, "less": v,
		"loosely": v, "loudly": v, "madly": v, "merrily": v,
		"monthly": v, "more": v, "mortally": v, "mysteriously": v,
		"nearly": v, "neatly": v, "nervously": v, "never": v, "noisily": v,
		"not": v, "obediently": v, "obnoxiously": v, "often": v,
		"only": v, "painfully": v, "perfectly": v, "politely": v,
		"poorly": v, "powerfully": v, "promptly": v, "punctually": v,
		"quickly": v, "quietly": v, "rapidly": v, "rarely": v,
		"really": v, "recklessly": v, "regularly": v, "reluctantly": v,
		"repeatedly": v, "rightfully": v, "roughly": v, "rudely": v,
		"sadly": v, "safely": v, "seldom": v, "selfishly": v,
		"seriously": v, "shakily": v, "sharply": v, "shrilly": v,
		"shyly": v, "silently": v, "sleepily": v, "slowly": v,
		"smoothly": v, "softly": v, "solemnly": v, "sometimes": v,
		"soon": v, "speedily": v, "stealthily": v, "sternly": v,
		"successfully": v, "suddenly": v, "suspiciously": v, "swiftly": v,
		"tenderly": v, "tensely": v, "thoughtfully": v, "tightly": v,
		"tomorrow": v, "too": v, "truthfully": v,
		"unexpectedly": v, "very": v, "victoriously": v, "violently": v,
		"vivaciously": v, "warmly": v, "weakly": v, "wearily": v,
		"well": v, "wildly": v, "yearly": v, "yesterday": v,
	} //End adverbs

	nlp.Conjunctions = map[string]struct{}{
		"and": v, "but": v, "or": v, "yet": v,
		"for": v, "nor": v, "so": v, "after": v, "although": v,
		"because": v, "before": v, "by the time": v, "as": v,
		"even if": v, "even though": v, "every time": v,
		"if": v, "in case": v, "now that": v, "once": v,
		"since": v, "so that": v, "than": v, "the first time": v,
		"unless": v, "until": v, "when": v, "whenever": v,
		"whether or not": v, "while": v, "why": v,
	} //End conjunctions

	nlp.Prepositions = map[string]struct{}{
		"about": v, "above": v, "across": v, "after": v, "against": v,
		"along": v, "among": v, "around": v, "at": v, "before": v,
		"behind": v, "below": v, "beneath": v, "beside": v, "between": v,
		"beyond": v, "but": v, "by": v, "concerning": v, "down": v,
		"during": v, "except": v, "for": v, "from": v, "in": v,
		"inside": v, "into": v, "like": v, "near": v, "of": v,
		"off": v, "on": v, "onto": v, "out": v, "outside": v,
		"over": v, "past": v, "regarding": v, "since": v, "through": v,
		"to": v, "toward": v, "towards": v, "under": v, "underneath": v,
		"until": v, "up": v, "upon": v, "with": v, "within": v, "without": v,
	} //End prepositions

	return nlp
} //End newNLPInstance()

//Save -: Saves instance of the NLP to path
func (nlp *NLP) Save(path string) {
	var file *os.File
	var err error

	file, err = os.OpenFile(path, os.O_RDWR, 0667)
	defer file.Close()
	if err != nil {
		err = nil
		file, err = os.Create(path)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	//Save File
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(nlp)
	if err != nil {
		panic(err)
	}

	//Write to file
	c, er := file.Write(buf.Bytes())
	if er != nil || c != buf.Len() {
		log.Println(er, c)
		os.Exit(1)
	}
} //End Save()
