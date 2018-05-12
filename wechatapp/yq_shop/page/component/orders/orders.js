// page/component/orders/orders.js
Page({
  data:{
    address:{},
    hasAddress: false,
    total:0,
    orders:[],
      domain:getApp().globalData.domain
  },
onLoad:function(options){
  this.setData({
    orders: JSON.parse(options.data)
  })
},
  onReady() {
    this.getTotalPrice();
  },
  
  onShow:function(){
    const self = this;
    wx.getStorage({
      key:'address',
      success(res) {
        self.setData({
          address: res.data,
          hasAddress: true
        })
      }
    })
  },

  /**
   * 计算总价
   */
  getTotalPrice() {
    let orders = this.data.orders;
    let total = 0;
    for(let i = 0; i < orders.length; i++) {
      total += orders[i].Goods_nums * orders[i].Goods_prices;
    }
    this.setData({
      total: total
    })
  },

  toPay() {
    wx.showModal({
      title: '提示',
      content: '本系统由个人开发，支付系统需要企业接口,暂未完成',
      text:'center',
      complete() {
        wx.switchTab({
          url: '/page/component/user/user'
        })
      }
    })
  }
})