# Dynamic Menu Operations Example

This example demonstrates the enhanced systray library with dynamic menu ordering and deletion capabilities.

## New Features

### 1. Dynamic Ordering
```go
// Set specific order
item.SetOrder(5)

// Get current order
order := item.GetOrder()

// Add item at specific position
item := systray.AddMenuItemAtIndex("Title", "Tooltip", 2)
```

### 2. Menu Item Deletion
```go
// Delete a menu item permanently
item.Delete()
```

### 3. Bulk Operations
```go
// Swap two items
systray.SwapMenuItems(item1, item2)

// Move item before another
systray.MoveMenuItemBefore(itemToMove, targetItem)

// Move item after another  
systray.MoveMenuItemAfter(itemToMove, targetItem)

// Reorder multiple items
items := []*systray.MenuItem{item1, item2, item3}
systray.ReorderMenuItems(items, 10) // Start from order 10

// Get all menu items
allItems := systray.GetAllMenuItems()

// Clear all menu items
systray.ClearAllMenuItems()
```

## Usage Example

```go
func onReady() {
    // Create items with specific order
    item1 := systray.AddMenuItemAtIndex("First", "First item", 1)
    item2 := systray.AddMenuItemAtIndex("Second", "Second item", 2) 
    item3 := systray.AddMenuItemAtIndex("Third", "Third item", 3)
    
    // Dynamic operations
    go func() {
        time.Sleep(2 * time.Second)
        
        // Change order dynamically
        item3.SetOrder(0) // Move to top
        
        // Swap items
        systray.SwapMenuItems(item1, item2)
        
        // Delete an item
        item2.Delete()
    }()
}
```

## Platform Support

- ✅ **macOS**: Uses `insertItem:atIndex:` and `removeItem:`
- ✅ **Windows**: Uses `InsertMenuItemW` position parameter and `RemoveMenu`
- ✅ **Linux**: Uses `gtk_menu_shell_insert` and `gtk_container_remove`

## Running the Example

```bash
cd example
go build .
./example
```

Try the dynamic_test.go for a comprehensive demonstration of all features.