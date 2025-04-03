// Package Env provides utilities for managing the runtime environment,
// including error logging, command logging, user information, and mounted partitions.
package Env

import (
	"backend/Classes/Structs"
	"backend/Classes/Utils"
	"sync"
)

// RuntimeError represents an error encountered during execution.
// It includes the source position (line and column), the type of command that caused the error,
// and an error message.
type RuntimeError struct {
	Line    int        // Line is the line number where the error occurred.
	Column  int        // Column is the column number where the error occurred.
	Command Utils.Type // Command is the type of command that generated the error.
	Msg     string     // Msg is the detailed error message.
}

// LoggedUser holds information about the currently logged in user
// and the partition that the user has mounted.
type LoggedUser struct {
	User             Structs.User             // User contains the user details.
	MountedPartition Structs.MountedPartition // MountedPartition is the partition currently in use.
}

// Global variables shared across requests or sessions.
// They hold runtime errors, command logs, the current user, and a list of mounted partitions.
var (
	Errors            []*RuntimeError             // Errors is a slice of runtime errors.
	CommandLogs       []*string                   // CommandLogs holds log messages for executed commands.
	CurrentUser       *LoggedUser                 // CurrentUser represents the currently logged in user.
	MountedPartitions []*Structs.MountedPartition // MountedPartitions is a slice of mounted partitions.
	mu                sync.Mutex                  // mu is a mutex to ensure thread safety.
)

// AddError safely adds a new runtime error to the global Errors slice.
// It locks the shared resource to prevent concurrent access issues.
func AddError(err RuntimeError) {
	mu.Lock()
	defer mu.Unlock()
	Errors = append(Errors, &err)
}

// GetErrors returns a copy of the current runtime errors.
// A copy is returned to prevent external modifications of the global Errors slice.
func GetErrors() []*RuntimeError {
	mu.Lock()
	defer mu.Unlock()
	return append([]*RuntimeError{}, Errors...)
}

// AddCommandLog safely adds a command log message to the global CommandLogs slice.
// It locks the shared resource to ensure thread safety.
func AddCommandLog(msg string) {
	mu.Lock()
	defer mu.Unlock()
	CommandLogs = append(CommandLogs, &msg)
}

// GetCommandLogs returns a copy of the current command log messages.
// A copy is returned to prevent external modifications of the global CommandLogs slice.
func GetCommandLogs() []*string {
	mu.Lock()
	defer mu.Unlock()
	return append([]*string{}, CommandLogs...)
}

// AddPartition safely adds a mounted partition to the global MountedPartitions list.
// It locks the shared resource to ensure thread safety.
func AddPartition(p *Structs.MountedPartition) {
	mu.Lock()
	defer mu.Unlock()
	MountedPartitions = append(MountedPartitions, p)
}

// GetPartitions returns a copy of the current mounted partitions.
// A copy is returned to prevent race conditions and external modifications.
func GetPartitions() []*Structs.MountedPartition {
	mu.Lock()
	defer mu.Unlock()
	return append([]*Structs.MountedPartition{}, MountedPartitions...)
}

// CleanConsole resets the global state for errors and command logs.
// It safely clears the Errors and CommandLogs slices.
func CleanConsole() {
	mu.Lock()
	defer mu.Unlock()
	Errors = []*RuntimeError{}
	CommandLogs = []*string{}
}
