package ifc

import (
	"context"
	"strings"
)

type entityEvalCache struct {
	values map[string]WhereValue
}

type entityEvalCacheKey struct{}

func newEntityEvalCache() *entityEvalCache {
	return &entityEvalCache{values: make(map[string]WhereValue)}
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
	step, ok := instanceStepName(inst)
	if !ok {
		return id + "|" + strings.ToLower(strings.TrimSpace(attrName))
	}
	return strings.ToUpper(step) + "|" + id + "|" + strings.ToLower(strings.TrimSpace(attrName))
}

func (c *entityEvalCache) lookup(inst SIMPLE_ENTITY_INSTANCE, attrName string) (WhereValue, bool) {
	if c == nil || inst == nil {
		return WhereValue{}, false
	}
	key := entityAttrCacheKey(inst, attrName)
	if key == "" {
		return WhereValue{}, false
	}
	v, ok := c.values[key]
	return v, ok
}

func (c *entityEvalCache) store(inst SIMPLE_ENTITY_INSTANCE, attrName string, value WhereValue) {
	if c == nil || inst == nil {
		return
	}
	key := entityAttrCacheKey(inst, attrName)
	if key == "" {
		return
	}
	c.values[key] = value
}
