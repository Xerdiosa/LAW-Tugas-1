from flask import Flask, render_template, request, render_template_string, make_response
import requests, os, json

app = Flask(__name__)

@app.route('/', methods=['GET'])
def index():
    print("\n\n\n\nmenerima request dari browser!", flush=True)

    a = request.args.get('a')
    b = request.args.get('b')

    rest_ip = os.environ.get("REST_IP")
    url = "http://" + rest_ip

    payload={'a': a,
            'b': b}
    
    print("mengirim request POST ke {}".format(rest_ip), flush=True)
    response = requests.request("POST", url, headers={}, data=payload, files=[])
    print("mendapat respon dari {}".format(rest_ip), flush=True)

    if response.status_code == 400:
        return "<html>invalid parameter</html>"

    parsed_response = json.loads(response.text)      

    return "<html>Hasil: {} </html>".format(parsed_response.get("hasil"))


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
