import yaml

from PySide2 import QtCore, QtWidgets, QtQml

class Patient(QtCore.QObject):
    def __init__(self, parent=None):
        super(Patient, self).__init__(parent)
        self._name = None
        self._age = None
        self._gender = None
        self._height = None
        self._weight = None
        self.initPatient()

    def initPatient(self):
        with open(r'patient.yaml') as file:
            # The FullLoader parameter handles the conversion from YAML
            # scalar values to Python the dictionary format
            patient_info = yaml.load(file, Loader=yaml.FullLoader)

            self._name = patient_info["name"]
            self._age = patient_info["age"]
            self._gender = patient_info["gender"]
            self._height = patient_info["height"]
            self._weight = patient_info["weight"]

    @QtCore.Property(str, constant=True)
    def name(self):
        return self._name

    @name.setter
    def setName(self, val):
        self._name = val

    @QtCore.Property(int, constant=True)
    def age(self):
        return self._age

    @age.setter
    def setAge(self, val):
        self._age = val

    @QtCore.Property(str, constant=True)
    def gender(self):
        return self._gender

    @gender.setter
    def setGender(self, val):
        self._gender = val

    @QtCore.Property(int, constant=True)
    def height(self):
        return self._height

    @gender.setter
    def setHeight(self, val):
        self._height = val

    @QtCore.Property(int, constant=True)
    def weight(self):
        return self._weight

    @gender.setter
    def setWeight(self, val):
        self._weight = val

    

if __name__ == "__main__":
    patient = Patient()
    patient.height=60
