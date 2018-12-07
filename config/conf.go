package config

import (
	"github.com/alecthomas/kingpin"
	"github.com/raedahgroup/webtail/util"
)

var (
	Dir      = kingpin.Arg("dir", "Directory path(s) to look for files").Default("./").ExistingFilesOrDirs()
	Port     = kingpin.Flag("port", "Port number to host the server").Short('p').Default("8080").Int()
	Restrict = kingpin.Flag("restrict", "Enforce PAM authentication (single level)").Short('r').Bool()
	Acl      = kingpin.Flag("acl", "enable Access Control List with users in the provided file").Short('a').ExistingFile()
	Cron     = kingpin.Flag("cron", "configure cron for re-indexing files, Supported durations:[h -> hours, d -> days]").Short('t').Default("0h").String()
	Secure   = kingpin.Flag("secure", "Run Server with TLS").Short('s').Bool()
	Cert     = kingpin.Flag("cert", "Server Certificate").Short('c').Default("server.crt").String()
	Key      = kingpin.Flag("key", "Server Key File").Short('k').Default("server.key").String()
)

func init() {
	kingpin.Parse()
}

func Parse() error {
	return util.ParseConfig(*Dir,*Restrict,*Acl,*Cron)
}