abcd allows you to bind to specific interfaces which enables you to setup
configurations with varying levels of complexity.  The listen parameter can be
specified on the command line as shown below with the -- prefix or in the
configuration file without the -- prefix (as can all long command line options).
The configuration file takes one entry per line.

**NOTE:** The listen flag can be specified multiple times to listen on multiple
interfaces as a couple of the examples below illustrate.

Command Line Examples:

|Flags|Comment|
|----------|------------|
|--listen=|all interfaces on default port which is changed by `--testnet` and `--regtest` (**default**)|
|--listen=0.0.0.0|all IPv4 interfaces on default port which is changed by `--testnet` and `--regtest`|
|--listen=::|all IPv6 interfaces on default port which is changed by `--testnet` and `--regtest`|
|--listen=:9527|all interfaces on port 9527|
|--listen=0.0.0.0:9527|all IPv4 interfaces on port 9527|
|--listen=[::]:9527|all IPv6 interfaces on port 9527|
|--listen=127.0.0.1:9527|only IPv4 localhost on port 9527|
|--listen=[::1]:9527|only IPv6 localhost on port 9527|
|--listen=:9524|all interfaces on non-standard port 9524|
|--listen=0.0.0.0:9524|all IPv4 interfaces on non-standard port 9524|
|--listen=[::]:9524|all IPv6 interfaces on non-standard port 9524|
|--listen=127.0.0.1:9523 --listen=[::1]:9527|IPv4 localhost on port 9523 and IPv6 localhost on port 9527|
|--listen=:9527 --listen=:9523|all interfaces on ports 9527 and 9523|

The following config file would configure abcd to only listen on localhost for both IPv4 and IPv6:

```text
[Application Options]

listen=127.0.0.1:9527
listen=[::1]:9527
```
