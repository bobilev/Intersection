package main

import (
	"github.com/bobilev/Intersection/util"
)

func main() {

}
func naivIntersection(n,m []int) []int {
	// Работает с (не)упорядочными множествами
	// O(n*m)
	// простой перебор, это самый не эффективный и скучный алгоритм
	// вставка нового элемента в итоговое множество производиться за О(1), эта константа не влияет на общию асимптотику
	result := []int{}
	for _,val := range n {
		for _,jval := range m {
			if val == jval {
				result = append(result,val)
			}
		}
	}
	return result
}
func sortIntersection(n,m []int) []int {
	// Работает с упорядочными множествами
	// O(n+m)
	// так как перед нами упорядоченные множества, мы можем этим воспользоваться.
	// если элементы отсортированы по возрастанию) то при проверки двух паралельных элементов двух множеств, алгоритм
	// может больше не возвращаться к началу т.к последующие элементы в двух множествах не могут там оказаться,
	// следовательно искать там нет смысла.
	// вставка нового элемента в итоговый массив производиться за О(1), эта константа не влияет на общию асимптотику
	result := []int{}
	var i,j int
	for i < len(n) && j < len(m){
		if n[i] == m[j] {
			result = append(result,n[i])
			i++;j++
		} else {
			if n[i] > m[j] {
				j++
			} else {
				i++
			}
		}
	}
	return result
}
func NoSortIntersection(n,m []int) []int {
	// Работает с упорядочными множествами
	// O()-> O(n log n) + (O(log n) * n) которое так же, как O(n log n)
	// тут сортируем меньшее множество, и ищем в нем элементы второго множества с помощью бинарного поиска
	result,bigInt,minInt := []int{},[]int{},[]int{}
	if len(n) > len(m) {
		minInt = util.QSort(m)
		bigInt = n
	} else {
		minInt = util.QSort(n)
		bigInt = m
	}
	for _,val := range bigInt {
		if findInt := util.BinarySearch(minInt,val); findInt == val {
			result = append(result,val)
		}
	}
	return result
}
func HashMapIntersection(n,m []int) []int {
	// Работает с (не) упорядочными множествами
	// O(n+m)
	// поместить первое(любой) множество в хеш таблицу, и искать элементы второго множества в хеш таблице.
	// поиск в хеш таблице занимает O(1) константное время и по этому не влияет на общую асимптотику.
	// преимущество этого алгоритма в том что мы получаем асимптотику O(n+m) как при отсортированном множестве, даже если
	// множества не отсортированны
	result := []int{}
	mapList := make(map[int]int)
	for _,val := range n {
		mapList[val] = val
	}
	for _,val := range m {
		if _ ,ok := mapList[val]; ok {
			result = append(result,val)
		}
	}
	return result
}
func HvangAndLinIntersection(n,m []int) []int {
	// Работает с упорядочными множествами
	// O(n + (m/4)log m)
	// Алгоритм Хванга и Лина (1972 г Ф. К. Хванга, С. Лин) -
	// Пересечение упорядоченных множеств неравных размеров с (возможно) меньшим количеством сравнений.
	// Это работает, вычисляя шаг от отношения m/n и делая сравнение с тем элементом в большом списке:
	// например, если m/n = 4, он будет сравнивать четвертый элемент большего списка с первым из другого, и либо исключит все 4 элемента,
	// либо выполнит в них бинарный поиск для правильной точки пересечения. И если (не)найдет то сдвинуть позицию n на 1
	result,bigInt,minInt := []int{},[]int{},[]int{}
	if len(n) > len(m) {
		bigInt = n
		minInt = m
	} else {
		bigInt = m
		minInt = n
	}
	var i,j int
	for i < len(minInt) {
		if minInt[i] == bigInt[j] {
			result = append(result,minInt[i])
			i++
		} else {
			if minInt[i] < bigInt[j] {
				// n < m в этих 4 числах возможно есть нужное нам число
				// бинарный поиск в этом отрезке
				var lowJ,highJ int
				if j >= 3 {
					lowJ = j - 3
				}
				if j != len(bigInt) {
					highJ = j +1
				}
				if findInt := util.BinarySearch(bigInt[lowJ:highJ],minInt[i]); findInt == minInt[i] {
					result = append(result,findInt)
				}
				i++
			} else {
				// Это обход ловушки в котрую мы можем попасть из-за особнености алгоритма
				// если последнй элемент n будет больше последнего элемента m
				// то это приведет к бесконечности т.к n и m будут каждый раз одинаковыми
				if i == len(minInt)-1 && j == len(bigInt)-1 {
					i++
				}
				// смело пропускаем эти 4 элемента
				if j == 0 {
					j = j+3
				} else {
					j = j+4
				}
				if j > len(bigInt)-1 {
					j = len(bigInt)-1
				}
			}
		}
	}
	return result
}
func MemHvangAndLinIntersection(n,m []int) []int {
	// Работает с упорядочными множествами
	// O(n + (m/4)log m)
	// Я подумал что будет уместно прикрутить мемоизацию, во время бинарного поиска запоминать элементы меньше чем n
	// чтобы не включать их в следующие m4 элемента повторно. Так мы сможем двигаться маленько быстрее.
	// Лучшем случаи и среднем случаи это (возможно) повлияет на производительность. Затрудняюсь вывести точную асимтотику
	result,bigInt,minInt := []int{},[]int{},[]int{}
	if len(n) > len(m) {
		bigInt = n
		minInt = m
	} else {
		bigInt = m
		minInt = n
	}
	var i,j int
	for i < len(minInt) {
		if minInt[i] == bigInt[j] {
			result = append(result,minInt[i])
			i++
		} else {
			if minInt[i] < bigInt[j] {
				// n < m в этих 4 числах возможно есть нужное нам число
				// бинарный поиск
				var lowJ,highJ int
				if j >= 3 {
					lowJ = j - 3
				}
				if j != len(bigInt) {
					highJ = j +1
				}
				findInt,lastInt := util.BinarySearchMem(bigInt[lowJ:highJ],minInt[i])
				if findInt == minInt[i] {
					result = append(result,findInt)
				}
				if lastInt != -1 {
					j = j+lastInt+1
					if j > len(bigInt)-1 {
						j = len(bigInt)-1
					}
				}
				i++
			} else {
				if i == len(minInt)-1 && j == len(bigInt)-1 {
					i++
				}
				if j == 0 {
					j = j+3
				} else{
					j = j+4
				}
				if j > len(bigInt)-1 {
					j = len(bigInt)-1
				}
			}
		}
	}
	return result
}
func FullHvangAndLinIntersection(n,m []int) []int {
	// Работает с упорядочными множествами
	// Мемоизация навела меня на мысль, а что если буду чередовать множества между собой,
	// ищу n число отрезке из 4 элементов m множества бинарным поиском,
	// запоминаю большее из меньших чисел которое пытались искать. И теперь меняем множества между собой
	// следующее из больших среди меньших в бинарном поиске становится новой n которое мы будем искать во втором множестве которое раньше было перым
	// Тем самым мы можем двигаться по двум множествам с O((x/4)log x) в лучшем случае  x - это n или m множество
	// O(n + (m/4)log m) - в среднем случаи как будто это MemHvangAndLinIntersection или просто HvangAndLinIntersection
	// O(n+m) в худшем случаи как и остальные два алгоритма. То есть мы не потеряли а только приобрели (возможный) прирост.
	// Наверное его надо на доске словами объяснять, мне кажеться я неявно описал его.
	result := []int{}
	var i,j,transLeft,transRight int
	var transit bool
	// transit: - позволит менять местами множества, так бинарный поиск будет работать в обоих множествах
	// 			true  = left <- n | right <- m
	// 			false = left <- m | right <- n
	// i,j - увеличиваються на 4, и их предел это конец своих массивов
	// lastN,lastM - большее из меньших значений из отрезка (m/n)4 в бинарном поиске и не больше значение которое пытаемся найти
	for i < len(n) && j < len(m){
		// выбор транзитных значений для N и M
		transLeft,transRight = n[i],m[j]
		if transit {
			transLeft,transRight = m[j],n[i]
		}
		if transLeft == transRight {
			result = append(result,transLeft)
			i++;j++
		} else {
			if transLeft < transRight {
				// n < m в этих 4 числах возможно есть нужное нам число
				// бинарный поиск
				var low,high,ti int // ti - transit increment - это может быть как i так и j
				left,right,tj,ti := n,m,j,i
				if transit {
					left,right,tj,ti  = m,n,i,j
				}
				if tj >= 3 {
					low = tj - 3
				}
				if tj != len(m) {
					high = tj+1
				}
				findInt,lastInt := util.BinarySearchFull(right[low:high],left[ti])
				if findInt == left[ti] {
					result = append(result,findInt)
				}
				if i != 0  {
					if !transit {
						minJ:=3
						if lastInt == -1 {
							j = j - minJ
						} else{
							j = (j-minJ)+(lastInt+1)
						}
						if j > len(m)-1 {
							j = len(m)-1
						}
					} else {
						minI:=3
						if lastInt == -1 {
							i = i - minI
						} else{
							i = (i-minI)+(lastInt+1)
						}
						if i > len(n)-1 {
							i = len(n)-1
						}
					}
				}
				if !transit {
					i = i+4
					if i > len(n)-1 {
						i = len(n)-1
					}
				} else {
					j = j+4
					if j > len(m)-1 {
						j = len(m)-1
					}
				}
				transit = !transit // переключаемся на другое множество
			} else {
				if !transit {
					if i == len(n)-1 && j == len(m)-1 {
						i++
					}
					if j == 0 {
						j = 3
					} else{
						j = j+4
					}
					if j > len(m)-1 {
						j = len(m)-1
					}
				} else {
					if i == len(n)-1 && j == len(m)-1 {
						j++
					}
					if i == 0 {
						i = 3
					} else{
						i = i+4
					}
					if i > len(n)-1 {
						i = len(n)-1
					}
				}
			}
		}
	}
	return result
}