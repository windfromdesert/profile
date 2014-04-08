def main():
    f = open('md','w')
    f.write('RiQi\n')
    f.write('Title\n')
    f.write('Tag\n')
    f.write('BeiZhu\n\n')
    f.write('ShiJian\n')
    f.close()
    ff = open('out.txt','w')
    ff.write('')
    ff.close()
    print('md 和 out.txt 文件已经重置')

if __name__ == "__main__":
    main()
