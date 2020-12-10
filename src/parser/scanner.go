package parser

//Scan will parse the input into a primary command and its arguments
func Scan(input string) (string, []string) {
	args := make([]string, 1)
	for _, c := range input {
		if c == ' ' || c == '\r' {
			//crate new argument
			if args[len(args)-1] != "" {
				args = append(args, "")
			}
		} else {
			args[len(args)-1] += string(c)
		}
	}

	return args[0], args[1:]
}
