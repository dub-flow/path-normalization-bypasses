from flask import Flask, request

app = Flask(__name__)

@app.route('/admin')
def admin():
    return "Admin area accessed!"

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
    return f"Accessed path: {path}"

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)
