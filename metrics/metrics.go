// handles statsd related things
package metrics

import (
	"fmt"
	statsd "github.com/etsy/statsd/examples/go"
	"time"
)

var (
	client *statsd.StatsdClient
)

func Connect() {
	host := `127.0.0.1`
	port := 8125
	fmt.Printf("Configured to send to statsd at %s:%d\n", host, port)
	client = statsd.New(host, port)
}

func Time(name string, start time.Time) {
	if client == nil {
		return
	}
	now := time.Now()
	duration := now.Sub(start)
	if duration > 5000*time.Millisecond {
		fmt.Printf("Latent measurement for %q: %s", name, duration)
	}
	milliseconds := int64(duration / time.Millisecond)
	toStatsd(func() {
		// record the duration
		client.Timing(name, milliseconds)
		// also record a count
		client.UpdateStats([]string{fmt.Sprintf("%s.count", name)}, 1, 1)
	})
}

func toStatsd(fn func()) {
	start := time.Now()
	fn()
	duration := time.Now().Sub(start)
	if duration > 250*time.Millisecond {
		fmt.Printf("Statsd time took %s", duration)
	}
}
