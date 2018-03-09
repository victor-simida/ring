# ring
Ring data structure implemented by golang.

## InitRing
 Init a ring data structure with number.  For example ,if number is 10, InitRing will init a ring with 10 nodes which can store the events that happens in different unix timestamp(at least 10 different timestamp).

## RingAdd
 Add a event to a ring, if the timestamp is newer than ring's lastest timestamp, store the event in the oldest node.

## RingGet
 Get a node by timestamp.
