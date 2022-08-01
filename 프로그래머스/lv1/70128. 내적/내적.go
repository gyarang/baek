func solution(a []int, b []int) int {
    result := 0
    
    for i, v := range(a){
        result += (v * b[i])
    }
    
    return result
}