package problem0161

func isOneEditDistance(s string, t string) bool {
    if len(s) - len(t) > 1 || len(t) - len(s) > 1 || s == t {
        return false
    }
    
    diffFound := false
    sPointer, tPointer := 0,0
    for sPointer < len(s) && tPointer < len(t){
        if s[sPointer] != t[tPointer]{
            if diffFound{
                return false
            }
            diffFound = true
            if len(s) < len(t){
                sPointer--
            }else if len(s) > len(t){
                tPointer--
            }
        }
        sPointer++
        tPointer++
    }
    
    return true
}