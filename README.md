Battleship
==========

##Protocol
Each message between the server and client consists of a 5byte header followed by the payload, a UTF-8 encoded string with json.

The Header is divided as follows:

              4bytes          |        1byte
------------------------------|--------------------
payload size in bytes (UInt32)| payload type (UInt8)

NOTE:
The headers are unsigned ints in big endian.


Payload Types are:

###Common Message Types
UInt8 | Type                        | Payload Format Example
------|-----------------------------|----------------
0     | Ping                        | None
1     | Ok                          | `{ "ok": "" }`
2     | Error                       | `{ "error": "" }`
3     | GameMove                    | `{ "player": 0, "x": 1, "y": 2 }`
4     | ChatMessage                 | `{ "msg": "" }`

###Client Message Types
UInt8 | Type                        | Payload Format Example
------|-----------------------------|----------------
5     | Connect                     | `{ "username": "" }`
6     | RequestOpenGamesList        | None
7     | CreateGame                  | None
8     | JoinGame                    | `{ "id": 100 }`
9     | AcceptGame                  | `{ "id": 100 }`
10    | RejectGame                  | `{ "id": 100 }`
11    | GameSetPiece                | `{ "piece": 0, "start": {"x": 0, "y": 1}, "end": "start": {"x": 0, "y": 1} }`
12    | RequestGameState            | None
13    | AbandonGame                 | None

###Server Message Types
UInt8 | Type                        | Payload Format Example
------|-----------------------------|----------------
14    | OpenGamesList               | `{"games": [{"id": 10, "username": "jonfk"}, {"id": 10, "username": "jonfk"}]}`
15    | GamePreGameStatus           | `{ "id": 10, "opponent": "jonfk" }`
16    | GameState                   | `{"p1":"jonfk","p2":"Gery","you":[[1,2,3,4,5,6,7,8,9,10],[1,2,3,4,5,6,7,8,9,10]],"opponnent":[[1,2,3,4,5,6,7,8,9,10],[1,2,3,4,5,6,7,8,9,10]]}`
17    | GameWon                     | None
18    | GameLost                    | None

####Note:
When there is no payload for a message, the payload length should be 0.


##Server dependencies
```bash
$ go get -u github.com/boltdb/bolt/...
```