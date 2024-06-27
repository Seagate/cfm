// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package datastore

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"

	"k8s.io/klog/v2"
)

const (
	mutexLocked = 1
)

func MutexLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

// Return our singleton dataStore
func DStore() *dataStoreHandler {
	once.Do(func() {
		dStore = &dataStoreHandler{
			filename:  DefaultDataStoreFile,
			dataStore: NewDataStore(),
		}
	})
	return dStore
}

// The Singleton instance of the DataStore
var (
	dStore *dataStoreHandler
	once   sync.Once
)

// Control for storing and restoring information to the data store
type dataStoreHandler struct {
	filename  string     // The file used to save the data store
	dataStore *DataStore // data store
	mutex     sync.Mutex // Mutex protection for concurrency
}

// GetDataStore: Return a pointer to the stored data store object
func (d *dataStoreHandler) GetDataStore() *DataStore {
	return d.dataStore
}

// NewDefaultDataStore: Create a new default DataStore object
func (d *dataStoreHandler) NewDefaultDataStore() {
	d.dataStore = NewDataStore()
}

// InitDataStore: Fill in any missing required data store
func (d *dataStoreHandler) InitDataStore() {

	// d.dataStore.SendEmail = true

	// if d.dataStore.IncomingFolder == "" {
	// 	d.dataStore.IncomingFolder = DefaultIncomingFolder
	// }
}

// Clear: Clear the data store reference
func (d *dataStoreHandler) Clear() (err error) {
	// Remove all audit items
	d.mutex.Lock()
	d.dataStore = nil
	d.mutex.Unlock()
	return nil
}

// Store: Save all data store information to a JSON file
func (d *dataStoreHandler) Store() (err error) {
	klog.V(3).InfoS("store data store data", "filename", d.filename)

	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Convert object into JSON format
	js, err := json.MarshalIndent(*d.dataStore, "", "    ")
	if err != nil {
		klog.ErrorS(err, "unable to translate data store to JSON")
		return fmt.Errorf("unable to translate data store to JSON, error: %v", err)
	}

	// Write to file
	// Set permissions so that owner can read/write (6), group can read (first 4), all others can read (second 4)
	err = os.WriteFile(d.filename, js, 0644)
	if err != nil {
		klog.ErrorS(err, "unable to write data store to file", "filename", d.filename)
		return fmt.Errorf("unable to write data store to file, filename (%s) error: %v", d.filename, err)
	}

	return nil
}

// Restore: Read data store from a JSON file, should be done once at startup
func (d *dataStoreHandler) Restore() (err error) {
	klog.V(0).InfoS("restore data store", "filename", d.filename)

	// On cfm-service startup, how handle "first-time" startup?
	// 1.) A "cfm-service installer" places a default "cfmservice.json" file in the correct location (so a failed read makes sense on the first try),
	// OR
	// 2.) Add a readfile retry, which deletes any existing "cfmservice.json" and then writes a default one (which the retry then reads)

	file, err := os.ReadFile(d.filename)
	if err != nil {
		d.NewDefaultDataStore()
		d.InitDataStore()
		os.Remove(d.filename)
		d.Store()
		file, err = os.ReadFile(d.filename)
		if err != nil {
			klog.ErrorS(err, "unable to restore data store", "filename", d.filename)
			return fmt.Errorf("unable to restore data store from file (%s), error: %v", d.filename, err)
		}
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	err = json.Unmarshal([]byte(file), d.dataStore)
	if err != nil {
		klog.ErrorS(err, "unable to unmarshal data store json file", "filename", d.filename)
		return fmt.Errorf("unable to unmarshal data store json file (%s), error: %v", d.filename, err)
	}

	d.InitDataStore()

	return nil
}

// Load: Load data store from a JSON file
func (d *dataStoreHandler) Load(filename string) (err error) {
	d.filename = filename
	klog.V(0).InfoS("load data store", "filename", d.filename)

	file, err := os.ReadFile(d.filename)
	if err != nil {
		d.NewDefaultDataStore()
		klog.ErrorS(err, "unable to load data store", "filename", d.filename)
		return fmt.Errorf("unable to load data store from file (%s), error: %v", d.filename, err)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	err = json.Unmarshal([]byte(file), d.dataStore)
	if err != nil {
		klog.ErrorS(err, "unable to unmarshal data store json file", "filename", d.filename)
		return fmt.Errorf("unable to unmarshal data store json file (%s), error: %v", d.filename, err)
	}

	d.InitDataStore()

	return nil
}
