'use strict';

// constants
const ENTER_KEY_CODE = 13;

// event set
$(".box .user-input").keydown( e => {
  if(e.keyCode === ENTER_KEY_CODE){
    var keyword = $(".user-input").val()
    $(".user-input").val("")
    requestOmikuji(keyword)
  }
})

// functions
var requestOmikuji = text => {

  var json = {
    body: "omikuji "+text,
    SenderName: ""
  }

  $.ajax({
    url: "/api/messages",
    type: "POST",
    data: JSON.stringify(json)
  })
  .done( data => {
    $.ajax({
      url: "/api/messages",
      type: "GET"
    })
    .done( data => {
      kujiRender(data.result[data.result.length -1].body)
    })
    .fail( err => {
      throw new Error(err)
    })
  })
  .fail( err => {
    throw new Error(err)
  })

}

var getOmukujiJson = () => {
  $.getJSON("/assets/omikuji.json", (data) => {
    console.log(data)
  })
}
getOmukujiJson()

var kujiRender = kuji_type => {

  var image_url = "/assets/images/凶.png"

  $(".kuji img").attr("src",image_url)
  changeView()

  kuji_type == "凶" ? doBadAnimation() : doAnimation()

}

var doAnimation = () => {
  TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
  TweenMax.to('.kuji', 1, { rotation: 360 });
  TweenMax.to(".cracker img", 0.5, { width: "100%" })
  TweenMax.to(".cracker img", 3, { delay: 0.5 ,autoAlpha: 0 })
}

var doBadAnimation = () => {
  TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
}

function changeView(){
  $(".kuji").removeClass("_hidden")
  $(".box").addClass("_hidden")
  $(".reload").removeClass("_hidden")
}