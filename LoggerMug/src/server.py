from flask import Flask, request
from Log import Log
from repository import SQLiteDB

app = Flask(__name__)
db = SQLiteDB("loggermug.db")
db.createSchema()

def startServer():
  app.run(debug=True)

@app.route('/hello', methods=['GET'])
def hello():
  return '<h1> Hello World </h1>'

@app.route('/event', methods=['POST'])
def postEvent():
  try:
    origin = str(request.form['origin'])
    info = str(request.form['info'])
    logType = Log.parse_log_type(request.form['type'])
    session = int(request.form['session'])
    if session < 1:
      raise ValueError("Session can't be less than one")
  except KeyError or ValueError as e:
    return f"<p> Incorrect Parameters {e}</p>"
  try:
    myLog = Log(info=info, logType=logType, session=session, origin=origin)
    db.insterEvent(myLog)
    return f"<p> Successfully logged event </p>"
  except Exception as e:
    return f"<p> {e} </p>"

@app.route('/getLogs', methods=['GET'])
def getEvents():
  try:
    page = int(request.args.get('page'))
  except KeyError or ValueError as e:
    return '<p> Invalid parameters </p>'

  queryTuples = db.getLastNEvents(10 * page, 10).split("), (")
  queryTuples[0] = queryTuples[0][2:]
  queryTuples[-1]= queryTuples[-1][:len(queryTuples[-1])-2]

  resp = ['<ol>']
  for log in queryTuples:
    resp.append(f"<li>{log}</li>")
  resp.append('</ol>')

  return '\n'.join(resp)