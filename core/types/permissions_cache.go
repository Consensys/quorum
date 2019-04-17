package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/golang-lru"
	"math/big"
	"strings"
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
	OrgPendingApproval OrgStatus = iota + 1
	OrgApproved
	OrgPendingSuspension
	OrgSuspended
	OrgRevokeSuspension
)

type OrgInfo struct {
	OrgId          string    `json:"orgId"`
	FullOrgId      string    `json:"fullOrgId"`
	ParentOrgId    string    `json:"parentOrgId"`
	UltimateParent string    `json:"ultimateParent"`
	Level          *big.Int  `json:"level"`
	SubOrgList     []string  `json:"subOrgList"`
	Status         OrgStatus `json:"status"`
}

type NodeStatus uint8

const (
	NodePendingApproval NodeStatus = iota + 1
	NodeApproved
	NodeDeactivated
	NodeActivated
	NodeBlackListed
)

type AcctStatus uint8

const (
	AcctPendingApproval AcctStatus = iota + 1
	AcctActive
	AcctInactive
	AcctSuspended
	AcctBlacklisted
)

type NodeInfo struct {
	OrgId  string     `json:"orgId"`
	Url    string     `json:"url"`
	Status NodeStatus `json:"status"`
}

type RoleInfo struct {
	OrgId   string     `json:"orgId"`
	RoleId  string     `json:"roleId"`
	IsVoter bool       `json:"isVoter"`
	Access  AccessType `json:"access"`
	Active  bool       `json:"active"`
}

type AccountInfo struct {
	OrgId      string         `json:"orgId"`
	RoleId     string         `json:"roleId"`
	AcctId     common.Address `json:"acctId"`
	IsOrgAdmin bool           `json:"isOrgAdmin"`
	Status     AcctStatus     `json:"status"`
}

type OrgDetailInfo struct {
	NodeList   []NodeInfo    `json:"nodeList"`
	RoleList   []RoleInfo    `json:"roleList"`
	AcctList   []AccountInfo `json:"acctList"`
	SubOrgList []string      `json:"subOrgList"`
}

type OrgStruct struct {
	OrgId string
	Keys  []string
}

// permission config for bootstrapping
type PermissionConfig struct {
	UpgrdAddress   common.Address
	InterfAddress  common.Address
	ImplAddress    common.Address
	NodeAddress    common.Address
	AccountAddress common.Address
	RoleAddress    common.Address
	VoterAddress   common.Address
	OrgAddress     common.Address
	NwAdminOrg     string
	NwAdminRole    string
	OrgAdminRole   string

	Accounts      []common.Address //initial list of account that need full access
	SubOrgDepth   big.Int
	SubOrgBreadth big.Int
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
	return pc.InterfAddress == common.HexToAddress("0x0") || pc.NodeAddress == common.HexToAddress("0x0") || pc.AccountAddress == common.HexToAddress("0x0")
}

// sets default access to ReadOnly
func SetDefaultAccess() {
	DefaultAccess = ReadOnly
}

func (o *OrgCache) UpsertOrg(orgId, parentOrg, ultimateParent string, level *big.Int, status OrgStatus) {
	defer o.mux.Unlock()
	o.mux.Lock()
	var key OrgKey
	if parentOrg == "" {
		key = OrgKey{OrgId: orgId}
	} else {
		key = OrgKey{OrgId: parentOrg + "." + orgId}
		pkey := OrgKey{OrgId: parentOrg}
		if ent, ok := o.c.Get(pkey); ok {
			porg := ent.(*OrgInfo)
			if !containsKey(porg.SubOrgList, key.OrgId) {
				porg.SubOrgList = append(porg.SubOrgList, key.OrgId)
				o.c.Add(pkey, porg)
			}
		}
	}

	norg := &OrgInfo{orgId, key.OrgId, parentOrg, ultimateParent, level, nil, status}
	o.c.Add(key, norg)
}

func containsKey(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (o *OrgCache) GetOrg(orgId string) *OrgInfo {
	defer o.mux.Unlock()
	o.mux.Lock()
	key := OrgKey{OrgId: orgId}
	if ent, ok := o.c.Get(key); ok {
		return ent.(*OrgInfo)
	}
	return nil
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
	n.c.Add(key, &NodeInfo{orgId, url, status})
}

func (n *NodeCache) GetNodeByUrl(url string) *NodeInfo {
	defer n.mux.Unlock()
	n.mux.Lock()
	for _, k := range n.c.Keys() {
		ent := k.(NodeKey)
		if ent.Url == url {
			v, _ := n.c.Get(ent)
			return v.(*NodeInfo)
		}
	}
	return nil
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
	key := AccountKey{acct}
	a.c.Add(key, &AccountInfo{orgId, role, acct, orgAdmin, status})
}

func (a *AcctCache) GetAccount(acct common.Address) *AccountInfo {
	defer a.mux.Unlock()
	a.mux.Lock()
	if v, ok := a.c.Get(AccountKey{acct}); ok {
		return v.(*AccountInfo)
	}
	return nil
}

func (a *AcctCache) GetAcctList() []AccountInfo {
	var alist []AccountInfo
	for _, k := range a.c.Keys() {
		v, _ := a.c.Get(k)
		vp := v.(*AccountInfo)
		alist = append(alist, *vp)
	}
	return alist
}

func (a *AcctCache) GetAcctListOrg(orgId string) []AccountInfo {
	var alist []AccountInfo
	for _, k := range a.c.Keys() {
		v, _ := a.c.Get(k)
		vp := v.(*AccountInfo)
		if vp.OrgId == orgId {
			alist = append(alist, *vp)
		}
	}
	return alist
}

func (a *AcctCache) GetAcctListRole(orgId, roleId string) []AccountInfo {
	var alist []AccountInfo
	for _, k := range a.c.Keys() {
		v, _ := a.c.Get(k)
		vp := v.(*AccountInfo)
		if vp.OrgId == orgId && vp.RoleId == roleId {
			alist = append(alist, *vp)
		}
	}
	return alist
}

func (r *RoleCache) UpsertRole(orgId string, role string, voter bool, access AccessType, active bool) {
	defer r.mux.Unlock()
	r.mux.Lock()
	key := RoleKey{orgId, role}
	r.c.Add(key, &RoleInfo{orgId, role, voter, access, active})

}

func (r *RoleCache) GetRole(orgId string, roleId string) *RoleInfo {
	defer r.mux.Unlock()
	r.mux.Lock()
	key := RoleKey{OrgId: orgId, RoleId: roleId}
	if ent, ok := r.c.Get(key); ok {
		return ent.(*RoleInfo)
	}
	return nil
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
	if a := AcctInfoMap.GetAccount(acctId); a != nil && a.Status == AcctActive {
		o := OrgInfoMap.GetOrg(a.OrgId)
		r := RoleInfoMap.GetRole(a.OrgId, a.RoleId)
		if o != nil && r != nil {
			if o.Status == OrgApproved && r.Active {
				return r.Access
			}
		}
	}
	return DefaultAccess
}

func ValidateNodeForTxn(enodeId string, from common.Address) bool {
	if enodeId == "" {
		return true
	}
	ac := AcctInfoMap.GetAccount(from)
	if ac == nil {
		return true
	}
	ultimateParent := OrgInfoMap.GetOrg(ac.OrgId).UltimateParent
	// scan through the node list and validate
	for _, n := range NodeInfoMap.GetNodeList() {
		if OrgInfoMap.GetOrg(n.OrgId).UltimateParent == ultimateParent && strings.Contains(n.Url, enodeId) {
			return true
		}
	}
	return false
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
