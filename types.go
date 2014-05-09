// © 2014 Steve McCoy. See LICENSE for details.

package main

// Term is either an Atom, Variable, or Compound Term.
// I don't think I'll ever do numbers or strings.
type Term interface{}

// Atom is a name such as cat, 'blah blah', etc. true and false are special, I think.
type Atom string

// Variable is a name starting with a capital letter, such as Cat, whose value may need to be resolved.
type Variable string

// Compound is a Term of Terms, e.g. list(a, b, bag(d, e, f)).
type Compound struct{
	Functor Atom
	Arguments []Term
}
