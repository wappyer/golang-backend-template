// pm2 启动配置文件
module.exports = {
    apps: [{
        "name": "wappyer-pro",
        "cwd": "./",
        "script": "./main",
        // "min_uptime": "60s", // 设置的时间内启动失败自动重新启动
        // "max_restarts": 3, // 重启次数
        "error_file": "./log/pm2-err.log",
        "out_file": "./log/pm2-out.log",
        "pid_file": "./pm2.pid",
        "args": ["-env", "pro"],
        "watch": ["main", "./config/pro/config.yaml"] // 监听文件，文件变动pm2会自动重载
    }]
};