package game

import (
	"fmt"
	"testing"
)

func TestSetPiece(t *testing.T) {
	game := NewGame(10, 10)
	err := game.SetPiece(Player1, Coord{X: 0, Y: 0}, Coord{X: 0, Y: 2}, Submarine)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Game: %v\n", game)
}

func TestSetPiece2(t *testing.T) {
	game := NewGame(10, 10)
	err := game.SetPiece(Player1, Coord{X: 9, Y: 7}, Coord{X: 9, Y: 9}, Battleship)
	if err == nil {
		t.Error("Should be invalid")
	}
}

func TestIsReadyToStart(t *testing.T) {
	game := NewGame(10, 10)
	if game.IsReadyToStart() {
		t.Error("Game should not be ready to start")
	}
	game.SetPlayer(Player1, "jonfk")
	game.SetPlayer(Player2, "gery")
	pieces := []Piece{
		Piece{Type: PatrolBoat, Start: Coord{0, 0}, End: Coord{0, 1}},
		Piece{Type: Destroyer, Start: Coord{0, 2}, End: Coord{0, 4}},
		Piece{Type: Submarine, Start: Coord{0, 5}, End: Coord{0, 7}},
		Piece{Type: Battleship, Start: Coord{1, 0}, End: Coord{1, 3}},
		Piece{Type: AircraftCarrier, Start: Coord{2, 0}, End: Coord{2, 4}},
	}
	for _, piece := range pieces {
		err := game.SetPiece(Player1, piece.Start, piece.End, piece.Type)
		if err != nil {
			t.Error(err)
		}
	}
	for _, piece := range pieces {
		err := game.SetPiece(Player2, piece.Start, piece.End, piece.Type)
		if err != nil {
			t.Error(err)
		}
	}
	if !game.IsReadyToStart() {
		t.Error("Game should be ready to start")
	}
	fmt.Printf("Game: %v\n", game)
}

func TestMove(t *testing.T) {
	game := NewGame(10, 10)
	game.SetPlayer(Player1, "jonfk")
	game.SetPlayer(Player2, "gery")
	pieces := []Piece{
		Piece{Type: PatrolBoat, Start: Coord{0, 0}, End: Coord{0, 1}},
		Piece{Type: Destroyer, Start: Coord{0, 2}, End: Coord{0, 4}},
		Piece{Type: Submarine, Start: Coord{0, 5}, End: Coord{0, 7}},
		Piece{Type: Battleship, Start: Coord{1, 0}, End: Coord{1, 3}},
		Piece{Type: AircraftCarrier, Start: Coord{2, 0}, End: Coord{2, 4}},
	}
	for _, piece := range pieces {
		err := game.SetPiece(Player1, piece.Start, piece.End, piece.Type)
		if err != nil {
			t.Error(err)
		}
	}
	for _, piece := range pieces {
		err := game.SetPiece(Player2, piece.Start, piece.End, piece.Type)
		if err != nil {
			t.Error(err)
		}
	}
	var err error
	// Turn 1
	p1T1 := Coord{X: 0, Y: 1}
	err = game.Move(Player1, p1T1)
	if err != nil {
		t.Error(err)
	}
	p2T1 := Coord{X: 0, Y: 2}
	err = game.Move(Player2, p2T1)
	if err != nil {
		t.Error(err)
	}
	// Turn 2
	p1T2 := Coord{X: 9, Y: 0}
	err = game.Move(Player1, p1T2)
	if err != nil {
		t.Error(err)
	}
	err = game.Move(Player2, p2T1)
	if err == nil {
		//Expect error
		t.Error("Expected Error on repeated move %v", p2T1)
	}
	p2T2 := Coord{X: 9, Y: 9}
	err = game.Move(Player2, p2T2)
	if err != nil {
		t.Error(err)
	}
	// Turn 3
	p1T3 := Coord{X: 10, Y: 1}
	err = game.Move(Player1, p1T3)
	if err == nil {
		t.Error("Expected Error on invalid move out of grid")
	}

	if game.Player2Grid[1][0] != HitGrid {
		t.Errorf("Coord %v should be %v instead it is %v", p1T1, HitGrid, game.Player1Grid[1][0])
	}
	if game.Player1Grid[2][0] != HitGrid {
		t.Errorf("Coord %v should be %v instead it is %v", p2T1, HitGrid, game.Player1Grid[2][0])
	}
	if game.Player2Grid[0][9] != EmptyHitGrid {
		t.Errorf("Coord %v should be %v instead it is %v", p1T2, HitGrid, game.Player1Grid[0][9])
	}
	if game.Player1Grid[9][9] != EmptyHitGrid {
		t.Errorf("Coord %v should be %v instead it is %v", p2T2, HitGrid, game.Player1Grid[9][9])
	}
}

func TestGameWon(t *testing.T) {
	game := NewGame(10, 10)
	game.SetPlayer(Player1, "jonfk")
	game.SetPlayer(Player2, "gery")
	piece := Piece{Type: PatrolBoat, Start: Coord{0, 0}, End: Coord{0, 1}}
	err := game.SetPiece(Player1, piece.Start, piece.End, piece.Type)
	if err != nil {
		t.Error(err)
	}
	err = game.SetPiece(Player2, piece.Start, piece.End, piece.Type)
	if err != nil {
		t.Error(err)
	}

	//Turn1
	err = game.Move(Player1, Coord{X: 9, Y: 9})
	if err != nil {
		t.Error(err)
	}
	err = game.Move(Player2, Coord{X: 0, Y: 0})
	if err != nil {
		t.Error(err)
	}
	// Turn2
	err = game.Move(Player1, Coord{X: 8, Y: 9})
	if err != nil {
		t.Error(err)
	}
	err = game.Move(Player2, Coord{X: 0, Y: 1})
	if err != nil {
		t.Error(err)
	}
	if game.HasPlayerWon(Player1) {
		t.Error("Player1 should not have won")
	}
	if !game.HasPlayerWon(Player2) {
		t.Error("Player2 should have won but hasn't")
	}

}
