package main

import "log"

type IgoChessman struct {
	Color string
}

func (ic IgoChessman) Display(coordinate Coordinate) {
	log.Printf("棋子颜色：%v 坐标：(%v,%v)\n", ic.Color, coordinate.X, coordinate.Y)
}

func NewIgoChessman(t string) *IgoChessman {
	switch t {
	case "white":
		return &IgoChessman{
			Color: "白色",
		}
	case "black":
		return &IgoChessman{
			Color: "黑色",
		}
	default:
		return nil
	}
}

type IgoChessmanFactory map[string]*IgoChessman

func (icf IgoChessmanFactory) GetChessman(t string) *IgoChessman {
	igoChessman := icf[t]
	if igoChessman == nil {
		igoChessman = NewIgoChessman(t)
		icf[t] = igoChessman
	}
	return igoChessman
}

var Maps IgoChessmanFactory

func NewIgoChessmanFactory() {
	if Maps == nil {
		Maps = make(map[string]*IgoChessman)
	}
}

type Coordinate struct {
	X int
	Y int
}

func main() {
	NewIgoChessmanFactory()
	tCoordinate := Coordinate{
		X: 1,
		Y: 2,
	}
	tIgoChessman := Maps.GetChessman("white")
	tIgoChessman.Display(tCoordinate)

	tIgoChessman = Maps.GetChessman("black")
	tIgoChessman.Display(tCoordinate)
	return
}
