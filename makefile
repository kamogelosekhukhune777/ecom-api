run: 
	go run api/ecom/main.go

# ===================================================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor