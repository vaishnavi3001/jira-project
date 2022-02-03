

from crypt import methods
from sre_constants import SUCCESS
from urllib import response
from flask import Flask, render_template, request, Response, json
from werkzeug.utils import secure_filename
from flask_cors import CORS, cross_origin


BUCKET_NAME='nau-upload-page'

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'
app.config['CORS_HEADERS'] = 'Allow'

@app.route('/update-project', methods=['POST'])
def updatedetails():
	if request.data:
		# request_data = request.get_json()
		return Response({'status': 'SUCCESS'}, status=200)
       


@app.route('/project')
def get_project():
    # if request.data:
    #     request_data = request.get_json()
    #     if  request_data['username'] == 'nau' and request_data['passwd'] == 'nau':
    #         return Response(status=200)
    #     else:
    #         return Response(status=401)
	data = [
  {'position': 1, 'name': 'Hydrogen', 'weight': 1.0079, 'symbol': 'H'},
  {'position': 2, 'name': 'Helium', 'weight': 4.0026, 'symbol': 'He'},
  {'position': 3, 'name': 'Lithium', 'weight': 6.941, 'symbol': 'Li'},
  {'position': 4, 'name': 'Beryllium', 'weight': 9.0122, 'symbol': 'Be'},
  {'position': 5, 'name': 'Boron', 'weight': 10.811, 'symbol': 'B'},
  {'position': 6, 'name': 'Carbon', 'weight': 12.0107, 'symbol': 'C'},
  {'position': 7, 'name': 'Nitrogen', 'weight': 14.0067, 'symbol': 'N'},
  {'position': 8, 'name': 'Oxygen', 'weight': 15.9994, 'symbol': 'O'},
  {'position': 9, 'name': 'Fluorine', 'weight': 18.9984, 'symbol': 'F'},
  {'position': 10, 'name': 'Neon', 'weight': 20.1797, 'symbol': 'Ne'},
]
	return Response(json.dumps({'response':data}), 200)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5000, debug = True)