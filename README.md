### Introduction

The project involves implementing a **key/value server** that provides reliable and consistent storage of key/value pairs in a distributed system environment. The server supports three primary operations:

- Get(key): Fetches the current value associated with a given key
- Put(key, value): Stores or updates the value associated with a given key.
- Append(key, value): Appends a value to the existing value of a given key and returns the old value.

The following principles are ensured:

- **Exactly-Once Semantics**: Each client operation (Get, Put, Append) is executed exactly once, even in the presence of network failures like dropped messages.
- **Linearizability**: Operations appear to occur atomically and in some total order that is consistent with the real-time ordering of operations, making the system's behavior predicatble and reliable for clients.
  [What is linearizability](https://pdos.csail.mit.edu/6.824/papers/linearizability-faq.txt)

### Key Concepts

- **Client Clerk**: encapsulates the client's interaction with the server. It sends RPCs (Remote Procedure Calls) to the server and handles retries in case of failures.
- **Server**: processes incoming client requests, maintains the key/value store, and ensures consistency and reliability.
- **Duplicate Detection**: detects and ignores duplicate requests (eg. due to client retries) to prevent executing the same operation multiple times by using `ClientID` and `SeqNum`.
- **Network Failures Handling**: handles scenarios where messages (requests or replies) are lost due to network issues.

### Challenges

- The system should work correctly even the network is unreliable and messages can be lost.
- Concurrent operations from multiple clients can be handled.
- Duplicate requests should be detected by maintaining the state on the server.

### Testing

Since this project originates from Lab 2 of MIT 6.5840 Distributed Systems, the lab itself includes test cases and files provided by the MIT faculty and staff. However, to deepen my understanding of the concepts involved in this project, I'm writing my own test files.
