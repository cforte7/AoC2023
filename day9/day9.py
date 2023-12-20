# D = open("input.txt").read().strip()
D = open("test.txt").read().strip()
s = D.split("\n")

score = 0
for r in s:
    row = [int(n) for n in r.split(" ")]
    end = []
    diff = [12321]
    while diff and not all(a == 0 for a in diff):
        for i in range(len(row)-1):
            diff.append(row[i+1]-row[i])
        print(diff)
        end.append(row[-1]+diff[-1])
        row = diff
    score += sum(end)
    print(score)
print(score)
        
