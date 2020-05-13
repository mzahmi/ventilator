'use strict';
// margin from the left
var leftX = 50;
// number of total components
var componentNumber;
// height of each component
// should be the same for all
var componentHeight = 90
// basic slider shape
var baseSlider = Qt.createComponent("BaseSlider.qml");
// basic radio group
var baseRadioGroup = Qt.createComponent("BaseRadioGroup.qml");

// classes that hold default values for
// every input, changes should be made here
var InsparotaryPressure = { 
    name: "Insparotary Pressure",
    initialVal:25,
    minVal:15,
    maxVal:40,
    stepSize:5,
}

var BreathPerMinute = { 
    name: "Breath Per Minute",
    initialVal:20,
    minVal:8,
    maxVal:40,
    stepSize:2,
}

var PMAX = { 
    name: "PMAX",
    initialVal:20,
    minVal:0,
    maxVal:40,
    stepSize:5,
}

var PEEP = { 
    name: "PEEP",
    initialVal:10,
    minVal:5,
    maxVal:20,
    stepSize:5,
}


var FIO2 = { 
    name: "FIO2%",
    initialVal:60,
    minVal:21,
    maxVal:100,
    stepSize:5,
}


// create a slider component given the slider class
function createSliderComponent(componentInstance, componentNumber) {
    var createComp = baseSlider.createObject(flickableItems, {
        x:leftX,
        y:componentNumber*componentHeight,
        name: componentInstance.name,
        initialVal: componentInstance.initialVal,
        minVal: componentInstance.minVal,
        maxVal: componentInstance.maxVal,
        stepSize: componentInstance.stepSize,
    })
    
    return createComp
}

// create IE ratio radio group
// change this if more radio groups are needed
function createIERatio(componentNumber) {
    var createComp = baseRadioGroup.createObject(flickableItems, {
        x:80,
        y:componentNumber*componentHeight
    })

    return createComp
}

// adds all components to view given a list of inputs
// the list must be a list of strings with names matching the component names
function addToView(componentList){
    for (var i = 0; i < componentList.length; i++) {
        const element = componentList[i];
        switch (element) {
            case "Insparotary Pressure":
                var component = createSliderComponent(InsparotaryPressure, componentNumber)
                if (component === null) {
                    // Error Handling
                    console.log("Error creating object "+element);
                }
                componentNumber++;
                break;
            case "Breath Per Minute":
                var component = createSliderComponent(BreathPerMinute, componentNumber);
                componentNumber++;
                if (component === null) {
                    // Error Handling
                    console.log("Error creating object "+element);
                }
                break
                case "IE":
                    var component = createIERatio(componentNumber);
                    componentNumber++;
                    if (component === null) {
                        // Error Handling
                        console.log("Error creating object "+element);
                    }
                    break
            case "PMAX":
                var component = createSliderComponent(PMAX, componentNumber)
                componentNumber++;
                if (component === null) {
                    // Error Handling
                    console.log("Error creating object "+element);
                }
                break
            case "PEEP":
                var component = createSliderComponent(PEEP, componentNumber)
                componentNumber++;
                if (component === null) {
                    // Error Handling
                    console.log("Error creating object "+element);
                }
                break
            case "FIO2":
                var component = createSliderComponent(FIO2, componentNumber)
                componentNumber++;
                if (component === null) {
                    // Error Handling
                    console.log("Error creating object "+element);
                }
                break
            default:
                console.log("Dont have the input "+ element)
                break;
        } 
    }
}


function getComponents(inputList) {
    console.log("pls work")

    componentNumber=1;
    var wantedComponents = inputList
    addToView(wantedComponents)

    return componentNumber
}

// takes in the view and sends name and value to python
function getComponentsValues(sliders){
    for(var i = 0; i < sliders.children.length; ++i)
        ModeSelect.sendValues(sliders.children[i].name, sliders.children[i].value)

}
