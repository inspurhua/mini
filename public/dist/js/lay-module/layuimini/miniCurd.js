layui.define(
  ["jquery", "laytpl", "form", "layer", "miniAdmin", "miniPage","layedit","conf"],
  function(exports) {
    var $ = layui.jquery;
    var laytpl = layui.laytpl;
    var form = layui.form;
    var miniAdmin = layui.miniAdmin;
    var miniPage = layui.miniPage;
    var layedit = layui.layedit;
    var layer = layui.layer;
    var config=layui.conf;
    if (typeof Object.assign != "function") {
      Object.assign = function(target) {
        "use strict";
        if (target == null) {
          throw new TypeError("Cannot convert undefined or null to object");
        }

        target = Object(target);
        for (var index = 1; index < arguments.length; index++) {
          var source = arguments[index];
          if (source != null) {
            for (var key in source) {
              if (Object.prototype.hasOwnProperty.call(source, key)) {
                target[key] = source[key];
              }
            }
          }
        }
        return target;
      };
    }

    var curd = function(options) {
        //event :add|edit|delete
        //display:function(){}
        //done:function(res){}
      var options = Object.assign(
        { dom:"#tpl_curd",event: "add", titile: "添加", url: "", data: {} },
        options || {}
      );
      var data = options.data;
      var type = options.event == "add" ? "POST" : "PUT";
      var url =
        config.server +
        options.url.replace(/\/$/, "") +
        (options.event == "add" ? "" : "/" + data.id);

      if (options.event == "delete") {
        layer.confirm("真的删除吗?", function(index) {
          $.ajax({
            type: "DELETE",
            headers: config.headers,
            url: url,
            dataType: "json"
          })
            .done(function(res) {
              if (res.code == 0) {
                options.done && options.done();
                miniAdmin.success(res.msg);
              } else {
                miniAdmin.error(res.msg);
              }
            })
            .fail(function() {
              miniAdmin.error("删除出错");
            });
          layer.close(index);
        });
      } else {
        var tpl = $(options.dom)
          .html()
          .trim();
        var content = laytpl(tpl).render(data);
        var openWH = miniPage.getOpenWidthHeight();

        var index = layer.open({
          title: options.title,
          type: 1,
          maxmin: true,
          shadeClose: true,
          area: [openWH[0] + "px", openWH[1] + "px"],
          offset: [openWH[2] + "px", openWH[3] + "px"],
          content: content
        });

        options.display && options.display(data);

        $(window).on("resize", function() {
          layer.full(index);
        });

        form.render();
        // 当前弹出层，防止ID被覆盖
        var parentIndex = layer.index;

        //监听提交
        form.on("submit(saveBtn)", function(data) {
          $.ajax({
            type: type,
            headers: config.headers,
            data: data.field,
            url: url,
            dataType: "json"
          })
            .done(function(res) {
              if (res.code == 0) {
                options.done && options.done();
                miniAdmin.success(res.msg);
              } else {
                miniAdmin.error(res.msg);
              }
            })
            .fail(function() {
              miniAdmin.error("保存出错");
            });
          layer.close(parentIndex);
          return false;
        });
      }
    };

    exports("miniCurd", curd);
  }
);
