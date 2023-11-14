package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//Creating Structure for Employee
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	PhoneNo   string
	Role      string
	Salary    float64
	BirthDate time.Time
}

//Creating structure for storing Employee Database
type EmployeeDB struct {
	employees []Employee
}

//Reciever function for Adding employee
func (db *EmployeeDB) AddEmployee(e Employee) {
	db.employees = append(db.employees, e)
}

//Reciever function for viewing employee
func (db *EmployeeDB) ViewEmployee(id int) (Employee, error) {
	for _, e := range db.employees {
		if e.ID == id {
			return e, nil
		}
	}
	return Employee{}, fmt.Errorf("Employee with ID %d not found", id)
}

//Reciever function for updating employee details
func (db *EmployeeDB) UpdateEmployee(id int, updatedEmployee Employee) error {
	for i, e := range db.employees {
		if e.ID == id {
			db.employees[i] = updatedEmployee
			return nil
		}
	}
	return fmt.Errorf("Employee with ID %d not found", id)
}

//Reciever function for deleting employee
func (db *EmployeeDB) DeleteEmployee(id int) error {
	for i, e := range db.employees {
		if e.ID == id {
			db.employees = append(db.employees[:i], db.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Employee with ID %d not found", id)
}

//Reciever function to list all employees
func (db *EmployeeDB) ListAllEmployees() string {
	employeesJSON, err := json.MarshalIndent(db.employees, "", "\t")
	if err != nil {
		return "Error generating JSON"
	}
	return string(employeesJSON)
}

//Reciever function to list employees sorted by any field
func (db *EmployeeDB) ListEmployeesSortedByField(field string) []Employee {
	employees := make([]Employee, len(db.employees))
	copy(employees, db.employees)

	switch field {
	case "FirstName":
		for i := 0; i < len(employees)-1; i++ {
			for j := 0; j < len(employees)-i-1; j++ {
				if employees[j].FirstName > employees[j+1].FirstName {
					employees[j], employees[j+1] = employees[j+1], employees[j]
				}
			}
		}
	case "LastName":
		for i := 0; i < len(employees)-1; i++ {
			for j := 0; j < len(employees)-i-1; j++ {
				if employees[j].LastName > employees[j+1].LastName {
					employees[j], employees[j+1] = employees[j+1], employees[j]
				}
			}
		}
	case "Email":
		for i := 0; i < len(employees)-1; i++ {
			for j := 0; j < len(employees)-i-1; j++ {
				if employees[j].Email > employees[j+1].Email {
					employees[j], employees[j+1] = employees[j+1], employees[j]
				}
			}
		}
	case "Salary":
		for i := 0; i < len(employees)-1; i++ {
			for j := 0; j < len(employees)-i-1; j++ {
				if employees[j].Salary < employees[j+1].Salary {
					employees[j], employees[j+1] = employees[j+1], employees[j]
				}
			}
		}
	}
	return employees
}

//Reciever function to list employees with upcoming birthdays
func (db *EmployeeDB) ListEmployeesWithUpcomingBirthday() []Employee {
	var upcomingEmployees []Employee
	today := time.Now()

	for _, emp := range db.employees {
		if emp.BirthDate.Month() == today.Month() && emp.BirthDate.Day() >= today.Day() {
			upcomingEmployees = append(upcomingEmployees, emp)
		}
	}
	return upcomingEmployees
}

//Reciever function to search employee
func (db *EmployeeDB) SearchEmployee(query string) []Employee {
	var result []Employee
	for _, emp := range db.employees {
		if emp.FirstName == query || emp.LastName == query || emp.Email == query || emp.Role == query {
			result = append(result, emp)
		}
	}
	return result
}

//Reciever function to check employee is admin or not
func (db *EmployeeDB) IsAdmin(employee Employee) bool {
	return employee.Role == "admin"
}

//Reciever function for login authentication
func (db *EmployeeDB) Login(username, password string) (Employee, error) {
	for _, emp := range db.employees {
		if emp.Email == username && emp.Password == password {
			return emp, nil
		}
	}
	return Employee{}, fmt.Errorf("Invalid username or password")
}

//Normal function for admin's operations
func adminOperations(manager EmployeeDB) {
	num := 0
	for num != 1 {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Admin Operations")
		fmt.Println("1. Add employee details")
		fmt.Println("2. View employee details")
		fmt.Println("3. Update employee details")
		fmt.Println("4. Delete employee details")
		fmt.Println("5. List all employees")
		fmt.Println("6. List employees sorted by a specific field")
		fmt.Println("7. List employees with upcoming birthdays")
		fmt.Println("8. Search employee")
		fmt.Println("--------------------------------------------------")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var employee Employee

			fmt.Println("*********************************************")
			fmt.Println("Enter employee details:")
			fmt.Print("ID: ")
			fmt.Scanln(&employee.ID)
			fmt.Print("First Name: ")
			fmt.Scanln(&employee.FirstName)
			fmt.Print("Last Name: ")
			fmt.Scanln(&employee.LastName)
			fmt.Print("Email: ")
			fmt.Scanln(&employee.Email)
			fmt.Print("Password: ")
			fmt.Scanln(&employee.Password)
			fmt.Print("Phone Number: ")
			fmt.Scanln(&employee.PhoneNo)
			fmt.Print("Role: ")
			fmt.Scanln(&employee.Role)
			fmt.Print("Salary: ")
			fmt.Scanln(&employee.Salary)
			fmt.Print("Birthdate (YYYY-MM-DD): ")
			var birthdateStr string
			fmt.Scanln(&birthdateStr)
			birthdate, err := time.Parse("2006-01-02", birthdateStr)
			if err != nil {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
				return
			}
			employee.BirthDate = birthdate
			manager.AddEmployee(employee)
			fmt.Println("Employee added successfully.")
			fmt.Println("************************************************")
			employees := manager.ListAllEmployees() 
			fmt.Println("All Employees after addition:", employees)
			fmt.Println("************************************************")
		case 2:
			
			var employeeID int
			fmt.Println("************************************************")
			fmt.Print("Enter Employee ID: ")
			fmt.Scanln(&employeeID)

			emp, err := manager.ViewEmployee(employeeID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			
			fmt.Println("Admin: Viewing employee details")
			fmt.Println("Employee Details:")
			fmt.Println("ID:", emp.ID)
			fmt.Println("First Name:", emp.FirstName)
			fmt.Println("Last Name:", emp.LastName)
			fmt.Println("Email:", emp.Email)
			fmt.Println("Phone Number:", emp.PhoneNo)
			fmt.Println("Role:", emp.Role)
			fmt.Println("Salary:", emp.Salary)
			fmt.Println("************************************************")

		case 3:
			fmt.Println("************************************************")
			var employeeID int
			fmt.Print("Enter Employee ID to update: ")
			fmt.Scanln(&employeeID)

			emp, err := manager.ViewEmployee(employeeID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Enter updated details for Employee ID", employeeID)

			var updateChoice int
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("1. Update First Name")
			fmt.Println("2. Update Last Name")
			fmt.Println("3. Update Email")
			fmt.Println("4. Update Phone Number")
			fmt.Println("5. Update Birthdate")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Print("Enter your update choice: ")
			fmt.Scanln(&updateChoice)

			newEmployee := emp

			fmt.Println("************************************************")
			switch updateChoice {
			case 1:
				fmt.Print("First Name: ")
				fmt.Scanln(&newEmployee.FirstName)
			case 2:
				fmt.Print("Last Name: ")
				fmt.Scanln(&newEmployee.LastName)
			case 3:
				fmt.Print("Email: ")
				fmt.Scanln(&newEmployee.Email)
			case 4:
				fmt.Print("Phone Number: ")
				fmt.Scanln(&newEmployee.PhoneNo)
			case 5:
				fmt.Print("Birthdate (YYYY-MM-DD): ")
				var birthdateStr string
				fmt.Scanln(&birthdateStr)
				birthdate, err := time.Parse("2006-01-02", birthdateStr)
				if err != nil {
					fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
					return
				}
				newEmployee.BirthDate = birthdate
			default:
				fmt.Println("Invalid update choice.")
				return
			}

			// Update the employee details
			if err := manager.UpdateEmployee(employeeID, newEmployee); err != nil {
				fmt.Println("Error updating employee details:", err)
				return
			}
			fmt.Println("Employee details updated successfully.")
			fmt.Println("************************************************")

		case 4:
			// Delete employee details
			var employeeID int
			fmt.Println("************************************************")
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scanln(&employeeID)

			if err := manager.DeleteEmployee(employeeID); err != nil {
				fmt.Println("Error deleting employee:", err)
				return
			}
			fmt.Println("Employee details deleted successfully.")
			fmt.Println("************************************************")

		case 5:
			// List all employees
			fmt.Println("************************************************")
			fmt.Println("Admin: Listing all employees")
			employeesJSON := manager.ListAllEmployees()
			fmt.Println(employeesJSON)
			fmt.Println("************************************************")
		case 6:
			// List employees sorted by a specific field
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("Sort employees by:")
			fmt.Println("1. First Name")
			fmt.Println("2. Last Name")
			fmt.Println("3. Email")
			fmt.Println("4. Salary")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

			var sortChoice int
			fmt.Print("Enter sorting choice: ")
			fmt.Scanln(&sortChoice)
			fmt.Println("************************************************")
			var field string
			switch sortChoice {
			case 1:
				field = "FirstName"
			case 2:
				field = "LastName"
			case 3:
				field = "Email"
			case 4:
				field = "Salary"
			default:
				fmt.Println("Invalid sorting choice")
				return
			}

			employees := manager.ListEmployeesSortedByField(field)
			fmt.Println("Employees sorted by", field, ":")
			for _, emp := range employees {
				fmt.Println(emp)
			}
			fmt.Println("************************************************")

		case 7:
			// List employees with upcoming birthdays
			fmt.Println("************************************************")
			upcomingEmployees := manager.ListEmployeesWithUpcomingBirthday()
			if len(upcomingEmployees) == 0 {
				fmt.Println("No employees have upcoming birthdays this month.")
			} else {
				fmt.Println("Employees with upcoming birthdays this month:")
				for _, emp := range upcomingEmployees {
					fmt.Println(emp)
				}
			}
			fmt.Println("************************************************")

		case 8:
			// Search employee (display all fields except password)
			var searchQuery string
			fmt.Println("************************************************")
			fmt.Print("Enter search query: ")
			fmt.Scanln(&searchQuery)

			foundEmployees := manager.SearchEmployee(searchQuery)
			if len(foundEmployees) == 0 {
				fmt.Println("No matching employees found.")
			} else {
				fmt.Println("Matching employees:")
				for _, emp := range foundEmployees {
					fmt.Println("ID:", emp.ID)
					fmt.Println("First Name:", emp.FirstName)
					fmt.Println("Last Name:", emp.LastName)
					fmt.Println("Email:", emp.Email)
					fmt.Println("Phone Number:", emp.PhoneNo)
					fmt.Println("Role:", emp.Role)
					fmt.Println("Salary:", emp.Salary)
					fmt.Println("------------")
				}
			}
			fmt.Println("************************************************")
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("Press 1 to Logout else 0 to continue >>>")
		fmt.Scanln(&num)
	}
}

func nonAdminOperations(manager EmployeeDB, loggedInEmployee Employee) {
	num1 := 0

	for num1 != 1 {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Non-Admin Operations for", loggedInEmployee.FirstName)
		fmt.Println("1. View your details")
		fmt.Println("2. Update your details")
		fmt.Println("3. Search employee")
		fmt.Println("--------------------------------------------------")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("************************************************")
			fmt.Println("Non-Admin: Viewing own details")
			fmt.Println("Employee Details:")
			fmt.Println("ID:", loggedInEmployee.ID)
			fmt.Println("First Name:", loggedInEmployee.FirstName)
			fmt.Println("Last Name:", loggedInEmployee.LastName)
			fmt.Println("Email:", loggedInEmployee.Email)
			fmt.Println("Phone Number:", loggedInEmployee.PhoneNo)
			fmt.Println("Role:", loggedInEmployee.Role)
			fmt.Println("Salary:", loggedInEmployee.Salary)
			fmt.Println("************************************************")

		case 2:
			// Update his/her details
			fmt.Println("Non-Admin: Updating own details")
			var updateField int
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("1. Update First Name")
			fmt.Println("2. Update Last Name")
			fmt.Println("3. Update Email")
			fmt.Println("4. Update Phone Number")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Print("Enter your update choice: ")
			fmt.Scanln(&updateField)

			var existingEmployee = loggedInEmployee 
			var updatedInfo string

			fmt.Println("************************************************")
			switch updateField {
			case 1:
				fmt.Print("Enter updated First Name: ")
				fmt.Scanln(&updatedInfo)
				existingEmployee.FirstName = updatedInfo
			case 2:
				fmt.Print("Enter updated Last Name: ")
				fmt.Scanln(&updatedInfo)
				existingEmployee.LastName = updatedInfo
			case 3:
				fmt.Print("Enter updated Email: ")
				fmt.Scanln(&updatedInfo)
				existingEmployee.Email = updatedInfo
			case 4:
				fmt.Print("Enter updated Phone Number: ")
				fmt.Scanln(&updatedInfo)
				existingEmployee.PhoneNo = updatedInfo
			default:
				fmt.Println("Invalid update choice.")
				return
			}

			// Update the employee details
			if err := manager.UpdateEmployee(loggedInEmployee.ID, existingEmployee); err != nil {
				fmt.Println("Error updating employee details:", err)
				return
			}
			fmt.Println("Employee details updated successfully.")
			fmt.Println("************************************************")

		case 3:
			// Search employee
			fmt.Println("************************************************")
			fmt.Println("Non-Admin: Searching employee")
			var searchQuery string
			fmt.Print("Enter search query: ")
			fmt.Scanln(&searchQuery)

			foundEmployees := manager.SearchEmployee(searchQuery)
			if len(foundEmployees) == 0 {
				fmt.Println("No matching employees found.")
			} else {
				fmt.Println("Matching employees:")
				for _, emp := range foundEmployees {
					fmt.Println("First Name:", emp.FirstName)
					fmt.Println("Last Name:", emp.LastName)
					fmt.Println("Email:", emp.Email)
					fmt.Println("Role:", emp.Role)
					fmt.Println("------------")
				}
			}
			fmt.Println("************************************************")

		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("Press 1 to Logout else 0 to continue >>>")
		fmt.Scanln(&num1)
	}
}

//Main function
func main() {
	employeesDB := &EmployeeDB{employees: []Employee{
		{ID: 1, FirstName: "Vishal", LastName: "Khatpe", Email: "vishal@example.com", Password: "password", PhoneNo: "1234567890", Role: "manager", Salary: 50000.0, BirthDate: time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC)},
		{ID: 2, FirstName: "Harshal", LastName: "Umasare", Email: "harshal@example.com", Password: "harsh151", PhoneNo: "9730077713", Role: "admin", Salary: 45000.0, BirthDate: time.Date(2001, time.May, 1, 0, 0, 0, 0, time.UTC)},
		{ID: 3, FirstName: "Varun", LastName: "Patil", Email: "varun@example.com", Password: "password", PhoneNo: "9876543210", Role: "developer", Salary: 40000.0, BirthDate: time.Date(1992, time.April, 5, 0, 0, 0, 0, time.UTC)},
		{ID: 4, FirstName: "Rohan", LastName: "Singh", Email: "rohan@example.com", Password: "password", PhoneNo: "5556667777", Role: "tester", Salary: 35000.0, BirthDate: time.Date(1988, time.July, 10, 0, 0, 0, 0, time.UTC)},
	}}
	var username, password string
	fmt.Println("╔═════════════════════════════════════════════════╗")
	fmt.Println("║          Welcome to the Employee Management     ║")
	fmt.Println("╠═════════════════════════════════════════════════╣")
	fmt.Print("║ Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("║ Enter password: ")
	fmt.Scanln(&password)
	fmt.Println("╚═════════════════════════════════════════════════╝")
	//Employee Authentication
	loggedInEmployee, err := employeesDB.Login(username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}
	//Checking employee is admin or not.
	if employeesDB.IsAdmin(loggedInEmployee) {
		adminOperations(*employeesDB)
	} else {
		nonAdminOperations(*employeesDB, loggedInEmployee)
	}
}
