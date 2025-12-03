function aoc {
    param(
        [Parameter(ValueFromRemainingArguments=$true)]
        [string[]]$Args
    )
    go run .\cmd\main.go @Args
}
