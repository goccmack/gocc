#!/bin/bash
for dir in $(find . -mindepth 1 -maxdepth 1 -type d) ; do ( echo -e "\n\n\n******\nTest: $(basename $dir)\n******" ;cd $dir ; rm -rf *.grammar ; make ; go test -v ) ; done