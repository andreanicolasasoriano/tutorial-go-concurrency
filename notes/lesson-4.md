# Channels
- A way for one goroutine to communicate with another;

## Syntax

## Send/receive only annotations

## Select statements
- If there are more cases that can satisfiy the case, it will just choose one at random.
```go

for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case 1", s1)
		case s2 := <-channel1:
			fmt.Println("Case 2", s2)
		case s3 := <-channel2:
			fmt.Println("Case 3", s3)
		case s4 := <-channel2:
			fmt.Println("Case 3", s4)
		default:
			
		}

	}
```

## Buffered channels
- Limited number of channels.
- Waits for a channel to go free before you let another one in.
- Effective at rate limiting. very nice.