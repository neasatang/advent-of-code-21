import pprint

numbers_drawn = []
boards = {}

def create_boards(): 
    with open("input.txt", "r") as input_file:
        lines = input_file.readlines()
        index = 0
        numbers_drawn = lines[0].strip().split(",")
        print(numbers_drawn)
        del lines[:2]

    for row in lines:
        row = [int(s) for s in row.split() if s.isdigit()]
        if not row and index in boards:
            existing_numbers = boards.get(index)
            boards[index] = existing_numbers + row
            index+=1

        elif index not in boards:
            boards[index] = row
        
        elif index in boards and row != []:
            existing_numbers = boards.get(index)
            boards[index] = existing_numbers + row

    return numbers_drawn


def check_boards(numbers_drawn): 
    test_list = [int(i) for i in numbers_drawn]
    for bingo_number in test_list:
        for k, v in boards.items():
            if bingo_number in v:
                for i in v:
                    if bingo_number == i:
                        
                        boards[k]=v

numbers_drawn = create_boards()
check_boards(numbers_drawn)
print(boards)

# print("Numbers drawn" + numbers_drawn)