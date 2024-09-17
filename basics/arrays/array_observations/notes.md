### Arrays as Value Types in Assignment

When you declare an array and assign it to another array, the entire array is copied element by element. This means that changes made to one array do not affect the other. However, direct assignment of arrays is not allowed in C, so you typically use loops or functions like `memcpy` to copy arrays.

### Example of Array Copying

```c
#include <stdio.h>
#include <string.h>

int main() {
    int originalArray[] = {1, 2, 3, 4, 5};
    int newArray[5];
    
    // Copying array using a loop
    for(int i = 0; i < 5; i++) {
        newArray[i] = originalArray[i];
    }
    
    // Modify the new array
    newArray[0] = 10;
    
    printf("Original array: ");
    for(int i = 0; i < 5; i++) {
        printf("%d ", originalArray[i]);
    }
    printf("\n");
    
    printf("New array: ");
    for(int i = 0; i < 5; i++) {
        printf("%d ", newArray[i]);
    }
    printf("\n");
    
    return 0;
}
```

Output:
```
Original array: 1 2 3 4 5 
New array: 10 2 3 4 5 
```

### Arrays in Function Parameters

When you pass an array to a function, what actually gets passed is a pointer to the first element of the array. This means that the function can modify the original array. In this context, arrays behave more like reference types.

### Example of Array Passed to Function

```c
#include <stdio.h>

void modifyArray(int arr[], int size) {
    arr[0] = 10; // Modify the first element
}

int main() {
    int originalArray[] = {1, 2, 3, 4, 5};
    
    printf("Original array before modification: ");
    for(int i = 0; i < 5; i++) {
        printf("%d ", originalArray[i]);
    }
    printf("\n");
    
   // after passing originalArray, the funciton points to the starting element of the array
    modifyArray(originalArray, 5);
    
    printf("Original array after modification: ");
    for(int i = 0; i < 5; i++) {
        printf("%d ", originalArray[i]);
    }
    printf("\n");
    
    return 0;
}
```

Output:
```
Original array before modification: 1 2 3 4 5 
Original array after modification: 10 2 3 4 5 
```

In summary, arrays in C are treated as value types when it comes to copying and assignment, but they behave like reference types when passed to functions. If you have any more questions or need further clarification, feel free to ask!
