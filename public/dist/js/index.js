layui.use(["jquery", "layer", "miniAdmin", "conf"], function () {
  var $ = layui.jquery,
    layer = layui.layer,
    config = layui.conf;
  miniAdmin = layui.miniAdmin;
  //检查登录
  if (!sessionStorage.getItem("token")) {
    window.location = "/login.html";
  }
  $("#user").text(sessionStorage.getItem("account"));
  var options = {
    iniUrl: config.server + "/api/menu?token=" + sessionStorage["token"], // 初始化接口
    clearUrl: "api/clear.json", // 缓存清理接口
    renderPageVersion: true, // 初始化页面是否加版本号
    bgColorDefault: 0, // 主题默认配置
    multiModule: true, // 是否开启多模块
    menuChildOpen: false, // 是否默认展开菜单
    loadingTime: 0, // 初始化加载时间
    pageAnim: true, // 切换菜单动画
  };
  miniAdmin.render(options);

  $(".login-out").on("click", function () {
    sessionStorage.removeItem("token");
    sessionStorage.removeItem("expire");
    sessionStorage.removeItem("account");
    window.location = "/login.html";
  });

  //5分钟刷新一次token
  setInterval(function () {
    var now = new Date().getTime() / 1000;
    if (+sessionStorage["expire"] - now < 600) {
      $.get(
        config.server + "/api/refresh?token=" + sessionStorage["token"]
      ).then(function (d) {
        sessionStorage.setItem("token", d.data.AccessToken);
        sessionStorage.setItem("expire", d.data.Expire);
      });
    }
  }, 1000 * 60 * 5);
});
