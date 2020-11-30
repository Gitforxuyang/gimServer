package utils

import "github.com/bwmarrin/snowflake"

var (
	node *snowflake.Node
)

func GetSnowflakeId() int64 {
	if node == nil {
		node, _ = snowflake.NewNode(1)
	}
	return node.Generate().Int64()
}
