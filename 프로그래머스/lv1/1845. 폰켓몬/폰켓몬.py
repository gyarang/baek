def solution(nums):
    s = set(nums)
    size = len(s)
    if (size > len(nums)/2):
        return len(nums)/2
    else:
        return size
    