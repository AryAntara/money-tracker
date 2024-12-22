package config

type CommandFlag struct {
	Args   []string
	source map[string]string
}

func (c *CommandFlag) ParseArgs() {
	args := c.Args
	flags := map[string]string{}
	for i, arg := range args {
		if arg[0] == '-' {
			nextIndex := i + 1
			if len(args) > nextIndex {
				flags[arg] = args[nextIndex]
			} else {
				flags[arg] = ""
			}
		}
	}

	c.source = flags
}

func (c CommandFlag) ValueOn(index int32) string {
	if len(c.Args) > int(index) {
		return c.Args[index]
	}

	return ""
}

func (c CommandFlag) GetValue(key string) string {
	if c.IsPresent(key) {
		return c.source[key]
	}

	return ""
}

func (c CommandFlag) IsPresent(key string) bool {
	_, ok := c.source[key]
	return ok
}

func NewCommandFlag(args []string) *CommandFlag {
	flag := &CommandFlag{
		Args: args,
	}

	flag.ParseArgs()
	return flag
}
