package main

import "fmt"

func partition(list []int, start int, end int) int {
  bottom := start - 1
  top := end
  pivot := list[end]

  done := false

  for done == false {
    for done == false {
      bottom += 1
      if bottom == top {
        done = true
        break
      }
      if list[bottom] > pivot {
        list[top] = list[bottom]
        break
      }
    }

    for done == false {
      top -= 1
      if top == bottom {
        done = true
        break
      }
      if list[top] < pivot {
        list[bottom] = list[top]
        break
      }
    }
  }
  list[top] = pivot
  return top
}

func quicksort(list []int, start int, end int) {
  if start < end {
    split := partition(list, start, end)
    quicksort(list, start, split - 1)
    quicksort(list, split + 1, end)
  } else {
    return
  }
}

func main() {
  unsorted := []int{5,1,4,2,3}
  fmt.Println("Unsorted =", unsorted)

  quicksort(unsorted, 0, len(unsorted) - 1)
  fmt.Println("Sorted =", unsorted)
}
