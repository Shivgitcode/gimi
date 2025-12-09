package helpers

import "fmt"


func GitDiffPrompt(diff string) string{
    return fmt.Sprintf(`
	you are an ai expert conventional commit generator 

	FORMAT:
	git commit -m "<type>(scope): <message>"


	RULES:
	- one line full git command with -m flag
	- No explanation.
	- No markdown.
	- No code block.
	- No extra whitespace.
	- No summary.
	- Only the commit command.
	- <type>: feat|fix|refactor|chore|docs|style|perf
	- Keep <message> <75 chars.

	Diff:
	%s

`, diff)
}

