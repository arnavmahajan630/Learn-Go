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

### Once
* `sync.Once` is not about doing something once in the program — it’s about doing it once when many goroutines may attempt it concurrently
* Use when events are triggered by demand
* Use Case: Lazy initialization where we can intialize something on demand
* once uses done, mutex after completion it marks as done and unlocks mutext so preemption causes no race conditions
* Sync.Once only executed once for the Instance of that specific Once
* TYPE: reentrant deadlock - goroutine tries to acquire a lock it already holds.
* If f calls Do again on the same Once instance, it will deadlock.

### Pool
* 