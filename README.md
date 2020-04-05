# gobermann

```Usage:
  gobermann [flags]

Flags:
  -a, --alg string     The domain generating algorithm to use. [locky, nymaim2] (default "locky")
  -b, --burst int      Number of requests in a burst of DNS traffic (default 15)
  -d, --delay int      Delay between requets in a burst in seconds (default 1)
  -s, --dns string     Target DNS Server (default "1.1.1.1")
  -r, --dryrun         Determines whether DNS traffic is sent over the wire (default true)
  -h, --help           help for gobermann
  -i, --interval int   Delay between bursts in minutes (default 720)
  -v, --verbose        Include verbose messages from program execution
```
