# GIMI - AI-Powered Commit Message Generator

[![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)](https://github.com/shivgitcode/gimi)

GIMI is a command-line interface (CLI) tool that uses AI to automatically generate conventional commit messages from your staged git changes. It's designed to streamline your workflow and help you write clear, concise, and well-formatted commit messages.

## Installation

You can install GIMI using `go install`:

```bash
go install github.com/shivgitcode/gimi@v0.1.0
```

Make sure you have Go installed on your system and that your `$GOPATH/bin` is included in your system's `PATH`.

## Usage

GIMI is simple to use and has two main commands: `init` and `generate`.

### 1. Initialize the Configuration

Before you can start generating commit messages, you need to initialize the configuration. This is where you'll choose your AI backend (Ollama or OpenAI) and provide your API key if necessary.

```bash
gimi init
```

This command will prompt you to select a backend.

- **Ollama**: If you have a local Ollama instance running, you can use it to generate commit messages. This is a free and fast option.
- **OpenAI**: If you prefer to use OpenAI's models, you can select this option. You'll be prompted to enter your OpenAI API key.

The configuration will be saved in `~/.gimi/config.json`.

### 2. Generate a Commit Message

Once you've staged your changes with `git add`, you can generate a commit message:

```bash
gimi generate
```

GIMI will read your staged changes, send them to the configured AI backend, and then print the generated commit message to the console.

## Configuration

GIMI's configuration is stored in `~/.gimi/config.json`. Here's an example of what the file might look like:

```json
{
  "Backend": "ollama"
}
```

Or, if you're using OpenAI:

```json
{
  "Backend": "openai",
  "apiKey": "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

You can manually edit this file if you need to change your backend or update your API key.

## Backends

GIMI supports the following AI backends:

- **Ollama**: For using local language models.
- **OpenAI**: For using OpenAI's powerful language models.

## Development

If you want to contribute to GIMI, you can clone the repository and build the project from source.

```bash
git clone https://github.com/shivgitcode/gimi.git
cd gimi
go build .
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
