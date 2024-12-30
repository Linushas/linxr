package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
	Path         string   `json:"path"`
	Created      string   `json:"created"`
	LastModified string   `json:"last-modified"`
	Tags         []string `json:"tags"`
	Git          string   `json:"git"`
	Repo         string   `json:"repo-url"`
}

type ProjectsWrapper struct {
	Projects []Project `json:"projects"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use 'linx help' to see a list of commands")
		return
	}

	// COMMANDS
	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			initCommand()
		} else if os.Args[1] == "help" {
			helpCommand()
		} else if os.Args[1] == "list" {
			listCommand()
		} else if os.Args[1] == "search" {
			searchCommand()
		} else if os.Args[1] == "update" {
			updateCommand()
		} else if os.Args[1] == "template" {
			templateCommand()
		} else if os.Args[1] == "jump" {
			jumpCommand()
		}
		// linxr add (add exesting project to list)
	}
}

func initCommand() {
	git := true

	if len(os.Args) == 2 {
		fmt.Printf("Succesfully initialized linxr project in current directory.\n")
	} else if len(os.Args) >= 3 {
		if len(os.Args) >= 5 && os.Args[3] == "-g" {
			if os.Args[4] == "disable" {
				git = false
			} else {
				git = true
			}
		}
		if os.Args[2] == "blank" {
			fmt.Printf("Succesfully initialized linxr project in current directory.\n")
		} else {
			fmt.Printf("Trying to copy template to current directory...\n\n")
			templateDir := filepath.Join(getTemplateDir(), os.Args[2])
			err := copyTemplate(templateDir, ".")
			if err != nil {
				fmt.Printf("Error initializing project: %v\n", err)
			} else {
				fmt.Println("Project files copied successfully.")

				if git {
					cmd := exec.Command("git", "init")
					output, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Printf("Error: %s\n", err)
						return
					}
					fmt.Printf("Git: %s\n", string(output))
				}

				var newProject Project
				newProject.Name = os.Args[2]
				newProject.Description = "A new project created from template"
				newProject.Status = "Active"
				newProject.Path, _ = os.Getwd()
				newProject.Created = "2024-12-30"
				newProject.LastModified = "2024-12-30"
				newProject.Tags = []string{}
				if git {
					newProject.Git = "true"
				} else {
					newProject.Git = "false"
				}
				newProject.Repo = "null"

				appendProjects(newProject)

				fmt.Printf("Succesfully initialized linxr project using the SDL2_C template in current directory.\n")
			}
		}
	}
}

func appendProjects(newProject Project) {
	projectsJsonPath := ""
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		projectsJsonPath = filepath.Join(appData, "linxr", "linxr_projects.json")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Could not determine home directory")
		}
		projectsJsonPath = filepath.Join(homeDir, ".linxr", "linxr_projects.json")
	}

	var wrapper ProjectsWrapper

	if _, err := os.Stat(projectsJsonPath); err == nil {
		data, err := ioutil.ReadFile(projectsJsonPath)
		if err == nil {
			json.Unmarshal(data, &wrapper)
		}
	}

	wrapper.Projects = append(wrapper.Projects, newProject)

	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		fmt.Printf("Error saving project: %v\n", err)
		return
	}
	err = ioutil.WriteFile(projectsJsonPath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Println("Project added to linxr projects list.")
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
		fmt.Printf("Linxr is a CLI tool for effortless project management in your terminal.\n- Usage: linxr <command>\n\n")
		fmt.Printf("List of commands:\n\n  \thelp <command>\t\t\t\tInformation about specific command\n  \tinit <template> <opts>\t\t\tCreate new project from template\n")
		fmt.Printf("\tlist <opts>\t\t\t\tList all your Linxr projects\n  \tsearch <string>\t\t\t\tSearch for Linxr projects\n\tupdate <project-name> <opts>\t\tEdit the status or description of a project\n")
		fmt.Printf("\ttemplate <template-name> <action>\tCommand to add or delete a template\n\n")
	} else if len(os.Args) == 3 && os.Args[2] == "init" {
		fmt.Printf("The init command is used to create a new project: 'linxr init <template> <opts>\n\n")
		fmt.Printf("Templates:\n\n  \tblank\t\tEmpty project (no files is created)\n  \tSDL2_C\t\tC project using the SDL2 library\n")
		fmt.Printf("\nOptions:\n\n  \t-g\t\tEnable/Disable automatic git init. Ex) 'linx init <template> -g disable' to \n\t\t\tcreate project without git initialization.\n\n  \t-l\t\tSpecify the main language for the project.\n\n")
		fmt.Printf("Note: Automatic git initialization is enabled by default. If no template is specified, the project defaults to blank (empty).\n\n")
	}
}

func newTemplate(name string, srcDir string) {
	templateDir := filepath.Join(getTemplateDir(), name)
	err := copyTemplate(srcDir, templateDir)
	if err != nil {
		fmt.Printf("Error creating template: %v\n", err)
	} else {
		fmt.Println("Template files copied successfully.")
	}
}

func templateCommand() {
	templateName := ""
	templateDir := "."
	if len(os.Args) == 2 {
		fmt.Printf("Please specify a template. See: 'linxr help template' for more information\n")
	} else if len(os.Args) >= 4 {
		templateName = os.Args[2]
		if os.Args[3] == "delete" {
			templatePath := filepath.Join(getTemplateDir(), templateName)
			err := os.RemoveAll(templatePath)
			if err != nil {
				fmt.Printf("Error deleting template: %v\n", err)
				return
			}
			fmt.Println("Template deleted successfully")
		} else if os.Args[3] == "new" {
			if len(os.Args) >= 5 {
				templateDir = os.Args[4]
			}
			newTemplate(templateName, templateDir)
		}
	} else if len(os.Args) >= 3 && os.Args[2] == "view" {
		allTemplatesDir := getTemplateDir()
		fmt.Printf("Templates:\n\n")
		err := filepath.Walk(allTemplatesDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && path != allTemplatesDir {
				relPath, _ := filepath.Rel(allTemplatesDir, path)
				parts := filepath.SplitList(relPath)

				if len(parts) == 1 {
					fmt.Println("\t", filepath.Base(path))
				}
				return filepath.SkipDir
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking the directory:", err)
		}
		fmt.Printf("\n")
	}

}

func listCommand() {
	projectsJsonPath := ""
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		projectsJsonPath = filepath.Join(appData, "linxr", "linxr_projects.json")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Could not determine home directory")
		}
		projectsJsonPath = filepath.Join(homeDir, ".linxr", "linxr_projects.json")
	}

	src, err := os.Open(projectsJsonPath)
	if err != nil {
		return
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var wrapper ProjectsWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Projects:")
	for _, project := range wrapper.Projects {
		fmt.Printf("\t[%s]   (%s)\n", project.Name, project.Status)
		if project.Description != "null" {
			fmt.Printf("\t- %s\n", project.Description)
		}
		if project.Repo != "null" {
			fmt.Printf("\t- %s\n\n", project.Repo)
		}
	}
}

func jumpCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Project name argument is missing.")
		return
	}
	projectName := os.Args[2]

	projectsJsonPath := ""
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		projectsJsonPath = filepath.Join(appData, "linxr", "linxr_projects.json")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Could not determine home directory")
		}
		projectsJsonPath = filepath.Join(homeDir, ".linxr", "linxr_projects.json")
	}

	src, err := os.Open(projectsJsonPath)
	if err != nil {
		return
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var wrapper ProjectsWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, project := range wrapper.Projects {
		if project.Name == projectName {
			if runtime.GOOS == "windows" {
				cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process powershell -ArgumentList \"cd %s; powershell\"; exit", project.Path))
				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				terminal := "gnome-terminal"
				cmd := exec.Command(terminal, "--working-directory="+project.Path)
				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	}
	fmt.Printf("Error: Project with name '%s' not found.\n", projectName)
}

func searchCommand() {

}

func updateCommand() {

}
