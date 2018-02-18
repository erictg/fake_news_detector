import urllib.request

fp = urllib.request.urlopen("https://en.wikipedia.org/wiki/Michel_Houellebecq")
mybytes = fp.read()

mystr = mybytes.decode("utf8")
fp.close()

print(mystr)