package simplepath

import (
	"github.com/payshares/go/psrkey"
	"github.com/payshares/go/xdr"
)

func makeAsset(typ xdr.AssetType, code string, issuer string) xdr.Asset {

	if typ == xdr.AssetTypeAssetTypeNative {
		result, _ := xdr.NewAsset(typ, nil)
		return result
	}

	an := xdr.AssetAlphaNum4{}
	copy(an.AssetCode[:], code[:])

	raw := psrkey.MustDecode(psrkey.VersionByteAccountID, issuer)
	var key xdr.Uint256
	copy(key[:], raw)

	an.Issuer, _ = xdr.NewAccountId(xdr.PublicKeyTypePublicKeyTypeEd25519, key)

	result, _ := xdr.NewAsset(typ, an)
	return result
}
