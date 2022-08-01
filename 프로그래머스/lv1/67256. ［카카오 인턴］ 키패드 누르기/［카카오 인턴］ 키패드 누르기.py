nums = [[1,3],
        [0,0], [1,0], [2,0],
        [0,1], [1,1], [2,1],
        [0,2], [1,2], [2,2]]

def getLength(p1, p2):
    x = abs(p1[0] - p2[0])
    y = abs(p1[1] - p2[1])
    return x + y

def solution(numbers, hand):
    lp = [0, 3]
    rp = [2, 3]
    
    result = ""
    
    for n in numbers:
        if n in [1, 4, 7]:
            lp = nums[n]
            result += "L"
        elif n in [3, 6, 9]:
            rp = nums[n]
            result += "R"
        else:
            lLen = getLength(lp, nums[n])
            rLen = getLength(rp, nums[n])
            if lLen > rLen:
                result += "R"
                rp = nums[n]
            elif lLen < rLen:
                result += "L"
                lp = nums[n]
            else:
                if hand == "right":
                    result += "R"
                    rp = nums[n]
                else:
                    result += "L"
                    lp = nums[n]
    
    return result