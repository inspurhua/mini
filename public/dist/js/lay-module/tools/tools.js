layui.define([], function (exports) {
  "use strict";

  var _MOD = "tools";
  var util = {
    param: function (url, key) {
      const theRequest = new Object();
      let pos = url.indexOf("?");
      if (pos != -1) {
        const str = url.substr(pos + 1);
        let strs = str.split("&");
        for (let i = 0; i < strs.length; i++) {
          theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
      }
      return theRequest[key];
    },
    open: function (url, target) {
      var a = document.createElement("a");
      a.setAttribute("href", url);
      if (target == null) {
        target = "";
      }
      a.setAttribute("target", target);
      document.body.appendChild(a);
      if (a.click) {
        a.click();
      } else {
        try {
          var evt = document.createEvent("Event");
          a.initEvent("click", true, true);
          a.dispatchEvent(evt);
        } catch (e) {
          window.open(url);
        }
      }
      document.body.removeChild(a);
    },
  };

  exports(_MOD, util);
});
