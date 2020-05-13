# https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-7
import argparse
import os
import random
import sys
import test
import time

from alarm_manager import AlarmManager

from PySide2 import QtCore, QtQml, QtWidgets

# from chart_manager import ChartManager as cm
import config
import mode_select as ms
from config import logging as logging
from patient import Patient
from chart_manager import ChartManager


parser = argparse.ArgumentParser(description='Run the main GUI code')
parser.add_argument('-r', '--redis', action='store_true',
                    help="run GUI and send information to redis")
parser.add_argument('-f', '--fullscreen', action='store_true',
                    help="run GUI in full screen, dont have a way to kill it using touch yet")
config.args = parser.parse_args()

r = config.r

if config.args.redis:
    r = config.r
    config.useredis = True

if config.args.fullscreen:
    config.fullscreen = True


def main():
    os.environ["QT_QUICK_CONTROLS_STYLE"] = "Material"
    app = QtWidgets.QApplication(sys.argv)

    chartManager = ChartManager()
    alarmManager = AlarmManager()
    patient = Patient()
    modeSelect = ms.ModeSelect()
    dp = 0

    app.aboutToQuit.connect(chartManager.stop)
    chartManager.start()
    alarmManager.start()
    engine = QtQml.QQmlApplicationEngine()

    ctx = engine.rootContext()
    ctx.setContextProperty("ChartManager", chartManager)
    ctx.setContextProperty("ModeSelect", modeSelect)
    ctx.setContextProperty("Patient", patient)
    ctx.setContextProperty("AlarmManager", alarmManager)
    ctx.setContextProperty("dp", dp)
    if config.args.fullscreen:
        logging.debug("Runnin in full screen")
        ctx.setContextProperty("fs", True)

    # engine.load('main.qml')
    engine.load('./qml/MainQt.qml')
    if not engine.rootObjects():
        sys.exit(-1)

    sys.exit(app.exec_())


if __name__ == "__main__":
    main()
