# WabiOS Documentation

General documentation for WabiOS. This project requires a lot of planning, so for each folder there's probably more DOCS.md files. Most of it is just documenting my rambling for personal reference and iteration, not for external understanding.

## WIP Implementation Plans:
- Stage 1: Implement and test syscalls from the command prompt over USB serial.
- Stage 2: Create shell language to better interact with OS.
- Stage 3: Create scripting language to handle complex tasks and games.
- Stage 4: Implement OS interaction with scripting language Cortado.
- Stage 5: Test and debug.
- Stage 6: Create applications in Cortado to make OS usable.

### Stretch Goals:
- Create PCB to run WabiOS on.
- Networking.
- Multicore??
- Find way to emulate so it can be tested by people??

## Design:

My experience with Linux is a major basis for how I want to design everything.

### Everything is a file??

Found this on the Plan 9 wikipedia page, and it talked about having virtual file systems for complex things that are usually interacted with through syscalls. This allows for applications to use simple file interactions to talk to complex systems. Since it's virtual, wouldn't pressure the flash storage. Might make implementing file syscalls weird since you have to search both vfs and the actual fs.

### Shells should pass more than strings

Maybe something closer to nushell than bash. Multiple return types. Somehow blend the simplicity and effectiveness of bash with the power of a scripting language. 

### Combine shell and scripting language???

It might be more beneficial to combine the shell and Cortado so I'm not implementing a bunch of the same things twice. Would be an interesting challenge. Turns it more into xonsh.

### Inherent design component: "immutable" kernel

Since I have to compile the kernel on a more capable device, nothing in Go can be written in "userspace". While configuring the kernel before you compile it is possible, you're going to be limited to actions you can perform through the shell, command prompt, and cortado.

## Notes:
- ...how are files preserved if reflashing the kernel wipes the flash...