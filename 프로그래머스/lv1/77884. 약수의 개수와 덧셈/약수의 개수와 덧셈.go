func cntDiv(val int) int {
    cnt := 0
    for i:=1; i <= val; i++ {
        if val%i == 0 {
            cnt += 1
        }
    }
    
    if cnt%2 == 0 {
        return val
    } else {
        return val * -1
    }
}

func solution(left int, right int) int {
    result := 0
    for i := left; i <= right; i++ {
        result += cntDiv(i)
    }
    
    return result
}