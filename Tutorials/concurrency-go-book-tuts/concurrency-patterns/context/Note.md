## The Context Package

* The context Package is used to manage the children goroutines so that they can be cancelled / preempted by the parent
* It provides WthCancel, Deadline and Timeouts i.e extention over done chan for managemnet of goruouties
* One a parent can preempt a child and not vice verca

### Signatures
* func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    - Returns a new context that Closees it's done channel after the returned Cancel() is called 
* func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
    - Closes it's done channel when the machine’s clock advances past the given deadline
* func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    - Returns a new context that clsoes it's done channel after the set timeout
* func WithValue(parent Context, key, val interface{})
    - Attaches the request metadata to the context instead of messy arguemnts

#### Never store instances of context.Context as variables but always pass them through functions

### Main's Context
* func Background() Context
    - Background returns empty Context
* func TODO() Context
    - Also returns empty which is not to b  e used in prod. it's temp context till you are waiting yet to recieve one ?

