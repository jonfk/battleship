// generated by stringer -type=PieceType; DO NOT EDIT

package game

import "fmt"

const _PieceType_name = "PatrolBoatDestroyerSubmarineBattleshipAircraftCarrier"

var _PieceType_index = [...]uint8{0, 10, 19, 28, 38, 53}

func (i PieceType) String() string {
	if i < 0 || i >= PieceType(len(_PieceType_index)-1) {
		return fmt.Sprintf("PieceType(%d)", i)
	}
	return _PieceType_name[_PieceType_index[i]:_PieceType_index[i+1]]
}
