# Hello to Concurrency in GO!

* Atomic Operaitons are “indivisible” and “uninterruptible.”
* Context is super important as a process might be Atomic in context of a group of process and not others

* An operation i++. fetch update store can be atomic in it's own context but all 3 as a whole in os's context it cannot be as preemption can happen
