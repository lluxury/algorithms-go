#!/usr/local/bin/python3
import subprocess

readme = '''# algorithms-go

[![Build Status](https://travis-ci.org/Chyroc/algorithms-go.svg?branch=master)](https://github.com/Chyroc/leetcode-badge)
[![codecov](https://codecov.io/gh/Chyroc/algorithms-go/branch/master/graph/badge.svg)](https://github.com/Chyroc/leetcode-badge)

[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&leetcode_badge_style=Leetcode%20|%20Ranking-{{.ranking}}-green.svg&refresh=new)](https://github.com/Chyroc/leetcode-badge)
[![leetcode badge](https://leetcode-badge.chyroc.cn/?name=chyroc&refresh=new)](https://github.com/Chyroc/leetcode-badge)

记录go算法题训练的地方
'''


class Question:
    right = ''
    lock = ''
    number = 0
    title = ''
    level = ''
    passed = ''

    def __str__(self):
        # return '<({})\t({})\t({})\t({})\t({})\t({})>'.format(self.right, self.lock, self.number, self.level,
        #                                                     self.passed, self.title)
        return '<({})\t({})\t({})\t({})\t({})\t({})>'.format(self.lock, self.right, self.number, self.title, self.level,
                                                             self.passed)


def _format_leetcode_line(line):
    q = Question()
    start = 0
    i = -1
    while i < len(line) - 1:
        i += 1

        char = line[i]
        print('i', i, 'start', start, 'char', char, char == '[', char == ']', char == '(', char == ')')

        if start <= 0:
            # if char == '✔':
            if char == '√':
                q.right = '✅'
                continue
            # elif char == '✘':
            elif char == '×':
                q.right = '❌'
                continue
            # elif char == '🔒':
            elif char == '$':
                # q.lock = '🔒'
                q.lock = '🔒'
                continue
            elif char == '[':
                start = 1
                continue
            else:
                continue

        if start <= 1:
            number = ''
            while True:
                if i < len(line) and line[i] == ']':
                    start = 2
                    break
                number += line[i]
                i += 1
            q.number = int(number)
            continue

        if start <= 2:
            title = ''
            title_start = False
            while True:
                if i + 1 < len(line) and line[i] == ' ' and line[i + 1] == ' ':
                    start = 3
                    break

                char = line[i]
                i += 1
                if title_start or char != ' ':
                    title += char
                    title_start = True
                elif char == ' ':
                    continue

            q.title = title
            continue

        if start <= 3:
            if char == ' ':
                continue
            start = 4
            continue

        if start <= 4:
            level = line[i - 1:i + 3]
            if level == 'Hard':
                q.level = '🔴🔴🔴'
                i += 4
            elif level == 'Easy':
                q.level = '🔴'
                i += 4
            else:
                q.level = '🔴🔴'
                i += 6
            start = 5
            continue

        if start <= 5:
            if char == ' ':
                continue
            elif char == '(':
                start = 6
                continue

        if start <= 6:
            q.passed = line[i:len(line) - 1]
            break

    return q


def get_leetcode_list():
    output = subprocess.check_output(['leetcode', 'list'])

    unlock_questions = []
    lock_questions = []
    for line in output.decode('utf-8').split('\n'):
        q = _format_leetcode_line(line)
        if q.number != 0:
            if q.lock == '':
                unlock_questions.append(q)
            else:
                lock_questions.append(q)
    unlock_questions.reverse()
    lock_questions.reverse()
    return unlock_questions, lock_questions
# 执行指定的命令，如果执行状态码为0则返回命令执行结果，否则抛出异常

def to_markdown(questions):
    l = []
    l.append('number | title | done | level | url | file')
    l.append('---  | --- | --- | --- | --- | ---')
    for q in questions:
        url = '[click](https://leetcode.com/problems/{}/description/)'.format(q.title.replace(' ', '-').lower())
        file = '[click](https://github.com/Chyroc/algorithms-go/blob/master/leetcode/{}_{}.go)'.format(q.number, q.title.replace(' ', '_').lower())
        l.append('{} | {} | {} | {} | {} | {}'.format(q.number, q.title, q.right, q.level, url, file))

    return '\n'.join(l)


if __name__ == '__main__':
    unlock_questions, lock_questions = get_leetcode_list()
    unlock_markdown = to_markdown(unlock_questions)
    lock_markdown = to_markdown(lock_questions)
    with open('README.md', 'w') as f:
        f.write(readme + '\n\n' + unlock_markdown + '\n\n## 🔒\n\n' + lock_markdown)
