from textblob import TextBlob
from textblob.taggers import NLTKTagger
import duckduckgo
import urllib2
from xml.etree import ElementTree


nltk_tagger = NLTKTagger()
#add this to docker 'python -m textblob.download_corpora'


def analyze(content):
    zen = TextBlob(content)
    pos = zen.pos_tags


def query(q):
    r = urllib2.urlopen('https://api.duckduckgo.com/?q=' + str(q).replace(" ", "+").lower() + '&t=hf&format=json&safesearch=false').read()
    print(r)
    xml = ElementTree.fromstring(r.read())
    r.close()
    return duckduckgo.Result(xml)

example = "i hate this"


query(example)
