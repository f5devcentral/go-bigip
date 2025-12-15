package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////            Syslog Configuration Testing      /////////////////////
/////////////////////////////////////////////////////////////////////////////////

func testSyslogRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Syslog Configuration Read Operation ===")

	fmt.Println("Getting Syslog configuration...")
	syslogConfig, err := f5.GetSyslogConfig()
	if err != nil {
		log.Printf("Error getting Syslog configuration: %v", err)
		return
	}
	if syslogConfig != nil {
		fmt.Printf("Retrieved Syslog Configuration:\n")
		fmt.Printf("  Include: %s\n", syslogConfig.Include)
		fmt.Printf("  Remote Servers: %d\n", len(syslogConfig.RemoteServers))
		for i, server := range syslogConfig.RemoteServers {
			fmt.Printf("    Server %d: %s:%d (LocalIP: %s)\n", i+1, server.Host, server.RemotePort, server.LocalIp)
		}
	}
}

func testSyslogCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Syslog Configuration Create Operation ===")

	syslogConfig := &bigip.SyslogConfig{
		RemoteServers: []bigip.SyslogRemoteServer{
			{
				Name:        "/Common/remote-syslog-01",
				Host:        "10.1.10.251",
				LocalIp:     "none",
				RemotePort:  1514,
				Description: "Test remote syslog server",
			},
		},
	}

	fmt.Println("Setting Syslog configuration...")
	err := f5.CreateSyslogConfig(syslogConfig)
	if err != nil {
		log.Printf("Error creating Syslog configuration: %v", err)
		return
	}
	fmt.Printf("Syslog configuration set successfully:\n")
	fmt.Printf("  Include: %s\n Remote Servers: %d\n",
		syslogConfig.Include, len(syslogConfig.RemoteServers))
}

func testSyslogUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Syslog Configuration Update Operation ===")

	syslogConfig := &bigip.SyslogConfig{
		Include: `
filter f_complete {
  facility(auth,authpriv);
};
destination d_syslog_server {
  udp(\"10.1.10.251\" port (1514));
};
log {
  source(s_syslog_pipe);
  filter(f_complete);
  destination(d_syslog_server);
};
`,
		RemoteServers: []bigip.SyslogRemoteServer{},
	}

	fmt.Println("Updating Syslog configuration...")
	err := f5.ModifySyslogConfig(syslogConfig)
	if err != nil {
		log.Printf("Error updating Syslog configuration: %v", err)
		return
	}
	fmt.Printf("Syslog configuration updated successfully:\n")
	fmt.Printf("  Include: %s\n Remote Servers: %d\n",
		syslogConfig.Include, len(syslogConfig.RemoteServers))
}

func testSyslogDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Syslog Configuration Delete Operation ===")

	fmt.Println("Resetting Syslog configuration to defaults...")
	err := f5.DeleteSyslogConfig()
	if err != nil {
		log.Printf("Error deleting Syslog configuration: %v", err)
		return
	}
	fmt.Printf("Syslog configuration reset to defaults.\n")
}
