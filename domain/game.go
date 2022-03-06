package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Phase int

const (
	Preparation Phase = 0
	Teaming     Phase = 1
	Bidding     Phase = 2
	Playing     Phase = 3
	Counting    Phase = 4
	Pause       Phase = 5
)

type BidValue int

const (
	Eighty           BidValue = 80
	Ninety           BidValue = 90
	Hundred          BidValue = 100
	HundredAndTen    BidValue = 110
	HundredAndTwenty BidValue = 120
	HundredAndThirty BidValue = 130
	HundredAndFourty BidValue = 140
	HundredAndFifty  BidValue = 150
	Capot            BidValue = 160
)

type Strength int

const (
	Seven  Strength = 1
	Eight  Strength = 2
	Nine   Strength = 3
	Jack   Strength = 4
	Queen  Strength = 5
	King   Strength = 6
	Ten    Strength = 7
	As     Strength = 8
	TSeven Strength = 11
	TEight Strength = 12
	TQueen Strength = 13
	TKing  Strength = 14
	TTen   Strength = 15
	TAs    Strength = 16
	TNine  Strength = 17
	TJack  Strength = 18
)

type Color string

const (
	Club     Color = "club"
	Diamond  Color = "diamond"
	Heart    Color = "heart"
	Spade    Color = "spade"
	NoTrump  Color = "noTrump"
	AllTrump Color = "allTrump"
)

type card struct {
	color         Color
	strength      Strength
	TrumpStrength Strength
}

type cardID string

const (
	C_7  cardID = "7-club"
	C_8  cardID = "8-club"
	C_9  cardID = "9-club"
	C_10 cardID = "10-club"
	C_J  cardID = "jack-club"
	C_Q  cardID = "queen-club"
	C_K  cardID = "king-club"
	C_A  cardID = "as-club"
	D_7  cardID = "7-diamond"
	D_8  cardID = "8-diamond"
	D_9  cardID = "9-diamond"
	D_10 cardID = "10-diamond"
	D_J  cardID = "jack-diamond"
	D_Q  cardID = "queen-diamond"
	D_K  cardID = "king-diamond"
	D_A  cardID = "as-diamond"
	H_7  cardID = "7-heart"
	H_8  cardID = "8-heart"
	H_9  cardID = "9-heart"
	H_10 cardID = "10-heart"
	H_J  cardID = "jack-heart"
	H_Q  cardID = "queen-heart"
	H_K  cardID = "king-heart"
	H_A  cardID = "as-heart"
	S_7  cardID = "7-spade"
	S_8  cardID = "8-spade"
	S_9  cardID = "9-spade"
	S_10 cardID = "10-spade"
	S_J  cardID = "jack-spade"
	S_Q  cardID = "queen-spade"
	S_K  cardID = "king-spade"
	S_A  cardID = "as-spade"
)

var cards = map[cardID]card{
	C_7:  {Club, Seven, TSeven},
	C_8:  {Club, Eight, TEight},
	C_9:  {Club, Nine, TNine},
	C_10: {Club, Ten, TTen},
	C_J:  {Club, Jack, TJack},
	C_Q:  {Club, Queen, TQueen},
	C_K:  {Club, King, TKing},
	C_A:  {Club, As, TAs},
	D_7:  {Diamond, Seven, TSeven},
	D_8:  {Diamond, Eight, TEight},
	D_9:  {Diamond, Nine, TNine},
	D_10: {Diamond, Ten, TTen},
	D_J:  {Diamond, Jack, TJack},
	D_Q:  {Diamond, Queen, TQueen},
	D_K:  {Diamond, King, TKing},
	D_A:  {Diamond, As, TAs},
	H_7:  {Heart, Seven, TSeven},
	H_8:  {Heart, Eight, TEight},
	H_9:  {Heart, Nine, TNine},
	H_10: {Heart, Ten, TTen},
	H_J:  {Heart, Jack, TJack},
	H_Q:  {Heart, Queen, TQueen},
	H_K:  {Heart, King, TKing},
	H_A:  {Heart, As, TAs},
	S_7:  {Spade, Seven, TSeven},
	S_8:  {Spade, Eight, TEight},
	S_9:  {Spade, Nine, TNine},
	S_10: {Spade, Ten, TTen},
	S_J:  {Spade, Jack, TJack},
	S_Q:  {Spade, Queen, TQueen},
	S_K:  {Spade, King, TKing},
	S_A:  {Spade, As, TAs},
}

func newDeck() []cardID {
	deck := []cardID{C_7, C_8, C_9, C_10, C_J, C_Q, C_K, C_A, D_7, D_8, D_9, D_10, D_J, D_Q, D_K, D_A, H_7, H_8, H_9, H_10, H_J, H_Q, H_K, H_A, S_7, S_8, S_9, S_10, S_J, S_Q, S_K, S_A}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

type Bid struct {
	Player  string
	Color   Color
	Coinche int
	Pass    int
}

type play struct {
	playerName string
	card       cardID
}

type turn struct {
	plays  []play
	winner string
}

func (turn *turn) setWinner(trump Color) {
	var winner string
	var strongerValue Strength
	var firstCard cardID
	for _, play := range turn.plays {
		if firstCard == "" {
			firstCard = play.card
		}
		cardValue := getCardValue(play.card, trump, firstCard)
		if cardValue > strongerValue {
			strongerValue = cardValue
			winner = play.playerName
		}
	}

	turn.winner = winner
}

func getCardValue(card cardID, trump Color, firstCard cardID) Strength {
	color := cards[card].color
	colorAsked := cards[firstCard].color

	if trump == color || trump == AllTrump {
		return cards[card].TrumpStrength
	} else if color == colorAsked {
		return cards[card].strength
	} else {
		return 0
	}
}

type Game struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Players   map[string]Player
	Phase     Phase
	Bids      map[BidValue]Bid
	trump     Color
	deck      []cardID
	turns     []turn
}

type Player struct {
	Team         string
	Order        int
	InitialOrder int
	Hand         []cardID
}

func (player Player) CanPlay() bool {
	return player.Order == 0
}

func NewGame(name string) Game {
	return Game{
		Name:    name,
		Players: map[string]Player{},
		Phase:   Preparation,
		Bids:    map[BidValue]Bid{},
		deck:    newDeck(),
	}
}

func (game *Game) checkPlayerTurn(playerName string) error {
	if game.Players[playerName].Order != 1 {
		return errors.New(ErrNotYourTurn)
	}
	return nil
}

func (game *Game) checkTeamTurn(playerName string) error {
	order := game.Players[playerName].Order
	if order != 1 && order != 3 {
		return errors.New(ErrNotYourTurn)
	}
	return nil
}

func (game *Game) setFirstPlayer(playerName string) {
	for i := 0; i < len(game.Players); i++ {
		if game.Players[playerName].Order == 1 {
			return
		}
		game.rotateOrder()
	}
}

func (game *Game) rotateOrder() {
	for name, player := range game.Players {
		if player.Order == 1 {
			player.Order = 4
		} else {
			player.Order--
		}

		game.Players[name] = player
	}
}
