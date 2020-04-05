package types

import (
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
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

var ErrNoError = errors.New("no error")

var syncStarted = false

var DefaultAccess = FullAccess
var QIP714BlockReached = false
var networkAdminRole string
var orgAdminRole string

var (
	OrgInfoMap  *OrgCache
	NodeInfoMap *NodeCache
	RoleInfoMap *RoleCache
	AcctInfoMap *AcctCache
)

type OrgKey struct {
	OrgId string
}

type OrgCache struct {
	c                 *lru.Cache
	mux               sync.Mutex
	evicted           bool
	populateCacheFunc func(orgId string) (*OrgInfo, error)
}

func (o *OrgCache) PopulateCacheFunc(cf func(string) (*OrgInfo, error)) {
	o.populateCacheFunc = cf
}

func NewOrgCache(cacheSize int) *OrgCache {
	orgCache := OrgCache{evicted: false}
	onEvictedFunc := func(k interface{}, v interface{}) {
		orgCache.evicted = true
	}
	orgCache.c, _ = lru.NewWithEvict(cacheSize, onEvictedFunc)
	return &orgCache
}

type RoleKey struct {
	OrgId  string
	RoleId string
}

type RoleCache struct {
	c                 *lru.Cache
	evicted           bool
	populateCacheFunc func(*RoleKey) (*RoleInfo, error)
}

func (r *RoleCache) PopulateCacheFunc(cf func(*RoleKey) (*RoleInfo, error)) {
	r.populateCacheFunc = cf
}

func NewRoleCache(cacheSize int) *RoleCache {
	roleCache := RoleCache{evicted: false}
	onEvictedFunc := func(k interface{}, v interface{}) {
		roleCache.evicted = true
	}
	roleCache.c, _ = lru.NewWithEvict(cacheSize, onEvictedFunc)
	return &roleCache
}

type NodeKey struct {
	OrgId string
	Url   string
}

type NodeCache struct {
	c                       *lru.Cache
	evicted                 bool
	populateCacheFunc       func(string) (*NodeInfo, error)
	populateAndValidateFunc func(string, string) bool
}

func (n *NodeCache) PopulateValidateFunc(cf func(string, string) bool) {
	n.populateAndValidateFunc = cf
}

func (n *NodeCache) PopulateCacheFunc(cf func(string) (*NodeInfo, error)) {
	n.populateCacheFunc = cf
}

func NewNodeCache(cacheSize int) *NodeCache {
	nodeCache := NodeCache{evicted: false}
	onEvictedFunc := func(k interface{}, v interface{}) {
		nodeCache.evicted = true

	}
	nodeCache.c, _ = lru.NewWithEvict(cacheSize, onEvictedFunc)
	return &nodeCache
}

type AccountKey struct {
	AcctId common.Address
}

type AcctCache struct {
	c                 *lru.Cache
	evicted           bool
	populateCacheFunc func(account common.Address) (*AccountInfo, error)
}

func (a *AcctCache) PopulateCacheFunc(cf func(common.Address) (*AccountInfo, error)) {
	a.populateCacheFunc = cf
}

func NewAcctCache(cacheSize int) *AcctCache {
	acctCache := AcctCache{evicted: false}
	onEvictedFunc := func(k interface{}, v interface{}) {
		acctCache.evicted = true
	}

	acctCache.c, _ = lru.NewWithEvict(cacheSize, onEvictedFunc)
	return &acctCache
}

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

func (o *OrgCache) UpsertOrgWithSubOrgList(orgRec *OrgInfo) {
	var key OrgKey
	if orgRec.ParentOrgId == "" {
		key = OrgKey{OrgId: orgRec.OrgId}
	} else {
		key = OrgKey{OrgId: orgRec.ParentOrgId + "." + orgRec.OrgId}
	}
	orgRec.FullOrgId = key.OrgId
	o.c.Add(key, orgRec)
}

func containsKey(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (o *OrgCache) GetOrg(orgId string) (*OrgInfo, error) {
	key := OrgKey{OrgId: orgId}
	if ent, ok := o.c.Get(key); ok {
		return ent.(*OrgInfo), ErrNoError
	}
	// check if the org cache is evicted. if yes we need
	// fetch the record from the contract
	if o.evicted {
		// call cache population function to populate from contract
		orgRec, err := o.populateCacheFunc(orgId)
		if err != ErrNoError {
			return nil, err
		}
		if orgRec == nil {
			return nil, err
		}
		// insert the received record into cache
		o.UpsertOrgWithSubOrgList(orgRec)
		//return the record
		return orgRec, err
	}
	return nil, ErrNoError
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

func (n *NodeCache) GetNodeByUrl(url string) (*NodeInfo, error) {
	for _, k := range n.c.Keys() {
		ent := k.(NodeKey)
		if ent.Url == url {
			v, _ := n.c.Get(ent)
			return v.(*NodeInfo), ErrNoError
		}
	}
	// check if the node cache is evicted. if yes we need
	// fetch the record from the contract
	if n.evicted {

		// call cache population function to populate from contract
		nodeRec, err := n.populateCacheFunc(url)
		if err != nil {
			return nil, err
		}

		if nodeRec == nil {
			return nil, err
		}

		// insert the received record into cache
		n.UpsertNode(nodeRec.OrgId, nodeRec.Url, nodeRec.Status)
		//return the record
		return nodeRec, err
	}
	return nil, ErrNoError
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

func (a *AcctCache) GetAccount(acct common.Address) (*AccountInfo, error) {
	if v, ok := a.c.Get(AccountKey{acct}); ok {
		return v.(*AccountInfo), ErrNoError
	}

	// check if the account cache is evicted. if yes we need
	// fetch the record from the contract
	if a.evicted {
		// call function to populate cache with the record
		acctRec, err := a.populateCacheFunc(acct)
		// insert the received record into cache
		if err != ErrNoError {
			return nil, err
		}
		if acctRec == nil {
			return nil, ErrNoError
		}
		a.UpsertAccount(acctRec.OrgId, acctRec.RoleId, acctRec.AcctId, acctRec.IsOrgAdmin, acctRec.Status)
		//return the record
		return acctRec, ErrNoError
	}
	return nil, nil
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

		orgRec, err := OrgInfoMap.GetOrg(vp.OrgId)
		if err != ErrNoError {
			return nil
		}

		if vp.RoleId == roleId && (vp.OrgId == orgId || (orgRec != nil && orgRec.UltimateParent == orgId)) {
			alist = append(alist, *vp)
		}
	}
	return alist
}

func (r *RoleCache) UpsertRole(orgId string, role string, voter bool, admin bool, access AccessType, active bool) {
	key := RoleKey{orgId, role}
	r.c.Add(key, &RoleInfo{orgId, role, voter, admin, access, active})

}

func (r *RoleCache) GetRole(orgId string, roleId string) (*RoleInfo, error) {
	key := RoleKey{OrgId: orgId, RoleId: roleId}
	if ent, ok := r.c.Get(key); ok {
		return ent.(*RoleInfo), ErrNoError
	}
	// check if the role cache is evicted. if yes we need
	// fetch the record from the contract
	if r.evicted {
		// call cache population function to populate from contract
		roleRec, err := r.populateCacheFunc(&RoleKey{RoleId: roleId, OrgId: orgId})
		if err != ErrNoError {
			return nil, err
		}
		if roleRec == nil {
			return nil, err
		}
		// insert the received record into cache
		r.UpsertRole(roleRec.OrgId, roleRec.RoleId, roleRec.IsVoter, roleRec.IsAdmin, roleRec.Access, roleRec.Active)

		//return the record
		return roleRec, err
	}
	return nil, ErrNoError
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
	a, _ := AcctInfoMap.GetAccount(acctId)
	if a != nil && a.Status == AcctActive {
		// get the org details and ultimate org details. check org status
		// if the org is not approved or pending suspension
		o, _ := OrgInfoMap.GetOrg(a.OrgId)
		if o != nil && (o.Status == OrgApproved || o.Status == OrgPendingSuspension) {
			u, _ := OrgInfoMap.GetOrg(o.UltimateParent)
			if u != nil && (u.Status == OrgApproved || u.Status == OrgPendingSuspension) {
				if a.RoleId == networkAdminRole || a.RoleId == orgAdminRole {
					return FullAccess
				}
				r, _ := RoleInfoMap.GetRole(a.OrgId, a.RoleId)
				if r != nil && r.Active {
					return r.Access
				}
				r, _ = RoleInfoMap.GetRole(o.UltimateParent, a.RoleId)
				if r != nil && r.Active {
					return r.Access
				}
			}
		}
	}
	return DefaultAccess
}

// validates if the account can transact from the current node
func ValidateNodeForTxn(hexnodeId string, from common.Address) bool {
	if !QIP714BlockReached || hexnodeId == "" {
		return true
	}

	passedEnodeId, err := enode.ParseV4(hexnodeId)
	if err != nil {
		return false
	}

	ac, _ := AcctInfoMap.GetAccount(from)
	if ac == nil {
		return true
	}

	acOrgRec, err := OrgInfoMap.GetOrg(ac.OrgId)
	if err != ErrNoError || acOrgRec == nil {
		return false
	}

	// scan through the node list and validate
	for _, n := range NodeInfoMap.GetNodeList() {
		orgRec, err := OrgInfoMap.GetOrg(n.OrgId)
		if err != ErrNoError || orgRec == nil {
			return false
		}
		if orgRec.UltimateParent == acOrgRec.UltimateParent {
			recEnodeId, _ := enode.ParseV4(n.Url)
			if recEnodeId.ID() == passedEnodeId.ID() && n.Status == NodeApproved {
				return true
			}
		}
	}
	if NodeInfoMap.evicted {
		return NodeInfoMap.populateAndValidateFunc(hexnodeId, acOrgRec.UltimateParent)
	}

	return false
}
