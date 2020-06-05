package types

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
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
	NodeBlackListed
	NodeRecoveryInitiated
)

type AcctStatus uint8

const (
	AcctPendingApproval AcctStatus = iota + 1
	AcctActive
	AcctInactive
	AcctSuspended
	AcctBlacklisted
	AdminRevoked
	AcctRecoveryInitiated
	AcctRecoveryCompleted
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
	IsAdmin bool       `json:"isAdmin"`
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

// permission config for bootstrapping
type PermissionConfig struct {
	UpgrdAddress   common.Address `json:"upgrdableAddress"`
	InterfAddress  common.Address `json:"interfaceAddress"`
	ImplAddress    common.Address `json:"implAddress"`
	NodeAddress    common.Address `json:"nodeMgrAddress"`
	AccountAddress common.Address `json:"accountMgrAddress"`
	RoleAddress    common.Address `json:"roleMgrAddress"`
	VoterAddress   common.Address `json:"voterMgrAddress"`
	OrgAddress     common.Address `json:"orgMgrAddress"`
	NwAdminOrg     string         `json:"nwAdminOrg"`
	NwAdminRole    string         `json:"nwAdminRole"`
	OrgAdminRole   string         `json:"orgAdminRole"`

	Accounts      []common.Address `json:"accounts"` //initial list of account that need full access
	SubOrgDepth   *big.Int         `json:"subOrgDepth"`
	SubOrgBreadth *big.Int         `json:"subOrgBreadth"`
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
	c *lru.Cache
}

type RoleCache struct {
	c *lru.Cache
}

type AcctCache struct {
	c *lru.Cache
}

func NewOrgCache() *OrgCache {
	c, _ := lru.New(defaultOrgMapLimit)
	return &OrgCache{c, sync.Mutex{}}
}

func NewNodeCache() *NodeCache {
	c, _ := lru.New(defaultNodeMapLimit)
	return &NodeCache{c}
}

func NewRoleCache() *RoleCache {
	c, _ := lru.New(defaultRoleMapLimit)
	return &RoleCache{c}
}

func NewAcctCache() *AcctCache {
	c, _ := lru.New(defaultAccountMapLimit)
	return &AcctCache{c}
}

var syncStarted = false

var DefaultAccess = FullAccess
var QIP714BlockReached = false
var networkAdminRole string
var orgAdminRole string

const defaultOrgMapLimit = 2000
const defaultRoleMapLimit = 2500
const defaultNodeMapLimit = 1000
const defaultAccountMapLimit = 6000

var OrgInfoMap = NewOrgCache()
var NodeInfoMap = NewNodeCache()
var RoleInfoMap = NewRoleCache()
var AcctInfoMap = NewAcctCache()

func (pc *PermissionConfig) IsEmpty() bool {
	return pc.InterfAddress == common.HexToAddress("0x0")
}

func SetSyncStatus() {
	syncStarted = true
}

func GetSyncStatus() bool {
	return syncStarted
}

// sets the default access to Readonly upon QIP714Blokc
func SetDefaultAccess() {
	DefaultAccess = ReadOnly
	QIP714BlockReached = true
}

// sets default access to readonly and initializes the values for
// network admin role and org admin role
func SetDefaults(nwRoleId, oaRoleId string) {
	networkAdminRole = nwRoleId
	orgAdminRole = oaRoleId
}

func GetDefaults() (string, string, AccessType) {
	return networkAdminRole, orgAdminRole, DefaultAccess
}

func GetNodeUrl(enodeId string, ip string, port uint16, raftport uint16) string {
	return fmt.Sprintf("enode://%s@%s:%d?raftport=%d&discport=0", enodeId, strings.Trim(ip, "\x00"), port, raftport)
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
	olist := make([]OrgInfo, len(o.c.Keys()))
	for i, k := range o.c.Keys() {
		v, _ := o.c.Get(k)
		vp := v.(*OrgInfo)
		olist[i] = *vp
	}
	return olist
}

func (n *NodeCache) UpsertNode(orgId string, url string, status NodeStatus) {
	key := NodeKey{OrgId: orgId, Url: url}
	n.c.Add(key, &NodeInfo{orgId, url, status})
}

func (n *NodeCache) GetNodeByUrl(url string) *NodeInfo {
	for _, k := range n.c.Keys() {
		ent := k.(NodeKey)
		if ent.Url == url {
			v, _ := n.c.Get(ent)
			return v.(*NodeInfo)
		}
	}
	return nil
}

func (n *NodeCache) GetNodeList() []NodeInfo {
	olist := make([]NodeInfo, len(n.c.Keys()))
	for i, k := range n.c.Keys() {
		v, _ := n.c.Get(k)
		vp := v.(*NodeInfo)
		olist[i] = *vp
	}
	return olist
}

func (a *AcctCache) UpsertAccount(orgId string, role string, acct common.Address, orgAdmin bool, status AcctStatus) {
	key := AccountKey{acct}
	a.c.Add(key, &AccountInfo{orgId, role, acct, orgAdmin, status})
}

func (a *AcctCache) GetAccount(acct common.Address) *AccountInfo {
	if v, ok := a.c.Get(AccountKey{acct}); ok {
		return v.(*AccountInfo)
	}
	return nil
}

func (a *AcctCache) GetAcctList() []AccountInfo {
	alist := make([]AccountInfo, len(a.c.Keys()))
	for i, k := range a.c.Keys() {
		v, _ := a.c.Get(k)
		vp := v.(*AccountInfo)
		alist[i] = *vp
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

		if vp.RoleId == roleId && (vp.OrgId == orgId || OrgInfoMap.GetOrg(vp.OrgId).UltimateParent == orgId) {
			alist = append(alist, *vp)
		}
	}
	return alist
}

func (r *RoleCache) UpsertRole(orgId string, role string, voter bool, admin bool, access AccessType, active bool) {
	key := RoleKey{orgId, role}
	r.c.Add(key, &RoleInfo{orgId, role, voter, admin, access, active})

}

func (r *RoleCache) GetRole(orgId string, roleId string) *RoleInfo {
	key := RoleKey{OrgId: orgId, RoleId: roleId}
	if ent, ok := r.c.Get(key); ok {
		return ent.(*RoleInfo)
	}
	return nil
}

func (r *RoleCache) GetRoleList() []RoleInfo {
	rlist := make([]RoleInfo, len(r.c.Keys()))
	for i, k := range r.c.Keys() {
		v, _ := r.c.Get(k)
		vp := v.(*RoleInfo)
		rlist[i] = *vp
	}
	return rlist
}

// Returns the access type for an account. If not found returns
// default access
func GetAcctAccess(acctId common.Address) AccessType {
	//if we have not reached QIP714 block return default access
	//which will be full access
	if !QIP714BlockReached {
		return DefaultAccess
	}

	// check if the org status is fine to do the transaction
	a := AcctInfoMap.GetAccount(acctId)
	if a != nil && a.Status == AcctActive {
		// get the org details and ultimate org details. check org status
		// if the org is not approved or pending suspension
		o := OrgInfoMap.GetOrg(a.OrgId)
		if o != nil && (o.Status == OrgApproved || o.Status == OrgPendingSuspension) {
			u := OrgInfoMap.GetOrg(o.UltimateParent)
			if u != nil && (u.Status == OrgApproved || u.Status == OrgPendingSuspension) {
				if a.RoleId == networkAdminRole || a.RoleId == orgAdminRole {
					return FullAccess
				}
				if r := RoleInfoMap.GetRole(a.OrgId, a.RoleId); r != nil && r.Active {
					return r.Access
				}
				if r := RoleInfoMap.GetRole(o.UltimateParent, a.RoleId); r != nil && r.Active {
					return r.Access
				}
			}
		}
	}
	return DefaultAccess
}

func ValidateNodeForTxn(hexnodeId string, from common.Address) bool {
	if !QIP714BlockReached || hexnodeId == "" {
		return true
	}

	passedEnodeId, err := enode.ParseV4(hexnodeId)
	if err != nil {
		return false
	}

	ac := AcctInfoMap.GetAccount(from)
	if ac == nil {
		return true
	}

	ultimateParent := OrgInfoMap.GetOrg(ac.OrgId).UltimateParent
	// scan through the node list and validate
	for _, n := range NodeInfoMap.GetNodeList() {
		if OrgInfoMap.GetOrg(n.OrgId).UltimateParent == ultimateParent {
			recEnodeId, _ := enode.ParseV4(n.Url)
			if recEnodeId.ID() == passedEnodeId.ID() {
				return true
			}
		}
	}
	return false
}
