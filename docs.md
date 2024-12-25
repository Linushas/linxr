# Linxr Docs
For a list of all commands (or a specific command) use:
```bash
linxr help <command>
```

For all listed commands, the command name should be appended right after *linxr* like this:
```bash
linxr <command>
```
---
## Commands

### init
To create a new project, use the *init* command directly in the project directory. Leave template and options empty to create a blank project. A blank project does not create any files. Only links the directory to linxr.
```bash
linxr init <template> <opts>
```

**Options:** 
```-g enable/disable``` Use *-g* to enable or disable automatic git initiazation.
```-d <description>``` Add a description to the project
```-o "project name"``` Set the project name, default is the name of the project directory.

Ex)
```bash
linxr init my-template -g disable -o "My Project"
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

---
### template
Create, delete or view your templates
```bash
linxr template <template-name> <action>
```
or to view all your saved templates
```bash
linxr template view
```

**Actions:**
```new``` Create new template from current directory.
```delete``` Delete specified template.

Ex)
```bash
linxr template my_template new
```

---
### jump
Open one of your linxr project directories in a terminal window.
```bash
linxr jump "project name"
```

Ex)
```bash
linxr jump "My Fantastic Project"
```

