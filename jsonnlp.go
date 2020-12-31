/**
 * JSONNLP package
 * (C) 2020 by Semiring Inc., Damir Cavar
 *
 * reading and writing JSON-NLP data.
 *
 * version 0.8.3
 */

package jsonnlp

import (
	"encoding/json"
	"io/ioutil"
)

const version string = "0.8.3"

// Meta contains the common meta information for the entire JSON-NLP or a single document.
// These are Dublin Core (DC) labels. See the DC documentation for details.
type Meta struct {
	DCConformsTo  string `json:"DC.conformsTo"`
	DCAuthor      string `json:"DC.author"`             //
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

// TokenFeatures is a data structure that containes all the detailed morphosyntactic token features.
type TokenFeatures struct {
	Overt          bool   `json:"overt,omitempty"`       // is the token overt? Invisible or covert words are words that are omitted in speech, subject to ellipsis, gapping, simple object, topic, or subject drop, etc.
	Stop           bool   `json:"stop,omitempty"`        // is the token a stop-word or not?
	Alpha          bool   `json:"alpha,omitempty"`       //
	Number         int    `json:"number,omitempty"`      // 1 = singular, 2 = dual, 3 or more = plural
	Gender         string `json:"gender,omitempty"`      // male, female, neuter
	Person         int    `json:"person,omitempty"`      // 1, 2, 3
	Tense          string `json:"tense,omitempty"`       // Tense of the token: past, present, future
	Perfect        bool   `json:"perfect,omitempty"`     // Aspect of the token
	Continuous     bool   `json:"continuous,omitempty"`  // is the token indicating continuous = ing
	Progressive    bool   `json:"progressive,omitempty"` // is the token indicating progressive = am + ...ing
	Case           string `json:"case,omitempty"`        // nom, acc, dat, gen, voc, loc, inst, ...
	Human          bool   `json:"human,omitempty"`       // yes/no
	Animate        bool   `json:"animate,omitempty"`     // yes/no
	Negated        bool   `json:"negated,omitempty"`     // word in scope og negation
	Countable      bool   `json:"countable,omitempty"`
	Factive        bool   `json:"factive,omitempty"` // factive verb
	Counterfactive bool   `json:"counterfactive,omitempty"`
	Irregular      bool   `json:"irregular,omitempty"` // irregular verb or noun form
	PhrasalVerb    bool   `json:"phrasalVerb,omitempty"`
	Mood           string `json:"mood,omitempty"` // indicative, imperative, subjunctive
	Foreign        bool   `json:"foreign,omitempty"`
	SpaceAfter     bool   `json:"spaceAfter,omitempty"` // space after token in orig text?
}

// Token structure contains all the token spoecific details.
type Token struct {
	ID                   int           `json:"id"`
	SentenceID           int           `json:"sentence_id"`
	Text                 string        `json:"text"`            // "John",
	Lemma                string        `json:"lemma,omitempty"` // "John",
	XPoS                 string        `json:"xpos,omitempty"`  // "NNP",
	XPoSProbability      float64       `json:"xpos_prob,omitempty"`
	UPoS                 string        `json:"upos,omitempty"` // "PROPN",
	UPoSProbability      float64       `json:"upos_prob,omitempty"`
	EntityIOB            string        `json:"entity_iob,omitempty"` // "B",
	CharacterOffsetBegin int           `json:"characterOffsetBegin,omitempty"`
	CharacterOffsetEnd   int           `json:"characterOffsetEnd,omitempty"`
	PropID               string        `json:"propID,omitempty"`            // PropBank ID
	PropIDProbability    float64       `json:"propIDProbability,omitempty"` // PropBank ID probability
	FrameID              int           `json:"frameID,omitempty"`
	FrameIDProbability   float64       `json:"frameIDProb,omitempty"`
	WordNetID            int           `json:"wordNetID,omitempty"`
	WordNetIDProbability float64       `json:"wordNetIDProb,omitempty"`
	VerbNetID            int           `json:"verbNetID,omitempty"`
	VerbNetIDProbability float64       `json:"verbNetIDProb,omitempty"`
	Lang                 string        `json:"lang,omitempty"`     // "en",
	Features             TokenFeatures `json:"features,omitempty"` //
	Shape                string        `json:"shape,omitempty"`    // "Xxxx",
	Entity               string        `json:"entity,omitempty"`   // "PERSON"
}

// Sentence is a new structure compared to the original JSON-NLP version.
type Sentence struct {
	ID                   int     `json:"id"`                      // sentence ID
	TokenFrom            int     `json:"tokenFrom,omitempty"`     // first token
	TokenTo              int     `json:"tokenTo,omitempty"`       // last token
	Tokens               []int   `json:"tokens,omitempty"`        // list of tokens in sentence
	Clauses              []int   `json:"clauses,omitempty"`       // list of clauses in sentence
	Type                 string  `json:"type,omitempty"`          // type of sentence: declarative, interrogative, exclamatory, imperative, instructive
	Sentiment            string  `json:"sentiment,omitempty"`     // sentiment type
	SentimentProbability float64 `json:"sentimentProb,omitempty"` //
}

// Clause contains information about clause level properties.
type Clause struct {
	ID                   int     `json:"id"`                  // clause ID
	SentenceID           int     `json:"sentenceID"`          // sentence ID
	TokenFrom            int     `json:"tokenFrom,omitempty"` // first token
	TokenTo              int     `json:"tokenTo,omitempty"`   // last token
	Tokens               []int   `json:"tokens,omitempty"`    // list of tokens
	Main                 bool    `json:"main,omitempty"`      // is it a main clause
	Governor             int     `json:"gov,omitempty"`       // the id of the governing clause
	Head                 int     `json:"head,omitempty"`      // token ID of root/head (main verb or predicate head
	Negation             bool    `json:"neg,omitempty"`       // clause negated
	Tense                string  `json:"tense,omitempty"`     //
	Mood                 string  `json:"mood,omitempty"`      //
	Perfect              bool    `json:"perfect,omitempty"`
	Continuous           bool    `json:"continuous,omitempty"`
	Aspect               string  `json:"aspect,omitempty"`        //
	Voice                string  `json:"voice,omitempty"`         //
	Sentiment            string  `json:"sentiment,omitempty"`     //
	SentimentProbability float64 `json:"sentimentProb,omitempty"` //
}

// Dependency tree encoding in JSON-NLP.
type Dependency struct {
	Label       string  `json:"lab"`
	Governor    int     `json:"gov"`
	Dependent   int     `json:"dep"`
	Probability float64 `json:"prob,omitempty"`
}

// DependencyTree is a dependency tree is redefined compared to the original version of JSON-NLP.
type DependencyTree struct {
	SentenceID    int          `json:"sentenceID"`
	Style         string       `json:"style,omitempty"`
	Dependencies  []Dependency `json:"dependencies,omitempty"`
	Probability   float64      `json:"prob,omitempty"`
	HashOverHeads int          `json:"hashhead,omitempty"`
}

// CoreferenceRepresentantive is
type CoreferenceRepresentantive struct {
	Tokens []int `json:"tokens"`
	Head   int   `json:"head,omitempty"`
}

// CoreferenceReferents is
type CoreferenceReferents struct {
	Tokens      []int   `json:"tokens"`
	Head        int     `json:"head,omitempty"`
	Probability float64 `json:"prob,omitempty"`
}

// Coreference is
type Coreference struct {
	ID             int                        `json:"id"`
	Representative CoreferenceRepresentantive `json:"representative"`
	Referents      []CoreferenceReferents     `json:"referents"`
}

// Scope is
type Scope struct {
	ID         int   `json:"id"`
	Governor   []int `json:"gov"`
	Dependents []int `json:"dep,omitempty"`
	Terminals  []int `json:"terminals,omitempty"`
}

// ConstituentParse is
type ConstituentParse struct {
	SentenceID        int     `json:"sentenceId"`
	Type              string  `json:"type,omitempty"`
	LabeledBracketing string  `json:"labeledBracketing"`
	Probability       float64 `json:"prob,omitempty"`
	Scopes            []Scope `json:"scopes,omitempty"`
}

// Expression is
type Expression struct {
	ID          int     `json:"id"`
	Type        string  `json:"type,omitempty"` // "NP"
	Head        int     `json:"head,omitempty"`
	Dependency  string  `json:"dependency,omitempty"` // "nsubj"
	TokenFrom   int     `json:"tokenFrom,omitempty"`  // first token
	TokenTo     int     `json:"tokenTo,omitempty"`    // last token
	Tokens      []int   `json:"tokens"`
	Probability float64 `json:"prob,omitempty"`
}

// Paragraph is
type Paragraph struct {
	ID        int   `json:"id"`
	TokenFrom int   `json:"tokenFrom,omitempty"`
	TokenTo   int   `json:"tokenTo,omitempty"`
	Tokens    []int `json:"tokens,omitempty"`
	Sentences []int `json:"sentences,omitempty"`
}

// Attribute is
type Attribute struct {
	Label string `json:"lab"`
	Value string `json:"val"`
}

// Entity is
type Entity struct {
	ID                   int         `json:"id"`
	Label                string      `json:"label,omitempty"`
	Type                 string      `json:"type"`
	URL                  string      `json:"url"`
	Head                 int         `json:"head,omitempty"`
	TokenFrom            int         `json:"tokenFrom,omitempty"`
	TokenTo              int         `json:"tokenTo,omitempty"`
	Tokens               []int       `json:"tokens,omitempty"`
	TripleID             int         `json:"tripleID,omitempty"`      // reified entity pointer to triple ID
	Sentiment            string      `json:"sentiment,omitempty"`     //
	SentimentProbability float64     `json:"sentimentProb,omitempty"` //
	Count                int         `json:"count,omitempty"`
	Attributes           []Attribute `json:"attributes"`
}

// Relation is
type Relation struct {
	ID                   int         `json:"id"`
	Label                string      `json:"label"`
	Type                 string      `json:"type"`
	URL                  string      `json:"url"`
	Head                 int         `json:"head,omitempty"`
	TokenFrom            int         `json:"tokenFrom,omitempty"`
	TokenTo              int         `json:"tokenTo,omitempty"`
	Tokens               []int       `json:"tokens,omitempty"`
	Sentiment            string      `json:"sentiment,omitempty"`     //
	SentimentProbability float64     `json:"sentimentProb,omitempty"` //
	Count                int         `json:"count,omitempty"`
	Attributes           []Attribute `json:"attributes"`
}

// Triple is
type Triple struct {
	ID               int     `json:"id"`
	FromEntity       int     `json:"fromEntity"`
	ToEntity         int     `json:"toEntity"`
	Relation         int     `json:"rel"`
	ClauseID         []int   `json:"clauseID,omitempty"`
	SentenceID       []int   `json:"sentenceID,omitempty"`
	Directional      bool    `json:"directional,omitempty"`
	EventID          int     `json:"eventID,omitempty"`
	TemporalSequence int     `json:"tempSeq,omitempty"`
	Probability      float64 `json:"prob,omitempty"`
	Syntactic        bool    `json:"syntactic,omitempty"`
	Implied          bool    `json:"implied,omitempty"`
	Presupposed      bool    `json:"presupposed,omitempty"`
	Count            int     `json:"count,omitempty"`
}

// Document is
type Document struct {
	MetaDocument    Meta               `json:"meta"`
	ID              int                `json:"id"`
	TokenList       []Token            `json:"tokenList,omitempty"`
	Clauses         []Clause           `json:"clauses,omitempty"`
	Sentences       []Sentence         `json:"sentences,omitempty"`
	Paragraphs      []Paragraph        `json:"paragraphs,omitempty"`
	DependencyTrees []DependencyTree   `json:"dependencyTrees,omitempty"`
	Coreferences    []Coreference      `json:"coreferences,omitempty"`
	Constituents    []ConstituentParse `json:"constituents,omitempty"`
	Expressions     []Expression       `json:"expressions,omitempty"`
	Entities        []Entity           `json:"entities,omitempty"`
	Relations       []Relation         `json:"relations,omitempty"`
	Triples         []Triple           `json:"triples,omitempty"`
}

// JSONNLP is
type JSONNLP struct {
	MetaData  Meta       `json:"meta,omitempty"`
	Documents []Document `json:"documents,omitempty"`
}

// FromString reads the JSON-NLP instance from a string.
func (data *JSONNLP) FromString(t string) {
	// TODO check whether data has any content
	_ = json.Unmarshal([]byte(t), data)
}

// FromFile reads the JSON-NLP instance from a file.
func (data *JSONNLP) FromFile(filename string) {
	// TODO check whether data has any content
	file, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(file), data)
}

// GetJSON returns the JSON-NLP instance as a byte array.
func (data *JSONNLP) GetJSON() ([]byte, error) {
	return json.Marshal(data)
}
