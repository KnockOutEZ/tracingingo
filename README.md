# Tracing in Go - From ৭-৮-৯ টেক আড্ডা Talk

This repository contains the example code and materials from the tracing presentation given at ৭-৮-৯ টেক আড্ডা. The talk demonstrates how to use tracing to understand and improve program performance, using a practical example of processing RSS news feeds.

## Talk Overview

We explore how tracing can help developers:
- Visualize program behavior
- Identify performance bottlenecks
- Understand resource utilization
- Make data-driven optimization decisions

## Example Application

The repository includes `main.go`, a program that:
- Searches through RSS news feeds for specific terms
- Demonstrates common performance patterns
- Shows how tracing can guide optimization

### Running the Example

1. Clone this repository
2. Build the program:
```bash
go build main.go
```

3. Run with tracing enabled:
```bash
tracingingo > trace.out
```

4. View the trace (requires Chrome browser):
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

## Following Along

The code is structured to follow the presentation flow:
1. Start with a basic sequential implementation
2. Use tracing to identify bottlenecks
3. Explore optimizations through concurrent implementation
4. Compare performance improvements

## Additional Resources

To learn more about Go's tracing capabilities:
- [Go Blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [Go Trace Documentation](https://pkg.go.dev/runtime/trace)

## License
This project is licensed under [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)