package args

type Args struct {
  Source        string
  BindAddress   string
  BindPort      string
  ServerAddress string
  ServerPort    string
  Mode          string
}

func Parse(line []string) Args {
  mappedArgs := mapArgs(line)
  args := Args{
    mappedArgs["--source"],
    mappedArgs["--bind-address"],
    mappedArgs["--bind-port"],
    mappedArgs["--server-address"],
    mappedArgs["--server-port"],
    mappedArgs["--mode"],
  }
  return args
}

func mapArgs(line []string) map[string]string {
  var mapping map[string]string
  mapping = make(map[string]string)
  argsArr := line[1:len(line)]

  for i := 0; i < len(argsArr) - 1; i++ {
    if i % 2 == 0 {
      mapping[argsArr[i]] = argsArr[i + 1]
    }
  }

  return mapping
}


