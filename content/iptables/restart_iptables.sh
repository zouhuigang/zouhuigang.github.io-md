#!/bin/bash
iptables -I  INPUT 1 -s 139.196.48.36 -p tcp --dport 8080 -j ACCEPT
iptables -I  INPUT 2 -s 139.196.16.67 -p tcp --dport 8080 -j ACCEPT
iptables -I  INPUT 3 -s 127.0.0.1 -p tcp --dport 8080 -j ACCEPT
iptables -I  INPUT 4 -s 10.174.113.12 -p tcp --dport 8080 -j ACCEPT
iptables -I  INPUT 5 -p tcp --dport 8080 -j DROP
