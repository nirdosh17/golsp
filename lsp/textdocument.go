package lsp

type TextDocumentItem struct {
	/**
	 * The text document's URI.
	 */
	URI string `json:"uri"`

	/**
	* The text document's language identifier.
	 */
	LanguageID string `json:"languageId"`

	/**
	* The version number of this document (it will increase after each
	* change, including undo/redo).
	 */
	Version int `json:"version"`

	/**
	* The content of the opened text document.
	 */
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentIdentifier
type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#versionedTextDocumentIdentifier
type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentPositionParams
type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#position
type Position struct {
	// index of line cursor is in. index of first line in the file is 0.
	Line uint `json:"line"`
	// index of character in the line where the cursor is. starts from zero
	Character uint `json:"character"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#location
type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
