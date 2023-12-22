from datetime import datetime
from enum import Enum

class LOGTYPE(Enum):
  TRACE = 'TRACE',
  DEBUG = 'DEBUG',
  INFO = 'INFO',
  WARN = 'WARN',
  ERROR = 'ERROR',
  FATAL = 'FATAL',

class Log:
  def __init__(self, timestamp=0, logType=None, origin='', info='', session=0):
    self.timestamp = datetime.timestamp(datetime.now()) if not timestamp else timestamp
    self.logType = logType
    self.origin = origin
    self.info = info
    self.session = 0

  def parse_log_type(log_type_str) -> LOGTYPE or ValueError:
    try:
        return LOGTYPE[log_type_str.upper()]
    except KeyError:
        raise ValueError(f"Invalid log type: {log_type_str}")