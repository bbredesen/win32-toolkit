package win32

import "fmt"

type Win32Error uint32

func (err *Win32Error) Error() string {
	return fmt.Sprintf("0x%.4x: %s", *err, err.String())
}

//go:generate stringer -type Win32Error

const (
	// The operation completed successfully.
	ERROR_SUCCESS Win32Error = 1
	// Incorrect function.
	ERROR_INVALID_FUNCTION Win32Error = 2
	// The system cannot find the file specified.
	ERROR_FILE_NOT_FOUND Win32Error = 3
	// The system cannot find the path specified.
	ERROR_PATH_NOT_FOUND Win32Error = 4
	// The system cannot open the file.
	ERROR_TOO_MANY_OPEN_FILES Win32Error = 5
	// Access is denied.
	ERROR_ACCESS_DENIED Win32Error = 6
	// The handle is invalid.
	ERROR_INVALID_HANDLE Win32Error = 7
	// The storage control blocks were destroyed.
	ERROR_ARENA_TRASHED Win32Error = 8
	// Not enough storage is available to process this command.
	ERROR_NOT_ENOUGH_MEMORY Win32Error = 9
	// The storage control block address is invalid.
	ERROR_INVALID_BLOCK Win32Error = 10
	// The environment is incorrect.
	ERROR_BAD_ENVIRONMENT Win32Error = 11
	// An attempt was made to load a program with an incorrect format.
	ERROR_BAD_FORMAT Win32Error = 12
	// The access code is invalid.
	ERROR_INVALID_ACCESS Win32Error = 13
	// The data is invalid.
	ERROR_INVALID_DATA Win32Error = 14
	// Not enough storage is available to complete this operation.
	ERROR_OUTOFMEMORY Win32Error = 15
	// The system cannot find the drive specified.
	ERROR_INVALID_DRIVE Win32Error = 16
	// The directory cannot be removed.
	ERROR_CURRENT_DIRECTORY Win32Error = 17
	// The system cannot move the file to a different disk drive.
	ERROR_NOT_SAME_DEVICE Win32Error = 18
	// There are no more files.
	ERROR_NO_MORE_FILES Win32Error = 19
	// The media is write protected.
	ERROR_WRITE_PROTECT Win32Error = 20
	// The system cannot find the device specified.
	ERROR_BAD_UNIT Win32Error = 21
	// The device is not ready.
	ERROR_NOT_READY Win32Error = 22
	// The device does not recognize the command.
	ERROR_BAD_COMMAND Win32Error = 23
	// Data error (cyclic redundancy check).
	ERROR_CRC Win32Error = 24
	// The program issued a command but the command length is incorrect.
	ERROR_BAD_LENGTH Win32Error = 25
	// The drive cannot locate a specific area or track on the disk.
	ERROR_SEEK Win32Error = 26
	// The specified disk or diskette cannot be accessed.
	ERROR_NOT_DOS_DISK Win32Error = 27
	// The drive cannot find the sector requested.
	ERROR_SECTOR_NOT_FOUND Win32Error = 28
	// The printer is out of paper.
	ERROR_OUT_OF_PAPER Win32Error = 29
	// The system cannot write to the specified device.
	ERROR_WRITE_FAULT Win32Error = 30
	// The system cannot read from the specified device.
	ERROR_READ_FAULT Win32Error = 31
	// A device attached to the system is not functioning.
	ERROR_GEN_FAILURE Win32Error = 32
	// The process cannot access the file because it is being used by another process.
	ERROR_SHARING_VIOLATION Win32Error = 33
	// The process cannot access the file because another process has locked a portion of the file.
	ERROR_LOCK_VIOLATION Win32Error = 34
	// The wrong diskette is in the drive. Insert %2 (Volume Serial Number: %3) into drive %1.
	ERROR_WRONG_DISK Win32Error = 36
	// Too many files opened for sharing.
	ERROR_SHARING_BUFFER_EXCEEDED Win32Error = 38
	// Reached the end of the file.
	ERROR_HANDLE_EOF Win32Error = 39
	// The disk is full.
	ERROR_HANDLE_DISK_FULL Win32Error = 50
	// The network request is not supported.
	ERROR_NOT_SUPPORTED Win32Error = 51
	// The remote computer is not available.
	ERROR_REM_NOT_LIST Win32Error = 52
	// A duplicate name exists on the network.
	ERROR_DUP_NAME Win32Error = 53
	// The network path was not found.
	ERROR_BAD_NETPATH Win32Error = 54
	// The network is busy.
	ERROR_NETWORK_BUSY Win32Error = 55
	// The specified network resource or device is no longer available.
	ERROR_DEV_NOT_EXIST Win32Error = 56
	// The network BIOS command limit has been reached.
	ERROR_TOO_MANY_CMDS Win32Error = 57
	// A network adapter hardware error occurred.
	ERROR_ADAP_HDW_ERR Win32Error = 58
	// The specified server cannot perform the requested operation.
	ERROR_BAD_NET_RESP Win32Error = 59
	// An unexpected network error occurred.
	ERROR_UNEXP_NET_ERR Win32Error = 60
	// The remote adapter is not compatible.
	ERROR_BAD_REM_ADAP Win32Error = 61
	// The printer queue is full.
	ERROR_PRINTQ_FULL Win32Error = 62
	// Space to store the file waiting to be printed is not available on the server.
	ERROR_NO_SPOOL_SPACE Win32Error = 63
	// Your file waiting to be printed was deleted.
	ERROR_PRINT_CANCELED Win32Error = 64
	// The specified network name is no longer available.
	ERROR_NETNAME_DELETED Win32Error = 65
	// Network access is denied.
	ERROR_NETWORK_ACCESS_DENIED Win32Error = 66
	// The network resource type is not correct.
	ERROR_BAD_DEV_TYPE Win32Error = 67
	// The network name cannot be found.
	ERROR_BAD_NET_NAME Win32Error = 68
	// The name limit for the local computer network adapter card was exceeded.
	ERROR_TOO_MANY_NAMES Win32Error = 69
	// The network BIOS session limit was exceeded.
	ERROR_TOO_MANY_SESS Win32Error = 70
	// The remote server has been paused or is in the process of being started.
	ERROR_SHARING_PAUSED Win32Error = 71
	// No more connections can be made to this remote computer at this time because there are already as many connections as the computer can accept.
	ERROR_REQ_NOT_ACCEP Win32Error = 72
	// The specified printer or disk device has been paused.
	ERROR_REDIR_PAUSED Win32Error = 80
	// The file exists.
	ERROR_FILE_EXISTS Win32Error = 82
	// The directory or file cannot be created.
	ERROR_CANNOT_MAKE Win32Error = 83
	// Fail on INT 24.
	ERROR_FAIL_I24 Win32Error = 84
	// Storage to process this request is not available.
	ERROR_OUT_OF_STRUCTURES Win32Error = 85
	// The local device name is already in use.
	ERROR_ALREADY_ASSIGNED Win32Error = 86
	// The specified network password is not correct.
	ERROR_INVALID_PASSWORD Win32Error = 87
	// The parameter is incorrect.
	ERROR_INVALID_PARAMETER Win32Error = 88
	// A write fault occurred on the network.
	ERROR_NET_WRITE_FAULT Win32Error = 89
	// The system cannot start another process at this time.
	ERROR_NO_PROC_SLOTS Win32Error = 100
	// Cannot create another system semaphore.
	ERROR_TOO_MANY_SEMAPHORES Win32Error = 101
	// The exclusive semaphore is owned by another process.
	ERROR_EXCL_SEM_ALREADY_OWNED Win32Error = 102
	// The semaphore is set and cannot be closed.
	ERROR_SEM_IS_SET Win32Error = 103
	// The semaphore cannot be set again.
	ERROR_TOO_MANY_SEM_REQUESTS Win32Error = 104
	// Cannot request exclusive semaphores at interrupt time.
	ERROR_INVALID_AT_INTERRUPT_TIME Win32Error = 105
	// The previous ownership of this semaphore has ended.
	ERROR_SEM_OWNER_DIED Win32Error = 106
	// Insert the diskette for drive %1.
	ERROR_SEM_USER_LIMIT Win32Error = 107
	// The program stopped because an alternate diskette was not inserted.
	ERROR_DISK_CHANGE Win32Error = 108
	// The disk is in use or locked by another process.
	ERROR_DRIVE_LOCKED Win32Error = 109
	// The pipe has been ended.
	ERROR_BROKEN_PIPE Win32Error = 110
	// The system cannot open the device or file specified.
	ERROR_OPEN_FAILED Win32Error = 111
	// The file name is too long.
	ERROR_BUFFER_OVERFLOW Win32Error = 112
	// There is not enough space on the disk.
	ERROR_DISK_FULL Win32Error = 113
	// No more internal file identifiers available.
	ERROR_NO_MORE_SEARCH_HANDLES Win32Error = 114
	// The target internal file identifier is incorrect.
	ERROR_INVALID_TARGET_HANDLE Win32Error = 117
	// The IOCTL call made by the application program is not correct.
	ERROR_INVALID_CATEGORY Win32Error = 118
	// The verify-on-write switch parameter value is not correct.
	ERROR_INVALID_VERIFY_SWITCH Win32Error = 119
	// The system does not support the command requested.
	ERROR_BAD_DRIVER_LEVEL Win32Error = 120
	// This function is not supported on this system.
	ERROR_CALL_NOT_IMPLEMENTED Win32Error = 121
	// The semaphore timeout period has expired.
	ERROR_SEM_TIMEOUT Win32Error = 122
	// The data area passed to a system call is too small.
	ERROR_INSUFFICIENT_BUFFER Win32Error = 123
	// The filename, directory name, or volume label syntax is incorrect.
	ERROR_INVALID_NAME Win32Error = 124
	// The system call level is not correct.
	ERROR_INVALID_LEVEL Win32Error = 125
	// The disk has no volume label.
	ERROR_NO_VOLUME_LABEL Win32Error = 126
	// The specified module could not be found.
	ERROR_MOD_NOT_FOUND Win32Error = 127
	// The specified procedure could not be found.
	ERROR_PROC_NOT_FOUND Win32Error = 128
	// There are no child processes to wait for.
	ERROR_WAIT_NO_CHILDREN Win32Error = 129
	// The %1 application cannot be run in Win32 mode.
	ERROR_CHILD_NOT_COMPLETE Win32Error = 130
	// Attempt to use a file handle to an open disk partition for an operation other than raw disk I/O.
	ERROR_DIRECT_ACCESS_HANDLE Win32Error = 131
	// An attempt was made to move the file pointer before the beginning of the file.
	ERROR_NEGATIVE_SEEK Win32Error = 132
	// The file pointer cannot be set on the specified device or file.
	ERROR_SEEK_ON_DEVICE Win32Error = 133
	// A JOIN or SUBST command cannot be used for a drive that contains previously joined drives.
	ERROR_IS_JOIN_TARGET Win32Error = 134
	// An attempt was made to use a JOIN or SUBST command on a drive that has already been joined.
	ERROR_IS_JOINED Win32Error = 135
	// An attempt was made to use a JOIN or SUBST command on a drive that has already been substituted.
	ERROR_IS_SUBSTED Win32Error = 136
	// The system tried to delete the JOIN of a drive that is not joined.
	ERROR_NOT_JOINED Win32Error = 137
	// The system tried to delete the substitution of a drive that is not substituted.
	ERROR_NOT_SUBSTED Win32Error = 138
	// The system tried to join a drive to a directory on a joined drive.
	ERROR_JOIN_TO_JOIN Win32Error = 139
	// The system tried to substitute a drive to a directory on a substituted drive.
	ERROR_SUBST_TO_SUBST Win32Error = 140
	// The system tried to join a drive to a directory on a substituted drive.
	ERROR_JOIN_TO_SUBST Win32Error = 141
	// The system tried to SUBST a drive to a directory on a joined drive.
	ERROR_SUBST_TO_JOIN Win32Error = 142
	// The system cannot perform a JOIN or SUBST at this time.
	ERROR_BUSY_DRIVE Win32Error = 143
	// The system cannot join or substitute a drive to or for a directory on the same drive.
	ERROR_SAME_DRIVE Win32Error = 144
	// The directory is not a subdirectory of the root directory.
	ERROR_DIR_NOT_ROOT Win32Error = 145
	// The directory is not empty.
	ERROR_DIR_NOT_EMPTY Win32Error = 146
	// The path specified is being used in a substitute.
	ERROR_IS_SUBST_PATH Win32Error = 147
	// Not enough resources are available to process this command.
	ERROR_IS_JOIN_PATH Win32Error = 148
	// The path specified cannot be used at this time.
	ERROR_PATH_BUSY Win32Error = 149
	// An attempt was made to join or substitute a drive for which a directory on the drive is the target of a previous substitute.
	ERROR_IS_SUBST_TARGET Win32Error = 150
	// System trace information was not specified in your CONFIG.SYS file, or tracing is disallowed.
	ERROR_SYSTEM_TRACE Win32Error = 151
	// The number of specified semaphore events for DosMuxSemWait is not correct.
	ERROR_INVALID_EVENT_COUNT Win32Error = 152
	// DosMuxSemWait did not execute; too many semaphores are already set.
	ERROR_TOO_MANY_MUXWAITERS Win32Error = 153
	// The DosMuxSemWait list is not correct.
	ERROR_INVALID_LIST_FORMAT Win32Error = 154
	// The volume label you entered exceeds the label character limit of the target file system.
	ERROR_LABEL_TOO_LONG Win32Error = 155
	// Cannot create another thread.
	ERROR_TOO_MANY_TCBS Win32Error = 156
	// The recipient process has refused the signal.
	ERROR_SIGNAL_REFUSED Win32Error = 157
	// The segment is already discarded and cannot be locked.
	ERROR_DISCARDED Win32Error = 158
	// The segment is already unlocked.
	ERROR_NOT_LOCKED Win32Error = 159
	// The address for the thread ID is not correct.
	ERROR_BAD_THREADID_ADDR Win32Error = 160
	// The argument string passed to DosExecPgm is not correct.
	ERROR_BAD_ARGUMENTS Win32Error = 161
	// The specified path is invalid.
	ERROR_BAD_PATHNAME Win32Error = 162
	// A signal is already pending.
	ERROR_SIGNAL_PENDING Win32Error = 164
	// No more threads can be created in the system.
	ERROR_MAX_THRDS_REACHED Win32Error = 167
	// Unable to lock a region of a file.
	ERROR_LOCK_FAILED Win32Error = 170
	// The requested resource is in use.
	ERROR_BUSY Win32Error = 173
	// A lock request was not outstanding for the supplied cancel region.
	ERROR_CANCEL_VIOLATION Win32Error = 174
	// The file system does not support atomic changes to the lock type.
	ERROR_ATOMIC_LOCKS_NOT_SUPPORTED Win32Error = 180
	// The system detected a segment number that was not correct.
	ERROR_INVALID_SEGMENT_NUMBER Win32Error = 182
	// The operating system cannot run %1.
	ERROR_INVALID_ORDINAL Win32Error = 183
	// Cannot create a file when that file already exists.
	ERROR_ALREADY_EXISTS Win32Error = 186
	// The flag passed is not correct.
	ERROR_INVALID_FLAG_NUMBER Win32Error = 187
	// The specified system semaphore name was not found.
	ERROR_SEM_NOT_FOUND Win32Error = 188
	// The operating system cannot run %1.
	ERROR_INVALID_STARTING_CODESEG Win32Error = 189
	// The operating system cannot run %1.
	ERROR_INVALID_STACKSEG Win32Error = 190
	// The operating system cannot run %1.
	ERROR_INVALID_MODULETYPE Win32Error = 191
	// Cannot run %1 in Win32 mode.
	ERROR_INVALID_EXE_SIGNATURE Win32Error = 192
	// The operating system cannot run %1.
	ERROR_EXE_MARKED_INVALID Win32Error = 193
	// %1 is not a valid Win32 application.
	ERROR_BAD_EXE_FORMAT Win32Error = 194
	// The operating system cannot run %1.
	ERROR_ITERATED_DATA_EXCEEDS_64k Win32Error = 195
	// The operating system cannot run %1.
	ERROR_INVALID_MINALLOCSIZE Win32Error = 196
	// The operating system cannot run this application program.
	ERROR_DYNLINK_FROM_INVALID_RING Win32Error = 197
	// The operating system is not presently configured to run this application.
	ERROR_IOPL_NOT_ENABLED Win32Error = 198
	// The operating system cannot run %1.
	ERROR_INVALID_SEGDPL Win32Error = 199
	// The operating system cannot run this application program.
	ERROR_AUTODATASEG_EXCEEDS_64k Win32Error = 200
	// The code segment cannot be greater than or equal to 64K.
	ERROR_RING2SEG_MUST_BE_MOVABLE Win32Error = 201
	// The operating system cannot run %1.
	ERROR_RELOC_CHAIN_XEEDS_SEGLIM Win32Error = 202
	// The operating system cannot run %1.
	ERROR_INFLOOP_IN_RELOC_CHAIN Win32Error = 203
	// The system could not find the environment option that was entered.
	ERROR_ENVVAR_NOT_FOUND Win32Error = 205
	// No process in the command subtree has a signal handler.
	ERROR_NO_SIGNAL_SENT Win32Error = 206
	// The filename or extension is too long.
	ERROR_FILENAME_EXCED_RANGE Win32Error = 207
	// The ring 2 stack is in use.
	ERROR_RING2_STACK_IN_USE Win32Error = 208
	// The global filename characters, * or ?, are entered incorrectly or too many global filename characters are specified.
	ERROR_META_EXPANSION_TOO_LONG Win32Error = 209
	// The signal being posted is not correct.
	ERROR_INVALID_SIGNAL_NUMBER Win32Error = 210
	// The signal handler cannot be set.
	ERROR_THREAD_1_INACTIVE Win32Error = 212
	// The segment is locked and cannot be reallocated.
	ERROR_LOCKED Win32Error = 214
	// Too many dynamic-link modules are attached to this program or dynamic-link module.
	ERROR_TOO_MANY_MODULES Win32Error = 215
	// Cannot nest calls to LoadModule.
	ERROR_NESTING_NOT_ALLOWED Win32Error = 216
	// The image file %1 is valid, but is for a machine type other than the current machine.
	ERROR_EXE_MACHINE_TYPE_MISMATCH Win32Error = 230
	// The pipe state is invalid.
	ERROR_BAD_PIPE Win32Error = 231
	// All pipe instances are busy.
	ERROR_PIPE_BUSY Win32Error = 232
	// The pipe is being closed.
	ERROR_NO_DATA Win32Error = 233
	// No process is on the other end of the pipe.
	ERROR_PIPE_NOT_CONNECTED Win32Error = 234
	// More data is available.
	ERROR_MORE_DATA Win32Error = 240
	// The session was canceled.
	ERROR_VC_DISCONNECTED Win32Error = 254
	// The specified extended attribute name was invalid.
	ERROR_INVALID_EA_NAME Win32Error = 255
	// The extended attributes are inconsistent.
	ERROR_EA_LIST_INCONSISTENT Win32Error = 258
	// The wait operation timed out.
	// WAIT_TIMEOUT
	// 259
	// No more data is available.
	ERROR_NO_MORE_ITEMS Win32Error = 266
	// The copy functions cannot be used.
	ERROR_CANNOT_COPY Win32Error = 267
	// The directory name is invalid.
	ERROR_DIRECTORY Win32Error = 275
	// The extended attributes did not fit in the buffer.
	ERROR_EAS_DIDNT_FIT Win32Error = 276
	// The extended attribute file on the mounted file system is corrupt.
	ERROR_EA_FILE_CORRUPT Win32Error = 277
	// The extended attribute table file is full.
	ERROR_EA_TABLE_FULL Win32Error = 278
	// The specified extended attribute handle is invalid.
	ERROR_INVALID_EA_HANDLE Win32Error = 282
	// The mounted file system does not support extended attributes.
	ERROR_EAS_NOT_SUPPORTED Win32Error = 288
	// Attempt to release mutex not owned by caller.
	ERROR_NOT_OWNER Win32Error = 298
	// Too many posts were made to a semaphore.
	ERROR_TOO_MANY_POSTS Win32Error = 299
	// Only part of a ReadProcessMemory or WriteProcessMemory request was completed.
	ERROR_PARTIAL_COPY Win32Error = 300
	// The oplock request is denied.
	ERROR_OPLOCK_NOT_GRANTED Win32Error = 301
	// An invalid oplock acknowledgment was received by the system.
	ERROR_INVALID_OPLOCK_PROTOCOL Win32Error = 317
	// The system cannot find message text for message number 0x%1 in the message file for %2.
	ERROR_MR_MID_NOT_FOUND Win32Error = 487
	// Attempt to access invalid address.
	ERROR_INVALID_ADDRESS Win32Error = 534
	// Arithmetic result exceeded 32 bits.
	ERROR_ARITHMETIC_OVERFLOW Win32Error = 535
	// There is a process on other end of the pipe.
	ERROR_PIPE_CONNECTED Win32Error = 536
	// Waiting for a process to open the other end of the pipe.
	ERROR_PIPE_LISTENING Win32Error = 994
	// Access to the extended attribute was denied.
	ERROR_EA_ACCESS_DENIED Win32Error = 995
	// The I/O operation has been aborted because of either a thread exit or an application request.
	ERROR_OPERATION_ABORTED Win32Error = 996
	// Overlapped I/O event is not in a signaled state.
	ERROR_IO_INCOMPLETE Win32Error = 997
	// Overlapped I/O operation is in progress.
	ERROR_IO_PENDING Win32Error = 998
	// Invalid access to memory location.
	ERROR_NOACCESS Win32Error = 999

// ERROR performing inpage operation.
// ERROR_SWAPERROR
)
