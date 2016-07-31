package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pickme")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS `customer` (`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(50) NOT NULL DEFAULT '0',`mobile_no` varchar(50) NOT NULL DEFAULT '0',`location` varchar(50) NOT NULL DEFAULT '0',`availability` tinyint(4) NOT NULL DEFAULT '0',PRIMARY KEY (`id`));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("customer Table successfully migrated....")
	}

  stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS `driver` (`id` int(11) NOT NULL AUTO_INCREMENT, `name` varchar(50) NOT NULL, `vehicle_no` varchar(50) NOT NULL, `availability` tinyint(4) NOT NULL, `location` varchar(50) NOT NULL, PRIMARY KEY (`id`));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("driver Table successfully migrated....")
	}

  stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS `customer_driver` (`id` int(11) NOT NULL AUTO_INCREMENT, `customer_id` int(11) NOT NULL, `driver_id` int(11) NOT NULL, `time_picked` datetime NOT NULL, `picked_location` varchar(50) NOT NULL, `destination` varchar(50) NOT NULL, `distance` double NOT NULL, `status` varchar(50) NOT NULL, PRIMARY KEY (`id`));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("customer_driver Table successfully migrated....")
	}

  stmt, err = db.Prepare("INSERT INTO `customer` (`id`, `name`, `mobile_no`, `location`, `availability`) VALUES(1, 'ashan', '754543742', 'kottawa', 1),	(2, 'chamitha', '321341242', 'nugegoda', 0);")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("insert cutomer values ....")
	}

  stmt, err = db.Prepare("INSERT INTO `driver` (`id`, `name`, `vehicle_no`, `availability`, `location`) VALUES(1, 'chamapaka', 'abc-2345', 1, 'panadura'),(2, 'saranga', 'cde-3345', 0, 'panadura');")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("insert driver values ....")
	}
}
