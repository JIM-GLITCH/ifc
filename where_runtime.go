package ifc

import (
	"math"
	"strconv"
	"strings"
)

type wvKind int

const (
	wvKindUnknown wvKind = iota
	wvKindAbsent
	wvKindInteger
	wvKindReal
	wvKindString
	wvKindEnum
	wvKindLogical
	wvKindList
	wvKindEnumType
	wvKindRecord
)

type WhereValue struct {
	kind        wvKind
	i           int64
	r           float64
	s           string
	list        []WhereValue
	log         int // 0=false, 1=true, 2=unknown
	entity      SIMPLE_ENTITY_INSTANCE
	entities    []SIMPLE_ENTITY_INSTANCE
	recordType  string
	recordAttrs map[string]WhereValue
}

func wvUnknownVal() WhereValue { return WhereValue{kind: wvKindUnknown} }
func wvAbsentVal() WhereValue  { return WhereValue{kind: wvKindAbsent} }
func wvUnknownLogical() WhereValue {
	return WhereValue{kind: wvKindLogical, log: 2}
}

func wvIsIndeterminate(v WhereValue) bool {
	if v.kind == wvKindUnknown || v.kind == wvKindAbsent {
		return true
	}
	return v.kind == wvKindLogical && v.log == 2
}

func wvFromString(s string) WhereValue {
	return WhereValue{kind: wvKindString, s: s}
}

func wvFromEnum(s string) WhereValue {
	return WhereValue{kind: wvKindEnum, s: s}
}

func wvFromEnumType(typeName string) WhereValue {
	return WhereValue{kind: wvKindEnumType, s: typeName}
}

func wvFromEnumStepLiteral(member string) WhereValue {
	member = normalizeEnumMember(member)
	if member == "" {
		return wvUnknownVal()
	}
	return wvFromEnum("." + member + ".")
}

func normalizeEnumMember(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, ".") && strings.HasSuffix(s, ".") {
		return strings.ToUpper(strings.Trim(s, "."))
	}
	if i := strings.LastIndex(s, "."); i >= 0 {
		s = s[i+1:]
	}
	return strings.Trim(strings.ToUpper(s), ".")
}

func wvEnumMemberLabel(v WhereValue) (string, bool) {
	if v.kind != wvKindEnum {
		return "", false
	}
	return normalizeEnumMember(v.s), true
}

func wvFromList(items []WhereValue) WhereValue {
	return WhereValue{kind: wvKindList, list: items}
}

func wvFromInt(n int) WhereValue {
	return WhereValue{kind: wvKindInteger, i: int64(n)}
}

func parseLiteralNumber(text string) WhereValue {
	text = strings.TrimSpace(strings.TrimSuffix(text, "."))
	if strings.Contains(text, ".") {
		f, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return wvUnknownVal()
		}
		return WhereValue{kind: wvKindReal, r: f}
	}
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindInteger, i: i}
}

func parseLiteralLogical(text string) WhereValue {
	switch strings.ToUpper(strings.TrimSpace(text)) {
	case "TRUE", ".T.":
		return WhereValue{kind: wvKindLogical, log: 1}
	case "FALSE", ".F.":
		return WhereValue{kind: wvKindLogical, log: 0}
	default:
		return wvUnknownLogical()
	}
}

func wvTruthy(v WhereValue) bool {
	switch v.kind {
	case wvKindLogical:
		return v.log == 1
	case wvKindInteger:
		return v.i != 0
	case wvKindReal:
		return v.r != 0
	case wvKindString, wvKindEnum:
		return v.s != ""
	case wvKindList:
		return len(v.list) > 0
	case wvKindUnknown, wvKindAbsent:
		return false
	default:
		return false
	}
}

func wvIsDefiniteTrue(v WhereValue) bool {
	if wvIsIndeterminate(v) {
		return false
	}
	return wvTruthy(v)
}

func wvIsDefiniteFalse(v WhereValue) bool {
	if wvIsIndeterminate(v) {
		return false
	}
	if v.kind == wvKindLogical {
		return v.log == 0
	}
	return !wvTruthy(v)
}

func wvAsNumber(v WhereValue) (float64, bool) {
	switch v.kind {
	case wvKindInteger:
		return float64(v.i), true
	case wvKindReal:
		return v.r, true
	case wvKindString:
		if strings.Contains(v.s, ".") {
			f, err := strconv.ParseFloat(v.s, 64)
			return f, err == nil
		}
		i, err := strconv.ParseInt(v.s, 10, 64)
		return float64(i), err == nil
	default:
		return 0, false
	}
}

func wvAsInt(v WhereValue) int {
	n, ok := wvAsNumber(v)
	if !ok {
		return 0
	}
	return int(n)
}

func wvAsString(v WhereValue) (string, bool) {
	switch v.kind {
	case wvKindString, wvKindEnum:
		return v.s, true
	case wvKindInteger:
		return strconv.FormatInt(v.i, 10), true
	case wvKindReal:
		return strconv.FormatFloat(v.r, 'f', -1, 64), true
	default:
		return "", false
	}
}

func wvCompareRel(left, right WhereValue, op string) WhereValue {
	ln, lok := wvAsNumber(left)
	rn, rok := wvAsNumber(right)
	if lok && rok {
		ok := false
		switch op {
		case "=":
			ok = ln == rn
		case "<>":
			ok = ln != rn
		case "<":
			ok = ln < rn
		case ">":
			ok = ln > rn
		case "<=":
			ok = ln <= rn
		case ">=":
			ok = ln >= rn
		default:
			return wvUnknownLogical()
		}
		return WhereValue{kind: wvKindLogical, log: boolToLog(ok)}
	}
	if le, lok := wvEnumMemberLabel(left); lok {
		if re, rok := wvEnumMemberLabel(right); rok {
			cmp := strings.Compare(le, re)
			return WhereValue{kind: wvKindLogical, log: boolToLog(relCompare(cmp, op))}
		}
	}
	if left.kind == wvKindLogical && right.kind == wvKindLogical {
		same := left.log == right.log
		switch op {
		case "=":
			return WhereValue{kind: wvKindLogical, log: boolToLog(same)}
		case "<>":
			return WhereValue{kind: wvKindLogical, log: boolToLog(!same)}
		default:
			return wvUnknownLogical()
		}
	}
	if ls, lok := wvAsString(left); lok {
		if rs, rok := wvAsString(right); rok {
			cmp := strings.Compare(ls, rs)
			return WhereValue{kind: wvKindLogical, log: boolToLog(relCompare(cmp, op))}
		}
	}
	if op == "=" || op == "<>" {
		leftNorm := wvNormalizeForCompare(left)
		rightNorm := wvNormalizeForCompare(right)
		if leftNorm.kind == wvKindRecord && rightNorm.kind == wvKindRecord {
			eq := wvRecordsEqual(leftNorm, rightNorm)
			if op == "=" {
				return WhereValue{kind: wvKindLogical, log: boolToLog(eq)}
			}
			return WhereValue{kind: wvKindLogical, log: boolToLog(!eq)}
		}
	}
	return wvUnknownLogical()
}

func relCompare(cmp int, op string) bool {
	switch op {
	case "=":
		return cmp == 0
	case "<>":
		return cmp != 0
	case "<":
		return cmp < 0
	case ">":
		return cmp > 0
	case "<=":
		return cmp <= 0
	case ">=":
		return cmp >= 0
	default:
		return true
	}
}

func boolToLog(b bool) int {
	if b {
		return 1
	}
	return 0
}

func wvAnd(a, b WhereValue) WhereValue {
	if wvIsDefiniteFalse(a) || wvIsDefiniteFalse(b) {
		return WhereValue{kind: wvKindLogical, log: 0}
	}
	if wvIsDefiniteTrue(a) && wvIsDefiniteTrue(b) {
		return WhereValue{kind: wvKindLogical, log: 1}
	}
	return wvUnknownLogical()
}

func wvOr(a, b WhereValue) WhereValue {
	if wvIsDefiniteTrue(a) || wvIsDefiniteTrue(b) {
		return WhereValue{kind: wvKindLogical, log: 1}
	}
	if wvIsDefiniteFalse(a) && wvIsDefiniteFalse(b) {
		return WhereValue{kind: wvKindLogical, log: 0}
	}
	return wvUnknownLogical()
}

func wvXor(a, b WhereValue) WhereValue {
	if wvIsIndeterminate(a) || wvIsIndeterminate(b) {
		return wvUnknownLogical()
	}
	return WhereValue{kind: wvKindLogical, log: boolToLog(wvTruthy(a) != wvTruthy(b))}
}

func wvNot(v WhereValue) WhereValue {
	if wvIsIndeterminate(v) {
		return wvUnknownLogical()
	}
	return WhereValue{kind: wvKindLogical, log: boolToLog(!wvTruthy(v))}
}

func wvIn(item, set WhereValue) WhereValue {
	itemStr, ok := wvAsString(item)
	if !ok {
		return wvUnknownLogical()
	}
	if set.kind != wvKindList {
		if s, ok := wvAsString(set); ok {
			return WhereValue{kind: wvKindLogical, log: boolToLog(strings.EqualFold(itemStr, s))}
		}
		return wvUnknownLogical()
	}
	for _, candidate := range set.list {
		if cs, ok := wvAsString(candidate); ok && strings.EqualFold(itemStr, cs) {
			return WhereValue{kind: wvKindLogical, log: 1}
		}
	}
	return WhereValue{kind: wvKindLogical, log: 0}
}

func wvValueEqual(a, b WhereValue) bool {
	if same := wvInstSame(a, b); same.kind == wvKindLogical && same.log == 1 {
		return true
	}
	ae, aok := wvEnumMemberLabel(a)
	be, bok := wvEnumMemberLabel(b)
	return aok && bok && ae == be
}

func wvSetUnion(a, b WhereValue) WhereValue {
	if a.kind != wvKindList {
		if a.kind == wvKindUnknown {
			return wvFromList([]WhereValue{b})
		}
		return wvUnknownVal()
	}
	out := append([]WhereValue{}, a.list...)
	for _, existing := range out {
		if wvValueEqual(existing, b) {
			return wvFromList(out)
		}
	}
	out = append(out, b)
	return wvFromList(out)
}

func wvEntityTypeTag(typeName string) WhereValue {
	return WhereValue{kind: wvKindRecord, recordType: typeName}
}

func wvNewDirection(ratios WhereValue) WhereValue {
	return WhereValue{
		kind:       wvKindRecord,
		recordType: "IfcDirection",
		recordAttrs: map[string]WhereValue{
			"DirectionRatios": ratios,
		},
	}
}

func wvNewVector(orientation, magnitude WhereValue) WhereValue {
	return WhereValue{
		kind:       wvKindRecord,
		recordType: "IfcVector",
		recordAttrs: map[string]WhereValue{
			"Orientation": orientation,
			"Magnitude":   magnitude,
		},
	}
}

func wvNewEntityRecord(typeName string, attrNames []string, args []WhereValue) WhereValue {
	attrs := make(map[string]WhereValue, len(attrNames))
	for i, name := range attrNames {
		if i >= len(args) {
			break
		}
		attrs[name] = args[i]
	}
	return WhereValue{kind: wvKindRecord, recordType: typeName, recordAttrs: attrs}
}

func wvRecordsEqual(a, b WhereValue) bool {
	if a.kind != wvKindRecord || b.kind != wvKindRecord {
		return false
	}
	if a.recordType != "" && b.recordType != "" && !strings.EqualFold(a.recordType, b.recordType) {
		return false
	}
	if len(a.recordAttrs) != len(b.recordAttrs) {
		return false
	}
	for name, av := range a.recordAttrs {
		bv, ok := b.recordAttrs[name]
		if !ok {
			return false
		}
		if !wvTruthy(wvCompareRel(av, bv, "=")) {
			return false
		}
	}
	return true
}

func wvEntityIntersect(a, b WhereValue) WhereValue {
	if b.kind == wvKindRecord || b.kind == wvKindList || b.entityInstance() != nil {
		return b
	}
	if a.kind == wvKindRecord || a.kind == wvKindList || a.entityInstance() != nil {
		return a
	}
	return b
}

func cloneRecord(v WhereValue) WhereValue {
	if v.kind != wvKindRecord {
		return v
	}
	attrs := make(map[string]WhereValue, len(v.recordAttrs))
	for k, val := range v.recordAttrs {
		attrs[k] = val
	}
	return WhereValue{kind: wvKindRecord, recordType: v.recordType, recordAttrs: attrs}
}

func wvAssignAttr(base WhereValue, attrName string, value WhereValue) WhereValue {
	out := cloneRecord(base)
	if out.kind != wvKindRecord {
		out = WhereValue{kind: wvKindRecord, recordAttrs: map[string]WhereValue{}}
	}
	if out.recordAttrs == nil {
		out.recordAttrs = map[string]WhereValue{}
	}
	out.recordAttrs[attrName] = value
	return out
}

func wvAssignListElem(base WhereValue, attrName string, index int, value WhereValue) WhereValue {
	out := cloneRecord(base)
	if out.kind != wvKindRecord {
		out = WhereValue{kind: wvKindRecord, recordAttrs: map[string]WhereValue{}}
	}
	if out.recordAttrs == nil {
		out.recordAttrs = map[string]WhereValue{}
	}
	list := out.recordAttrs[attrName]
	items := []WhereValue{}
	if list.kind == wvKindList {
		items = append(items, list.list...)
	}
	for len(items) < index {
		items = append(items, wvUnknownVal())
	}
	if index < 1 {
		return out
	}
	items[index-1] = value
	out.recordAttrs[attrName] = wvFromList(items)
	return out
}

func wvAssignAtIndex(base WhereValue, index int, value WhereValue) WhereValue {
	items := []WhereValue{}
	if base.kind == wvKindList {
		items = append(items, base.list...)
	}
	for len(items) < index {
		items = append(items, wvUnknownVal())
	}
	if index < 1 {
		return base
	}
	items[index-1] = value
	return wvFromList(items)
}

func wvRepeatList(item WhereValue, count WhereValue) WhereValue {
	n := wvAsInt(count)
	if n < 0 {
		return wvUnknownVal()
	}
	items := make([]WhereValue, n)
	for i := range items {
		items[i] = item
	}
	return wvFromList(items)
}

func wvAdd(a, b WhereValue) WhereValue {
	an, aok := wvAsNumber(a)
	bn, bok := wvAsNumber(b)
	if aok && bok {
		return WhereValue{kind: wvKindReal, r: an + bn}
	}
	if a.kind == wvKindList {
		return wvSetUnion(a, b)
	}
	return wvUnknownVal()
}

func wvSub(a, b WhereValue) WhereValue {
	an, aok := wvAsNumber(a)
	bn, bok := wvAsNumber(b)
	if !aok || !bok {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindReal, r: an - bn}
}

func wvMul(a, b WhereValue) WhereValue {
	an, aok := wvAsNumber(a)
	bn, bok := wvAsNumber(b)
	if aok && bok {
		return WhereValue{kind: wvKindReal, r: an * bn}
	}
	return wvSetIntersect(a, b)
}

func wvDiv(a, b WhereValue) WhereValue {
	an, aok := wvAsNumber(a)
	bn, bok := wvAsNumber(b)
	if !aok || !bok || bn == 0 {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindReal, r: an / bn}
}

func wvMod(a, b WhereValue) WhereValue {
	an, aok := wvAsNumber(a)
	bn, bok := wvAsNumber(b)
	if !aok || !bok || bn == 0 {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindInteger, i: int64(math.Mod(an, bn))}
}

func wvAbs(v WhereValue) WhereValue {
	n, ok := wvAsNumber(v)
	if !ok {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindReal, r: math.Abs(n)}
}

func wvSqrt(v WhereValue) WhereValue {
	n, ok := wvAsNumber(v)
	if !ok {
		return wvUnknownVal()
	}
	return WhereValue{kind: wvKindReal, r: math.Sqrt(n)}
}

func wvSizeOf(v WhereValue) WhereValue {
	if v.kind == wvKindList {
		return WhereValue{kind: wvKindInteger, i: int64(len(v.list))}
	}
	if v.kind == wvKindEntityList {
		return WhereValue{kind: wvKindInteger, i: int64(len(v.entities))}
	}
	return wvUnknownVal()
}

func wvHiIndex(v WhereValue) WhereValue {
	return wvSizeOf(v)
}

func wvLoIndex(v WhereValue) WhereValue {
	return wvFromInt(1)
}

func wvExists(v WhereValue) WhereValue {
	if v.kind == wvKindUnknown || v.kind == wvKindAbsent {
		return WhereValue{kind: wvKindLogical, log: 0}
	}
	return WhereValue{kind: wvKindLogical, log: 1}
}

func wvTypeOf(v WhereValue) WhereValue {
	inst := v.entityInstance()
	if inst == nil {
		return wvUnknownVal()
	}
	step, ok := instanceStepName(inst)
	if !ok {
		return wvUnknownVal()
	}
	return entityTypeOfValue(step)
}

func wvEntityInTypeOf(v WhereValue, expectedExpressName string) WhereValue {
	if inst := v.entityInstance(); inst != nil {
		step, ok := instanceStepName(inst)
		if !ok {
			return wvUnknownLogical()
		}
		return WhereValue{kind: wvKindLogical, log: boolToLog(entityIsOrSubtypes(step, expectedExpressName))}
	}
	if step, ok := whereValueExpressType(v); ok {
		return WhereValue{kind: wvKindLogical, log: boolToLog(entityIsOrSubtypes(step, expectedExpressName))}
	}
	return wvUnknownLogical()
}

func typeOfLabelSet(v WhereValue) map[string]struct{} {
	if v.kind != wvKindList {
		return nil
	}
	set := make(map[string]struct{}, len(v.list))
	for _, item := range v.list {
		if s, ok := wvAsString(item); ok {
			set[strings.ToLower(s)] = struct{}{}
		}
	}
	return set
}

func wvTypeOfEqual(a, b WhereValue) WhereValue {
	sa := typeOfLabelSet(a)
	sb := typeOfLabelSet(b)
	if sa == nil || sb == nil {
		return wvUnknownLogical()
	}
	if len(sa) != len(sb) {
		return WhereValue{kind: wvKindLogical, log: 0}
	}
	for k := range sa {
		if _, ok := sb[k]; !ok {
			return WhereValue{kind: wvKindLogical, log: 0}
		}
	}
	return WhereValue{kind: wvKindLogical, log: 1}
}

func wvTypeOfNotEqual(a, b WhereValue) WhereValue {
	eq := wvTypeOfEqual(a, b)
	if eq.kind == wvKindLogical && eq.log == -1 {
		return eq
	}
	return WhereValue{kind: wvKindLogical, log: boolToLog(eq.log != 1)}
}
func wvNVL(a, b WhereValue) WhereValue {
	if a.kind != wvKindUnknown && a.kind != wvKindAbsent {
		return a
	}
	return b
}
func wvBLength(v WhereValue) WhereValue {
	if s, ok := wvAsString(v); ok {
		return WhereValue{kind: wvKindInteger, i: int64(len(s))}
	}
	if v.kind == wvKindList {
		return WhereValue{kind: wvKindInteger, i: int64(len(v.list))}
	}
	return wvUnknownVal()
}
func wvInstSame(a, b WhereValue) WhereValue {
	aInst := a.entityInstance()
	bInst := b.entityInstance()
	if aInst != nil && bInst != nil {
		nameA := strings.TrimSpace(aInst.Name())
		nameB := strings.TrimSpace(bInst.Name())
		same := nameA != "" && nameA == nameB
		return WhereValue{kind: wvKindLogical, log: boolToLog(same)}
	}
	as, aok := wvAsString(a)
	bs, bok := wvAsString(b)
	if aok && bok {
		return WhereValue{kind: wvKindLogical, log: boolToLog(strings.EqualFold(as, bs))}
	}
	return wvUnknownLogical()
}
func wvSetIntersect(a, b WhereValue) WhereValue {
	sb := typeOfLabelSet(b)
	if sb == nil {
		if s, ok := wvAsString(b); ok {
			sb = map[string]struct{}{strings.ToLower(s): {}}
		} else {
			return wvUnknownVal()
		}
	}
	if a.kind != wvKindList {
		if s, ok := wvAsString(a); ok {
			if _, ok := sb[strings.ToLower(s)]; ok {
				return wvFromList([]WhereValue{wvFromString(s)})
			}
			return wvFromList([]WhereValue{})
		}
		return wvUnknownVal()
	}
	out := make([]WhereValue, 0, len(a.list))
	for _, item := range a.list {
		if s, ok := wvAsString(item); ok {
			if _, ok := sb[strings.ToLower(s)]; ok {
				out = append(out, item)
			}
		}
	}
	return wvFromList(out)
}
func wvPower(a, _ WhereValue) WhereValue { return a }
func wvQuery(v WhereValue, pred func(WhereValue) WhereValue) WhereValue {
	var elements []WhereValue
	switch v.kind {
	case wvKindList:
		elements = v.list
	case wvKindEntityList:
		elements = make([]WhereValue, len(v.entities))
		for i, inst := range v.entities {
			elements[i] = wvFromEntityInstance(inst)
		}
	default:
		return wvUnknownVal()
	}
	out := make([]WhereValue, 0, len(elements))
	for _, el := range elements {
		if wvTruthy(pred(el)) {
			out = append(out, el)
		}
	}
	return wvFromList(out)
}

type TypeWhereEval struct {
	self WhereValue
}

func NewTypeWhereEval(self WhereValue) *TypeWhereEval {
	return &TypeWhereEval{self: self}
}

func (ec *TypeWhereEval) Self() WhereValue {
	return ec.self
}

func (ec *TypeWhereEval) QualIndex(v WhereValue, index int) WhereValue {
	if v.kind != wvKindList || index < 1 || index > len(v.list) {
		return wvUnknownVal()
	}
	return v.list[index-1]
}

func (ec *TypeWhereEval) QualAttr(_ WhereValue, _ string, _ int) WhereValue {
	return wvUnknownVal()
}

func (ec *TypeWhereEval) QualEntity(_ WhereValue, _ string) WhereValue {
	return wvUnknownVal()
}

func (ec *TypeWhereEval) Attr(_ string) WhereValue {
	return wvUnknownVal()
}

func (ec *TypeWhereEval) Derived(_ string) WhereValue {
	return wvUnknownVal()
}

func (ec *TypeWhereEval) CallSchema(_ string, _ []WhereValue) WhereValue {
	return wvUnknownVal()
}

func LookupTypeWhereRules(typeName string) ([]TypeWhereRuleSpec, bool) {
	if rules, ok := TypeWhereRules[typeName]; ok {
		return rules, true
	}
	for name, rules := range TypeWhereRules {
		if strings.EqualFold(name, typeName) {
			return rules, true
		}
	}
	return nil, false
}

func CheckTypeWhereRules(typeName string, self WhereValue) (ok bool, detail string) {
	rules, found := LookupTypeWhereRules(typeName)
	if !found || len(rules) == 0 {
		return true, ""
	}
	ec := NewTypeWhereEval(self)
	for _, rule := range rules {
		if rule.Check == nil {
			continue
		}
		if pass, msg := rule.Check(ec); !pass {
			return false, msg
		}
	}
	return true, ""
}
