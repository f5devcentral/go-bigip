/*
Original work Copyright © 2015 Scott Ware
Modifications Copyright 2019 F5 Networks Inc
Licensed under the Apache License, Version 2.0 (the "License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.
*/
package bigip

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	//"strings"
	"time"
)

type Version struct {
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
	Entries  struct {
		HTTPSLocalhostMgmtTmCliVersion0 struct {
			NestedStats struct {
				Entries struct {
					Active struct {
						Description string `json:"description"`
					} `json:"active,omitempty"`
					Latest struct {
						Description string `json:"description"`
					} `json:"latest,omitempty"`
					Supported struct {
						Description string `json:"description"`
					} `json:"supported,omitempty"`
				} `json:"entries,omitempty"`
			} `json:"nestedStats,omitempty"`
		} `json:"https://localhost/mgmt/tm/cli/version/0,omitempty"`
	} `json:"entries,omitempty"`
}

type NTPs struct {
	NTPs []NTP `json:"items"`
}

type NTP struct {
	Description string   `json:"description"`
	Servers     []string `json:"servers"`
	Timezone    string   `json:"timezone,omitempty"`
}

type BigipCommand struct {
	Command       string `json:"command"`
	UtilCmdArgs   string `json:"utilCmdArgs"`
	CommandResult string `json:"commandResult,omitempty"`
}

type BigipCmdResp struct {
	Code       int           `json:"code"`
	Message    string        `json:"message"`
	ErrorStack []interface{} `json:"errorStack"`
	APIError   int           `json:"apiError"`
}

type DNSs struct {
	DNSs []DNS `json:"items"`
}

type DNS struct {
	Description  string   `json:"description"`
	NameServers  []string `json:"nameServers,"`
	NumberOfDots int      `json:"numberOfDots,omitempty"`
	Search       []string `json:"search"`
}

type Provisions struct {
	Provisions []Provision `json:"items"`
}

type Provision struct {
	Name        string `json:"name,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	CpuRatio    int    `json:"cpuRatio,omitempty"`
	DiskRatio   int    `json:"diskRatio,omitempty"`
	Level       string `json:"level,omitempty"`
	MemoryRatio int    `json:"memoryRatio,omitempty"`
}

type SNMPs struct {
	SNMPs []SNMP `json:"items"`
}

type SNMP struct {
	SysContact       string   `json:"sysContact,omitempty"`
	SysLocation      string   `json:"sysLocation,omitempty"`
	AllowedAddresses []string `json:"allowedAddresses,omitempty"`
}

type TRAPs struct {
	SNMPs []SNMP `json:"items"`
}

type TRAP struct {
	Name                     string `json:"name,omitempty"`
	AuthPasswordEncrypted    string `json:"authPasswordEncrypted,omitempty"`
	AuthProtocol             string `json:"authProtocol,omitempty"`
	Community                string `json:"community,omitempty"`
	Description              string `json:"description,omitempty"`
	EngineId                 string `json:"engineId,omitempty"`
	Host                     string `json:"host,omitempty"`
	Port                     int    `json:"port,omitempty"`
	PrivacyPassword          string `json:"privacyPassword,omitempty"`
	PrivacyPasswordEncrypted string `json:"privacyPasswordEncrypted,omitempty"`
	PrivacyProtocol          string `json:"privacyProtocol,omitempty"`
	SecurityLevel            string `json:"securityLevel,omitempty"`
	SecurityName             string `json:"SecurityName,omitempty"`
	Version                  string `json:"version,omitempty"`
}

type Bigiplicenses struct {
	Bigiplicenses []Bigiplicense `json:"items"`
}

type Bigiplicense struct {
	Registration_key string `json:"registrationKey,omitempty"`
	Command          string `json:"command,omitempty"`
}

type LogIPFIXs struct {
	LogIPFIXs []LogIPFIX `json:"items"`
}
type LogIPFIX struct {
	AppService                 string `json:"appService,omitempty"`
	Name                       string `json:"name,omitempty"`
	PoolName                   string `json:"poolName,omitempty"`
	ProtocolVersion            string `json:"protocolVersion,omitempty"`
	ServersslProfile           string `json:"serversslProfile,omitempty"`
	TemplateDeleteDelay        int    `json:"templateDeleteDelay,omitempty"`
	TemplateRetransmitInterval int    `json:"templateRetransmitInterval,omitempty"`
	TransportProfile           string `json:"transportProfile,omitempty"`
}
type LogPublishers struct {
	LogPublishers []LogPublisher `json:"items"`
}
type LogPublisher struct {
	Name  string `json:"name,omitempty"`
	Dests []Destinations
}

type Destinations struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
}

type destinationsDTO struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	Dests     struct {
		Items []Destinations `json:"items,omitempty"`
	} `json:"destinationsReference,omitempty"`
}

type ExternalDGFile struct {
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	SourcePath string `json:"sourcePath"`
	Type       string `json:"type"`
}

type OCSP struct {
	Name                       string `json:"name,omitempty"`
	FullPath                   string `json:"fullPath,omitempty"`
	Partition                  string `json:"partition,omitempty"`
	ProxyServerPool            string `json:"proxyServerPool,omitempty"`
	DnsResolver                string `json:"dnsResolver,omitempty"`
	RouteDomain                string `json:"routeDomain,omitempty"`
	ConcurrentConnectionsLimit int64  `json:"concurrentConnectionsLimit,omitempty"`
	ResponderUrl               string `json:"responderUrl,omitempty"`
	ConnectionTimeout          int64  `json:"timeout,omitempty"`
	TrustedResponders          string `json:"trustedResponders,omitempty"`
	ClockSkew                  int64  `json:"clockSkew,omitempty"`
	StatusAge                  int64  `json:"statusAge,omitempty"`
	StrictRespCertCheck        string `json:"strictRespCertCheck,omitempty"`
	CacheTimeout               string `json:"cacheTimeout,omitempty"`
	CacheErrorTimeout          int64  `json:"cacheErrorTimeout,omitempty"`
	SignerCert                 string `json:"signerCert,omitempty"`
	SignerKey                  string `json:"signerKey,omitempty"`
	Passphrase                 string `json:"passphrase,omitempty"`
	SignHash                   string `json:"signHash,omitempty"`
}

func (p *LogPublisher) MarshalJSON() ([]byte, error) {
	return json.Marshal(destinationsDTO{
		Name: p.Name,
		Dests: struct {
			Items []Destinations `json:"items,omitempty"`
		}{Items: p.Dests},
	})
}

func (p *LogPublisher) UnmarshalJSON(b []byte) error {
	var dto destinationsDTO
	err := json.Unmarshal(b, &dto)
	if err != nil {
		return err
	}

	p.Name = dto.Name
	p.Dests = dto.Dests.Items
	return nil
}

const (
	uriMgmtRoute       = "management-route"
	uriMgmtIpRules     = "management-ip-rules"
	uriRules           = "rules"
	uriLdap            = "ldap"
	uriAuthSrc         = "source"
	uriRemoteUser      = "remote-user"
	uriGlobalSettings  = "global-settings"
	uriSshd            = "sshd"
	uriHttpd           = "httpd"
	uriSys             = "sys"
	uriTm              = "tm"
	uriCli             = "cli"
	uriUtil            = "util"
	uriBash            = "bash"
	uriVersion         = "version"
	uriNtp             = "ntp"
	uriDNS             = "dns"
	uriProvision       = "provision"
	uriAfm             = "afm"
	uriAsm             = "asm"
	uriApm             = "apm"
	uriAvr             = "avr"
	uriAuth            = "auth"
	uriPartition       = "partition"
	uriRemoteRole      = "remote-role"
	uriRoleInfo        = "role-info"
	uriFolder          = "folder"
	uriIlx             = "ilx"
	uriSyslog          = "syslog"
	uriSnmp            = "snmp"
	uriTraps           = "traps"
	uriLicense         = "license"
	uriLogConfig       = "logConfig"
	uriDestination     = "destination"
	uriIPFIX           = "ipfix"
	uriPublisher       = "publisher"
	uriFile            = "file"
	uriSslCert         = "ssl-cert"
	uriSslKey          = "ssl-key"
	uriDataGroup       = "data-group"
	uriTransaction     = "transaction"
	REST_DOWNLOAD_PATH = "/var/config/rest/downloads"
)

// Certificates represents a list of installed SSL certificates.
type Certificates struct {
	Certificates []Certificate `json:"items,omitempty"`
}

// Certificate represents an SSL Certificate.
type Certificate struct {
	AppService              string                  `json:"appService,omitempty"`
	CachePath               string                  `json:"cachePath,omitempty"`
	CertificateKeyCurveName string                  `json:"certificateKeyCurveName,omitempty"`
	CertificateKeySize      int                     `json:"certificateKeySize,omitempty"`
	CertValidationOptions   []string                `json:"certValidationOptions,omitempty"`
	Checksum                string                  `json:"checksum,omitempty"`
	CreatedBy               string                  `json:"createdBy,omitempty"`
	CreateTime              string                  `json:"createTime,omitempty"`
	Email                   string                  `json:"email,omitempty"`
	ExpirationDate          int64                   `json:"expirationDate,omitempty"`
	ExpirationString        string                  `json:"expirationString,omitempty"`
	Fingerprint             string                  `json:"fingerprint,omitempty"`
	FullPath                string                  `json:"fullPath,omitempty"`
	Generation              int                     `json:"generation,omitempty"`
	IsBundle                string                  `json:"isBundle,omitempty"`
	IsDynamic               string                  `json:"isDynamic,omitempty"`
	Issuer                  string                  `json:"issuer,omitempty"`
	IssuerCert              string                  `json:"issuerCert,omitempty"`
	KeyType                 string                  `json:"keyType,omitempty"`
	LastUpdateTime          string                  `json:"lastUpdateTime,omitempty"`
	Mode                    int                     `json:"mode,omitempty"`
	Name                    string                  `json:"name,omitempty"`
	Partition               string                  `json:"partition,omitempty"`
	Revision                int                     `json:"revision,omitempty"`
	SerialNumber            string                  `json:"serialNumber,omitempty"`
	Size                    uint64                  `json:"size,omitempty"`
	SourcePath              string                  `json:"sourcePath,omitempty"`
	Subject                 string                  `json:"subject,omitempty"`
	SubjectAlternativeName  string                  `json:"subjectAlternativeName,omitempty"`
	SystemPath              string                  `json:"systemPath,omitempty"`
	UpdatedBy               string                  `json:"updatedBy,omitempty"`
	Version                 int                     `json:"version,omitempty"`
	CertValidatorRef        *CertValidatorReference `json:"certValidatorsReference,omitempty"`
}

type CertValidatorReference struct {
	Items []CertValidatorState `json:"items,omitempty"`
}

type CertValidatorState struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	FullPath  string `json:"fullPath,omitempty"`
}

// Keys represents a list of installed keys.
type Keys struct {
	Keys []Key `json:"items,omitempty"`
}

// Key represents a private key associated with a certificate.
type Key struct {
	AppService     string `json:"appService,omitempty"`
	CachePath      string `json:"cachePath,omitempty"`
	Checksum       string `json:"checksum,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CurveName      string `json:"curveName,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	IsDynamic      string `json:"isDynamic,omitempty"`
	KeySize        int    `json:"keySize,omitempty"`
	KeyType        string `json:"keyType,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	Mode           int    `json:"mode,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	Passphrase     string `json:"passphrase,omitempty"`
	Revision       int    `json:"revision,omitempty"`
	SecurityType   string `json:"securityType,omitempty"`
	Size           uint64 `json:"size,omitempty"`
	SourcePath     string `json:"sourcePath,omitempty"`
	SystemPath     string `json:"systemPath,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
}

type Transaction struct {
	TransID          int64  `json:"transId,omitempty"`
	State            string `json:"state,omitempty"`
	TimeoutSeconds   int64  `json:"timeoutSeconds,omitempty"`
	AsyncExecution   bool   `json:"asyncExecution,omitempty"`
	ValidateOnly     bool   `json:"validateOnly,omitempty"`
	ExecutionTimeout int64  `json:"executionTimeout,omitempty"`
	ExecutionTime    int64  `json:"executionTime,omitempty"`
	FailureReason    string `json:"failureReason,omitempty"`
}

type Partition struct {
	Name        string `json:"name,omitempty"`
	RouteDomain int    `json:"defaultRouteDomain"`
	Description string `json:"description,omitempty"`
}

type RoleInfo struct {
	Name          string `json:"name,omitempty"`
	Attribute     string `json:"attribute"`
	Console       string `json:"console,omitempty"`
	Deny          string `json:"deny,omitempty"`
	Description   string `json:"description,omitempty"`
	LineOrder     int    `json:"lineOrder"`
	Role          string `json:"role,omitempty"`
	UserPartition string `json:"userPartition,omitempty"`
}

// Certificates returns a list of certificates.
func (b *BigIP) Certificates() (*Certificates, error) {
	var certs Certificates
	err, _ := b.getForEntity(&certs, uriSys, uriFile, uriSslCert)
	if err != nil {
		return nil, err
	}

	return &certs, nil
}

// AddCertificate installs a certificate.
func (b *BigIP) AddCertificate(cert *Certificate) error {
	return b.post(cert, uriSys, uriFile, uriSslCert)
}

// AddExternalDatagroupfile adds datagroup file
func (b *BigIP) AddExternalDatagroupfile(dataGroup *ExternalDGFile) error {
	return b.post(dataGroup, uriSys, uriFile, uriDataGroup)
}

// DeleteExternalDatagroupfile removes a Datagroup file.
func (b *BigIP) DeleteExternalDatagroupfile(name string) error {
	return b.delete(uriSys, uriFile, uriDataGroup, name)
}

// ModifyExternalDatagroupfile modify datagroup file
func (b *BigIP) ModifyExternalDatagroupfile(dgName string, dataGroup *ExternalDGFile) error {
	return b.patch(dataGroup, uriSys, uriFile, uriDataGroup, dgName)
}

// ModifyCertificate installs a certificate.
func (b *BigIP) ModifyCertificate(certName string, cert *Certificate) error {
	return b.patch(cert, uriSys, uriFile, uriSslCert, certName, "?expandSubcollections=true")
}

// UploadCertificate copies a certificate local disk to BIGIP
func (b *BigIP) UploadCertificate(certpath string, cert *Certificate) error {
	certbyte := []byte(certpath)
	_, err := b.UploadBytes(certbyte, cert.Name)
	if err != nil {
		return err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + cert.Name
	log.Printf("[DEBUG] sourcepath :%+v", sourcepath)

	cert.SourcePath = sourcepath
	log.Printf("cert: %+v\n", cert)
	err = b.AddCertificate(cert)
	if err != nil {
		return err
	}
	return nil
}

// GetCertificate retrieves a Certificate by name. Returns nil if the certificate does not exist
func (b *BigIP) GetCertificate(name string) (*Certificate, error) {
	var cert Certificate
	err, ok := b.getForEntity(&cert, uriSys, uriFile, uriSslCert, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &cert, nil
}

// DeleteCertificate removes a certificate.
func (b *BigIP) DeleteCertificate(name string) error {
	return b.delete(uriSys, uriFile, uriSslCert, name)
}

// UpdateCertificate copies a certificate local disk to BIGIP
func (b *BigIP) UpdateCertificate(certpath string, cert *Certificate) error {
	certbyte := []byte(certpath)
	_, err := b.UploadBytes(certbyte, cert.Name)
	if err != nil {
		return err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + cert.Name

	cert.SourcePath = sourcepath
	certName := fmt.Sprintf("/%s/%s", cert.Partition, cert.Name)
	log.Printf("certName: %+v\n", certName)
	err = b.ModifyCertificate(certName, cert)
	if err != nil {
		return err
	}
	return nil
}

// UploadKey copies a certificate key from local disk to BIGIP
func (b *BigIP) UploadKey(keyname, keypath string) (string, error) {
	keybyte := []byte(keypath)
	_, err := b.UploadBytes(keybyte, keyname)
	if err != nil {
		return "", err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + keyname
	log.Println("[DEBUG] string:", sourcepath)
	return sourcepath, nil
}

// UpdateKey copies a certificate key from local disk to BIGIP
func (b *BigIP) UpdateKey(keyname, keypath, partition string) error {
	keybyte := []byte(keypath)
	_, err := b.UploadBytes(keybyte, keyname)
	if err != nil {
		return err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + keyname
	log.Println("[DEBUG]string:", sourcepath)
	certkey := Key{
		Name:       keyname,
		SourcePath: sourcepath,
		Partition:  partition,
	}
	keyName := fmt.Sprintf("/%s/%s", partition, keyname)
	log.Printf("[DEBUG]keyName: %+v\n", keyName)
	err = b.ModifyKey(keyName, &certkey)
	if err != nil {
		return err
	}
	return nil
}

// Keys returns a list of keys.
func (b *BigIP) Keys() (*Keys, error) {
	var keys Keys
	err, _ := b.getForEntity(&keys, uriSys, uriFile, uriSslKey)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

// AddKey installs a key.
func (b *BigIP) AddKey(config *Key) error {
	return b.post(config, uriSys, uriFile, uriSslKey)
}

// ModifyKey Updates a key.
func (b *BigIP) ModifyKey(keyName string, config *Key) error {
	return b.patch(config, uriSys, uriFile, uriSslKey, keyName)
}

// GetKey retrieves a key by name. Returns nil if the key does not exist.
func (b *BigIP) GetKey(name string) (*Key, error) {
	var key Key
	err, ok := b.getForEntity(&key, uriSys, uriFile, uriSslKey, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &key, nil
}

// DeleteKey removes a key.
func (b *BigIP) DeleteKey(name string) error {
	return b.delete(uriSys, uriFile, uriSslKey, name)
}

type IFile struct {
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	SubPath        string `json:"subPath,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Checksum       string `json:"checksum,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	Mode           int    `json:"mode,omitempty"`
	Revision       int    `json:"revision,omitempty"`
	Size           int    `json:"size,omitempty"`
	SourcePath     string `json:"sourcePath,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
}

func (b *BigIP) ImportIfile(ifile *IFile, fileData, opCall string) error {
	fileByte := []byte(fileData)
	_, err := b.UploadBytes(fileByte, ifile.Name)
	if err != nil {
		return err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + ifile.Name
	log.Println("[DEBUG]string:", sourcepath)
	ifile.SourcePath = sourcepath
	// fileName := fmt.Sprintf("/%s/%s", ifile.Partition, ifile.Name)
	// log.Printf("[DEBUG]fileName: %+v\n", fileName)
	if opCall == "POST" {
		err = b.CreateIFile(ifile)
		if err != nil {
			return err
		}
	}
	if opCall == "PUT" {
		err = b.UpdateIFile(ifile.FullPath, ifile)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create iFile
func (b *BigIP) CreateIFile(ifile *IFile) error {
	return b.post(ifile, uriSys, uriFile, "ifile")
}

// Get iFile
func (b *BigIP) GetIFile(name string) (*IFile, error) {
	var ifile IFile
	err, ok := b.getForEntity(&ifile, uriSys, uriFile, "ifile", name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &ifile, nil
}

// Update iFile
func (b *BigIP) UpdateIFile(name string, ifile *IFile) error {
	return b.put(ifile, uriSys, uriFile, "ifile", name)
}

// Delete iFile
func (b *BigIP) DeleteIFile(name string) error {
	return b.delete(uriSys, uriFile, "ifile", name)
}

// Add to the existing sys.go file after the IFile struct definition
type LtmIFile struct {
	Name              string             `json:"name,omitempty"`
	Partition         string             `json:"partition,omitempty"`
	SubPath           string             `json:"subPath,omitempty"`
	FullPath          string             `json:"fullPath,omitempty"`
	FileName          string             `json:"fileName,omitempty"`
	FileNameReference *FileNameReference `json:"fileNameReference,omitempty"`
}

type FileNameReference struct {
	Link string `json:"link,omitempty"`
}

// Create LTM iFile
func (b *BigIP) CreateLtmIFile(ltmIfile *LtmIFile) error {
	return b.post(ltmIfile, uriMgmt, uriTm, uriLtm, "ifile")
}

// Get LTM iFile
func (b *BigIP) GetLtmIFile(name string) (*LtmIFile, error) {
	var ltmIfile LtmIFile
	err, ok := b.getForEntity(&ltmIfile, uriMgmt, uriTm, uriLtm, "ifile", name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &ltmIfile, nil
}

// Update LTM iFile
func (b *BigIP) UpdateLtmIFile(ltmIfile *LtmIFile) error {
	return b.put(ltmIfile, uriMgmt, uriTm, uriLtm, "ifile", ltmIfile.FullPath)
}

// Delete LTM iFile
func (b *BigIP) DeleteLtmIFile(name string) error {
	return b.delete(uriMgmt, uriTm, uriLtm, "ifile", name)
}

func (b *BigIP) CreateNTP(description string, servers []string, timezone string) error {
	config := &NTP{
		Description: description,
		Servers:     servers,
		Timezone:    timezone,
	}

	return b.patch(config, uriSys, uriNtp)
}

func (b *BigIP) ModifyNTP(config *NTP) error {
	return b.patch(config, uriSys, uriNtp)
}

func (b *BigIP) NTPs() (*NTP, error) {
	var ntp NTP
	err, _ := b.getForEntity(&ntp, uriSys, uriNtp)

	if err != nil {
		return nil, err
	}
	return &ntp, nil
}

func (b *BigIP) BigipVersion() (*Version, error) {
	var bigipversion Version
	err, _ := b.getForEntity(&bigipversion, uriMgmt, uriTm, uriCli, uriVersion)

	if err != nil {
		return nil, err
	}
	return &bigipversion, nil
}

func (b *BigIP) RunCommand(config *BigipCommand) (*BigipCommand, error) {
	var respRef BigipCommand
	resp, err := b.postReq(config, uriMgmt, uriTm, uriUtil, uriBash)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp, &respRef)
	return &respRef, nil
}

func (b *BigIP) CreateDNS(description string, nameservers []string, numberofdots int, search []string) error {
	config := &DNS{
		Description:  description,
		NameServers:  nameservers,
		NumberOfDots: numberofdots,
		Search:       search,
	}
	return b.patch(config, uriSys, uriDNS)
}

func (b *BigIP) ModifyDNS(config *DNS) error {
	return b.patch(config, uriSys, uriDNS)
}

// DNS & NTP resource does not support Delete API
func (b *BigIP) DNSs() (*DNS, error) {
	var dns DNS
	err, _ := b.getForEntity(&dns, uriSys, uriDNS)

	if err != nil {
		return nil, err
	}

	return &dns, nil
}

func (b *BigIP) CreateProvision(name string, fullPath string, cpuRatio int, diskRatio int, level string, memoryRatio int) error {
	config := &Provision{
		Name:        name,
		FullPath:    fullPath,
		CpuRatio:    cpuRatio,
		DiskRatio:   diskRatio,
		Level:       level,
		MemoryRatio: memoryRatio,
	}
	if name == "asm" {
		return b.put(config, uriSys, uriProvision, uriAsm)
	}
	if name == "afm" {
		return b.put(config, uriSys, uriProvision, uriAfm)

	}
	if name == "gtm" {
		return b.put(config, uriSys, uriProvision, uriGtm)
	}

	if name == "apm" {
		return b.put(config, uriSys, uriProvision, uriApm)
	}

	if name == "avr" {
		return b.put(config, uriSys, uriProvision, uriAvr)
	}
	if name == "ilx" {
		return b.put(config, uriSys, uriProvision, uriIlx)
	}
	return nil
}

func (b *BigIP) ProvisionModule(config *Provision) error {
	log.Printf(" Module Provision:%v", config)
	if config.Name == "asm" {
		return b.put(config, uriSys, uriProvision, uriAsm)
	}
	if config.Name == "afm" {
		return b.put(config, uriSys, uriProvision, uriAfm)
	}
	if config.Name == "gtm" {
		return b.put(config, uriSys, uriProvision, uriGtm)
	}
	if config.Name == "apm" {
		return b.put(config, uriSys, uriProvision, uriApm)
	}
	if config.Name == "avr" {
		return b.put(config, uriSys, uriProvision, uriAvr)
	}
	if config.Name == "ilx" {
		return b.put(config, uriSys, uriProvision, uriIlx)
	}
	return nil
}

func (b *BigIP) DeleteProvision(name string) error {
	// Delete API does not exists for resource Provision
	return b.delete(uriSys, uriProvision, uriIlx, name)
}

func (b *BigIP) Provisions(name string) (*Provision, error) {
	var provision Provision
	if name == "afm" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriAfm)

		if err != nil {
			return nil, err
		}
	}
	if name == "asm" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriAsm)

		if err != nil {
			return nil, err
		}
	}
	if name == "gtm" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriGtm)

		if err != nil {
			return nil, err
		}
	}
	if name == "apm" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriApm)

		if err != nil {
			return nil, err
		}
	}
	if name == "avr" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriAvr)

		if err != nil {
			return nil, err
		}

	}
	if name == "ilx" {
		err, _ := b.getForEntity(&provision, uriSys, uriProvision, uriIlx)

		if err != nil {
			return nil, err
		}

	}

	log.Println("Display ****************** provision  ", provision)
	return &provision, nil
}

func (b *BigIP) CreateSNMP(sysContact string, sysLocation string, allowedAddresses []string) error {
	config := &SNMP{
		SysContact:       sysContact,
		SysLocation:      sysLocation,
		AllowedAddresses: allowedAddresses,
	}

	return b.patch(config, uriSys, uriSnmp)
}

func (b *BigIP) ModifySNMP(config *SNMP) error {
	return b.put(config, uriSys, uriSnmp)
}

func (b *BigIP) SNMPs() (*SNMP, error) {
	var snmp SNMP
	err, _ := b.getForEntity(&snmp, uriSys, uriSnmp)

	if err != nil {
		return nil, err
	}

	return &snmp, nil
}

func (b *BigIP) CreateTRAP(name string, authPasswordEncrypted string, authProtocol string, community string, description string, engineId string, host string, port int, privacyPassword string, privacyPasswordEncrypted string, privacyProtocol string, securityLevel string, securityName string, version string) error {
	config := &TRAP{
		Name:                     name,
		AuthPasswordEncrypted:    authPasswordEncrypted,
		AuthProtocol:             authProtocol,
		Community:                community,
		Description:              description,
		EngineId:                 engineId,
		Host:                     host,
		Port:                     port,
		PrivacyPassword:          privacyPassword,
		PrivacyPasswordEncrypted: privacyPasswordEncrypted,
		PrivacyProtocol:          privacyProtocol,
		SecurityLevel:            securityLevel,
		SecurityName:             securityName,
		Version:                  version,
	}
	return b.post(config, uriSys, uriSnmp, uriTraps)
}

func (b *BigIP) StartTransaction() (*Transaction, error) {
	b.Transaction = ""
	body := make(map[string]interface{})
	resp, err := b.postReq(body, uriMgmt, uriTm, uriTransaction)

	if err != nil {
		return nil, fmt.Errorf("error encountered while starting transaction: %v", err)
	}
	transaction := &Transaction{}
	err = json.Unmarshal(resp, transaction)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Transaction: %v", transaction)
	b.Transaction = fmt.Sprint(transaction.TransID)
	return transaction, nil
}

func (b *BigIP) CommitTransaction(tId int64) error {
	b.Transaction = ""
	commitTransaction := map[string]interface{}{
		"state": "VALIDATING",
	}
	log.Printf("[INFO] Commiting Transaction with TransactionID: %v", tId)

	err := b.patch(commitTransaction, uriMgmt, uriTm, uriTransaction, strconv.Itoa(int(tId)))
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func (b *BigIP) ModifyTRAP(config *TRAP) error {
	return b.patch(config, uriSys, uriSnmp, uriTraps)
}

func (b *BigIP) TRAPs(name string) (*TRAP, error) {
	var traps TRAP
	err, _ := b.getForEntity(&traps, uriSys, uriSnmp, uriTraps, name)

	if err != nil {
		return nil, err
	}

	return &traps, nil
}

func (b *BigIP) DeleteTRAP(name string) error {
	return b.delete(uriSys, uriSnmp, uriTraps, name)
}

func (b *BigIP) Bigiplicenses() (*Bigiplicense, error) {
	var bigiplicense Bigiplicense
	err, _ := b.getForEntity(&bigiplicense, uriSys, uriLicense)

	if err != nil {
		return nil, err
	}

	return &bigiplicense, nil
}

func (b *BigIP) GetBigipLiceseStatus() (map[string]interface{}, error) {
	bigipLicense := make(map[string]interface{})
	err, _ := b.getForEntityNew(&bigipLicense, uriMgmt, uriTm, uriSys, uriLicense)
	c := 0
	for err != nil {
		time.Sleep(10 * time.Second)
		c++
		err, _ = b.getForEntityNew(&bigipLicense, uriMgmt, uriTm, uriSys, uriLicense)
		if c == 15 {
			log.Printf("[DEBUG] Device is not up even after waiting for 120 seconds")
			return nil, err
		}
	}
	return bigipLicense, nil
}

func (b *BigIP) CreateBigiplicense(command, registration_key string) error {
	config := &Bigiplicense{
		Command:          command,
		Registration_key: registration_key,
	}

	return b.post(config, uriSys, uriLicense)
}

func (b *BigIP) ModifyBigiplicense(config *Bigiplicense) error {
	return b.put(config, uriSys, uriLicense)
}

func (b *BigIP) LogIPFIXs() (*LogIPFIX, error) {
	var logipfix LogIPFIX
	err, _ := b.getForEntity(&logipfix, uriSys, uriLogConfig, uriDestination, uriIPFIX)

	if err != nil {
		return nil, err
	}

	return &logipfix, nil
}

func (b *BigIP) CreateLogIPFIX(name, appService, poolName, protocolVersion, serversslProfile string, templateDeleteDelay, templateRetransmitInterval int, transportProfile string) error {
	config := &LogIPFIX{
		Name:                       name,
		AppService:                 appService,
		PoolName:                   poolName,
		ProtocolVersion:            protocolVersion,
		ServersslProfile:           serversslProfile,
		TemplateDeleteDelay:        templateDeleteDelay,
		TemplateRetransmitInterval: templateRetransmitInterval,
		TransportProfile:           transportProfile,
	}

	return b.post(config, uriSys, uriLogConfig, uriDestination, uriIPFIX)
}

func (b *BigIP) ModifyLogIPFIX(config *LogIPFIX) error {
	return b.put(config, uriSys, uriLogConfig, uriDestination, uriIPFIX)
}

func (b *BigIP) DeleteLogIPFIX(name string) error {
	return b.delete(uriSys, uriLogConfig, uriDestination, uriIPFIX, name)
}

func (b *BigIP) LogPublisher() (*LogPublisher, error) {
	var logpublisher LogPublisher
	err, _ := b.getForEntity(&logpublisher, uriSys, uriLogConfig, uriPublisher)

	if err != nil {
		return nil, err
	}

	return &logpublisher, nil
}

func (b *BigIP) CreateLogPublisher(r *LogPublisher) error {
	return b.post(r, uriSys, uriLogConfig, uriPublisher)
}

func (b *BigIP) ModifyLogPublisher(r *LogPublisher) error {
	return b.put(r, uriSys, uriLogConfig, uriPublisher)
}

func (b *BigIP) DeleteLogPublisher(name string) error {
	return b.delete(uriSys, uriLogConfig, uriPublisher, name)
}

// UploadDatagroup copies a template set from local disk to BIGIP
func (b *BigIP) UploadDatagroup(tmplpath *os.File, dgname, partition, dgtype string, createDg bool) error {
	_, err := b.UploadDataGroupFile(tmplpath, dgname)
	if err != nil {
		return err
	}
	sourcepath := "file://" + REST_DOWNLOAD_PATH + "/" + dgname
	log.Printf("[DEBUG] sourcepath :%+v", sourcepath)
	dataGroup := ExternalDGFile{
		Name:       dgname,
		SourcePath: sourcepath,
		Partition:  partition,
		Type:       dgtype,
	}
	log.Printf("External DG: %+v\n", dataGroup)
	if createDg {
		err = b.AddExternalDatagroupfile(&dataGroup)
		if err != nil {
			return err
		}
	} else {
		err = b.ModifyExternalDatagroupfile(fmt.Sprintf("/%s/%s", partition, dgname), &dataGroup)
		if err != nil {
			return err
		}
	}

	dataGroup2 := ExternalDG{
		Name:             dgname,
		ExternalFileName: fmt.Sprintf("/%s/%s", partition, dgname),
		FullPath:         fmt.Sprintf("/%s/%s", partition, dgname),
	}
	if createDg {
		err = b.AddExternalDataGroup(&dataGroup2)
		if err != nil {
			return err
		}
	} else {
		err = b.ModifyExternalDataGroup(&dataGroup2)
		if err != nil {
			return err
		}
	}
	return nil
}

// Upload a file
func (b *BigIP) UploadDataGroupFile(f *os.File, tmpName string) (*Upload, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	log.Printf("tmpName:%+v", tmpName)
	return b.Upload(f, info.Size(), uriShared, uriFileTransfer, uriUploads, tmpName)
}

func (b *BigIP) CreateOCSP(ocsp *OCSP) error {
	return b.post(ocsp, uriSys, "crypto", "cert-validator", "ocsp")
}

func (b *BigIP) ModifyOCSP(name string, ocsp *OCSP) error {
	return b.put(ocsp, uriSys, "crypto", "cert-validator", "ocsp", name)
}

func (b *BigIP) GetOCSP(name string) (*OCSP, error) {
	var ocsp OCSP
	err, _ := b.getForEntity(&ocsp, uriSys, "crypto", "cert-validator", "ocsp", name)

	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ocsp)

	if err != nil {
		return nil, fmt.Errorf("error encountered while marshalling ocsp: %v", err)
	}

	if string(js) == "{}" {
		return nil, nil
	}

	return &ocsp, nil
}

func (b *BigIP) DeleteOCSP(name string) error {
	return b.delete(uriSys, "crypto", "cert-validator", "ocsp", name)
}

func (b *BigIP) CreatePartition(partition *Partition) error {
	return b.post(partition, uriAuth, uriPartition)
}

func (b *BigIP) GetPartition(name string) (*Partition, error) {
	var partition Partition
	err, ok := b.getForEntity(&partition, uriAuth, uriPartition, name)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, nil
	}

	return &partition, err
}

func (b *BigIP) ModifyPartition(name string, partition *Partition) error {
	return b.patch(partition, uriAuth, uriPartition, name)
}

func (b *BigIP) DeletePartition(name string) error {
	return b.delete(uriAuth, uriPartition, name)
}

func (b *BigIP) ModifyFolderDescription(partition string, body map[string]string) error {
	partition = fmt.Sprintf("~%s", partition)
	return b.patch(body, uriSys, uriFolder, partition)
}

func (b *BigIP) CreateRoleInfo(roleInfo *RoleInfo) error {
	return b.post(roleInfo, uriAuth, uriRemoteRole, uriRoleInfo)
}

func (b *BigIP) GetRoleInfo(name string) (*RoleInfo, error) {
	var roleInfo RoleInfo
	err, ok := b.getForEntity(&roleInfo, uriAuth, uriRemoteRole, uriRoleInfo, name)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, nil
	}

	return &roleInfo, err
}

func (b *BigIP) ModifyRoleInfo(name string, roleInfo *RoleInfo) error {
	return b.patch(roleInfo, uriAuth, uriRemoteRole, uriRoleInfo, name)
}

func (b *BigIP) DeleteRoleInfo(name string) error {
	return b.delete(uriAuth, uriRemoteRole, uriRoleInfo, name)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////               Management Route             /////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// ManagementRoutes represents a collection of BIG-IP management routes.
type ManagementRoutes struct {
	ManagementRoutes []ManagementRoute `json:"items"`
}

// ManagementRoute represents a BIG-IP management route configuration.
type ManagementRoute struct {
	Name        string `json:"name,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	MTU         int    `json:"mtu,omitempty"`
	Network     string `json:"network,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetManagementRoutes returns a list of management routes.
func (b *BigIP) GetManagementRoutes() (*ManagementRoutes, error) {
	var mgmtroute ManagementRoutes
	err, ok := b.getForEntity(&mgmtroute, uriSys, uriMgmtRoute)
	if err != nil {
		if !ok {
			return &ManagementRoutes{}, nil
		}
		return nil, err
	}

	return &mgmtroute, nil
}

// GetManagementRoute returns a named Management Route.
func (b *BigIP) GetManagementRoute(managementroute string) (*ManagementRoute, error) {
	var mgmtroute ManagementRoute
	err, ok := b.getForEntity(&mgmtroute, uriSys, uriMgmtRoute, managementroute)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &mgmtroute, nil
}

// CreateManagementRoute adds a new management route to the BIG-IP system. <dest> must include the
// subnet mask in CIDR notation, i.e.: "10.1.1.0/24".
func (b *BigIP) CreateManagementRoute(config *ManagementRoute) error {
	return b.post(config, uriSys, uriMgmtRoute)
}

// DeleteManagementRoute removes a management route.
func (b *BigIP) DeleteManagementRoute(name string) error {
	return b.delete(uriSys, uriMgmtRoute, name)
}

// ModifyManagementRoute allows for a change of any attribute of a management route. Fields that
// can be modified are referenced in the ManagementRoute struct.
func (b *BigIP) ModifyManagementRoute(name string, config *ManagementRoute) error {
	return b.put(config, uriSys, uriMgmtRoute, name)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////          Management Firewall Rule          /////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// MgmtFirewallRules represents a collection of BIG-IP management firewall rules.
type MgmtFirewallRules struct {
	MgmtFirewallRules []MgmtFirewallRule `json:"items"`
}

// MgmtFwRuleAddress represents an IP address or address range for firewall matching.
type MgmtFwRuleAddress struct {
	Name string `json:"name,omitempty"`
}

// MgmtFwRulePort represents a port or port range for firewall matching.
type MgmtFwRulePort struct {
	Name string `json:"name,omitempty"`
}

// MgmtFwRuleICMP represents an ICMP type and code specification for firewall rules.
// Format is "type:code" (e.g., "3:1" for destination unreachable / host unreachable).
type MgmtFwRuleICMP struct {
	Name string `json:"name,omitempty"`
}

// MgmtFwRuleIpPortData contains source or destination IP and port matching criteria for firewall rules.
type MgmtFwRuleIpPortData struct {
	AddressLists []string            `json:"addressLists,omitempty"`
	Addresses    []MgmtFwRuleAddress `json:"addresses,omitzero"`
	PortLists    []string            `json:"portLists,omitempty"`
	Ports        []MgmtFwRulePort    `json:"ports,omitzero"`
}

// MgmtFirewallRule represents a BIG-IP management firewall rule configuration.
// These rules control access to the management interface.
type MgmtFirewallRule struct {
	Name        string               `json:"name,omitempty"`
	FullPath    string               `json:"fullPath,omitempty"`
	Description string               `json:"description,omitempty"`
	UUID        string               `json:"uuid,omitempty"`
	Action      string               `json:"action,omitempty" validate:"oneof=accept drop reject"`
	IpProtocol  string               `json:"ipProtocol,omitempty"`
	Log         string               `json:"log,omitempty" validate:"oneof=yes no"`
	PlaceAfter  string               `json:"placeAfter,omitempty"`
	PlaceBefore string               `json:"placeBefore,omitempty"`
	Status      string               `json:"status,omitempty"`
	Schedule    string               `json:"schedule,omitempty"`
	Destination MgmtFwRuleIpPortData `json:"destination,omitempty"`
	Source      MgmtFwRuleIpPortData `json:"source,omitempty"`
	ICMPs       []MgmtFwRuleICMP     `json:"icmp,omitzero"`
}

func validateMgmtRulePlacing(config *MgmtFirewallRule) error {
	hasAfter := config.PlaceAfter != ""
	hasBefore := config.PlaceBefore != ""

	if hasAfter && hasBefore {
		return errors.New("cannot set both 'PlaceAfter' and 'PlaceBefore', choose one")
	}
	if !hasAfter && !hasBefore {
		return errors.New("must set either 'PlaceAfter' or 'PlaceBefore' for rule placement")
	}
	return nil
}

// GetManagementFwRules returns a list of management firewall rules.
func (b *BigIP) GetManagementFwRules() (*MgmtFirewallRules, error) {
	var mgmtfwrules MgmtFirewallRules
	err, ok := b.getForEntity(&mgmtfwrules, uriSecurity, uriFirewall, uriMgmtIpRules, uriRules)
	if err != nil {
		if !ok {
			return &MgmtFirewallRules{}, nil
		}
		return nil, err
	}

	return &mgmtfwrules, nil
}

// GetManagementFwRule returns a named Management Firewall Rule.
func (b *BigIP) GetManagementFwRule(name string) (*MgmtFirewallRule, error) {
	var mgmtfwrule MgmtFirewallRule
	err, ok := b.getForEntity(&mgmtfwrule, uriSecurity, uriFirewall, uriMgmtIpRules, uriRules, name)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &mgmtfwrule, nil
}

// CreateManagementFwRule adds a new management firewall rule to the BIG-IP system.
func (b *BigIP) CreateManagementFwRule(config *MgmtFirewallRule) error {
	err := validateMgmtRulePlacing(config)
	if err != nil {
		return err
	}
	return b.post(config, uriSecurity, uriFirewall, uriMgmtIpRules, uriRules)
}

// ModifyManagementFwRule allows you to change any attribute of a management firewall rule.
func (b *BigIP) ModifyManagementFwRule(name string, config *MgmtFirewallRule) error {
	return b.put(config, uriSecurity, uriFirewall, uriMgmtIpRules, uriRules, name)
}

// DeleteManagementFwRule removes a management firewall rule.
func (b *BigIP) DeleteManagementFwRule(name string) error {
	return b.delete(uriSecurity, uriFirewall, uriMgmtIpRules, uriRules, name)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////        Authentication - Remote Role         ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// RemoteRoles represent a collection of the user roles on the BIGIP system.
type RemoteRoles struct {
	RemoteRoles []RemoteRole `json:"items"`
}

// // RemoteRoles represents a specific user role on the BIGIP system.
type RemoteRole struct {
	Name          string `json:"name,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Attribute     string `json:"attribute,omitempty"`
	Console       string `json:"console,omitempty"`
	Deny          string `json:"deny,omitempty"`
	Description   string `json:"description,omitempty"`
	LineOrder     int    `json:"lineOrder,omitempty"`
	Role          string `json:"role,omitempty"`
	UserPartition string `json:"userPartition,omitempty"`
}

// GetRemoteRoles returns a list of all remote role configurations.
func (b *BigIP) GetRemoteRoles() (*RemoteRoles, error) {
	var remoteRoles RemoteRoles
	err, ok := b.getForEntity(&remoteRoles, uriAuth, uriRemoteRole, uriRoleInfo)
	if err != nil {
		if !ok {
			return &RemoteRoles{}, nil
		}
		return nil, err
	}

	return &remoteRoles, nil
}

// GetRemoteRole returns a named remote role configuration.
func (b *BigIP) GetRemoteRole(name string) (*RemoteRole, error) {
	var remoteRole RemoteRole
	err, ok := b.getForEntity(&remoteRole, uriAuth, uriRemoteRole, uriRoleInfo, name)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &remoteRole, nil
}

func validateRemoteRole(config *RemoteRole) error {
	if config.Attribute == "" {
		return errors.New("attribute is required (e.g., 'memberof=cn=group,dc=example,dc=com')")
	}
	if config.LineOrder == 0 {
		return errors.New("lineOrder is required")
	}
	return nil
}

// CreateRemoteRole adds a new remote role configuration to the BIG-IP system.
func (b *BigIP) CreateRemoteRole(config *RemoteRole) error {
	err := validateRemoteRole(config)
	if err != nil {
		return err
	}
	return b.post(config, uriAuth, uriRemoteRole, uriRoleInfo)
}

// ModifyRemoteRole allows the update of the attributes.
func (b *BigIP) ModifyRemoteRole(name string, config *RemoteRole) error {
	return b.patch(config, uriAuth, uriRemoteRole, uriRoleInfo, name)
}

// DeleteRemoteRole removes a remote role configuration.
func (b *BigIP) DeleteRemoteRole(name string) error {
	return b.delete(uriAuth, uriRemoteRole, uriRoleInfo, name)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////            LDAP Authentication              ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// LDAPConfig represents LDAP configuration used for authentication to BIGIP system.
type LdapConfig struct {
	Name                  string   `json:"name,omitempty" validate:"eq=system-auth"`
	BindDn                string   `json:"bindDn,omitempty"`
	BindPw                string   `json:"bindPw,omitempty"`
	BindTimeout           int      `json:"bindTimeout,omitempty"`
	CheckHostAttr         string   `json:"checkHostAttr,omitempty"`
	CheckRolesGroup       string   `json:"checkRolesGroup,omitempty"`
	Debug                 string   `json:"debug,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	GroupDn               string   `json:"groupDn,omitempty"`
	GroupMemberAttribute  string   `json:"groupMemberAttribute,omitempty"`
	IdleTimeout           int      `json:"idleTimeout,omitempty"`
	IgnoreAuthInfoUnavail string   `json:"ignoreAuthInfoUnavail,omitempty"`
	IgnoreUnknownUser     string   `json:"ignoreUnknownUser,omitempty"`
	LoginAttribute        string   `json:"loginAttribute,omitempty"`
	Port                  int      `json:"port,omitempty"`
	Referrals             string   `json:"referrals,omitempty"`
	Scope                 string   `json:"scope,omitempty"`
	SearchBaseDn          string   `json:"searchBaseDn,omitempty"`
	SearchTimeout         int      `json:"searchTimeout,omitempty"`
	Servers               []string `json:"servers,omitempty"`
	Ssl                   string   `json:"ssl,omitempty"`
	SslCaCertFile         string   `json:"sslCaCertFile,omitempty"`
	SslCheckPeer          string   `json:"sslCheckPeer,omitempty"`
	SslCiphers            string   `json:"sslCiphers,omitempty"`
	SslClientCert         string   `json:"sslClientCert,omitempty"`
	SslClientKey          string   `json:"sslClientKey,omitempty"`
	UserTemplate          string   `json:"userTemplate,omitempty"`
	Version               int      `json:"version,omitempty"`
	Warnings              string   `json:"warnings,omitempty"`
}

// GetLdapConfig returns a named LDAP authentication configuration.
func (b *BigIP) GetLdapConfig(name string) (*LdapConfig, error) {
	var ldapConfig LdapConfig
	err, ok := b.getForEntity(&ldapConfig, uriAuth, uriLdap, name)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &ldapConfig, nil
}

func validateLdapConfig(config *LdapConfig) error {
	if len(config.Servers) == 0 {
		return errors.New("servers is required option, at least one LDAP server must be specified")
	}
	if config.Name != "" && config.Name != "system-auth" {
		return errors.New("name must be 'system-auth'")
	}
	return nil
}

// CreateLdapConfig adds a new LDAP authentication configuration to the BIG-IP system.
// BIG-IP only allows the name "system-auth" for LDAP configurations
func (b *BigIP) CreateLdapConfig(config *LdapConfig) error {
	err := validateLdapConfig(config)
	if err != nil {
		return err
	}
	return b.post(config, uriAuth, uriLdap)
}

// ModifyLdapConfig allows for a change of attribute in LDAP configuration
func (b *BigIP) ModifyLdapConfig(name string, config *LdapConfig) error {
	return b.patch(config, uriAuth, uriLdap, name)
}

// DeleteLdapConfig removes an LDAP authentication configuration.
func (b *BigIP) DeleteLdapConfig(name string) error {
	return b.delete(uriAuth, uriLdap, name)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////          Authentication Source              ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// AuthSource represents authentication source used to access BIGIP system.
type AuthSource struct {
	Fallback string `json:"fallback,omitempty"`
	Type     string `json:"type,omitempty"`
}

// GetAuthSource retrieves the current authentication source configuration.
func (b *BigIP) GetAuthSource() (*AuthSource, error) {
	var authSource AuthSource
	err, ok := b.getForEntity(&authSource, uriAuth, uriAuthSrc)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &authSource, nil
}

// CreateAuthSource sets the authentication source configuration.
func (b *BigIP) CreateAuthSource(config *AuthSource) error {
	return b.patch(config, uriAuth, uriAuthSrc)
}

// ModifyAuthSource updates the authentication source configuration.
func (b *BigIP) ModifyAuthSource(config *AuthSource) error {
	return b.patch(config, uriAuth, uriAuthSrc)
}

// DeleteAuthSource resets the authentication source to default values.
func (b *BigIP) DeleteAuthSource() error {
	defaultConfig := &AuthSource{
		Type:     "local",
		Fallback: "false",
	}
	return b.patch(defaultConfig, uriAuth, uriAuthSrc)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////            Authentication Remote User         //////////////////
//////////////////////////////////////////////////////////////////////////////////////

// RemoteUser represents the authorization data for external users
type RemoteUser struct {
	DefaultPartition    string `json:"defaultPartition,omitempty"`
	DefaultRole         string `json:"defaultRole,omitempty"`
	RemoteConsoleAccess string `json:"remoteConsoleAccess,omitempty"`
}

// GetRemoteUser retrieves the current remote user configuration.
func (b *BigIP) GetRemoteUser() (*RemoteUser, error) {
	var remoteUser RemoteUser
	err, ok := b.getForEntity(&remoteUser, uriAuth, uriRemoteUser)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &remoteUser, nil
}

// CreateRemoteUser sets the remote user configuration.
func (b *BigIP) CreateRemoteUser(config *RemoteUser) error {
	return b.patch(config, uriAuth, uriRemoteUser)
}

// ModifyRemoteUser updates the remote user configuration.
func (b *BigIP) ModifyRemoteUser(config *RemoteUser) error {
	return b.patch(config, uriAuth, uriRemoteUser)
}

// DeleteRemoteUser resets the remote user configuration to default values.
func (b *BigIP) DeleteRemoteUser() error {
	defaultConfig := &RemoteUser{
		DefaultPartition:    "all",
		DefaultRole:         "no-access",
		RemoteConsoleAccess: "disabled",
	}
	return b.patch(defaultConfig, uriAuth, uriRemoteUser)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////            System Syslog Configuration      ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// SyslogRemoteServer represents a remote syslog server configuration.
type SyslogRemoteServer struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Host        string `json:"host,omitempty"`
	LocalIp     string `json:"localIp,omitempty"`
	RemotePort  int    `json:"remotePort,omitempty"`
}

// SyslogConfig represents the BIG-IP system syslog configuration.
type SyslogConfig struct {
	AuthPrivFrom         string               `json:"authPrivFrom,omitempty"`
	AuthPrivTo           string               `json:"authPrivTo,omitempty"`
	ClusteredHostSlot    string               `json:"clusteredHostSlot,omitempty"`
	ClusteredMessageSlot string               `json:"clusteredMessageSlot,omitempty"`
	ConsoleLog           string               `json:"consoleLog,omitempty"`
	CronFrom             string               `json:"cronFrom,omitempty"`
	CronTo               string               `json:"cronTo,omitempty"`
	DaemonFrom           string               `json:"daemonFrom,omitempty"`
	DaemonTo             string               `json:"daemonTo,omitempty"`
	Description          string               `json:"description,omitempty"`
	Include              string               `json:"include,omitempty"`
	IsoDate              string               `json:"isoDate,omitempty"`
	KernFrom             string               `json:"kernFrom,omitempty"`
	KernTo               string               `json:"kernTo,omitempty"`
	Local6From           string               `json:"local6From,omitempty"`
	Local6To             string               `json:"local6To,omitempty"`
	MailFrom             string               `json:"mailFrom,omitempty"`
	MailTo               string               `json:"mailTo,omitempty"`
	MessagesFrom         string               `json:"messagesFrom,omitempty"`
	MessagesTo           string               `json:"messagesTo,omitempty"`
	UserLogFrom          string               `json:"userLogFrom,omitempty"`
	UserLogTo            string               `json:"userLogTo,omitempty"`
	RemoteServers        []SyslogRemoteServer `json:"remoteServers,omitzero"`
}

// GetSyslogConfig retrieves the current system syslog configuration.
func (b *BigIP) GetSyslogConfig() (*SyslogConfig, error) {
	var syslogConfig SyslogConfig
	err, ok := b.getForEntity(&syslogConfig, uriSys, uriSyslog)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &syslogConfig, nil
}

// CreateSyslogConfig sets the system syslog configuration.
func (b *BigIP) CreateSyslogConfig(config *SyslogConfig) error {
	return b.patch(config, uriSys, uriSyslog)
}

// ModifySyslogConfig updates the system syslog configuration.
func (b *BigIP) ModifySyslogConfig(config *SyslogConfig) error {
	return b.patch(config, uriSys, uriSyslog)
}

// DeleteSyslogConfig resets the system syslog configuration to default values.
func (b *BigIP) DeleteSyslogConfig() error {
	defaultConfig := &SyslogConfig{
		AuthPrivFrom:         "notice",
		AuthPrivTo:           "emerg",
		ClusteredHostSlot:    "enabled",
		ClusteredMessageSlot: "disabled",
		ConsoleLog:           "enabled",
		CronFrom:             "warning",
		CronTo:               "emerg",
		DaemonFrom:           "notice",
		DaemonTo:             "emerg",
		Description:          "none",
		Include:              "none",
		IsoDate:              "disabled",
		KernFrom:             "debug",
		KernTo:               "emerg",
		Local6From:           "notice",
		Local6To:             "emerg",
		MailFrom:             "notice",
		MailTo:               "emerg",
		MessagesFrom:         "notice",
		MessagesTo:           "warning",
		UserLogFrom:          "notice",
		UserLogTo:            "emerg",
		RemoteServers:        []SyslogRemoteServer{},
	}
	return b.patch(defaultConfig, uriSys, uriSyslog)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////         System Global Settings              ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// GlobalSettings represents the BIG-IP system global settings configuration.
// Only a subset of attributes has been implemented.
type GlobalSettings struct {
	GuiSecurityBanner     string `json:"guiSecurityBanner,omitempty"`
	GuiSecurityBannerText string `json:"guiSecurityBannerText,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
}

// GetGlobalSettings retrieves the current system global settings configuration.
func (b *BigIP) GetGlobalSettings() (*GlobalSettings, error) {
	var globalSettings GlobalSettings
	err, ok := b.getForEntity(&globalSettings, uriSys, uriGlobalSettings)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &globalSettings, nil
}

// CreateGlobalSettings sets the system global settings configuration using PATCH.
func (b *BigIP) CreateGlobalSettings(config *GlobalSettings) error {
	return b.patch(config, uriSys, uriGlobalSettings)
}

// ModifyGlobalSettings updates the system global settings configuration.
func (b *BigIP) ModifyGlobalSettings(config *GlobalSettings) error {
	return b.patch(config, uriSys, uriGlobalSettings)
}

// DeleteGlobalSettings resets the system global settings configuration to default values.
func (b *BigIP) DeleteGlobalSettings() error {
	defaultConfig := &GlobalSettings{
		GuiSecurityBanner:     "enabled",
		GuiSecurityBannerText: "Welcome to the BIG-IP Configuration Utility.\n\nLog in with your username and password using the fields on the left.",
		Hostname:              "bigip1",
	}
	return b.patch(defaultConfig, uriSys, uriGlobalSettings)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////           SSHD Configuration                ////////////////////
//////////////////////////////////////////////////////////////////////////////////////

// SSHDConfig represents the BIG-IP SSHD (SSH daemon) configuration.
type SSHDConfig struct {
	Allow             []string `json:"allow,omitempty"`
	Banner            string   `json:"banner,omitempty"`
	BannerText        string   `json:"bannerText,omitempty"`
	FipsCipherVersion int      `json:"fipsCipherVersion,omitempty"`
	InactivityTimeout int      `json:"inactivityTimeout,omitempty"`
	Include           string   `json:"include,omitempty"`
	LogLevel          string   `json:"logLevel,omitempty"`
	Login             string   `json:"login,omitempty"`
	Port              int      `json:"port,omitempty"`
}

// GetSSHDConfig retrieves the current SSHD configuration.
func (b *BigIP) GetSSHDConfig() (*SSHDConfig, error) {
	var sshdConfig SSHDConfig
	err, ok := b.getForEntity(&sshdConfig, uriSys, uriSshd)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &sshdConfig, nil
}

// CreateSSHDConfig sets the SSHD configuration.
func (b *BigIP) CreateSSHDConfig(config *SSHDConfig) error {
	return b.patch(config, uriSys, uriSshd)
}

// ModifySSHDConfig updates the SSHD configuration.
func (b *BigIP) ModifySSHDConfig(config *SSHDConfig) error {
	return b.patch(config, uriSys, uriSshd)
}

// DeleteSSHDConfig resets the SSHD configuration to default values.
func (b *BigIP) DeleteSSHDConfig() error {
	defaultConfig := &SSHDConfig{
		Allow:             []string{"ALL"},
		Banner:            "disabled",
		BannerText:        "none",
		InactivityTimeout: 0,
		Include:           "none",
		LogLevel:          "info",
		Login:             "enabled",
		Port:              22,
	}
	return b.patch(defaultConfig, uriSys, uriSshd)
}

//////////////////////////////////////////////////////////////////////////////////////
/////////////////////           HTTPD Configuration                ///////////////////
//////////////////////////////////////////////////////////////////////////////////////

// HTTPDConfig represents the BIG-IP HTTPD (HTTP daemon) configuration.
type HTTPDConfig struct {
	Allow                    []string `json:"allow,omitempty"`
	AuthName                 string   `json:"authName,omitempty"`
	AuthPamDashboardTimeout  string   `json:"authPamDashboardTimeout,omitempty"`
	AuthPamIdleTimeout       int      `json:"authPamIdleTimeout,omitempty"`
	AuthPamValidateIp        string   `json:"authPamValidateIp,omitempty"`
	FastcgiTimeout           int      `json:"fastcgiTimeout,omitempty"`
	FipsCipherVersion        int      `json:"fipsCipherVersion,omitempty"`
	HostnameLookup           string   `json:"hostnameLookup,omitempty"`
	Include                  string   `json:"include,omitempty"`
	LogLevel                 string   `json:"logLevel,omitempty"`
	MaxClients               int      `json:"maxClients,omitempty"`
	RedirectHttpToHttps      string   `json:"redirectHttpToHttps,omitempty"`
	RequestBodyMaxTimeout    int      `json:"requestBodyMaxTimeout,omitempty"`
	RequestBodyMinRate       int      `json:"requestBodyMinRate,omitempty"`
	RequestBodyTimeout       int      `json:"requestBodyTimeout,omitempty"`
	RequestHeaderMaxTimeout  int      `json:"requestHeaderMaxTimeout,omitempty"`
	RequestHeaderMinRate     int      `json:"requestHeaderMinRate,omitempty"`
	RequestHeaderTimeout     int      `json:"requestHeaderTimeout,omitempty"`
	SslCaCertFile            string   `json:"sslCaCertFile,omitempty"`
	SslCertchainfile         string   `json:"sslCertchainfile,omitempty"`
	SslCertfile              string   `json:"sslCertfile,omitempty"`
	SslCertkeyfile           string   `json:"sslCertkeyfile,omitempty"`
	SslCiphersuite           string   `json:"sslCiphersuite,omitempty"`
	SslInclude               string   `json:"sslInclude,omitempty"`
	SslOcspDefaultResponder  string   `json:"sslOcspDefaultResponder,omitempty"`
	SslOcspEnable            string   `json:"sslOcspEnable,omitempty"`
	SslOcspOverrideResponder string   `json:"sslOcspOverrideResponder,omitempty"`
	SslOcspResponderTimeout  int      `json:"sslOcspResponderTimeout,omitempty"`
	SslOcspResponseMaxAge    int      `json:"sslOcspResponseMaxAge,omitempty"`
	SslOcspResponseTimeSkew  int      `json:"sslOcspResponseTimeSkew,omitempty"`
	SslPort                  int      `json:"sslPort,omitempty"`
	SslProtocol              string   `json:"sslProtocol,omitempty"`
	SslVerifyClient          string   `json:"sslVerifyClient,omitempty"`
	SslVerifyDepth           int      `json:"sslVerifyDepth,omitempty"`
}

// GetHTTPDConfig retrieves the current HTTPD configuration.
func (b *BigIP) GetHTTPDConfig() (*HTTPDConfig, error) {
	var httpdConfig HTTPDConfig
	err, ok := b.getForEntity(&httpdConfig, uriSys, uriHttpd)
	if err != nil {
		if !ok {
			return nil, nil
		}
		return nil, err
	}

	return &httpdConfig, nil
}

// CreateHTTPDConfig sets the HTTPD configuration.
func (b *BigIP) CreateHTTPDConfig(config *HTTPDConfig) error {
	return b.patch(config, uriSys, uriHttpd)
}

// ModifyHTTPDConfig updates the HTTPD configuration.
func (b *BigIP) ModifyHTTPDConfig(config *HTTPDConfig) error {
	return b.patch(config, uriSys, uriHttpd)
}

// DeleteHTTPDConfig resets the HTTPD configuration to default values.
func (b *BigIP) DeleteHTTPDConfig() error {
	defaultConfig := &HTTPDConfig{
		Allow:                    []string{"All"},
		AuthName:                 "BIG-IP",
		AuthPamDashboardTimeout:  "off",
		AuthPamIdleTimeout:       1200,
		AuthPamValidateIp:        "on",
		FastcgiTimeout:           300,
		FipsCipherVersion:        0,
		HostnameLookup:           "off",
		Include:                  "none",
		LogLevel:                 "warn",
		MaxClients:               10,
		RedirectHttpToHttps:      "disabled",
		RequestBodyMaxTimeout:    0,
		RequestBodyMinRate:       500,
		RequestBodyTimeout:       60,
		RequestHeaderMaxTimeout:  40,
		RequestHeaderMinRate:     500,
		RequestHeaderTimeout:     20,
		SslCaCertFile:            "none",
		SslCertchainfile:         "none",
		SslCertfile:              "/etc/httpd/conf/ssl.crt/server.crt",
		SslCertkeyfile:           "/etc/httpd/conf/ssl.key/server.key",
		SslCiphersuite:           "ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA:ECDHE-ECDSA-AES128-SHA256:ECDHE-ECDSA-AES256-SHA384:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA:AES256-SHA:AES128-SHA256:AES256-SHA256",
		SslInclude:               "none",
		SslOcspDefaultResponder:  "http://127.0.0.1",
		SslOcspEnable:            "off",
		SslOcspOverrideResponder: "off",
		SslOcspResponderTimeout:  300,
		SslOcspResponseMaxAge:    -1,
		SslOcspResponseTimeSkew:  300,
		SslPort:                  443,
		SslProtocol:              "all -SSLv2 -SSLv3 -TLSv1",
		SslVerifyClient:          "no",
		SslVerifyDepth:           10,
	}
	return b.patch(defaultConfig, uriSys, uriHttpd)
}
