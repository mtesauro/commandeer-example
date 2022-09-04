package main

import (
	c "github.com/mtesauro/commandeer"
)

// An example of how to create command packages

// GetCmdPackage takes a pointer to a Target and sets the
// package commands for the provided Target
func GetCmdPackage(cPkg *c.Target) {
	// If you wanted, you could have a switch statement
	// here to handle multiple OSes or Linux distros

	// Create an arbitrary command package as an example
	pCmds := []c.SingleCmd{
		{
			// List the files
			Cmd:    "ls",
			Errmsg: "ls failed",
			Hard:   false,
		},
		{
			// Check uptime
			Cmd:    "uptime",
			Errmsg: "uptime failed",
			Hard:   false,
		},
		{
			// Check uptime
			Cmd:    "pstree",
			Errmsg: "pstree failed",
			Hard:   false,
		},
	}

	// Set the Target's package commands
	cPkg.PkgCmds = pCmds

}
