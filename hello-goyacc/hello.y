%{
package main
import "fmt"
%}

%union{
type_string string
}

%type <type_string>expression

%token <type_string> PRINT HELLO

%%
expression:HELLO
{
fmt.Println("Hello World")
}
%%