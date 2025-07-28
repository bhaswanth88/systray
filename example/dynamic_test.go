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
	systray.SetTitle("Dynamic Menu Test")
	systray.SetTooltip("Testing Dynamic Menu Operations")

	// Create initial menu items
	item1 := systray.AddMenuItemAtIndex("Item 1", "First item", 1)
	item2 := systray.AddMenuItemAtIndex("Item 2", "Second item", 2)
	item3 := systray.AddMenuItemAtIndex("Item 3", "Third item", 3)
	
	systray.AddSeparator()
	
	// Add control items
	swapBtn := systray.AddMenuItem("Swap 1&2", "Swap items 1 and 2")
	moveBtn := systray.AddMenuItem("Move 3 Before 1", "Move item 3 before item 1")
	deleteBtn := systray.AddMenuItem("Delete Item 2", "Delete the second item")
	reorderBtn := systray.AddMenuItem("Reorder All", "Reorder all items")
	clearBtn := systray.AddMenuItem("Clear All", "Clear all menu items")
	
	systray.AddSeparator()
	quitBtn := systray.AddMenuItem("Quit", "Exit the test")

	// Track deleted state
	item2Deleted := false

	// Handle menu interactions
	go func() {
		for {
			select {
			case <-item1.ClickedCh:
				fmt.Printf("Item 1 clicked (order: %d)\n", item1.GetOrder())
			case <-item2.ClickedCh:
				if !item2Deleted {
					fmt.Printf("Item 2 clicked (order: %d)\n", item2.GetOrder())
				}
			case <-item3.ClickedCh:
				fmt.Printf("Item 3 clicked (order: %d)\n", item3.GetOrder())
				
			case <-swapBtn.ClickedCh:
				if !item2Deleted {
					fmt.Println("Swapping items 1 and 2...")
					systray.SwapMenuItems(item1, item2)
					fmt.Printf("After swap - Item 1 order: %d, Item 2 order: %d\n", 
						item1.GetOrder(), item2.GetOrder())
				} else {
					fmt.Println("Cannot swap - Item 2 has been deleted")
				}
				
			case <-moveBtn.ClickedCh:
				fmt.Println("Moving item 3 before item 1...")
				systray.MoveMenuItemBefore(item3, item1)
				fmt.Printf("After move - Item 3 order: %d, Item 1 order: %d\n", 
					item3.GetOrder(), item1.GetOrder())
				
			case <-deleteBtn.ClickedCh:
				if !item2Deleted {
					fmt.Println("Deleting item 2...")
					item2.Delete()
					item2Deleted = true
					deleteBtn.SetTitle("Item 2 Deleted")
					deleteBtn.Disable()
				}
				
			case <-reorderBtn.ClickedCh:
				fmt.Println("Reordering all items...")
				allItems := []*systray.MenuItem{item3, item1}
				if !item2Deleted {
					allItems = append(allItems, item2)
				}
				systray.ReorderMenuItems(allItems, 10) // Start from order 10
				fmt.Println("Items reordered starting from order 10")
				
			case <-clearBtn.ClickedCh:
				fmt.Println("Clearing all menu items...")
				// Note: This will also clear the control buttons, so we quit immediately
				systray.ClearAllMenuItems()
				time.Sleep(1 * time.Second)
				systray.Quit()
				return
				
			case <-quitBtn.ClickedCh:
				fmt.Println("Quit requested")
				systray.Quit()
				return
			}
		}
	}()

	// Demonstrate dynamic operations over time
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Auto-demo: Changing order of Item 1 to 0 (should move to top)")
		item1.SetOrder(0)
		
		time.Sleep(3 * time.Second)
		fmt.Println("Auto-demo: Moving Item 2 after Item 3")
		if !item2Deleted {
			systray.MoveMenuItemAfter(item2, item3)
		}
		
		time.Sleep(3 * time.Second)
		fmt.Println("Auto-demo complete. Try the manual controls!")
	}()
}

func onExit() {
	fmt.Println("Dynamic menu test finished")
}