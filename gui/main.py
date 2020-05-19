# https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-7
import argparse
import json
import os
import random
import sys

from PySide2 import QtCore, QtQml, QtWidgets

import config
import mode_select as ms
from alarm_manager import AlarmManager
from chart_manager1 import ChartManager1
from chart_manager2 import ChartManager2
from chart_manager3 import ChartManager3
from config import logging as logging
from input_manager import UserInput

# from chart_manager import ChartManager as cm

# from patient import Patient


parser = argparse.ArgumentParser(description='Run the main GUI code')
parser.add_argument('-r', '--redis', action='store_true',
                    help="run GUI and send information to redis")
parser.add_argument('-f', '--fullscreen', action='store_true',
                    help="run GUI in full screen, dont have a way to kill it using touch yet")
config.args = parser.parse_args()

if config.args.redis:
    config.useredis = True

if config.args.fullscreen:
    config.fullscreen = True


def main():
    os.environ["QT_IM_MODULE"] = "qtvirtualkeyboard"

    app = QtWidgets.QApplication(sys.argv)

    chartManager1 = ChartManager1()
    chartManager2 = ChartManager2()
    chartManager3 = ChartManager3()

    app.aboutToQuit.connect(chartManager1.stop)
    app.aboutToQuit.connect(chartManager2.stop)
    app.aboutToQuit.connect(chartManager3.stop)

    chartManager1.start()

    alarmManager = AlarmManager()
    alarmManager.start()
    # patient = Patient()
    userInput = UserInput()
    modeSelect = ms.ModeSelect()
    dp = 0

    engine = QtQml.QQmlApplicationEngine()
    ctx = engine.rootContext()

    ctx.setContextProperty("ChartManager1", chartManager1)
    ctx.setContextProperty("ChartManager2", chartManager2)
    ctx.setContextProperty("ChartManager3", chartManager3)

    ctx.setContextProperty("ModeSelect", modeSelect)
    # ctx.setContextProperty("Patient", patient)
    ctx.setContextProperty("UserInput", userInput)
    ctx.setContextProperty("AlarmManager", alarmManager)

    ctx.setContextProperty("dp", dp)
    ctx.setContextProperty("fs", False)

    # if redis exists take the userinput
    if config.useredis:
        params = config.r.get("PARAMS")
        params = json.loads(params)
        ctx.setContextProperty("Params", params)

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
