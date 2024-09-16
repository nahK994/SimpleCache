package test

import "github.com/nahK994/TinyCache/pkg/errors"

var malformedRawCmds []string = []string{
	"SET",
	"SET age ",
	"GET age val",
	"GET",
	"EXISTS",
	"EXISTS age val",
	"TEST",
	"PING haha",
	// Malformed INCR/DECR/DEL commands
	"INCR",         // Missing key argument
	"DECR",         // Missing key argument
	"DEL",          // Missing key argument
	"INCR age val", // Too many arguments for INCR
	"DECR age val", // Too many arguments for DECR
	"DEL key1 key2 key3",
}

var errType = errors.GetErrorTypes()
var testSerializedCmds = []struct {
	name      string
	input     string
	expectErr error
}{
	{
		name:      "Valid Command",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
		expectErr: nil,
	},
	{
		name:      "Empty Command",
		input:     "",
		expectErr: errors.Err{Type: errType.IncompleteCommand},
	},
	{
		name:      "Incorrect Starting Character",
		input:     "3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Malformed Length Specification",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue",
		expectErr: errors.Err{Type: errType.MissingCRLF},
	},
	{
		name:      "Unexpected Characters",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\nextra",
		expectErr: errors.Err{Type: errType.CommandLengthMismatch},
	},
	{
		name:      "Incorrect CRLF Placement",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\n",
		expectErr: errors.Err{Type: errType.MissingCRLF},
	},
	{
		name:      "Command Length Mismatch",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n\r\n",
		expectErr: errors.Err{Type: errType.CommandLengthMismatch},
	},
	{
		name:      "Missing value",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n",
		expectErr: errors.Err{Type: errType.WrongNumberOfArguments},
	},
	{
		name:      "Unexpected character in parsing number",
		input:     "*2$3SET$3foo\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Unexpected character in parsing number 2",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$-1\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Array count mismatch",
		input:     "*2\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
		expectErr: errors.Err{Type: errType.CommandLengthMismatch},
	},
	{
		name:      "Unexpected character in array count",
		input:     "*x\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Missing CRLF after a bulk string",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar",
		expectErr: errors.Err{Type: errType.MissingCRLF},
	},
	{
		name:      "Missing CRLF in position",
		input:     "*2\r\n$3\r\nSET\r\n$5\r\nkeyvalue\r\n",
		expectErr: errors.Err{Type: errType.MissingCRLF},
	},
	{
		name:      "Invalid array format (extra CRLF at the end)",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n\r\n",
		expectErr: errors.Err{Type: errType.CommandLengthMismatch},
	},
	{
		name:      "Missing bulk string length specifier",
		input:     "*3\r\n$3\r\nSET\r\nfoo\r\n$3\r\nbar\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Invalid bulk string format (missing $ sign)",
		input:     "*3\r\n$3\r\nSET\r\n3\r\nfoo\r\n$3\r\nbar\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Incomplete command (missing arguments)",
		input:     "*1\r\n$3\r\nSET\r\n",
		expectErr: errors.Err{Type: errType.WrongNumberOfArguments},
	},
	{
		name:      "Negative bulk string length (Invalid)",
		input:     "-3\r\n$3\r\nSET\r\n$3\r\nage\r\n$3\r\n123\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Invalid bulk string length in the middle",
		input:     "*3\r\n-3\r\nSET\r\n$3\r\nage\r\n$3\r\n123\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Invalid bulk string length with negative number",
		input:     "*3\r\n$3\r\nSET\r\n-3\r\nage\r\n$3\r\n123\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Invalid bulk string length with invalid number in the middle",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nage\r\n-3\r\n123\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Invalid bulk string format with incorrect character",
		input:     "*3\r\n$3\r\nSET\r\n$3\r\nage\r\n$3\r\a123\r\n",
		expectErr: errors.Err{Type: errType.UnexpectedCharacter},
	},
	{
		name:      "Valid command with extra characters",
		input:     "*3\r\n$3\r\nGET\r\n$3\r\nage\r\n$3\r\n123\r\n",
		expectErr: errors.Err{Type: errType.WrongNumberOfArguments},
	},
}
