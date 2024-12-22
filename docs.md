# Linxr Docs
For a list of all commands (or a specific command) use:
```bash
linxr help <commands>
```

For all listed commands, the command name should be appended right after *linxr* like this:
```bash
linxr <command>
```
---
## Commands

### init
To create a new project, use the *init* command directly in the project directory. Leave template and options emty to create a blank project. A blank project does not create any files. Only links the directory to linxr.
```bash
linxr init <template> <opts>
```

**Options:** 
```-g enable/disable``` Use *-g* to enable or disable automatic git initiazation.
```-l <language>``` Use *-l* to specify the main language for the project. This is used to search and find projects that is using a certain language.
```-d <description>``` Add a description to the project

Ex)
```bash
linxr init my-template -g disable -l python
```
---
### list
See a list of all your Linxr projects. Leave the options blank to view all projects.
```bash
linxr list <opts>
```

**Options:**
```-l <language>``` Filter list to only see projects with the specified language.

Ex)
```bash
linxr list -l c++
```