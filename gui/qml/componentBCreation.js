// basic button 
var baseButton = Qt.createComponent("BaseLargeButton.qml");
var inputPage;
var inputList;
var sprite;
var inputListHeight;

var activeButtons = ["Pressure A/C"]


function createBreathButtons(buttonList) {
    for (var i = 0; i < buttonList.length; i++) {
        const mode = buttonList[i];
        var createButtonComp = baseButton.createObject(rowBreath, {
            title: mode
        })
    }

    return createButtonComp
}

function createTriggerButtons(buttonList) {
    for (var i = 0; i < buttonList.length; i++) {
        const mode = buttonList[i];
        var createButtonComp = baseButton.createObject(rowTrigger, {
            title: mode
        })
    }

    return createButtonComp
}

function toArray(string) {
    var result = string.split(",")
    return result
}

function createComponent(chosenButton) {
    if (ModeSelect.mode===""){
        // first click
        ModeSelect.mode = chosenButton
        heading.text="Select Breathe Type"
        // show breath buttons
        rowBreath.visible=true
    } else if(ModeSelect.breath===""){
        heading.text="Select Trigger Type"
        ModeSelect.breath = chosenButton
        // show trigger buttons
        rowTrigger.visible=true
    } else if (ModeSelect.trigger===""){
        heading.text="Select Input"
        ModeSelect.trigger = chosenButton
        // show input page
        flickableItems.visible=true
    }

    // list is updated depending if mode, breath and trigger are populated
    var list = toArray(ModeSelect.buttonList)
    // for(var i = rowButtons.children.length; i > 0 ; i--) {
    //     rowButtons.children[i-1].visible=false
    // }

    if (ModeSelect.trigger!=="") {
        // hide trigger buttons
        rowTrigger.visible=false
        // make root page scrollable
        flickablePage.interactive=true
        // make sliders using list
        createInputs(list)
        // dynamicall resize page
        flickablePage.contentHeight = flickableItems.children[0].contentHeight
    }else if(ModeSelect.breath!==""){
        // make trigger row
        console.log("making trigger")
        createTriggerButtons(list)
        // remove breath buttons
        rowBreath.visible=false
    }else if(ModeSelect.mode!==""){
        // make breath row
        console.log("making breath")
        createBreathButtons(list)
        // remove mode row
        rowButtons.visible=false
    }
}

// when back button is clicked
// delete and clean depending on 
// the current view
function backButton() {
    if (ModeSelect.trigger!==""){
        // set trigger to empty
        ModeSelect.trigger=""
        // hide input page
        flickableItems.visible=false
        // make page non interactive
        flickablePage.interactive=false
        // change heading
        heading.text="Select Trigger Type"
        // show trigger
        rowTrigger.visible=true
    
    // from trigger to breath
    } else if (ModeSelect.breath!==""){
        // set breath to empty
        ModeSelect.breath=""
        // hide trigger buttons
        rowTrigger.visible=false
        // hide previous buttons
        for(var i = rowTrigger.children.length; i > 0 ; i--) {
            console.log("destroying trigger")
            rowTrigger.children[i-1].height=0
        }
        // change heading
        heading.text="Select Breath Type"
        // show trigger
        rowBreath.visible=true
    
    // from breath to mode
    }else if (ModeSelect.mode!==""){
        // set mode to empty
        ModeSelect.mode=""
        // hide breath buttons
        rowBreath.visible=false
        // hide previous buttons
        for(var i = rowBreath.children.length; i > 0 ; i--) {
            console.log("destroying breath")
            rowBreath.children[i-1].height=0
        }
        // change heading
        heading.text="Select Mode"
        // show trigger
        rowButtons.visible=true
    
    }
}

function finishCreation() {
    if (inputPage.status == Component.Ready) {
        sprite = inputPage.createObject(flickableItems, {
            inputList:inputList
        });
        if (sprite == null) {
            // Error Handling
            console.log("Error creating object");
        }
    } else if (inputPage.status == Component.Error) {
        // Error Handling
        console.log("Error loading component:", inputPage.errorString());
    }
}

function createInputs(newInputList) {
    inputList = newInputList

    inputPage = Qt.createComponent("InputList.qml");
    if (inputPage.status == Component.Ready){
        finishCreation();
    }
    else{
    inputPage.statusChanged.connect(finishCreation);
    }
}
