package main 
import (
	"log"
	"os"
	"fmt"

	"github.com/wangpy1489/DNative/pkg/apis"
	"github.com/wangpy1489/DNative/pkg/controller"

	"github.com/operator-framework/operator-sdk/pkg/restmapper"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"github.com/wangpy1489/DNative/pkg/router"
)

var logger = log.New(os.Stdout,"", log.LstdFlags|log.Llongfile)

var (
	metricsHost               = "0.0.0.0"
	metricsPort         int32 = 8383
	operatorMetricsPort int32 = 8686
)

func init(){
	os.Setenv(k8sutil.ForceRunModeEnv, string(k8sutil.LocalRunMode))
	os.Setenv( k8sutil.WatchNamespaceEnvVar, "default")
}

func main ()  {
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logger.Fatal(err, "Failed to get watch namespace")
		os.Exit(1)
	}
	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal(err, "")
		os.Exit(1)
	}

	mgr, err := manager.New(cfg, manager.Options{
		Namespace:          namespace,
		MapperProvider:     restmapper.NewDynamicRESTMapper,
		MetricsBindAddress: fmt.Sprintf("%s:%d", metricsHost, metricsPort),
	})
	if err != nil {
		logger.Fatal(err, "")
		os.Exit(1)
	}

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		logger.Fatal(err, "")
		os.Exit(1)
	}

	// Setup all Controllers
	if err := controller.AddToManager(mgr); err != nil {
		logger.Fatal(err, "")
		os.Exit(1)
	}

	logger.Println("Registering Components.")

	rou, err := router.MakeRouter(logger,mgr.GetClient())
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("server start")
	go rou.Serve(8000)
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Fatal(err, "Manager exited non-zero")
		os.Exit(1)
	}
	logger.Fatal("server done")
}