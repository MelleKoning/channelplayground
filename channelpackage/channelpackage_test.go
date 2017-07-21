package channelpackage

import (
	"fmt"
	"testing"
)

// example taken from
// https://blog.golang.org/pipelines
func TestChannelGen_SquareValues(t *testing.T) {
	// Melle: producer or generator of a pipeline that sends a number of values on to a channel
	// in this example we just put four integers on the channel
	in := gen(3, 4)
	// with 'in' we now have a handle on the channel for the generator

	results := sq(in)

	fmt.Println(<-results)
	fmt.Println(<-results)
	fmt.Println(<-results)
	fmt.Println(<-results)

	//results := sq(sq(in))

	for n := range results {
		fmt.Println(n)
	}

	for n := range results {
		fmt.Println(n)
	}

}

func TestChannelGen_SquareValues_Twice(t *testing.T) {
	// producer or generator of a pipeline that sends a number of values on to a channel
	// in this example we just put four integers on the channel
	in := gen(3, 4)

	// with 'in' we now have a handle on the channel for the generator
	results := sq(in)
	// the results are also <- chan int, so we can square those again,
	// as we do in the for range loop below

	for n := range sq(results) {
		fmt.Println(n) // exports 3*3 * 3*3 (81) and 4*4 * 4*4 (256)
	}

}

func TestGenerateTags(t *testing.T) {

	// These would be incoming tags
	tag1 := &Tag{Key: "one", Value: 1}
	tag2 := &Tag{Key: "two", Value: 2}
	tagChannel := generateTags(*tag1, *tag2)
	//tagChannel := generateTags()

	// we now have a tagChannel*/
	results := processTagChannel(tagChannel)

	for n := range results {
		fmt.Println(n.RequestKey)
	}
}

func TestSendTags_GetResponsesOnChannel(t *testing.T) {
	tag1 := &Tag{Key: "one", Value: 1}
	tag2 := &Tag{Key: "two", Value: 2}

	requestsChannel := generateTags(*tag1, *tag2)

	resultChannel := make(chan TagResult) // buffer size 2, set it or not
	requestsSent := processTagChannel(requestsChannel)

	fmt.Println("start awaiting results")
	awaitResults(requestsSent, resultChannel)
	fmt.Println("done  awaiting results")

	for n := range resultChannel {
		fmt.Println(n.Result)
	}

}
