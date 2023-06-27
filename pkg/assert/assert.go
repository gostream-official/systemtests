package assert

import "log"

func Expect(condition bool, failedMessage string) {
	if !condition {
		log.Fatalf("failed expectation: %s", failedMessage)
	}
}
