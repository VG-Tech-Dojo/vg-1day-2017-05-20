'use strict';

import Utils from "./Utils"

(function($){

  let utils = new Utils()

  $(".box .user-input").keydown( e => {
    if(e.keyCode === ENTER_KEY_CODE){
      var keyword = $(".user-input").val()
      $(".user-input").val("")
      utils.requestOmikuji(keyword)
    }
  })

})(jQuery)