
# ecs-blox

[![godoc][godoc-badge]][godoc]

This library is an ECS-compatible API client that delegates to [blox][blox]
(cluster-state-service) when possible, but uses ECS directly in other cases or
if a failure is encountered.

The primary goal is to allow for applications to use ECS for running and
monitoring tasks, but avoiding the pretty low rate-limits that the API has.


## Usage

Where you are using the `*ecs.ECS` type, use the `ecsiface.ECSAPI` interface
type instead. Then, you can initialize this client and drop it in place. (or
quickly switch back to a "vanilla" ECS client)


[blox]: https://github.com/blox/blox
[godoc-badge]: https://godoc.segment.com/github.com/segmentio/ecs-blox?status.svg
[godoc]: https://godoc.segment.com/github.com/segmentio/ecs-blox
