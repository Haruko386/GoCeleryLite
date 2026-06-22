# GoTaskFlow

GoTaskFlow is a lightweight asynchronous task queue system written in Go.

It is designed as a learning project for Go concurrency, including goroutines,
channels, worker pools, context cancellation, task retries, timeout control,
and task state management.

## Features

- In-memory task queue
- Worker pool based on goroutines
- Task status tracking
- HTTP API for task submission and query
- Built-in task handlers
- Context-based timeout and cancellation
- Future support for SQLite and Redis

## Roadmap

- [ ] In-memory queue
- [ ] Worker pool
- [ ] HTTP API
- [ ] Built-in sleep task
- [ ] Built-in echo task
- [ ] Built-in hash task
- [ ] Task timeout
- [ ] Task cancellation
- [ ] Retry mechanism
- [ ] SQLite persistence
- [ ] Redis broker
- [ ] Web dashboard