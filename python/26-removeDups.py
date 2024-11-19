nums = [1, 1, 2, 3, 4, 5, 6, 6]


def RemoveDups(nums):
    dupcount = 0
    for index in range(0, len(nums) - 2, 1):
        print(index)
        value = nums[index]
        nextvalue = nums[index + 1]
        print(value, nextvalue)
        if value == nextvalue:
            nums.pop(value)
            nums.append(value)
            print(f"removing {value}")
            dupcount += 1

    return dupcount


print(RemoveDups(nums))
print(nums)
