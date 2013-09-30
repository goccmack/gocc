/*
This package contains the Abstract Syntax Tree (AST) elements used by gocc to generate a target lexer and parser.

The top-level node is Grammar in grammar.go

The EBNF accepted by gocc consists of two parts:

1. The lexical part, containing the defintion of tokens.

2. The grammar or syntax part, containing the grammar or syntax of the language. Files containing grammar objects are prefixed with "g", e.g.: galts.go
*/
package ast
