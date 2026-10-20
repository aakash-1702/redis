1. GoRoutines : One of the many workers who are doing their tasks.(Independent of other workers);
2. For starting a new GoRoutine , we use keyword `go`;
3. Channels are essentially the connection medium between GoRoutines.
4. These channels helps in transferring data from one routine to another.
5. If the channel has no buffer , then it does one thing -> it blocks sender until reciever is ready , it blocks reciever until sender is ready.

## JOBS is the very important part of goroutines and channels

It is like creating workers , and each worker independently can consume or basically execute that job/event.
It does not depends upon each other.
If a job comes , it could be picked by any worker free then , no gurrantee of equal distribution out of box here.
