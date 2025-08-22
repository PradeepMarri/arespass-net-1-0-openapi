package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// About represents the About schema from the OpenAPI specification
type About struct {
	Availablelanguagesiso639_1 string `json:"availableLanguagesIso639_1,omitempty"` // **The list of available languages.** Each language is identified by its ISO 639-1, two-letter code. The list elements are comma-separated and sorted in ascending order.
	Apireleasedateiso8601 string `json:"apiReleaseDateIso8601,omitempty"` // **The release date of this API, ISO 8601 format.**
	Apiversion map[string]interface{} `json:"apiVersion,omitempty"`
}

// Ec represents the Ec schema from the OpenAPI specification
type Ec struct {
	Entropy float64 `json:"entropy,omitempty"` // **The entropy calculated for the input password.** It is measured in bits.
	Summary []string `json:"summary,omitempty"`
	Total []map[string]interface{} `json:"total,omitempty"` // **The total penalty applied to each character.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Efficiency float64 `json:"efficiency,omitempty"` // **The ratio entropy / idealEntropy.** It is a float number in the range [0, 1].
	Penalty float64 `json:"penalty,omitempty"` // **The penalty applied to each character that has been found to be part of a word, number sequence, alphabet sequence, etc.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1. Its value is equal to the value of the input parameter *penalty*.
	Apiversion string `json:"apiVersion,omitempty"` // **This API version number.**
	Nonuniformentropydistributionpenalty float64 `json:"nonUniformEntropyDistributionPenalty,omitempty"` // **The penalty applied to the whole password because of irregular entropy distribution.** This penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Requestid string `json:"requestId,omitempty"` // **The identifier of the request that corresponds to this response.**
	Alphabetsequence []map[string]interface{} `json:"alphabetSequence,omitempty"` // **The penalty applied to each character that has been found to be part of an alphabet sequence.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Numbersequence []map[string]interface{} `json:"numberSequence,omitempty"` // **The penalty applied to each character that has been found to be part of a number sequence.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Requesttimestamp float64 `json:"requestTimestamp,omitempty"` // **The timestamp for this response.** Milliseconds from the epoch of 1970-01-01T00:00:00Z.
	Idealentropy float64 `json:"idealEntropy,omitempty"` // **The Shannon entropy.** The Shannon entropy is the entropy calculated if no penalizations - words, number sequence, alphabet sequence, etc - were found in the password. It is measured in bits.
	Detectedkeyboard string `json:"detectedKeyboard,omitempty"` // **The detected keyboard, QWERTY or Dvorak.**
	L33tpassword string `json:"l33tPassword,omitempty"` // The analyzed password after the l33t substitution.
	Password string `json:"password,omitempty"` // The analyzed password.
	Repeatedchars []map[string]interface{} `json:"repeatedChars,omitempty"` // **The penalty applied to each character that are repeated** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Keyboardsequence []map[string]interface{} `json:"keyboardSequence,omitempty"` // **The penalty applied to each character that has been found to be part of a keyboard sequence.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Words []map[string]interface{} `json:"words,omitempty"` // **The penalty applied to each character that has been found to be part of a word.** The penalty is a float number in the range [0, 1]. Full penalty, 0; no penalty, 1.
	Entropydistribution []map[string]interface{} `json:"entropyDistribution,omitempty"` // **The distribution of the calculated entropy among the password characters.**
	Passwordlength int `json:"passwordLength,omitempty"` // The number of characters the password has.
}
