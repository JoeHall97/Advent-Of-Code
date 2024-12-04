import re

def puzzle_two():
    with open('puzzle_input.txt') as f:
        contents = f.read()
        found: list[str] = re.findall(r"(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))", contents)
        sum = 0
        do = True
        for f in found:
            # check for do()
            if len(f) == 4:
                do = True
            # check for don't()
            elif len(f) == 7:
                do = False
            elif do:
                nums = f[4:len(f)-1].split(',')
                sum += int(nums[0]) * int(nums[1])
        print(sum)

def puzzle_one():
    with open('puzzle_input.txt') as f:
        contents = f.read()
        found: list[str] = re.findall(r"(mul\(\d{1,3},\d{1,3}\))", contents)
        sum = 0
        for f in found:
            nums = f[4:len(f)-1].split(',')
            sum += int(nums[0]) * int(nums[1])
        print(sum)

# puzzle_one()
puzzle_two()