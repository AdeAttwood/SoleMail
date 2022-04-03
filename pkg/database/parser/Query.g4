grammar Query;

query: expression;

expression
 : expressionItem                #boolExpression
 | expression join expression    #joinExpression
 | expression expression         #joinExpression
 | '(' expression ')'            #embedExpression
 ;

expressionItem
 : Word Colon operator? matchValue
 ;

matchValue
 : Word
 | QuotedString
 ;

join
 : 'and'
 | 'or'
 ;

operator
 : NotEqual
 ;

NotEqual: '!';

Colon:  ':';
Word:   ([a-zA-Z])+;
QuotedString: '"' (~[\\"] | '\\' [\\"])* '"';

SPACE: [ \t\r\n]+ -> skip;