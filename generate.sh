#!/usr/bin/env bash

antlr4='java -jar /usr/local/lib/antlr-4.7.1-complete.jar'
${antlr4} -Dlanguage=Go $1.g4
${antlr4} -Dlanguage=Go -visitor $1.g4