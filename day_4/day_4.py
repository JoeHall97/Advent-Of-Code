# list of checks to do
# each list is an (x, y) offset that need to be applied when iterating
checks: list[list[int]] = [
    # check right
    [1, 0],
    # check left
    [-1, 0],
    # check down
    [0, 1],
    # check up
    [0, -1],
    # check diagonal down right
    [1, 1],
    # check diagonal down left
    [-1, 1],
    # check diagonal up right
    [1, -1],
    # check diagonal up left
    [-1, -1]
]
word = 'MAS'
word_set = set(['M','A','S'])


def boundsCheck(x, y, con):
    return y >= 0 and y < len(con) and x >= 0 and x < len(con[y])


def performChecks(x: int, y: int, con: list[str]) -> int:
    count = 0
    for cx, cy in checks:
        tempX = x + cx
        tempY = y + cy
        currCharPos = 0
        while boundsCheck(tempX, tempY, con) and currCharPos < len(word) \
                and con[tempY][tempX] == word[currCharPos]:
            tempX += cx
            tempY += cy
            currCharPos += 1
        if currCharPos == len(word):
            count += 1
    return count


def performChecks2(x: int, y: int, con: list[str]) -> int:
    count = 0
    if boundsCheck(x-1, y-1, con) and boundsCheck(x+1,y+1,con):
        l_diag = set([con[y][x],con[y+1][x+1],con[y-1][x-1]])
        r_diag = set([con[y][x],con[y+1][x-1],con[y-1][x+1]])
        if l_diag == word_set and r_diag == word_set:
            count = 1
    return count


def puzzleTwo():
    with open('puzzle_input.txt') as f:
        contents = f.read().split('\n')
    total = 0
    for i, line in enumerate(contents):
        for j, ch in enumerate(line):
            if ch != 'A':
                continue
            total += performChecks2(j, i, contents)
    print(total)


def puzzleOne():
    with open('puzzle_input.txt') as f:
        contents = f.read().split('\n')
    total = 0
    for i, line in enumerate(contents):
        for j, ch in enumerate(line):
            if ch != 'X':
                continue
            total += performChecks(j, i, contents)
    print(total)


puzzleOne()
puzzleTwo()