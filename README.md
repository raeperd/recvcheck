# recvcheck
Golang linter for check receiver type in method (WIP)

## Motivtation
- [Is there a way to detect following data race code using golangci-lint or other linter?? · golangci/golangci-lint · Discussion #5006](https://github.com/golangci/golangci-lint/discussions/5006)
    - [Wednesday pop quiz: spot the race | Dave Cheney](https://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race)

From [Go Wiki: Go Code Review Comments - The Go Programming Language](https://go.dev/wiki/CodeReviewComments#receiver-type)
> Don’t mix receiver types. Choose either pointers or struct types for all available method


