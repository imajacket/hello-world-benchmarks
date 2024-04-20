# Simple Hello World Benchmarks
Testing both the executable size and run times for simple hello world programs in multiple languages.

Tests are all compiled to native executables - interpreted & VM languages are compiled using additional tools.

All build commands are at their simplest form, no attempts at optimization are made.

# Results
### File sizes (on macOS - M2 Pro):
    Go:     1.4MB
    C:      33KB
    Rust:   440KB
    Python: 6.8MB
    Java:   6.2MB



### Run times:
```
./benchmarkTool -runs=500
    
    Go:     1.637288ms
    C:      1.646728ms
    Rust:   1.650977ms
    Python: 1.656732ms
    Java:   1.670468ms

```

## Languages and Build Commands Used:
* Java:
    ```
    // GraalVM
    javac HelloWorld.java
    native-image HelloWorld
    ```
* Python:
    ```
    // Pyinstaller
    pyinstaller -F -n test_python test_python.py
    ```
* Go:
    ```
    go build test_go.go
    ```
* Rust:
    ```
    rustc test_rust.rs
    ```
* C:
    ```
    gcc -o test_c test_c.c
    ```
