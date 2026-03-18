package api

import (
	"context"

	"github.com/labstack/gommon/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Server struct {
	clientset *kubernetes.Clientset
}

func NewServer() Server {
	// local
	//var kubeconfig string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = filepath.Join(home, ".kube", "config")
	//}
	//
	//config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//if err != nil {
	//	log.Error(err.Error())
	//	panic(err.Error())
	//}

	// k8s
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}
	//

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

	return Server{
		clientset: clientset,
	}
}

func namespaceFromV1Namespace(namespace *corev1.Namespace) *Namespace {
	return &Namespace{
		CreatedAt: &namespace.CreationTimestamp.Time,
		Labels:    namespace.Labels,
		Name:      namespace.Name,
		Status:    (*NamespaceStatus)(&namespace.Status.Phase),
	}
}

func (s Server) ListNamespaces(ctx context.Context, request ListNamespacesRequestObject) (ListNamespacesResponseObject, error) {
	namespaces, err := s.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error(err.Error())
		return ListNamespaces500JSONResponse{}, nil
	}

	decoratedNamespaces := []Namespace{}
	for _, ns := range namespaces.Items {
		decoratedNamespaces = append(decoratedNamespaces, *namespaceFromV1Namespace(&ns))
	}

	return ListNamespaces200JSONResponse{
		Items: &decoratedNamespaces,
	}, nil
}

func (s Server) CreateNamespace(ctx context.Context, request CreateNamespaceRequestObject) (CreateNamespaceResponseObject, error) {
	namespaceDraft := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   request.Body.Name,
			Labels: request.Body.Labels,
		},
	}

	namespace, err := s.clientset.CoreV1().Namespaces().Create(ctx, &namespaceDraft, metav1.CreateOptions{})
	if err != nil {
		log.Error(err.Error())
		return CreateNamespace500Response{}, nil
	}

	return CreateNamespace201JSONResponse{
		CreatedAt: &namespace.CreationTimestamp.Time,
		Labels:    namespace.Labels,
		Name:      namespace.Name,
		Status:    (*NamespaceStatus)(&namespace.Status.Phase),
	}, nil
}

func (s Server) DeleteNamespace(ctx context.Context, request DeleteNamespaceRequestObject) (DeleteNamespaceResponseObject, error) {

	err := s.clientset.CoreV1().Namespaces().Delete(
		ctx,
		request.NamespaceName,
		metav1.DeleteOptions{},
	)

	if err != nil {
		log.Error(err.Error())
		return DeleteNamespace500Response{}, nil
	}

	return DeleteNamespace204Response{}, nil
}

func (s Server) GetNamespace(ctx context.Context, request GetNamespaceRequestObject) (GetNamespaceResponseObject, error) {
	return GetNamespace200JSONResponse{}, nil
}

func (s Server) UpdateNamespace(ctx context.Context, request UpdateNamespaceRequestObject) (UpdateNamespaceResponseObject, error) {
	return UpdateNamespace200JSONResponse{}, nil
}
