package problem0386

func lexicalOrder(n int) []int {
    result := make([]int,0);
    dfs(&result, 0, n);
    return result
}
func dfs(result *[]int, current int, n int){
    for i:=0;i<10;i++{
        temp := current*10 + i;
        if temp != 0 && temp <= n {
            *result = append(*result,temp);
            dfs(result,temp,n);
        }
    }
}