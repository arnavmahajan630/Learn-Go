## Error Propogation

* An error is composed of the following properties:
- what happened: exact messasge log of what is the issue
- When and Where is occured: A Complete stack trace starting with how the call
was initiated and ending with where the error was instantiated. The error should contain information regarding the context it’s running within. This pertaints to the where eg machine on which it  occured. Additionally Time in UTC
- A Friendly User Facing Message: abbrevated and relevant information from the above two points.Must be human centric
- How the user can get more information: Must provide an error ID So More details reagrding the error can be discovered by running the Error message against the Log File. It is helpful to include a hash of the stack trace to aid in aggregating like issues in bug trackers.eg. Hashing the stack trace is better so we can count the occurences of the same issue. No Duplicaiton 

### Categories for Errors

* Bugs:raw errors that are not customized to the system. An error can become a bug because it might leak Logic
* Known Edge Cases: Errors that are designed for.


- Any error leaving the application boundary should be represented as an AppError (or some equivalent application-level error) before it reaches the user.

## err Error
* error.As: It traverses the Erorrchain and finds the correct level of error and extracts the message required for display. 
* error.Is: walks the propogation chain and returns a boolean value when compared with a specific error
    - eg: error.Is(err, os.ErrorNotExisit)

