from dataclasses import dataclass, field
# D = open("input.txt").read().strip()
D = open("test.txt").read().strip()
input = D.split("\n\n")

@dataclass
class Part:
    a: int
    m: int
    s: int
    x: int

    @property
    def rat(self):
        return self.a + self.m + self.s + self.x


TOTAL = 4000
class Workflow:
    def __init__(self, a):
        if ":" in a:
            op, dest = a.split(":")
        else:
            dest = a
            op = None

        self.op = op
        self.dest = dest

    def __repr__(self) -> str:
        return f"{self.op} - {self.dest}"
    
    def consume(self, part: Part) -> str|None:
        if not self.op:
            return self.dest
        
        a = part.a
        m = part.m
        s = part.s
        x = part.x
        if eval(self.op):
            return self.dest
        return None
    
    # need to carry forward the constraints
    # track the tightest upper and lower bounds
    # for each letter and multiple those ranges to get
    # the number of possibilities
    # and then once you get to A that is the number of that node
    # no need to 
    def calc_combinations(self, past: int) -> int:
        if self.op is None:
            combs = pow(TOTAL,4) - past
            return combs
        if "<" in self.op:

            _, numb = self.op.split("<")
            combs = (int(numb)-1) * pow(TOTAL,3) - past
            # breakpoint()
            ...
            return combs
        if ">" in self.op:
            _, numb = self.op.split(">")
            combs = (TOTAL-int(numb)) * pow(TOTAL,3) - past
            return combs

        
        


    
workflows: dict[str, list[Workflow]] = {}
for f in input[0].split("\n"):
    flow = f.split("{")
    name = flow[0]
    dirs = [Workflow(a) for a in flow[1].rstrip("}").split(",")]
    workflows[name] = dirs



START = "in"
acc = 0
for p in input[1].split("\n"):
    pa = [int(a.split("=")[1]) for a in sorted(p.lstrip("{").rstrip("}").split(","))]
    part =Part(*pa)
    workflow = workflows[START]
    accrej = None
    while accrej is None:
        for w in workflow:
            if dest := w.consume(part):
                if dest == 'R':
                    accrej = False
                elif dest == 'A':
                    accrej = True
                else:
                    workflow = workflows[dest]
                break
    acc += part.rat if accrej else 0



@dataclass
class Node:
    val: int
    label: str
    nodes: list["Node"] = field(default_factory=lambda: [])

start = Node(val=1, label="in", nodes=[])
total_part_2 = 0
def add_nodes(node: Node, prev_max:int):
    combs_so_far = 0

    if node.label in workflows:
        for wf in workflows[node.label]:
            count = wf.calc_combinations(combs_so_far)
            combs_so_far += count
            new_node = Node(count, wf.dest, [])
            node.nodes.append(new_node)
            add_nodes(new_node)

add_nodes(start, pow(4000,4))

def calc_total(node: Node) -> int:
    total = 0
    for n in node.nodes:
        v = calc_total(n)
        total += v 
    print(total)
    return total * node.val

print(calc_total(start))