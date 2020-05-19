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

if useredis:
    r = config.r


class ModeSelect(QtCore.QObject):
    # signal when mode is selected
    modeSelected = QtCore.Signal(str, name='modeSelected')
    # stop ventilation signal
    stopVent = QtCore.Signal(name='stopVent')

    def __init__(self, parent=None):
        super(ModeSelect, self).__init__(parent)
        self._currMode = ""
        self._currBreath = ""
        self._currTrigger = ""
        self._status = ""
        self._threader = None
        self._delay = 1
        self._goOn = True
        self._pip = {"name": "PIP", "value": 3}

    @QtCore.Property(str)
    def mode(self):
        return self._currMode

    @mode.setter
    def setMode(self, mode):
        self._currMode = mode
        if mode == "":
            logging.debug("mode reset")
        else:
            logging.debug(f'mode set: {mode}')

    @QtCore.Property('QVariant')
    def pip(self):
        return self._pip

    @pip.setter
    def setPip(self, val):
        print(val)
        self._pip = val

    @QtCore.Property(str)
    def breath(self):
        return self._currBreath

    @breath.setter
    def setBreath(self, breath):
        self._currBreath = breath
        if breath == "":
            logging.debug("breath reset")
        else:
            logging.debug(f'breath type set: {breath}')

    @QtCore.Property(str)
    def trigger(self):
        return self._currTrigger

    @trigger.setter
    def setTrigger(self, trigger):
        self._currTrigger = trigger
        if trigger == "":
            logging.debug("trigger reset")
        else:
            logging.debug(f'trigger type set: {trigger}')

    @QtCore.Property(str)
    def status(self):
        return self._status

    @status.setter
    def setStatus(self, status):
        logging.debug("status set: {}".format(status))
        self._status = status

    @QtCore.Slot()
    def stopVentilation(self):
        self._currMode = ""
        self._currTrigger = ""
        self._currBreath = ""
        logging.warning("Stopping Ventilation")
        # send status 'stop'
        if useredis:
            r.mset({"status": "stop"})
        self.stopVent.emit()

    @QtCore.Slot()
    def startVentilation(self):
        self.status = "start"
        logging.warning("Starting Ventilation")
        if useredis:
            r.mset({"status": self._status})

    @QtCore.Slot(str, str)
    def sendString(self, keystring, valstring):
        if useredis:
            params = r.get("PARAMS")
            params = json.loads(params)
            params[keystring] = valstring
            paramsdump = json.dumps(params)
            r.mset({"PARAMS": paramsdump})

    @QtCore.Slot(str, int)
    def sendInt(self, mystring, myint):

        if useredis:
            params = r.get("PARAMS")
            params = json.loads(params)

            params[mystring] = myint
            paramsdump = json.dumps(params)
            r.mset({"PARAMS": paramsdump})

        logging.debug(f'Input {mystring} set: {myint}')
        self.modeSelected.emit(self._currMode)

    def start(self):
        self._threader = Threader(self.core, self)
        self._threader.start()

    def core(self):
        while self._goOn:
            # sends signal and then waits for delay
            # print("on thread")
            time.sleep(self._delay)

# -------------------------------------------------


class Threader(QtCore.QThread):
    def __init__(self, core, parent=None):
        super(Threader, self).__init__(parent)
        self._core = core

    def run(self):
        self._core()


if __name__ == "__main__":
    for key in mode_breath:
        print(key)

    # choose mode
    mode = input("choose mode\n")

    # list of breath
    breath_list = mode_breath[mode]
    print("Breathe Types ", breath_list)

    chosen_breathe = input("Choose breathe\n")
    print("Trigger Types", breath_trigger[chosen_breathe])

    chosen_trigger = input("Choose trigger\n")
    print("Inputs ", trigger_input[chosen_trigger])
