package main

func main() {

	var or func(channels ... <-chan any) <-chan any
	or = func(channels ... <-chan any) <-chan any {
		switch len(channels) {
		case 0: return nil
		case 1: return channels[0]
		}

		orDone := make(chan any)

		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2: select {
			case <-channels[0]:
			case <- channels[1]:
			}
			default: 
				select {
				case <- channels[0]:
				case <- channels[1]:
				case <- channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			
			}
		}()
		return orDone
	}
}