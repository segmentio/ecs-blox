package ecs

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/go-openapi/runtime"
	blox "github.com/segmentio/ecs-blox/blox/client"
	"github.com/segmentio/ecs-blox/blox/client/operations"
	"github.com/segmentio/ecs-blox/blox/models"
	"github.com/segmentio/ecs-blox/testutils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestClientSuite struct {
	suite.Suite
	ecs    *testutils.ECSAPI
	blox   *testutils.BloxTransport
	client *Client
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}

func (suite *TestClientSuite) SetupSuite() {
	suite.ecs = new(testutils.ECSAPI)
	suite.blox = new(testutils.BloxTransport)
	suite.client = &Client{
		blox: blox.New(suite.blox, nil),
		ecs:  suite.ecs,
	}
}

func (suite *TestClientSuite) TearDownTest() {
	t := suite.T()
	suite.ecs.AssertExpectations(t)
	suite.blox.AssertExpectations(t)
}

func (suite *TestClientSuite) TestImplements() {
	c := new(Client)
	suite.Implements((*ecsiface.ECSAPI)(nil), c)
}

func (suite *TestClientSuite) TestCreateCluster() {
	input := ecs.CreateClusterInput{}
	output := ecs.CreateClusterOutput{}
	suite.ecs.On("CreateCluster", &input).Once().Return(&output, nil)

	out, err := suite.client.CreateCluster(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestCreateClusterWithContext() {
	ctx := context.TODO()
	input := ecs.CreateClusterInput{}
	output := ecs.CreateClusterOutput{}
	suite.ecs.On("CreateClusterWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.CreateClusterWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestCreateClusterRequest() {
	input := ecs.CreateClusterInput{}
	request := request.Request{}
	output := ecs.CreateClusterOutput{}
	suite.ecs.On("CreateClusterRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.CreateClusterRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestCreateService() {
	input := ecs.CreateServiceInput{}
	output := ecs.CreateServiceOutput{}
	suite.ecs.On("CreateService", &input).Once().Return(&output, nil)

	out, err := suite.client.CreateService(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestCreateServiceWithContext() {
	ctx := context.TODO()
	input := ecs.CreateServiceInput{}
	output := ecs.CreateServiceOutput{}
	suite.ecs.On("CreateServiceWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.CreateServiceWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestCreateServiceRequest() {
	input := ecs.CreateServiceInput{}
	request := request.Request{}
	output := ecs.CreateServiceOutput{}
	suite.ecs.On("CreateServiceRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.CreateServiceRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteAttributes() {
	input := ecs.DeleteAttributesInput{}
	output := ecs.DeleteAttributesOutput{}
	suite.ecs.On("DeleteAttributes", &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteAttributes(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteAttributesWithContext() {
	ctx := context.TODO()
	input := ecs.DeleteAttributesInput{}
	output := ecs.DeleteAttributesOutput{}
	suite.ecs.On("DeleteAttributesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteAttributesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteAttributesRequest() {
	input := ecs.DeleteAttributesInput{}
	request := request.Request{}
	output := ecs.DeleteAttributesOutput{}
	suite.ecs.On("DeleteAttributesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DeleteAttributesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteCluster() {
	input := ecs.DeleteClusterInput{}
	output := ecs.DeleteClusterOutput{}
	suite.ecs.On("DeleteCluster", &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteCluster(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteClusterWithContext() {
	ctx := context.TODO()
	input := ecs.DeleteClusterInput{}
	output := ecs.DeleteClusterOutput{}
	suite.ecs.On("DeleteClusterWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteClusterWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteClusterRequest() {
	input := ecs.DeleteClusterInput{}
	request := request.Request{}
	output := ecs.DeleteClusterOutput{}
	suite.ecs.On("DeleteClusterRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DeleteClusterRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteService() {
	input := ecs.DeleteServiceInput{}
	output := ecs.DeleteServiceOutput{}
	suite.ecs.On("DeleteService", &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteService(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteServiceWithContext() {
	ctx := context.TODO()
	input := ecs.DeleteServiceInput{}
	output := ecs.DeleteServiceOutput{}
	suite.ecs.On("DeleteServiceWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DeleteServiceWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeleteServiceRequest() {
	input := ecs.DeleteServiceInput{}
	request := request.Request{}
	output := ecs.DeleteServiceOutput{}
	suite.ecs.On("DeleteServiceRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DeleteServiceRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterContainerInstance() {
	input := ecs.DeregisterContainerInstanceInput{}
	output := ecs.DeregisterContainerInstanceOutput{}
	suite.ecs.On("DeregisterContainerInstance", &input).Once().Return(&output, nil)

	out, err := suite.client.DeregisterContainerInstance(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterContainerInstanceWithContext() {
	ctx := context.TODO()
	input := ecs.DeregisterContainerInstanceInput{}
	output := ecs.DeregisterContainerInstanceOutput{}
	suite.ecs.On("DeregisterContainerInstanceWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DeregisterContainerInstanceWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterContainerInstanceRequest() {
	input := ecs.DeregisterContainerInstanceInput{}
	request := request.Request{}
	output := ecs.DeregisterContainerInstanceOutput{}
	suite.ecs.On("DeregisterContainerInstanceRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DeregisterContainerInstanceRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterTaskDefinition() {
	input := ecs.DeregisterTaskDefinitionInput{}
	output := ecs.DeregisterTaskDefinitionOutput{}
	suite.ecs.On("DeregisterTaskDefinition", &input).Once().Return(&output, nil)

	out, err := suite.client.DeregisterTaskDefinition(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterTaskDefinitionWithContext() {
	ctx := context.TODO()
	input := ecs.DeregisterTaskDefinitionInput{}
	output := ecs.DeregisterTaskDefinitionOutput{}
	suite.ecs.On("DeregisterTaskDefinitionWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DeregisterTaskDefinitionWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDeregisterTaskDefinitionRequest() {

	input := ecs.DeregisterTaskDefinitionInput{}
	request := request.Request{}
	output := ecs.DeregisterTaskDefinitionOutput{}
	suite.ecs.On("DeregisterTaskDefinitionRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DeregisterTaskDefinitionRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeClusters() {
	input := ecs.DescribeClustersInput{}
	output := ecs.DescribeClustersOutput{}
	suite.ecs.On("DescribeClusters", &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeClusters(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeClustersWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeClustersInput{}
	output := ecs.DescribeClustersOutput{}
	suite.ecs.On("DescribeClustersWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeClustersWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeClustersRequest() {
	input := ecs.DescribeClustersInput{}
	request := request.Request{}
	output := ecs.DescribeClustersOutput{}
	suite.ecs.On("DescribeClustersRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DescribeClustersRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeContainerInstancesBloxLayer() {
	// m := testutils.NewECS()
	// b := testutils.NewBlox()
	//
	// b.Transport.On("Submit", mock.Anything).Once().Return(nil, nil)
	//
	// input := ecs.DescribeContainerInstancesInput{
	// 	Cluster: aws.String("default"),
	// }
	// output := ecs.DescribeContainerInstancesOutput{}
	// suite.ecs.On("DescribeContainerInstances", &input).Once().Return(&output, nil)
	//
	// c := Client{ecs: m, blox: b.Client}
	// actual, err := suite.client.DescribeContainerInstances(&input)
	// assert.NoError(t, err)
	// assert.EqualValues(t, &output, actual)

}

func (suite *TestClientSuite) TestDescribeContainerInstancesWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeContainerInstancesInput{}
	output := ecs.DescribeContainerInstancesOutput{}
	suite.ecs.On("DescribeContainerInstancesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeContainerInstancesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeContainerInstancesRequest() {
	input := ecs.DescribeContainerInstancesInput{}
	request := request.Request{}
	output := ecs.DescribeContainerInstancesOutput{}
	suite.ecs.On("DescribeContainerInstancesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DescribeContainerInstancesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeServices() {
	input := ecs.DescribeServicesInput{}
	output := ecs.DescribeServicesOutput{}
	suite.ecs.On("DescribeServices", &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeServices(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeServicesWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeServicesInput{}
	output := ecs.DescribeServicesOutput{}
	suite.ecs.On("DescribeServicesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeServicesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeServicesRequest() {
	input := ecs.DescribeServicesInput{}
	request := request.Request{}
	output := ecs.DescribeServicesOutput{}
	suite.ecs.On("DescribeServicesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DescribeServicesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTaskDefinition() {
	input := ecs.DescribeTaskDefinitionInput{}
	output := ecs.DescribeTaskDefinitionOutput{}
	suite.ecs.On("DescribeTaskDefinition", &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeTaskDefinition(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTaskDefinitionWithContext() {

	ctx := context.TODO()
	input := ecs.DescribeTaskDefinitionInput{}
	output := ecs.DescribeTaskDefinitionOutput{}
	suite.ecs.On("DescribeTaskDefinitionWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeTaskDefinitionWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTaskDefinitionRequest() {
	input := ecs.DescribeTaskDefinitionInput{}
	request := request.Request{}
	output := ecs.DescribeTaskDefinitionOutput{}
	suite.ecs.On("DescribeTaskDefinitionRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DescribeTaskDefinitionRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTasks() {
	input := ecs.DescribeTasksInput{}
	output := ecs.DescribeTasksOutput{}
	suite.ecs.On("DescribeTasks", &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeTasks(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTasksWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeTasksInput{}
	output := ecs.DescribeTasksOutput{}
	suite.ecs.On("DescribeTasksWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DescribeTasksWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDescribeTasksRequest() {
	input := ecs.DescribeTasksInput{}
	request := request.Request{}
	output := ecs.DescribeTasksOutput{}
	suite.ecs.On("DescribeTasksRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DescribeTasksRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDiscoverPollEndpoint() {
	input := ecs.DiscoverPollEndpointInput{}
	output := ecs.DiscoverPollEndpointOutput{}
	suite.ecs.On("DiscoverPollEndpoint", &input).Once().Return(&output, nil)

	out, err := suite.client.DiscoverPollEndpoint(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDiscoverPollEndpointWithContext() {
	ctx := context.TODO()
	input := ecs.DiscoverPollEndpointInput{}
	output := ecs.DiscoverPollEndpointOutput{}
	suite.ecs.On("DiscoverPollEndpointWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.DiscoverPollEndpointWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestDiscoverPollEndpointRequest() {
	input := ecs.DiscoverPollEndpointInput{}
	request := request.Request{}
	output := ecs.DiscoverPollEndpointOutput{}
	suite.ecs.On("DiscoverPollEndpointRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.DiscoverPollEndpointRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListAttributes() {
	input := ecs.ListAttributesInput{}
	output := ecs.ListAttributesOutput{}
	suite.ecs.On("ListAttributes", &input).Once().Return(&output, nil)

	out, err := suite.client.ListAttributes(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListAttributesWithContext() {
	ctx := context.TODO()
	input := ecs.ListAttributesInput{}
	output := ecs.ListAttributesOutput{}
	suite.ecs.On("ListAttributesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListAttributesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListAttributesRequest() {
	input := ecs.ListAttributesInput{}
	request := request.Request{}
	output := ecs.ListAttributesOutput{}
	suite.ecs.On("ListAttributesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListAttributesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListClusters() {
	input := ecs.ListClustersInput{}
	output := ecs.ListClustersOutput{}
	suite.ecs.On("ListClusters", &input).Once().Return(&output, nil)

	out, err := suite.client.ListClusters(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListClustersWithContext() {
	ctx := context.TODO()
	input := ecs.ListClustersInput{}
	output := ecs.ListClustersOutput{}
	suite.ecs.On("ListClustersWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListClustersWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListClustersRequest() {
	input := ecs.ListClustersInput{}
	request := request.Request{}
	output := ecs.ListClustersOutput{}
	suite.ecs.On("ListClustersRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListClustersRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListClustersPages() {
	input := ecs.ListClustersInput{}
	fn := func(output *ecs.ListClustersOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListClustersPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListClustersPages(&input, fn))
}

func (suite *TestClientSuite) TestListClustersPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListClustersInput{}
	fn := func(output *ecs.ListClustersOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListClustersPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListClustersPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestListContainerInstancesBlox() {
	input := ecs.ListContainerInstancesInput{
		Cluster: aws.String("default"),
		Status:  aws.String("RUNNING"),
	}
	output := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []*string{aws.String("instance_1")},
	}

	request := runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params: &operations.ListInstancesParams{
			Cluster: aws.String("default"),
			Status:  aws.String("RUNNING"),
		},
		Reader: new(operations.ListInstancesReader),
	}
	response := operations.ListInstancesOK{
		Payload: &models.ContainerInstances{
			Items: []*models.ContainerInstance{
				&models.ContainerInstance{
					Entity: &models.ContainerInstanceDetail{
						ContainerInstanceARN: aws.String("instance_1"),
					},
				},
			},
		},
	}
	suite.blox.On("Submit", &request).Once().Return(&response, nil)

	out, err := suite.client.ListContainerInstances(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListContainerInstancesBloxError() {
	input := ecs.ListContainerInstancesInput{
		Cluster: aws.String("default"),
		Status:  aws.String("RUNNING"),
	}
	output := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []*string{aws.String("instance_1")},
	}
	suite.ecs.On("ListContainerInstances", &input).Once().Return(&output, nil)

	request := runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params: &operations.ListInstancesParams{
			Cluster: aws.String("default"),
			Status:  aws.String("RUNNING"),
		},
		Reader: new(operations.ListInstancesReader),
	}
	response := operations.NewListInstancesInternalServerError()
	suite.blox.On("Submit", &request).Once().Return(nil, response)

	out, err := suite.client.ListContainerInstances(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListContainerInstancesWithContextBlox() {
	ctx := context.TODO()
	input := ecs.ListContainerInstancesInput{
		Cluster: aws.String("default"),
		Status:  aws.String("RUNNING"),
	}
	output := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []*string{aws.String("instance_1")},
	}

	request := runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params: &operations.ListInstancesParams{
			Cluster: aws.String("default"),
			Status:  aws.String("RUNNING"),
		},
		Reader: new(operations.ListInstancesReader),
	}
	response := operations.ListInstancesOK{
		Payload: &models.ContainerInstances{
			Items: []*models.ContainerInstance{
				&models.ContainerInstance{
					Entity: &models.ContainerInstanceDetail{
						ContainerInstanceARN: aws.String("instance_1"),
					},
				},
			},
		},
	}
	suite.blox.On("Submit", &request).Once().Return(&response, nil)

	out, err := suite.client.ListContainerInstancesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListContainerInstancesWithContextBloxError() {
	ctx := context.TODO()
	input := ecs.ListContainerInstancesInput{
		Cluster: aws.String("default"),
		Status:  aws.String("RUNNING"),
	}
	output := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []*string{aws.String("instance_1")},
	}
	suite.ecs.On("ListContainerInstancesWithContext", ctx, &input).Once().Return(&output, nil)

	request := runtime.ClientOperation{
		ID:                 "ListInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params: &operations.ListInstancesParams{
			Cluster: aws.String("default"),
			Status:  aws.String("RUNNING"),
		},
		Reader: new(operations.ListInstancesReader),
	}
	response := operations.NewListInstancesInternalServerError()
	suite.blox.On("Submit", &request).Once().Return(nil, response)

	out, err := suite.client.ListContainerInstancesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListContainerInstancesRequest() {

	input := ecs.ListContainerInstancesInput{}
	request := request.Request{}
	output := ecs.ListContainerInstancesOutput{}
	suite.ecs.On("ListContainerInstancesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListContainerInstancesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListContainerInstancesPages() {
	input := ecs.ListContainerInstancesInput{}
	fn := func(output *ecs.ListContainerInstancesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListContainerInstancesPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListContainerInstancesPages(&input, fn))
}

func (suite *TestClientSuite) TestListContainerInstancesPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListContainerInstancesInput{}
	fn := func(output *ecs.ListContainerInstancesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListContainerInstancesPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListContainerInstancesPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestListServices() {
	input := ecs.ListServicesInput{}
	output := ecs.ListServicesOutput{}
	suite.ecs.On("ListServices", &input).Once().Return(&output, nil)

	out, err := suite.client.ListServices(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListServicesWithContext() {
	ctx := context.TODO()
	input := ecs.ListServicesInput{}
	output := ecs.ListServicesOutput{}
	suite.ecs.On("ListServicesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListServicesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListServicesRequest() {
	input := ecs.ListServicesInput{}
	request := request.Request{}
	output := ecs.ListServicesOutput{}
	suite.ecs.On("ListServicesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListServicesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListServicesPages() {
	input := ecs.ListServicesInput{}
	fn := func(output *ecs.ListServicesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListServicesPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListServicesPages(&input, fn))
}

func (suite *TestClientSuite) TestListServicesPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListServicesInput{}
	fn := func(output *ecs.ListServicesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListServicesPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListServicesPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestListTaskDefinitionFamilies() {
	input := ecs.ListTaskDefinitionFamiliesInput{}
	output := ecs.ListTaskDefinitionFamiliesOutput{}
	suite.ecs.On("ListTaskDefinitionFamilies", &input).Once().Return(&output, nil)

	out, err := suite.client.ListTaskDefinitionFamilies(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionFamiliesWithContext() {
	ctx := context.TODO()
	input := ecs.ListTaskDefinitionFamiliesInput{}
	output := ecs.ListTaskDefinitionFamiliesOutput{}
	suite.ecs.On("ListTaskDefinitionFamiliesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListTaskDefinitionFamiliesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionFamiliesRequest() {
	input := ecs.ListTaskDefinitionFamiliesInput{}
	request := request.Request{}
	output := ecs.ListTaskDefinitionFamiliesOutput{}
	suite.ecs.On("ListTaskDefinitionFamiliesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListTaskDefinitionFamiliesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionFamiliesPages() {
	input := ecs.ListTaskDefinitionFamiliesInput{}
	fn := func(output *ecs.ListTaskDefinitionFamiliesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTaskDefinitionFamiliesPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTaskDefinitionFamiliesPages(&input, fn))
}

func (suite *TestClientSuite) TestListTaskDefinitionFamiliesPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListTaskDefinitionFamiliesInput{}
	fn := func(output *ecs.ListTaskDefinitionFamiliesOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTaskDefinitionFamiliesPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTaskDefinitionFamiliesPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestListTaskDefinitions() {
	input := ecs.ListTaskDefinitionsInput{}
	output := ecs.ListTaskDefinitionsOutput{}
	suite.ecs.On("ListTaskDefinitions", &input).Once().Return(&output, nil)

	out, err := suite.client.ListTaskDefinitions(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionsWithContext() {
	ctx := context.TODO()
	input := ecs.ListTaskDefinitionsInput{}
	output := ecs.ListTaskDefinitionsOutput{}
	suite.ecs.On("ListTaskDefinitionsWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListTaskDefinitionsWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionsRequest() {
	input := ecs.ListTaskDefinitionsInput{}
	request := request.Request{}
	output := ecs.ListTaskDefinitionsOutput{}
	suite.ecs.On("ListTaskDefinitionsRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListTaskDefinitionsRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTaskDefinitionsPages() {
	input := ecs.ListTaskDefinitionsInput{}
	fn := func(output *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTaskDefinitionsPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTaskDefinitionsPages(&input, fn))
}

func (suite *TestClientSuite) TestListTaskDefinitionsPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListTaskDefinitionsInput{}
	fn := func(output *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTaskDefinitionsPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTaskDefinitionsPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestListTasks() {
	input := ecs.ListTasksInput{}
	output := ecs.ListTasksOutput{}
	suite.ecs.On("ListTasks", &input).Once().Return(&output, nil)

	out, err := suite.client.ListTasks(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTasksWithContext() {
	ctx := context.TODO()
	input := ecs.ListTasksInput{}
	output := ecs.ListTasksOutput{}
	suite.ecs.On("ListTasksWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.ListTasksWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTasksRequest() {
	input := ecs.ListTasksInput{}
	request := request.Request{}
	output := ecs.ListTasksOutput{}
	suite.ecs.On("ListTasksRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.ListTasksRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestListTasksPages() {
	input := ecs.ListTasksInput{}
	fn := func(output *ecs.ListTasksOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTasksPages", &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTasksPages(&input, fn))
}

func (suite *TestClientSuite) TestListTasksPagesWithContext() {
	ctx := context.TODO()
	input := ecs.ListTasksInput{}
	fn := func(output *ecs.ListTasksOutput, lastPage bool) bool {
		return lastPage
	}
	suite.ecs.On("ListTasksPagesWithContext", ctx, &input, mock.Anything).Once().Return(nil)

	suite.NoError(suite.client.ListTasksPagesWithContext(ctx, &input, fn))
}

func (suite *TestClientSuite) TestPutAttributes() {
	input := ecs.PutAttributesInput{}
	output := ecs.PutAttributesOutput{}
	suite.ecs.On("PutAttributes", &input).Once().Return(&output, nil)

	out, err := suite.client.PutAttributes(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestPutAttributesWithContext() {
	ctx := context.TODO()
	input := ecs.PutAttributesInput{}
	output := ecs.PutAttributesOutput{}
	suite.ecs.On("PutAttributesWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.PutAttributesWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestPutAttributesRequest() {
	input := ecs.PutAttributesInput{}
	request := request.Request{}
	output := ecs.PutAttributesOutput{}
	suite.ecs.On("PutAttributesRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.PutAttributesRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterContainerInstance() {
	input := ecs.RegisterContainerInstanceInput{}
	output := ecs.RegisterContainerInstanceOutput{}
	suite.ecs.On("RegisterContainerInstance", &input).Once().Return(&output, nil)

	out, err := suite.client.RegisterContainerInstance(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterContainerInstanceWithContext() {
	ctx := context.TODO()
	input := ecs.RegisterContainerInstanceInput{}
	output := ecs.RegisterContainerInstanceOutput{}
	suite.ecs.On("RegisterContainerInstanceWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.RegisterContainerInstanceWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterContainerInstanceRequest() {
	input := ecs.RegisterContainerInstanceInput{}
	request := request.Request{}
	output := ecs.RegisterContainerInstanceOutput{}
	suite.ecs.On("RegisterContainerInstanceRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.RegisterContainerInstanceRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterTaskDefinition() {
	input := ecs.RegisterTaskDefinitionInput{}
	output := ecs.RegisterTaskDefinitionOutput{}
	suite.ecs.On("RegisterTaskDefinition", &input).Once().Return(&output, nil)

	out, err := suite.client.RegisterTaskDefinition(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterTaskDefinitionWithContext() {
	ctx := context.TODO()
	input := ecs.RegisterTaskDefinitionInput{}
	output := ecs.RegisterTaskDefinitionOutput{}
	suite.ecs.On("RegisterTaskDefinitionWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.RegisterTaskDefinitionWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRegisterTaskDefinitionRequest() {
	input := ecs.RegisterTaskDefinitionInput{}
	request := request.Request{}
	output := ecs.RegisterTaskDefinitionOutput{}
	suite.ecs.On("RegisterTaskDefinitionRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.RegisterTaskDefinitionRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRunTask() {
	input := ecs.RunTaskInput{}
	output := ecs.RunTaskOutput{}
	suite.ecs.On("RunTask", &input).Once().Return(&output, nil)

	out, err := suite.client.RunTask(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRunTaskWithContext() {
	ctx := context.TODO()
	input := ecs.RunTaskInput{}
	output := ecs.RunTaskOutput{}
	suite.ecs.On("RunTaskWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.RunTaskWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestRunTaskRequest() {
	input := ecs.RunTaskInput{}
	request := request.Request{}
	output := ecs.RunTaskOutput{}
	suite.ecs.On("RunTaskRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.RunTaskRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStartTask() {
	input := ecs.StartTaskInput{}
	output := ecs.StartTaskOutput{}
	suite.ecs.On("StartTask", &input).Once().Return(&output, nil)

	out, err := suite.client.StartTask(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStartTaskWithContext() {
	ctx := context.TODO()
	input := ecs.StartTaskInput{}
	output := ecs.StartTaskOutput{}
	suite.ecs.On("StartTaskWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.StartTaskWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStartTaskRequest() {

	input := ecs.StartTaskInput{}
	request := request.Request{}
	output := ecs.StartTaskOutput{}
	suite.ecs.On("StartTaskRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.StartTaskRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStopTask() {
	input := ecs.StopTaskInput{}
	output := ecs.StopTaskOutput{}
	suite.ecs.On("StopTask", &input).Once().Return(&output, nil)

	out, err := suite.client.StopTask(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStopTaskWithContext() {
	ctx := context.TODO()
	input := ecs.StopTaskInput{}
	output := ecs.StopTaskOutput{}
	suite.ecs.On("StopTaskWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.StopTaskWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestStopTaskRequest() {
	input := ecs.StopTaskInput{}
	request := request.Request{}
	output := ecs.StopTaskOutput{}
	suite.ecs.On("StopTaskRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.StopTaskRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitContainerStateChange() {
	input := ecs.SubmitContainerStateChangeInput{}
	output := ecs.SubmitContainerStateChangeOutput{}
	suite.ecs.On("SubmitContainerStateChange", &input).Once().Return(&output, nil)

	out, err := suite.client.SubmitContainerStateChange(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitContainerStateChangeWithContext() {
	ctx := context.TODO()
	input := ecs.SubmitContainerStateChangeInput{}
	output := ecs.SubmitContainerStateChangeOutput{}
	suite.ecs.On("SubmitContainerStateChangeWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.SubmitContainerStateChangeWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitContainerStateChangeRequest() {
	input := ecs.SubmitContainerStateChangeInput{}
	request := request.Request{}
	output := ecs.SubmitContainerStateChangeOutput{}
	suite.ecs.On("SubmitContainerStateChangeRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.SubmitContainerStateChangeRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitTaskStateChange() {
	input := ecs.SubmitTaskStateChangeInput{}
	output := ecs.SubmitTaskStateChangeOutput{}
	suite.ecs.On("SubmitTaskStateChange", &input).Once().Return(&output, nil)

	out, err := suite.client.SubmitTaskStateChange(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitTaskStateChangeWithContext() {
	ctx := context.TODO()
	input := ecs.SubmitTaskStateChangeInput{}
	output := ecs.SubmitTaskStateChangeOutput{}
	suite.ecs.On("SubmitTaskStateChangeWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.SubmitTaskStateChangeWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestSubmitTaskStateChangeRequest() {
	input := ecs.SubmitTaskStateChangeInput{}
	request := request.Request{}
	output := ecs.SubmitTaskStateChangeOutput{}
	suite.ecs.On("SubmitTaskStateChangeRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.SubmitTaskStateChangeRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerAgent() {
	input := ecs.UpdateContainerAgentInput{}
	output := ecs.UpdateContainerAgentOutput{}
	suite.ecs.On("UpdateContainerAgent", &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateContainerAgent(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerAgentWithContext() {
	ctx := context.TODO()
	input := ecs.UpdateContainerAgentInput{}
	output := ecs.UpdateContainerAgentOutput{}
	suite.ecs.On("UpdateContainerAgentWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateContainerAgentWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerAgentRequest() {
	input := ecs.UpdateContainerAgentInput{}
	request := request.Request{}
	output := ecs.UpdateContainerAgentOutput{}
	suite.ecs.On("UpdateContainerAgentRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.UpdateContainerAgentRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerInstancesState() {
	input := ecs.UpdateContainerInstancesStateInput{}
	output := ecs.UpdateContainerInstancesStateOutput{}
	suite.ecs.On("UpdateContainerInstancesState", &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateContainerInstancesState(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerInstancesStateWithContext() {
	ctx := context.TODO()
	input := ecs.UpdateContainerInstancesStateInput{}
	output := ecs.UpdateContainerInstancesStateOutput{}
	suite.ecs.On("UpdateContainerInstancesStateWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateContainerInstancesStateWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateContainerInstancesStateRequest() {
	input := ecs.UpdateContainerInstancesStateInput{}
	request := request.Request{}
	output := ecs.UpdateContainerInstancesStateOutput{}
	suite.ecs.On("UpdateContainerInstancesStateRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.UpdateContainerInstancesStateRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateService() {
	input := ecs.UpdateServiceInput{}
	output := ecs.UpdateServiceOutput{}
	suite.ecs.On("UpdateService", &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateService(&input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateServiceWithContext() {
	ctx := context.TODO()
	input := ecs.UpdateServiceInput{}
	output := ecs.UpdateServiceOutput{}
	suite.ecs.On("UpdateServiceWithContext", ctx, &input).Once().Return(&output, nil)

	out, err := suite.client.UpdateServiceWithContext(ctx, &input)
	suite.NoError(err)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestUpdateServiceRequest() {
	input := ecs.UpdateServiceInput{}
	request := request.Request{}
	output := ecs.UpdateServiceOutput{}
	suite.ecs.On("UpdateServiceRequest", &input).Once().Return(&request, &output)

	req, out := suite.client.UpdateServiceRequest(&input)
	suite.EqualValues(&request, req)
	suite.EqualValues(&output, out)
}

func (suite *TestClientSuite) TestWaitUntilServicesInactive() {
	input := ecs.DescribeServicesInput{}
	suite.ecs.On("WaitUntilServicesInactive", &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilServicesInactive(&input))
}

func (suite *TestClientSuite) TestWaitUntilServicesInactiveWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeServicesInput{}
	suite.ecs.On("WaitUntilServicesInactiveWithContext", ctx, &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilServicesInactiveWithContext(ctx, &input))
}

func (suite *TestClientSuite) TestWaitUntilServicesStable() {
	input := ecs.DescribeServicesInput{}
	suite.ecs.On("WaitUntilServicesStable", &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilServicesStable(&input))
}

func (suite *TestClientSuite) TestWaitUntilServicesStableWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeServicesInput{}
	suite.ecs.On("WaitUntilServicesStableWithContext", ctx, &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilServicesStableWithContext(ctx, &input))
}

func (suite *TestClientSuite) TestWaitUntilTasksRunning() {
	input := ecs.DescribeTasksInput{}
	suite.ecs.On("WaitUntilTasksRunning", &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilTasksRunning(&input))
}

func (suite *TestClientSuite) TestWaitUntilTasksRunningWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeTasksInput{}
	suite.ecs.On("WaitUntilTasksRunningWithContext", ctx, &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilTasksRunningWithContext(ctx, &input))
}

func (suite *TestClientSuite) TestWaitUntilTasksStopped() {
	input := ecs.DescribeTasksInput{}
	suite.ecs.On("WaitUntilTasksStopped", &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilTasksStopped(&input))
}

func (suite *TestClientSuite) TestWaitUntilTasksStoppedWithContext() {
	ctx := context.TODO()
	input := ecs.DescribeTasksInput{}
	suite.ecs.On("WaitUntilTasksStoppedWithContext", ctx, &input).Once().Return(nil)

	suite.NoError(suite.client.WaitUntilTasksStoppedWithContext(ctx, &input))
}
