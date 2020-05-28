/**
 * JSONNLP package
 * (C) 2020 by Semiring Inc., Damir Cavar
 *
 * reading and writing JSON-NLP data.
 */

package jsonnlp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Meta struct {
	DCConformsTo string `json:"DC.conformsTo"`
	DCCreated    string `json:"DC.created"` // "2020-05-28T02:15:19"
	DCDate       string `json:"DC.date"`    // "2020-05-28T02:15:19"
}

type MetaDocument struct {
	DCConformsTo string `json:"DC.conformsTo"`
	DCCreated    string `json:"DC.created"`  // "2020-05-28T02:15:19"
	DCDate       string `json:"DC.date"`     // "2020-05-28T02:15:19"
	DCSource     string `json:"DC.source"`   // "NLP1 2.2.3",
	DCLanguage   string `json:"DC.language"` // "en"
}

type TokenFeatures struct {
	Overt bool // `json: "Overt"`
	Stop  bool // `json: "Stop"`
	Alpha bool // `json: "Alpha"`
	//NounTypeProp bool `json:"NounType_prop"`
	Number         int    // `json:"Number"` // 1 = singular, 2 = dual, 3 or more = plural
	Gender         string `json:"gender"` // male, female, neuter
	Person         int    `json:"person"` // 1, 2, 3
	Tense          string `json:"tense"`  // past, present, future
	Perfect        bool   `json:"perfect"`
	Continuous     bool   `json:"continuous"`
	Case           string `json:"case"`    // nom, acc, dat, gen, voc, loc, inst, ...
	Human          bool   `json:"human"`   // yes/no
	Animate        bool   `json:"animate"` // yes/no
	Negated        bool   `json:"negated"` // word in scope og negation
	Countable      bool   `json:"countable"`
	Factive        bool   `json:"factive"` // factive verb
	Counterfactive bool   `json:"counterfactive"`
	Irregular      bool   `json:"irregular"` // irregular verb or noun form
	PhrasalVerb    bool   `json:"phrasalVerb"`
	Mood           string `json:"mood"` // indicative, imperative, subjunctive
	Foreign        bool   // `json: "Foreign"`
	SpaceAfter     bool   // "SpaceAfter": true
}

type TokenMisc struct {
	SpaceAfter bool // "SpaceAfter": true
}

type TokenList struct {
	ID                   int           `json:"id"`
	SentenceID           int           `json:"sentence_id"`
	Text                 string        `json:"text"`  // "John",
	Lemma                string        `json:"lemma"` // "John",
	XPoS                 string        `json:"xpos"`  // "NNP",
	XPoSProbability      float64       `json:"xpos_prob"`
	UPoS                 string        `json:"upos"` // "PROPN",
	UPoSProbability      float64       `json:"upos_prob"`
	EntityIOB            string        `json:"entity_iob"` // "B",
	CharacterOffsetBegin int           `json:"characterOffsetBegin"`
	CharacterOffsetEnd   int           `json:"characterOffsetEnd"`
	PropID               string        `json:"propID"`            // PropBank ID
	PropIDProbability    string        `json:"propIDProbability"` // PropBank ID probability
	Lang                 string        `json:"lang"`              // "en",
	Features             TokenFeatures `json:"features"`          //
	// Misc                 TokenMisc     `json:"misc"`
	Shape  string `json:"shape"`  // "Xxxx",
	Entity string `json:"entity"` // "PERSON"
}

// this is a new structure compared to the original JSON-NLP version
type Sentence struct {
	ID        int   `json:"id"`        //
	TokenFrom int   `json:"tokenFrom"` //
	TokenTo   int   `json:"tokenTo"`   //
	Tokens    []int `json:"tokens"`    //
}

type Dependency struct {
	Label     string `json:"lab"`
	Governor  int    `json:"gov"`
	Dependent int    `json:"dep"`
}

// a dependency tree is redefined compared to the original version of JSON-NLP
type DependencyTree struct {
	SentenceID   int          `json:"sentenceID"`
	Style        string       `json:"style"`
	Dependencies []Dependency `json:"dependencies"`
}

type CoreferenceRepresentantive struct {
	Tokens []int `json:"tokens"`
	Head   int   `json:"head"`
}

type CoreferenceReferents struct {
	Tokens []int `json:"tokens"`
	Head   int   `json:"head"`
}

type Coreference struct {
	ID              int                        `json:"id"`
	Representartive CoreferenceRepresentantive `json:"representative"`
	Referents       []CoreferenceReferents     `json:"referents"`
}

type ConstituentParse struct {
	SentenceID        int     `json:"sentenceId"`
	Type              string  `json:"type"`
	LabeledBracketing string  `json:"labeledBracketing"`
	Probability       float64 `json:"prob"`
}

type Expression struct {
	ID         int    `json:"id"`
	Type       string `json:"type"` // "NP"
	Head       int    `json:"head"`
	Dependency string `json:"dependency"` // "nsubj"
	Tokens     []int  `json:"tokens"`
}

type Document struct {
	MetaDocument    MetaDocument       `json:"meta"`
	ID              int                `json:"id"`
	TokenList       []TokenList        `json:"tokenList"`
	Sentences       []Sentence         `json:"sentences"`
	DependencyTrees []DependencyTree   `json:"dependencyTrees"`
	Coreferences    []Coreference      `json:"coreferences"`
	Constituents    []ConstituentParse `json:"constituents"`
	Expressions     []Expression       `json:"expressions"`
}

type JSONNLP struct {
	MetaData  Meta       `json:"meta"`
	Documents []Document `json:"documents"`
}

// func loadJSON()

func main() {
	file, _ := ioutil.ReadFile("example1.json")

	data := JSONNLP{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Documents); i++ {
		fmt.Println("Product Id: ", data.Documents[i].MetaDocument)
		fmt.Println("Quantity: ", data.Documents[i].Constituents[0].LabeledBracketing)
	}
}
