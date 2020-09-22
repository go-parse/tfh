
# tfh - getting data about horse racing
 
## Examples

###### **Parsing the results page**

```Go
package main

import "github.com/go-parse/tfh"

func main() {

    path := "/horse-racing/result/chelmsford-city/2020-09-20/1300/28/1"

    proxy := "95.174.67.50:18080"

	race := tfh.GetResultByPath(path, proxy)

	fmt.Println("Datetime", race.Datetime)
	fmt.Println("Distance", race.Distance)
	fmt.Println("Class", race.Class)
	fmt.Println("Title", race.Title)
	fmt.Println("Currency", race.Currency)
	fmt.Println("Winner", race.Winner)
	fmt.Println("RaceType", race.RaceType)
	fmt.Println("Category", race.Category)
	fmt.Println("Surface", race.Surface)
	fmt.Println("Racecourse", race.Racecourse)
	fmt.Println("Going", race.Going)
	fmt.Println("Types", race.Types)

	for _, runner := range race.Runners {

		fmt.Println("Isp", runner.Isp)
		fmt.Println("Wgt", runner.Wgt)
		fmt.Println("Number", runner.Number)
		fmt.Println("Draw", runner.Draw)
		fmt.Println("Age", runner.Age)
		fmt.Println("Rating", runner.Rating)
		fmt.Println("HorseID", runner.HorseID)
		fmt.Println("JockeyID", runner.JockeyID)
		fmt.Println("TrainerID", runner.TrainerID)
		fmt.Println("Pos", runner.Pos)
		fmt.Println("Horse", runner.Horse)
		fmt.Println("Jockey", runner.Jockey)
		fmt.Println("Trainer", runner.Trainer)
		fmt.Println("")
    }
}
```

###### **Parsing the results for the whole day.**

```Go
func main() {

	date := time.Date(2019, time.January, 30, 0, 0, 0, 0, time.UTC)

	proxy := "95.174.67.50:18080"

	for _, race := range tfh.GetResultByDate(date, proxy) {

		fmt.Println("Datetime", race.Datetime)
		fmt.Println("Distance", race.Distance)
		fmt.Println("Class", race.Class)
		fmt.Println("Title", race.Title)
		fmt.Println("Currency", race.Currency)
		fmt.Println("Winner", race.Winner)
		fmt.Println("RaceType", race.RaceType)
		fmt.Println("Category", race.Category)
		fmt.Println("Surface", race.Surface)
		fmt.Println("Racecourse", race.Racecourse)
		fmt.Println("Going", race.Going)
		fmt.Println("Types", race.Types)

		for _, runner := range race.Runners {

			fmt.Println("Isp", runner.Isp)
			fmt.Println("Wgt", runner.Wgt)
			fmt.Println("Number", runner.Number)
			fmt.Println("Draw", runner.Draw)
			fmt.Println("Age", runner.Age)
			fmt.Println("Rating", runner.Rating)
			fmt.Println("HorseID", runner.HorseID)
			fmt.Println("JockeyID", runner.JockeyID)
			fmt.Println("TrainerID", runner.TrainerID)
			fmt.Println("Pos", runner.Pos)
			fmt.Println("Horse", runner.Horse)
			fmt.Println("Jockey", runner.Jockey)
			fmt.Println("Trainer", runner.Trainer)
			fmt.Println("")
		}
	}
}
```