'use strict';

export default class Utils{

  constructor() {
    this.ENTER_KEY_CODE = 13;
    $.getJSON("/assets/omikuji.json", (data) => {
      console.log(data)
    })
  }

  requestOmikuji(text){

    let json = {
      body: "omikuji "+text,
      SenderName: ""
    }

    this.post("/api/messages","POST",JSON.stringify(json))

  }

  post(url,method,data){
    $.ajax({
      url: url,
      type: method,
      data: data
    })
    .done( data => {
      this.get(url,"GET")
    })
    .fail( err => {
      throw new Error(err)
    })
  }

  get(url,method){
    $.ajax({
      url: url,
      type: method
    })
    .done( data => {
      this.kujiRender(data.result[data.result.length - 1].body)
    })
    .fail( err => {
      throw new Error(err)
    })
  }

  kujiRender(kuji_type){

    let image_url = "/assets/images/凶.png"

    $(".kuji img").attr("src",image_url)
    changeView()

    kuji_type == "凶" ? this.doBadAnimation() : this.doAnimation()

  }

  doAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
    TweenMax.to('.kuji', 1, { rotation: 360 });
    TweenMax.to(".cracker img", 0.5, { width: "100%" })
    TweenMax.to(".cracker img", 3, { delay: 0.5 ,autoAlpha: 0 })
  }

  doBadAnimation(){
    TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
  }

  changeView(){
    $(".kuji").removeClass("_hidden")
    $(".box").addClass("_hidden")
    $(".reload").removeClass("_hidden")
  }

}