package database

import "github.com/xnacly/meal-planner/api/models"

type Database struct {
	storage map[int]models.Meal
}

var internal *Database

func Get() *Database {
	if internal == nil {
		return &Database{
			storage: make(map[int]models.Meal),
		}
	}
	return internal
}

func (d *Database) AddMeal(m models.Meal) {

}

func (d *Database) RemoveMeal(m models.Meal) {

}

func (d *Database) RandomMeal(m models.Meal) {

}
