import Utils from "./Utils"

let utils = new Utils()
const ENTER_KEY_CODE = 13;

let userInput = document.querySelector(".box .user-input")

// keydownをとってEnterキーを押したら送信
document.addEventListener("keydown", e => {
  if(e.keyCode === ENTER_KEY_CODE){
    let keyword = userInput.value
    utils.requestOmikuji(keyword)
  }
})