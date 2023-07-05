package logic

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

// 生成唯一ID
func GenerateID() uint64 {
	return uint64(node.Generate().Int64())
}

func GenerateID32() uint {
	var id = uint64(node.Generate().Int64())
	id32 := uint(id & 7)
	id >>= 9
	id32 += uint(id & 7 << 3)
	id >>= 3
	id32 += uint(id & 7 << 6)
	id >>= 8
	id32 += uint(id & 3 << 9)
	id >>= 2
	id32 += uint(id & 2047 << 11)
	id >>= 31
	id32 += uint(id & 1023 << 22)
	return id32
}
