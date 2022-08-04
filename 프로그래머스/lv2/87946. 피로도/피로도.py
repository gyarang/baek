import copy

max = 0
def dfs(cnt, k, dungeons):
    global max
    if cnt > max :
        max = cnt

    if len(dungeons) == 0 :
        return

    for i, dungeon in enumerate(dungeons) :
        if k >= dungeon[0] :
            remainDungeons = copy.copy(dungeons)
            del remainDungeons[i]
            dfs(cnt + 1, k - dungeon[1], remainDungeons)


def solution(k, dungeons):
    global max
    dfs(0, k, dungeons)
    return max