
import random
import time
# from main import args

import config
from PySide2 import QtCore, QtQml, QtWidgets

r = config.r


class AlarmManager(QtCore.QObject):
    # create a signal
    alarmStatus = QtCore.Signal(str, name='alarmStatus')

    # initial values

    def __init__(self, parent=None, r=None):
        # if 'parent' is given then it will inherit it
        super(AlarmManager, self).__init__(parent)
        self._status = None
        self._title = None
        self._info = None
        self._delay = 0.5
        self._starter = False
        self._goOn = False
        self._threader = None

    @QtCore.Property(bool)
    def status(self):
        return self._status

    # set the 'starter' property
    @status.setter
    def setStatus(self, val):
        self._status = val

    @QtCore.Property(bool)
    def title(self):
        return self._title

    # set the 'starter' property
    @title.setter
    def setTitle(self, val):
        self._title = val

    @QtCore.Property(bool)
    def info(self):
        return self._info

    # set the 'starter' property
    @info.setter
    def setinfo(self, val):
        self._info = val

    def stop(self):
        self._goOn = False
        # checks threader, if still alive, stays inside till dead
        if self._threader is not None:
            while self._threader.isRunning():
                time.sleep(0.1)

    def start(self):
        self._goOn = True
        self._threader = Threader(self.core, self)
        self._threader.start()

    def core(self):
        while self._goOn:
            if config.useredis:
                self._status = config.r.get("alarm_status")
                self._title = config.r.get("alarm_title")
                self._info = config.r.get("alarm_info")
                print("emitting alarm ", self._info)
                self.alarmStatus.emit(config.r.get("alarm_info"))
                time.sleep(self._delay)

# -------------------------------------------------


class Threader(QtCore.QThread):
    def __init__(self, core, parent=None):
        super(Threader, self).__init__(parent)
        self._core = core

    def run(self):
        self._core()
