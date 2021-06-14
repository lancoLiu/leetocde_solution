package main

//给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
func nextGreatestLetter(letters []byte, target byte) byte {
 l, h := 0, len(letters)-1
    for l <= h {
        m:= l + (h-l)/2
        if letters[m] <= target {
            l = m + 1
        } else {
            h = m -1
        }
    }
    if l < len(letters) {
        return letters[l]
    }
    return letters[0]

	//找最右侧边界
	l,r := 0,len(letters) - 1

	for l <= r {
		mid := (r - l ) / 2 +  l 

		if letters[mid] <= target  {
			l = mid+1
		}else{
			r = mid - 1
		}
	
	}
		//[r,l]
	 if l < len(letters) {
        return letters[l]
    }
    return letters[0]
}

func main()  {
	
}

