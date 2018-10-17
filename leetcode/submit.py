import sys
from subprocess import call

def rename_file(s):
    fs = s.split('_')
    return '/tmp/' + fs[0] + '.' + '-'.join(fs[1:])


def fix_content(s):
    # package
    s = s.lstrip('package leetcode')

    return s


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
