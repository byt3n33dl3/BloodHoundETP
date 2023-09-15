// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by Cuelang code gen. DO NOT EDIT!
// Cuelang source: github.com/specterops/bloodhound/-/tree/main/packages/cue/schemas/

package ad

import (
	"errors"
	graph "github.com/specterops/bloodhound/dawgs/graph"
)

var (
	Entity                          = graph.StringKind("Base")
	User                            = graph.StringKind("User")
	Computer                        = graph.StringKind("Computer")
	Group                           = graph.StringKind("Group")
	GPO                             = graph.StringKind("GPO")
	OU                              = graph.StringKind("OU")
	Container                       = graph.StringKind("Container")
	Domain                          = graph.StringKind("Domain")
	LocalGroup                      = graph.StringKind("ADLocalGroup")
	LocalUser                       = graph.StringKind("ADLocalUser")
	Owns                            = graph.StringKind("Owns")
	GenericAll                      = graph.StringKind("GenericAll")
	GenericWrite                    = graph.StringKind("GenericWrite")
	WriteOwner                      = graph.StringKind("WriteOwner")
	WriteDACL                       = graph.StringKind("WriteDacl")
	MemberOf                        = graph.StringKind("MemberOf")
	ForceChangePassword             = graph.StringKind("ForceChangePassword")
	AllExtendedRights               = graph.StringKind("AllExtendedRights")
	AddMember                       = graph.StringKind("AddMember")
	HasSession                      = graph.StringKind("HasSession")
	Contains                        = graph.StringKind("Contains")
	GPLink                          = graph.StringKind("GPLink")
	AllowedToDelegate               = graph.StringKind("AllowedToDelegate")
	GetChanges                      = graph.StringKind("GetChanges")
	GetChangesAll                   = graph.StringKind("GetChangesAll")
	GetChangesInFilteredSet         = graph.StringKind("GetChangesInFilteredSet")
	TrustedBy                       = graph.StringKind("TrustedBy")
	AllowedToAct                    = graph.StringKind("AllowedToAct")
	AdminTo                         = graph.StringKind("AdminTo")
	CanPSRemote                     = graph.StringKind("CanPSRemote")
	CanRDP                          = graph.StringKind("CanRDP")
	ExecuteDCOM                     = graph.StringKind("ExecuteDCOM")
	HasSIDHistory                   = graph.StringKind("HasSIDHistory")
	AddSelf                         = graph.StringKind("AddSelf")
	DCSync                          = graph.StringKind("DCSync")
	ReadLAPSPassword                = graph.StringKind("ReadLAPSPassword")
	ReadGMSAPassword                = graph.StringKind("ReadGMSAPassword")
	DumpSMSAPassword                = graph.StringKind("DumpSMSAPassword")
	SQLAdmin                        = graph.StringKind("SQLAdmin")
	AddAllowedToAct                 = graph.StringKind("AddAllowedToAct")
	WriteSPN                        = graph.StringKind("WriteSPN")
	AddKeyCredentialLink            = graph.StringKind("AddKeyCredentialLink")
	LocalToComputer                 = graph.StringKind("LocalToComputer")
	MemberOfLocalGroup              = graph.StringKind("MemberOfLocalGroup")
	RemoteInteractiveLogonPrivilege = graph.StringKind("RemoteInteractiveLogonPrivilege")
	SyncLAPSPassword                = graph.StringKind("SyncLAPSPassword")
	WriteAccountRestrictions        = graph.StringKind("WriteAccountRestrictions")
)

type Property string

const (
	AdminCount              Property = "admincount"
	DistinguishedName       Property = "distinguishedname"
	DomainFQDN              Property = "domain"
	DomainSID               Property = "domainsid"
	Sensitive               Property = "sensitive"
	HighValue               Property = "highvalue"
	BlocksInheritance       Property = "blocksinheritance"
	IsACL                   Property = "isacl"
	IsACLProtected          Property = "isaclprotected"
	Enforced                Property = "enforced"
	Department              Property = "department"
	HasSPN                  Property = "hasspn"
	UnconstrainedDelegation Property = "unconstraineddelegation"
	LastLogon               Property = "lastlogon"
	LastLogonTimestamp      Property = "lastlogontimestamp"
	IsPrimaryGroup          Property = "isprimarygroup"
	HasLAPS                 Property = "haslaps"
	DontRequirePreAuth      Property = "dontreqpreauth"
	LogonType               Property = "logontype"
	HasURA                  Property = "hasura"
	PasswordNeverExpires    Property = "pwdneverexpires"
	PasswordNotRequired     Property = "passwordnotreqd"
	FunctionalLevel         Property = "functionallevel"
	TrustType               Property = "trusttype"
	SidFiltering            Property = "sidfiltering"
	TrustedToAuth           Property = "trustedtoauth"
)

func AllProperties() []Property {
	return []Property{AdminCount, DistinguishedName, DomainFQDN, DomainSID, Sensitive, HighValue, BlocksInheritance, IsACL, IsACLProtected, Enforced, Department, HasSPN, UnconstrainedDelegation, LastLogon, LastLogonTimestamp, IsPrimaryGroup, HasLAPS, DontRequirePreAuth, LogonType, HasURA, PasswordNeverExpires, PasswordNotRequired, FunctionalLevel, TrustType, SidFiltering, TrustedToAuth}
}
func ParseProperty(source string) (Property, error) {
	switch source {
	case "admincount":
		return AdminCount, nil
	case "distinguishedname":
		return DistinguishedName, nil
	case "domain":
		return DomainFQDN, nil
	case "domainsid":
		return DomainSID, nil
	case "sensitive":
		return Sensitive, nil
	case "highvalue":
		return HighValue, nil
	case "blocksinheritance":
		return BlocksInheritance, nil
	case "isacl":
		return IsACL, nil
	case "isaclprotected":
		return IsACLProtected, nil
	case "enforced":
		return Enforced, nil
	case "department":
		return Department, nil
	case "hasspn":
		return HasSPN, nil
	case "unconstraineddelegation":
		return UnconstrainedDelegation, nil
	case "lastlogon":
		return LastLogon, nil
	case "lastlogontimestamp":
		return LastLogonTimestamp, nil
	case "isprimarygroup":
		return IsPrimaryGroup, nil
	case "haslaps":
		return HasLAPS, nil
	case "dontreqpreauth":
		return DontRequirePreAuth, nil
	case "logontype":
		return LogonType, nil
	case "hasura":
		return HasURA, nil
	case "pwdneverexpires":
		return PasswordNeverExpires, nil
	case "passwordnotreqd":
		return PasswordNotRequired, nil
	case "functionallevel":
		return FunctionalLevel, nil
	case "trusttype":
		return TrustType, nil
	case "sidfiltering":
		return SidFiltering, nil
	case "trustedtoauth":
		return TrustedToAuth, nil
	default:
		return "", errors.New("Invalid enumeration value: " + source)
	}
}
func (s Property) String() string {
	switch s {
	case AdminCount:
		return string(AdminCount)
	case DistinguishedName:
		return string(DistinguishedName)
	case DomainFQDN:
		return string(DomainFQDN)
	case DomainSID:
		return string(DomainSID)
	case Sensitive:
		return string(Sensitive)
	case HighValue:
		return string(HighValue)
	case BlocksInheritance:
		return string(BlocksInheritance)
	case IsACL:
		return string(IsACL)
	case IsACLProtected:
		return string(IsACLProtected)
	case Enforced:
		return string(Enforced)
	case Department:
		return string(Department)
	case HasSPN:
		return string(HasSPN)
	case UnconstrainedDelegation:
		return string(UnconstrainedDelegation)
	case LastLogon:
		return string(LastLogon)
	case LastLogonTimestamp:
		return string(LastLogonTimestamp)
	case IsPrimaryGroup:
		return string(IsPrimaryGroup)
	case HasLAPS:
		return string(HasLAPS)
	case DontRequirePreAuth:
		return string(DontRequirePreAuth)
	case LogonType:
		return string(LogonType)
	case HasURA:
		return string(HasURA)
	case PasswordNeverExpires:
		return string(PasswordNeverExpires)
	case PasswordNotRequired:
		return string(PasswordNotRequired)
	case FunctionalLevel:
		return string(FunctionalLevel)
	case TrustType:
		return string(TrustType)
	case SidFiltering:
		return string(SidFiltering)
	case TrustedToAuth:
		return string(TrustedToAuth)
	default:
		panic("Invalid enumeration case: " + string(s))
	}
}
func (s Property) Name() string {
	switch s {
	case AdminCount:
		return "Admin Count"
	case DistinguishedName:
		return "Distinguished Name"
	case DomainFQDN:
		return "Domain FQDN"
	case DomainSID:
		return "Domain SID"
	case Sensitive:
		return "Sensitive"
	case HighValue:
		return "High Value"
	case BlocksInheritance:
		return "Blocks Inheritance"
	case IsACL:
		return "Is ACL"
	case IsACLProtected:
		return "ACL Inheritance Denied"
	case Enforced:
		return "Enforced"
	case Department:
		return "Department"
	case HasSPN:
		return "Has SPN"
	case UnconstrainedDelegation:
		return "Allows Unconstrained Delegation"
	case LastLogon:
		return "Last Logon"
	case LastLogonTimestamp:
		return "Last Logon (Replicated)"
	case IsPrimaryGroup:
		return "Is Primary Group"
	case HasLAPS:
		return "LAPS Enabled"
	case DontRequirePreAuth:
		return "Do Not Require Pre-Authentication"
	case LogonType:
		return "Logon Type"
	case HasURA:
		return "Has User Rights Assignment Collection"
	case PasswordNeverExpires:
		return "Password Never Expires"
	case PasswordNotRequired:
		return "Password Not Required"
	case FunctionalLevel:
		return "Functional Level"
	case TrustType:
		return "Trust Type"
	case SidFiltering:
		return "SID Filtering Enabled"
	case TrustedToAuth:
		return "Trusted For Constrained Delegation"
	default:
		panic("Invalid enumeration case: " + string(s))
	}
}
func (s Property) Is(others ...graph.Kind) bool {
	for _, other := range others {
		if value, err := ParseProperty(other.String()); err == nil && value == s {
			return true
		}
	}
	return false
}
func Nodes() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser}
}
func Relationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, GetChanges, GetChangesAll, GetChangesInFilteredSet, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, LocalToComputer, MemberOfLocalGroup, RemoteInteractiveLogonPrivilege, SyncLAPSPassword, WriteAccountRestrictions}
}
func ACLRelationships() []graph.Kind {
	return []graph.Kind{AllExtendedRights, ForceChangePassword, AddMember, AddAllowedToAct, GenericAll, WriteDACL, WriteOwner, GenericWrite, ReadLAPSPassword, ReadGMSAPassword, Owns, AddSelf, WriteSPN, AddKeyCredentialLink, GetChanges, GetChangesAll, GetChangesInFilteredSet, WriteAccountRestrictions, SyncLAPSPassword, DCSync}
}
func PathfindingRelationships() []graph.Kind {
	return []graph.Kind{Owns, GenericAll, GenericWrite, WriteOwner, WriteDACL, MemberOf, ForceChangePassword, AllExtendedRights, AddMember, HasSession, Contains, GPLink, AllowedToDelegate, TrustedBy, AllowedToAct, AdminTo, CanPSRemote, CanRDP, ExecuteDCOM, HasSIDHistory, AddSelf, DCSync, ReadLAPSPassword, ReadGMSAPassword, DumpSMSAPassword, SQLAdmin, AddAllowedToAct, WriteSPN, AddKeyCredentialLink, SyncLAPSPassword, WriteAccountRestrictions}
}
func IsACLKind(s graph.Kind) bool {
	for _, acl := range ACLRelationships() {
		if s == acl {
			return true
		}
	}
	return false
}
func NodeKinds() []graph.Kind {
	return []graph.Kind{Entity, User, Computer, Group, GPO, OU, Container, Domain, LocalGroup, LocalUser}
}
