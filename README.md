# d-clock

![Go](https://github.com/jabolina/d-clock/workflows/Go/badge.svg?branch=main)

A simple distributed logical clock using the [Raft Protocol](https://raft.github.io/) for
consistency.

## Available

There is available two different options for the logical clock.

### In-Memory

A simpler and cheaper implementation, that is thread safe and can be used
by a system inside a single host.

### Distributed

A distributed version, more expensive and strongly consistent. Should be used only when the
application is spread across multiple hosts.

#### Configuration

When the clock is distributed, it needs to communicate with the other hosts, so the need
for a configuration is needed.
