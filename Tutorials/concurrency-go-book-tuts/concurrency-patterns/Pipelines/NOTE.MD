## Pipelines

* Pipeline in Go is a tool that we can use to perform abstraction in our system.
* used when we need to process streams and batches of data
* We take in a series of dat -> process it ->   and pass the data out.
* each op is a stage in the pipeline.
* Independent stages are easy to maintain and we can handle concurrency gracefully

### Properties of Pipelies
* A stage consumes and returns the same type.
* A stage must be reified2 by the language so that it may be passed around. Functions in Go are reified and fit this purpose nicely

#### What are monads ? later let's look into this

* it's always better to modify functions i.e our designed pipelines instead of procedural code for the change required

### Channels as Pipelines
