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

func entityTypeOfQualifiedNames(stepName string) []string {
	def, ok := LookupEntity(strings.ToUpper(strings.TrimSpace(stepName)))
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

func lookupStepByExpressName(expressName string) (string, bool) {
	for step, def := range Entities {
		if strings.EqualFold(def.ExpressName, expressName) {
			return step, true
		}
	}
	return "", false
}

func lookupEntityByExpressName(expressName string) (EntityDef, bool) {
	step, ok := lookupStepByExpressName(expressName)
	if !ok {
		return EntityDef{}, false
	}
	return LookupEntity(step)
}
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

func entityTypeOfQualifiedNames(stepName string) []string {
	def, ok := LookupEntity(strings.ToUpper(strings.TrimSpace(stepName)))
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

func lookupStepByExpressName(expressName string) (string, bool) {
	for step, def := range Entities {
		if strings.EqualFold(def.ExpressName, expressName) {
			return step, true
		}
	}
	return "", false
}

func lookupEntityByExpressName(expressName string) (EntityDef, bool) {
	step, ok := lookupStepByExpressName(expressName)
	if !ok {
		return EntityDef{}, false
	}
	return LookupEntity(step)
}
