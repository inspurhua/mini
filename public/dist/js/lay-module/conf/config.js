layui.define([], function (exports) {
    "use strict";

    var _MOD = 'conf';
    var config = {
        server: "http://127.0.0.1:8000",
        headers:{
            token:sessionStorage['token']
        }
    };


    exports(_MOD, config);
});