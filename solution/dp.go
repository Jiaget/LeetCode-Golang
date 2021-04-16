package solution

// 扰乱字符串 （递归） https://leetcode-cn.com/problems/scramble-string/
// 超时，不推荐
func IsScramble_recursion(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	mapp := make(map[int8]int8)
	// 判断两个字符串字符出现次数是否相同
	for i := range s1 {
		mapp[int8(s1[i]-'a')]++
		mapp[int8(s2[i]-'a')]--
	}
	for _, e := range mapp {
		if e != 0 {
			return false
		}
	}
	n := len(s1)
	// 递归，只要所有分支有一个true，就是扰乱字符串
	for k := 1; k < n; k++ {
		// 分割，将分割下来的左右两个字符串分别进入递归
		// 未交换的情况
		if IsScramble_recursion(s1[:k], s2[:k]) && IsScramble_recursion(s1[k:], s2[k:]) {
			return true
		}
		// 交换的情况
		if IsScramble_recursion(s1[:k], s2[n-k:]) && IsScramble_recursion(s1[k:], s2[:n-k]) {
			return true
		}
	}
	return false
}

// 扰乱字符串 （动规） https://leetcode-cn.com/problems/scramble-string/
// 区间递归通常使用dp[i][j] ,i 代表区间起点，j代表长度
// 本题有两个字符串，即两个区间，本应该使用四维数据dp[i][j][k][l]。不过可以进行压缩。因为两个字符串长度一样，需要的区间长度也相同。压缩成dp[i][j][len]
// 由于字符串交换位置与否是随机的，因此我们需要分成 1. s1部分交换后能否变成s2的部分以及；2.s1不交换后能否变成s2的部分。 两种情况讨论。
// 两个字符串如果字符个数出现不相同的情况，可以直接判断false。可以使用map来实现。
func IsScramble_dp(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	mapp := make(map[int8]int8)
	// 判断两个字符串字符出现次数是否相同
	for i := range s1 {
		mapp[int8(s1[i]-'a')]++
		mapp[int8(s2[i]-'a')]--
	}
	for _, e := range mapp {
		if e != 0 {
			return false
		}
	}
	// 初始化dp数组
	n := len(s1)
	dp := make([][][]bool, n)
	for i := range s1 {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}
	// i, j , k 分别代表，s1的起点，s2的起点，区间长度。其中k 作为区间长度必然 大于等于1
	for k := 1; k <= n; k++ {
		// s1 的区间
		for i := 0; i < n-k+1; i++ {
			// s2 的区间
			for j := 0; j < n-k+1; j++ {
				// 两个区间里都只有一个字符，直接比较就可以了
				if k == 1 {
					if s1[i] == s2[j] {
						dp[i][j][k] = true
					}
				} else {
					// 由于扰乱字符串是将字符串的某两个不相交的字串交换顺序产生的新字符串
					// 所以，当子串进行错位对比，必然会会出现成功匹配的情况。
					// 可以将匹配的情况分为
					// 1.dp[i][j][u] && dp[i + u][j + u][k - u] 子串没有交换的情况
					// 2.dp[i][j+k-u][u] && dp[i+u][j][k-u] 子串发生交换的情况
					// 这两种情况，只要任一匹配成功，即可视为发生过一次扰乱字符串。
					for u := 1; u < k; u++ {
						if dp[i][j][u] && dp[i+u][j+u][k-u] ||
							dp[i][j+k-u][u] && dp[i+u][j][k-u] {
							dp[i][j][k] = true
							break
						}
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
