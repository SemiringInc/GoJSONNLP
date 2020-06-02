/**
 * JSONNLP package
 * (C) 2020 by Semiring Inc., Damir Cavar
 *
 * reading and writing JSON-NLP data.
 *
 * version 0.4
 */

package jsonnlp

import (
	"encoding/json"
	"io/ioutil"
)

const version string = "0.4"

type Meta struct {
	DCConformsTo  string `json:"DC.conformsTo"`
	DCCreated     string `json:"DC.created"`            // "2020-05-28T02:15:19"
	DCDate        string `json:"DC.date,omitempty"`     // "2020-05-28T02:15:19"
	DCSource      string `json:"DC.source,omitempty"`   // "NLP1 2.2.3"
	DCLanguage    string `json:"DC.language,omitempty"` // "en"
	DCCreator     string `json:"DC.creator,omitempty"`
	DCPublisher   string `json:"DC.publisher,omitempty"`
	DCTitle       string `json:"DC.title,omitempty"`
	DCDescription string `json:"DC.description,omitempty"`
	DCIdentifier  string `json:"DC.identifier,omitempty"`
}

type TokenFeatures struct {
	Overt bool `json:"overt,omitempty"`
	Stop  bool `json:"stop,omitempty"`
	Alpha bool `json:"alpha,omitempty"`
	//NounTypeProp bool `json:"NounType_prop"`
	Number         int    `json:"number,omitempty"` // 1 = singular, 2 = dual, 3 or more = plural
	Gender         string `json:"gender,omitempty"` // male, female, neuter
	Person         int    `json:"person,omitempty"` // 1, 2, 3
	Tense          string `json:"tense,omitempty"`  // past, present, future
	Perfect        bool   `json:"perfect,omitempty"`
	Continuous     bool   `json:"continuous,omitempty"`
	Case           string `json:"case,omitempty"`    // nom, acc, dat, gen, voc, loc, inst, ...
	Human          bool   `json:"human,omitempty"`   // yes/no
	Animate        bool   `json:"animate,omitempty"` // yes/no
	Negated        bool   `json:"negated,omitempty"` // word in scope og negation
	Countable      bool   `json:"countable,omitempty"`
	Factive        bool   `json:"factive,omitempty"` // factive verb
	Counterfactive bool   `json:"counterfactive,omitempty"`
	Irregular      bool   `json:"irregular,omitempty"` // irregular verb or noun form
	PhrasalVerb    bool   `json:"phrasalVerb,omitempty"`
	Mood           string `json:"mood,omitempty"` // indicative, imperative, subjunctive
	Foreign        bool   `json:"foreign,omitempty"`
	SpaceAfter     bool   `json:"spaceAfter,omitempty"` //: true
}

type TokenMisc struct {
	SpaceAfter bool // "SpaceAfter": true
}

type TokenList struct {
	ID                   int           `json:"id"`
	SentenceID           int           `json:"sentence_id"`
	Text                 string        `json:"text"`           // "John",
	Lemma                string        `json:"lemma"`          // "John",
	XPoS                 string        `json:"xpos,omitempty"` // "NNP",
	XPoSProbability      float64       `json:"xpos_prob,omitempty"`
	UPoS                 string        `json:"upos,omitempty"` // "PROPN",
	UPoSProbability      float64       `json:"upos_prob,omitempty"`
	EntityIOB            string        `json:"entity_iob,omitempty"` // "B",
	CharacterOffsetBegin int           `json:"characterOffsetBegin,omitempty"`
	CharacterOffsetEnd   int           `json:"characterOffsetEnd,omitempty"`
	PropID               string        `json:"propID,omitempty"`            // PropBank ID
	PropIDProbability    string        `json:"propIDProbability,omitempty"` // PropBank ID probability
	Lang                 string        `json:"lang,omitempty"`              // "en",
	Features             TokenFeatures `json:"features,omitempty"`          //
	// Misc                 TokenMisc     `json:"misc"`
	Shape  string `json:"shape,omitempty"`  // "Xxxx",
	Entity string `json:"entity,omitempty"` // "PERSON"
}

// this is a new structure compared to the original JSON-NLP version
type Sentence struct {
	ID        int    `json:"id"`                  // sentence ID
	TokenFrom int    `json:"tokenFrom,omitempty"` // first token
	TokenTo   int    `json:"tokenTo,omitempty"`   // last token
	Tokens    []int  `json:"tokens,omitempty"`    // list of tokens in sentence
	Clauses   []int  `json:"clauses,omitempty"`   // list of clauses in sentence
	Type      string `json:"type,omitempty"`      // type of sentence: declarative, interrogative, exclamatory, imperative, instructive
}

type Clause struct {
	ID        int    `json:"id"`                  // clause ID
	TokenFrom int    `json:"tokenFrom,omitempty"` // first token
	TokenTo   int    `json:"tokenTo,omitempty"`   // last token
	Tokens    []int  `json:"tokens,omitempty"`    // list of tokens
	Main      bool   `json:"main"`                // is it a main clause
	Governor  int    `json:"gov,omitempty"`       // the id of the governing clause
	Root      int    `json:"root,omitempty"`      // token ID of root (main verb or predicate head
	Negation  bool   `json:"neg,omitempty"`       // clause negated
	Tense     string `json:"tense,omitempty"`     //
	Voice     string `json:"voice,omitempty"`     //
	Mood      string `json:"mood,omitempty"`      //
}

type Dependency struct {
	Label     string `json:"lab"`
	Governor  int    `json:"gov"`
	Dependent int    `json:"dep"`
}

// a dependency tree is redefined compared to the original version of JSON-NLP
type DependencyTree struct {
	SentenceID   int          `json:"sentenceID"`
	Style        string       `json:"style,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type CoreferenceRepresentantive struct {
	Tokens []int `json:"tokens"`
	Head   int   `json:"head,omitempty"`
}

type CoreferenceReferents struct {
	Tokens []int `json:"tokens"`
	Head   int   `json:"head,omitempty"`
}

type Coreference struct {
	ID              int                        `json:"id"`
	Representartive CoreferenceRepresentantive `json:"representative"`
	Referents       []CoreferenceReferents     `json:"referents"`
}

type ConstituentParse struct {
	SentenceID        int     `json:"sentenceId"`
	Type              string  `json:"type,omitempty"`
	LabeledBracketing string  `json:"labeledBracketing"`
	Probability       float64 `json:"prob,omitempty"`
}

type Expression struct {
	ID         int    `json:"id"`
	Type       string `json:"type,omitempty"` // "NP"
	Head       int    `json:"head,omitempty"`
	Dependency string `json:"dependency,omitempty"` // "nsubj"
	Tokens     []int  `json:"tokens"`
}

type Document struct {
	MetaDocument    Meta               `json:"meta"`
	ID              int                `json:"id"`
	TokenList       []TokenList        `json:"tokenList,omitempty"`
	Sentences       []Sentence         `json:"sentences,omitempty"`
	Clauses         []Clause           `json:"clauses,omitempty"`
	DependencyTrees []DependencyTree   `json:"dependencyTrees,omitempty"`
	Coreferences    []Coreference      `json:"coreferences,omitempty"`
	Constituents    []ConstituentParse `json:"constituents,omitempty"`
	Expressions     []Expression       `json:"expressions,omitempty"`
}

type JSONNLP struct {
	MetaData  Meta       `json:"meta,omitempty"`
	Documents []Document `json:"documents,omitempty"`
}

func loadJSON(filename string) JSONNLP {
	file, _ := ioutil.ReadFile(filename)
	data := JSONNLP{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func getJSON(data JSONNLP) ([]byte, error) {
	return json.Marshal(data)
}
