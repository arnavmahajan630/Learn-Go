package patterns

import "fmt"

func PatternCon()  {
		 charchan := make(chan string, 3) //buffered channel with only 3 chars
		// async communications are buffered
		mychars := []string{"a", "b", "c"}
		
		for _, c := range mychars {
			select {
			case charchan <- c:
				
			}
		}
		close(charchan)

		for result := range charchan {
			fmt.Println(result)
		} // looping over closed channel
}