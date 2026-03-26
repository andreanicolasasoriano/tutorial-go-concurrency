# Race Conditions
- Access/modify the thing then they come out funny;

# Mutexes
sync.Mutex()
- deals with race conditions;
- tries to hit the same piece of data so that you lock/unlock data as necessary;
- You don't always really see race conditions unless you run go run -race .

# Channels
"share memory by communicating, rather than communicating by sharing memory"
- Preferred way of dealing with concurrency in go;