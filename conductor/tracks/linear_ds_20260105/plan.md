# Track Plan: Foundational Linear Data Structures

## Phase 1: Project Setup & Linked Lists [checkpoint: 919b632]
建立项目基础结构，并实现核心的链表结构。

- [x] Task: Initialize Project Structure 02777b2
    - [ ] Subtask: Initialize go module `go mod init <module_name>` (User to provide or infer `go-algorithm`)
    - [ ] Subtask: Create standard layout directories: `pkg/structures`, `pkg/algorithms`, `cmd`.
    - [ ] Subtask: Configure `golangci-lint`.
- [x] Task: Implement Singly Linked List (Generic) 897a7ac
    - [ ] Subtask: Define `Node[T]` and `LinkedList[T]` struct in `pkg/structures/list`.
    - [ ] Subtask: Write Failing Tests for `PushFront`, `PushBack`, `Size`.
    - [ ] Subtask: Implement `PushFront`, `PushBack`, `Size`.
    - [ ] Subtask: Write Failing Tests for `Get`, `InsertAt`, `Remove`.
    - [ ] Subtask: Implement `Get`, `InsertAt`, `Remove`.
    - [ ] Subtask: Refactor and Verify Coverage.
- [x] Task: Implement Doubly Linked List (Generic) c20d629
    - [ ] Subtask: Define `DoublyNode[T]` and `DoublyLinkedList[T]`.
    - [ ] Subtask: Write Failing Tests for bidirectional operations.
    - [ ] Subtask: Implement operations (Add `Prev` pointer logic).
    - [ ] Subtask: Verify Coverage (>80%).
- [x] Task: Conductor - User Manual Verification 'Phase 1: Project Setup & Linked Lists' (Protocol in workflow.md) 919b632

## Phase 2: Stacks and Queues
基于切片和链表实现栈与队列，并处理环形队列逻辑。

- [x] Task: Implement Stack (Slice & List based) f4c62ea
    - [ ] Subtask: Create `pkg/structures/stack`.
    - [ ] Subtask: Write Failing Tests for `SliceStack` (Push/Pop/Peek).
    - [ ] Subtask: Implement `SliceStack`.
    - [ ] Subtask: Write Failing Tests for `ListStack` (using `pkg/structures/list`).
    - [ ] Subtask: Implement `ListStack`.
- [x] Task: Implement Queue (Slice & List based) 7171296
    - [ ] Subtask: Create `pkg/structures/queue`.
    - [ ] Subtask: Write Failing Tests for `SliceQueue` and `ListQueue`.
    - [ ] Subtask: Implement `SliceQueue` (handle underlying array shifting or logical pointers).
    - [ ] Subtask: Implement `ListQueue`.
- [ ] Task: Implement Circular Queue
    - [ ] Subtask: Define `CircularQueue[T]` with fixed size.
    - [ ] Subtask: Write Failing Tests for `IsFull`, `Enqueue` (wrap around), `Dequeue` (wrap around).
    - [ ] Subtask: Implement Circular Queue logic.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Stacks and Queues' (Protocol in workflow.md)

## Phase 3: Documentation, Visualization & Performance
完善文档，添加图解，并进行详细的性能评估。

- [ ] Task: Add Mermaid Visualizations
    - [ ] Subtask: Create `pkg/structures/list/README.md` with Singly/Doubly list diagrams.
    - [ ] Subtask: Create `pkg/structures/stack/README.md` with Stack LIFO diagrams.
    - [ ] Subtask: Create `pkg/structures/queue/README.md` with Queue/Circular diagrams.
- [ ] Task: Implement Benchmarks
    - [ ] Subtask: Create `pkg/structures/stack/benchmark_test.go` (Slice vs List).
    - [ ] Subtask: Create `pkg/structures/queue/benchmark_test.go` (Slice vs List vs Circular).
    - [ ] Subtask: Run benchmarks and record results.
- [ ] Task: Finalize Documentation
    - [ ] Subtask: Update package READMEs with Benchmark results and Big O analysis.
    - [ ] Subtask: Create root `README.md` linking to all modules.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Documentation, Visualization & Performance' (Protocol in workflow.md)
