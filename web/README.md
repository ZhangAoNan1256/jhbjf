# JHB积分商城 - 前端使用说明

## 📁 项目结构

```
web/
├─ global.js           # 全局配置（API地址、工具函数）
├─ index.html          # 首页（入口导航）
├─ user.html           # 用户管理
├─ traffic.html        # 里程管理
├─ point.html          # 积分中心
├─ coupon.html         # 优惠券中心
├─ goods.html          # 商品列表
└─ order.html          # 订单管理
```

## 🚀 快速开始

### 1. 启动后端服务

确保后端Go服务已经启动在 `http://localhost:8888`

```bash
cd d:\enviroment\Go\GoCode\gocode\jhb1\api
go run jhb.go
```

### 2. 打开前端页面

**直接双击任意HTML文件即可运行！**

推荐从 `index.html` 开始：
```
双击 web/index.html
```

或者在浏览器中打开：
```
file:///d:/enviroment/Go/GoCode/gocode/jhb1/web/index.html
```

## 📖 使用流程

### 第1步：初始化用户
1. 打开 `user.html`
2. 点击"初始化用户"按钮创建测试数据
3. 点击任意用户行，自动选择该用户（userId会保存到localStorage）

### 第2步：录入里程
1. 打开 `traffic.html`
2. userId已自动填充（如果之前选择了用户）
3. 填写车牌号、里程、时间
4. 点击"提交里程记录"

### 第3步：核算积分
1. 打开 `point.html`
2. userId已自动填充
3. 点击"核算积分"按钮
4. 查看积分余额和日志

### 第4步：兑换优惠券
1. 打开 `coupon.html`
2. userId已自动填充
3. 点击"查看可兑换券"浏览可用优惠券
4. 点击"立即兑换"按钮
5. 在"我的优惠券"区域查看已兑换的券

### 第5步：下单购物
1. 打开 `goods.html` 查看商品列表
2. 点击"立即下单"跳转到订单页面
3. 或直接打开 `order.html`
4. 填写商品ID（和可选的优惠券ID）
5. 点击"提交订单"
6. 在下方查看订单列表

## 🔧 技术栈

- **HTML5**: 页面结构
- **Bootstrap 5**: UI组件库（CDN引入）
- **Axios**: HTTP请求库（CDN引入）
- **原生JavaScript**: 逻辑交互
- **localStorage**: 用户状态共享

## ✨ 特性

✅ **零配置**：无需安装Node.js、npm等环境  
✅ **零编译**：无需打包构建，双击即用  
✅ **零依赖**：所有库通过CDN引入  
✅ **状态共享**：userId通过localStorage在所有页面共享  
✅ **响应式设计**：基于Bootstrap，支持移动端  

## 🌐 API配置

如需修改后端地址，编辑 `global.js` 文件：

```javascript
const BASE_URL = "http://localhost:8888";  // 修改为你的后端地址
```

## 📝 注意事项

1. **必须先启动后端服务**，否则所有API请求都会失败
2. **建议使用Chrome或Edge浏览器**以获得最佳体验
3. **首次使用需要先初始化用户**（在user.html页面）
4. **所有页面的userId会自动同步**，在一个页面选择用户后，其他页面会自动填充

## 🎯 业务流程

```
用户管理 → 录入里程 → 核算积分 → 兑换优惠券 → 下单购物
   ↓          ↓          ↓           ↓           ↓
user.html → traffic.html → point.html → coupon.html → order.html
```

## 💡 提示

- 每个页面右上角显示当前选中的用户ID
- 点击用户行的"选择"按钮会复制userId到剪贴板
- 所有操作都有成功/失败提示
- 列表数据可以手动刷新

---

**开发完成时间**: 2026-05-12  
**技术选型**: 原生HTML + Bootstrap 5 + Axios
