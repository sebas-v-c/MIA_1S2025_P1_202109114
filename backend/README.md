<!-- Auto-generated Markdown from Godoc -->

# Application Architecture Overview

This document describes how the application’s frontend, backend, and underlying disk resources interact based on the provided diagram.

![Application Architecture Diagram](/docs/diagram.png "Architecture Diagram")

---

## High-Level Flow

1. **Frontend (React Application)**

   - The user interacts with a React-based web client.
   - The client allows the user to submit code or commands, typically via a form or code editor in the UI.

2. **Backend (API)**

   - The frontend sends a POST request to `http://localhost:8090/parse` (or a similar endpoint).
   - The request body usually contains JSON data, such as:
     ```json
     {
       "code": "mkdir -p /folder"
     }
     ```
   - The backend receives the request, processes the code or commands, and may interact with one or more mounted disks.
   - The backend responds with a JSON payload indicating success, logs, or errors:
     ```json
     {
       "logs": "success",
       "errors": "syntactic error",
       "runtime": "runtime error"
     }
     ```
   - These results are then used by the frontend to display feedback to the user.

3. **Disks (Filesystem Resources)**
   - The backend manages several “disks” (or disk images) which can be mounted, unmounted, and manipulated by commands.
   - The backend’s command parser (or interpreter) reads and writes data to these disks—creating directories, files, or performing other filesystem operations.
   - The user’s commands (e.g., `mkdir`, `rm`, `mkfs`) are executed against these virtual disks, and the results are then sent back through the API response.

---

## Detailed Interactions

1. **User Action (Frontend)**

   - A user writes or selects a command (like `mkdir -p /folder`) in the React app.
   - When they submit, the frontend constructs a JSON payload:
     ```json
     {
       "code": "mkdir -p /folder"
     }
     ```
   - The frontend makes an HTTP POST to `http://localhost:8090/parse`.

2. **Command Parsing (Backend)**

   - The backend exposes an API endpoint (`/parse`) to handle code or command submissions.
   - Upon receiving the request:
     1. The backend reads the `"code"` field from the JSON.
     2. The command parser or interpreter checks syntax, validity, and handles any arguments.
     3. If needed, the parser executes the command against the internal data structures or mounted disks.

3. **Filesystem Operations (Disks)**

   - If the command involves disk manipulation (e.g., creating a directory, listing partitions, mounting a disk), the backend accesses one or more disk files.
   - The backend might:
     - Open a disk file.
     - Read or write a Master Boot Record (MBR).
     - Update metadata in a superblock.
     - Create or modify directories and files inside the disk’s data structures.
   - Any changes or read operations on the disk are done at this step.

4. **Result Logging and Error Handling (Backend)**

   - The backend collects logs (e.g., command success messages, debugging info).
   - If any errors occur (syntax errors, runtime errors, insufficient permissions), the backend includes them in the response.
   - The response body might look like:
     ```json
     {
       "logs": "Directory '/folder' created successfully.",
       "errors": "",
       "runtime": ""
     }
     ```
   - Alternatively, if an error occurred:
     ```json
     {
       "logs": "",
       "errors": "Syntax error near '-p'",
       "runtime": ""
     }
     ```

5. **Response to Frontend**
   - The backend returns a JSON response containing logs and errors.
   - The frontend receives the response and updates the UI accordingly—showing success messages or errors to the user.

---

## Diagram Reference

In the diagram:

1. **Frontend (React)**

   - Depicted as a window labeled “React” on the left.
   - Sends a POST request to the API with `{"code":"mkdir -p /folder"}` or similar.

2. **API**

   - Shown as a server in the center with an “API” label.
   - It processes incoming requests on the `/parse` endpoint.
   - It returns a JSON response like `{"logs":"success","errors":"syntactic error","runtime":"runtime error"}`.

3. **Disks**
   - On the right side, the diagram shows multiple disk icons labeled “DISC.”
   - The API has direct access to these disks.
   - Commands that manipulate the filesystem on these disks are executed by the API’s command logic.

---

## Summary

- The **frontend** (React) is responsible for providing a user interface to enter and submit commands.
- The **backend** (API) receives these commands, parses them, and executes the requested operations (e.g., creating directories, mounting disks) on the **disks** (which are managed by the backend).
- The result (success, logs, or errors) is returned to the frontend in JSON format, allowing the user to see immediate feedback about their commands.

# Project Packages

# package Commands

```go
import "backend/Classes/Commands"
```

Package Commands provides implementations for various filesystem commands.

## FUNCTIONS

```go
func CreateDirs(dirTree *Structs.DirTree, splitedPath []string, createDirs bool, consoleString *strings.Builder, file *os.File, superBlock Structs.SuperBlock) (error, strings.Builder)
    CreateDirs is a helper function that creates directories along a specified
    path. It starts from the root ("/") and iterates through each directory in
    the provided path. If a directory does not exist and the createDirs flag is
    set (or if it's the last directory), it creates the directory and updates
    the parent inode accordingly. It returns an error (if any) and the updated
    console log.
```

## TYPES

### CAT

```go
type Cat struct {
        Interfaces.CommandStruct          // Embeds common command metadata (Type, Line, Column, etc.)
        Files                    []string // Files is a slice of file paths to be displayed.
}
    Cat represents the CAT command, which is used to display the contents of one
    or more files. It embeds the common CommandStruct for shared metadata and
    behavior, and holds a slice of file paths to be read.
```

```go
func NewCat(line, column int, files []string) *Cat
    NewCat creates a new instance of the Cat command with the specified source
    location and list of file paths.
```

```go
func (c *Cat) Exec()
    Exec executes the CAT command. It verifies that a user is logged in,
    checks the disk integrity, and then iterates over the provided file paths.
    For each file, it verifies read permissions, retrieves the file content,
    updates the access time, and appends the content to the console log.
```

### CHGRP

```go
type Chgrp struct {
        Interfaces.CommandStruct                   // Embedded command structure for metadata (Type, Line, Column, etc.)
        Params                   map[string]string // Params holds the command parameters (e.g., "user" and "grp").
}
    Chgrp represents the CHGRP command, which is used to change the group
    associated with a user in the "/users.txt" file. It embeds the common
    CommandStruct for shared command metadata and behavior, and holds a map of
    command-specific parameters.
```

```go
func NewChgrp(line, column int, params map[string]string) *Chgrp
    NewChgrp creates a new instance of the Chgrp command with the specified
    source location and parameters.
```

```go
func (c *Chgrp) Exec()
    Exec executes the CHGRP command. It validates parameters, ensures the logged
    user is root, verifies the disk integrity, retrieves the "/users.txt" file,
    and then updates the group for the specified user.
```

### FDISK

```go
type Fdisk struct {
        Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
        Params                   map[string]string // Params holds command-specific parameters like "size", "path", "name", "type", "fit", and "unit".
}
    Fdisk represents the FDISK command, which is used to create a new partition
    on a disk file. It embeds the common CommandStruct for shared command
    metadata and behavior, and holds a map of command parameters.
```

```go
func NewFdisk(line, column int, params map[string]string) *Fdisk
    NewFdisk creates a new Fdisk command instance with the specified source
    location and parameters.
```

```go
func (f *Fdisk) Exec()
    Exec executes the FDISK command to create a new partition. It validates
    parameters, opens the disk file, reads the disk MBR, and creates a new
    primary, extended, or logical partition based on the provided parameters.
    The command calculates the appropriate start position and size for the
    partition, updates the MBR, and (if needed) creates an EBR for logical
    partitions.
```

### LOGIN

```go
type Login struct {
        Interfaces.CommandStruct                   // Embeds common command metadata (Type, Line, Column, etc.)
        Params                   map[string]string // Params contains command-specific parameters (e.g., "user", "pass", "id").
}
    Login represents the LOGIN command, which is used to authenticate a user
    against the filesystem's user records stored in "/users.txt". It embeds the
    common CommandStruct for shared behavior and holds a map of parameters.
```

```go
func NewLogin(line, column int, params map[string]string) *Login
    NewLogin creates a new instance of the Login command with the specified
    source location and parameters.
```

```go
func (l *Login) Exec()
    Exec executes the LOGIN command. It validates the provided parameters,
    retrieves the mounted partition, checks the partition's integrity,
    and attempts to locate and authenticate the user using the information
    from the "/users.txt" file. If successful, it sets the global current user;
    otherwise, it logs an appropriate error.
```

### LOGOUT

```go
type Logout struct {
        Interfaces.CommandStruct // Embeds common command behavior and metadata.
}
    Logout represents the LOGOUT command, which is used to log out the currently
    logged-in user. It embeds the common CommandStruct to inherit shared
    metadata (such as Type, Line, and Column).
```

```go
func NewLogout(line, column int) *Logout
    NewLogout creates a new instance of the Logout command with the specified
    source location. It initializes the command with the LOGOUT type, and sets
    the line and column where the command is defined.
```

```go
func (l *Logout) Exec()
    Exec executes the LOGOUT command. It checks whether a user is currently
    logged in; if not, it logs an error. If a user is logged in, it sets the
    global current user to nil (logging out), then logs a success message to the
    console.
```

### MKDIR

```go
type Mkdir struct {
        Interfaces.CommandStruct                   // Embeds common command metadata (Type, Line, Column, etc.)
        Params                   map[string]string // Params holds command parameters (e.g., "path").
}
    Mkdir represents the MKDIR command, which is used to create a new directory
    (or directories) in the filesystem. It embeds the common CommandStruct to
    inherit shared metadata and behavior, and holds a map of parameters for the
    command.
```

```go
func NewMkdir(line, column int, params map[string]string) *Mkdir
    NewMkdir creates a new instance of the Mkdir command with the given source
    location and parameters.
```

```go
func (m *Mkdir) Exec()
    Exec executes the MKDIR command to create a new directory. It validates
    parameters, verifies disk integrity, and then attempts to create the
    directory (or directories) along the provided path. The final updated
    superblock is written to disk, and a log is generated.
```

### MKDISK

```go
type Mkdisk struct {
        Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
        Params                   map[string]string // Params holds command-specific parameters such as "size", "path", "fit", and "unit".
}
    Mkdisk represents the MKDISK command, which is used to create a new
    disk file with a specified size and configuration. It embeds the common
    CommandStruct to inherit shared metadata and behavior, and holds a map for
    command parameters.
```

```go
func NewMkdisk(line, column int, params map[string]string) *Mkdisk
    NewMkdisk creates a new Mkdisk command instance with the specified source
    location and parameters.
```

```go
func (m *Mkdisk) Exec()
    Exec executes the MKDISK command. It validates the parameters, creates a
    disk file with the specified size, writes empty data in batches to the file,
    creates a new MBR structure, and writes the MBR to the disk. Finally,
    it logs the outcome.
```

### MKFILE

```go
type Mkfile struct {
        Interfaces.CommandStruct                   // Embedded command structure providing metadata (Type, Line, Column, etc.)
        Params                   map[string]string // Params holds command-specific parameters, such as "path", "size", and "cont".
}
    Mkfile represents the MKFILE command, which is used to create a new file in
    the filesystem. It embeds the common CommandStruct for shared behavior and
    includes a map for command parameters.
```

```go
func NewMkfile(line, column int, params map[string]string) *Mkfile
    NewMkfile creates a new Mkfile command instance with the given source
    location and parameters.
```

```go
func (m *Mkfile) Exec()
    Exec executes the MKFILE command to create a new file in the filesystem.
    It validates parameters, verifies disk integrity, optionally creates missing
    directories, prepares file content (from a given size or an external file),
    locates the parent directory inode, and creates a new file inode with the
    specified content. The updated inode and file content are then written to
    disk, and a log message is generated.
```

### MKFS

```go
type Mkfs struct {
        Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
        Params                   map[string]string // Params contains command parameters such as "id" and "type".
}
    Mkfs represents the MKFS command, which is used to format a mounted
    partition with an EXT2 filesystem. It embeds the common CommandStruct for
    shared metadata and includes a map for command parameters.
```

```go
func NewMkfs(line, column int, params map[string]string) *Mkfs
    NewMkfs creates a new Mkfs command instance with the specified source
    location and parameters.
```

```go
func (m *Mkfs) Exec()
    Exec executes the MKFS command to format the partition with an EXT2
    filesystem. It validates parameters, verifies the mounted partition,
    calculates filesystem parameters, creates a new superblock, initializes
    inodes and blocks, creates root and users files, updates bitmaps, and writes
    the final superblock back to disk.
```

### MKGRP

```go
type Mkgrp struct {
        Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
        Params                   map[string]string // Params contains command parameters, e.g., "name" of the group.
}
    Mkgrp represents the MKGRP command, which is used to create a new group.
    It embeds the common CommandStruct to inherit shared metadata and behavior,
    and holds a map of parameters required for creating the group.
```

```go
func NewMkgrp(line, column int, params map[string]string) *Mkgrp
    NewMkgrp creates a new Mkgrp command instance with the specified source
    location and parameters.
```

```go
func (m *Mkgrp) Exec()
    Exec executes the MKGRP command to create a new group. It validates
    parameters, ensures that the current user is root, verifies disk integrity,
    locates the "/users.txt" inode that stores user and group entries, and then
    appends a new group entry if it does not already exist.
```

### MKUSR

```go
type Mkusr struct {
        Interfaces.CommandStruct                   // Embedded command structure (provides Type, Line, Column, etc.)
        Params                   map[string]string // Params contains command parameters such as "user", "pass", and "grp".
}
    Mkusr represents the MKUSR command, which is used to create a new user.
    It embeds the common CommandStruct to inherit shared metadata and behavior,
    and it holds a map of parameters required for user creation.
```

```go
func NewMkusr(line, column int, params map[string]string) *Mkusr
    NewMkusr creates a new Mkusr command instance with the specified source
    location and parameters.
```

```go
func (m *Mkusr) Exec()
    Exec executes the MKUSR command to create a new user. It validates
    parameters, checks that the current user is root, verifies disk integrity,
    locates the "/users.txt" inode, and appends a new user entry to that file.
```

### MOUNT

```go
type Mount struct {
        Interfaces.CommandStruct                   // Embeds common command metadata (type, line, column, etc.)
        Params                   map[string]string // Params holds the command parameters, such as "path" and "name".
}
    Mount represents the MOUNT command, which is used to mount a disk. It embeds
    the common CommandStruct for shared behavior and includes a map for command
    parameters.
```

```go
func NewMount(line, column int, params map[string]string) *Mount
    NewMount creates a new Mount command instance with the provided source
    position and parameters.
```

```go
func (m *Mount) Exec()
    Exec executes the MOUNT command. It validates parameters, opens the disk
    file, reads the MBR, and mounts a primary partition. It then updates the
    partition's status and ID, adds the partition to the global RAM partitions,
    writes the updated MBR back to disk, and logs the operation.
```

### MOUNTED

```go
type Mounted struct {
        Interfaces.CommandStruct // Embedded command structure for basic metadata.
}
    Mounted represents the command that logs information about currently mounted
    partitions. It embeds the common CommandStruct which provides metadata like
    the command type, source code line, and column numbers.
```

```go
func NewMounted(line, column int) *Mounted
    NewMounted creates a new instance of the Mounted command with the specified
    line and column numbers. It returns a pointer to the Mounted command.
```

```go
func (m *Mounted) Exec()
    Exec executes the Mounted command. It builds a console string that lists
    all currently loaded mounted partitions. If no partitions are mounted,
    it logs an appropriate message.
```

### REP

```go
type Rep struct {
        Interfaces.CommandStruct
        Params map[string]string
}
```

```go
func NewRep(line, column int, params map[string]string) *Rep
```

```go
func (r *Rep) Exec()
```

### RMDISK

```go
type Rmdisk struct {
        Interfaces.CommandStruct                   // Embedded command structure providing metadata (type, line, column, etc.)
        Params                   map[string]string // Params holds the command parameters (e.g., file path for removal).
}
    Rmdisk represents the command to remove a disk file. It embeds the common
    CommandStruct for basic command metadata and includes a map for command
    parameters.
```

```go
func NewRmdisk(line, column int, params map[string]string) *Rmdisk
    NewRmdisk creates a new instance of the Rmdisk command with the specified
    source position and parameters.
```

```go
func (r *Rmdisk) Exec()
    Exec executes the Rmdisk command logic. It validates parameters, then
    attempts to remove the specified disk file. If the removal is successful,
    it logs a success message; otherwise, it appends an error.
```

### RMGRP

```go
type Rmgrp struct {
        Interfaces.CommandStruct                   // Embedded command structure providing basic metadata.
        Params                   map[string]string // Params holds the command parameters (e.g., the group name to remove).
}
    Rmgrp represents the command to remove a group. It embeds the common
    CommandStruct and includes a map of parameters.
```

```go
func NewRmgrp(line, column int, params map[string]string) *Rmgrp
    NewRmgrp creates a new instance of the Rmgrp command with the specified
    source position and parameters.
```

```go
func (r *Rmgrp) Exec()
    Exec executes the Rmgrp command logic. It validates parameters, ensures that
    the current user is root, verifies disk integrity, locates the "/users.txt"
    inode (which also contains group entries), and then marks the specified
    group as removed in the file content.
```

### RMUSR

```go
type Rmusr struct {
        Interfaces.CommandStruct                   // Embedded command structure providing basic command metadata.
        Params                   map[string]string // Params holds the command parameters, e.g., the username to remove.
}
    Rmusr represents the command to remove a user. It embeds the common
    CommandStruct and includes a map of parameters.
```

```go
func NewRmusr(line, column int, params map[string]string) *Rmusr
    NewRmusr creates a new instance of the Rmusr command with the given source
    position and parameters.
```

```go
func (r *Rmusr) Exec()
    Exec executes the Rmusr command logic. It validates parameters, checks if
    the current user is root, verifies disk integrity, locates the "/users.txt"
    inode, and then marks the specified user as removed in the file content.
```

# package Env

```go
import "backend/Classes/Env"
```

Package Env provides utilities for handling disk and partition status.

## VARIABLES

```go
var (
        Errors            []*RuntimeError             // Errors is a slice of runtime errors.
        CommandLogs       []*string                   // CommandLogs holds log messages for executed commands.
        CurrentUser       *LoggedUser                 // CurrentUser represents the currently logged in user.
        MountedPartitions []*Structs.MountedPartition // MountedPartitions is a slice of mounted partitions.

)
    Global variables shared across requests or sessions. They hold runtime
    errors, command logs, the current user, and a list of mounted partitions.
```

## FUNCTIONS

```go
func AddCommandLog(msg string)
    AddCommandLog safely adds a command log message to the global CommandLogs
    slice. It locks the shared resource to ensure thread safety.
```

```go
func AddError(err RuntimeError)
    AddError safely adds a new runtime error to the global Errors slice.
    It locks the shared resource to prevent concurrent access issues.
```

```go
func AddPartition(p *Structs.MountedPartition)
    AddPartition safely adds a mounted partition to the global MountedPartitions
    list. It locks the shared resource to ensure thread safety.
```

```go
func CalculatePermissions(perm byte) (bool, bool, bool)
    CalculatePermissions converts a permission byte (expressed as a character
    '0'-'7') into boolean values representing read, write, and execute
    permissions. The permission digit corresponds to a specific combination:
    '0': no permissions, '1': execute, '2': write, '3': write and execute,
    '4': read, '5': read and execute, '6': read and write, '7': read, write and
    execute.
```

```go
func CheckFilePermissions(loggedUser LoggedUser, fileInode *Structs.Inode) (bool, bool, bool, error)
    CheckFilePermissions checks if the given logged user has read, write,
    and execute permissions for the file represented by the provided inode. It
    returns three booleans corresponding to read, write, and execute permissions
    respectively, along with an error if one occurs. The function determines
    permissions based on the file's owner (UID), group (GID), and a lookup of
    the user's group ID from the mounted partition.
```

```go
func CleanConsole()
    CleanConsole resets the global state for errors and command logs. It safely
    clears the Errors and CommandLogs slices.
```

```go
func GetCommandLogs() []*string
    GetCommandLogs returns a copy of the current command log messages. A copy is
    returned to prevent external modifications of the global CommandLogs slice.
```

```go
func GetPartitions() []*Structs.MountedPartition
    GetPartitions returns a copy of the current mounted partitions. A copy is
    returned to prevent race conditions and external modifications.
```

```go
func GetUserGID(groupName string, mountedPartition Structs.MountedPartition) (int32, error)
    GetUserGID retrieves the group ID (GID) for a user based on the provided
    group name and mounted partition. It opens the file system corresponding
    to the mounted partition, reads the superblock and directory tree,
    and then searches for the "/users.txt" file, which contains user and group
    information. If found, the function returns the group ID as an int32;
    otherwise, it returns an error.
```

```go
func VerifyDiscStatus(partition [4]byte) (*Structs.MountedPartition, *Structs.Partition, *os.File, error)
    VerifyDiscStatus checks the status of a disk partition given its ID.
    It attempts to retrieve the mounted partition from RAM using the provided
    partition ID, then validates the health of the partition by reading the
    disk's MBR. It returns pointers to the mounted partition, its corresponding
    MBR partition entry, the disk file, and an error if any step fails.
```

## TYPES

### LoggedUser

```go
type LoggedUser struct {
        User             Structs.User             // User contains the user details.
        MountedPartition Structs.MountedPartition // MountedPartition is the partition currently in use.
}
    LoggedUser holds information about the currently logged in user and the
    partition that the user has mounted.
```

### RuntimeError

```go
type RuntimeError struct {
        Line    int        // Line is the line number where the error occurred.
        Column  int        // Column is the column number where the error occurred.
        Command Utils.Type // Command is the type of command that generated the error.
        Msg     string     // Msg is the detailed error message.
}
    RuntimeError represents an error encountered during execution. It includes
    the source position (line and column), the type of command that caused the
    error, and an error message.
```

```go
func GetErrors() []*RuntimeError
    GetErrors returns a copy of the current runtime errors. A copy is returned
    to prevent external modifications of the global Errors slice.
```

# Package Interfaces

```go
import "backend/Classes/Interfaces"
```

Package Interfaces defines the interfaces and shared structures for executable
commands.

## TYPES

### Command

```go
type Command interface {
        // GetLine returns the line number where the command is defined.
        GetLine() int

// GetColumn returns the column number where the command is defined.
        GetColumn() int

// GetType returns the type of the command.
        GetType() Utils.Type

// Exec executes the command's logic.
        Exec()
}
    Command defines the interface for all executable commands. Each command must
    provide methods for retrieving its source position, determining its type,
    and executing its logic.
```

### CommandStruct

```go
type CommandStruct struct {
        Type   Utils.Type // Type is the type of the command.
        Line   int        // Line is the line number where the command is defined.
        Column int        // Column is the column number where the command is defined.
}
    CommandStruct provides shared behavior for commands. It serves as a minimal
    base structure that includes the command type, and the line and column
    numbers where the command is defined.
```

```go
func (c *CommandStruct) AppendError(msg string)
    AppendError logs a runtime error associated with this command. It creates a
    RuntimeError entry with the command's source position and type.
```

```go
func (c *CommandStruct) Exec()
    Exec provides a default execution method for commands. This method should be
    overridden by specific command implementations. Calling this default method
    results in a runtime panic.
```

```go
func (c *CommandStruct) GetColumn() int
    GetColumn returns the column number where the command is defined.
```

```go
func (c *CommandStruct) GetLine() int
    GetLine returns the line number where the command is defined.
```

```go
func (c *CommandStruct) GetType() Utils.Type
    GetType returns the type of the command.
```

```go
func (c *CommandStruct) LogConsole(msg string)
    LogConsole logs a message to the command log console.
```

# package Structs

```go
import "backend/Classes/Structs"
```

Package Structs provides data structures for representing various filesystem
components.

## TYPES

### Content

```go
type Content struct {
        Name  [12]byte // Name holds the content's name as a fixed-size byte array.
        Inode int32    // Inode is the index of the inode associated with this content entry.
}
    Content represents an entry within a folder block. It contains a fixed-size
    name and an inode number that points to the associated file or directory.
```

```go
func (c *Content) ToString() string
    ToString returns a formatted string representation of the Content.
    It converts the Name field from a byte array to a string and displays the
    inode number.
```

### EBR

```go
type EBR struct {
        Mount byte     // Mount indicates the mount status of the partition ('0' for unmounted).
        Fit   byte     // Fit specifies the partition's fit type (e.g., first, best, or worst fit).
        Start int32    // Start is the starting byte position of the partition.
        Size  int32    // Size indicates the total size of the partition in bytes.
        Next  int32    // Next points to the starting byte of the next EBR (-1 if none).
        Name  [16]byte // Name holds the partition's name.
}
    EBR represents an Extended Boot Record. It stores metadata for a logical
    partition including its mount status, fit type, starting position, size,
    pointer to the next EBR, and its name.
```

```go
func NewEBR(fit byte, start, size int32, name [16]byte) *EBR
    NewEBR creates a new Extended Boot Record with the provided fit type,
    starting position, size, and name. The Mount field is initialized to '0'
    (unmounted) and Next is set to -1, indicating that there is no subsequent
    EBR.
```

```go
func (e *EBR) ToString() string
    ToString returns a formatted string representation of the EBR. It includes
    details such as mount status, start position, size, name, and the pointer to
    the next EBR.
```

### FILEBLOCK

```go
type FileBlock struct {
        Content [64]byte // Content holds the raw file data.
}
    FileBlock represents a block of data for file storage. It contains a
    fixed-size array of 64 bytes that holds file content.
```

```go
func (fb *FileBlock) ToString() string
    ToString returns a formatted string representation of the FileBlock's
    content. It trims the byte array at the first null byte (0) and returns the
    resulting string. This is useful for converting the fixed-size byte array
    into a human-readable string.
```

### FOLDERBLOCK

```go
type FolderBlock struct {
        Content [4]Content // Content is an array of 4 Content items representing directory entries.
}
    FolderBlock represents a block in the filesystem that holds directory
    content. It contains an array of Content structures, each representing an
    entry in a folder.
```

```go
func NewFolderBlock() *FolderBlock
    NewFolderBlock creates and returns a new FolderBlock instance. It
    initializes the Content array with default values where each Content's Inode
    is set to -1, indicating that the entry is unused.
```

```go
func (fb *FolderBlock) ToString() string
    ToString returns a formatted string representation of the FolderBlock.
    It iterates over the Content array and concatenates each entry's string
    representation.
```

### INODE

```go
type Inode struct {
        UID    int32     // UID is the user identifier that owns the inode.
        GID    int32     // GID is the group identifier that owns the inode.
        Size   int32     // Size represents the size of the file or directory in bytes.
        ATime  [17]byte  // ATime is the last access time in "YYYY-MM-DD HH:MM" format.
        CTime  [17]byte  // CTime is the creation time in "YYYY-MM-DD HH:MM" format.
        MTime  [17]byte  // MTime is the last modification time in "YYYY-MM-DD HH:MM" format.
        IBlock [15]int32 // IBlock is an array of pointers to data blocks.
        Type   [1]byte   // Type denotes the inode type: '0' for folder and '1' for file.
        Perm   [3]byte   // Perm represents the permission settings (e.g., "644").
}
    Inode represents a file or directory inode in the filesystem. It contains
    metadata such as the user and group IDs, file size, timestamps for access,
    creation, and modification, data block pointers, the type of the inode,
    and its permissions.
```

```go
func NewInode(inodeType [1]byte) *Inode
    NewInode creates a new inode with default parameters. The inodeType
    parameter must be '0' for a folder and '1' for a file. It initializes the
    timestamps to the current time and sets default permissions.
```

```go
func (i *Inode) ToString() string
    ToString returns a formatted string representation of the inode. It includes
    details such as UID, GID, size, timestamps, block pointers, inode type,
    and permissions.
```

### MBR

```go
type MBR struct {
        MbrSize      int32        // MbrSize indicates the total size of the MBR in bytes.
        CreationDate [10]byte     // CreationDate holds the creation date in "YYYY-MM-DD" format.
        Signature    int32        // Signature is a randomly generated identifier for the MBR.
        Fit          [1]byte      // Fit represents the partition fit type (e.g., first, best, or worst fit).
        Partitions   [4]Partition // Partitions is an array containing up to 4 partitions.
}
    MBR represents the Master Boot Record of a disk. It stores information about
    the disk size, creation date, a unique signature, the partition fit type,
    and an array of partitions.
```

```go
func NewMBR(size int32, fit [1]byte) *MBR
    NewMBR creates a new MBR instance with the provided size and fit parameters.
    It initializes the creation date with the current date, generates a random
    signature, and sets up an empty partitions array.
```

```go
func (m *MBR) ToString() string
    ToString returns a formatted string representation of the MBR. The string
    includes details such as the MBR size, creation date, fit type, and a list
    of each partition's string representation.
```

### MOUNTEDPARTITION

```go
type MountedPartition struct {
        Partition            // Embedded Partition struct.
        DiscSignature int32  // DiscSignature is a numeric signature for the disk.
        DiscTag       rune   // DiscTag is a character identifier for the disk.
        Path          string // Path is the filesystem path where the partition is mounted.
}
    MountedPartition represents a partition that has been mounted on the system.
    It embeds a Partition and adds disk-specific information such as disk
    signature, disk tag, and the mount path.
```

```go
func (m *MountedPartition) ToString() string
    ToString returns a formatted string representation of the MountedPartition.
    It includes both the partition details and the disk-specific metadata.
```

### PARTITION

```go
type Partition struct {
        Status      [1]byte  // Status indicates the partition's status (e.g., active/inactive).
        Type        [1]byte  // Type specifies the partition type (primary, extended, etc.).
        Fit         [1]byte  // Fit defines the partition's fit type (best, first, or worst fit).
        Start       int32    // Start is the starting byte position of the partition on the disk.
        Size        int32    // Size indicates the total size of the partition in bytes.
        Name        [16]byte // Name holds the partition's name.
        Correlative int32    // Correlative is a sequential number used for identification.
        Id          [4]byte  // Id is a unique identifier for the partition.
}
    Partition represents a disk partition. It holds metadata including its
    status, type, fit, starting position, size, name, a correlative number,
    and a unique identifier.
```

```go
func NewPartition(status, type_, fit [1]byte, start, size int32, name [16]byte, correlative int32) *Partition
    NewPartition creates a new Partition with the provided parameters.
    It returns a pointer to the newly created Partition.
```

```go
func (p *Partition) ToString() string
    ToString returns a formatted string representation of the Partition. The
    output includes details such as start position, status, size, name, type,
    fit, ID, and correlative number.
```

### POINTERBLOCK

```go
type PointerBlock struct {
        Pointers [16]int32 // Pointers holds an array of pointer values.
}
    PointerBlock represents a block of pointers in the filesystem. It contains
    an array of 16 pointers represented as int32 values.
```

```go
func (p *PointerBlock) ToString() string
    ToString returns a string representation of the PointerBlock. It formats the
    Pointers array into a readable string.
```

### SUPERBLOCK

```go
type SuperBlock struct {
        FilesystemType  int32    // FilesystemType indicates the type of the filesystem.
        InodesCount     int32    // InodesCount is the total number of inodes in the filesystem.
        BlocksCount     int32    // BlocksCount is the total number of blocks in the filesystem.
        FreeBlocksCount int32    // FreeBlocksCount is the number of free blocks available.
        FreeInodesCount int32    // FreeInodesCount is the number of free inodes available.
        MTime           [17]byte // MTime represents the last modification time (e.g., "2025-09-20 12:00:00").
        UMTime          [17]byte // UMTime represents the last unmount time.
        MntCount        int32    // MntCount is the count of how many times the filesystem has been mounted.
        Magic           int32    // Magic holds a magic number for identifying or validating the filesystem.
        InodeSize       int32    // InodeSize specifies the size of each inode in bytes.
        BlockSize       int32    // BlockSize specifies the size of each block in bytes.
        FirstInode      int32    // FirstInode is the index of the first inode.
        FirstBlock      int32    // FirstBlock is the index of the first block.
        BmInodeStart    int32    // BmInodeStart is the starting position of the inode bitmap.
        BmBlockStart    int32    // BmBlockStart is the starting position of the block bitmap.
        InodeStart      int32    // InodeStart is the starting position of the inodes.
        BlockStart      int32    // BlockStart is the starting position of the blocks.
}
    SuperBlock holds the metadata for a filesystem's superblock. It contains
    various fields describing the filesystem's properties such as inode and
    block counts, timestamps, and layout information.
```

```go
func (sb *SuperBlock) ToString() string
    ToString returns a formatted multi-line string representation of the
    SuperBlock. It formats all the fields of the superblock into a readable
    string.
```

### USER

```go
type User struct {
        Id       int32  // Id is the unique identifier for the user.
        Group    string // Group represents the user's group or role.
        Name     string // Name is the user's full name.
        Password string // Password stores the user's password.
}
    User represents a user in the system. It contains the user's ID, group,
    name, and password.
```

```go
func (u *User) ToString() string
    ToString returns a string representation of the User. The output format is:
    "<Id>,U,<Group>,<Name>,<Password>"
```

### DirTree

```go
type DirTree struct {
        SuperBlock  *SuperBlock // Pointer to the superblock of the filesystem.
        File        *os.File    // Reference to the file representing the disk.
        BlockBitMap []byte      // BlockBitMap indicates the allocation status of each block.
        InodeBitMap []byte      // InodeBitMap indicates the allocation status of each inode.
}
    DirTree represents a directory tree structure for the filesystem.
    It contains a pointer to the superblock, a reference to the disk file,
    and bitmaps for blocks and inodes.
```

```go
func NewDirTree(superBlock SuperBlock, file *os.File) *DirTree
    NewDirTree creates a new DirTree instance using the provided superblock and
    file. It reads the block and inode bitmaps from the disk file based on the
    superblock information.
```

```go
func (dt *DirTree) AppendToFileInode(s string, inode *Inode) error
    AppendToFileInode appends the given string content to a file represented
    by the inode. It determines the last file block, appends the new content,
    splits the result into chunks, and writes updated file blocks to disk. Note:
    Indirect addressing is not fully supported.
```

```go
func (dt *DirTree) CreateNewDir(index int32, inode *Inode, dirName string, uid int32, gid int32) (int32, *Inode, error)
    CreateNewDir creates a new directory under a parent directory represented
    by the given inode. It creates a new inode for the directory, initializes
    its folder block with default entries ("." and ".."), writes the new
    folder block and inode to disk, and updates the parent's inode. Returns
    the new directory inode index, a pointer to the new Inode, and an error if
    encountered.
```

```go
func (dt *DirTree) CreateNewFile(index int32, inode *Inode, fileName string, content string, uid int32, gid int32) (int32, *Inode, error)
    CreateNewFile creates a new file under a directory represented by the given
    inode. It creates a new inode for the file, appends the initial content to
    the file's inode, writes the new inode to disk, and returns the new file
    inode index, pointer, and error if any.
```

```go
func (dt *DirTree) CreateNewInode(index int32, inode *Inode, name string, UID, GID int32, inodeType [1]byte) (int32, *Inode, error)
    CreateNewInode creates a new inode entry within a directory represented
    by the parent inode. It adds the new inode entry to a folder block,
    updates bitmaps, writes the new inode to disk, and updates the parent inode
    accordingly. Returns the new inode index, a pointer to the new Inode,
    and an error if encountered.
```

```go
func (dt *DirTree) FreeInodeBlockContent(inode *Inode) error
    FreeInodeBlockContent frees all data blocks associated with an inode
    by updating the block bitmap and resetting the inode's block pointers.
    It writes the updated bitmap to disk.
```

```go
func (dt *DirTree) GetAvailableBlockAddress() (int32, error)
    GetAvailableBlockAddress searches the block bitmap for a free block and
    returns its address.
```

```go
func (dt *DirTree) GetAvailableInodeAdress() (int32, error)
    GetAvailableInodeAdress searches the inode bitmap for a free inode and
    returns its address.
```

```go
func (dt *DirTree) GetFileContentByInode(inode *Inode) (string, error)
    GetFileContentByInode retrieves the content of a file represented by the
    given inode. It verifies the inode type, then reads and concatenates file
    blocks (including indirect blocks) into a single string. The inode's access
    time is updated upon retrieval.
```

```go
func (dt *DirTree) GetInodeByIndex(index int32) (*Inode, error)
    GetInodeByIndex retrieves an inode from the disk given its index. It reads
    the inode from the disk using the superblock's inode start position.
```

```go
func (dt *DirTree) GetInodeByPath(path string) (int32, *Inode, error)
    GetInodeByPath retrieves an inode by traversing the directory tree using
    the provided path. It returns the inode index, a pointer to the Inode,
    and an error if the inode cannot be found.
```

```go
func (dt *DirTree) SplitRawBytes(s string, chunkSize int) []string
    SplitRawBytes splits a string into chunks of the given size and returns a
    slice of string chunks.
```
