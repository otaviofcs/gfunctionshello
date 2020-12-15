package main
import (
	"log"
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"example.com/function"
  smartenv "github.com/otaviofcs/smartenv"
)
func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", function.HelloWorld); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}
	port := smartenv.Env("PORT", "8080")
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
