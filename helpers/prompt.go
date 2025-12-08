package helpers

import "fmt"


func GitDiffPrompt(diff string) string{
	p:=fmt.Sprintf(`
					Assistant, generate a git commit message in Conventional Commits format 
					Rules:
					- Keep subject under 75 chars
					- No explaining the diff
					- don't summarise the diff 
					- just single line commit message
					- Only output the commit message
					based on this diff:

					%s
					
					`, diff)

	return p
}