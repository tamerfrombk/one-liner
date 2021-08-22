#!/usr/bin/env python

import sys

def seed(n):
    with open("lorem-ipsum.txt", "r") as r:
        contents = r.read()
        with open("large-file.txt", "w") as w:
            for i in range(n):
                w.write(contents)


def parse_iteration_count(argv):
    argc = len(argv)

    if argc == 0:
        return 10000
    
    if argc > 1:
        raise ValueError("only one valid integer argument allowed")
    
    try:
        return int(argv[0])
    except:
        raise ValueError("{} is not a valid number".format(argv[0]))


def main():
    try:
        n = parse_iteration_count(sys.argv[1:]) 

        seed(n)
    except Exception as e:
        print(e)
        exit(1)


if __name__ == '__main__':
    main()

