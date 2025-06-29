# 🚀 Go Concurrent Primitives

This repository contains practical, hands-on examples demonstrating how to use Go’s powerful concurrency primitives. From goroutines to mutexes, wait groups to semaphores — it's all here!

---

## 📦 Contents

| Folder / File        | Description                                  |
|----------------------|----------------------------------------------|
| `mutex/`             | Mutex and RWMutex examples to prevent races  |
| `waitgroup/`         | Coordinating goroutines using WaitGroups     |
| `semaphore/`         | Limiting concurrency with semaphores         |
| `channels/`          | Communicating between goroutines             |
| `once/`              | Ensuring one-time operations with sync.Once  |
| `atomic/`            | Using sync/atomic for low-level concurrency  |

> Each example is isolated and self-contained, making it easy to explore and test individually.

---

## 🧠 What You'll Learn

- The difference between **concurrency** and **parallelism**
- Safe access to shared memory using **mutexes**
- Efficient coordination using **WaitGroups**
- How **RWMutex** optimizes concurrent reads
- Managing goroutine execution with **semaphores**
- Atomic operations and thread-safe counters
- And much more!

---

## 🛠 How to Run

Clone the repo and run any file using:

```bash
go run <file>.go
