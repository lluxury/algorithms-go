# encoding: utf8

import os
import subprocess
import sys


def exist_question(number):
    files = os.listdir('./')

    for file in files:
        if file.startswith(number + '_') and file.endswith('.go') and not file.endswith('_test.go'):
            return True
    return False


def find_generate_file(number, rec=True):
    files = os.listdir('./')

    generate_file = ''
    for file in files:
        if file.startswith(number + '.'):
            generate_file = file

    if not rec:
        return generate_file

    if generate_file == '':
        command = ['leetcode', 'show', number, '-gx', '-l', 'golang']
        p = subprocess.Popen(command)
        p.wait()

    return find_generate_file(number, False)


def find_test(number):
    files = os.listdir('./')

    for file in files:
        if file.startswith(number + '_') and file.endswith('_test.go'):
            return file
    return ''


if __name__ == '__main__':
    argv = sys.argv
    if len(argv) < 2:
        print('请输入参数')
        exit(1)

    question_number = argv[1]

    exist = exist_question(question_number)
    if exist:
        print('exist')
        sys.exit(0)

    file = find_generate_file(question_number)

    newname = file.replace('-', '_')
    newname = newname.replace(question_number + '.', question_number + '_')

    os.rename(file, newname)

    content = ['package leetcode', '\n', '\n']
    with open(newname, 'r') as f:
        for i in f:
            t = i.replace('*ListNode', '*lib.ListNode')
            content.append(t)

    with open(newname, 'w') as f:
        f.write(''.join(content))

    testfile = find_test(question_number)
    if testfile=='':
        with open(question_number+'_test.go', 'w') as f:
            f.write('''package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_{}(t *testing.T) {{
	test.Runs(t, testfunction, []*test.Case{{
	    {{Input: ``, Output: ``}},
	}})
}}

'''.format(question_number))
    elif testfile != '' and testfile != question_number + '_test.go':
        os.rename(testfile, question_number + '_test.go')
