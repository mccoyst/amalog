package grammar

type lex struct{
	line []byte
	peek rune
}

func (l *lex) Lex(lval *yySymType) int {
	return 0
}

func (l *lex) Error(e string) {

}
