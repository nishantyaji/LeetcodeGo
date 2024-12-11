package main 

import(
  "fmt"
)

func maximumBeauty(nums []int, k int) int {
    limit := 100002
    var dp[100002] int
    for i := 0; i < limit; i++ {
        dp[i] = 0
    }
    for i := 0; i < len(nums); i++ {
        dp[max(0, nums[i] - k)] += 1
        dp[min(limit -1, nums[i] + k + 1)] -= 1
    }
    runSum := 0
    mx := 0
    for i := 0; i < limit; i++ {
        runSum += dp[i]
        mx = max(mx, runSum)
    }
    return mx
}

func main() {
  var arr = [4]int{4,6,1,2}
  fmt.Println(maximumBeauty(arr, 2))  // 3
}
