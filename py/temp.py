import sys
import re
from datetime import datetime
from pymongo import Connection                  # pymongo.connection
from pymongo.errors import ConnectionFailure    # pymongo.errors.ConnectionFailure

def main():
    try:
        c = Connection(host="localhost", port=27017)
    except ConnectionFailure,e: 
        sys.stderr.write("Could not connect to MongoDB: %s" % e)
        sys.exit(1)
    dbh = c["profile"]
    assert dbh.connection == c
    profile_docs = dbh.profile.find({"riqi":re.compile("^201402.*")})  # search input
    if not profile_docs:
        print 'no data found'
    f = open('out.txt','w')
    for profile_doc in profile_docs:
        f.write(profile_doc['riqi'])
        f.write(profile_doc['title'].encode('utf-8'))
        f.write(profile_doc['tag'].encode('utf-8'))
        f.write('\n')
        f.write(profile_doc['shijian'].encode('utf-8'))
        f.write('\n')
        f.write(profile_doc['beizhu'].encode('utf-8'))
        f.write('\n=====\n\n')
    f.close()
    print "search is done. Output into out.txt"

if __name__ == "__main__":
        main()

