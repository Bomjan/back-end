# Pointers, Maps, and Real-World Applications

Learn about Go's pointer system and maps, with practical API examples.

## Learning Objectives

- Understand pointers and memory addresses
- Use the `*` (dereference) and `&` (address-of) operators
- Understand maps (Go's key-value store)
- See how slices and structs work in real-world API scenarios

## Key Concepts

### Pointers: References to Memory

A **pointer** is a variable that stores the **memory address** of another variable.

#### The `&` Operator: Get Address

```go
a := 42           // a is an integer variable
ptr := &a         // ptr is a pointer to a (stores the address)
```

#### The `*` Operator: Dereference

```go
a := 42
ptr := &a
fmt.Println(*ptr)  // Prints 42 (dereferences ptr to get the value)
```

#### Modifying Through Pointers

```go
a := 42
ptr := &a
*ptr = 100        // Changes a through the pointer
fmt.Println(a)    // 100
```

### Why Pointers Matter

**Without pointers** (copy by value):
```go
a := 42
b := a        // b gets a COPY of the value
b = 100
fmt.Println(a)  // 42 (unchanged)
```

**With pointers** (reference):
```go
a := 42
ptr := &a
modify(ptr)
fmt.Println(a)  // Can be changed by function

func modify(p *int) {
    *p = 100
}
```

### Maps: Key-Value Storage

Maps are Go's equivalent to JavaScript objects or Python dictionaries.

```go
// Create a map: map[KeyType]ValueType
mymap := make(map[string]string)

// Add key-value pairs
mymap["name"] = "Alice"
mymap["email"] = "alice@example.com"

// Retrieve values
fmt.Println(mymap["name"])  // Alice

// Check if key exists
if email, exists := mymap["email"]; exists {
    fmt.Println("Email:", email)
}

// Delete a key
delete(mymap, "name")

// Iterate over map
for key, value := range mymap {
    fmt.Println(key, "=>", value)
}
```

## JavaScript vs Go: Pointers

### JavaScript (No Explicit Pointers)
```javascript
let a = { name: "Alice" };
let b = a;                   // References same object
b.name = "Bob";
console.log(a.name);         // "Bob" (both reference same object)
```

### Go (Explicit Pointers)
```go
type User struct {
    name string
}

a := User{name: "Alice"}
b := a                       // Copies the struct
b.name = "Bob"
fmt.Println(a.name)          // "Alice" (unchanged)

// To reference same struct:
ptr := &a
ptr.name = "Bob"             // Now a.name is "Bob"
```

## Real-World Example: API User Response

Here's how pointers and structs work with API data:

```go
// Define a User struct (like a JSON model)
type User struct {
    ID       int
    Name     string
    Email    string
    Password string  // Should NOT be sent in API responses!
}

func main() {
    // Create a user
    user := User{1, "Sundra Bomjan", "bomjan@gmail.com", "secret"}
    
    // When sending API response, you might:
    // 1. Use a pointer to avoid copying large structs
    sendUserResponse(&user)
    
    // 2. Create a DTO (Data Transfer Object) without sensitive fields
    // 3. Use JSON tags to exclude fields
}

func sendUserResponse(user *User) {
    // Do something with user pointer
    fmt.Printf("User %s (%s) sent\n", user.Name, user.Email)
}
```

## Slices and Real-World Usage

### Removing an Item from a Slice

```go
mySlice := []int{3, 5, 7}    // [3, 5, 7]

// Remove element at index 1 (value 5)
newSlice := append(mySlice[:1], mySlice[2:]...)
// mySlice[:1] = [3]
// mySlice[2:] = [7]
// Result: [3, 7]
```

### Slice of Structs (Common in APIs)

```go
type User struct {
    ID    int
    Name  string
    Email string
}

// Create a slice of users
var users []User
users = append(users, User{1, "Alice", "alice@example.com"})
users = append(users, User{2, "Bob", "bob@example.com"})

fmt.Println(users)
// Output: [{1 Alice alice@example.com} {2 Bob bob@example.com}]
```

## Running This Example

This example demonstrates:
1. **Pointers**: Creating and dereferencing pointers with `&` and `*`
2. **Maps**: Creating and using key-value storage
3. **Commented code**: Real-world patterns (slices of structs, API responses)

```bash
go run main.go
```

## Important Notes

⚠️ **Uncommented code shows**:
- Pointer usage with `int8` type
- Map creation and usage

📝 **Commented code shows**:
- How to remove items from slices
- How to use structs for API responses
- How to create slices of structs

Uncomment the commented sections to see full examples!

## Maps vs JavaScript Objects

| Feature | JavaScript | Go |
|---------|-----------|-----|
| Creation | `{}` or `new Map()` | `make(map[K]V)` |
| Key type | Any (usually string) | Explicit type required |
| Value type | Any | Explicit type required |
| Syntax | `obj.key` or `obj["key"]` | `map["key"]` |
| Iteration | `for...in` or `for...of` | `for range` |

## Key Takeaways

✅ **Pointers** reference memory addresses  
✅ **`&` gets the address**, `*` dereferences it  
✅ **Maps are typed** key-value stores  
✅ **Slices can contain structs** (useful for API responses)  
✅ **Pointers avoid copying large data** (efficient for large structs)  

## Next Steps

1. Run the current program
2. Uncomment the slice removal example and see how it works
3. Uncomment the User struct example
4. Try creating your own map with different key/value types
5. Create a struct with multiple fields and a slice of them

## Try It

Exercises:
```go
// 1. Create a map of city populations
cities := make(map[string]int)
cities["New York"] = 8000000
cities["Tokyo"] = 13000000
// Print the map

// 2. Create a struct for a Product
type Product struct {
    ID    int
    Name  string
    Price float64
}

// 3. Make a slice of products and append items
// 4. Use a pointer to modify product price through function
```

---

**Related Modules**: 
- [01_basics - Variables](../../01_basics/03_variables/) (type system)
- [02_intermediate/HTTP](../HTTP/) (JSON with structs)
- [02_intermediate/mymath_package](../mymath_package/) (imports)
