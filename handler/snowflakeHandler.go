package handler

import (
	"time"
)

// Solution copied from https://canary.discord.com/channels/81384788765712384/381889600978681856/385510659175088129 in the Discord API server
func CreationTime(snowflake int64) int64  {
	timestamp := (snowflake >> 22) + 1420070400000
	time := time.Unix(timestamp/1000, 0)
	return time.Unix()
	
}