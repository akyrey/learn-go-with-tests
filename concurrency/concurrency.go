package concurrency

type (
	WebsiteChecker func(string) bool
	result         struct {
		// anonymous values, just declaring the type
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// Channel of result
	resultChannel := make(chan result)

	for _, url := range urls {
		// We often use anonymous functions when we want to start a goroutine
		// We need to bind the parameter of the anonymous function or they'd all use the same one
		go func(u string) {
			// We "send" the result to the channel (send statement: channel on the left and value on the right)
			// This operation is in parallel inside its own process
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression (variable on the left, channel on the right)
		// Each result is taken out one at a time, avoiding writing to the map concurrently
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
