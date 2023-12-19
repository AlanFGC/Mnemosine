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
  return db.getLastNEvents(32, 2)


