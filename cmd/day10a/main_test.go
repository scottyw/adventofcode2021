package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {

	assert.Equal(t, 26397, verify(strings.Split(input, "\n")))

}

// Some of the lines aren't corrupted, just incomplete; you can ignore these lines for now. The remaining five lines are corrupted:

//     {([(<{}[<>[]}>{[]{[(<()> - Expected ], but found } instead.
//     [[<[([]))<([[{}[[()]]] - Expected ], but found ) instead.
//     [{[{({}]{}}([{[{{{}}([] - Expected ), but found ] instead.
//     [<(<(<(<{}))><([]([]() - Expected >, but found ) instead.
//     <{([([[(<>()){}]>(<<{{ - Expected ], but found > instead.

var input = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`
