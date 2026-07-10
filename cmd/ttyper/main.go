package main 

import(
	"fmt"
	"log"
	"github.com/Adarshrai24/ttyper/internal/ui"
	"github.com/Adarshrai24/ttyper/internal/data"
	"github.com/Adarshrai24/ttyper/internal/test"
	"github.com/Adarshrai24/ttyper/internal/storage"
)

func main() {
	db, err := storage.Open()
	if err != nil {
		log.Fatal(err) 
	}
	defer db.Close()
	for {
		fmt.Println("---------------Welcome to ttyper------------------")

		choice := ui.ShowMainMenu()
		
		if choice == 1 {
			timeChoice := ui.ShowTimeMenu()
			passageset := data.RandomParagraphPick()
			result := test.Start(passageset.Passages, passageset.Current, timeChoice)
			fmt.Println("GWPM: ", result.GWPM)
			fmt.Println("NWPM: ", result.NWPM)
			fmt.Println("Accuracy: ", result.Accuracy)
			fmt.Println("TimeFormat: ", timeChoice)
			err := storage.SaveSession(db, result, timeChoice)
			if err != nil {
				log.Fatal(err)
			}
		} else if choice == 2 {
			history, err := storage.GetHistory(db)
			if err != nil {
				log.Fatal(err)
			}
			ui.PrintHistory(history)
		} else if choice == 3 {
			// for starting I am for now ignoring errors will handle later properly
			durations := []int{15, 30, 60, 120}

			var maxStats []storage.Stats
			var avgStats []storage.Stats
			
			for _, d := range durations {
				maximum, _ := storage.MaximumStats(db, d)
				average, _ := storage.AverageStats(db, d)

				maxStats = append(maxStats, maximum)
				avgStats = append(avgStats, average)
			}	
			ui.PrintMaxStats()
			ui.PrintStats(maxStats)
			ui.PrintAvgStats()
			ui.PrintStats(avgStats)
		} else {
			break
		}
	}
}
