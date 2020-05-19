from PySide2 import QtCore, QtWidgets, QtQml
import json
import redis
import config
from main import logging
import time
import json

if config.useredis:
    logging.info("using redis")
    useredis = True
else:
    logging.info("not using redis")
    useredis = False


InsparotaryPressure = {
    "name": "Insparotary Pressure",
    "initialVal": 25,
    "minVal": 15,
    "maxVal": 40,
    "stepSize": 5,
}

BreathPerMinute = {
    "name": "Breath Per Minute",
    "initialVal": 20,
    "minVal": 8,
    "maxVal": 40,
    "stepSize": 2,
}

PMAX = {
    "name": "PMAX",
    "initialVal": 20,
    "minVal": 0,
    "maxVal": 40,
    "stepSize": 5,
}

PEEP = {
    "name": "PEEP",
    "initialVal": 10,
    "minVal": 5,
    "maxVal": 20,
    "stepSize": 5,
}

VT = {
    "name": "Tidal Volume",
    "initialVal": 200,
    "minVal": 100,
    "maxVal": 800,
    "stepSize": 10,
}


FIO2 = {
    "name": "FIO2%",
    "initialVal": 60,
    "minVal": 21,
    "maxVal": 100,
    "stepSize": 5,
}


class UserInput(QtCore.QObject):
    def __init__(self, parent=None):
        super(UserInput, self).__init__(parent)
        self._PEEP = PEEP
        self._InsparotaryPressure = InsparotaryPressure
        self._FIO2 = FIO2
        self._PMAX = PMAX
        self._VT = VT
        self._BreathPerMinute = BreathPerMinute

        self._mode = None
        self._trigger = None

    @QtCore.Property(str)
    def mode(self):
        return self._mode

    @mode.setter
    def setMode(self, mode):
        self._mode = mode
        if mode == "":
            logging.debug("mode reset")
        else:
            logging.debug(f'mode set: {mode}')

    @QtCore.Property(str)
    def trigger(self):
        return self._trigger

    @trigger.setter
    def setTrigger(self, trigger):
        self._trigger = trigger
        if trigger == "":
            logging.debug("trigger reset")
        else:
            logging.debug(f'trigger type set: {trigger}')

    @QtCore.Property('QVariant', constant=True)
    def InsparotaryPressure(self):
        return self._InsparotaryPressure

    @QtCore.Property('QVariant', constant=True)
    def PEEP(self):
        return self._PEEP

    @QtCore.Property('QVariant', constant=True)
    def FIO2(self):
        return self._FIO2

    @QtCore.Property('QVariant', constant=True)
    def BreathPerMinute(self):
        return self._BreathPerMinute

    @QtCore.Property('QVariant', constant=True)
    def VT(self):
        return self._VT
