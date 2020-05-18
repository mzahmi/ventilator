import os
import sys
import time
from PySide2 import QtCore, QtWidgets, QtQml


class Manager(QtCore.QObject):
    dataReady = QtCore.Signal(QtCore.QPointF, name='dataReady')

    def __init__(self, parent=None):
        super(Manager, self).__init__(parent)
        self._currX = 0
        self._currY = 0
        self._delay = 0.5
        self._multiplier = 1.0
        self._power = 1.0
        self._xIncrement = 1.0
        self._starter = True
        self._goOn = False
        self._threader = None

    @QtCore.Property(bool)
    def starter(self):
        return self._starter

    @starter.setter
    def setStarter(self, val):
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

    @multiplier.setter
    def setMultiplier(self, val):
        if self._multiplier == val:
            return
        print(val)
        self._multiplier = val

    @QtCore.Property(int)
    def power(self):
        return self._power

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
        self._currX += self._xIncrement
        self._currY = self._multiplier*(self._currX**self._power)

        return self._currX, self._currY

    def stop(self):
        self._goOn = False
        if self._threader is not None:
            while self._threader.isRunning():
                time.sleep(0.1)

    def start(self):
        self._goOn = True
        self._threader = Threader(self.core, self)
        self._threader.start()

    def core(self):
        while self._goOn:
            p = QtCore.QPointF(*self.generatePoint())
            self.dataReady.emit(p)
            time.sleep(self._delay)

# -------------------------------------------------


class Threader(QtCore.QThread):
    def __init__(self, core, parent=None):
        super(Threader, self).__init__(parent)
        self._core = core

    def run(self):
        self._core()


if __name__ == "__main__":
    os.environ["QT_QUICK_CONTROLS_STYLE"] = "Material"
    app = QtWidgets.QApplication(sys.argv)
    manager = Manager()
    app.aboutToQuit.connect(manager.stop)
    manager.start()
    engine = QtQml.QQmlApplicationEngine()
    ctx = engine.rootContext()
    ctx.setContextProperty("Manager", manager)
    engine.load('test.qml')
    if not engine.rootObjects():
        sys.exit(-1)

    sys.exit(app.exec_())
