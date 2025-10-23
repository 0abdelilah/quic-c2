package models

import (
	"log"
	"sync"
)

var Manager = &DeviceManager{}

type DeviceManager struct {
	mu      sync.RWMutex
	Devices []Device
}

var SelectedDevice string

// adds cunstructed name, conn, stream to Devices struct
// adds the device to DeviceList
func (m *DeviceManager) Add(device Device) {

	// Add to manager slice
	m.mu.Lock()
	m.Devices = append(m.Devices, device)
	m.mu.Unlock()

	// Add to TUI device list
	App.QueueUpdateDraw(func() {
		DeviceList.AddItem(device.Name, "", 0, func() {
			SelectedDevice = device.Name
		})
	})
}

// retreive a device from devices struct by name
func (m *DeviceManager) Get(name string) (Device, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, d := range m.Devices {
		if d.Name == name {
			return d, true
		}
	}
	return Device{}, false
}

// Remove deletes a device and updates the TUI list.
func (m *DeviceManager) Remove(d Device) {
	m.mu.Lock()
	for i, dev := range m.Devices {
		if dev.Name == d.Name {
			m.Devices = append(m.Devices[:i], m.Devices[i+1:]...)
			break
		}
	}
	m.mu.Unlock()

}

// monitor checks if the device disconnects.
func (m *DeviceManager) Monitor(d Device) {
	buf := make([]byte, 1)
	for {
		if _, err := d.Stream.Read(buf); err != nil {
			log.Printf("Device %s disconnected: %v", d.Name, err)
			m.Remove(d)
			return
		}
	}
}
