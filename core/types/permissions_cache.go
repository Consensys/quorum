package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/golang-lru"
	"sync"
)

type AccessType uint8

const (
	ReadOnly AccessType = iota
	Transact
	ContractDeploy
	FullAccess
)

type OrgStatus uint8

const (
	OrgProposed OrgStatus = iota + 1
	OrgApproved
	OrgPendingSuspension
	OrgSuspended
	OrgRevokeSuspension
)

type OrgInfo struct {
	OrgId  string
	Status OrgStatus
}

type NodeStatus uint8

const (
	NodePendingApproval NodeStatus = iota + 1
	NodeApproved
	NodeDeactivated
	Blacklisted
)

type AcctStatus uint8

const (
	AcctPendingApproval AcctStatus = iota + 1
	AcctActive
	AcctInactive
)

type NodeInfo struct {
	OrgId  string
	Url    string
	Status NodeStatus
}

type RoleInfo struct {
	OrgId   string
	RoleId  string
	IsVoter bool
	Access  AccessType
	Active  bool
}

type AccountInfo struct {
	OrgId      string
	RoleId     string
	AcctId     common.Address
	IsOrgAdmin bool
	Status     AcctStatus
}

type OrgStruct struct {
	OrgId string
	Keys  []string
}

// permission config for bootstrapping
type PermissionConfig struct {
	UpgrdAddress   string
	InterfAddress  string
	ImplAddress    string
	NodeAddress    string
	AccountAddress string
	RoleAddress    string
	VoterAddress   string
	OrgAddress     string
	NwAdminOrg     string
	NwAdminRole    string
	OrgAdminRole   string

	Accounts []string //initial list of account that need full access
}

type OrgKey struct {
	OrgId string
}

type NodeKey struct {
	OrgId string
	Url   string
}

type RoleKey struct {
	OrgId  string
	RoleId string
}

type AccountKey struct {
	OrgId  string
	RoleId string
	AcctId common.Address
}

type OrgCache struct {
	c   *lru.Cache
	mux sync.Mutex
}

type NodeCache struct {
	c   *lru.Cache
	mux sync.Mutex
}

type RoleCache struct {
	c   *lru.Cache
	mux sync.Mutex
}

type AcctCache struct {
	c   *lru.Cache
	mux sync.Mutex
}

func NewOrgCache() *OrgCache {
	c, _ := lru.New(defaultMapLimit)
	return &OrgCache{c, sync.Mutex{}}
}

func NewNodeCache() *NodeCache {
	c, _ := lru.New(defaultMapLimit)
	return &NodeCache{c, sync.Mutex{}}
}

func NewRoleCache() *RoleCache {
	c, _ := lru.New(defaultMapLimit)
	return &RoleCache{c, sync.Mutex{}}
}

func NewAcctCache() *AcctCache {
	c, _ := lru.New(defaultMapLimit)
	return &AcctCache{c, sync.Mutex{}}
}

var DefaultAccess = FullAccess

const orgKeyMapLimit = 100

const defaultMapLimit = 100

var OrgKeyMap, _ = lru.New(orgKeyMapLimit)

var OrgInfoMap = NewOrgCache()
var NodeInfoMap = NewNodeCache()
var RoleInfoMap = NewRoleCache()
var AcctInfoMap = NewAcctCache()

var orgKeyLock sync.Mutex

func (pc *PermissionConfig) IsEmpty() bool {
	return pc.InterfAddress == "" || pc.NodeAddress == "" || pc.AccountAddress == ""
}

// sets default access to ReadOnly
func SetDefaultAccess() {
	DefaultAccess = ReadOnly
}

func (o *OrgCache) UpsertOrg(orgId string, status OrgStatus) {
	defer o.mux.Unlock()
	o.mux.Lock()
	key := OrgKey{OrgId: orgId}
	if _, ok := o.c.Get(key); ok {
		log.Info("AJ-OrgId already exists. update it", "orgId", orgId)
		o.c.Add(key, &OrgInfo{orgId, status})
	} else {
		log.Info("AJ-OrgId does not exist. add it", "orgId", orgId)
		o.c.Add(key, &OrgInfo{orgId, status})
	}
}

func (o *OrgCache) GetOrg(orgId string) *OrgInfo {
	defer o.mux.Unlock()
	o.mux.Lock()
	key := OrgKey{OrgId: orgId}
	if ent, ok := o.c.Get(key); ok {
		log.Info("AJ-OrgFound", "orgId", orgId)
		return ent.(*OrgInfo)
	}
	return nil
}

func (o *OrgCache) Show() {
	for i, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		log.Info("AJ-Org", "i", i, "key", k, "value", v)
	}
}
func (o *OrgCache) GetOrgList() []OrgInfo {
	var olist []OrgInfo
	for _, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		vp := v.(*OrgInfo)
		olist = append(olist, *vp)
	}
	return olist
}

func (n *NodeCache) UpsertNode(orgId string, url string, status NodeStatus) {
	defer n.mux.Unlock()
	n.mux.Lock()
	key := NodeKey{OrgId: orgId, Url: url}
	if _, ok := n.c.Get(key); ok {
		log.Info("AJ-Node already exists. update it", "orgId", orgId, "url", url)
		n.c.Add(key, &NodeInfo{orgId, url, status})
	} else {
		log.Info("AJ-Node does not exist. add it", "orgId", orgId, "url", url)
		n.c.Add(key, &NodeInfo{orgId, url, status})
	}
}

func (n *NodeCache) GetNodeByUrl(url string) *NodeInfo {
	defer n.mux.Unlock()
	n.mux.Lock()
	var key NodeKey
	var found = false
	for _, k := range n.c.Keys() {
		ent := k.(NodeKey)
		if ent.Url == url {
			key = ent
			found = true
			break
		}
	}
	if found {
		v, _ := n.c.Get(key)
		ent := v.(*NodeInfo)
		log.Info("AJ-NodeFound", "url", ent.Url, "orgId", ent.OrgId)
		return ent
	}
	return nil
}

func (o *NodeCache) Show() {
	for i, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		log.Info("AJ-Node", "i", i, "key", k, "value", v)
	}
}

func (o *NodeCache) GetNodeList() []NodeInfo {
	var olist []NodeInfo
	for _, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		vp := v.(*NodeInfo)
		olist = append(olist, *vp)
	}
	return olist
}

func (a *AcctCache) UpsertAccount(orgId string, role string, acct common.Address, orgAdmin bool, status AcctStatus) {
	defer a.mux.Unlock()
	a.mux.Lock()
	key := AccountKey{orgId, role, acct}
	if _, ok := a.c.Get(key); ok {
		log.Info("AJ-account already exists. update it", "orgId", orgId, "role", role, "acct", acct)
		a.c.Add(key, &AccountInfo{orgId, role, acct, orgAdmin, status})
	} else {
		log.Info("AJ-account does not exist. add it", "orgId", orgId, "role", role, "acct", acct)
		a.c.Add(key, &AccountInfo{orgId, role, acct, orgAdmin, status})
	}
}

func (a *AcctCache) GetAccountByAccount(acct common.Address) *AccountInfo {
	defer a.mux.Unlock()
	a.mux.Lock()
	var key AccountKey
	var found = false
	for _, k := range a.c.Keys() {
		ent := k.(AccountKey)
		if ent.AcctId == acct {
			key = ent
			found = true
			break
		}
	}
	if found {
		v, _ := a.c.Get(key)
		ent := v.(*AccountInfo)
		log.Info("AJ-AccountFound", "org", ent.OrgId, "role", ent.RoleId, "acct", ent.AcctId)
		return ent
	}
	return nil
}

func (o *AcctCache) Show() {
	for i, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		log.Info("AJ-Accounts", "i", i, "key", k, "value", v)
	}
}

func (o *AcctCache) GetAcctList() []AccountInfo {
	var olist []AccountInfo
	for _, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		vp := v.(*AccountInfo)
		olist = append(olist, *vp)
	}
	return olist
}

func (r *RoleCache) UpsertRole(orgId string, role string, voter bool, access AccessType, active bool) {
	defer r.mux.Unlock()
	r.mux.Lock()
	key := RoleKey{orgId, role}
	if _, ok := r.c.Get(key); ok {
		log.Info("AJ-role already exists. update it", "orgId", orgId, "role", role, "access", access, "voter", voter, "active", active)
		r.c.Add(key, &RoleInfo{orgId, role, voter, access, active})
	} else {
		log.Info("AJ-role does not exist. add it", "orgId", orgId, "role", role, "access", access, "voter", voter, "active", active)
		r.c.Add(key, &RoleInfo{orgId, role, voter, access, active})
	}
}

func (r *RoleCache) GetRole(orgId string, roleId string) *RoleInfo {
	defer r.mux.Unlock()
	r.mux.Lock()
	key := RoleKey{OrgId: orgId, RoleId: roleId}
	if ent, ok := r.c.Get(key); ok {
		log.Info("AJ-RoleFound", "orgId", orgId, "roleId", roleId)
		return ent.(*RoleInfo)
	}
	return nil
}

func (r *RoleCache) Show() {
	for i, k := range r.c.Keys() {
		v, _ := r.c.Get(k)
		log.Info("AJ-Role", "i", i, "key", k, "value", v)
	}
}

func (o *RoleCache) GetRoleList() []RoleInfo {
	var olist []RoleInfo
	for _, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		vp := v.(*RoleInfo)
		olist = append(olist, *vp)
	}
	return olist
}

// Returns the access type for an account. If not found returns
// default access
func GetAcctAccess(acctId common.Address) AccessType {
	log.Info("AJ-get acct access ", "acct", acctId)
	if a := AcctInfoMap.GetAccountByAccount(acctId); a != nil {
		log.Info("AJ-Acct found", "a", a)
		o := OrgInfoMap.GetOrg(a.OrgId)
		r := RoleInfoMap.GetRole(a.OrgId, a.RoleId)
		if o != nil && r != nil {
			log.Info("AJ-org role found")
			if (o.Status == OrgApproved || o.Status == OrgRevokeSuspension) && r.Active {
				log.Info("AJ-access found", "access", r.Access)
				return r.Access
			} else {
				log.Info("AJ-access org or role invalid")
			}
		} else {
			log.Info("AJ-access org or role is missing")
		}
	} else {
		log.Info("AJ-Acct not found", "def access", DefaultAccess)
	}
	return DefaultAccess
}

// Adds org key details to cache
func AddOrgKey(orgId string, key string) {
	if OrgKeyMap.Len() != 0 {
		if val, ok := OrgKeyMap.Get(orgId); ok {
			orgKeyLock.Lock()
			defer orgKeyLock.Unlock()
			// Org record exists. Append the key only
			vo := val.(*OrgStruct)
			vo.Keys = append(vo.Keys, key)
			return
		}
	}
	OrgKeyMap.Add(orgId, &OrgStruct{OrgId: orgId, Keys: []string{key}})
}

// deletes org key details from cache
func DeleteOrgKey(orgId string, key string) {
	if val, ok := OrgKeyMap.Get(orgId); ok {
		orgKeyLock.Lock()
		defer orgKeyLock.Unlock()
		vo := val.(*OrgStruct)
		for i, keyVal := range vo.Keys {
			if keyVal == key {
				vo.Keys = append(vo.Keys[:i], vo.Keys[i+1:]...)
				break
			}
		}
	}
}

// Givens a orgid returns the linked keys for the org
func ResolvePrivateForKeys(orgId string) []string {
	var keys []string
	if val, ok := OrgKeyMap.Get(orgId); ok {
		vo := val.(*OrgStruct)
		if len(vo.Keys) > 0 {
			keys = vo.Keys
		} else {
			keys = append(keys, orgId)
		}
		return keys
	}
	keys = append(keys, orgId)
	return keys
}
