# Markdown Manager (MDM)

This command-line tool helps teams **manage collaborative documents** in a GitHub-first workflow.  
It focuses on clear structure, smooth collaboration, and traceable contributions—while minimizing merge conflicts.

## Features

### Document structure driven by schema.yaml

The document layout is fully defined by a schema file that describes sections and subsections using a recursive data model.

- **Directories for major sections:** Any section containing children is represented as a directory.
- **Files for leaf sections:** Individual leaf sections are stored as standalone files.
- **Aliases:** Optional short identifiers (`alias`) can be provided for convenience, shortening the file name without changing the inner title on the document.

This approach ensures a predictable, file-tree-like structure that is easy to navigate, version, and extend.

Example schema:

```yaml
- title: Summary
- title: Installation and usage
  alias: installation
- title: System Architecture
  alias: architecture
- title: Components
  children:
    - title: Backend
    - title: Frontend
    - title: Database
```

### Automated builds with GitHub Actions

A workflow and a companion script:
- Combine the sections defined in `schema.yaml` into a single output document.
- Preserve a clean contribution trail (commits/PRs) for each collaborator.
- Reduce the likelihood of merge conflicts by encouraging isolated, section-based edits.

### Consistent naming with snake_case 

All section identifiers use snake_case to prevent encoding issues in Git-based version control and to keep diffs readable.

## Installation and Usage

This project is written in **Go**, so you will need to have [Golang](https://go.dev/doc/install) installed on your system.  

### Build from source

1. Clone the repository into your local environment.  
2. Build the binary with:  

```bash
go build -o mdm
```

### Add to PATH

Optionally you could add it to path and use it the same way as other CLI applications such as Git.

## Basic usage

### Initializing repository

```bash
# This requires a schema.yaml file to be present
mdm init
# Optionally you can also generate a sample schema
mdm init --sample
```

This will generate the basic repository structure with:

- **schema.yaml** (schema definition file)
- Section directories and files following snake_case conventions
- A workflow + script for automated builds and CI integration

### Updating sections folder

Whenever the **schema.yaml** file changes, you can generate missing files or folders

```bash
# Adds and deletes sections accordingly
mdm update schema
# Prevents deleting files and directories
mdm update schema --no-delete
```

## Coming soon

The current release ships with the core functionality: schema-driven document structure and GitHub-ready workflows.
Future planned features include:

- **Variable support** across documents for reusable metadata and consistent references.
- **Section-by-section grammar checks** to improve collaboration quality.
- **Custom format transformations** (e.g. .md to PDF or DOCX) to generate polished outputs directly from the CLI.
