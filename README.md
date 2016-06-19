remgo
=====

Remote command execution for lazy people, using a small config file (toml) where you can define the remote hosts in roles and set tasks, indicating the server role, user if it's different from the running one and of course the command to run in the task.

ConfigFile (remgo.toml):

```toml
title = "string"
logdir = "path"

[servers]
[servers.ServerGroupName]
IPs = ["ipaddr", "hostname", "fqdn"]

[tasks]
[tasks.SomeName]
user = "username"
role = "ServerGroupName"
command = "SomeShellCommand"
log = true
```

logdir, user and log are optionals.

if logs is set to true in a task but logdir is not set, then the logs will be generated in the same dir where you are running the app.

Remember to set the ssh keys in your servers.

```bash
$ ./remgo
#####################################
#####################################
##.----.-----.--------.-----.-----.##
##|   _|  -__|        |  _  |  _  |##
##|__| |_____|__|__|__|___  |_____|##
######################|_____|########
#####################################

Running Example of remgo Configuration
###############################################################
Task: CheckDir
###############################################################
Servers Role: External
===============================================================
===============================================================
IP: 192.168.125.200
===============================================================
Can't Dial
--- FAILED ---
dial tcp 192.168.125.200:22: getsockopt: no route to host
===============================================================
IP: 192.168.125.60
===============================================================
Can't Dial
--- FAILED ---
dial tcp 192.168.125.60:22: getsockopt: no route to host
###############################################################
Task: WhoamIandIP
###############################################################
Servers Role: Mixed
===============================================================
===============================================================
IP: 192.168.125.100
===============================================================
+++ SUCCESS +++
root
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
inet 127.0.0.1/8 scope host lo
valid_lft forever preferred_lft forever
inet6 ::1/128 scope host
valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
link/ether 08:00:27:0a:24:3f brd ff:ff:ff:ff:ff:ff
inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic enp0s3
valid_lft 83407sec preferred_lft 83407sec
inet6 fe80::a00:27ff:fe0a:243f/64 scope link
valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
link/ether 08:00:27:ab:d9:b8 brd ff:ff:ff:ff:ff:ff
inet 192.168.125.100/24 brd 192.168.125.255 scope global enp0s8
valid_lft forever preferred_lft forever
inet6 fe80::a00:27ff:feab:d9b8/64 scope link
valid_lft forever preferred_lft forever

===============================================================
IP: 192.168.125.200
===============================================================
Can't Dial
--- FAILED ---
dial tcp 192.168.125.200:22: getsockopt: no route to host
===============================================================
IP: aldoca.remote
===============================================================
+++ SUCCESS +++
root
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
inet 127.0.0.1/8 scope host lo
valid_lft forever preferred_lft forever
inet6 ::1/128 scope host
valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
link/ether 08:00:27:0a:24:3f brd ff:ff:ff:ff:ff:ff
inet 10.0.2.150/24 brd 10.0.2.255 scope global dynamic enp0s3
valid_lft 83404sec preferred_lft 83404sec
inet6 fe80::a00:27ff:fe0a:243f/64 scope link
valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
link/ether 08:00:27:ab:d9:b8 brd ff:ff:ff:ff:ff:ff
inet 192.168.125.150/24 brd 192.168.125.255 scope global enp0s8
valid_lft forever preferred_lft forever
inet6 fe80::a00:27ff:feab:d9b8/64 scope link
valid_lft forever preferred_lft forever

###############################################################
Task: WhoamIaldo
###############################################################
Servers Role: Mixed
===============================================================
===============================================================
IP: 192.168.125.100
===============================================================
+++ SUCCESS +++
aldo

===============================================================
IP: 192.168.125.200
===============================================================
Can't Dial
--- FAILED ---
dial tcp 192.168.125.200:22: getsockopt: no route to host
===============================================================
IP: aldoca.remote
===============================================================
+++ SUCCESS +++
aldoca.remote

###############################################################
Task: CheckHostname
###############################################################
Servers Role: Internal
===============================================================
===============================================================
IP: 192.168.125.100
===============================================================
+++ SUCCESS +++
test.example.local

===============================================================
IP: aldoca.local
===============================================================
+++ SUCCESS +++
aldoca.local

===============================================================
IP: 127.0.0.1
===============================================================
+++ SUCCESS +++
test.example.local
