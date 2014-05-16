%union {
	atom string
	op string
}

%token <atom> ATOM
%token <op> OP

%%

top:
	atom
|	atom op atom
