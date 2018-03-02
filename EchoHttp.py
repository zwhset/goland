'''gunicorn -k gevent -w 4 -b 127.0.0.1:7070 EchoHttp:app'''

from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def index():
    return jsonify(dict(code=1, message='successed'))

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=7070)