package problem1213

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    ptr1, ptr2, ptr3 := 0, 0, 0
    result := []int{}
    for ptr1 < len(arr1) && ptr2 < len(arr2) && ptr3 < len(arr3){
        if arr1[ptr1] == arr2[ptr2] && arr1[ptr1] == arr3[ptr3]{
            result = append(result,arr1[ptr1])
            ptr1++
            ptr2++
            ptr3++
        }else{
            if arr1[ptr1] <= arr2[ptr2] && arr1[ptr1] <= arr3[ptr3]{
                ptr1++
            }else if arr2[ptr2] <= arr1[ptr1] && arr2[ptr2] <= arr3[ptr3]{
                ptr2++
            }else{
                ptr3++
            }
        }
    } 
    return result
}
