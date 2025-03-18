package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// Init 初始化系统时间起点和机器节点。
// 该函数接收开始时间字符串和机器ID，解析开始时间，计算时间戳，并初始化机器节点,用于生成雪花id。
// 参数:
//
//	startTime - 表示时间起点的字符串，格式为"2006-01-02"。
//	machineID - 机器ID，用于标识不同的机器。
//
// 返回值:
//
//	可能发生的错误。
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GenID 生成雪花ID。
// 该函数返回一个64位整数，表示雪花ID。
// 返回值:
//
//	雪花ID。
func GenID() int64 {
	return node.Generate().Int64()
}
