%{
package grammar

import (
	"github.com/mccoyst/amalog"
)

%}

%union {
	atom amalog.Atom
	var amalog.Variable
	op string
	term amalog.Term
	comp amalog.Compound
	stmts []Statement
}

%token <atom> ATOM
%token <op> OP

%type <stmts> top

%%

top:
	stmts

stmts:
	stmts stmt
	stmt

stmt:
	ATOM
|	ATOM op term
|	term op term

term:
	ATOM
|	var
|	comp

comp:
	ATOM '(' terms ')'

terms:
	terms ',' term
|	term
