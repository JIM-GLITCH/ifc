package ifc

import "sync"

// This file holds the memoization caches for schema-derived lookups.
//
// Every cache below is derived exclusively from package-level generated data
// (Entities, Types, ForwardInverses, entityAttrNames), which is populated
// during package init and never mutated afterwards. The derived results are
// therefore stable for the lifetime of the process and can be computed once
// and shared by every document and every validation pass.
//
// Before memoization, the hottest of these derivations (forward-inverse
// merging, attribute index resolution, SELECT member resolution and the
// supertype closure used for global entity sets) ran per entity instance,
// which dominated validation time on large IFC files.
//
// sync.Map is used instead of a plain map plus mutex so the caches are safe
// for concurrent access; this keeps the door open for sharding validation
// across goroutines later on.

// mergedForwardInversesCache maps an upper-case STEP entity name to its merged
// forward-inverse map (own entries plus entries inherited from supertypes).
var mergedForwardInversesCache sync.Map // string -> map[int][]ForwardInverseTarget

// attrIndexCache maps an upper-case STEP entity name to a lookup from
// lower-cased attribute name to its parameter index.
var attrIndexCache sync.Map // string -> map[string]int

// resolvedTypeValueKindCache maps a type name to its resolved value kind,
// following alias chains through the Types table.
var resolvedTypeValueKindCache sync.Map // string -> ParamKind

// allowedSelectKindsCache maps a SELECT type name to the parameter kinds that
// are acceptable for values of that type.
var allowedSelectKindsCache sync.Map // string -> []ParamKind

// expressClosureCache maps an upper-case STEP entity name to the EXPRESS names
// of the entity itself and all of its supertypes, deduplicated, self first.
var expressClosureCache sync.Map // string -> []string

// typeOfQualifiedNamesCache maps an upper-case STEP entity name to the
// qualified TYPEOF names ("<schema>.<STEP name>") of the entity and all of its
// supertypes.
var typeOfQualifiedNamesCache sync.Map // string -> []string

// stepByExpressNameIndex is the reverse index from lower-cased EXPRESS entity
// name to upper-case STEP name. It is built once on first use.
var (
	stepByExpressNameOnce sync.Once
	stepByExpressNameMap  map[string]string
)
