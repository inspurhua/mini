layui.define([], function (exports) {
    "use strict";

    var _MOD = 'tools';
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
    };


    exports(_MOD, util);
});