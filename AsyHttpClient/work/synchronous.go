package work

import (
	"fmt"
	"time"

	"github.com/Naive-vs-Reactive/agency"
)

func synchronous() [10]agency.Infomation {
	start := time.Now()
	customer := agency.GetCustomers()
	destinations := agency.GetDestinations(customer)
	var infos [10]agency.Infomation
	for index, dest := range destinations {
		q := agency.GetQuoting(dest)
		w := agency.GetWeather(dest)
		infos[index] = agency.Infomation{W: w, D: dest, Q: q}
	}
	fmt.Println(time.Since(start))
	return infos
}
