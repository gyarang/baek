func solution(progresses []int, speeds []int) []int {
    var result []int

    for len(progresses) > 0 {
        if progresses[0] >= 100 {
            i := 1
            for ; i < len(progresses); i++ {
                if progresses[i] < 100 {
                    break
                }
            }
            progresses = progresses[i:]
            speeds = speeds[i:]
            result = append(result, i)
        } else {
            for i := 0; i < len(progresses); i++ {
                progresses[i] += speeds[i]
            }
        }
    }

    return result
}