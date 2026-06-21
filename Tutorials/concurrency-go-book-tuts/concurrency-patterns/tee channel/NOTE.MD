## The Tee channel

* For a channel called Commands if 2 goroutines are reading from it only 1 will get a value 
eg commands: 1,2,3,4
G1: 1,3
G2: 2,4 
Each vlaue can only go to one receiver
This is a fundamental property of Go

* Tee pattern Gaurentees that both can consume all the values

* Working: 
* Say we have commands login, delete and consumers Executor and Auditor
- as long as  login is not receved by both Executor and Auditor the delete will not be recieved by any other consumer

* Channel mutliplexes between the vlaue for writing to two other channels
* It writes to any possible and set's it's vlaue to nil
* if a value is nil nothing can be written to it. SO if 1 was written again 1 cannot be written
* After 1 is written to both channels then and only then can we write 2.


### Don't add done until you have a goroutine that can outlive its usefulness.
### Done channel = A Signal for "I don't care anymore"
### Finished Channel  = A Signal of Completion

