package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// a struct to store the data in memory

type employees struct {
	Employee_ID int64  `json:"emlpoyee_id"`
	Department  string `json:"department"`
	Name        string `json:"name"`
	Salary      int64  `json:"salary"`
}

func getData(GET *gin.Context) {
	rows, err := db.Query("SELECT emlpoyee_id, department, name, salary FROM employees")
	if err != nil {
		GET.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}
	defer rows.Close()
	var data []employees

	for rows.Next() {
		var dt employees
		if err := rows.Scan(&dt.Employee_ID, &dt.Department, &dt.Name, &dt.Salary); err != nil {
			GET.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
			return
		}
		data = append(data, dt)
	}
	GET.JSON(http.StatusOK, data)
}

func addData(POST *gin.Context) {
	var nemp employees
	if err := POST.BindJSON(&nemp); err != nil {
		POST.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := db.Exec("INSERT INTO employees (emlpoyee_id, department, name, salary) VALUES (?, ?, ?, ?)",
		nemp.Employee_ID, nemp.Department, nemp.Name, nemp.Salary)
	if err != nil {
		POST.JSON(http.StatusBadRequest, gin.H{"error": "failed to add data"})
		return
	}

	POST.JSON(http.StatusCreated, "data added")
}

func delEmp(DEL *gin.Context) {
	emp := DEL.Param("id")
	result, err := db.Exec("DELETE FROM employees WHERE emlpoyee_id = ?", emp)
	if err != nil {
		DEL.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		DEL.JSON(http.StatusBadRequest, gin.H{"Error": "failed to check"})
		return
	}

	if rowsAffected == 0 {
		DEL.JSON(http.StatusNotFound, gin.H{"Error": "Employee not found"})
		return
	}

	DEL.JSON(http.StatusOK, "Employee with Id"+emp+"is deleted")
}

func getEmpById(GETe *gin.Context) {
	empId := GETe.Param("id")
	var employee employees
	err := db.QueryRow("SELECT emlpoyee_id, department, name, salary FROM employees WHERE emlpoyee_id = ?", empId).Scan(
		&employee.Employee_ID, &employee.Department, &employee.Name, &employee.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			GETe.JSON(http.StatusNotFound, "No employee with ID: "+empId+"found")
		} else {
			GETe.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		}
		return
	}

	GETe.JSON(http.StatusOK, employee)

}
func main() {
	db = connect()
	defer db.Close()
	router := gin.Default()
	router.GET("/getData", getData)
	router.POST("/addData", addData)
	router.DELETE("/delete/:id", delEmp)
	router.GET("/getEmp/:id", getEmpById)
	router.Run("localhost:8083")
}
