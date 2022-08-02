def solution(numbers):
    results = []
    for i in range(len(numbers)):
        for j in range(i+1, len(numbers)):
            results.append(numbers[i] + numbers[j])
    
    results = list(set(results))
    results.sort()
    return results
