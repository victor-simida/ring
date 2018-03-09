package ring

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Event struct {
	Timestamp int64
	Value     interface{}
}

type Node struct {
	Previous  *Node
	Next      *Node
	EventList []Event
	Timestamp int64
}

func (n *Node) String() string {
	return fmt.Sprintf("%v:%+v", n.Timestamp, n.EventList)
}

type Ring struct {
	sync.RWMutex
	Number int
	Now    *Node
}

func InitRing(number int) *Ring {
	var result Ring
	result.Number = number
	head := new(Node)
	head.Timestamp = time.Now().Unix()
	result.Now = head
	temp := head
	for i := 1; i < number; i++ {
		temp.Next = new(Node)
		temp.Next.Previous = temp
		temp = temp.Next
	}
	temp.Next = head
	head.Previous = temp

	return &result
}

func (r *Ring) RingAdd(value interface{}) {
	r.Lock()

	now := time.Now().Unix()
	if r.Now.Timestamp != now {
		r.Now = r.Now.Next
		r.Now.Timestamp = now
		r.Now.EventList = []Event{}
	}
	r.Now.EventList = append(r.Now.EventList, Event{
		Timestamp: time.Now().Unix(),
		Value:     value,
	})

	r.Unlock()
}

func (r *Ring) RingGet(timestamp int64) (*Node, bool) {
	r.RLock()
	defer r.RUnlock()
	temp := r.Now

	if r.Now.Timestamp < timestamp {
		return nil, false
	}

	for i := 0; i < r.Number; i++ {
		if temp.Timestamp == timestamp {
			return temp, true
		}
		temp = temp.Previous
		continue

	}

	return nil, false
}

func (r *Ring) String() string {
	r.RLock()
	defer r.RUnlock()
	var result []string
	temp := r.Now.Next
	for i := 0; i < r.Number; i++ {
		result = append(result, temp.String())
		temp = temp.Next
	}

	return strings.Join(result, ";")
}
