

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

@app.route('/create-project', methods=['POST'])
def createProject():
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
  {'id':1,'name':'Jira Clone', 'key':'SE', 'type':'Software', 'lead': 'Sample Person 1'},
  {'id':2,'name':'Corona Virus Tracker', 'key':'DBMS', 'type':'Software', 'lead': 'Sample Person 2'},
  {'id':3,'name':'Twitter Clone', 'key':'DOS', 'type':'Software', 'lead': 'Sample Person 3'},
  {'id':4,'name':'Bit Torrent Simulation', 'key':'CN', 'type':'Distributed Software', 'lead': 'Sample Person 4'},
  {'id':5,'name':'Gossip Protocol Implementation', 'key':'DOS', 'type':'Distributed Software', 'lead': 'Sample Person 5'},
]
	return Response(json.dumps({'response':data}), 200)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5000, debug = True)