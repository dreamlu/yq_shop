// page/component/new-pages/user/user.js
var app = getApp()

Page({
  data:{
    thumb:'',
    nickname:'',
    orders:[],
    hasAddress:false,
    address:{},
    domain:app.globalData.domain
  },
  onLoad(){
    var self = this;
    /**
     * 获取用户信息,升级参考：https://developers.weixin.qq.com/community/develop/doc/0000a26e1aca6012e896a517556c01
     */
    // wx.getUserInfo({
    //   success: function(res){
    //     self.setData({
    //       thumb: res.userInfo.avatarUrl,
    //       nickname: res.userInfo.nickName
    //     })
    //   }
    // })

    /**
     * 发起请求获取订单列表信息
     */
    let openid = wx.getStorageSync('openid')
    wx.request({
      url:app.globalData.domain+'/orders/'+openid,
      success(res){  
        self.setData({
          orders: res.data
        })
      }
    })
  },
  onShow(){
    var self = this;
    /**
     * 获取本地缓存 地址信息
     */
    wx.getStorage({
      key: 'address',
      success: function(res){
        self.setData({
          hasAddress: true,
          address: res.data
        })
      }
    })
    /**
     * 获取订单信息
     */
  },
  /**
   * 发起支付请求
   */
  payOrders(){
    wx.requestPayment({
      timeStamp: 'String1',
      nonceStr: 'String2',
      package: 'String3',
      signType: 'MD5',
      paySign: 'String4',
      success: function(res){
        console.log(res)
      },
      fail: function(res) {
        wx.showModal({
          title:'支付提示',
          content:'<text>',
          showCancel: false
        })
      }
    })
  }
})