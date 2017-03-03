import re
import sys

report_re = r'Bufsize(\s|\S)*? read'

buffersize_re=r'(?<=Bufsize: )(\d)+(?= )'
threads_re=r'(?<=Threads: )(\d)+(?= )'
sessiongs_re=r'(?<=Sessions: )(\d)+'
writen_bytes_re=r'(\d| )+(?= total bytes written)'
read_bytes_re=r'(\d| )+(?= total bytes read)'

if __name__== "__main__":
    if len(sys.argv) != 2:
        print('input error')
        sys.exit(0)

    print('|Buffer|Threads|Sessions|Writes|Read|')
    print('|-|-|-|-|-|')
    all_report = sys.argv[1]
    matches = re.finditer(report_re, all_report)
    for match in matches:
        one_report = match.group()

        write = int(int(re.search(writen_bytes_re,one_report).group())/1024/1024/10)
        read = int(int(re.search(read_bytes_re,one_report).group())/1024/1024/10)

        one_row = '|{Buffer} BYTES|{Threads}|{Sessions}|{Write} MB/s|{Read} MB/s|'.format(Buffer=re.search(buffersize_re,one_report).group(),
                    Threads=re.search(threads_re,one_report).group(),
                    Sessions=re.search(sessiongs_re,one_report).group(),
                    Write=write,
                    Read=read)
        print(one_row)
