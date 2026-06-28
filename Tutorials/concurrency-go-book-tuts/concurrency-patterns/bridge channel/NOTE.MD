## The Bridge Channel

* For a system producing multiple channels it's better to send through another channel
Visual

chanStream
    |
    +----> stream1
    |
    +----> stream2
    |
    +----> stream3


SO: <-chan <-chan int is a channel that deliver's other channels

The problem is that without a bridge we use double for loop to iterate 
* A bridge falttens a  `<-chan <-chan int` into a `<-chan int` i.e a single continuous stream
* Multiple workers generate their own channels while a common consumer is waiting to easily read the data
* It's essentiall a `vector<vector<Type T>> ` being condensed to `vector<Type T>` synchronously in a concurrent environment