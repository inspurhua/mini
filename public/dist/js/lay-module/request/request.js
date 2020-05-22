layui.define(["jquery", "layer"],
    function (exports) {
        var $ = layui.jquery,
            layer = layui.layer;
        let request = function (url, method = "get", data = {}) {
            if (method.toLocaleLowerCase() == "get") {
                url = url + "?_" + new Date().getTime();
            }
            if (method.toLocaleLowerCase() == "post" || method.toLocaleLowerCase() == "put") {
                if (typeof data != 'string') {
                    data = JSON.stringify(data);
                }
            }
            return new Promise((resolve, reject) => {
                $.ajax({
                    url,
                    type: method,
                    data,
                    contentType: 'application/json',
                    headers: {
                        token: sessionStorage.getItem("token")
                    },
                    success(res) {
                        resolve({
                            code: 200,
                            msg: res.msg,
                            data: res.data
                        })
                    },
                    error(err) {
                        err.responseJSON && layer.msg(err.responseJSON.msg);
                        if (err.responseJSON.code == 401) {
                            window.location.href = "/admin/login.html";
                        }
                        reject({
                            code: 200,
                            msg: "",
                            data: []
                        })
                    }
                })
            })
        }
        exports('request', request);
    });
