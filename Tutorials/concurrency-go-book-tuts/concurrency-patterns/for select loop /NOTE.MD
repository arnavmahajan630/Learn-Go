### The For Select Loop

* forver loop over select i.e keep monitoring the channels forever for activity

* FS.Go implements a cancellable goroutine
- Idea is we want the goroutine to keep working but also get cancelled when the parent wants it to
- we create an unbuffered channel and idea is we wait till someone closes this channel or writes to it.
- select allows us to this while we  goroutine can keep working
- We try to read from the nil and this would block forever so this goroutine keeps running
- we launch a routine that closes done so this means that done is now ready so all those waiting for done wake up
- Then are able to cancel this routine.
- <-terminated means we are waiting to read from terminated so we are blocking here till terminated becoems available this waiting. this is similar to wg.wait() but thas was a primitive and this is modern channels based communication
- This patterns allows us manual cancellation