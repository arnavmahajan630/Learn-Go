## The Or-Channel

* combine one or more done channels into a single done channel that closes if any of its component channels close.
* We can write a select but if we don't know the number of Or channels we are working with then it's better to write the Or Channel Pattern

* creates a composite done channel through recursion and goroutines
