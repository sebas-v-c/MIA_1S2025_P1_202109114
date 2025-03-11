package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"sync"
)

// RuntimeError represents an error encountered during execution.
type RuntimeError struct {
	Line    int
	Column  int
	Command Utils.Type
	Msg     string
}

// Global variables (shared across requests)
var (
	Errors            []*RuntimeError
	CommandLogs       []*string
	MountedPartitions []*Structs.Partition
	mu                sync.Mutex // Mutex for thread safety
)

// AddError safely adds an error to the global Errors slice.
func AddError(err RuntimeError) {
	mu.Lock()
	defer mu.Unlock()
	Errors = append(Errors, &err)
}

// GetErrors returns a copy of the current errors (to prevent external modification).
func GetErrors() []*RuntimeError {
	mu.Lock()
	defer mu.Unlock()
	return append([]*RuntimeError{}, Errors...)
}

// AddError safely adds an error to the global Errors slice.
func AddCommandLog(msg string) {
	mu.Lock()
	defer mu.Unlock()
	CommandLogs = append(CommandLogs, &msg)
}

// GetErrors returns a copy of the current errors (to prevent external modification).
func GetCommandLogs() []*string {
	mu.Lock()
	defer mu.Unlock()
	return append([]*string{}, CommandLogs...)
}

// AddPartition safely adds a partition to the MountedPartitions list.
func AddPartition(p *Structs.Partition) {
	mu.Lock()
	defer mu.Unlock()
	MountedPartitions = append(MountedPartitions, p)
}

// GetPartitions returns a copy of the MountedPartitions list (to prevent race conditions).
func GetPartitions() []*Structs.Partition {
	mu.Lock()
	defer mu.Unlock()
	return append([]*Structs.Partition{}, MountedPartitions...)
}

// CleanConsole resets the global state safely.
func CleanConsole() {
	mu.Lock()
	defer mu.Unlock()
	Errors = []*RuntimeError{}
	CommandLogs = []*string{}
}
