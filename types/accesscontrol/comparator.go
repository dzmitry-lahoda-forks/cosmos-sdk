package accesscontrol

import (
	"fmt"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	// Param Store Values can only be set during genesis and updated
	// through a gov proposal and those are always processed sequentially
	ConcurrentSafeIdentifiers = map[string]bool{
		"bank/SendEnabled": true,
		"bank/DefaultSendEnabled": true,
		"staking/BondDenom": true,
		"staking/UnbondingTime": true,
	}
)

type Comparator struct {
	AccessType AccessType
	Identifier 	string
	StoreKey string
}

func AccessTypeStringToEnum(accessType string) AccessType {
	switch strings.ToUpper(accessType) {
	case "WRITE":
		return AccessType_WRITE
	case "READ":
		return AccessType_READ
	default:
		panic(fmt.Sprintf("unknown accessType=%s", accessType))
	}
}

func BuildComparatorFromEvents(events []abci.Event, storeKeyToResourceTypePrefixMap StoreKeyToResourceTypePrefixMap) (comparators []Comparator) {
	for _, event := range events {
		if event.Type != "resource_access" {
			continue
		}
		attributes := event.GetAttributes()

		identifier := ""
		accessType := AccessType_UNKNOWN
		storeKey := ""
		for _, attribute := range attributes {
			if attribute.Key == "key" {
				identifier = attribute.Value
			}
			if attribute.Key == "access_type" {
				accessType =  AccessTypeStringToEnum(attribute.Value)
			}
			if attribute.Key == "store_key" {
				storeKey =  attribute.Value
			}
		}

		comparators = append(comparators, Comparator{
			AccessType: accessType,
			Identifier: identifier,
			StoreKey: storeKey,

		})
	}
	return comparators
}

func (c *Comparator) DependencyMatch(accessOp AccessOperation, prefix []byte) bool {
	// If the resource prefixes are the same, then its just the access type, if they're not the same
	// then they do not match. Make this the first condition check to avoid additional matching
	// as most of the time this will be enough to determine if they're dependency matches
	if c.AccessType != accessOp.AccessType && accessOp.AccessType != AccessType_UNKNOWN {
		return false
	}

	// The resource type was found in the parent store mapping or the child mapping
	if accessOp.GetIdentifierTemplate() == "*" {
		return true
	}

	// Both Identifiers should be starting with the same prefix expected for the resource type
	// e.g if the StoreKey and resource type is ResourceType_KV_BANK_BALANCES, then they both must start with BalancesPrefix
	if !strings.HasPrefix(c.Identifier, string(prefix)) || !strings.HasPrefix(accessOp.GetIdentifierTemplate(), string(prefix)) {
		return false
	}

	// With the same prefix, c.Identififer should be a superset of IdentifierTemplate, it not equal
	if !strings.Contains(c.Identifier, accessOp.GetIdentifierTemplate()) {
		return false
	}

	// At this point the resource identifier matches, but the access type is unknown, therefore
	// it will match both read and writes
	if accessOp.AccessType == AccessType_UNKNOWN {
		return true
	}

	return true
}

func (c *Comparator) IsConcurrentSafeIdentifier() bool {
	// Params are safe as they need to be updated through gov proposals
	if c.StoreKey == "params" {
		return true
	}
	if val, ok := ConcurrentSafeIdentifiers[c.Identifier]; ok {
		return val
	}
	return false
}

func (c *Comparator) String() string {
	return fmt.Sprintf("AccessType=%s, Identifier=%s\n", c.AccessType, c.Identifier)
}
