from flask import Flask, request
import json
from k import means

app = Flask(__name__)


@app.route('/rest/overall', methods=["POST"])
def index():
    print(request.data)
    data = request.data
    print(data)
    ddict = json.loads(data)

    ratio = ddict['ratio']
    sentiment = ddict['sentiment']
    res = means.compute(ratio, sentiment)

    data = {'truthiness': str(res)}
    return json.dumps(data)


if __name__ == '__main__':
    app.run(debug=True, port=8003, host='0.0.0.0')

