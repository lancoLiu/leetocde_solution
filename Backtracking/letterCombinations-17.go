package Backtracking

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func LetterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}
	res = []string{}
	backTrackLetterCombinations(digits, 0, "")
	return res

}

//回溯递归
func backTrackLetterCombinations(digits string, index int, tmpStr string) {

	//递归终止条件，存储结果
	if index == len(digits) {
		res = append(res, tmpStr)
		return
	}
	//递归逻辑
	digit := string(digits[index])
	letters := phoneMap[digit]
	letterCount := len(letters)
	for i := 0; i < letterCount; i++ {
		backTrackLetterCombinations(digits, index+1, tmpStr+string(letters[i]))

	}
}
