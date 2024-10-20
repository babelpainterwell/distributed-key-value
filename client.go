package kvsrv

import (
	"crypto/rand"
	"math/big"
	"time"

	"keyvalueserverwell/labrpc"
)


type Clerk struct {
	server *labrpc.ClientEnd // in order t simulate dropped messages or network failure
	clientID int64 // unique identifier for the client 
	seqNum   int64 // the latest sequence number to use for requests 
}

func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	x := bigx.Int64()
	return x
}

func MakeClerk(server *labrpc.ClientEnd) *Clerk {
	ck := new(Clerk)
	ck.server = server
	ck.clientID = nrand()
	ck.seqNum = 0
	return ck
}


func (ck *Clerk) Get(key string) string {

	args := GetArgs{Key: key}
	var reply GetReply

	for {
		ok := ck.server.Call("KVServer.Get", &args, &reply)

		// Time-out? - prevent from getting stuck if the server is not responsive
		// Not necessary, Call will return ok == false if the server doesn't respond within a certain time frame. 

		if ok {
			return reply.Value
		}

		// retry after a short delay of RPC fails 
		time.Sleep(100 * time.Millisecond)
	}
	
}


func (ck *Clerk) PutAppend(key string, value string, op string) string {
	ck.seqNum ++
	args := PutAppendArgs{
		Key: key, 
		Value: value,
		Op: op,
		ClientID: ck.clientID,
		SeqNum: ck.seqNum,
	}
	var reply PutAppendReply

	for {
		ok := ck.server.Call("KVServer."+op, &args, &reply)
		if ok {
			return reply.Value
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func (ck *Clerk) Put(key string, value string) {
	ck.PutAppend(key, value, "Put")
}

// Append value to key's value and return that value
func (ck *Clerk) Append(key string, value string) string {
	return ck.PutAppend(key, value, "Append")
}