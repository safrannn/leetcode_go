package problem0170

type TwoSum struct {
    data map[int]int 
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
    data := make(map[int]int)
    return TwoSum{data:data}
}


/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int)  {
    this.data[number]++
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    for i:= range this.data{
        _,prs := this.data[value-i]
        if prs {
            if value - i == i{
                if this.data[value-i] > 1{
                    return true
                }else{
                    continue
                }
            }else{
                return true
            }
        }
    }
    return false
}
