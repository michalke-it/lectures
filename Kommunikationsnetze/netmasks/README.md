# Let's explore netmasks
How do netmasks work? What happens if I change them?
This sandbox aims at answering these questions by deploying an environment in which you can safely play around with the subnetmasks.

## How to
With the command "vagrant up", two machines will be created; 'netmask1' and 'netmask2'. You can log into them via "vagrant ssh netmask1" and "vagrant ssh netmask2".

## What you should learn
You should learn that there might be different outcomes when the IP address configuration of a device is changed:
- no communication is possible anymore; this happens when both devices are in completely separate networks, i.e. fo reach device, the IP address of the other one is not inside their own network
- only unidirectional communication (e.g. UDP traffic) from one host to another is possible but not vice versa; this happens when A's IP is in B's network, but B's IP is outside of A's network
- bidirectional communication (e.g. ping)is possible; both devices have IP addresses that are inside of the other's address range (detemined by the other's netmask

## Tools
In this experiment, we use 'ping' and 'iperf3' to send traffic between the hosts. While ping requires bidirectional communication (that's why ping shows the results at the sending machine), iperf3 can generate UDP traffic that can be send unidirectional (that's why we only see the results at the receiver; there is no response to the sender, so the sender never knows if the packets were received).
