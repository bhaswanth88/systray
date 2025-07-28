package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Order Test")
	systray.SetTooltip("Testing Menu Order")

	// Test ordering by adding items with specific order values
	item3 := systray.AddMenuItemAtIndex("Third Item", "Should appear third", 3)
	item1 := systray.AddMenuItemAtIndex("First Item", "Should appear first", 1)
	item2 := systray.AddMenuItemAtIndex("Second Item", "Should appear second", 2)

	// Add a regular item (should appear last with default order)
	itemDefault := systray.AddMenuItem("Default Order", "Should appear last")

	// Test changing order of existing item
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Changing order of Third Item to 0 (should move to first)")
		item3.SetOrder(0)
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the test app")

	// Handle clicks
	go func() {
		for {
			select {
			case <-item1.ClickedCh:
				fmt.Println("First Item clicked")
			case <-item2.ClickedCh:
				fmt.Println("Second Item clicked")
			case <-item3.ClickedCh:
				fmt.Println("Third Item clicked")
			case <-itemDefault.ClickedCh:
				fmt.Println("Default Order Item clicked")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	fmt.Println("Test finished")
}