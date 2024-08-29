package main

import (
    "hysteria-manager/routers"
    "hysteria-manager/database"
    "log"
    "net/http"
)

func main() {
    // 初始化数据库
    database.InitDB()

    // 启动定时任务
    utils.StartCleanupTask()

    // 设置路由
    router := routers.SetupRouter()

    // 启动服务
    log.Println("Starting server on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
