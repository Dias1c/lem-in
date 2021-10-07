// To Do Frame Log on HTML

var btn_SendContent = document.getElementById("btn_SendContent")
var btn_SetParams = document.getElementById("btn_SetParams")
var btn_Animate = document.getElementById("btn_Animate")
var tb_FrameIndex = document.getElementById("tb_FrameIndex")
var l_FrameInfo = document.getElementById("l_FrameInfo")
var tb_Content = document.getElementById("tb_Content")
var tb_Result = document.getElementById("tb_Result")
var IsFramePlay = false

var popup = {
    Window: document.getElementById("window"),
    Title: document.getElementById("windowTitle"),
    Text: document.getElementById("windowText"),
    Btn_Ok: document.getElementById("windowAction"),
    Btn_Close: document.getElementById("windowClose")
}

let Drawwer = new GraphDrawer("#b_result", new Graphs(document.getElementById("tb_Result").value))


window.addEventListener("load", OnLoad)

// Events
function OnLoad() {
    InitEvents()
    InitPopup()
    SetDefaultValuesOnHtml()
}

async function btn_SendContent_OnClick() {
    btn_SendContent.disabled = true
    try {
        let data = { "Content": tb_Content.value }
        const response = await fetch("/api/lemin", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        await response
        if (!response.ok) {
            let error = { message: (await response.text()).toString() }
            if (error.message == "") {
                error.message = "The server is not responding"
            }
            throw error
        }
        tb_Result.value = await response.text();//(await response.text()).toString().split('\\n').join("\n")
        // Start Show Graph
        btn_SendContent.disabled = false
        btn_SetParams_OnClick()
    } catch (error) {
        //Show Pop Up window
        let enableBtn = () => { btn_SendContent.disabled = false }
        ShowPopupWindow("Error", error.message, enableBtn, enableBtn)
    }
}
function btn_SetParams_OnClick() {
    try {
        let transition = parseInt(d3.select("#tb_Transition").node().value)
        let tileSize = parseInt(d3.select("#tb_TileSize").node().value)
        let circleRadius = parseInt(d3.select("#tb_CircleRadius").node().value)
        if (Number.isNaN(transition) || transition < 4) {
            throw { message: "Parametr: Transition < 4ms\nParametrs Not Accepted!" }
        } else if (Number.isNaN(tileSize) || tileSize < 1) {
            throw { message: "Parametr: Tile Size < 1px\nParametrs Not Accepted!" }
        } else if (Number.isNaN(circleRadius) || circleRadius < 1) {
            throw { message: "Parametr: Circle Radius < 1px\nParametrs Not Accepted!" }
        }
        Drawwer.TileSize = tileSize
        Drawwer.CircleRadius = parseInt(d3.select("#tb_CircleRadius").node().value)
        Drawwer.LineWidth = Drawwer.CircleRadius / 2
        Drawwer.Transition = transition

        Drawwer.Graph = new Graphs(document.getElementById("tb_Result").value)
        Drawwer.ReDrawGraph()
        ResetFrameValues()
    } catch (error) {
        ShowPopupWindow("Error", error.message)
    }
}
function btn_Animate_OnClick() {
    IsFramePlay = true
    EnambleAnimationControllers(IsFramePlay)
    function animate(frame) {
        let min = parseInt(tb_FrameIndex.min)
        let max = parseInt(tb_FrameIndex.max)

        Drawwer.ShowFrame(frame)
        // l_FrameInfo.innerHTML = `${tb_FrameIndex.value}/${tb_FrameIndex.max}`
        if (frame <= min) {
            frame = min
        } else if (frame >= max) {
            tb_FrameIndex.value = frame
            frame = max
            IsFramePlay = false
        }
        tb_FrameIndex.value = frame
        l_FrameInfo.innerHTML = `${tb_FrameIndex.value}/${tb_FrameIndex.max}`
        if (IsFramePlay) {
            setTimeout(() => { animate(frame + 1) }, Drawwer.Transition)
        } else {
            EnambleAnimationControllers(IsFramePlay)
        }
    }
    animate(0)
}


// FUNCTIONS
function InitEvents() {
    btn_SendContent.addEventListener("click", btn_SendContent_OnClick)
    btn_SetParams.addEventListener("click", btn_SetParams_OnClick)
    btn_Animate.addEventListener("click", btn_Animate_OnClick)
    tb_FrameIndex.addEventListener("input", () => {
        Drawwer.ShowFrame(parseInt(tb_FrameIndex.value))
        l_FrameInfo.innerHTML = `${tb_FrameIndex.value}/${tb_FrameIndex.max}`
    })
}

function InitPopup() {
    popup.Btn_Close.addEventListener("click", () => { popup.Window.style.display = "none" })
    popup.Btn_Ok.addEventListener("click", () => { popup.Window.style.display = "none" })
}

function ShowPopupWindow(title, message, onClickOk, onClickClose) {
    // Do Old Window Event
    if (popup.Window.style.display == "block") {
        popup.Btn_Close.click()
    }
    // Set Values
    if (onClickOk != null) {
        popup.Btn_Ok.addEventListener("click", () => {
            popup.Window.style.display = "none"
            onClickOk()
        })
    } else {
        popup.Btn_Ok.addEventListener("click", () => {
            popup.Window.style.display = "none"
        });
    }
    if (onClickClose != null) {
        popup.Btn_Close.addEventListener("click", () => {
            popup.Window.style.display = "none"
            onClickClose()
        })
    } else {
        popup.Btn_Close.addEventListener("click", () => {
            popup.Window.style.display = "none"
        });
    }
    popup.Title.textContent = title
    popup.Text.textContent = message
    popup.Window.style.display = "block"
}

function SetDefaultValuesOnHtml() {
    tb_Result.value = `20
a 0 0
##start
b 0 4
c 2 2
d 4 0
e 3 4
f 6 2
g 6 4
h 5 7
i 9 0
j 9 4
k 9 6
l 7 8
##end
m 12 4
n 12 8
a-b
b-e
e-g
g-j
j-m
a-d
d-i
i-j
b-c
c-f
f-g
e-h
h-l
l-n
n-m
g-k
k-m

L1-a L2-e L3-c 
L1-d L2-h L3-f L4-a L5-e L6-c 
L1-i L2-l L3-g L4-d L5-h L6-f L7-a L8-e L9-c 
L1-j L2-n L3-k L4-i L5-l L6-g L7-d L8-h L9-f L10-a L11-e L12-c 
L1-m L2-m L3-m L4-j L5-n L6-k L7-i L8-l L9-g L10-d L11-h L12-f L13-a L14-e L15-c 
L4-m L5-m L6-m L7-j L8-n L9-k L10-i L11-l L12-g L13-d L14-h L15-f L16-a L17-e L18-c 
L7-m L8-m L9-m L10-j L11-n L12-k L13-i L14-l L15-g L16-d L17-h L18-f L19-a L20-e 
L10-m L11-m L12-m L13-j L14-n L15-k L16-i L17-l L18-g L19-d L20-h 
L13-m L14-m L15-m L16-j L17-n L18-k L19-i L20-l 
L16-m L17-m L18-m L19-j L20-n 
L19-m L20-m 
`
    btn_SetParams_OnClick()
    ResetFrameValues()
}

function EnambleAnimationControllers(enable) {
    btn_Animate.disabled = enable
    tb_FrameIndex.disabled = enable
}

function ResetFrameValues() {
    let min = 0,
        value = tb_FrameIndex.value,
        max = Drawwer.Graph.Steps.length - 1;
    max = max < 0 ? 0 : max;
    if (value > max || value < min) {
        value = min
        console.log("Value changed")
    }

    tb_FrameIndex.min = min;
    tb_FrameIndex.value = value;
    tb_FrameIndex.max = max;
    l_FrameInfo.innerHTML = `${value}/${max}`
    Drawwer.ShowFrame(value)
    IsFramePlay = false
}