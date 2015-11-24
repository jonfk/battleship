//go:generate stringer -type=PieceType
//go:generate stringer -type=Player
package game

import (
	"bytes"
	"fmt"
)

type Game struct {
	Size         Coord
	Player1Grid  [][]GridState
	Player2Grid  [][]GridState
	Player1      *string
	Player2      *string
	Player1Ships []Piece
	Player2Ships []Piece
	CurrentTurn  Player
}

type Coord struct {
	X int
	Y int
}

type PieceType int

const (
	PatrolBoat PieceType = iota
	Destroyer
	Submarine
	Battleship
	AircraftCarrier
)

func (piece PieceType) Length() int {
	switch piece {
	case PatrolBoat:
		return 2
	case Destroyer:
		return 3
	case Submarine:
		return 3
	case Battleship:
		return 4
	case AircraftCarrier:
		return 5
	default:
		return -1
	}
}

type Piece struct {
	Type  PieceType
	Start Coord
	End   Coord
}

func (piece Piece) Length() int {
	return piece.Type.Length()
}

type Player int

const (
	Player1 Player = iota
	Player2
)

func (player Player) IsValid() bool {
	if player < Player1 || player > Player2 {
		return false
	}
	return true
}

type GridState int

const (
	EmptyGrid GridState = iota
	ShipGrid
	HitGrid
	EmptyHitGrid
)

func NewGame(x, y int) *Game {
	newGame := &Game{Size: Coord{X: x, Y: y}}
	newGame.Player1Grid = make([][]GridState, y)
	newGame.Player2Grid = make([][]GridState, y)

	for i := 0; i < y; i++ {
		newGame.Player1Grid[i] = make([]GridState, x)
		newGame.Player2Grid[i] = make([]GridState, x)
	}
	return newGame
}

func (game *Game) SetPiece(player Player, start, end Coord, piece PieceType) error {
	if piece.Length() < 0 {
		return fmt.Errorf("SetPiece: piece %v is invalid", piece.String())
	}
	if !game.IsValidCoord(start) {
		return fmt.Errorf("SetPiece: start coordinate %#v is invalid", start)
	}
	if !game.IsValidCoord(end) {
		return fmt.Errorf("SetPiece: end coordinate %#v is invalid", end)
	}
	if !player.IsValid() {
		return fmt.Errorf("SetPiece: player %v invalid", player)
	}
	pieceLength := piece.Length()
	var grid [][]GridState
	//var pieceList []Piece
	switch player {
	case Player1:
		grid = game.Player1Grid
		//pieceList = game.Player1Ships
	case Player2:
		grid = game.Player2Grid
		//pieceList = game.Player2Ships
	}

	if start.X == end.X && abs(start.Y-end.Y) == (pieceLength-1) {
		minY := min(start.Y, end.Y)
		if minY == end.Y {
			start, end = end, start
		}
		// check grid for obstructing piece
		for i := start.Y; i <= end.Y; i++ {
			if grid[i][start.X] != EmptyGrid {
				return fmt.Errorf("SetPiece: piece already at %v obstructs start %v and end %v locations for piece %v", Coord{X: start.X, Y: i}, start, end, piece)
			}
		}

		for i := start.Y; i <= end.Y; i++ {
			grid[i][start.X] = ShipGrid
		}
	} else if start.Y == end.Y && abs(start.X-end.X) == (pieceLength-1) {
		minX := min(start.X, end.X)
		if minX == end.X {
			start, end = end, start
		}
		// check grid for obstructing piece
		for i := start.X; i <= end.X; i++ {
			if grid[start.Y][i] != EmptyGrid {
				return fmt.Errorf("SetPiece: piece already at %v obstructs start %v and end %v locations for piece %v", Coord{X: i, Y: start.Y}, start, end, piece)
			}
		}

		for i := start.X; i <= end.X; i++ {
			grid[start.Y][i] = ShipGrid
		}

	} else {
		return fmt.Errorf("SetPiece: invalid start (%v) and end(%v) locations for piece length %d", start, end, pieceLength)
	}
	switch player {
	case Player1:
		game.Player1Ships = append(game.Player1Ships, Piece{Type: piece, Start: start, End: end})
	case Player2:
		game.Player2Ships = append(game.Player2Ships, Piece{Type: piece, Start: start, End: end})
	}
	return nil
}

func (game *Game) Move(player Player, coord Coord) error {
	if game.CurrentTurn != player {
		return fmt.Errorf("Move: Cannot execute move %v for %v. Currently %v's turn", coord, player, game.CurrentTurn)
	}
	if !game.IsValidCoord(coord) {
		return fmt.Errorf("Move: Invalid move coordinate %v", coord)
	}
	var grid [][]GridState
	switch player {
	case Player1:
		grid = game.Player2Grid
	case Player2:
		grid = game.Player1Grid
	}
	if grid[coord.Y][coord.X] == EmptyGrid {
		grid[coord.Y][coord.X] = EmptyHitGrid
	} else if grid[coord.Y][coord.X] == ShipGrid {
		grid[coord.Y][coord.X] = HitGrid
	} else {
		return fmt.Errorf("Move: Invalid move %v has already been executed before", coord)
	}
	game.changeTurn()
	return nil

}

func (game *Game) IsReadyToStart() bool {
	var p1Ships map[PieceType]bool = make(map[PieceType]bool)
	var p2Ships map[PieceType]bool = make(map[PieceType]bool)
	for _, val := range game.Player1Ships {
		p1Ships[val.Type] = true
	}
	for _, val := range game.Player2Ships {
		p2Ships[val.Type] = true
	}
	return game.Player1 != nil && game.Player2 != nil && len(p1Ships) >= 5 && len(p2Ships) >= 5
}

func (game *Game) SetPlayer(player Player, name string) {
	var playerName *string = new(string)
	*playerName = name
	if player == Player1 {
		game.Player1 = playerName
	} else {
		game.Player2 = playerName
	}
}

func (game *Game) HasPlayerWon(player Player) bool {
	var grid [][]GridState
	switch player {
	case Player1:
		grid = game.Player2Grid
	case Player2:
		grid = game.Player1Grid
	default:
		return false
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == ShipGrid {
				return false
			}
		}
	}
	return true
}

func (game *Game) changeTurn() {
	if game.CurrentTurn == Player1 {
		game.CurrentTurn = Player2
	} else {
		game.CurrentTurn = Player1
	}
}

// Util functions

func (game *Game) IsValidCoord(coord Coord) bool {
	if coord.X < 0 || coord.X >= game.Size.X {
		return false
	}
	if coord.Y < 0 || coord.Y >= game.Size.Y {
		return false
	}
	return true
}

func (game *Game) String() string {
	buf := new(bytes.Buffer)

	var (
		player1Name string
		player2Name string
	)
	if game.Player1 != nil {
		player1Name = *game.Player1
	} else {
		player1Name = "<nil>"
	}
	if game.Player2 != nil {
		player2Name = *game.Player2
	} else {
		player2Name = "<nil>"
	}

	buf.WriteString(fmt.Sprintf("Size: x: %d, y: %d\n", game.Size.X, game.Size.Y))
	buf.WriteString(fmt.Sprintf("Players: 1: %v, 2: %v\n", player1Name, player2Name))
	buf.WriteString(fmt.Sprint("Players 1 Ships:\n"))
	for _, v := range game.Player1Ships {
		buf.WriteString(fmt.Sprintf("\t%#v\n", v))
	}
	buf.WriteString(fmt.Sprint("Players 2 Ships:\n"))
	for _, v := range game.Player2Ships {
		buf.WriteString(fmt.Sprintf("\t%#v\n", v))
	}

	buf.WriteString(fmt.Sprint("Players 1 Grid:\n"))
	for _, v := range game.Player1Grid {
		buf.WriteString(fmt.Sprintf("\t%#v\n", v))
	}
	buf.WriteString(fmt.Sprint("Players 2 Grid:\n"))
	for _, v := range game.Player2Grid {
		buf.WriteString(fmt.Sprintf("\t%#v\n", v))
	}
	return buf.String()

}

func (coord Coord) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", coord.X, coord.Y)
}

func AllPieceTypes() <-chan PieceType {
	// You can define constraints for the iterator in one place
	var first PieceType = PatrolBoat
	var last PieceType = AircraftCarrier

	// Sequential values of the iterator are communicated via channel
	ch := make(chan PieceType)

	// Spawn a goroutine to iterate over the values of the iota constant
	go func() {
		for piece := first; piece <= last; piece++ {
			ch <- piece
		}

		// Indicate to consumers the iteration of the enumeration is complete
		close(ch)
	}()

	return ch
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
