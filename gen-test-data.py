#!/usr/bin/env python

import os

def test_file_path(name):
    return os.path.join(os.getcwd(), "test-data", name)

def seed():
    fileToIterationCount = {
        'small.txt': 5,
        'medium.txt': 500,
        'large.txt': 50_000,
        'huge.txt': 500_000
    }

    with open(test_file_path("lorem-ipsum.txt"), "r") as r:
        contents = r.read()
        for [name, count] in fileToIterationCount.items():
            with open(test_file_path(name), "w") as w:
                w.write(contents * count)


def main():
    seed()


if __name__ == '__main__':
    main()

