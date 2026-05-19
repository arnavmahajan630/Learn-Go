# Building blocks of  Concurrency in Golang

## Goroutines

### Introduction
* invoked with the go keyword. direct or as anonymous assigned / unassigned functions
* Green Threads are scheduled by a VM. os see's like it's running some go runtime stuff. but it's actually our green threads multiplexed on the os threams by runtime.
* Goroutines are Fully Preemptive Green Threads
* Goroutines are super lightweight (few kbs) comapred to OS threads. can spin up thousands in same address space
* The GC does not clean zombie goroutines.

### fork-join model

* a process is forked() fro mthe main as child and then needs to meet back with the parent.
* sync.Waitgroup is one of the ways to join a child back with it's parent

### Escape Analysis
* This is a feature of the go runtime which decides if a variable should be moved from stack to heap.
* Closures take the variables in lexical scope and put's it on heap if required. but this creates unnecesary instances of the variable on heap. instead of possible we can refference the lexical variables this avoids excess allocations improving performance. (execution time of gorouties is nlon deterministic)

## Sync Package
* The classic way of Concurrency using locks for memory synchronization

### WaitGroups
* To be used when we don't care about the result of a goroutine or have other means of gathering the result. if not then use channels instead.

### Mutex and RWMutex
* Stands for mutual Exclusion and is used to guard critical sections of programs from racing
* RWMutex means a goroutine can have access to reading a resource if another goroutine is not holding it for writing. meaning multiple writes can be performed parellelly.

### Cond
* a rendezvous point for goroutines waiting for or announcing the occurrence
of an event.
* sleep some goroutines unitl an annourcer routine does not deliver the message. this is done via sync.Cond
* Example: A worker checks tasks if there are no tasks it will call cond.wait(). This will unlock the mutex and put the goroutine to sleep.
* For notifying all the routines who are waiting for a task producer uses cond.Signal. this wakes the sleeping processes for task processing.
* Must always signal inside the critical section. This is imprtant as it might trigger the `Lost-wakeup problem`.
* The cond.Signal() wakes only one goroutine that has been sleeping for longest time. If we need multiple routines then we should call the Cond.Broadcast() This wakes all the sleeping routines registered to that mutex.

### Once
* `sync.Once` is not about doing something once in the program — it’s about doing it once when many goroutines may attempt it concurrently
* Use when events are triggered by demand
* Use Case: Lazy initialization where we can intialize something on demand
* once uses done, mutex after completion it marks as done and unlocks mutext so preemption causes no race conditions
* Sync.Once only executed once for the Instance of that specific Once
* TYPE: reentrant deadlock - goroutine tries to acquire a lock it already holds.
* If f calls Do again on the same Once instance, it will deadlock.
* even if there are two functions icnrement / decrement even if we do once.Do on either that once object is over and not valid for any of the other fucntions

### Pool
* Pool is a concurrent-safe implementation of the object pool pattern.
* Pool's `Get` will first check whether there are any available instances within the pool to return to the caller, and if not, call its New member variable to create a new one.
* After using the Caller calls the Put method to put the instance back in the Pool to be used by other prcoesses.
* `sync.Pool` is a temporary object reuse cache
* The GC does not gaurentee that the object state is saved and cleans it up. Should never use it for persistant state wide storage purposes. temporary buffers are the best.
* fmt writes to a buffer from pool and returns the buffer.
* we never test data being received from an object from the pool. always use reset to clean data.
* POINTERS TO REMEMBER PERTAINING TO POOL
    * When instantiating sync.Pool, give it a New member variable that is thread-safe
    when called.
    * When you receive an instance from Get, make no assumptions regarding the
    state of the object you receive back.
    * Make sure to call Put when you’re finished with the object you pulled out of the
    pool. Otherwise, the Pool is useless. Usually this is done with defer.
    * Objects in the pool must be roughly uniform in makeup.

## Channels
* channel serves as a conduit for a stream of information
* R_Only or W_Only chan means under that refference a routine is allwoed to perform only restricted action
* Unbuffered channels: 
    * don't have any storage. they need both processes to recieve and write at the same time. either waits in a queue and when other comes to life the waiting process is woken up and data is exchanged.
    * Multiple processes can share single unbuffered channel but 1-1 transaction always
    * A channel cannot read from an open stream using for range. 
    * To unblock multiple goroutines waiting on a chan we can just close it. they get auto unblocked. We did use cond but that was explicilty to unblock after certain activity. here whenever chanel closes we can read so it performs similar functionality

* Buffered channels:
    * perform max N writes where n is the size of the buffer.

* We need to differentiate between chan owners (writer) and chan utilizers (readers)
* chann1 shows creating a readonly view of a channel where producer is encapsulated.


## Select

* Just like we have wg for primitive select is a sync point for multiple channels
* IT blocks till one of the channels is ready and we can process
* an event multiplexer for channel operations. wait on multiple possible communication events simultaneously  
* It's expensive that waitgroups
* SELECT Block contains multiple cases much similar to switch case statements. but only 1 of ready among the is chosen pseudo randomly for processing.
* What if none of the routines are ready ? cancellations and timeouts exist for this reason
* How to do something else when all are blocking ? use the default clause
#### one-time event arbitration not constant monitoring

* but for select ? hehe boi now that's monitoring forever


## GOMAXPRCS LEVER

* controls the exact number of logical processors available to the Go runtime
* Tells the maximum number of OS threads that can run these
* Research more on the GMP Model where this sets the `P`
* For a Goroutine (G) to actually run, an OS thread (M) must acquire a logical processor (P). Therefore, if GOMAXPROCS is set to 4, there are only 4 Ps available. Even if you spin up 10,000 Goroutines, only 4 OS threads will be actively executing Go code at any given microsecond.

* Under ristricted infra such as rstrictions of logical CPU by docker or kubernetes we need to manually change this value.
