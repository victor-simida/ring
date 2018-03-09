# ring
Ring data structure implemented by golang.

## InitRing
 Init a ring data structure with number.  For example ,if number is 10, InitRing will init a ring with 10 nodes which can store the events that happens in different unix timestamp(at least 10 different timestamp).

## RingAdd
 Add a event to a ring, if the timestamp is newer than ring's lastest timestamp, store the event in the oldest node.

## RingGet
 Get a node by timestamp.


## 中文
ring是一个消息环，以秒级存储，例如，在过去一秒钟内发生了两个消息，那么这两个消息依次存储在同一个node的eventlist中，一个ring最大的node数在初始化时确定，如果每个node都已经填满，则使用最旧最老的那个node存储最新的消息。，
