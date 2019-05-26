package util

func BinarySearch(list []int,item int) int{
	low  := 0
	high := len(list)-1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		quess := list[mid]
		if quess == item {
			return quess
		}
		if quess > item {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return -1
}

func QSort(items []int) []int{
	if len(items) < 2 {
		return items
	} else {
		op := items[0]//опорный элемент
		var left,right []int
		for _ ,el := range items[1:] {
			if el < op {
				left = append(left,el)
			} else {
				right = append(right,el)
			}
		}
		leftSort  := QSort(left)
		rightSort := QSort(right)
		var res []int
		res = append(res,leftSort...)
		res = append(res,op)
		res = append(res,rightSort...)
		return res
	}
}
func BinarySearchMem(list []int,item int) (int,int) {
	// возвращаем вторым числом индекс числа которое большее из меньших
	low  := 0
	high := len(list)-1
	var mid int
	lastInt := -1
	lastIndex := -1
	for low <= high {
		mid = (low + high) / 2
		quess := list[mid]
		if quess == item {
			return quess,lastIndex
		}
		if quess > item {
			high = mid-1
		} else {
			if quess > lastInt {
				lastInt = quess
				lastIndex = mid
			}
			low = mid+1
		}
	}
	return -1,lastIndex
}
func BinarySearchFull(list []int,item int) (int,int) {
	// возвращаем вторым числом индекс числа которое большее из меньших
	low  := 0
	high := len(list)-1
	var mid int
	lastInt := -1
	lastIndex := -1
	for low <= high {
		mid = (low + high) / 2
		quess := list[mid]
		if quess == item {
			lastIndex = mid
			return quess,lastIndex
		}
		if quess > item {
			high = mid-1
		} else {
			if quess > lastInt {
				lastInt = quess
				lastIndex = mid
			}
			low = mid+1
		}
	}
	return -1,lastIndex
}