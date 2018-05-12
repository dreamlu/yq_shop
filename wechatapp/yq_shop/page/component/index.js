var app = getApp()

Page({
  data: {
    imgUrls: [
      '/image/b1.jpg',
      '/image/b2.jpg',
      '/image/b3.jpg'
    ],
    indicatorDots: false,
    autoplay: false,
    interval: 3000,
    duration: 800,
    goods: {},//数组[]也行，区别呢
    domain: app.globalData.domain//域名
  },
  onLoad: function () {
    var that = this//后面不能加逗号,否则出错
      wx.request({
        url: app.globalData.domain + '/goods',
        success: function (res) {
          that.setData({//如果在sucess直接写this就变成了wx.request()的this了
            goods: res.data
          })
        }
      })
  }
})