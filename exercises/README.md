# Go Learning Exercises

This repository contains comprehensive Go programming exercises covering all major concepts from go-by-example and more. Each exercise is designed to help you learn Go through hands-on practice.

## Structure

The exercises are organized into numbered directories under `exercises/`, each focusing on specific Go concepts:

1. **Hello World** - Basic Go program structure
2. **Values & Variables** - Working with different data types
3. **For Loops** - Different types of loops in Go
4. **If/Else** - Conditional statements
5. **Switch** - Switch statements and expressions
6. **Arrays & Slices** - Working with arrays and slices
7. **Maps** - Hash maps/dictionaries
8. **Range** - Iterating over data structures
9. **Functions** - Function definitions and usage
10. **Closures** - Functions that capture variables
11. **Recursion** - Recursive function calls
12. **Pointers** - Memory addresses and pointers
13. **Structs** - Custom data types
14. **Methods** - Functions associated with types
15. **Interfaces** - Interface definitions and implementations
16. **Errors** - Error handling patterns
17. **Goroutines** - Concurrent programming basics
18. **Channels** - Communication between goroutines
19. **Worker Pools** - Advanced concurrency patterns
20. **Mutexes** - Thread-safe operations
21. **Context** - Context for cancellation and timeouts
22. **JSON** - JSON encoding and decoding
23. **Strings** - String manipulation and formatting
24. **Regular Expressions** - Pattern matching
25. **File I/O** - File operations

## Getting Started

### Prerequisites

- Go 1.19 or later installed
- A text editor or IDE
- Basic understanding of programming concepts

### Running the Exercises

1. **Run all tests** to see what needs to be implemented:
   ```bash
   make test
   ```

2. **Run a specific exercise**:
   ```bash
   make test-single EX=1
   ```

3. **Check your code quality**:
   ```bash
   make lint
   ```

### How to Complete Exercises

1. Open an exercise file (e.g., `exercises/1/01_hello_test.go`)
2. Find the `TODO` comments indicating what you need to implement
3. Replace the placeholder code with your implementation
4. Run the tests to verify your solution:
   ```bash
   make test-single EX=1
   ```
5. Repeat until all tests pass

### Example

In `exercises/1/01_hello_test.go`, you'll find:

```go
func HelloWorld() string {
    // TODO: Return "Hello, World!"
    return ""
}
```

Replace the empty return with:

```go
func HelloWorld() string {
    // TODO: Return "Hello, World!"
    return "Hello, World!"
}
```

## Learning Path

### Beginner (Exercises 1-8)
Start with basic Go syntax and data structures:
- Hello World, Variables, Control Flow
- Arrays, Slices, Maps, Range

### Intermediate (Exercises 9-16)
Learn about functions and object-oriented concepts:
- Functions, Closures, Recursion
- Pointers, Structs, Methods, Interfaces
- Error Handling

### Advanced (Exercises 17-25)
Master concurrency and advanced topics:
- Goroutines, Channels, Worker Pools
- Synchronization, Context
- JSON, String Processing, File I/O

## Tips for Success

1. **Read the TODO comments carefully** - they contain specific instructions
2. **Look at the test cases** - they show expected behavior
3. **Start simple** - implement the basic functionality first
4. **Run tests frequently** - to get immediate feedback
5. **Don't skip exercises** - each builds on previous concepts
6. **Read Go documentation** - [golang.org/doc](https://golang.org/doc/)

## Test Output

When you run tests, you'll see output like:

```
=== RUN   TestHelloWorld
--- FAIL: TestHelloWorld (0.00s)
    01_hello_test.go:13: HelloWorld() = "", want "Hello, World!"
FAIL
```

This tells you:
- Which test failed (`TestHelloWorld`)
- What your function returned (`""`)
- What was expected (`"Hello, World!"`)
- Which line of code to look at (`01_hello_test.go:13`)

## Common Go Patterns You'll Learn

- **Error handling**: `if err != nil { return err }`
- **Multiple return values**: `func divide(a, b int) (int, error)`
- **Pointer receivers**: `func (p *Person) SetAge(age int)`
- **Interface satisfaction**: Implementing methods to satisfy interfaces
- **Channel patterns**: Producer-consumer, fan-out/fan-in
- **Context usage**: Timeouts, cancellation, value passing

## Additional Resources

- [A Tour of Go](https://tour.golang.org/) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Annotated example programs
- [Effective Go](https://golang.org/doc/effective_go.html) - Writing idiomatic Go
- [Go Documentation](https://golang.org/doc/) - Official documentation

## Getting Help

If you're stuck:

1. Read the error messages carefully
2. Check the test cases for expected behavior
3. Look up Go documentation for relevant packages
4. Break the problem into smaller parts
5. Use `fmt.Printf` for debugging

## Progress Tracking

Keep track of your progress:

- [ ] Exercises 1-8: Basic Go (Beginner)
- [ ] Exercises 9-16: Functions & OOP (Intermediate)  
- [ ] Exercises 17-25: Concurrency & Advanced (Advanced)

Good luck with your Go learning journey! 🚀

## Make Commands

- `make test` - Run all exercises
- `make test-single EX=N` - Run specific exercise (replace N with number)
- `make lint` - Check code quality
- `make clean` - Clean test cache
- `make help` - Show available commands
