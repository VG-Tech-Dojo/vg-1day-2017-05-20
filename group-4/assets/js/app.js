// constants
var ENTER_KEY_CODE = 13;
var kuji_map = {
  "大吉": "/assets/images/大吉.png",
  "吉": "/assets/images/吉.png",
  "中吉": "/assets/images/中吉.png",
  "小吉": "/assets/images/小吉.png",
  "末吉": "/assets/images/末吉.png",
  "凶": "/assets/images/凶.png",
}

// event set
$(".box .user-input").keydown(function(e){
  if(e.keyCode === ENTER_KEY_CODE){
    var keyword = $(".user-input").val()
    $(".user-input").val("")
    requestOmikuji(keyword)
  }
})

// http
function requestOmikuji(text){

  var json = {
    body: "omikuji "+text,
    SenderName: ""
  }

  $.ajax({
    url: "/api/messages",
    type: "POST",
    data: JSON.stringify(json)
  })
  .done(function(data){
    $.ajax({
      url: "/api/messages",
      type: "GET"
    })
    .done(function(data){
      kujiRender(data.result[data.result.length -1].body)
    })
    .fail(function(err){
      throw new Error(err)
    })
  })
  .fail(function(err){
    throw new Error(err)
  })

}

function kujiRender(kuji_type){

  var image_url = kuji_map[kuji_type]

  $(".kuji img").attr("src",image_url)
  changeView()

  kuji_type == "凶" ? doBadAnimation() : doAnimation()

}

function doAnimation(){
  TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
  TweenMax.to('.kuji', 1, { rotation: 360 });
  TweenMax.to(".cracker img", 0.5, { width: "100%" })
  TweenMax.to(".cracker img", 3, { delay: 0.5 ,autoAlpha: 0 })
}

function doBadAnimation(){
  TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
}

function changeView(){
  $(".kuji").removeClass("_hidden")
  $(".box").addClass("_hidden")
  $(".reload").removeClass("_hidden")
}