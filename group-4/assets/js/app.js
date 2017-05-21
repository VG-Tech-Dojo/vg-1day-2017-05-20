// constants
var ENTER_KEY_CODE = 13;
var kuji_map = {
  "daikichi": "/assets/images/kuji-daikichi.png"
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
  // $.ajax({
  //   url: "",
  //   type: "GET",
  //   data: {
  //     text: text
  //   }
  // })
  // .done(function(data){
  //   // console.log(data)
  //   kujiRender()
  // })
  // .fail(function(err){
  //   throw new Error(err)
  // })

  kujiRender(text)
}

function kujiRender(kuji_type){

  var image_url = kuji_map[kuji_type]

  $(".kuji img").attr("src",image_url)
  changeView()

  TweenMax.to('.kuji', 1.5, { autoAlpha: 1, ease: Expo.easeInOut });
}

function changeView(){
  $(".kuji").removeClass("_hidden")
  $(".box").addClass("_hidden")
  $(".reload").removeClass("_hidden")
}