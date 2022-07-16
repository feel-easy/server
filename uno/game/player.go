package game

import (
	"github.com/ratel-online/server/uno/card"
	"github.com/ratel-online/server/uno/card/color"
)

type Player interface {
	PlayerID() int64
	NickName() string
	PickColor(gameState State) color.Color
	Play(playableCards []card.Card, gameState State) (card.Card, error)
	NotifyCardsDrawn(drawnCards []card.Card)
	NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card)
}
