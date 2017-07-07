//中央服务器

package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"ipc"
)

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}
