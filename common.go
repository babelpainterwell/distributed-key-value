package kvsrv

// Put or Append arguments
type PutAppendArgs struct {
	Key 		string
	Value 		string
	Op    		string // "Put" or "Append"
	ClientID    int64 
	SeqNum      int64  // Sequence number of the request (duplcate detection)
}

type PutAppendReply struct {
	Value       string // Append operation returns the old value
}

// Get arguments 
type GetArgs struct {
	Key         string 
}

type GetReply struct {
	Value 		string
}