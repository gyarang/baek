def solution(board, moves):
    stack = []
    result = 0

    for m in moves: 
        for line in board:
            if line[m-1] != 0:
                stack.append(line[m-1])
                line[m-1] = 0
                if len(stack) >= 2:
                    if stack[-1] == stack[-2]:
                        result += 2
                        stack.pop()
                        stack.pop()
                break
    return result