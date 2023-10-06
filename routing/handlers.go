package routing

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/XessiveObserver/zoo/db"
	"github.com/XessiveObserver/zoo/model"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

// GetAnimals retrieves all animals in the database and returns in JSON
func GetAnimals(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.DB.Query("SELECT * FROM animalias")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var animals []model.Animal
	for rows.Next() {
		var animal model.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Kind, &animal.Diet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		animals = append(animals, animal)
	}
	json.NewEncoder(w).Encode(animals)
}

// GetAnimal retrieves a specific animal from the database by UUID and returns in JSON
func GetAnimal(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.Parse(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var animal model.Animal
	err = db.DB.QueryRow("SELECT * FROM animalias WHERE id = $1", id).Scan(&animal.ID, &animal.Name, &animal.Kind, &animal.Diet)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Animal not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

// CreateAnimal insterts a new animal into the database and returns animal as JSON
func CreateAnimal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var animal model.Animal
	_ = json.NewDecoder(r.Body).Decode(&animal)

	// Insert new animal into database
	_, err := db.DB.Exec("INSERT INTO animalias (id, name, kind, diet) VALUES ($1, $2, $3, $4)", animal.ID, animal.Name, animal.Kind, animal.Diet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(animal)
}

// UpdateAnimal updates an animal in the database by UUID and returns the updated animal in JSON
func UpdateAnimal(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.Parse(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var animal model.Animal
	_ = json.NewDecoder(r.Body).Decode(&animal)

	// Upadate the animal in the database
	_, err = db.DB.Exec("UPDATE animalias SET name = $1, kind = $2, diet = $3 WHERE id = $4", animal.Name, animal.Kind, animal.Diet, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	animal.ID = id
	json.NewEncoder(w).Encode(animal)
}

// DeleteAnimal deletes a animal from the database by UUID and returns the deleted animal as JSON
func DeleteAnimal(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.Parse(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var animal model.Animal
	err = db.DB.QueryRow("DELETE FROM animalias WHERE id = $1 RETURNING id, name, kind, diet", id).Scan(&animal.ID, &animal.Name, &animal.Kind, &animal.Diet)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(animal)
}
