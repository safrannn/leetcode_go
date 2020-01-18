package problem0373

type Pair struct{
    Index1 int
    Index2 int
    Val int
}

type PairHeap []Pair

func (ph *PairHeap) Len() int {
    return len(*ph)
}

func (ph *PairHeap) Less(i, j int) bool {
    return (*ph)[i].Val < (*ph)[j].Val
}

func (ph *PairHeap) Swap(i, j int) {
    (*ph)[i],(*ph)[j] = (*ph)[j],(*ph)[i]
}

func (ph *PairHeap) Push(v interface{}) {
    *ph = append(*ph, v.(Pair))
}
func (ph *PairHeap) Pop() interface{} {
    root := (*ph)[len(*ph)-1]
    *ph = (*ph)[:len(*ph)-1]
    return root
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return nil
    }
    // make heap
    pairHeap := new(PairHeap)
    heap.Init(pairHeap)
    
    for i := range nums2{
        heap.Push(pairHeap, Pair{0, i, nums1[0]+nums2[i]})
    }
    
    result := make([][]int, 0)
    if k > len(nums1) * len(nums2){
        k = len(nums1) * len(nums2)
    }
    
    for i := 0; i < k ;i++{
        pair := heap.Pop(pairHeap).(Pair)
        result = append(result, []int{nums1[pair.Index1], nums2[pair.Index2]})
        if pair.Index1 < len(nums1) - 1{
            heap.Push(pairHeap, Pair{pair.Index1 + 1, pair.Index2, nums1[pair.Index1 + 1] + nums2[pair.Index2]})
        }
    }
    return result
}



