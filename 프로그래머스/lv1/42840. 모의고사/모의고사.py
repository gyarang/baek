from numpy import asarray
from collections import Counter

def set_supo_answers(supo_pattern, count):
    share, rest = count // len(supo_pattern), count % len(supo_pattern)
    answers = (supo_pattern * share) + supo_pattern[:rest]

    return answers
    
def solution(answers):
    supo1 = [1, 2, 3, 4, 5]
    supo2 = [2, 1, 2, 3, 2, 4, 2, 5]
    supo3 = [3, 3, 1, 1, 2, 2, 4, 4, 5, 5]

    answers = asarray(answers)

    supo1_answers = asarray(set_supo_answers(supo1, len(answers)))
    supo2_answers = asarray(set_supo_answers(supo2, len(answers)))
    supo3_answers = asarray(set_supo_answers(supo3, len(answers)))

    supo1_point = answers - supo1_answers
    print(supo1_point)
    supo2_point = answers - supo2_answers
    supo3_point = answers - supo3_answers

    supo1_point = Counter(supo1_point)[0]
    supo2_point = Counter(supo2_point)[0]
    supo3_point = Counter(supo3_point)[0]

    points = {'1': supo1_point, '2': supo2_point, '3': supo3_point}
    points = sorted(points.items(), key=(lambda x:x[1]), reverse=True)
    print(points)
    
    top_score = points[0][1]

    result = []
    for point in points:
        if point[1] == top_score:
            result.append(int(point[0]))
    
    return result