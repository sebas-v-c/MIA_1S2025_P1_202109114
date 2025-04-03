package Commands

import (
	env "backend/Classes/Env"
	"backend/Classes/Interfaces"
	"backend/Classes/Utils"
	"fmt"
)

// Mounted represents the command that logs information about currently mounted partitions.
// It embeds the common CommandStruct which provides metadata like the command type,
// source code line, and column numbers.
type Mounted struct {
	Interfaces.CommandStruct // Embedded command structure for basic metadata.
}

// NewMounted creates a new instance of the Mounted command with the specified line and column numbers.
// It returns a pointer to the Mounted command.
func NewMounted(line, column int) *Mounted {
	return &Mounted{
		CommandStruct: Interfaces.CommandStruct{
			Type:   Utils.MOUNTED, // Set command type to MOUNTED.
			Line:   line,          // Source line where the command is defined.
			Column: column,        // Source column where the command is defined.
		},
	}
}

// Exec executes the Mounted command.
// It builds a console string that lists all currently loaded mounted partitions.
// If no partitions are mounted, it logs an appropriate message.
func (m *Mounted) Exec() {
	// Start the console output with a header.
	consoleString := "=================MOUNTED=================\n"

	// Check if there are any mounted partitions.
	if len(env.GetPartitions()) == 0 {
		consoleString += "No partition has been mounted...\n" +
			"=================END MOUNTED=================\n"
		// Log the console output and exit.
		m.LogConsole(consoleString)
		return
	}

	// If partitions exist, add a header for listing them.
	consoleString += "Loaded partition(s):\n"
	// Iterate through each mounted partition and append its string representation.
	for _, partition := range env.GetPartitions() {
		consoleString += fmt.Sprintf("%s\n", partition.ToString())
	}

	// End the console output with a footer.
	consoleString += "=================END MOUNTED=================\n"
	// Log the complete console output.
	m.LogConsole(consoleString)
}
