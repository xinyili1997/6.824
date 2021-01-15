package mr

//
// RPC definitions.
//

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type RPCArgs struct {
	Key   string
	Value string
}

type RPCReply struct {
	//Err Err
  Value string
}
// Add your RPC definitions here.

