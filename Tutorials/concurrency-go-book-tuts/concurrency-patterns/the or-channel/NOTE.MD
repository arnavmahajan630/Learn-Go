## The Or-Channel

* combine one or more done channels into a single done channel that closes if any of its component channels close.
* We can write a select but if we don't know the number of Or channels we are working with then it's better to write the Or Channel Pattern

* creates a composite done channel through recursion and goroutines
* Motive is to create a Channel such that chan C depends of X , y, z any one closes then the main channel closes

* say we want to unblock when any of the channel closes we do have select
* but if we have variadic ? 

### Explanaition for Or.go

* we create a function signature called Or that takes variadic channels and returns 1 reader channel 
* we have the switch statement this is for base case
    - case 0 means there are no channels so we can return 0
    - case 1 means only 1 channel exists so we can return that chan i.e that is the signal
* For channels if greater than 2 we create an Or done channel. the main chan (C)
* The we lanuch go func() defer close(orDone). we wait for events from other channels and if they close we close this main channel 
* if we have exactly 2 channels we wait for either to finish first so we can close
* under the default we handlde the 2+ Case
    - Wait for 0,1,2 or we make the recursive call
    - <-or(append(channel[3:0], orDone) ...) => This means say we started with or(a,b,c,d,e,f,g) then we wait for a, b, c then call another or(d,e,f, orDone) this orDone is shared everywhere
    - then d,e,f we wait for and we call gor(g, orDone). 
    - Then we wait for g and orDone meaning `is the parent signalling or is g getting signalled`
* Signaling : Now this creates a single linear tree or(a,b,c,orDone1) -> or(d,e,f,orDone2) -> or(g, orDone3)
* each is waiting for some orDone and each function is monitoring 3 channels + an orDone channel
* The moment any of these fumble say Done1 that specific routine closes it. and Done1 get's signalled since it closed. 
* Here we play with 2 channels. 1. inherited from parent 2. passed to child for cancellation
