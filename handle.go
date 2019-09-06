package main

import (
	"fmt"
	"sync"

	"github.com/gautamrege/gochat/api"
)

// Ensure that users are added / removed using a mutex!
type PeerHandleMapSync struct {
	sync.RWMutex
	PeerHandleMap map[string]api.Handle
}

// Insert user if not exists already then add it
func (hs *PeerHandleMapSync) Insert(newHandle api.Handle) (err error) {
	hs.Lock()

	// TODO-WORKSHOP-STEP-3: This code should insert the handle into the PeerHandleMap
	_, ok := hs.PeerHandleMap[newHandle.Name]
	if !ok {
		hs.PeerHandleMap[newHandle.Name] = newHandle
		fmt.Printf("New user joined chat: @%s \n", newHandle.Name)
	}

	hs.Unlock()
	return nil
}

// get the user details from the map with given name
func (hs *PeerHandleMapSync) Get(name string) (handle api.Handle, ok bool) {
	hs.Lock()
	// TODO-WORKSHOP-STEP-4: This code should fetch the handle from the PeerHandleMap based on the key name
	// TODO-THINK: Why is this in a Lock() method?
	handle, ok = hs.PeerHandleMap[name]
	hs.Unlock()
	return
}

// delete the user from map
func (hs *PeerHandleMapSync) Delete(name string) {
	hs.Lock()
	// TODO-WORKSHOP-STEP-5: This code should remove the handle from the PeerHandleMap based on the key name
	delete(hs.PeerHandleMap, name)
	hs.Unlock()
	fmt.Println("UserHandle Removed for ", name)
}

func String(h api.Handle) string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs PeerHandleMapSync) String() string {
	var users string
	// TODO-WORKSHOP-STEP-6: This code should print the list of all names of the handles in the map
	// TODO-THINK: Do we need a Lock here?
	hs.Lock()
	defer hs.Unlock()

	for name, _ := range hs.PeerHandleMap {
		users += fmt.Sprintf("@%s\n", name)
	}

	return users
}
