import sqlite3
import pypika
from Log import LOGTYPE, Log


LOG_TABLE = 'apilogger'
ID_COL = 'id'
LOGTYPE_COL = 'logType'
TIMESTAMP_COL = 'timestamp'
INFO_COL = 'info'
SESSION_COL = 'session'
ORIGIN_COL = 'origin'

class SQLiteDB():
  def __init__(self, url: str) -> None:
    self.connection = sqlite3.connect(url, check_same_thread=False)

  def createSchema(self):
    cursor = self.connection.cursor()
    try:
      cursor.execute(f"""
            CREATE TABLE {LOG_TABLE} (
                {ID_COL} INTEGER PRIMARY KEY,
                {ORIGIN_COL} VARCHAR(255),
                {LOGTYPE_COL} VARCHAR(5) CHECK(logType IN
                ('{LOGTYPE.TRACE.value[0]}', '{LOGTYPE.DEBUG.value[0]}', '{LOGTYPE.INFO.value[0]}', '{LOGTYPE.WARN.value[0]}', '{LOGTYPE.ERROR.value[0]}',
                '{LOGTYPE.FATAL.value[0]}')),
                {TIMESTAMP_COL} DATETIME DEFAULT CURRENT_TIMESTAMP,
                {INFO_COL} VARCHAR(255),
                {SESSION_COL} MEDIUMINT DEFAULT 1
            )
        """)
      print(f"{LOG_TABLE} successfully created.")
    except Exception as e:
      print(f"FAILED to create log table: f{e}")

  def insterEvent(self, log: Log):
    try:
      cursor = self.connection.cursor()
      q = pypika.Query.into(LOG_TABLE)\
        .columns(ORIGIN_COL, LOGTYPE_COL, TIMESTAMP_COL, INFO_COL, SESSION_COL)\
        .insert(log.origin, log.logType.value[0], log.timestamp, log.info, log.session)
      cursor.execute(str(q))
      self.connection.commit()
    except Exception as e:
      print(f"Failed to insert: {e}")
      cursor.close()
      raise ValueError(f"Incorrect parameters: {e}")

    cursor.close()

  def getLastNEvents(self, offset, limit):
    try:
        cursor = self.connection.cursor()
        q = pypika.Query.from_(LOG_TABLE).select(ID_COL, ORIGIN_COL, LOGTYPE_COL, TIMESTAMP_COL, INFO_COL, SESSION_COL).offset(offset).limit(limit)
        c = cursor.execute(str(q))
        res = []
        for row in c:
            res.append(row)
        return str(res)
    except Exception as e:
        print(f"Failed to get last events: {e}")
    finally:
        cursor.close()

  def getEventsFromClientName(self, host):
    pass

  def getEventFromClientIP(self, IP):
    pass

  def ___del__(self):
    self.connection.close()