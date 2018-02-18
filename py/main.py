from flask import Flask, request
import json
from similarity import sim

app = Flask(__name__)


@app.route('/rest/matching', methods=["POST"])
def index():
    data = request.data
    ddict = json.loads(data)

    content = ddict['content']
    res = sim.analyze(content)

    data = {'total': res[0], 'score': res[1]}
    return json.dumps(data)


if __name__ == '__main__':
    app.run(debug=True, port=8002)
