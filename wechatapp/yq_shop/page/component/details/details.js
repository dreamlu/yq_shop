// page/component/details/details.js
var app = getApp()

Page({
  data: {
    goods: {
      id: 1,
      image: '/image/goods1.png',
      title: '新鲜梨花带雨',
      price: 0.01,
      stock: '有货',
      detail: '这里是浙传校服详情。',
      parameter: '125g/个',
      service: '不支持退货'
    },
    num: 1,
    totalNum: 1,
    hasCarts: false,
    curIndex: 0,
    show: false,
    scaleCart: false,
    domain: app.globalData.domain
  },
  onLoad: function (obj) {
    var that = this
    this.setData({
      'goods.id': obj.id
    }),
      wx.request({
        url: app.globalData.domain + '/goods/' + that.data.goods.id,
        success: function (res) {
          that.setData({//如果在sucess直接写this就变成了wx.request()的this了
            'goods.title': res.data.Goods_name,
            'goods.image': res.data.Goods_picture_address,
            'goods.price': res.data.Goods_prices
          })
        }
      })
  },
  onShareAppMessage(options){
    return {
      title: '校服购',
      path: '/details/details?id='+this.data.goods.id,
      success: function (res) {
        // 转发成功
      },
      fail: function (res) {
        // 转发失败
      }
    }
  },
  //增加数量
  addCount() {
    let num = this.data.num;
    num++;
    this.setData({
      num: num
    })
  },
  //加入购物车
  addToCart() {
    const self = this;
    const num = this.data.num;
    let total = this.data.totalNum-1;
    //加入购物车
    //that = this
    let openid = wx.getStorageSync('openid')
    wx.request({
      url: app.globalData.domain + '/cart',
      header: { 'content-type': 'application/x-www-form-urlencoded' },//只有meithod没有heder无法传值
      method: "POST",
      data: {
        openid: openid,
        goods_id: self.data.goods.id,
        goods_nums: self.data.totalNum
      }
    }),

    //加入购物车的特效动画
    self.setData({
      show: true
    })
    setTimeout(function () {
      self.setData({
        show: false,
        scaleCart: true
      }),setTimeout(function () {
        self.setData({
          scaleCart: false,
          hasCarts: true,
          totalNum: num + total
        })
      }, 200)
    }, 300)
  },
  //详情
  bindTap(e) {
    const index = parseInt(e.currentTarget.dataset.index);
    this.setData({
      curIndex: index
    })
  }

})