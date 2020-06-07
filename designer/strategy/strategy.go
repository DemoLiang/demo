package main

import "log"

type Staplefood interface {
	Eat()
}

type RiceStaplefood struct {
}

type NoodleStaplefood struct {
}

type BreadStaplefood struct {
}
type EatContext struct {
	staplefood Staplefood
}

func (RiceStaplefood)Eat(){
	log.Printf("吃米饭\n")
}

func (NoodleStaplefood)Eat(){
	log.Printf("吃面条\n")
}

func (BreadStaplefood)Eat(){
	log.Printf("吃面包\n")
}

func (context EatContext)Eatfood(){
	context.staplefood.Eat()
}

func NewEatContext(staplefood Staplefood)*EatContext{
	return &EatContext{
		staplefood:staplefood,
	}
}

func main(){
	eat:=NewEatContext(new(RiceStaplefood))
	eat.Eatfood()
	eat = NewEatContext(new(NoodleStaplefood))
	eat.Eatfood()
	eat = NewEatContext(new(BreadStaplefood))
	eat.Eatfood()
}