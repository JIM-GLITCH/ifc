package ifc

import "strings"

func mergedInverseDefs(stepName string) []InverseDef {
	seen := map[string]bool{}
	var out []InverseDef
	var collect func(string)
	collect = func(name string) {
		upper := strings.ToUpper(strings.TrimSpace(name))
		if upper == "" || seen[upper] {
			return
		}
		seen[upper] = true
		out = append(out, EntityInverses[upper]...)
		def, ok := LookupEntity(upper)
		if !ok {
			return
		}
		for _, super := range def.SuperTypes {
			collect(super)
		}
	}
	collect(stepName)
	return out
}

func entityHasInverse(stepName, inverseName string) bool {
	needle := strings.ToLower(strings.TrimSpace(inverseName))
	if needle == "" {
		return false
	}
	for _, inv := range mergedInverseDefs(stepName) {
		if strings.ToLower(inv.Name) == needle {
			return true
		}
	}
	return false
}

func (ec *EntityWhereEval) inverseOn(inst SIMPLE_ENTITY_INSTANCE, inverseName string) WhereValue {
	if ec.pop == nil || inst == nil {
		return wvFromEntityInstances(nil)
	}
	items := ec.pop.InverseRelations(inst.Name(), inverseName)
	return wvFromEntityInstances(items)
}
