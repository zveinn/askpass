# installation
download the askpass binary in the root of this repository or build it from source:
```bash 
$ go build .
```

# fprintd-enroll
Use fprintd-entroll to enroll whichever finger you want to use for authentication.
Askpass does not care which finger it is.

# /etc/sudoers
Remove original user and %sudo lines if you want to lock sudo to fingerprint.

```
USERNAME  ALL=(ALL:ALL) NOPASSWD: /usr/local/bin/askpass
%sudo   ALL=(ALL:ALL) NOPASSWD: /usr/local/bin/askpass
```



