import requests
import sys

def main():
    urls = sys.argv[1:]
    for url in urls:
        r = requests.get(url)
        print r.text

main()