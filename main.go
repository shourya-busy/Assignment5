package main

import (
	"fmt"
	//"slices"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)


type Manager struct {
	EmployeeID uint64  `pg:"employee_id"`
	ManagerID  uint64  `pg:"manager_id"`
}

func main() {
	db := loadDatabase()

	var managers []Manager
	err := db.Model(&managers).Select()
	if err != nil {

		fmt.Println(err.Error())
		
	}

	adjList := make(map[uint64][]uint64)
	for _, manager := range managers {
		adjList[manager.EmployeeID] = append(adjList[manager.EmployeeID], manager.ManagerID)
	}

	visited := make(map[uint64]bool)

	var empID,managerID uint64

	fmt.Println("Enter the empID :")
	fmt.Scanln(&empID)
	fmt.Println("Enter the managerID :")
	fmt.Scanln(&managerID)

	// if !slices.Contains(adjList[empID],managerID) {
	// 	fmt.Printf("%d is not the Manager of Employee : %d \n",managerID,empID)
	// 	return
	// }

	adjList[empID]  = append(adjList[empID], managerID)

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

func dfsDetectCycle(empID uint64 , managerID uint64, adjList map[uint64][]uint64, visited map[uint64]bool) bool {

	if visited[managerID] {
		return true
	}

	visited[empID] = true;
	visited[managerID] = true;

	for _, manager := range adjList[managerID] {
		if dfsDetectCycle(managerID, manager,adjList,visited) {
			return true
		}
	}

	return false
}