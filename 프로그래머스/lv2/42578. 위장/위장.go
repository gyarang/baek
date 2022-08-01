func solution(clothes [][]string) int {
    counter := make(map[string]int)
    
    for _, cloth := range clothes {
        counter[cloth[1]]++
    }
    
    sum := 1
    
    for _, val := range counter {
        sum *= (val + 1)
    }
    
    return sum - 1
}