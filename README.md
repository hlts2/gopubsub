# gpubsub

gpubsub is a [topic-based](http://en.wikipedia.org/wiki/Publish–subscribe_pattern#Message_filtering) [publish/subscribe](http://en.wikipedia.org/wiki/Publish/subscribe) library written in golang.

## Requirement

Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/gopubsub
```

## Example

To subscribe:

```go
ps := gopubsub.NewPubSub()

subscriber := ps.Subscribe("t1")
```

To add subscribe:

```go
subscriber := ps.Subscribe("t1")

// Adds subscriptions
ps.AddSubsrcibe("t2", subscriber)
```

To publish:

```go
// publish message to topic asyncronously
ps.Publish("t1", "Hello World!!")

// Because the message type is `interface{}`, you can publish anything
ps.Publish("t1", func() {
  fmt.Println("Hello World!!")
})
```

To Cancel specific subscription:

```go
ps.UnSubscribe("t1")
```

To fetch published message

```go
message := <-subscriber.Read()
```

## Author
[hlts2](https://github.com/hlts2)
