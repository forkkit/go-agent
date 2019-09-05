// Package nrnats instruments https://github.com/nats-io/nats.go.
//
// This package can be used to simplify instrumenting NATS publishers and subscribers.
//
// NATS publishers
//
// To to
//// generate an external segment for any method that publishes or responds to a NATS message, use the
// `StartPublishSegment` method. The resulting segment will also need to be ended. Example:
//
// nc, _ := nats.Connect(nats.DefaultURL)
// txn := currentTransaction()  // current newrelic.Transaction
// subject := "testing.subject"
// seg := nrnats.StartPublishSegment(txn, nc, subject)
// err := nc.Publish(subject, []byte("Hello World"))
// if nil != err {
//   panic(err)
// }
// seg.End()
//
// Or:
//
// nc, _ := nats.Connect(nats.DefaultURL)
// txn := currentTransaction()  // current newrelic.Transaction
// subject := "testing.subject"
// defer nrnats.StartPublishSegment(txn, nc, subject).End()
// _ := nc.Publish(subject, []byte("Hello World"))
//
//
// NATS subscribers
//
// The `nrnats.SubWrapper` function can be used to wrap the function for `nats.Subscribe` and `nats.QueueSubscribe`.
// If the `newrelic.Application` parameter is non-nil, it will create a `newrelic.Transaction` and end the transaction
// when the passed function is complete.  Example:
//
// nc, _ := nats.Connect(nats.DefaultURL)
// app := createNRApp()  // newrelic.Application
// subject := "testing.subject"
// nc.Subscribe(subject, nrnats.SubWrapper(app, myMessageHandler))
//
// Full Publisher/Subscriber example:
// https://github.com/newrelic/go-agent/go-agent/blob/master/_integrations/nrnats/examples/main.go
package nrnats

import "github.com/newrelic/go-agent/internal"

func init() { internal.TrackUsage("integration", "framework", "nats") }
