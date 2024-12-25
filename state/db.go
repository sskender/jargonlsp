package state

type DocumentItem struct {
	LanguageId string
	Text       string
	Version    uint
}

type StateDB struct {
	Documents map[string]*DocumentItem
}
