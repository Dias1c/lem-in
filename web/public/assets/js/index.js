var btn_SendContent = document.getElementById("btn_SendContent")
var tb_Content = document.getElementById("tb_Content")
var tb_Result = document.getElementById("tb_Result")
var popup = {
    Window: document.getElementById("window"),
    Title: document.getElementById("windowTitle"),
    Text: document.getElementById("windowText"),
    Btn_Ok: document.getElementById("windowAction"),
    Btn_Close: document.getElementById("windowClose")
}

window.addEventListener("load", OnLoad)

// Events
function OnLoad() {
    btn_SendContent.addEventListener("click", btn_SendContent_OnClick)
    InitPopup()
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
    } catch (error) {
        //Show Pop Up window
        let enableBtn = () => { btn_SendContent.disabled = false }
        ShowPopupWindow("Error", error.message, enableBtn, enableBtn)
    }
}

function InitPopup() {
    popup.Btn_Close.addEventListener("click", () => { popup.Window.style.display = "none" })
    popup.Btn_Ok.addEventListener("click", () => { popup.Window.style.display = "none" })
}

function ShowPopupWindow(title, message, onClickOk, onClickClose) {
    if (onClickOk != null) {
        popup.Btn_Ok.addEventListener("click", () => {
            popup.Window.style.display = "none"
            onClickOk()
        })
    }
    if (onClickClose != null) {
        popup.Btn_Close.addEventListener("click", () => {
            popup.Window.style.display = "none"
            onClickClose()
        })
    }
    popup.Title.textContent = title
    popup.Text.textContent = message
    popup.Window.style.display = "block"
}