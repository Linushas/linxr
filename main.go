package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use 'linx help' to see a list of commands")
		return
	}

	// COMMANDS
	if len(os.Args) > 1 && os.Args[1] == "init" {
		initCommand()
	} else if len(os.Args) > 1 && os.Args[1] == "help" {
		helpCommand()
	} else if len(os.Args) > 1 && os.Args[1] == "list" {
		listCommand()
	} else if len(os.Args) > 1 && os.Args[1] == "search" {
		searchCommand()
	} else if len(os.Args) > 1 && os.Args[1] == "update" {
		updateCommand()
	}
	// linxr add (add exesting project to list)
}

func initCommand() {
	if len(os.Args) == 2 {
		// create blank project
		fmt.Printf("Succesfully initialized emty project in current directory.\n")
	} else if len(os.Args) == 3 {
		if os.Args[2] == "blank" {
			// create blank project
			fmt.Printf("Succesfully initialized emty project in current directory.\n")
		} else if os.Args[2] == "SDL2_C" {
			// create SDL2 project
			fmt.Printf("Succesfully initialized project using the SDL2_C template in current directory.\n")
		}

	}
}

func helpCommand() {
	if len(os.Args) == 2 {
		fmt.Printf("Linxr is a CLI tool for effortless project management in your terminal.\n- Usage: 'linxr <command>\n\n")
		fmt.Printf("List of commands:\n\n  \thelp <command>\t\t\tInformation about specific command\n  \tinit <template> <opts>\t\tCreate new project from template\n")
		fmt.Printf("\tlist <opts>\t\t\tList all your Linxr projects\n  \tsearch <string>\t\t\tSearch for Linxr projects\n\tupdate <project-name> <opts>\tEdit the status or description of a project\n\n")
	} else if len(os.Args) == 3 && os.Args[2] == "init" {
		fmt.Printf("The init command is used to create a new project: 'linxr init <template> <opts>\n\n")
		fmt.Printf("Templates:\n\n  \tblank\t\tEmty project (no files is created)\n  \tSDL2_C\t\tC project using the SDL2 library\n")
		fmt.Printf("\nOptions:\n\n  \t-g\t\tEnable/Disable automatic git init. Ex) 'linx init <template> -g disable' to \n\t\t\tcreate project without git initialization.\n\n  \t-l\t\tSpecify the main language for the project.\n\n")
		fmt.Printf("Note: Automatic git initialization is enabled by default. If no template is specified, the project defaults to 'blank'.\n\n")
	}
}

func listCommand() {

}

func searchCommand() {

}

func updateCommand() {

}
