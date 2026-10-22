package ifc

import (
	"context"
	"strings"
)

// entityEvalEntry caches an attribute value together with the reason its
// conversion failed, if any.
//
// The two pieces have different lifetimes: the value is a pure memo (safe to
// share across evaluations), while the failure reason drives eval-unknown
// reporting. Keeping the reason next to the value lets a cache hit replay the
// note on the current evaluation, so which evaluation reports a failed
// attribute no longer depends on who computed it first. That matters under
// parallel validation, where per-worker caches make "first reader"
// scheduling-dependent.
type entityEvalEntry struct {
	value      WhereValue
	failReason string
}

type entityEvalCache struct {
	values map[string]entityEvalEntry
}

type entityEvalCacheKey struct{}

func newEntityEvalCache() *entityEvalCache {
	return &entityEvalCache{values: make(map[string]entityEvalEntry)}
}

func WithEntityEvalCache(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if entityEvalCacheFromContext(ctx) != nil {
		return ctx
	}
	return context.WithValue(ctx, entityEvalCacheKey{}, newEntityEvalCache())
}

func entityEvalCacheFromContext(ctx context.Context) *entityEvalCache {
	if ctx == nil {
		return nil
	}
	cache, _ := ctx.Value(entityEvalCacheKey{}).(*entityEvalCache)
	return cache
}

func entityAttrCacheKey(inst SIMPLE_ENTITY_INSTANCE, attrName string) string {
	if inst == nil {
		return ""
	}
	id := strings.TrimSpace(inst.Name())
	if id == "" {
		return ""
	}
	// instanceStepName already returns an upper-case, trimmed STEP name.
	step, ok := instanceStepName(inst)
	if !ok {
		return id + "|" + strings.ToLower(strings.TrimSpace(attrName))
	}
	return step + "|" + id + "|" + strings.ToLower(strings.TrimSpace(attrName))
}

func (c *entityEvalCache) lookup(inst SIMPLE_ENTITY_INSTANCE, attrName string) (entityEvalEntry, bool) {
	if c == nil || inst == nil {
		return entityEvalEntry{}, false
	}
	key := entityAttrCacheKey(inst, attrName)
	if key == "" {
		return entityEvalEntry{}, false
	}
	v, ok := c.values[key]
	return v, ok
}

func (c *entityEvalCache) store(inst SIMPLE_ENTITY_INSTANCE, attrName string, entry entityEvalEntry) {
	if c == nil || inst == nil {
		return
	}
	key := entityAttrCacheKey(inst, attrName)
	if key == "" {
		return
	}
	c.values[key] = entry
}
