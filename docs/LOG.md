# Logging

Notes on logging within the `cfm-service` and `cfm-cli` application.

## General

The logging mdoule used by this project is klog.

All logging calls are to use the strucutred logging log commands (ending in "S")
    ex: `InfoS()`

## Logging Errors

Errors are to be logged without using the verbosity level so as to ensure that they are always logged.
    ex: `klog.ErrorS("my error")`

## Logging Verbosity Levels

All other logging will use logging levels 1-4, as defined below:

- 0 - logging OFF (Do **not** use)
- 1 - Warning
- 2 - Info
- 3 - Debug
- 4 - All

Usage example:
    `klog.V(3).InfoS("log message", "value", value)`
