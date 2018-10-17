import os
import sys
from subprocess import call


def rename_file(s):
    fs = s.split('_')
    return '/tmp/' + fs[0] + '.' + '-'.join(fs[1:])


def trim(s):
    s = s.strip()
    s = s.replace('import', '')
    s = s.replace(' ', '')
    s = s.replace('\t', '')
    return s


def split_imports_contents(files):
    contents = []
    imports = {}
    for f in files:
        start_import = False
        for i in f.split('\n'):
            if i.startswith('package'):
                continue
            if i.startswith('import '):
                if '"' in i:
                    imports[trim(i)] = True  # import "fmt"
                else:
                    start_import = True  # import (
                continue
            if start_import and ')' in i:
                start_import = False  # import ... )
                continue
            if start_import:
                imports[trim(i)] = True
                continue

            contents.append(i.replace('Interval','Interval_').replace('ListNode','ListNode_').replace('TreeNode','TreeNode_'))

    x = []
    for k, v in imports.items():
        if '/lib' not in k:
            x.append(k)

    return x, contents


def read_libs():
    fs = os.listdir('../lib')
    files = []
    for v in fs:
        if v.endswith('_test.go'):
            continue
        if not v.endswith('.go'):
            continue

        f = open('../lib/' + v).read()
        files.append(f)

    return files


def fix_content(s):
    libs = []
    if 'lib.' in s:
        libs = read_libs()
    libs.append(s.replace('lib.TreeNode', '12345678901234567890-1'))

    imports, contents = split_imports_contents(libs)

    new = []
    if len(imports) > 0:
        new.append('import (')
    for i in imports:
        new.append(i)
    if len(imports) > 0:
        new.append(')')
    for i in contents:
        new.append(i)

    return '\n'.join(new).replace('12345678901234567890-1', 'TreeNode')


def read_file(file):
    with open(file) as f:
        return f.read()


def write_file(file, content):
    with open(file, 'w') as f:
        f.write(content)


if __name__ == '__main__':
    argv = sys.argv
    if len(argv) < 2:
        print('请输入参数')
        exit(1)

    filename = argv[1]
    newname = rename_file(filename)

    content = read_file(filename)

    write_file(newname, fix_content(content))

    call(["leetcode", "submit", newname])
