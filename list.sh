#!/bin/bash
<services_list pv -L 1 -l --quiet | sudo ./lzr --handshakes http -sendSYNs -sourceIP 192.168.71.128 -gatewayMac 00:0c:29:a0:cf:67 -sendInterface "ens33" -blockedList "blockList"