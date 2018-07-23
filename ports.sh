#!/bin/bash
sudo iptables -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4010
sudo iptables -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4011
sudo ip6tables -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4010
sudo ip6tables -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4011