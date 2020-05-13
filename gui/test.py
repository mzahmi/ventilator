"""
https://stackoverflow.com/questions/57619227/connect-qml-signal-to-pyside2-slot
https://stackoverflow.com/questions/54010254/connect-python-signal-to-qml-ui-slot-with-pyside2

"""

from PySide2 import QtCore, QtWidgets, QtQml


class UserInput(QtCore.QObject):
    def __init__(self,startval=42, parent=None):
        super(UserInput, self).__init__(parent)
        self.ppval = startval

    def readPP(self):
        return self.ppval

    def setPP(self,val):
        self.ppval = val

        UserInput.pp = QtCore.Property(int, UserInput.readPP, UserInput.setPP)

if __name__ == "__main__":
    ui = UserInput()
    print(ui.ppval)
