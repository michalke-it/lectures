# Let's explore netmasks
How do source addresses change along a packet's transmission path? TBD

## How to
With the command "vagrant up", two machines will be created; 'host1', 'switch', 'router', and 'host2'.

## Network

  ┌────────┐       ┌────────┐
  │        ├───────┤        │
  │ switch │       │ router │
  └───┬────┘       └────┬───┘
      │                 │
      │                 │
  ┌───┴────┐       ┌────┴───┐
  │        │       │        │
  │  host1 │       │  host2 │
  └────────┘       └────────┘


## What you should learn
You should learn which network devices replace which address in packets that cross them. TBD

## Tools
Open vSwitch
TBD
