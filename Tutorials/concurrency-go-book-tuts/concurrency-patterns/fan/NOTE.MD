## Fan In Fan Out

* Fan-out is a term to describe the process of starting multiple goroutines to handle
input from the pipeline
* Fan-in is a term to describe the process of combining
multiple results into one channel.

When to Fan out our pipeline ?
- It doesn’t rely on values that the stage had calculated before.
- It takes a long time to run.

#### Fan Out 
* create an optimal number of channels with propotion to the number of CPU's
* Launch concurrent generators taking done and the reader channel


