## HeartBeats

* Heartbeats are a way for concurrent processes to signal life to outside parties.
* They can occur at a specific time interval or beginning of a unit of work 

* Essentially heartbeats are a communication channel that pings it's caller with "alive"

Using default block for heartbeats is very important. 
- If a there is no reader to read the unbuffered heartbeak worker might block. we need to skip the pulse. Default makes that possible.
- Pulse is only a metric for observation and missing is fine as long as we get another one timely.
- But blocking for someone to read the hearbeat would be hazardous.

worker dies
     |
     v
no heartbeat
     |
     v
timeout
     |
     v
failure detected

