# External Signer

Quorum has made enhancement on Ethereum's Clef and extended its support of private transaction signing.

## How to Use Clef

#### Step 1: Start Clef

```shell script
> clef -keystore ./keys -chainid 10

WARNING!

Clef is an account management tool. It may, like any software, contain bugs.

Please take care to
- backup your keystore files,
- verify that the keystore(s) can be opened with your password.

Clef is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR
PURPOSE. See the GNU General Public License for more details.
```

#### Step 2: Start Quorum Geth with `--signer` flag

#### Step 3: Send private transaction from Quorum console and approve on Clef

```shell script
--------- Transaction request-------------
to:    <contact creation>
from:     0xed9d02e382b34818e88B88a309c7fe71E65f419d [chksum ok]
value:    0 wei
gas:      0x47b760 (4700000)
gasprice: 0 wei
nonce:    0x70 (112)
data:     0xcc05ff79f3e9735cc3d60d72b8a35658eaf7856a91e216a95819bc26d9b463399d99eb23ea74b0603deff44589046d1e977b0b847e0406a013adb7dc6664013b

Transaction validation:
  * Info : Quorum private transaction. Key: 0xcc05ff79f3e9735cc3d60d72b8a35658eaf7856a91e216a95819bc26d9b463399d99eb23ea74b0603deff44589046d1e977b0b847e0406a013adb7dc6664013b


Request context:
	NA -> NA -> NA

Additional HTTP header data, provided by the external caller:
	User-Agent:
	Origin:
-------------------------------------------
Approve? [y/N]:
> y
## Account password

Please enter the password for account 0xed9d02e382b34818e88B88a309c7fe71E65f419d
> -----------------------
Transaction signed:
 {
    "nonce": "0x70",
    "gasPrice": "0x0",
    "gas": "0x47b760",
    "to": null,
    "value": "0x0",
    "input": "0xcc05ff79f3e9735cc3d60d72b8a35658eaf7856a91e216a95819bc26d9b463399d99eb23ea74b0603deff44589046d1e977b0b847e0406a013adb7dc6664013b",
    "v": "0x25",
    "r": "0xe2bb99d9697a8fec92486b51ebf91b5a012b10a0bc3c67ff7bdb2bcd9a50b028",
    "s": "0x667e65da1b712f4cc0c83010b2f25d0203e259440281f90136e93eaeb37091b7",
    "hash": "0x6d01b5c45ed1259cee62e5f9c3ef68be9fcfb98b74279777609319c5f269c0ae"
  }
DEBUG[03-17|16:52:28.814] Served account_signTransaction           reqid=4 t=6.941322354s
```

For more information on Clef as external signer, please refer to 
[Ethereum's Clef](https://github.com/ethereum/go-ethereum/tree/master/cmd/clef)