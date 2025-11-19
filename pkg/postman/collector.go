package postman

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// RouteInfo holds information about a route
type RouteInfo struct {
	Method        string
	Path          string
	Name          string
	PathParams    []string
	RequiresAuth  bool
	RequiresAdmin bool
	Description   string
}

// RouteCollector collects routes from a mux router
type RouteCollector struct {
	routes []RouteInfo
}

// NewRouteCollector creates a new route collector
func NewRouteCollector() *RouteCollector {
	return &RouteCollector{
		routes: []RouteInfo{},
	}
}

// CollectRoutes extracts all routes from a mux router
func (rc *RouteCollector) CollectRoutes(router *mux.Router) []RouteInfo {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		methods, err := route.GetMethods()
		if err != nil {
			// Skip routes without methods
			return nil
		}

		// Get route name if available
		name := route.GetName()

		for _, method := range methods {
			routeInfo := RouteInfo{
				Method:        method,
				Path:          path,
				Name:          name,
				PathParams:    extractPathParams(path),
				RequiresAuth:  detectAuthMiddleware(route),
				RequiresAdmin: detectAdminMiddleware(route),
			}
			rc.routes = append(rc.routes, routeInfo)
		}

		return nil
	})

	return rc.routes
}

// extractPathParams extracts path parameters from a route path
func extractPathParams(path string) []string {
	var params []string
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			param := strings.Trim(part, "{}")
			params = append(params, param)
		}
	}
	return params
}

// detectAuthMiddleware attempts to detect if a route requires authentication
// This is a heuristic based on common patterns
func detectAuthMiddleware(route *mux.Route) bool {
	path, _ := route.GetPathTemplate()
	// Simple heuristic: auth routes typically don't require auth
	if strings.Contains(path, "/auth/") {
		return false
	}
	// Most other routes in protected APIs require auth
	// This can be enhanced based on your middleware detection needs
	return strings.Contains(path, "/user") ||
		strings.Contains(path, "/role") ||
		strings.Contains(path, "/permission")
}

// detectAdminMiddleware attempts to detect if a route requires admin privileges
func detectAdminMiddleware(route *mux.Route) bool {
	path, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()

	// Heuristic: admin routes typically involve modifications
	for _, method := range methods {
		if method == http.MethodPost ||
			method == http.MethodPut ||
			method == http.MethodPatch ||
			method == http.MethodDelete {
			// User management, role management typically require admin
			if strings.Contains(path, "/user") ||
				strings.Contains(path, "/role") ||
				strings.Contains(path, "/permission") {
				return true
			}
		}
	}
	return false
}

// GroupByPrefix groups routes by their path prefix
func (rc *RouteCollector) GroupByPrefix() map[string][]RouteInfo {
	grouped := make(map[string][]RouteInfo)

	for _, route := range rc.routes {
		prefix := extractPrefix(route.Path)
		grouped[prefix] = append(grouped[prefix], route)
	}

	return grouped
}

// extractPrefix extracts the primary prefix from a path
func extractPrefix(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return "root"
}

// GenerateName creates a human-readable name for a route
func GenerateName(route RouteInfo) string {
	if route.Name != "" {
		return route.Name
	}

	// Generate name from method and path
	path := strings.Trim(route.Path, "/")
	path = strings.ReplaceAll(path, "/", " ")
	path = strings.ReplaceAll(path, "{", ":")
	path = strings.ReplaceAll(path, "}", "")

	return fmt.Sprintf("%s %s", route.Method, path)
}
