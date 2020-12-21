class Matching(object):
    def __init__(self, size):
        self.p_to_a = [-1]*size
        self.a_to_p = [-1]*size

    def set_match(self, p, a):
        # if p already has an a, unassign the a's p
        if self.p_to_a[p] != -1:
            self.a_to_p[self.p_to_a[p]] = -1
        # if a already has a p, unassign the p's a
        if self.a_to_p[a] != -1:
            self.p_to_a[self.a_to_p[a]] = -1
        self.p_to_a[p] = a
        self.a_to_p[a] = p

    def get_a(self, idx):
        return self.a_to_p[idx]

    def get_p(self, idx):
        return self.p_to_a[idx]

    def write(self):
        print('p to a', self.p_to_a)
        print('a to p', self.a_to_p)

def construct_table(preferences):
    table = []
    for i in range(len(preferences)):
        table.append([-1]*len(preferences))
    for p_idx, p in enumerate(preferences):
        for rank_idx, a in enumerate(p):
            table[p_idx][a] = rank_idx
    return table

# Gale-Shapley algorithm
def match(proposers, acceptors):
    matching = Matching(len(proposers))
    free_p = range(len(proposers))
    p_index = [0]*len(proposers)
    a_table = construct_table(acceptors)
    while len(free_p) > 0:
        p = free_p.pop(0)
        a = proposers[p][p_index[p]]
        existing_p = matching.get_a(a)
        if existing_p == -1:
            matching.set_match(p, a)
        else:
            if a_table[a][p] < a_table[a][existing_p]:
                matching.set_match(p, a)
                free_p.append(existing_p)
            else:
                free_p.append(p)

        p_index[p] += 1

    return matching

if __name__ == '__main__':
    proposers = [
        [1, 0, 2, 3],
        [3, 0, 1, 2],
        [0, 2, 1, 3],
        [1, 2, 0, 3],
    ]
    acceptors = [
        [0, 2, 1, 3],
        [2, 3, 0, 1],
        [3, 1, 2, 0],
        [2, 1, 0, 3],
    ]
    # These two stable matchings are the two ends of a lattice of stable matchings.
    match(proposers, acceptors).write()
    match(acceptors, proposers).write()
