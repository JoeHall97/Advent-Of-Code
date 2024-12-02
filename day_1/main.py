import bisect
from functools import reduce

def diff(l: int, r: int) -> int:
    return l - r

def puzzle_one():
    left_list, right_list = [], []
    with open('puzzle_input.txt') as f:
        line = f.readline()
        while line:
            l, r = line.split()
            bisect.insort(left_list, int(l))
            bisect.insort(right_list, int(r))
            line = f.readline()

    print(sum(abs(l-r) for l,r in zip(left_list, right_list)))

def puzzle_two():
    left_list, right_list = [], []
    with open('puzzle_input.txt') as f:
        line = f.readline()
        while line:
            l, r = line.split()
            left_list.append(int(l))
            right_list.append(int(r))
            line = f.readline()

    right_dict = {}
    for i in set(right_list):
        right_dict[i] = right_list.count(i)
    
    sum_duplicates = 0
    for l in left_list:
        if right_dict.get(l):
            sum_duplicates += l * right_dict[l]
    print(sum_duplicates)

if __name__ == "__main__":
    # puzzle_one()
    puzzle_two()