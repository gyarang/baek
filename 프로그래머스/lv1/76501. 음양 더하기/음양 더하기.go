func solution(absolutes []int, signs []bool) int {
    result := 0
    
    for i, v := range absolutes {
        input := v
        if !signs[i]{
            input = v * -1
        }
        result += input
    }
    
    return result
}