## Timeouts and cancellations

### Reasons for having Timeouts

* System Saturation: if our system is saturated (i.e., if its ability to process requests is at capacity), we may want requests at the edges of our system to time out rather than take a long time to field them

* Stale Data: Man times we want our goroutine to produce result in a specific time window and if not time it out

* A timeout is an implicit cancellation / cancellation is one of the retults of Timeouts

### Reasons for having Cancellations

* User Intervention: Allow users to cancel concurrent operations they have started
* Parent Cancellation: Ofcourse a parent should cancel it's child
* Replicated Requests: Send data to multiple concurrent requests and when first responds we cancel the rest of them.

### How Cancellations Affect Downstream Consumers ?

* A Goroutine must be broken down into small atmoic operations that are >= desired cancellation requirements
* If a goroute modifies shared space db,file,cache and is cancelled should the changes rollback ? 

### Duplicate Message Problem
* A generator observers a reader to be non performant and launches second reader. The prior one can be slow and later reads same value read by the new reader. Now we have duplicate reading

* One solution is mark it with ID and store results in set based on that ID.
* We can use bidirectional communication with the parent to explicitly request permission to send the message. eg after A proceesse it asks parent should i send the message or not. Generator holds the authority

