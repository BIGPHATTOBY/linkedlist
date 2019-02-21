# LinkedList
This repo implements a linked-list and a few functions to manipulate it with.

### Functions:
* Length // Returns length of linked list
* PrintAll // Returns slice containing node values
* GetValueByIndex // Returns value from node by index
* Append // Append node to the list
* Insert // Insert node at index
* Swap // Swap two nodes by index
* Delete // Delete node by index

## Example:
```GO
package main

import(
  "fmt"
  ll "github.com/bigphattoby/linkedlist"
)
func main() {
  ll := ll.NewLinkedList()
  ll.Append("foo")
  ll.Append("bar")
  ll.Append("cowsay")
  ll.Delete(2)
  s, err := ll.PrintAll()
  if err != nil {
    fmt.Println(err)
  }
  for _, list := range s {
    fmt.Println(list)
  }
}
```
