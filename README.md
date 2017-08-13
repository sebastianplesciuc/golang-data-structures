# Data Structures written in Golang

[![Build Status](https://travis-ci.org/sebastianplesciuc/golang-data-structures.svg?branch=master)](https://travis-ci.org/sebastianplesciuc/golang-data-structures)

Personal notes on implementing various data structures in Golang.

1. Linked list
    - doubly linked
    - last element is known

## TODO

- Queue
- Stack
- Priority Queue
- Binary Tree
- Binary Search Tree
- Huffman Tree
- Hash table
- Heap

## Building the examples

```./bin/build.sh```

## Running the tests

```./bin/test.sh```

## Running from Docker

```docker-compose up --build --force-recreate -d```

```docker-compose exec datastructs bash```

While inside the container run `bin/test.sh`.