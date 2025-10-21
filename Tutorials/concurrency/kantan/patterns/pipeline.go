package patterns

import "fmt"

/*
 => break problem into different stages
  start to end depends on the main func
*/

/*
Since channels are unbuffered they can only send 1 val and thy block
How the code works is as follows := 

# We call slice to data with nums and return a channel
# We then make the channel to return
# Then we fire off a anon function that iterates over the nums and writes each int to out channel
Well when it writes the first int it is going to go in the out and out is not full so the routine will block till it's freed
So once that happens Slice to data is ofcourse not waiting for anon so it retuns out
# We now have the first int in the datachan variable and we call sq which takes this chan
# Now we iterate over this chan and sqaure the int (only 1 present) and read it to new chan
# Now the out has been read so the SliceToData can modify it
# We return the sq int to finalchan and now we need to wait till we recieve a new int to process it

*/

func SliceToData(nums []int) <- chan int {
	out := make(chan int)
	go func(){
		for _, x := range nums {
		out <- x
		}
		close(out)
	}()
	return out
}


func Sq(in <- chan int) <-chan int {
	out := make(chan int)
	go func(){
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}


func Pipleline() {
	// input
	nums := []int{1,2,3,4,5,6}

	// stage 1 {get each element in data}
	datachan := SliceToData(nums) 
	
	// stage 2 {pass data to another stage}
	finalchan := Sq(datachan);
	// stage 3 {output the result}
	for n := range finalchan {
		fmt.Println(n)
	}
}