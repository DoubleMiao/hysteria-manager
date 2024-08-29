package services

import (
    "hysteria-manager/database"
    "hysteria-manager/models"
)

func CreateInstance(instance *models.Instance) error {
    return database.DB.Create(instance).Error
}

func GetInstance(id uint) (*models.Instance, error) {
    var instance models.Instance
    err := database.DB.First(&instance, id).Error
    return &instance, err
}

// 类似地实现 UpdateInstance 和 DeleteInstance
