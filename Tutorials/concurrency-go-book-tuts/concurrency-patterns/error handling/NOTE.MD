## Error Handling

* In Concurrent Systems we should send errors to another part of our program that has complete information
about the state of our program, and can make a more informed decision about what to do.

* Here we created a result struct and held the error and passde it to the caller and printed it from there. 
* Since we never handle the erorr in the goroutine but it's handled by the producer we have more control an context
