# -*- coding: utf-8 -*-
"""
    package.module
    ~~~~~~~~~~~~~~

    A brief description goes here.

    :copyright: (c) YEAR by zwhset.
    :license: GOMEOPS, see LICENSE_FILE for more details.
"""

import sys

def main():
    counts = {}
    args = sys.argv[1:]
    if not args:
        print('please use filename')
        sys.exit(1)
    for filename in args:
        countLines(filename, counts)

    for k, v in counts.items():
        if v > 1:
            print(v, '\t', k)

def countLines(f, counts):
    with open(f) as fd:
        for line in fd:
            key = line.rstrip()[1:4]
            if counts.get(key, 0):
                counts[key] += 1
            else:
                counts[key] = 1

main()