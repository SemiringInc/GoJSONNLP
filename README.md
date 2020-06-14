# GoJSONNLP - JSON-NLP Go Code

(C) 2020 by [Semiring Inc.], [Damir Cavar]

Package version 0.5


## Introduction

This repository provides the [Go] package *jsonnlp* for reading and writing [JSON-NLP] Schema conform data. [JSON-NLP] encodes outputs from Natural Language Processing (NLP) pipelines, functioning as some form of a middleware.

[JSON-NLP] wrappers for the output formats from various NLP pipelines are available:

- [Flair](https://github.com/flairNLP/flair) and [Flair-JSON-NLP](https://github.com/dcavar/Flair-JSON-NLP)
- [NLTK](http://nltk.org/) and [NLTK-JSON-NLP](https://github.com/dcavar/NLTK-JSON-NLP)
- [Polyglot](https://github.com/aboSamoor/polyglot) and [Polyglot-JSON-NLP](https://github.com/dcavar/Polyglot-JSON-NLP)
- [spaCy](https://spacy.io/) and [spaCy-JSON-NLP](https://github.com/dcavar/spaCy-JSON-NLP)
- [Xrenner](https://github.com/amir-zeldes/xrenner) and [Xrenner-JSON-NLP](https://github.com/dcavar/Xrenner-JSON-NLP)
- [Stanford CoreNLP](https://stanfordnlp.github.io/CoreNLP/)
- [OpenNLP](https://opennlp.apache.org/)

Many other wrappers and modules likely exist or will be made available.

[JSON-NLP] processing and validation modules exist for other languages as well, as for example:

- Java: [J-JSON-NLP](https://github.com/dcavar/J-JSON-NLP)
- Python: [Py-JSON-NLP](https://github.com/dcavar/Py-JSON-NLP)


## Installation

Install the *jsonnlp* [Go] package using:

    go get github.com/SemiringInc/GoJSONNLP

Update to new version using:

    go get -u github.com/SemiringInc/GoJSONNLP



[Semiring Inc.]: https://semiring.com/ "Semiring Inc."
[Damir Cavar]: http://damir.cavar.me/ "Damir Cavar"
[JSON-NLP]: https://github.com/SemiringInc/JSON-NLP "JSON-NLP"
[Go]: https://golang.org/ "Golang"
