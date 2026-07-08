package main 

import(
	"fmt"
	"github.com/Adarshrai24/ttyper/internal/ui"
)

func main() {
	for {
		fmt.Println("---------------Welcome to ttyper------------------")
		
		choice := ui.ShowMainMenu()

		if choice == 1 {
			timeChoice := ui.ShowTimeMenu()
			if timeChoice >= 0 {
				break;
			}
		} else if choice == 2 {

		} else if choice == 3 {

		} else {
			break
		}
	}
}
