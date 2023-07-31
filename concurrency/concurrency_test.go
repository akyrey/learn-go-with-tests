package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// Mock used for testing
func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

// Stub used to benchmark the function
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("wanted %v, got %v", expected, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
    // Generate 100 random urls
    urls := make([]string, 100)

    for i := 0; i < len(urls); i++ {
        urls[i]= "a url"
    }

    // Reset the timer before actually running the benchmark
    b.ResetTimer()

    // Benchmark using the slow stub
    for i := 0; i < b.N; i++ {
        CheckWebsites(slowStubWebsiteChecker, urls)
    }
}
