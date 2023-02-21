// This package contains logic for multi-tenancy feature.
// It encapsulates how to isAuthorized a PSI of a tenant and the scope design specification
// to be used by Operator when setting up access token scope for a client.
//
// # Query param `node.eoa` and `self.eoa` can be multiple
//
// Scope examples:
//   - Wild card
//     `psi://MY_PSI?node.eoa=0x0`: any node-managed EOA
//     `psi://MY_PSI?self.eoa=0x0`: any self-managed EOA
//     `psi://MY_PSI?node.eoa=0x0&self.node=0x0`: any node-managed EOA
//   - Specific:
//     `psi://MY_PSI?node.eoa=0xdf08aad9d60f2227fdaed44dffd22753faf3d676`
//     `psi://MY_PSI?self.eoa=0x1234aad9d60f2227fdaed44dffd22753faf3d676`
package multitenancy
