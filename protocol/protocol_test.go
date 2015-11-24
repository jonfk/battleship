package protocol

import (
	"testing"
)

func TestProtocolMessages(t *testing.T) {
	messages := []BattleMsg{
		// Common Messages
		PingMsg{},
		OkMsg{Ok: "hello world"},
		ErrorMsg{Error: "This is not an error"},
		GameMoveMsg{Player: 0, X: 1, Y: 2},
		ChatMessageMsg{Msg: "This is not a message"},
		// Client Messages
		ConnectMsg{Username: "jonfk"},
		RequestOpenGamesListMsg{},
		CreateGameMsg{},
		JoinGameMsg{Id: 99},
		AcceptGameMsg{Id: 99},
		RejectGameMsg{Id: 99},
		GameSetPieceMsg{Piece: 2, Start: Coord{X: 0, Y: 0}, End: Coord{X: 99, Y: 100}},
		RequestGameStateMsg{},
		AbandonGameMsg{},
		// Server Messages
		OpenGamesListMsg{Games: []Game{Game{Id: 9919, Username: "gery"}, Game{Id: 91823, Username: "dad"}}},
		GamePreGameStatusMsg{Id: 2838, Opponent: ""},
		GameStateMsg{P1: "jonfk!", P2: "-Gery",
			YourGrid:     [][]int{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			OpponentGrid: [][]int{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
		GameWonMsg{},
		GameLostMsg{},
	}

	for _, msg := range messages {
		msgTypeB, msgB, err := Msg2Raw(msg)
		if err != nil {
			t.Error(err)
		}
		//t.Logf("\n%s\n\n", string(msgB))
		parsedMsg, err := Raw2Msg(msgTypeB, msgB)
		if err != nil {
			t.Error(err)
		}
		if !BattleMsgEquals(parsedMsg, msg) {
			t.Errorf("parsed message %#v should be the same as original %#v", parsedMsg, msg)
		}
	}
}
