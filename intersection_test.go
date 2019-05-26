package main

import (
	"testing"
	"github.com/franela/goblin"
)
func TestNaivIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("naivIntersection O(n*m)",func() {
		g.It("naivIntersection [6]", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := naivIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("naivIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := naivIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("naivIntersection {3,23,48,22}", func() {
			result := []int{3,23,48,22}
			unsortN := []int{3,43,23,58,48,50,22,107,207,40,18,99,110,666,71,81} // len 16
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := naivIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("naivIntersection {}", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := naivIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("naivIntersection {34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := naivIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
func TestSortIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("sortIntersection O(n+m)",func() {
		g.It("sortIntersection {6}", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := sortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("sortIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := sortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("sortIntersection {}", func() {
			// Алгоритм не работает с неотсортированными множествами
			// если ничего не нашел тогда тест пройден
			result := []int{}
			unsortN := []int{43,23,58,48,50,22,107,207,40,18,99,110,666,71,81} // len 16
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := sortIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("sortIntersection {}", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := sortIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("sortIntersection {34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := sortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
func TestNoSortIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("NoSortIntersection O(n log n)",func() {
		g.It("NoSortIntersection {6}", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := NoSortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("NoSortIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := NoSortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("NoSortIntersection 3", func() {
			result := []int{23,48,22}
			unsortN := []int{43,23,58,48,50,22,107,207,40,18,99,110,666,71,81} // len 16
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := NoSortIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("NoSortIntersection []", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := NoSortIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("NoSortIntersection {34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := NoSortIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
//HashMapIntersection
func TestHashMapIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("HashMapIntersection O(n+m)", func() {
		g.It("HashMapIntersection [6]", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := HashMapIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("HashMapIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := HashMapIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("HashMapIntersection {23,48,22}", func() {
			result := []int{23,48,22}
			unsortN := []int{43,23,58,48,50,22,107,207,40,18,99,110,666,71,81} // len 16
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := HashMapIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("HashMapIntersection []", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := HashMapIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("HashMapIntersection 34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := HashMapIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
func TestHvangAndLinIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("HvangAndLinIntersection O(n + (m/4)log m)", func() {
		g.It("HvangAndLinIntersection [6]", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := HvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("HvangAndLinIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := HvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("HvangAndLinIntersection {23,48,22}", func() {
			result := []int{}
			unsortN := []int{43,23,58,50,22,107,207,40,18,99,110,666,71,81} // len 16
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := HvangAndLinIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
			//g.Assert(res).Eql(result)
		})
		g.It("HvangAndLinIntersection []", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := HvangAndLinIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("HvangAndLinIntersection 34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := HvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
func TestMemHvangAndLinIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("MemHvangAndLinIntersection O(n + (m/4)log m)", func() {
		g.It("MemHvangAndLinIntersection [6]", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := MemHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("MemHvangAndLinIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := MemHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("MemHvangAndLinIntersection []", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := MemHvangAndLinIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("MemHvangAndLinIntersection 34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := MemHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}
func TestFullHvangAndLinIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("FullHvangAndLinIntersection O(n + (m/4)log m)", func() {
		g.It("FullHvangAndLinIntersection [6]", func() {
			result := []int{6}
			sortN := []int{1,3,6,9,11} // len 5
			sortM := []int{2,5,6,8,12} // len 5
			res := FullHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("FullHvangAndLinIntersection {6,35,57}", func() {
			result := []int{6,35,57}
			sortN := []int{1,3,6,9,11,27,30,35,49,57,60,89}    // len 12
			sortM := []int{2,5,6,8,12,13,17,23,35,36,48,57,68} // len 13
			res := FullHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
		g.It("FullHvangAndLinIntersection []", func() {
			result := []int{}
			unsortN := []int{} // len 0
			unsortM := []int{3,23,15,7,48,22,44,17,98,0,6,77} // len 12
			res := FullHvangAndLinIntersection(unsortN,unsortM)
			g.Assert(res).Equal(result)
		})
		g.It("FullHvangAndLinIntersection 34, 147, 352, 457, 520}", func() {
			result := []int{34, 147, 352, 457, 520}
			sortN := []int{15,23,34,38,42,65,71,113,147,200,212,222,280,341,352,387,394,457,515,516,520,569} // len 16 [34,147,352,457,520]
			sortM := []int{17,19,24,30,34,37,52,89,97,147,198,220,313,352,457,520} // len 14
			res := FullHvangAndLinIntersection(sortN,sortM)
			g.Assert(res).Equal(result)
		})
	})
}