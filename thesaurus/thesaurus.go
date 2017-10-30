package thesaurus

type Thesaurus interface {
	Synonyms(term String) ([]string, error)
}
