package coin_change

func GetWays(money int32, coins [] int64) int64 {
	dp := make([]int, money+1)
	dp[0] = 1
	for i:=0; i<len(coins); i++{
		start := coins[i]
		for j := start; j<=int64(money); j++{
			dp[j] += dp[j-start]
		}
	}
	return int64(dp[money])
}
