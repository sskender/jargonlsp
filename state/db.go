package state

type DocumentItem struct {
	LanguageId string
	Text       string
	Version    int
}

type StateDB struct {
	Documents map[string]*DocumentItem
}
