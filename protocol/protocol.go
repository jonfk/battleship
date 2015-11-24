package protocol

import (
	"encoding/json"
	"fmt"
	"net"
)

func Raw2Msg(msgType uint8, msg []byte) (BattleMsg, error) {
	switch MsgType(msgType) {
	// Common Messages
	case Ping:
		return PingMsg{}, nil
	case Ok:
		var structMsg OkMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case Error:
		var structMsg ErrorMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case GameMove:
		var structMsg GameMoveMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case ChatMessage:
		var structMsg ChatMessageMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil

	// Client Messages
	case Connect:
		var connectMsg ConnectMsg
		err := json.Unmarshal(msg, &connectMsg)
		if err != nil {
			goto Error
		}
		return connectMsg, nil
	case RequestOpenGamesList:
		return RequestOpenGamesListMsg{}, nil
	case CreateGame:
		return CreateGameMsg{}, nil
	case JoinGame:
		var structMsg JoinGameMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case AcceptGame:
		var structMsg AcceptGameMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case RejectGame:
		var structMsg RejectGameMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case GameSetPiece:
		var structMsg GameSetPieceMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case RequestGameState:
		return RequestGameStateMsg{}, nil
	case AbandonGame:
		return AbandonGameMsg{}, nil
	// Server Messages
	case OpenGamesList:
		var structMsg OpenGamesListMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case GamePreGameStatus:
		var structMsg GamePreGameStatusMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case GameState:
		var structMsg GameStateMsg
		err := json.Unmarshal(msg, &structMsg)
		if err != nil {
			goto Error
		}
		return structMsg, nil
	case GameWon:
		return GameWonMsg{}, nil
	case GameLost:
		return GameLostMsg{}, nil
	default:
		return nil, fmt.Errorf("Unknown msg type %v", MsgType(msgType))
	}
Error:
	return nil, fmt.Errorf("Cannot Unmarshal message. Expected %v msg but Received %v", MsgType(msgType), string(msg))
}

func ReadMsg(conn net.Conn) (BattleMsg, error) {
	msgType, msg, err := readRawMsg(conn)
	if err != nil {
		return nil, err
	}
	return Raw2Msg(msgType, msg)
}

func Msg2Raw(message BattleMsg) (uint8, []byte, error) {
	switch message.(type) {
	// Common Messages
	case PingMsg:
		return uint8(Ping), []byte{}, nil
	case OkMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(Ok), byteMsg, nil
	case ErrorMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(Error), byteMsg, nil
	case GameMoveMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(GameMove), byteMsg, nil
	case ChatMessageMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(ChatMessage), byteMsg, nil

	// Client Messages
	case ConnectMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(Connect), byteMsg, nil
	case RequestOpenGamesListMsg:
		return uint8(RequestOpenGamesList), []byte{}, nil
	case CreateGameMsg:
		return uint8(CreateGame), []byte{}, nil
	case JoinGameMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(JoinGame), byteMsg, nil
	case AcceptGameMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(AcceptGame), byteMsg, nil
	case RejectGameMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(RejectGame), byteMsg, nil
	case GameSetPieceMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(GameSetPiece), byteMsg, nil
	case RequestGameStateMsg:
		return uint8(RequestGameState), []byte{}, nil
	case AbandonGameMsg:
		return uint8(AbandonGame), []byte{}, nil

	// Server Messages
	case OpenGamesListMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(OpenGamesList), byteMsg, nil
	case GamePreGameStatusMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(GamePreGameStatus), byteMsg, nil
	case GameStateMsg:
		byteMsg, err := json.Marshal(message)
		if err != nil {
			return 0, nil, err
		}
		return uint8(GameState), byteMsg, nil
	case GameWonMsg:
		return uint8(GameWon), []byte{}, nil
	case GameLostMsg:
		return uint8(GameLost), []byte{}, nil
	default:
		return 0, nil, fmt.Errorf("Unknown msg type %v cannot be sent", message)
	}
}

func WriteMsg(conn net.Conn, message BattleMsg) error {
	msgType, payload, err := Msg2Raw(message)
	if err != nil {
		return err
	}
	return writeRawMsg(conn, msgType, payload)
}
