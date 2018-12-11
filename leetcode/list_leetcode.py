#!/usr/local/bin/python3
# encoding: utf8

import subprocess


def get_leetcode_list():
    outout = subprocess.check_output(['leetcode', 'list', '-x'])
    ls = outout.decode("utf-8").split('\n')

    i = 0
    for i in range(len(ls)):
        if i % 2 == 0 and i + 1 < len(ls):
            a = ls[i].strip()
            b = ls[i + 1].strip()

            if not b.startswith('algorithms'):
                continue
            if a.startswith('✔'):
                continue
            if a.startswith('🔒'):
                continue
            if a.startswith('✘'):
                continue

            skip = False
            for i in [141, 190, 191, 235, 237, 160, 278]:
                if a.startswith('[{}]'.format(i)):
                    skip = True
                    break
            if skip:
                continue

            if 'Easy' in a:
                print(a)


if __name__ == '__main__':
    get_leetcode_list()
