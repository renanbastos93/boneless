package main

import "github.com/renanbastos93/boneless"

func init() {
	boneless.ValidateLatestVersion()
}

func main() {
	// flag.Usage = func() { fmt.Fprint(os.Stderr, cli.Usage) }
	// flag.Parse()
	// if len(flag.Args()) == 0 {
	// 	fmt.Fprint(os.Stderr, cli.Usage)
	// 	os.Exit(1)
	// }

	// if execute, ok := cli.CmdToRun[flag.Arg(0)]; ok {
	// 	execute()
	// } else {
	// 	flag.Usage()
	// }
}
