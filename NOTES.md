1. GoRoutines : One of the many workers who are doing their tasks.(Independent of other workers);
2. For starting a new GoRoutine , we use keyword `go`;
3. Channels are essentially the connection medium between GoRoutines.
4. These channels helps in transferring data from one routine to another.
5. If the channel has no buffer , then it does one thing -> it blocks sender until reciever is ready , it blocks reciever until sender is ready.
6.
