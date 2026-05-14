import request from '@/utils/request'

// 用户相关接口
export const userApi = {
  // 初始化用户
  initUsers() {
    return request.post('/user/init')
  },
  
  // 获取用户列表
  getUserList() {
    return request.get('/user/list')
  },
  
  // 获取用户信息
  getUserInfo(userId) {
    return request.get(`/user/info/${userId}`)
  }
}

// 里程相关接口
export const trafficApi = {
  // 添加里程记录
  addTraffic(data) {
    return request.post('/traffic/add', data)
  },
  
  // 获取里程列表
  getTrafficList(userId) {
    return request.get(`/traffic/list/${userId}`)
  }
}

// 积分相关接口
export const pointApi = {
  // 核算积分
  calculatePoints(userId) {
    return request.post('/point/calculate', { userId })
  },
  
  // 获取积分余额
  getPointBalance(userId) {
    return request.get(`/point/balance/${userId}`)
  },
  
  // 获取积分日志
  getPointLogs(userId) {
    return request.get(`/point/logs/${userId}`)
  }
}

// 优惠券相关接口
export const couponApi = {
  // 获取优惠券列表
  getCouponList() {
    return request.get('/coupon/list')
  },
  
  // 兑换优惠券
  exchangeCoupon(data) {
    return request.post('/coupon/exchange', data)
  },
  
  // 获取我的优惠券
  getMyCoupons(userId) {
    return request.get(`/coupon/my/${userId}`)
  }
}

// 商品相关接口
export const goodsApi = {
  // 获取商品列表
  getGoodsList() {
    return request.get('/goods/list')
  }
}

// 订单相关接口
export const orderApi = {
  // 创建订单
  createOrder(data) {
    return request.post('/order/create', data)
  },
  
  // 获取订单列表
  getOrderList(userId) {
    return request.get(`/order/list/${userId}`)
  }
}
