//go:generate stringer -type=MsgType
package protocol

import ()

type MsgType uint8

const (
	// Common Messages
	Ping MsgType = iota
	Ok
	Error
	GameMove
	ChatMessage
	//Client Messages
	Connect
	RequestOpenGamesList
	CreateGame
	JoinGame
	AcceptGame
	RejectGame
	GameSetPiece
	RequestGameState
	AbandonGame
	//Server Messages
	OpenGamesList
	GamePreGameStatus
	GameState
	GameWon
	GameLost
)

func AllMsgTypes() <-chan MsgType {
	// You can define constraints for the iterator in one place
	var first MsgType = Ping
	var last MsgType = GameLost

	// Sequential values of the iterator are communicated via channel
	ch := make(chan MsgType)

	// Spawn a goroutine to iterate over the values of the iota constant
	go func() {
		for msg := first; msg <= last; msg++ {
			ch <- msg
		}

		// Indicate to consumers the iteration of the enumeration is complete
		close(ch)
	}()

	return ch
}

type BattleMsg interface {
	BattleMsg()
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

/*
 * Common Messages
 */

type PingMsg struct{}
type OkMsg struct {
	Ok string `json:"ok,omitempty"`
}
type ErrorMsg struct {
	Error string `json:"error:omitempty"`
}
type GameMoveMsg struct {
	Player int `json:"player"`
	X      int `json:"x"`
	Y      int `json:"y"`
}
type ChatMessageMsg struct {
	Msg string `json:"msg"`
}

/*
 * Client Messages
 */

type ConnectMsg struct {
	Username string `json:"username"`
}
type RequestOpenGamesListMsg struct{}
type CreateGameMsg struct{}
type JoinGameMsg struct {
	Id int `json:"id"`
}
type AcceptGameMsg struct {
	Id int `json:"id"`
}
type RejectGameMsg struct {
	Id int `json:"id"`
}
type GameSetPieceMsg struct {
	Piece int   `json:"piece"`
	Start Coord `json:"start"`
	End   Coord `json:"end"`
}
type RequestGameStateMsg struct{}
type AbandonGameMsg struct{}

/*
 * Server Messages
 */
type Game struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type OpenGamesListMsg struct {
	Games []Game `json:"games"`
}
type GamePreGameStatusMsg struct {
	Id       int    `json:"id"`
	Opponent string `json:"opponent"`
}
type GameStateMsg struct {
	P1           string  `json:"p1"`
	P2           string  `json:"p2"`
	YourGrid     [][]int `json:"you"`
	OpponentGrid [][]int `json:"opponnent"`
}
type GameWonMsg struct{}
type GameLostMsg struct{}

/*
 * Methods to satisfy BattleShipMsg interface
 */

// Common
func (m PingMsg) BattleMsg()        {}
func (m OkMsg) BattleMsg()          {}
func (m ErrorMsg) BattleMsg()       {}
func (m GameMoveMsg) BattleMsg()    {}
func (m ChatMessageMsg) BattleMsg() {}

// Client
func (m ConnectMsg) BattleMsg()              {}
func (m RequestOpenGamesListMsg) BattleMsg() {}
func (m CreateGameMsg) BattleMsg()           {}
func (m JoinGameMsg) BattleMsg()             {}
func (m AcceptGameMsg) BattleMsg()           {}
func (m RejectGameMsg) BattleMsg()           {}
func (m GameSetPieceMsg) BattleMsg()         {}
func (m RequestGameStateMsg) BattleMsg()     {}
func (m AbandonGameMsg) BattleMsg()          {}

// Server
func (m OpenGamesListMsg) BattleMsg()     {}
func (m GamePreGameStatusMsg) BattleMsg() {}
func (m GameStateMsg) BattleMsg()         {}
func (m GameWonMsg) BattleMsg()           {}
func (m GameLostMsg) BattleMsg()          {}

// Deep comparison of BattleMsg
func BattleMsgEquals(a, b BattleMsg) bool {
	switch a := a.(type) {
	case PingMsg:
		if _, ok := b.(PingMsg); ok {
			return true
		}
	case OkMsg:
		if b, ok := b.(OkMsg); ok && b.Ok == a.Ok {
			return true
		}
	case ErrorMsg:
		if b, ok := b.(ErrorMsg); ok && b.Error == a.Error {
			return true
		}
	case GameMoveMsg:
		if b, ok := b.(GameMoveMsg); ok &&
			b.Player == a.Player &&
			b.X == a.X &&
			b.Y == a.Y {
			return true
		}
	case ChatMessageMsg:
		if b, ok := b.(ChatMessageMsg); ok && b.Msg == a.Msg {
			return true
		}
	case ConnectMsg:
		if b, ok := b.(ConnectMsg); ok && b.Username == a.Username {
			return true
		}
	case RequestOpenGamesListMsg:
		if _, ok := b.(RequestOpenGamesListMsg); ok {
			return true
		}
	case CreateGameMsg:
		if _, ok := b.(CreateGameMsg); ok {
			return true
		}
	case JoinGameMsg:
		if b, ok := b.(JoinGameMsg); ok && b.Id == a.Id {
			return true
		}
	case AcceptGameMsg:
		if b, ok := b.(AcceptGameMsg); ok && b.Id == a.Id {
			return true
		}
	case RejectGameMsg:
		if b, ok := b.(RejectGameMsg); ok && b.Id == a.Id {
			return true
		}
	case GameSetPieceMsg:
		if b, ok := b.(GameSetPieceMsg); ok && b.Piece == a.Piece &&
			b.Start.X == a.Start.X && b.Start.Y == a.Start.Y &&
			b.End.X == a.End.X && b.End.Y == a.End.Y {
			return true
		}
	case RequestGameStateMsg:
		if _, ok := b.(RequestGameStateMsg); ok {
			return true
		}
	case AbandonGameMsg:
		if _, ok := b.(AbandonGameMsg); ok {
			return true
		}
	case OpenGamesListMsg:
		if b, ok := b.(OpenGamesListMsg); ok && len(b.Games) == len(a.Games) {
			for i := range b.Games {
				if b.Games[i].Id != a.Games[i].Id || b.Games[i].Username != a.Games[i].Username {
					return false
				}
			}
			return true
		}
	case GamePreGameStatusMsg:
		if b, ok := b.(GamePreGameStatusMsg); ok && b.Id == a.Id && b.Opponent == a.Opponent {
			return true
		}
	case GameStateMsg:
		if b, ok := b.(GameStateMsg); ok && b.P1 == a.P1 && b.P2 == a.P2 &&
			len(b.OpponentGrid) == len(a.OpponentGrid) &&
			len(b.YourGrid) == len(a.YourGrid) {

			for i := range b.YourGrid {
				if len(b.YourGrid[i]) != len(a.YourGrid[i]) {
					return false
				}
				for j := range b.YourGrid[i] {
					if b.YourGrid[i][j] != a.YourGrid[i][j] {
						return false
					}
				}
			}

			for i := range b.OpponentGrid {
				if len(b.OpponentGrid[i]) != len(a.OpponentGrid[i]) {
					return false
				}
				for j := range b.OpponentGrid[i] {
					if b.OpponentGrid[i][j] != a.OpponentGrid[i][j] {
						return false
					}
				}
			}
			return true
		}
	case GameWonMsg:
		if _, ok := b.(GameWonMsg); ok {
			return true
		}
	case GameLostMsg:
		if _, ok := b.(GameLostMsg); ok {
			return true
		}

	default:
		return false
	}
	return false
}
