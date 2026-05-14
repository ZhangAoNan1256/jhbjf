const http = require('http');
const httpProxy = require('http-proxy');
const fs = require('fs');
const path = require('path');

// 创建代理实例
const proxy = httpProxy.createProxyServer({});

// 目标后端地址
const TARGET = 'http://localhost:8888';

// 静态文件目录
const STATIC_DIR = __dirname;

const server = http.createServer((req, res) => {
    // 1. 如果是 /api 开头的请求，转发给后端
    if (req.url.startsWith('/api')) {
        console.log(`[Proxy] Forwarding ${req.method} ${req.url} to ${TARGET}`);
        proxy.web(req, res, { target: TARGET }, (err) => {
            console.error('[Proxy Error]', err);
            res.writeHead(500);
            res.end('Proxy Error');
        });
    } 
    // 2. 其他请求，作为静态文件返回
    else {
        // 提取路径部分（去掉查询参数）
        const urlPath = req.url.split('?')[0];
        
        let filePath = path.join(STATIC_DIR, urlPath);
        
        // 如果路径以 / 结尾，默认访问 index.html
        if (urlPath.endsWith('/') || urlPath === '/') {
            filePath = path.join(STATIC_DIR, 'index.html');
        }

        // 检查文件是否存在
        fs.access(filePath, fs.constants.F_OK, (err) => {
            if (err) {
                // 文件不存在，尝试添加 .html 后缀（仅当路径没有扩展名时）
                if (!path.extname(filePath)) {
                    filePath += '.html';
                    
                    // 再次检查添加后缀后的文件是否存在
                    fs.access(filePath, fs.constants.F_OK, (err2) => {
                        if (err2) {
                            // 仍然不存在，返回 404
                            res.writeHead(404);
                            res.end('File not found: ' + req.url);
                        } else {
                            serveFile(filePath, res);
                        }
                    });
                } else {
                    // 有扩展名但文件不存在，返回 404
                    res.writeHead(404);
                    res.end('File not found: ' + req.url);
                }
            } else {
                // 文件存在，直接返回
                serveFile(filePath, res);
            }
        });
    }
});

// 辅助函数：读取并返回文件
function serveFile(filePath, res) {
    fs.readFile(filePath, (error, content) => {
        if (error) {
            if (error.code == 'ENOENT') {
                res.writeHead(404);
                res.end('File not found');
            } else {
                res.writeHead(500);
                res.end('Server Error: ' + error.code);
            }
        } else {
            // 设置 Content-Type
            const extname = String(path.extname(filePath)).toLowerCase();
            const mimeTypes = {
                '.html': 'text/html',
                '.js': 'text/javascript',
                '.css': 'text/css',
                '.json': 'application/json',
                '.png': 'image/png',
                '.jpg': 'image/jpg',
                '.gif': 'image/gif',
                '.svg': 'image/svg+xml',
                '.ico': 'image/x-icon',
                '.woff': 'application/font-woff',
                '.ttf': 'application/font-ttf',
                '.eot': 'application/vnd.ms-fontobject',
                '.otf': 'application/font-otf',
                '.wasm': 'application/wasm'
            };
            const contentType = mimeTypes[extname] || 'application/octet-stream';
            
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content, 'utf-8');
        }
    });
}

const PORT = 3000;
server.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}/`);
    console.log(`API requests will be proxied to ${TARGET}`);
});
