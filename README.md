# Understanding Program Behavior with Tracing

This repository contains example code demonstrating how to use tracing to understand and improve program performance. We'll explore how tracing can reveal both what is and isn't happening in your program during execution.

## Overview

We use a simple RSS feed processor as our example application. The program searches through multiple XML files for specific terms, demonstrating common performance challenges in real-world applications.

## Getting Started

### Prerequisites
- Go 1.11 or later
- Chrome browser (required for viewing traces)

### Running the Example

1. Clone this repository
2. Build the program:
```bash
go build
```

3. Run with tracing enabled:
```bash
tracingingo./ > trace.out
```

4. View the trace:
```bash
go tool trace trace.out
```

## Understanding the Code

The program implements two search strategies:
- `searchSequential`: Processes files one at a time
- `searchConcurrent`: An alternative implementation we'll explore using trace analysis

To switch between implementations, uncomment/comment the relevant lines in `main()`.

## Trace Analysis

The trace viewer will show several important metrics:
- Goroutine execution
- Heap memory usage
- Garbage collection activity
- CPU utilization

Key points to observe in the traces:
- Goroutine states (running, blocked, waiting)
- Memory allocation patterns
- CPU utilization across available cores

## Additional Resources

To learn more about Go's tracing capabilities:
- [Go Blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [Go Trace Documentation](https://pkg.go.dev/runtime/trace)

## License
This project is licensed under [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)