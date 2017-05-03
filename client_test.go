package ecs

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/segmentio/ecs-blox/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientImplementsECS(t *testing.T) {
	c := new(Client)
	assert.Implements(t, (*ecsiface.ECSAPI)(nil), c)
}

func TestClientCreateCluster(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.CreateClusterInput)
	output := new(ecs.CreateClusterOutput)
	m.On("CreateCluster", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.CreateCluster(input)

	m.AssertExpectations(t)
}

func TestClientCreateClusterWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.CreateClusterInput)
	output := new(ecs.CreateClusterOutput)
	m.On("CreateClusterWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.CreateClusterWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientCreateClusterRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.CreateClusterInput)
	request := new(request.Request)
	output := new(ecs.CreateClusterOutput)
	m.On("CreateClusterRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.CreateClusterRequest(input)

	m.AssertExpectations(t)
}

func TestClientCreateService(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.CreateServiceInput)
	output := new(ecs.CreateServiceOutput)
	m.On("CreateService", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.CreateService(input)

	m.AssertExpectations(t)
}

func TestClientCreateServiceWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.CreateServiceInput)
	output := new(ecs.CreateServiceOutput)
	m.On("CreateServiceWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.CreateServiceWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientCreateServiceRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.CreateServiceInput)
	request := new(request.Request)
	output := new(ecs.CreateServiceOutput)
	m.On("CreateServiceRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.CreateServiceRequest(input)

	m.AssertExpectations(t)
}

func TestClientDeleteAttributes(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteAttributesInput)
	output := new(ecs.DeleteAttributesOutput)
	m.On("DeleteAttributes", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteAttributes(input)

	m.AssertExpectations(t)
}

func TestClientDeleteAttributesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DeleteAttributesInput)
	output := new(ecs.DeleteAttributesOutput)
	m.On("DeleteAttributesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteAttributesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDeleteAttributesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteAttributesInput)
	request := new(request.Request)
	output := new(ecs.DeleteAttributesOutput)
	m.On("DeleteAttributesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DeleteAttributesRequest(input)

	m.AssertExpectations(t)
}

func TestClientDeleteCluster(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteClusterInput)
	output := new(ecs.DeleteClusterOutput)
	m.On("DeleteCluster", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteCluster(input)

	m.AssertExpectations(t)
}

func TestClientDeleteClusterWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DeleteClusterInput)
	output := new(ecs.DeleteClusterOutput)
	m.On("DeleteClusterWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteClusterWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDeleteClusterRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteClusterInput)
	request := new(request.Request)
	output := new(ecs.DeleteClusterOutput)
	m.On("DeleteClusterRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DeleteClusterRequest(input)

	m.AssertExpectations(t)
}

func TestClientDeleteService(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteServiceInput)
	output := new(ecs.DeleteServiceOutput)
	m.On("DeleteService", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteService(input)

	m.AssertExpectations(t)
}

func TestClientDeleteServiceWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DeleteServiceInput)
	output := new(ecs.DeleteServiceOutput)
	m.On("DeleteServiceWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeleteServiceWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDeleteServiceRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeleteServiceInput)
	request := new(request.Request)
	output := new(ecs.DeleteServiceOutput)
	m.On("DeleteServiceRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DeleteServiceRequest(input)

	m.AssertExpectations(t)
}

func TestClientDeregisterContainerInstance(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeregisterContainerInstanceInput)
	output := new(ecs.DeregisterContainerInstanceOutput)
	m.On("DeregisterContainerInstance", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeregisterContainerInstance(input)

	m.AssertExpectations(t)
}

func TestClientDeregisterContainerInstanceWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DeregisterContainerInstanceInput)
	output := new(ecs.DeregisterContainerInstanceOutput)
	m.On("DeregisterContainerInstanceWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeregisterContainerInstanceWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDeregisterContainerInstanceRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeregisterContainerInstanceInput)
	request := new(request.Request)
	output := new(ecs.DeregisterContainerInstanceOutput)
	m.On("DeregisterContainerInstanceRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DeregisterContainerInstanceRequest(input)

	m.AssertExpectations(t)
}

func TestClientDeregisterTaskDefinition(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeregisterTaskDefinitionInput)
	output := new(ecs.DeregisterTaskDefinitionOutput)
	m.On("DeregisterTaskDefinition", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeregisterTaskDefinition(input)

	m.AssertExpectations(t)
}

func TestClientDeregisterTaskDefinitionWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DeregisterTaskDefinitionInput)
	output := new(ecs.DeregisterTaskDefinitionOutput)
	m.On("DeregisterTaskDefinitionWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DeregisterTaskDefinitionWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDeregisterTaskDefinitionRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DeregisterTaskDefinitionInput)
	request := new(request.Request)
	output := new(ecs.DeregisterTaskDefinitionOutput)
	m.On("DeregisterTaskDefinitionRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DeregisterTaskDefinitionRequest(input)

	m.AssertExpectations(t)
}

func TestClientDescribeClusters(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeClustersInput)
	output := new(ecs.DescribeClustersOutput)
	m.On("DescribeClusters", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeClusters(input)

	m.AssertExpectations(t)
}

func TestClientDescribeClustersWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeClustersInput)
	output := new(ecs.DescribeClustersOutput)
	m.On("DescribeClustersWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeClustersWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDescribeClustersRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeClustersInput)
	request := new(request.Request)
	output := new(ecs.DescribeClustersOutput)
	m.On("DescribeClustersRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DescribeClustersRequest(input)

	m.AssertExpectations(t)
}

func TestClientDescribeContainerInstances(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeContainerInstancesInput)
	output := new(ecs.DescribeContainerInstancesOutput)
	m.On("DescribeContainerInstances", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeContainerInstances(input)

	m.AssertExpectations(t)
}

func TestClientDescribeContainerInstancesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeContainerInstancesInput)
	output := new(ecs.DescribeContainerInstancesOutput)
	m.On("DescribeContainerInstancesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeContainerInstancesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDescribeContainerInstancesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeContainerInstancesInput)
	request := new(request.Request)
	output := new(ecs.DescribeContainerInstancesOutput)
	m.On("DescribeContainerInstancesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DescribeContainerInstancesRequest(input)

	m.AssertExpectations(t)
}

func TestClientDescribeServices(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeServicesInput)
	output := new(ecs.DescribeServicesOutput)
	m.On("DescribeServices", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeServices(input)

	m.AssertExpectations(t)
}

func TestClientDescribeServicesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeServicesInput)
	output := new(ecs.DescribeServicesOutput)
	m.On("DescribeServicesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeServicesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDescribeServicesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeServicesInput)
	request := new(request.Request)
	output := new(ecs.DescribeServicesOutput)
	m.On("DescribeServicesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DescribeServicesRequest(input)

	m.AssertExpectations(t)
}

func TestClientDescribeTaskDefinition(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTaskDefinitionInput)
	output := new(ecs.DescribeTaskDefinitionOutput)
	m.On("DescribeTaskDefinition", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeTaskDefinition(input)

	m.AssertExpectations(t)
}

func TestClientDescribeTaskDefinitionWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeTaskDefinitionInput)
	output := new(ecs.DescribeTaskDefinitionOutput)
	m.On("DescribeTaskDefinitionWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeTaskDefinitionWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDescribeTaskDefinitionRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTaskDefinitionInput)
	request := new(request.Request)
	output := new(ecs.DescribeTaskDefinitionOutput)
	m.On("DescribeTaskDefinitionRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DescribeTaskDefinitionRequest(input)

	m.AssertExpectations(t)
}

func TestClientDescribeTasks(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTasksInput)
	output := new(ecs.DescribeTasksOutput)
	m.On("DescribeTasks", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeTasks(input)

	m.AssertExpectations(t)
}

func TestClientDescribeTasksWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeTasksInput)
	output := new(ecs.DescribeTasksOutput)
	m.On("DescribeTasksWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DescribeTasksWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDescribeTasksRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTasksInput)
	request := new(request.Request)
	output := new(ecs.DescribeTasksOutput)
	m.On("DescribeTasksRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DescribeTasksRequest(input)

	m.AssertExpectations(t)
}

func TestClientDiscoverPollEndpoint(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DiscoverPollEndpointInput)
	output := new(ecs.DiscoverPollEndpointOutput)
	m.On("DiscoverPollEndpoint", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DiscoverPollEndpoint(input)

	m.AssertExpectations(t)
}

func TestClientDiscoverPollEndpointWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DiscoverPollEndpointInput)
	output := new(ecs.DiscoverPollEndpointOutput)
	m.On("DiscoverPollEndpointWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.DiscoverPollEndpointWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientDiscoverPollEndpointRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DiscoverPollEndpointInput)
	request := new(request.Request)
	output := new(ecs.DiscoverPollEndpointOutput)
	m.On("DiscoverPollEndpointRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.DiscoverPollEndpointRequest(input)

	m.AssertExpectations(t)
}

func TestClientListAttributes(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListAttributesInput)
	output := new(ecs.ListAttributesOutput)
	m.On("ListAttributes", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListAttributes(input)

	m.AssertExpectations(t)
}

func TestClientListAttributesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListAttributesInput)
	output := new(ecs.ListAttributesOutput)
	m.On("ListAttributesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListAttributesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListAttributesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListAttributesInput)
	request := new(request.Request)
	output := new(ecs.ListAttributesOutput)
	m.On("ListAttributesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListAttributesRequest(input)

	m.AssertExpectations(t)
}

func TestClientListClusters(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListClustersInput)
	output := new(ecs.ListClustersOutput)
	m.On("ListClusters", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListClusters(input)

	m.AssertExpectations(t)
}

func TestClientListClustersWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListClustersInput)
	output := new(ecs.ListClustersOutput)
	m.On("ListClustersWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListClustersWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListClustersRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListClustersInput)
	request := new(request.Request)
	output := new(ecs.ListClustersOutput)
	m.On("ListClustersRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListClustersRequest(input)

	m.AssertExpectations(t)
}

func TestClientListClustersPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListClustersInput)
	fn := func(output *ecs.ListClustersOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListClustersPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListClustersPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListClustersPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListClustersInput)
	fn := func(output *ecs.ListClustersOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListClustersPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListClustersPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientListContainerInstances(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListContainerInstancesInput)
	output := new(ecs.ListContainerInstancesOutput)
	m.On("ListContainerInstances", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListContainerInstances(input)

	m.AssertExpectations(t)
}

func TestClientListContainerInstancesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListContainerInstancesInput)
	output := new(ecs.ListContainerInstancesOutput)
	m.On("ListContainerInstancesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListContainerInstancesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListContainerInstancesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListContainerInstancesInput)
	request := new(request.Request)
	output := new(ecs.ListContainerInstancesOutput)
	m.On("ListContainerInstancesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListContainerInstancesRequest(input)

	m.AssertExpectations(t)
}

func TestClientListContainerInstancesPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListContainerInstancesInput)
	fn := func(output *ecs.ListContainerInstancesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListContainerInstancesPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListContainerInstancesPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListContainerInstancesPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListContainerInstancesInput)
	fn := func(output *ecs.ListContainerInstancesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListContainerInstancesPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListContainerInstancesPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientListServices(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListServicesInput)
	output := new(ecs.ListServicesOutput)
	m.On("ListServices", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListServices(input)

	m.AssertExpectations(t)
}

func TestClientListServicesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListServicesInput)
	output := new(ecs.ListServicesOutput)
	m.On("ListServicesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListServicesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListServicesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListServicesInput)
	request := new(request.Request)
	output := new(ecs.ListServicesOutput)
	m.On("ListServicesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListServicesRequest(input)

	m.AssertExpectations(t)
}

func TestClientListServicesPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListServicesInput)
	fn := func(output *ecs.ListServicesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListServicesPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListServicesPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListServicesPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListServicesInput)
	fn := func(output *ecs.ListServicesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListServicesPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListServicesPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionFamilies(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionFamiliesInput)
	output := new(ecs.ListTaskDefinitionFamiliesOutput)
	m.On("ListTaskDefinitionFamilies", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionFamilies(input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionFamiliesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTaskDefinitionFamiliesInput)
	output := new(ecs.ListTaskDefinitionFamiliesOutput)
	m.On("ListTaskDefinitionFamiliesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionFamiliesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionFamiliesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionFamiliesInput)
	request := new(request.Request)
	output := new(ecs.ListTaskDefinitionFamiliesOutput)
	m.On("ListTaskDefinitionFamiliesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListTaskDefinitionFamiliesRequest(input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionFamiliesPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionFamiliesInput)
	fn := func(output *ecs.ListTaskDefinitionFamiliesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTaskDefinitionFamiliesPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionFamiliesPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionFamiliesPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTaskDefinitionFamiliesInput)
	fn := func(output *ecs.ListTaskDefinitionFamiliesOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTaskDefinitionFamiliesPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionFamiliesPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitions(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionsInput)
	output := new(ecs.ListTaskDefinitionsOutput)
	m.On("ListTaskDefinitions", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTaskDefinitions(input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionsWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTaskDefinitionsInput)
	output := new(ecs.ListTaskDefinitionsOutput)
	m.On("ListTaskDefinitionsWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionsWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionsRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionsInput)
	request := new(request.Request)
	output := new(ecs.ListTaskDefinitionsOutput)
	m.On("ListTaskDefinitionsRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListTaskDefinitionsRequest(input)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionsPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTaskDefinitionsInput)
	fn := func(output *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTaskDefinitionsPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionsPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListTaskDefinitionsPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTaskDefinitionsInput)
	fn := func(output *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTaskDefinitionsPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTaskDefinitionsPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientListTasks(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTasksInput)
	output := new(ecs.ListTasksOutput)
	m.On("ListTasks", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTasks(input)

	m.AssertExpectations(t)
}

func TestClientListTasksWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTasksInput)
	output := new(ecs.ListTasksOutput)
	m.On("ListTasksWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.ListTasksWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientListTasksRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTasksInput)
	request := new(request.Request)
	output := new(ecs.ListTasksOutput)
	m.On("ListTasksRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.ListTasksRequest(input)

	m.AssertExpectations(t)
}

func TestClientListTasksPages(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.ListTasksInput)
	fn := func(output *ecs.ListTasksOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTasksPages", input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTasksPages(input, fn)

	m.AssertExpectations(t)
}

func TestClientListTasksPagesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.ListTasksInput)
	fn := func(output *ecs.ListTasksOutput, lastPage bool) bool {
		return lastPage
	}
	m.On("ListTasksPagesWithContext", ctx, input, mock.Anything).Once().Return(nil)

	c := Client{ecs: m}
	c.ListTasksPagesWithContext(ctx, input, fn)

	m.AssertExpectations(t)
}

func TestClientPutAttributes(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.PutAttributesInput)
	output := new(ecs.PutAttributesOutput)
	m.On("PutAttributes", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.PutAttributes(input)

	m.AssertExpectations(t)
}

func TestClientPutAttributesWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.PutAttributesInput)
	output := new(ecs.PutAttributesOutput)
	m.On("PutAttributesWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.PutAttributesWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientPutAttributesRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.PutAttributesInput)
	request := new(request.Request)
	output := new(ecs.PutAttributesOutput)
	m.On("PutAttributesRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.PutAttributesRequest(input)

	m.AssertExpectations(t)
}

func TestClientRegisterContainerInstance(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RegisterContainerInstanceInput)
	output := new(ecs.RegisterContainerInstanceOutput)
	m.On("RegisterContainerInstance", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RegisterContainerInstance(input)

	m.AssertExpectations(t)
}

func TestClientRegisterContainerInstanceWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.RegisterContainerInstanceInput)
	output := new(ecs.RegisterContainerInstanceOutput)
	m.On("RegisterContainerInstanceWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RegisterContainerInstanceWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientRegisterContainerInstanceRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RegisterContainerInstanceInput)
	request := new(request.Request)
	output := new(ecs.RegisterContainerInstanceOutput)
	m.On("RegisterContainerInstanceRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.RegisterContainerInstanceRequest(input)

	m.AssertExpectations(t)
}

func TestClientRegisterTaskDefinition(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RegisterTaskDefinitionInput)
	output := new(ecs.RegisterTaskDefinitionOutput)
	m.On("RegisterTaskDefinition", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RegisterTaskDefinition(input)

	m.AssertExpectations(t)
}

func TestClientRegisterTaskDefinitionWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.RegisterTaskDefinitionInput)
	output := new(ecs.RegisterTaskDefinitionOutput)
	m.On("RegisterTaskDefinitionWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RegisterTaskDefinitionWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientRegisterTaskDefinitionRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RegisterTaskDefinitionInput)
	request := new(request.Request)
	output := new(ecs.RegisterTaskDefinitionOutput)
	m.On("RegisterTaskDefinitionRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.RegisterTaskDefinitionRequest(input)

	m.AssertExpectations(t)
}

func TestClientRunTask(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RunTaskInput)
	output := new(ecs.RunTaskOutput)
	m.On("RunTask", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RunTask(input)

	m.AssertExpectations(t)
}

func TestClientRunTaskWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.RunTaskInput)
	output := new(ecs.RunTaskOutput)
	m.On("RunTaskWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.RunTaskWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientRunTaskRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.RunTaskInput)
	request := new(request.Request)
	output := new(ecs.RunTaskOutput)
	m.On("RunTaskRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.RunTaskRequest(input)

	m.AssertExpectations(t)
}

func TestClientStartTask(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.StartTaskInput)
	output := new(ecs.StartTaskOutput)
	m.On("StartTask", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.StartTask(input)

	m.AssertExpectations(t)
}

func TestClientStartTaskWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.StartTaskInput)
	output := new(ecs.StartTaskOutput)
	m.On("StartTaskWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.StartTaskWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientStartTaskRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.StartTaskInput)
	request := new(request.Request)
	output := new(ecs.StartTaskOutput)
	m.On("StartTaskRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.StartTaskRequest(input)

	m.AssertExpectations(t)
}

func TestClientStopTask(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.StopTaskInput)
	output := new(ecs.StopTaskOutput)
	m.On("StopTask", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.StopTask(input)

	m.AssertExpectations(t)
}

func TestClientStopTaskWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.StopTaskInput)
	output := new(ecs.StopTaskOutput)
	m.On("StopTaskWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.StopTaskWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientStopTaskRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.StopTaskInput)
	request := new(request.Request)
	output := new(ecs.StopTaskOutput)
	m.On("StopTaskRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.StopTaskRequest(input)

	m.AssertExpectations(t)
}

func TestClientSubmitContainerStateChange(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.SubmitContainerStateChangeInput)
	output := new(ecs.SubmitContainerStateChangeOutput)
	m.On("SubmitContainerStateChange", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.SubmitContainerStateChange(input)

	m.AssertExpectations(t)
}

func TestClientSubmitContainerStateChangeWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.SubmitContainerStateChangeInput)
	output := new(ecs.SubmitContainerStateChangeOutput)
	m.On("SubmitContainerStateChangeWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.SubmitContainerStateChangeWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientSubmitContainerStateChangeRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.SubmitContainerStateChangeInput)
	request := new(request.Request)
	output := new(ecs.SubmitContainerStateChangeOutput)
	m.On("SubmitContainerStateChangeRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.SubmitContainerStateChangeRequest(input)

	m.AssertExpectations(t)
}

func TestClientSubmitTaskStateChange(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.SubmitTaskStateChangeInput)
	output := new(ecs.SubmitTaskStateChangeOutput)
	m.On("SubmitTaskStateChange", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.SubmitTaskStateChange(input)

	m.AssertExpectations(t)
}

func TestClientSubmitTaskStateChangeWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.SubmitTaskStateChangeInput)
	output := new(ecs.SubmitTaskStateChangeOutput)
	m.On("SubmitTaskStateChangeWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.SubmitTaskStateChangeWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientSubmitTaskStateChangeRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.SubmitTaskStateChangeInput)
	request := new(request.Request)
	output := new(ecs.SubmitTaskStateChangeOutput)
	m.On("SubmitTaskStateChangeRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.SubmitTaskStateChangeRequest(input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerAgent(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateContainerAgentInput)
	output := new(ecs.UpdateContainerAgentOutput)
	m.On("UpdateContainerAgent", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateContainerAgent(input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerAgentWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.UpdateContainerAgentInput)
	output := new(ecs.UpdateContainerAgentOutput)
	m.On("UpdateContainerAgentWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateContainerAgentWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerAgentRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateContainerAgentInput)
	request := new(request.Request)
	output := new(ecs.UpdateContainerAgentOutput)
	m.On("UpdateContainerAgentRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.UpdateContainerAgentRequest(input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerInstancesState(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateContainerInstancesStateInput)
	output := new(ecs.UpdateContainerInstancesStateOutput)
	m.On("UpdateContainerInstancesState", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateContainerInstancesState(input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerInstancesStateWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.UpdateContainerInstancesStateInput)
	output := new(ecs.UpdateContainerInstancesStateOutput)
	m.On("UpdateContainerInstancesStateWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateContainerInstancesStateWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientUpdateContainerInstancesStateRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateContainerInstancesStateInput)
	request := new(request.Request)
	output := new(ecs.UpdateContainerInstancesStateOutput)
	m.On("UpdateContainerInstancesStateRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.UpdateContainerInstancesStateRequest(input)

	m.AssertExpectations(t)
}

func TestClientUpdateService(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateServiceInput)
	output := new(ecs.UpdateServiceOutput)
	m.On("UpdateService", input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateService(input)

	m.AssertExpectations(t)
}

func TestClientUpdateServiceWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.UpdateServiceInput)
	output := new(ecs.UpdateServiceOutput)
	m.On("UpdateServiceWithContext", ctx, input).Once().Return(output, nil)

	c := Client{ecs: m}
	c.UpdateServiceWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientUpdateServiceRequest(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.UpdateServiceInput)
	request := new(request.Request)
	output := new(ecs.UpdateServiceOutput)
	m.On("UpdateServiceRequest", input).Once().Return(request, output)

	c := Client{ecs: m}
	c.UpdateServiceRequest(input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilServicesInactive(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeServicesInput)
	m.On("WaitUntilServicesInactive", input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilServicesInactive(input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilServicesInactiveWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeServicesInput)
	m.On("WaitUntilServicesInactiveWithContext", ctx, input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilServicesInactiveWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilServicesStable(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeServicesInput)
	m.On("WaitUntilServicesStable", input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilServicesStable(input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilServicesStableWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeServicesInput)
	m.On("WaitUntilServicesStableWithContext", ctx, input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilServicesStableWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilTasksRunning(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTasksInput)
	m.On("WaitUntilTasksRunning", input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilTasksRunning(input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilTasksRunningWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeTasksInput)
	m.On("WaitUntilTasksRunningWithContext", ctx, input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilTasksRunningWithContext(ctx, input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilTasksStopped(t *testing.T) {
	m := new(testutils.ECSAPI)

	input := new(ecs.DescribeTasksInput)
	m.On("WaitUntilTasksStopped", input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilTasksStopped(input)

	m.AssertExpectations(t)
}

func TestClientWaitUntilTasksStoppedWithContext(t *testing.T) {
	m := new(testutils.ECSAPI)

	ctx := context.TODO()
	input := new(ecs.DescribeTasksInput)
	m.On("WaitUntilTasksStoppedWithContext", ctx, input).Once().Return(nil)

	c := Client{ecs: m}
	c.WaitUntilTasksStoppedWithContext(ctx, input)

	m.AssertExpectations(t)
}
