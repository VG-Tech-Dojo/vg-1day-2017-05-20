export default class Utils{

  constructor() {
    this.getKujiJson()
  }

  /**
   *  getKujiJson
   *  クジのマップをjsonファイルから取得する
   */
  getKujiJson(){
    this.getJSON("/assets/omikuji.json")
  }

  /**
   *  getJSON
   *  URLをもとにjsonファイルの中身を取得する
   *
   *  @param { string } url
   */
  getJSON(url){
    let xhr = new XMLHttpRequest()
    xhr.open("GET",url)
    xhr.send()

    xhr.onload = data => {
      this.omikujiMap = JSON.parse(data.target.response)
    }
  }

  /**
   *  requestOmikuji
   *  次のおみくじをリクエストする
   *
   *  @param { string } text
   */
  requestOmikuji(text){

    let json = {
      body: "omikuji " + text,
      SenderName: ""
    }

    this.post("/api/messages","POST",JSON.stringify(json))

  }

  /**
   *  post
   *  XMLHttpRequwstを用いて、値の更新を行う
   *
   *  @param { string } url
   *  @param { string } method
   *  @param { string } data
   */
  post(url,method,data){
    this.request(url,method,data, () => {
      this.get(url,"GET")
    })
  }

  /**
   *  get
   *  XMLHttpRequestを用いて値の取得を行う
   *
   *  @param { string } url
   *  @param { string } method
   */
  get(url,method){
    this.request(url,method,null, res => {
      let response = JSON.parse(res.target.response)
      this.kujiRender(response.result[response.result.length -1].body)
    })
  }

  /**
   *  kujiRender
   *  おみくじをレンダリングする
   *
   *  @param { string } kuji_type
   */
  kujiRender(kuji_type){

    let image_url = this.omikujiMap[kuji_type]

    let kuji_image = document.querySelector(".kuji img")
    kuji_image.src = image_url

    this.changeView()

    kuji_type == "凶" ? this.doBadAnimation() : this.doAnimation()

  }

  /**
   *  doAnimation
   *  クジとクラッカーを表示するアニメーション
   */
  doAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut })
    TweenMax.to('.kuji', 1, { rotation: 360 })
    TweenMax.to(".cracker img", 0.5, { width: "100%" })
    TweenMax.to(".cracker img", 3, { delay: 0.5 ,autoAlpha: 0 })
  }

  /**
   *  doBadAnimation
   *  凶の時は、クラッカーのアニメーションはなくす
   */
  doBadAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut })
  }

  /**
   *  changeView
   *  クジとリロードボタンを表示して、クジを引く箱を非表示にする
   */
  changeView(){
    document.querySelector(".kuji").classList.remove("_hidden")
    document.querySelector(".box").classList.add("_hidden")
    document.querySelector(".reload").classList.remove("_hidden")
  }

  /**
   *  request
   *  XMLHttpRequestを用いで、データを取得し、callbackを実行する
   *
   *  @param { string } url
   *  @param { string } method
   *  @param { string } data
   *  @param { function } callback
   */
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