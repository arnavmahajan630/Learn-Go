## Queuing

* Process of accepting work for pipeline even if it's not ready is called queueing
* Once a stage has completed some work. it stores it in a temp location in memory so that other stages can retrieve it later and the stage does not need to hold refference to it.

* Adding queueing is not really good and it one of the last optimizations we can think about
* It can pre‐maturely can hide synchronization issues such as deadlocks and livelocks,

* Queueing never will increase the runtime performance of the program
* A queue Does decopouling but cannot buy speed 


### Trafeoffs for queues

* Improve Utilization but introduce Latency
* consume memory

### Little's Law
L=λW ; work in system = arrival rate * Time in system


Problem statement: 
1 Cashier (takes orders)
2 Baristas (make coffee)
Customers arrive randomly.
Cashier is fast and Barista's are slow

