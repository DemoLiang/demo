package main

import "fmt"

type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (this *IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"indian", "black", "hindi"}
}

func (this *KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"korean", "brown", "korean"}
}

func getGirlFriend(typeGf string) GirlFriend {
	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = &IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = &KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}

func main() {
	a := getGirlFriend("Indian")
	fmt.Print(a.eyesColor)
}
