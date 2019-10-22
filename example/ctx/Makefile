regenerate:
	gocc -a -v -o ctx0.grammar/ctx0 ctx.bnf  	>ctx0.log 2>&1
	gocc -a -v -o ctx1.grammar/ctx1 ctx1.bnf		>ctx1.log 2>&1
	gocc -a -v -o ctx2.grammar/ctx2 -bugs=lexer_regdefs:ignore ctx.bnf	>ctx2.log 2>&1
