package main

import (
	"fmt"
	"time"

	"github.com/gookit/event"
)

func main() {
	TestGookitEvent()
	TestSyncGookitEvent()
}

func TestGookitEvent() {
	// Register event listener
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle normal event: %s\n", e.Name())
		return nil
	}), event.Normal)

	// Register multiple listeners
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle high event: %s\n", e.Name())
		return nil
	}), event.High)

	// ... ...

	// Trigger event
	// Note: The second listener has a higher priority, so it will be executed first.
	event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})

	fmt.Println("TestGookitEvent done")
}

func TestSyncGookitEvent() {
	defer event.CloseWait()

	// Register event listener
	event.On("app.evt_async", event.ListenerFunc(func(e event.Event) error {
		time.Sleep(1 * time.Second)
		fmt.Printf("handle normal event: %s\n", e.Name())
		return nil
	}), event.Normal)

	// Register multiple listeners
	event.On("app.evt_async", event.ListenerFunc(func(e event.Event) error {
		time.Sleep(1 * time.Second)
		fmt.Printf("handle high event: %s\n", e.Name())
		return nil
	}), event.High)

	// ... ...

	// Trigger event
	// Note: The second listener has a higher priority, so it will be executed first.
	event.Async("app.evt_async", event.M{"arg0": "val0", "arg1": "val1"})

	fmt.Println("TestSyncGookitEvent done")
}
