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
