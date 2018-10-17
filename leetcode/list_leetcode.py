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

            if 'Easy' in a \
                    and not a.startswith('[141]') \
                    and not a.startswith('[190]') \
                    and not a.startswith('[191]') \
                    and not a.startswith('[235]') \
                    and not a.startswith('[237]') \
                    and not a.startswith('[160]'):
                print(a)


if __name__ == '__main__':
    get_leetcode_list()
