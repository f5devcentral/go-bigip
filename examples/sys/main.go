package main

import (
	"fmt"
	"os"

	"github.com/f5devcentral/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	// Replace with your actual BIG-IP credentials and hostname
	config := &bigip.Config{
		//Address:           "https://192.168.1.1",
		Username: "admin",
		//Password:          "admin",
		CertVerifyDisable: true, // Disable certificate verification for testing purposes
	}

	config.Address = os.Getenv("BIGIP_ADDRESS")
	if config.Address != "" {
		fmt.Println("BIGIP_ADDRESS:", config.Address)
	} else {
		fmt.Println("BIGIP_ADDRESS is not set.")
	}

	config.Password = os.Getenv("BIGIP_PASSWORD")
	if config.Password != "" {
		fmt.Println("BIGIP_PASSWORD:", config.Password)
	} else {
		fmt.Println("BIGIP_PASSWORD is not set.")
	}

	f5 := bigip.NewSession(config)

	mgmtRouteTesting := false
	mgmtFwRuleTesting := false
	remoteRoleTesting := false
	ldapAuthTesting := false
	authSourceTesting := false
	remoteUserTesting := false
	syslogTesting := false
	globalSettingsTesting := false
	sshdTesting := false
	httpdTesting := false
	snmpConfigTesting := false
	snmpCommunitiesTesting := false
	snmpUsersTesting := true

	if mgmtRouteTesting {
		testManagementRouteCreation(f5)
		testManagementRouteRead(f5)
		testManagementRouteUpdate(f5)
		testManagementRouteRead(f5)
		testManagementRouteDelete(f5)
		testManagementRouteRead(f5)
	}

	if mgmtFwRuleTesting {
		testManagementFwRulesRead(f5)
		testManagementFwRuleCreation(f5)
		testManagementFwRulesRead(f5)
		testManagementFwRuleModify(f5)
		testManagementFwRulesRead(f5)
		testManagementFwRuleDelete(f5)
		testManagementFwRulesRead(f5)
	}
	if remoteRoleTesting {
		testRemoteRoleRead(f5)
		testRemoteRoleCreation(f5)
		testRemoteRoleRead(f5)
		testRemoteRoleUpdate(f5)
		testRemoteRoleRead(f5)
		testRemoteRoleDelete(f5)
		testRemoteRoleRead(f5)
	}
	if ldapAuthTesting {
		testLdapAuthRead(f5)
		testLdapAuthCreation(f5)
		testLdapAuthRead(f5)
		testLdapAuthUpdate(f5)
		testLdapAuthRead(f5)
		testLdapAuthDelete(f5)
		testLdapAuthRead(f5)
	}
	if authSourceTesting {
		testAuthSourceRead(f5)
		testAuthSourceCreate(f5)
		testAuthSourceRead(f5)
		testAuthSourceUpdate(f5)
		testAuthSourceRead(f5)
		testAuthSourceDelete(f5)
		testAuthSourceRead(f5)
	}
	if remoteUserTesting {
		testRemoteUserRead(f5)
		testRemoteUserCreate(f5)
		testRemoteUserRead(f5)
		testRemoteUserUpdate(f5)
		testRemoteUserRead(f5)
		testRemoteUserDelete(f5)
		testRemoteUserRead(f5)
	}
	if syslogTesting {
		testSyslogRead(f5)
		testSyslogCreate(f5)
		testSyslogRead(f5)
		testSyslogUpdate(f5)
		testSyslogRead(f5)
		testSyslogDelete(f5)
		testSyslogRead(f5)
	}
	if globalSettingsTesting {
		testGlobalSettingsRead(f5)
		testGlobalSettingsCreate(f5)
		testGlobalSettingsRead(f5)
		testGlobalSettingsUpdate(f5)
		testGlobalSettingsRead(f5)
		testGlobalSettingsDelete(f5)
		testGlobalSettingsRead(f5)
	}
	if sshdTesting {
		testSSHDRead(f5)
		testSSHDCreate(f5)
		testSSHDRead(f5)
		testSSHDUpdate(f5)
		testSSHDRead(f5)
		testSSHDDelete(f5)
		testSSHDRead(f5)
	}
	if httpdTesting {
		testHTTPDRead(f5)
		testHTTPDCreate(f5)
		testHTTPDRead(f5)
		testHTTPDUpdate(f5)
		testHTTPDRead(f5)
		testHTTPDDelete(f5)
		testHTTPDRead(f5)
	}
	if snmpConfigTesting {
		testSnmpConfigRead(f5)
		testSnmpConfigCreate(f5)
		testSnmpConfigRead(f5)
		testSnmpConfigUpdate(f5)
		testSnmpConfigRead(f5)
		testSnmpConfigDelete(f5)
		testSnmpConfigRead(f5)
	}
	if snmpCommunitiesTesting {
		testSnmpCommunitiesRead(f5)
		testSnmpCommunityCreate(f5)
		testSnmpCommunitiesRead(f5)
		testSnmpCommunityUpdate(f5)
		testSnmpCommunitiesRead(f5)
		testSnmpCommunityDelete(f5)
		testSnmpCommunitiesRead(f5)
	}
	if snmpUsersTesting {
		testSnmpUsersRead(f5)
		testSnmpUserCreation(f5)
		testSnmpUsersRead(f5)
		testSnmpUserUpdate(f5)
		testSnmpUsersRead(f5)
		testSnmpUserDelete(f5)
		testSnmpUsersRead(f5)
	}

}
