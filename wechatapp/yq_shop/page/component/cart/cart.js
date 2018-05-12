// page/component/new-pages/cart/cart.js
var app = getApp()

Page({
  data: {
    carts: [],                    //购物车列表
    orders :[],                  // 选中的订单列表
    hasList: false,          // 列表是否有数据
    totalPrice: 0,           // 总价，初始为0
    selectAllStatus: false,    // 全选状态，默认全选
    obj: {
      name: "hello"
    },
    domain: app.globalData.domain
  },
  onShow() {
    let that = this
    let openid = wx.getStorageSync('openid')
    wx.request({
      url: app.globalData.domain + '/cart/' + openid,
      success: function (res) {
        that.setData({
          hasList: true,
          carts: res.data
        })
      }
    })
  },
  // onHide(){
  //   updateCart()
  // },
  // onUnload(){
  //   updateCart()
  // },
  // //更新购物车商品数量
  // updateCart(){

  // },
  /**
   * 删除购物车当前商品
   */
  deleteList(e) {
    const index = e.currentTarget.dataset.index;
    let carts = this.data.carts;

    //提交服务器
    let that = this
    let openid = wx.getStorageSync('openid')
    wx.request({
      url: app.globalData.domain + '/cartDelete',
      header: { 'content-type': 'application/x-www-form-urlencoded' },
      method: "POST",
      data: {
        openid: openid,
        goods_id: that.data.carts[index].Goods_id
      }
    })
    carts.splice(index, 1);//删除goods
    this.setData({
      carts: carts
    });
    if (!carts.length) {
      this.setData({
        hasList: false
      });
    } else {
      this.getTotalPrice();
    }
  },
  /**
  * 当前商品选中事件
  */
  selectList(e) {
    const index = e.currentTarget.dataset.index;
    let carts = this.data.carts;
    const selected = carts[index].selected;
    carts[index].selected = !selected;
    this.setData({
      carts: carts
    });
    this.getTotalPrice();
  },
  /**
   * 购物车全选事件
   */
  selectAll(e) {
    let selectAllStatus = this.data.selectAllStatus;
    selectAllStatus = !selectAllStatus;
    let carts = this.data.carts;

    for (let i = 0; i < carts.length; i++) {
      carts[i].selected = selectAllStatus;
    }
    this.setData({
      selectAllStatus: selectAllStatus,
      carts: carts
    });
    this.getTotalPrice();
  },
  //下订单(获取选中商品属性)
  orderTap(){
    let carts = this.data.carts;
    //遍历添加order
    for (let i = 0; i < carts.length; i++) {
      if(carts[i].selected){
        this.data.orders.push(carts[i])
      }
    }
    var data = JSON.stringify(this.data.orders);
    let that = this
    wx.navigateTo({
      url: '../orders/orders?data='+data,
      success:function(){   
        //订单完成后删除购物车相应信息并创建订单(本地缓存吧,懒人)
        //wx.setStorageSync("orders", data)
        //重置orders,防止push会增添数据
        that.setData({
          orders:[],
          totalPrice:0
        })
      }
    })
  },
  /**
   * 绑定加数量事件
   */
  addCount(e) {
    const index = e.currentTarget.dataset.index;
    let carts = this.data.carts;
    let num = carts[index].Goods_nums;
    num = num + 1;
    carts[index].Goods_nums = num;
    this.setData({
      carts: carts
    });
    this.updateCart(index);
    this.getTotalPrice();
  },
  /**
   * 绑定减数量事件
   */
  minusCount(e) {
    const index = e.currentTarget.dataset.index;
    const obj = e.currentTarget.dataset.obj;
    let carts = this.data.carts;
    let num = carts[index].Goods_nums;
    if (num <= 1) {
      return false;
    }
    num = num - 1;
    carts[index].Goods_nums = num;
    this.setData({
      carts: carts
    });
    this.updateCart(index);
    this.getTotalPrice();
  },
  //更新购物车商品数量,应该通过onHide(),onUnload()中统一提交数据
  //偷懒者:lucheng
  updateCart(index) {
    let that = this
    let openid = wx.getStorageSync('openid')
    wx.request({
      url: app.globalData.domain + '/cartUpdate',
      header: { 'content-type': 'application/x-www-form-urlencoded' },
      method: "POST",
      data: {
        openid: openid,
        goods_id: that.data.carts[index].Goods_id,
        goods_nums: that.data.carts[index].Goods_nums
      }
    })
  },
  /**
   * 计算总价
   */
  getTotalPrice() {
    let carts = this.data.carts;                  // 获取购物车列表
    let total = 0;
    for (let i = 0; i < carts.length; i++) {         // 循环列表得到每个数据
      if (carts[i].selected) {                     // 判断选中才会计算价格
        total += carts[i].Goods_nums * carts[i].Goods_prices;   // 所有价格加起来
      }
    }
    this.setData({                                // 最后赋值到data中渲染到页面
      carts: carts,
      totalPrice: total.toFixed(2)
    });
  }

})