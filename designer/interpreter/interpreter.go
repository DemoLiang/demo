package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type TOKEN_TYPE int

const(
	TYPE_NONE TOKEN_TYPE = iota
	TYPE_INTEGER
	TYPE_PLUS
	TYPE_MINUS
	TYPE_MUL
	TYPE_DIV
	TYPE_EOF
	)

type Token struct {
	TokenType TOKEN_TYPE
	TokenValue interface{}
}

type Lexer struct {
	Text []rune
	Pos int
	CurrentChar rune
}

type Interpreter struct {
	InterLexer *Lexer
	CurrentToken Token
}

func (lexer *Lexer)Advance(){
	lexer.Pos +=1
	if lexer.Pos>len(lexer.Text)-1{
		lexer.CurrentChar = 0
	}else{
		lexer.CurrentChar = lexer.Text[lexer.Pos]
	}
}

func StrToInt(s string)(int){
	if ret,err:=strconv.Atoi(s);err==nil{
		return ret
	}
	panic(fmt.Sprintf("cannot convert %s to int!",s))
}

func (lexer *Lexer)Integer()int{
	result:=""
	for lexer.CurrentChar!=0&&unicode.IsDigit(lexer.CurrentChar){
		result+=string(lexer.CurrentChar)
		lexer.Advance()
	}
	return StrToInt(result)
}

func (lexer *Lexer)SkipSpace(){
	for lexer.CurrentChar!=0&& unicode.IsSpace(lexer.CurrentChar){
		lexer.Advance()
	}
}

func (lexer *Lexer)GetNextToken()Token{
	for lexer.CurrentChar!=0{
		if unicode.IsSpace(lexer.CurrentChar){
			lexer.SkipSpace()
		}
		if unicode.IsDigit(lexer.CurrentChar){
			return Token{TYPE_INTEGER,lexer.Integer()}
		}
		if lexer.CurrentChar=='*'{
			lexer.Advance()
			return Token{TYPE_MUL,'*'}
		}
		if lexer.CurrentChar=='/'{
			lexer.Advance()
			return Token{TYPE_DIV,'/'}
		}
		panic("invalid text")
	}
	return Token{TYPE_EOF,nil}
}

func (interpreter *Interpreter)Eat(token_type TOKEN_TYPE){
	if interpreter.CurrentToken.TokenType == token_type{
		interpreter.CurrentToken = interpreter.InterLexer.GetNextToken()
	}else {
		panic("invalid token type")
	}
}

func (interpreter *Interpreter)Factor()int{
	token:=interpreter.CurrentToken
	interpreter.Eat(TYPE_INTEGER)
	return token.TokenValue.(int)
}

func (interpreter *Interpreter)Expr(){
	defer func() {
		if r:=recover();r!=nil{
			log.Printf("[Error]:%v\n",r)
		}
	}()
	interpreter.CurrentToken = interpreter.InterLexer.GetNextToken()

	result:=interpreter.Factor()

	for {
		if interpreter.CurrentToken.TokenType== TYPE_MUL{
			interpreter.Eat(TYPE_MUL)
			result*=interpreter.Factor()
		}else if interpreter.CurrentToken.TokenType==TYPE_DIV{
			interpreter.Eat(TYPE_DIV)
			result/=interpreter.Factor()
		}else{
			break
		}
	}
	log.Printf("> ",result)
}

func main(){
	log.Print("-----\n")
	reader:=bufio.NewReader(os.Stdin)
	for{
		log.Print("[calc]-> ")
		text,_:=reader.ReadString('\n')
		text=strings.ToLower(strings.TrimSpace(strings.TrimSuffix(text,"\n")))
		if len(text) == 0{
			continue
		}
		if text == "exit"{
			log.Printf("--------\n")
			log.Print("-----end----\n")
			os.Exit(0)
		}
		my_text:=[]rune(text)
		my_lexer:=Lexer{my_text,0,my_text[0]}
		my_Interpreter:=Interpreter{&my_lexer,Token{}}
		my_Interpreter.Expr()
	}
	return
}
