package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/redis/go-redis/v9"
)


func main() {
    client := redis.NewClient(&redis.Options{
		//TODO: This options should be environment variables instead of hardcoded
        Addr:     "localhost:6379",
        Password: "", 
        DB:       0,  
    })

	for {
		//TODO: Validate x and y
		//TODO: Add way to convert from grid to points
		var x, y float64
   		fmt.Print("Type x: ") 
		fmt.Scan(&x)
   		fmt.Print("Type y: ") 
		fmt.Scan(&y)

		val, err := client.Get(context.Background(), fmt.Sprintf("%.4f,%.4f", x, y)).Result()
		if err == redis.Nil {
			fmt.Println("key does not exist")
			resp, err := http.Get(fmt.Sprintf("https://api.weather.gov/points/%.4f,%.4f", x, y))
			if err != nil {
				fmt.Printf("Http get failed: %s", err)
				return
			}
			responseData, err := io.ReadAll(resp.Body)
		    if err != nil {
				fmt.Printf("ReadAll failed: %s", err)
				return
    		}
			//TODO: Write case when points are not found in website
   			fmt.Println(string(responseData))
			//TODO: Look for properties.forecast url and then return geometry.periods[0].shortForecast

			err = client.Set(context.Background(), fmt.Sprintf("%.4f,%.4f", x, y), string(responseData), 0).Err()
			if err != nil {
				fmt.Printf("failed to set key")
				return
			}
		} else if err != nil {
			fmt.Println("Get failed", err)
			return
		} else {
			fmt.Println("Key found")
			fmt.Println(val)
		}
		
	}

}