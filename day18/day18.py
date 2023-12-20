D = open("input.txt").read().strip()
# D = open("test.txt").read().strip()
input = [r.split(" ") for r in D.split("\n")]

max_right = 0
min_right = 0
curr_right = 0
max_down = 0
min_down = 0
curr_down = 0

for r in input:
    if r[0] == 'R':
        curr_right += int(r[1])
        max_right = max(max_right, curr_right)
    if r[0] == 'L':
        curr_right -= int(r[1])
        min_right = min(min_right, curr_right)
    if r[0] == 'D':
        curr_down += int(r[1])
        max_down = max(max_down, curr_down)
    if r[0] == 'U':
        curr_down -= int(r[1])
        min_down = min(min_down, curr_down)

trench = [['.' for _ in range(0,max_right-min_right+1)] for __ in range(0, max_down-min_down+1)]
pos = [abs(min_down),abs(min_right)]
print(pos)
trench[pos[0]][pos[1]] = '#'
for dirrow in input:
    dist = int(dirrow[1])
    try:
        if dirrow[0] == 'R':
            for col in range(pos[1]+1, pos[1] + dist+1):
                trench[pos[0]][col] = '#'
                pos[1] = col
        elif dirrow[0] == 'L':
            for col in range(pos[1]-1, pos[1] - dist-1, -1):
                trench[pos[0]][col] = '#'
                pos[1] = col
        elif dirrow[0] == 'D':
            for row in range(pos[0]+1, pos[0] + dist+1):
                trench[row][pos[1]] = '#'
                pos[0] = row
        elif dirrow[0] == 'U':
            for row in range(pos[0]-1, pos[0] - dist-1,-1):
                trench[row][pos[1]] = '#'
                pos[0] = row
    except:
        with open('output.txt', 'w+') as r:
            for t in trench:
                r.write(''.join(t))
                r.write("\n")
                new = ''.join(t).lstrip('.').rstrip('.')

        breakpoint()
with open('output.txt', 'w+') as r:
    for t in trench:
        r.write(''.join(t))
        r.write("\n")
        new = ''.join(t).lstrip('.').rstrip('.')
scores = []
for t in trench:
    tscoe = 0
    in_trench= False
    i = 0

    last_t = 0
    for j, r in enumerate(t):
        # print(i, r)
        if r == '#':
            last_t = j
    # print(i, last_t)


    while i <= last_t:
        tscoe += 1 if in_trench else 0
        if t[i] == "#":
            while i < len(t) and t[i] == "#":
                tscoe += 1
                i += 1
            in_trench = not in_trench
        i += 1
    # print(tscoe)
    # t.append(f" {str(tscoe)}")
    scores.append(tscoe)

print(sum(scores))

with open('output.txt', 'w+') as r:
    for i,t in enumerate(trench):
        t.append(f" {str(scores[i])}")
        r.write(''.join(t))
        r.write("\n")
        new = ''.join(t).lstrip('.').rstrip('.')