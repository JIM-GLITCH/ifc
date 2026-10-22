package ifc

import "strings"

func entityTypeOfValue(stepName string) WhereValue {
	names := entityTypeOfQualifiedNames(stepName)
	if len(names) == 0 {
		return wvUnknownVal()
	}
	items := make([]WhereValue, len(names))
	for i, name := range names {
		items[i] = wvFromString(name)
	}
	return wvFromList(items)
}

// entityTypeOfQualifiedNames returns the qualified type names
// ("<schema>.<STEP name>") of the given entity and all of its supertypes.
//
// The result is derived from the immutable generated Entities table and cached
// per entity name; WHERE rules evaluate TYPEOF(...) for every entity instance,
// so caching removes both the supertype walk and the reverse name lookups from
// the per-instance hot path. The returned slice is shared and read-only.
func entityTypeOfQualifiedNames(stepName string) []string {
	key := strings.ToUpper(strings.TrimSpace(stepName))
	if cached, ok := typeOfQualifiedNamesCache.Load(key); ok {
		names, _ := cached.([]string)
		return names
	}
	names := computeEntityTypeOfQualifiedNames(key)
	typeOfQualifiedNamesCache.Store(key, names)
	return names
}

func computeEntityTypeOfQualifiedNames(stepName string) []string {
	def, ok := LookupEntity(stepName)
	if !ok {
		return nil
	}
	visited := map[string]bool{}
	var names []string
	var walk func(EntityDef)
	walk = func(def EntityDef) {
		if visited[def.ExpressName] {
			return
		}
		visited[def.ExpressName] = true
		if step, ok := lookupStepByExpressName(def.ExpressName); ok {
			names = append(names, SchemaName+"."+step)
		}
		for _, superName := range def.SuperTypes {
			if superDef, ok := lookupEntityByExpressName(superName); ok {
				walk(superDef)
			}
		}
	}
	walk(def)
	return names
}

// stepByExpressNameIndex returns the reverse index from lower-cased EXPRESS
// entity name to upper-case STEP name. It is built once from the immutable
// generated Entities table, replacing a linear EqualFold scan over all
// entities that used to run on every TYPEOF(...) evaluation.
func stepByExpressNameIndex() map[string]string {
	stepByExpressNameOnce.Do(func() {
		stepByExpressNameMap = make(map[string]string, len(Entities))
		for step, def := range Entities {
			lower := strings.ToLower(def.ExpressName)
			if _, exists := stepByExpressNameMap[lower]; !exists {
				stepByExpressNameMap[lower] = step
			}
		}
	})
	return stepByExpressNameMap
}

func lookupStepByExpressName(expressName string) (string, bool) {
	step, ok := stepByExpressNameIndex()[strings.ToLower(expressName)]
	return step, ok
}

func lookupEntityByExpressName(expressName string) (EntityDef, bool) {
	step, ok := lookupStepByExpressName(expressName)
	if !ok {
		return EntityDef{}, false
	}
	return LookupEntity(step)
}
