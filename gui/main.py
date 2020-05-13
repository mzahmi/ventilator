# https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-7
import os
import sys
import time
from PySide2 import QtCore, QtWidgets, QtQml
import redis
import mode_select as ms
from patient import Patient
import logging
import test
import argparse
logging.basicConfig(filename='log.log',format='%(levelname)s:%(asctime)s - %(message)s', datefmt='%H:%M:%S',level=logging.NOTSET)


parser = argparse.ArgumentParser(description='Run the main GUI code')
parser.add_argument('-r', '--redis', action='store_true', help="run GUI and send information to redis")
parser.add_argument('-f', '--fullscreen', action='store_true', help="run GUI in full screen, dont have a way to kill it using touch yet")
args = parser.parse_args()

if args.redis:
    r = redis.StrictRedis(
        host='localhost',
        port=6379,
        password='',
        decode_responses=True)




class Manager(QtCore.QObject):
    # create a signal
    dataReady = QtCore.Signal(QtCore.QPointF, name='dataReady')
    

    # initial values
    def __init__(self, parent=None):
        # if 'parent' is given then it will inherit it
        super(Manager, self).__init__(parent)
        self._currX = 0
        self._currY = 0
        self._delay = 0.2
        self._multiplier = 1.0
        self._power = 1.0
        self._xIncrement = 1.0
        self._starter = False
        self._goOn = False
        self._threader = None

    # property 'starter' can be seen in qml
    # connected to the button start
    @QtCore.Property(bool)
    def starter(self):
        return self._starter

    # set the 'starter' property
    @starter.setter
    def setStarter(self, val):
        # val is returned from qml
        if self._multiplier == val:
            return
        print(val)
        if val:
            self.start()
        else:
            self.stop()
        self._starter = val

    
    @QtCore.Property(float)
    def multiplier(self):
        return self._multiplier

    # 'multiplier' can be set from qml
    # being set when slider 'multiplierSlider' is changed
    @multiplier.setter
    def setMultiplier(self, val):
        if self._multiplier == val:
            return
        print(val)
        self._multiplier = val

    # makes 'power'
    @QtCore.Property(int)
    def power(self):
        return self._power

    # sets power 
    @power.setter
    def setPower(self, val):
        if self._power == val:
            return
        print(val)
        self._power = val

    @QtCore.Property(float)
    def delay(self):
        return self._delay

    @delay.setter
    def setDelay(self, val):
        if self._delay == val:
            return
        print(val)
        self._delay = val

    @QtCore.Property(float)
    def xIncrement(self):
        return self._xIncrement

    @xIncrement.setter
    def setXIncrement(self, val):
        if self._xIncrement == val:
            return
        print(val)
        self._xIncrement = val

    def generatePoint(self):
        # increments and returns x and y

        self._currX += self._xIncrement
        if not args.redis:
            self._currY = self._currY+3
        else:
            self._currY = float(r.get("pressure"))

        return self._currX,self._currY

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
            # makes an XY point object using generatepoint
            # using 'self._currX,self._currY'
            p = QtCore.QPointF(*self.generatePoint())
            # sends signal and then waits for delay
            self.dataReady.emit(p)
            time.sleep(self._delay)

# ------------------------------------------------- 

class Threader(QtCore.QThread):
    def __init__(self,core,parent=None):
        super(Threader, self).__init__(parent)
        self._core = core

    def run(self):
        self._core()

def main():
    os.environ["QT_QUICK_CONTROLS_STYLE"] = "Material"
    app = QtWidgets.QApplication(sys.argv)

    manager = Manager()
    patient = Patient()
    modeSelect = ms.ModeSelect()
    dp = 0


    app.aboutToQuit.connect(manager.stop)
    manager.start()
    engine = QtQml.QQmlApplicationEngine()

    ctx = engine.rootContext()
    ctx.setContextProperty("Manager", manager)
    ctx.setContextProperty("ModeSelect", modeSelect)
    ctx.setContextProperty("Patient", patient)
    ctx.setContextProperty("dp", dp)
    if args.fullscreen:
        logging.debug("Runnin in full screen")
        ctx.setContextProperty("fs", True)



    # engine.load('main.qml')
    engine.load('./qml/MainQt.qml')
    if not engine.rootObjects():
        sys.exit(-1)

    sys.exit(app.exec_())
if __name__ == "__main__":
    main()
