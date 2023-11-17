package notifier

import "fmt"

type Notifier interface {
	Notify(ip string) error
    CheckEnv() error
    Auth() error
}

// Will return a new notifier
func CreateNotifierBasedOnInput(input string) (Notifier, error) {
	var class Notifier
	switch input {
	case "webhook":
		class = &WebhookNotifier{}
	default:
		return nil, nil 
	}

	if err := class.CheckEnv(); err != nil {
		return nil, fmt.Errorf("error while validating provider (%s) environment: %s", input, err)
	}

	if err := class.Auth(); err != nil {
		return nil, fmt.Errorf("error while trying to auth: %s", err)
	}

	return class, nil
}
