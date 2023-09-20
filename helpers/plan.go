package helpers

import (
	"log"

	"github.com/MhmoudGit/go-digital-menu/models"
	"gorm.io/gorm"
)

func GetPlans(db *gorm.DB) ([]models.Plan, error) {
	var plans []models.Plan
	result := db.Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}
	return plans, nil
}

func GetPlan(db *gorm.DB, id uint) (models.Plan, error) {
	var plan models.Plan
	result := db.First(&plan, id)
	if result.Error != nil {
		return plan, result.Error
	}
	return plan, nil
}

func CreatePlan(db *gorm.DB, Plan *models.Plan) error {
	// Create the Plan in the database
	result := db.Create(Plan)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Plan was created successfully....")
	return nil
}

func UpdatePlan(db *gorm.DB, updatePlan *models.Plan, id uint) error {
	var PlanToUpdate models.Plan
	result := db.First(&PlanToUpdate, id).Save(updatePlan)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Plan was updated successfully....")
	return nil
}

func DeletePlan(db *gorm.DB, id uint) error {
	var Plan models.Plan
	result := db.Delete(&Plan, id)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Plan was deleted successfully....")
	return nil
}
