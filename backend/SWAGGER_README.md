# Swagger API Documentation

This project uses Swagger (OpenAPI) for API documentation. The Swagger UI is available at `/swagger/index.html` when the server is running.

## Accessing the Documentation

Once the server is running, you can access the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

## Regenerating Documentation

After making changes to API endpoints, controllers, or request/response models, you need to regenerate the Swagger documentation:

```bash
cd backend
swag init -g cmd/main.go -o docs
```

This will update the `docs/` directory with the latest API documentation.

## Adding Swagger Annotations

### Controller Functions

Add annotations above each controller function:

```go
// FunctionName godoc
// @Summary      Brief summary
// @Description  Detailed description
// @Tags         TagName
// @Accept       json
// @Produce      json
// @Security     BearerAuth  // If authentication required
// @Param        paramName  body      domain.Model  true  "Description"
// @Success      200         {object}  shared.ApiResponse{data=domain.Model}
// @Failure      400         {object}  shared.ErrorResponse
// @Router       /endpoint [method]
func (c *Controller) FunctionName(ctx *gin.Context) {
    // ...
}
```

### Security

For endpoints that require authentication, add:
```go
// @Security     BearerAuth
```

The Bearer token authentication is configured in `main.go`. When testing in Swagger UI, click the "Authorize" button and enter your JWT token in the format: `Bearer <your-token>`

## Swagger Annotations Reference

- `@Summary`: Brief description of the endpoint
- `@Description`: Detailed description
- `@Tags`: Group endpoints (e.g., "Auth", "Hotels", "Users", "Bookings")
- `@Accept`: Content type accepted (usually `json`)
- `@Produce`: Content type produced (usually `json`)
- `@Param`: Request parameters (path, query, body, header)
- `@Success`: Success response format
- `@Failure`: Error response format
- `@Router`: Endpoint path and HTTP method
- `@Security`: Authentication requirement

## Model Documentation

Add example values to struct fields for better documentation:

```go
type Model struct {
    Field string `json:"field" example:"example_value"`
}
```

