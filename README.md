## gobermann

Basic usage of Gobermann

```
Usage:
  gobermann [flags]

Flags:
  -a, --alg string        The domain generating algorithm to use.[locky, nymaim2, necurs, monero] (default "locky")
  -b, --burst int         Number of requests in a burst of DNS traffic (default 15)
  -d, --delay int         Delay between requets in a burst in milliseconds (default 500)
  -s, --dns string        Target DNS Server (default "1.1.1.1")
  -r, --dryrun            When enabled dns traffic will not be sent over the wire (default true)
  -h, --help              help for gobermann
  -i, --interval int      Delay between bursts in minutes (default 720)
  -l, --loglevel string   Include verbose messages from program execution [error, warn, info, debug] (default "info")
  -v, --version           version for gobermann
```

## Purpose
Gobermann is a utility for emulating traffic patterns seen by the likes of malware using domain generation algorithms to reach out to and identify C2 servers. The Go language was chosen for is flexibility of compliation and deployment.

This project is inspired by [Maltese](https://github.com/HPE-AppliedSecurityResearch/maltese) by Hewlett Packard Enterprise which has not been maintained for many years. 
