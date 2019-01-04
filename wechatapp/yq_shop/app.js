App({
  onLaunch: function () {
    var that = this
    wx.login({
      success: function (res) {
        if (res.code) {
          //发起网络请求,获取openid
          wx.request({
            url: that.globalData.domain + '/login',
            data: {
              code: res.code//临时code验证
            },
            success: function (obj) {
              //本地存储openid,unionid更好
              wx.setStorage({
                key: "openid",
                data: obj.data.Openid
              })
              // wx.showModal({
              //   title: '测试',
              //   content: "返回数据"+obj.data.OpenId
              // })
            }
          })
        } else {
          console.log('登录失败！' + res.errMsg)
        }
      }
    })
  }
,
  globalData: {
    hasLogin: false,
    domain: 'http://192.168.31.193:8009',
  }
})
