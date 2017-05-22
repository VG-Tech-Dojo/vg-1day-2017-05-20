(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory();
	else if(typeof define === 'function' && define.amd)
		define([], factory);
	else {
		var a = factory();
		for(var i in a) (typeof exports === 'object' ? exports : root)[i] = a[i];
	}
})(this, function() {
return /******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// identity function for calling harmony imports with the correct context
/******/ 	__webpack_require__.i = function(value) { return value; };
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


var _Utils = __webpack_require__(1);

var _Utils2 = _interopRequireDefault(_Utils);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

(function ($) {

  var utils = new _Utils2.default();
  var ENTER_KEY_CODE = 13;

  $(".box .user-input").keydown(function (e) {
    if (e.keyCode === ENTER_KEY_CODE) {
      var keyword = $(".user-input").val();
      $(".user-input").val("");
      utils.requestOmikuji(keyword);
    }
  });
})(jQuery);

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var Utils = function () {
  function Utils() {
    _classCallCheck(this, Utils);

    this.getKujiJson();
  }

  _createClass(Utils, [{
    key: "getKujiJson",
    value: function getKujiJson() {
      var _this = this;

      $.getJSON("/assets/omikuji.json", function (data) {
        _this.omikujiMap = data;
      });
    }
  }, {
    key: "requestOmikuji",
    value: function requestOmikuji(text) {

      var json = {
        body: "omikuji " + text,
        SenderName: ""
      };

      this.post("/api/messages", "POST", JSON.stringify(json));
    }
  }, {
    key: "post",
    value: function post(url, method, data) {
      var _this2 = this;

      $.ajax({
        url: url,
        type: method,
        data: data
      }).done(function (data) {
        _this2.get(url, "GET");
      }).fail(function (err) {
        throw new Error(err);
      });
    }
  }, {
    key: "get",
    value: function get(url, method) {
      var _this3 = this;

      $.ajax({
        url: url,
        type: method
      }).done(function (data) {
        _this3.kujiRender(data.result[data.result.length - 1].body);
      }).fail(function (err) {
        throw new Error(err);
      });
    }
  }, {
    key: "kujiRender",
    value: function kujiRender(kuji_type) {

      var image_url = this.omikujiMap[kuji_type];

      $(".kuji img").attr("src", image_url);
      this.changeView();

      kuji_type == "凶" ? this.doBadAnimation() : this.doAnimation();
    }
  }, {
    key: "doAnimation",
    value: function doAnimation() {
      TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
      TweenMax.to('.kuji', 1, { rotation: 360 });
      TweenMax.to(".cracker img", 0.5, { width: "100%" });
      TweenMax.to(".cracker img", 3, { delay: 0.5, autoAlpha: 0 });
    }
  }, {
    key: "doBadAnimation",
    value: function doBadAnimation() {
      TweenMax.to('.kuji', 1, { autoAlpha: 1, ease: Expo.easeInOut });
    }
  }, {
    key: "changeView",
    value: function changeView() {
      $(".kuji").removeClass("_hidden");
      $(".box").addClass("_hidden");
      $(".reload").removeClass("_hidden");
    }
  }]);

  return Utils;
}();

exports.default = Utils;

/***/ })
/******/ ]);
});