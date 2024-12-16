package db

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"

	"github.com/jpradass/whattoeat/models"
)

func GetRandomRecipe() map[string]models.Recipe {
	conn := getConn()

	defer conn.Close()

	db := sql.OpenDB(conn)
	defer db.Close()

	recipes := getNumRecipes(db)
	recipeID := selectRandomRecipeId(db, recipes)

	rows, err := db.Query(
		fmt.Sprintf(`SELECT r.Id, r.Title, r.Image, r.Description, 
      i.Name AS Ingredient, 
      s.StepOrder, s.StepDescription
    FROM Recipe r
    LEFT JOIN Ingredient i ON r.Id = i.RecipeId
    LEFT JOIN Step s ON r.Id = s.RecipeId
    WHERE r.Id = '%s'
    ORDER BY s.StepOrder;`, recipeID))
	if err != nil {
		fmt.Println("Error executing query:", err)
		os.Exit(1)
	}

	defer rows.Close()

	var recipe models.Recipe
	ingredientMap := make(map[string]bool) // To handle duplicate ingredients
	stepMap := make(map[int]string)

	for rows.Next() {
		var (
			id, title, image, description, ingredientName, stepDescription string
			stepOrder                                                      int
		)

		err := rows.Scan(&id, &title, &image, &description, &ingredientName, &stepOrder, &stepDescription)
		if err != nil {
			panic(err)
		}

		// Set Recipe fields
		if recipe.Id == "" {
			recipe.Id = id
			recipe.Title = title
			recipe.Image = image
			recipe.Description = description
		}

		// Add unique ingredients
		if ingredientName != "" && !ingredientMap[ingredientName] {
			recipe.Ingredients = append(recipe.Ingredients, models.Ingredient{Name: ingredientName})
			ingredientMap[ingredientName] = true
		}

		// Add steps in order
		if stepDescription != "" {
			stepMap[stepOrder] = stepDescription
		}
	}

	for i := 1; i <= len(stepMap); i++ {
		recipe.Steps = append(recipe.Steps, models.Step{StepDescription: stepMap[i]})
	}

	return map[string]models.Recipe{"Recipe": recipe}
}

func getConn() *libsql.Connector {
	dbName := "maindb.db"
	dbURL := os.Getenv("WHATTOEAT_DB_URL")
	// dbURL := "libsql://maindb-jpradass.turso.io"
	authToken := os.Getenv("WHATTOEAT_DB_AUTH")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(dir, dbName)
	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbURL, libsql.WithAuthToken(authToken))
	if err != nil {
		fmt.Println("Error getting connector:", err)
		os.Exit(1)
	}

	return connector
}

func selectRandomRecipeId(db *sql.DB, offset int) string {
	recipeNumber := rand.IntN(offset - 1 + 1)
	var recipeId string

	row := db.QueryRow("SELECT Id FROM Recipe LIMIT 1 OFFSET ?", recipeNumber)
	row.Scan(&recipeId)

	return recipeId
}

func getNumRecipes(db *sql.DB) int {
	rows, err := db.Query(
		`SELECT COUNT(*) AS recipes FROM Recipe;`)
	if err != nil {
		fmt.Println("Error executing query:", err)
		os.Exit(1)
	}

	defer rows.Close()

	var recipes int
	for rows.Next() {
		err := rows.Scan(&recipes)
		if err != nil {
			panic(err)
		}
	}

	return recipes
}
