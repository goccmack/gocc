Lexical and Parser user-defined context data
============================================

For advanced use cases, it is sometimes helpful to be able to have
additional data available in either the parsing context, such as
sdt actions, or the lexical context, on the lexer or the tokens
it produces.

Distinction between Lexical and Parser Context
----------------------------------------------

Both `parser` and `token` define `Context` as `interface{}`.

Gocc does not care if the parser and lexical contexts are the same, or
whether you use one and not the other.

Lexical Context
---------------

When the lexer.Lexer object is configured with a non-nil Context value,
the value will be assigned to the Pos.Context property of every token
the lexer generates.

If you are parsing multiple files, this could allow you to inform the
user that tokenA is redeclared in file2.txt after being previously
declared in file1.txt, rather than simply "I've seen tokenA before".

Assigning:
```
	// assigning
	l := lexer.NewLexer()
	l.Context = &MyContext{ /* populate fields */ }
	// or just
	l.Context = filename
```

Use (assumes MyContext has method 'Source()' that returns filename)
```
	func duplicate(newToken, oldToken *token.Token) error {
		newPos, oldPos := newToken.Pos, oldToken.Pos

		return fmt.Errorf("%s:%d:%d: error: '%s' redefined.\n%s:%d:%d: previous definition",
			newPos.Context.(*MyContext).Source(), newPos.Line, newPos.Column,
			string(newToken.Lit),
			oldPos.Context.(*MyContext).Source(), oldPos.Line, oldPos.Column)
```

Parser Context
--------------

This allows you to surface contextual information when invoking SDT actions. It
corresponds to the `Context` property of the `parser.Parser` class, and can be
accessed with the SDT token `$Context`.

Example:

_bnf_
```
	Identifier: identifier << ast.NewIdentifier($0, $Context);
```

_ast/main.go_
```
	func NewIdentifier(nameAttr Attrib, context interface{}) (string, error) {
		name := string(nameAttr.(*token.Token).Lit)
		if previous, exists := context.(*ParseContext).Identifiers[name]; exists {
			...
```

