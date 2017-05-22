export default class Utils{

  constructor() {
    this.getKujiJson()
  }

  getKujiJson(){
    this.getJSON("/assets/omikuji.json")
  }

  getJSON(url){
    let xhr = new XMLHttpRequest()
    xhr.open("GET",url)
    xhr.send()

    xhr.onload = (data) => {
      this.omikujiMap = JSON.parse(data.target.response)
    }
  }

  requestOmikuji(text){

    let json = {
      body: "omikuji " + text,
      SenderName: ""
    }

    this.post("/api/messages","POST",JSON.stringify(json))

  }

  post(url,method,data){
    this.request(url,method,data, () => {
      this.get(url,"GET")
    })
  }

  get(url,method){
    this.request(url,method,null,(res) => {
      let response = JSON.parse(res.target.response)
      this.kujiRender(response.result[response.result.length -1].body)
    })
  }

  kujiRender(kuji_type){

    let image_url = this.omikujiMap[kuji_type]

    let kuji_image = document.querySelector(".kuji img")
    kuji_image.src = image_url
    
    this.changeView()

    kuji_type == "凶" ? this.doBadAnimation() : this.doAnimation()

  }

  doAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut })
    TweenMax.to('.kuji', 1, { rotation: 360 })
    TweenMax.to(".cracker img", 0.5, { width: "100%" })
    TweenMax.to(".cracker img", 3, { delay: 0.5 ,autoAlpha: 0 })
  }

  doBadAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut })
  }

  changeView(){
    document.querySelector(".kuji").classList.remove("_hidden")
    document.querySelector(".box").classList.add("_hidden")
    document.querySelector(".reload").classList.remove("_hidden")
  }

  request(url,method,data,callback){
    let xhr = new XMLHttpRequest()
    xhr.open(method,url)
    xhr.setRequestHeader("Content-type","application/json")

    xhr.send(data)

    xhr.onload = (data) => {
      callback(data)
    }
  }

}