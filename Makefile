build: build_java build_python build_go build_rust build_c

build_java:
	@javac HelloWorld.java
	@native-image HelloWorld
	@mv helloworld test_java
	@rm HelloWorld.class

build_python:
	@pyinstaller -F -n test_python test_python.py
	@cp dist/test_python test_python
	@rm test_python.spec
	@rm -rf build
	@rm -rf dist

build_go:
	@go build test_go.go

build_rust:
	@rustc test_rust.rs

build_c:
	@gcc -o test_c test_c.c

clean:
	@rm test_java
	@rm test_python
	@rm test_go
	@rm test_rust
	@rm test_c