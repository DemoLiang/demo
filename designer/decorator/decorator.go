package main

import "log"

type Component interface {
	Display()
}

type Window struct {
}

func (w Window) Display() {
	log.Printf("show Windows\n")
	return
}

type TextBox struct {
}

func (t TextBox) Display() {
	log.Printf("show TextBox\n")
	return
}

type ListBox struct {
}

func (l ListBox) Display() {
	log.Printf("show list box\n")
	return
}

type ScrollBarDecorator struct {
	Component
}

func (sbd ScrollBarDecorator) Display() {
	log.Printf("scrollbar display\n")
	sbd.Component.Display()
	return
}

type BlackBoarderDecorator struct {
	Component
}

func (bbd BlackBoarderDecorator) Display() {
	log.Printf("black border\n")
	bbd.Component.Display()
	return
}

func NewDecorator(t string, decorator Component) Component {
	switch t {
	case "sbd":
		return ScrollBarDecorator{
			Component: decorator,
		}
	case "bbd":
		return BlackBoarderDecorator{
			Component: decorator,
		}
	default:
		return nil
	}
}

func main() {
	component := Window{}
	tScrollBarDecorator := NewDecorator("sbd", component)
	tScrollBarDecorator.Display()

	tBlackBorderDecorator := NewDecorator("bbd", tScrollBarDecorator)
	tBlackBorderDecorator.Display()
	return
}
