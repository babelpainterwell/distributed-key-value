package kvsrv

import (
	"log"
	"sync"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}


type KVServer struct {
	mu 			sync.Mutex
	data 		map[string]string 
	lastSeq     map[int64]int64         // for Put operation, ClientID -> sequence number 
	lastReply   map[int64]string        // for Append Operation, ClientID -> the last cached reply for this client

}

// Get is not affected by diplcaite messages
func (kv *KVServer) Get(args *GetArgs, reply *GetReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	// retrieve the data 
	// return empty string if the key doesn't exist 
	value, ok := kv.data[args.Key]
	if ok {
		reply.Value = value
	} else {
		reply.Value = ""
	}
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	// check if the request is duplicated 
	lastSeqNum, seenClient := kv.lastSeq[args.ClientID]
	if seenClient && args.SeqNum <= lastSeqNum {
		return 
	}

	// if not duplicates, perform the put operation 
	kv.data[args.Key] = args.Value
	kv.lastSeq[args.ClientID] = args.SeqNum

}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	// check if the request is duplicated 
	lastSeqNum, seenClient := kv.lastSeq[args.ClientID]
	if seenClient && args.SeqNum <= lastSeqNum {
		reply.Value = kv.lastReply[args.ClientID]
		return
	}

	// Otherwise, perform the Append operation, and return the old value 
	old_value := kv.data[args.Key]
	new_value := old_value + args.Value
	kv.data[args.Key] = new_value

	reply.Value = old_value

	// update the last reply and sequence 
	kv.lastReply[args.ClientID] = reply.Value
	kv.lastSeq[args.ClientID] = args.SeqNum

}

func StartKVServer() *KVServer {
	kv := new(KVServer)

	kv.data = make(map[string]string)
	kv.lastSeq = make(map[int64]int64)
	kv.lastReply = make(map[int64]string)
	
	return kv
}