package models

import (
	"database/sql"
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID           int       `json:"id"`
	Password     string    `json:"password"`
	SapUserCode  string    `json:"sap_user_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

func CreateUser(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        user.CreatedAt = time.Now()
        // user.UpdatedAt = time.Now()

        query := `
            INSERT INTO Haus_Inventory_System_Dev.dbo.[User] (password, sap_user_code, created_at)
            VALUES (@password, @sap_user_code, @created_at);
            SELECT ID = convert(bigint, SCOPE_IDENTITY());
        `
        var id int64
        err := db.QueryRow(query, sql.Named("password", user.Password), sql.Named("sap_user_code", user.SapUserCode), sql.Named("created_at", user.CreatedAt)).Scan(&id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        user.ID = int(id)
        c.JSON(http.StatusCreated, user)
    }
}

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		query := `SELECT id, password, sap_user_code, created_at, updated_at, deleted_at FROM Haus_Inventory_System_Dev.dbo.[User] WHERE deleted_at IS NULL`
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Password, &user.SapUserCode, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user User
		query := `SELECT id, password, sap_user_code, created_at, updated_at, deleted_at FROM Haus_Inventory_System_Dev.dbo.[User] WHERE id = @id AND deleted_at IS NULL`
		err := db.QueryRow(query, sql.Named("id", id)).Scan(&user.ID, &user.Password, &user.SapUserCode, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.UpdatedAt = time.Now()

		query := `UPDATE Haus_Inventory_System_Dev.dbo.[User] SET password = @password, sap_user_code = @sap_user_code, updated_at = @updated_at WHERE id = @id AND deleted_at IS NULL`
		result, err := db.Exec(query, sql.Named("password", user.Password), sql.Named("sap_user_code", user.SapUserCode), sql.Named("updated_at", user.UpdatedAt), sql.Named("id", id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		deletedAt := time.Now()

		query := `UPDATE Haus_Inventory_System_Dev.dbo.[User] SET deleted_at = @deleted_at WHERE id = @id AND deleted_at IS NULL`
		result, err := db.Exec(query, sql.Named("deleted_at", deletedAt), sql.Named("id", id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
