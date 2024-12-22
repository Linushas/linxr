package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// type Project interface {
// }

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
		fmt.Printf("Succesfully initialized linxr project in current directory.\n")
	} else if len(os.Args) == 3 {
		if os.Args[2] == "blank" {
			fmt.Printf("Succesfully initialized linxr project in current directory.\n")
		} else {
			fmt.Printf("Trying to copy template to current directory...\n")
			templateDir := filepath.Join(getTemplateDir(), os.Args[2])
			err := copyTemplate(templateDir, ".")
			if err != nil {
				fmt.Printf("Error initializing project: %v\n", err)
			} else {
				fmt.Println("Project files copied successfully.")
				fmt.Printf("Succesfully initialized linxr project using the SDL2_C template in current directory.\n")
			}
		}
	}
}

func getTemplateDir() string {
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		return filepath.Join(appData, "linxr", "templates")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Could not determine home directory")
		}
		return filepath.Join(homeDir, ".linxr", "templates")
	}
}

func copyTemplate(template, dest string) error {
	return filepath.Walk(template, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(template, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dest, relPath)
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		return copyFile(path, targetPath)
	})
}

func copyFile(templateFile, destFile string) error {
	src, err := os.Open(templateFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	templateInfo, err := os.Stat(templateFile)
	if err != nil {
		return err
	}
	return os.Chmod(destFile, templateInfo.Mode())
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
