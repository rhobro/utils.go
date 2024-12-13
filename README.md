# utils.go

General utilities for use in Go programs.

Includes the following:
- `coll` - operations on slices of standard datatypes and strings.
- `fileio` - unzipping ZIP files, generating tempfile paths and directory operations.
- `httputil` - persistent HTTP requests, generating random User-Agents
- `timetrack` - used for timing parts of programs.
- `services/cfgcat` - basic wrapper for connecting to ConfigCat for feature flags.
- `services/sentree` - basic wrapper for integrating with Sentry error tracking.
- `services/storaje` - wrapper for connecting to the Storj decentralised platform.
