[![Build Status](https://drone.io/github.com/goccmack/gocc/status.png)](https://drone.io/github.com/goccmack/gocc/latest)

Release Notes for gocc 2.1
==========================

Changes:
========

1. no_lexer option added to suppress generation of lexer. See the user guide.

2. Unreachable code removed from generated code.

Bugs fixed:
===========

1. gocc 2.1 does not support string_lit symbols with the same value as production names of the BNF. E.g. (t2.bnf):

A : "a" | "A" ;

string_lit "A" is not allowed. 

Previously gocc silently ignored the conflicting string_lit. Now it generates an ugly panic:

> gocc t2.bnf
panic: string_lit "A" conflicts with production name A

This issue will be properly resolved in a future release.

