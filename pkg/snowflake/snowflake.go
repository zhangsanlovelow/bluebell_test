package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// 指定开始时间和机器ID初始化snowflake
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	//解析开始时间
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	//这个纪元可以自己设置，是一个int64的毫秒级别的时间戳
	//纳秒除以1000000，转换成秒数
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GenID 生成64位唯一分布式ID
func GenID() int64 {
	return node.Generate().Int64()
}
