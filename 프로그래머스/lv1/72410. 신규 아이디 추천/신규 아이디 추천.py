import re

def solution(new_id: str):
    answer = new_id.lower()
    answer = re.sub(r"[^a-z0-9\-\_\.]","",answer)
    while ".." in answer:
        answer = answer.replace("..", ".")

    if answer[0] == ".":
        answer = answer[1:]
    if len(answer) != 0:
        if answer[-1] == ".":
            answer = answer[:-1]
    if len(answer) == 0:
        answer = "a"
    print(answer)
    if len(answer) >= 16:
        answer = answer[:15]
        if answer[-1] == ".":
            answer = answer[:-1]
    print(answer)

    while len(answer) < 3:
        answer = answer + answer[-1]
    print(answer)

    return answer