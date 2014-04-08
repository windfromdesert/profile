f = open('md')
while 1:
    lines = f.readlines(10000)
    if not lines:
        break
    print lines
    print len(lines)
    print lines[3]
