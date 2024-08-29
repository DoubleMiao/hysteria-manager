package utils

import (
    "hysteria-manager/database"
    "hysteria-manager/models"
    "log"
    "time"
)

func StartCleanupTask() {
    ticker := time.NewTicker(1 * time.Hour)
    go func() {
        for {
            select {
            case <-ticker.C:
                cleanUpExpiredInstances()
            }
        }
    }()
}

func cleanUpExpiredInstances() {
    var instances []models.Instance
    now := time.Now()
    database.DB.Where("expires_at < ?", now).Find(&instances)

    for _, instance := range instances {
        if err := database.DB.Delete(&instance).Error; err != nil {
            log.Printf("Failed to delete instance %v: %v", instance.ID, err)
        } else {
            log.Printf("Deleted expired instance %v", instance.ID)
        }
    }
}
