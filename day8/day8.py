D = open("input.txt").read().strip()
# D = open("test.txt").read().strip()
s = D.split("\n")
LR = s[0]
turns = s[2:]
import math

dir_lk = {}
for l in turns:
    node, dirs = l.split(" = ")
    dir_lk[node] = dirs.strip("(").strip(")").split(", ")

# curr_node = "AAA"
# i = 0
# while curr_node != "ZZZ":
    # LR_ind = i % len(LR)
    # dirr = dir_lk[curr_node]
    # if LR[LR_ind] == "L":
    #     curr_node = dirr[0]
    # else:
    #     curr_node = dirr[1]
#     print(curr_node)
#     i += 1
# print("steps: ", i)

starts = []
starts_copy = []
for l in turns:
    node, dirs = l.split(" = ")
    if node[-1] == "A":
        starts.append(node)
        starts_copy.append(node)
    dir_lk[node] = dirs.strip("(").strip(")").split(", ")

i = 0
# print(starts)
# print(dir_lk)

print(len(starts))

z_lens = {}
# while not all(n[-1] == "Z" for n in starts):
while len(z_lens) < len(starts_copy):
    LR_ind = i % len(LR)
    i += 1
    for j in range(len(starts)):
        if z_lens.get(starts[j]):
            continue
        dirr = dir_lk[starts[j]]
        if LR[LR_ind] == "L":
            starts[j] = dirr[0]
        else:
            starts[j] = dirr[1]
        if starts[j][-1] == "Z":
            z_lens[starts_copy[j]] = i

values = [v for v in z_lens.values()]
print(values)
lcm = math.lcm(*values)
print(lcm)
# for x in range(2, len(values)):
#     lcm = math.lcm(lcm, x)
# print(lcm)
# print(math.lcm(13770,  math.lcm(17286,  math.lcm(17872,  math.lcm(19630,  math.lcm(20802, 23146))))))