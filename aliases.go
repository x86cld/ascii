package ascii

// Control code aliases in byte form.
const (
	// Soft end of file which "can be used to hide file content when the file is entered at the console or opened in editors".
	// See https://en.wikipedia.org/wiki/Substitute_character#End_of_file for additional information.
	SoftEOF = SUB
)

// Control code aliases in string form.
const (
	StringSoftEOF = StringSUB
)
