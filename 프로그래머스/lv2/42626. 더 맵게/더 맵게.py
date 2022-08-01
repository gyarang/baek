import heapq

def solution(scoville, K):
    count = 0
    heapq.heapify(scoville)
    
    while scoville[0] <= K:
        if len(scoville) == 1:
            return -1
        min1 = heapq.heappop(scoville)
        min2 = heapq.heappop(scoville)
        
        heapq.heappush(scoville, min1 + (min2 * 2))
        count += 1

    return count