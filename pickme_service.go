package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	type Customer struct {
		Id            int
		Name          string
		Mobile_No     string
        Location      string
        Availability  bool
	}
  type Driver struct {
    Id            int
    Name          string
    Vehicle_No    string
    Availability  bool
    Location      string

  }
	router := gin.Default()

	// GET a Customer detail
	router.GET("/customer/:id", func(c *gin.Context) {
		var (
			customer Customer
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, name, mobile_no, location, availability from customer where id = ?;", id)
		err = row.Scan(&customer.Id, &customer.Name, &customer.Mobile_No, &customer.Location, &customer.Availability)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": customer,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
  // GET all customers
	router.GET("/customer", func(c *gin.Context) {
		var (
			customer  Customer
			customers []Customer
		)
		rows, err := db.Query("select id, name, mobile_no, location, availability from customer;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&customer.Id, &customer.Name, &customer.Mobile_No, &customer.Location, &customer.Availability)
			customers = append(customers, customer)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": customers,
			"count":  len(customers),
		})
	})
  // GET a driver detail
  router.GET("/driver/:id", func(c *gin.Context) {
    var (
      driver Driver
      result gin.H
    )
    id := c.Param("id")
    row := db.QueryRow("select id, name, vehicle_no, availability, location from driver where id = ?;", id)
    err = row.Scan(&driver.Id, &driver.Name, &driver.Vehicle_No, &driver.Availability, &driver.Location)
    if err != nil {
      // If no results send null
      result = gin.H{
        "result": nil,
        "count":  0,
      }
    } else {
      result = gin.H{
        "result": driver,
        "count":  1,
      }
    }
    c.JSON(http.StatusOK, result)
  })
  // GET all drivers
  router.GET("/driver", func(c *gin.Context) {
    var (
      driver  Driver
      drivers []Driver
    )
    rows, err := db.Query("select id, name, vehicle_no, availability, location from driver;")
    if err != nil {
      fmt.Print(err.Error())
    }
    for rows.Next() {
      err = rows.Scan(&driver.Id, &driver.Name, &driver.Vehicle_No, &driver.Availability, &driver.Location)
      drivers = append(drivers, driver)
      if err != nil {
        fmt.Print(err.Error())
      }
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "result": drivers,
      "count":  len(drivers),
    })
  })
  // GET driver availability
  router.GET("/driveravailability/:id", func(c *gin.Context) {
    var (
      availability bool
      result gin.H
    )
    id := c.Param("id")
    row := db.QueryRow("select availability from driver where id = ?;", id)
    err = row.Scan(&availability)
    if err != nil {
      // If no results send null
      result = gin.H{
        "result": nil,
      }
    } else {
      result = gin.H{
        "result": availability,
      }
    }
    c.JSON(http.StatusOK, result)
  })
  // GET customer availability
	router.GET("/customeravailability/:id", func(c *gin.Context) {
		var (
			availability bool
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select availability from customer where id = ?;", id)
		err = row.Scan(&availability)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
			}
		} else {
			result = gin.H{
				"result": availability,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	router.Run(":3000")
}
