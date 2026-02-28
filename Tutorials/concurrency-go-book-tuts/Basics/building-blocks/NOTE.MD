# Building blocks of  Concurrency in Golang

## Goroutines

### Introduction
* invoked with the go keyword. direct or as anonymous assigned / unassigned functions
* Green Threads are scheduled by a VM. os see's like it's running some go runtime stuff. but it's actually our green threads multiplexed on the os threams by runtime.
* Goroutines are Fully Preemptive Green Threads
* Goroutines are super lightweight (few kbs) comapred to OS threads. can spin up thousands in same address space
* The GC does not clean zombie goroutines.

### for-join model

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
* 
