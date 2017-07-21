package channelpackage

import (
	"fmt"
	"time"
)

// generates a list of ints on a channel that can be consumed
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// squaring the ints on a given int channel and puts out the results on another int channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Generates tags that can be read from the Tag channel
func generateTags(tags ...Tag) <-chan Tag {
	out := make(chan Tag)
	go func() {
		for _, n := range tags {
			fmt.Println("generate tag ", n.Key)
			out <- n
		}
		close(out)
	}()
	return out
}

// processes a tag channel and generates tagResponses for each of the Tag
func processTagChannel(in <-chan Tag) <-chan TagResponse {
	out := make(chan TagResponse)
	go func() {
		for n := range in {
			fmt.Println("SendTagRequest for ", n.Key)
			tagResponse := SendTagRequest(n)
			out <- tagResponse
		}
		close(out)
	}()
	return out
}

// awaitResults takes a TagResponse channel and will emit TagResult structs on the results channel
func awaitResults(req <-chan TagResponse, results chan TagResult) {
	go func() {
		// put some results on the TagResults channel
		for n := range req {

			// just sleep for one second
			// consider this is another method being triggered
			// that puts results on the results channel,
			// which only works if the channel was setup
			// to contain the proper buffer size
			// otherwise the flow is blocked.
			time.Sleep(1 * time.Second)

			fmt.Println("a result is received..")
			aresult := &TagResult{Result: n.RequestKey}
			results <- *aresult // blocks for a result to arrive, or blocks when buffer size not set for channel
		}
		close(results)
	}()
}

// Tag is just a dummy struct
type Tag struct {
	Key   string
	Value int
}

// TagResponse is a dummy response structure
type TagResponse struct {
	RequestKey string
}

// TagResult ...
type TagResult struct {
	Result string
}

// SendTagRequest is a dummy Call for a tag and a response
func SendTagRequest(in Tag) TagResponse {
	tagResponse := &TagResponse{RequestKey: in.Key}
	return *tagResponse
}
