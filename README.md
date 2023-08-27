# taylz.io/log

a configurable logger for golang

## `log.Writer`

interface for sending lines to the log

supports sourcing, timestamps, fields, and levels

## `log.Liner`

interface for processing a `log.Line`

supports logging middleware

## `log.Formatter`

interface for converting a `log.Line` to `[]byte`

supports `log.IOLiner` ( with `io.WriteCloser` )
