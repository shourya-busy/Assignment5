package main

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)


type Manager struct {
	EmployeeID int      `pg:"employee_id"`
	ManagerID  int      `pg:"manager_id"`
}

func main() {
	db := loadDatabase()

	var managers []Manager
	err := db.Model(&managers).Select()
	if err != nil {

		fmt.Println(err.Error())
		
	}

	adjList := make(map[int]int)
	for _, manager := range managers {
		adjList[manager.EmployeeID] = manager.ManagerID
	}

	visited := make(map[int]bool)

	var empID,managerID int

	fmt.Println("Enter the empID :")
	fmt.Scanln(&empID)
	fmt.Println("Enter the managerID :")
	fmt.Scanln(&managerID)

	isCycleDetected := dfsDetectCycle(empID,managerID,adjList,visited)

	fmt.Println("Cycle Detected :",isCycleDetected)

}


func loadDatabase() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "shouryagautam",
		Password: "shourya",
		Database: "postgres",
	})


	err := db.Model((*Manager)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}

func dfsDetectCycle(empID int , managerID int, adjList map[int]int, visited map[int]bool) bool {

	if adjList[managerID] == 0{
		return false
	}

	if visited[managerID] {
		return true
	}

	visited[empID] = true;
	visited[managerID] = true;

	return dfsDetectCycle(managerID,adjList[managerID],adjList,visited)
}