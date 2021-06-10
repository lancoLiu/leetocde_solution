package main

//颠倒给定的 32 位无符号整数的二进制位。
//输入: 00000010100101000001111010011100
//输出: 00111001011110000010100101000000
func reverseBits(num uint32) (rev uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		//tmp = num & 1
		//tmp2 = tmp << (31 - i)
		//rev |= tmp2
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return
}

func main() {

}
