// 全局配置
// 使用反向代理时，baseURL 留空或设为当前域名，请求会发给 localhost:3000
const BASE_URL = "";

// 初始化axios默认配置
axios.defaults.baseURL = BASE_URL;
axios.defaults.headers['Content-Type'] = 'application/json';

// 统一提示函数
function tip(msg, type = 'info') {
    alert(msg);
}

// 获取当前用户ID
function getCurrentUserId() {
    return localStorage.getItem("userId");
}

// 设置当前用户ID
function setCurrentUserId(userId) {
    localStorage.setItem("userId", userId);
}

// 通用错误处理
function handleError(error) {
    console.error('请求错误:', error);
    if (error.response) {
        tip(`错误: ${error.response.data?.message || error.response.statusText}`);
    } else {
        tip('网络请求失败，请检查后端服务是否启动');
    }
}
