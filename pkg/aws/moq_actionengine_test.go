// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package aws

import (
	"sync"

	"github.com/integr8ly/cluster-service/pkg/clusterservice"
)

var (
	lockActionEngineMockDeleteResourcesForCluster sync.RWMutex
	lockActionEngineMockGetName                   sync.RWMutex
)

// Ensure, that ActionEngineMock does implement ClusterResourceManager.
// If this is not the case, regenerate this file with moq.
var _ ClusterResourceManager = &ActionEngineMock{}

// ActionEngineMock is a mock implementation of ClusterResourceManager.
//
//     func TestSomethingThatUsesActionEngine(t *testing.T) {
//
//         // make and configure a mocked ClusterResourceManager
//         mockedActionEngine := &ActionEngineMock{
//             DeleteResourcesForClusterFunc: func(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error) {
// 	               panic("mock out the DeleteResourcesForCluster method")
//             },
//             GetNameFunc: func() string {
// 	               panic("mock out the GetName method")
//             },
//         }
//
//         // use mockedActionEngine in code that requires ClusterResourceManager
//         // and then make assertions.
//
//     }
type ActionEngineMock struct {
	// DeleteResourcesForClusterFunc mocks the DeleteResourcesForCluster method.
	DeleteResourcesForClusterFunc func(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error)

	// GetNameFunc mocks the GetName method.
	GetNameFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// DeleteResourcesForCluster holds details about calls to the DeleteResourcesForCluster method.
		DeleteResourcesForCluster []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
			// Tags is the tags argument value.
			Tags map[string]string
			// DryRun is the dryRun argument value.
			DryRun bool
		}
		// GetName holds details about calls to the GetName method.
		GetName []struct {
		}
	}
}

// DeleteResourcesForCluster calls DeleteResourcesForClusterFunc.
func (mock *ActionEngineMock) DeleteResourcesForCluster(clusterId string, tags map[string]string, dryRun bool) ([]*clusterservice.ReportItem, error) {
	if mock.DeleteResourcesForClusterFunc == nil {
		panic("ActionEngineMock.DeleteResourcesForClusterFunc: method is nil but ClusterResourceManager.DeleteResourcesForCluster was just called")
	}
	callInfo := struct {
		ClusterId string
		Tags      map[string]string
		DryRun    bool
	}{
		ClusterId: clusterId,
		Tags:      tags,
		DryRun:    dryRun,
	}
	lockActionEngineMockDeleteResourcesForCluster.Lock()
	mock.calls.DeleteResourcesForCluster = append(mock.calls.DeleteResourcesForCluster, callInfo)
	lockActionEngineMockDeleteResourcesForCluster.Unlock()
	return mock.DeleteResourcesForClusterFunc(clusterId, tags, dryRun)
}

// DeleteResourcesForClusterCalls gets all the calls that were made to DeleteResourcesForCluster.
// Check the length with:
//     len(mockedActionEngine.DeleteResourcesForClusterCalls())
func (mock *ActionEngineMock) DeleteResourcesForClusterCalls() []struct {
	ClusterId string
	Tags      map[string]string
	DryRun    bool
} {
	var calls []struct {
		ClusterId string
		Tags      map[string]string
		DryRun    bool
	}
	lockActionEngineMockDeleteResourcesForCluster.RLock()
	calls = mock.calls.DeleteResourcesForCluster
	lockActionEngineMockDeleteResourcesForCluster.RUnlock()
	return calls
}

// GetName calls GetNameFunc.
func (mock *ActionEngineMock) GetName() string {
	if mock.GetNameFunc == nil {
		panic("ActionEngineMock.GetNameFunc: method is nil but ClusterResourceManager.GetName was just called")
	}
	callInfo := struct {
	}{}
	lockActionEngineMockGetName.Lock()
	mock.calls.GetName = append(mock.calls.GetName, callInfo)
	lockActionEngineMockGetName.Unlock()
	return mock.GetNameFunc()
}

// GetNameCalls gets all the calls that were made to GetName.
// Check the length with:
//     len(mockedActionEngine.GetNameCalls())
func (mock *ActionEngineMock) GetNameCalls() []struct {
} {
	var calls []struct {
	}
	lockActionEngineMockGetName.RLock()
	calls = mock.calls.GetName
	lockActionEngineMockGetName.RUnlock()
	return calls
}
