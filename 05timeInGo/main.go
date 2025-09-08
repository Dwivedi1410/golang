package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is a time section")

	currentTime := time.Now()
	fmt.Println(currentTime) //2025-09-08 11:50:42.7069195 +0530 IST m=+0.000521201

	presentTime := currentTime.Format("2006-01-02 Monday 15:04:05")
	fmt.Println(presentTime)   //2025-09-08 Monday 11:55:29

	pastDate := time.Date(2022, time.Month(05) , 22 , 12, 34, 20, 50, time.UTC)
	fmt.Println(pastDate)  //2022-05-22 12:34:20.00000005 +0000 UTC
	fmt.Println(pastDate.Format("2006-01-02 Monday 15:04:05"))  //2022-05-22 Sunday 12:34:20

}
