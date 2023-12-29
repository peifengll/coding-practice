package leetcode

/**
不同字符不能映射到同一个字符上
相同只能映射到同一个
字符可以映射到自己


遇到一个直接映射，
如果已经映射了但是不一样，返回false
下边的已经被映射了，返回false
*/

func isIsomorphic(s string, t string) bool {
	//	 这道题就是看长得像不像
	dict := make(map[byte]byte)
	dict2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		_, ok1 := dict[s[i]]
		_, ok2 := dict2[t[i]]
		if ok1 && ok2 {
			if dict[s[i]] != t[i] || dict2[t[i]] != s[i] {
				return false
			}
		} else if ok1 || ok2 {
			return false
		} else {
			dict[s[i]] = t[i]
			dict2[t[i]] = s[i]
		}
	}

	return true
}
