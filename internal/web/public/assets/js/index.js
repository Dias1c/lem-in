// To Do Frame Log on HTML

var btn_SendContent = document.getElementById("btn_SendContent")
var btn_SetParams = document.getElementById("btn_SetParams")
var btn_Animate = document.getElementById("btn_Animate")
var btn_ResultResize = document.getElementById("btn_ResultResize")
var tb_FrameIndex = document.getElementById("tb_FrameIndex")
var l_FrameInfo = document.getElementById("l_FrameInfo")
var tb_Content = document.getElementById("tb_Content")
var b_Result = document.getElementById("b_result")
var tb_Result = document.getElementById("tb_Result")
var IsFramePlay = false
var cb_ShowHideText = document.getElementById("cb_ShowHideText")

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

    const controller = new AbortController()
    const timeoutId = setTimeout(() => controller.abort(), 60000)
    try {

        let data = { "Content": tb_Content.value }
        const response = await fetch("/api/lemin", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            signal: controller.signal,
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
    clearTimeout(timeoutId)
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
        cb_ShowHideText_OnChange()
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
function btn_ResultResize_OnClick() {
    if (btn_ResultResize.checked == null) {
        btn_ResultResize.checked = true
    }
    switch (btn_ResultResize.checked) {
        case true:
            b_Result.classList.add("result--maximize")
            btn_ResultResize.classList.add("btn--green")
            break;
        case false:
            b_Result.classList.remove("result--maximize")
            btn_ResultResize.classList.remove("btn--green")
            break;
    }
    btn_ResultResize.checked = !btn_ResultResize.checked
}
function cb_ShowHideText_OnChange(e) {
    let nameElements = document.querySelectorAll("text")
    let display = "block"
    if (cb_ShowHideText.checked) {
        display = "none"
    }
    if (document.querySelector("text") == null || document.querySelector("text").style.display == display) {
        return
    }
    nameElements.forEach((element) => { element.style.display = display });
}


// FUNCTIONS
function InitEvents() {
    btn_SendContent.addEventListener("click", btn_SendContent_OnClick)
    btn_ResultResize.addEventListener("click", btn_ResultResize_OnClick)
    btn_SetParams.addEventListener("click", btn_SetParams_OnClick)
    btn_Animate.addEventListener("click", btn_Animate_OnClick)
    tb_FrameIndex.addEventListener("input", () => {
        Drawwer.ShowFrame(parseInt(tb_FrameIndex.value))
        l_FrameInfo.innerHTML = `${tb_FrameIndex.value}/${tb_FrameIndex.max}`
    })
    cb_ShowHideText.addEventListener("change", cb_ShowHideText_OnChange)
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
    tb_Result.value = `10
a 1 1
##start
b 1 3
c 2 2
d 3 1
e 2 3
f 4 2
g 4 3
h 3 4
i 5 1
j 5 3
k 5 4
l 4 5
##end
m 6 3
n 6 5
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

L1-a L2-c L3-e 
L1-d L2-f L3-h L4-a L5-c L6-e 
L1-i L2-g L3-l L4-d L5-f L6-h L7-a L8-c L9-e 
L1-j L2-k L3-n L4-i L5-g L6-l L7-d L8-f L9-h L10-a 
L1-m L2-m L3-m L4-j L5-k L6-n L7-i L8-g L9-l L10-d 
L4-m L5-m L6-m L7-j L8-k L9-n L10-i 
L7-m L8-m L9-m L10-j 
L10-m
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
