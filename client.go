package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	blox "github.com/blox/blox/cluster-state-service/swagger/v1/generated/client"
)

type Client struct {
	blox *blox.BloxCSS
	ecs  ecsiface.ECSAPI
}

// New creates a new client instance.
func New(config *Config) *Client {
	return &Client{
		blox: bloxClient(config),
		ecs:  ecsClient(config),
	}
}

// Create Cluster

func (c *Client) CreateCluster(input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	return c.ecs.CreateCluster(input)
}

func (c *Client) CreateClusterWithContext(ctx aws.Context, input *ecs.CreateClusterInput, options ...request.Option) (*ecs.CreateClusterOutput, error) {
	return c.ecs.CreateClusterWithContext(ctx, input, options...)
}

func (c *Client) CreateClusterRequest(input *ecs.CreateClusterInput) (*request.Request, *ecs.CreateClusterOutput) {
	return c.ecs.CreateClusterRequest(input)
}

// Create Service

func (c *Client) CreateService(input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	return c.ecs.CreateService(input)
}

func (c *Client) CreateServiceWithContext(ctx aws.Context, input *ecs.CreateServiceInput, options ...request.Option) (*ecs.CreateServiceOutput, error) {
	return c.ecs.CreateServiceWithContext(ctx, input, options...)
}

func (c *Client) CreateServiceRequest(input *ecs.CreateServiceInput) (*request.Request, *ecs.CreateServiceOutput) {
	return c.ecs.CreateServiceRequest(input)
}

// Delete Attributes

func (c *Client) DeleteAttributes(input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
	return c.ecs.DeleteAttributes(input)
}

func (c *Client) DeleteAttributesWithContext(ctx aws.Context, input *ecs.DeleteAttributesInput, options ...request.Option) (*ecs.DeleteAttributesOutput, error) {
	return c.ecs.DeleteAttributesWithContext(ctx, input, options...)
}

func (c *Client) DeleteAttributesRequest(input *ecs.DeleteAttributesInput) (*request.Request, *ecs.DeleteAttributesOutput) {
	return c.ecs.DeleteAttributesRequest(input)
}

// Delete Cluster

func (c *Client) DeleteCluster(input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	return c.ecs.DeleteCluster(input)
}

func (c *Client) DeleteClusterWithContext(ctx aws.Context, input *ecs.DeleteClusterInput, options ...request.Option) (*ecs.DeleteClusterOutput, error) {
	return c.ecs.DeleteClusterWithContext(ctx, input, options...)
}

func (c *Client) DeleteClusterRequest(input *ecs.DeleteClusterInput) (*request.Request, *ecs.DeleteClusterOutput) {
	return c.ecs.DeleteClusterRequest(input)
}

// Delete Service

func (c *Client) DeleteService(input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
	return c.ecs.DeleteService(input)
}

func (c *Client) DeleteServiceWithContext(ctx aws.Context, input *ecs.DeleteServiceInput, options ...request.Option) (*ecs.DeleteServiceOutput, error) {
	return c.ecs.DeleteServiceWithContext(ctx, input, options...)
}

func (c *Client) DeleteServiceRequest(input *ecs.DeleteServiceInput) (*request.Request, *ecs.DeleteServiceOutput) {
	return c.ecs.DeleteServiceRequest(input)
}

// Deregister Container Instance

func (c *Client) DeregisterContainerInstance(input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
	return c.ecs.DeregisterContainerInstance(input)
}

func (c *Client) DeregisterContainerInstanceWithContext(ctx aws.Context, input *ecs.DeregisterContainerInstanceInput, options ...request.Option) (*ecs.DeregisterContainerInstanceOutput, error) {
	return c.ecs.DeregisterContainerInstanceWithContext(ctx, input, options...)
}

func (c *Client) DeregisterContainerInstanceRequest(input *ecs.DeregisterContainerInstanceInput) (*request.Request, *ecs.DeregisterContainerInstanceOutput) {
	return c.ecs.DeregisterContainerInstanceRequest(input)
}

// Deregister Task Definition

func (c *Client) DeregisterTaskDefinition(input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
	return c.ecs.DeregisterTaskDefinition(input)
}

func (c *Client) DeregisterTaskDefinitionWithContext(ctx aws.Context, input *ecs.DeregisterTaskDefinitionInput, options ...request.Option) (*ecs.DeregisterTaskDefinitionOutput, error) {
	return c.ecs.DeregisterTaskDefinitionWithContext(ctx, input, options...)
}

func (c *Client) DeregisterTaskDefinitionRequest(input *ecs.DeregisterTaskDefinitionInput) (*request.Request, *ecs.DeregisterTaskDefinitionOutput) {
	return c.ecs.DeregisterTaskDefinitionRequest(input)
}

// Describe Clusters

func (c *Client) DescribeClusters(input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
	return c.ecs.DescribeClusters(input)
}

func (c *Client) DescribeClustersWithContext(ctx aws.Context, input *ecs.DescribeClustersInput, options ...request.Option) (*ecs.DescribeClustersOutput, error) {
	return c.ecs.DescribeClustersWithContext(ctx, input, options...)
}

func (c *Client) DescribeClustersRequest(input *ecs.DescribeClustersInput) (*request.Request, *ecs.DescribeClustersOutput) {
	return c.ecs.DescribeClustersRequest(input)
}

// Describe Container Instances

func (c *Client) DescribeContainerInstances(input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	return c.ecs.DescribeContainerInstances(input)
}

func (c *Client) DescribeContainerInstancesWithContext(ctx aws.Context, input *ecs.DescribeContainerInstancesInput, options ...request.Option) (*ecs.DescribeContainerInstancesOutput, error) {
	return c.ecs.DescribeContainerInstancesWithContext(ctx, input, options...)
}

func (c *Client) DescribeContainerInstancesRequest(input *ecs.DescribeContainerInstancesInput) (*request.Request, *ecs.DescribeContainerInstancesOutput) {
	return c.ecs.DescribeContainerInstancesRequest(input)
}

// Describe Services

func (c *Client) DescribeServices(input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
	return c.ecs.DescribeServices(input)
}

func (c *Client) DescribeServicesWithContext(ctx aws.Context, input *ecs.DescribeServicesInput, options ...request.Option) (*ecs.DescribeServicesOutput, error) {
	return c.ecs.DescribeServicesWithContext(ctx, input, options...)
}

func (c *Client) DescribeServicesRequest(input *ecs.DescribeServicesInput) (*request.Request, *ecs.DescribeServicesOutput) {
	return c.ecs.DescribeServicesRequest(input)
}

// Describe Task Definition

func (c *Client) DescribeTaskDefinition(input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	return c.ecs.DescribeTaskDefinition(input)
}

func (c *Client) DescribeTaskDefinitionWithContext(ctx aws.Context, input *ecs.DescribeTaskDefinitionInput, options ...request.Option) (*ecs.DescribeTaskDefinitionOutput, error) {
	return c.ecs.DescribeTaskDefinitionWithContext(ctx, input, options...)
}

func (c *Client) DescribeTaskDefinitionRequest(input *ecs.DescribeTaskDefinitionInput) (*request.Request, *ecs.DescribeTaskDefinitionOutput) {
	return c.ecs.DescribeTaskDefinitionRequest(input)
}

// Describe Tasks

func (c *Client) DescribeTasks(input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	return c.ecs.DescribeTasks(input)
}

func (c *Client) DescribeTasksWithContext(ctx aws.Context, input *ecs.DescribeTasksInput, options ...request.Option) (*ecs.DescribeTasksOutput, error) {
	return c.ecs.DescribeTasksWithContext(ctx, input, options...)
}

func (c *Client) DescribeTasksRequest(input *ecs.DescribeTasksInput) (*request.Request, *ecs.DescribeTasksOutput) {
	return c.ecs.DescribeTasksRequest(input)
}

// Discover Poll Endpoint

func (c *Client) DiscoverPollEndpoint(input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	return c.ecs.DiscoverPollEndpoint(input)
}

func (c *Client) DiscoverPollEndpointWithContext(ctx aws.Context, input *ecs.DiscoverPollEndpointInput, options ...request.Option) (*ecs.DiscoverPollEndpointOutput, error) {
	return c.ecs.DiscoverPollEndpointWithContext(ctx, input, options...)
}

func (c *Client) DiscoverPollEndpointRequest(input *ecs.DiscoverPollEndpointInput) (*request.Request, *ecs.DiscoverPollEndpointOutput) {
	return c.ecs.DiscoverPollEndpointRequest(input)
}

// List Attributes

func (c *Client) ListAttributes(input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
	return c.ecs.ListAttributes(input)
}

func (c *Client) ListAttributesWithContext(ctx aws.Context, input *ecs.ListAttributesInput, options ...request.Option) (*ecs.ListAttributesOutput, error) {
	return c.ecs.ListAttributesWithContext(ctx, input, options...)
}

func (c *Client) ListAttributesRequest(input *ecs.ListAttributesInput) (*request.Request, *ecs.ListAttributesOutput) {
	return c.ecs.ListAttributesRequest(input)
}

// List Clusters

func (c *Client) ListClusters(input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	return c.ecs.ListClusters(input)
}

func (c *Client) ListClustersWithContext(ctx aws.Context, input *ecs.ListClustersInput, options ...request.Option) (*ecs.ListClustersOutput, error) {
	return c.ecs.ListClustersWithContext(ctx, input, options...)
}

func (c *Client) ListClustersRequest(input *ecs.ListClustersInput) (*request.Request, *ecs.ListClustersOutput) {
	return c.ecs.ListClustersRequest(input)
}

func (c *Client) ListClustersPages(input *ecs.ListClustersInput, fn func(*ecs.ListClustersOutput, bool) bool) error {
	return c.ecs.ListClustersPages(input, fn)
}

func (c *Client) ListClustersPagesWithContext(ctx aws.Context, input *ecs.ListClustersInput, fn func(*ecs.ListClustersOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListClustersPagesWithContext(ctx, input, fn, options...)
}

// List Container Instances

func (c *Client) ListContainerInstances(input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	return c.ecs.ListContainerInstances(input)
}

func (c *Client) ListContainerInstancesWithContext(ctx aws.Context, input *ecs.ListContainerInstancesInput, options ...request.Option) (*ecs.ListContainerInstancesOutput, error) {
	return c.ecs.ListContainerInstancesWithContext(ctx, input, options...)
}

func (c *Client) ListContainerInstancesRequest(input *ecs.ListContainerInstancesInput) (*request.Request, *ecs.ListContainerInstancesOutput) {
	return c.ecs.ListContainerInstancesRequest(input)
}

func (c *Client) ListContainerInstancesPages(input *ecs.ListContainerInstancesInput, fn func(*ecs.ListContainerInstancesOutput, bool) bool) error {
	return c.ecs.ListContainerInstancesPages(input, fn)
}

func (c *Client) ListContainerInstancesPagesWithContext(ctx aws.Context, input *ecs.ListContainerInstancesInput, fn func(*ecs.ListContainerInstancesOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListContainerInstancesPagesWithContext(ctx, input, fn, options...)
}

// List Services

func (c *Client) ListServices(input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	return c.ecs.ListServices(input)
}

func (c *Client) ListServicesWithContext(ctx aws.Context, input *ecs.ListServicesInput, options ...request.Option) (*ecs.ListServicesOutput, error) {
	return c.ecs.ListServicesWithContext(ctx, input, options...)
}

func (c *Client) ListServicesRequest(input *ecs.ListServicesInput) (*request.Request, *ecs.ListServicesOutput) {
	return c.ecs.ListServicesRequest(input)
}

func (c *Client) ListServicesPages(input *ecs.ListServicesInput, fn func(*ecs.ListServicesOutput, bool) bool) error {
	return c.ecs.ListServicesPages(input, fn)
}

func (c *Client) ListServicesPagesWithContext(ctx aws.Context, input *ecs.ListServicesInput, fn func(*ecs.ListServicesOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListServicesPagesWithContext(ctx, input, fn, options...)
}

// List Task Definition Families

func (c *Client) ListTaskDefinitionFamilies(input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	return c.ecs.ListTaskDefinitionFamilies(input)
}

func (c *Client) ListTaskDefinitionFamiliesWithContext(ctx aws.Context, input *ecs.ListTaskDefinitionFamiliesInput, options ...request.Option) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	return c.ecs.ListTaskDefinitionFamiliesWithContext(ctx, input, options...)
}

func (c *Client) ListTaskDefinitionFamiliesRequest(input *ecs.ListTaskDefinitionFamiliesInput) (*request.Request, *ecs.ListTaskDefinitionFamiliesOutput) {
	return c.ecs.ListTaskDefinitionFamiliesRequest(input)
}

func (c *Client) ListTaskDefinitionFamiliesPages(input *ecs.ListTaskDefinitionFamiliesInput, fn func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool) error {
	return c.ecs.ListTaskDefinitionFamiliesPages(input, fn)
}

func (c *Client) ListTaskDefinitionFamiliesPagesWithContext(ctx aws.Context, input *ecs.ListTaskDefinitionFamiliesInput, fn func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListTaskDefinitionFamiliesPagesWithContext(ctx, input, fn, options...)
}

// List Task Definitions

func (c *Client) ListTaskDefinitions(input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return c.ecs.ListTaskDefinitions(input)
}

func (c *Client) ListTaskDefinitionsWithContext(ctx aws.Context, input *ecs.ListTaskDefinitionsInput, options ...request.Option) (*ecs.ListTaskDefinitionsOutput, error) {
	return c.ecs.ListTaskDefinitionsWithContext(ctx, input, options...)
}

func (c *Client) ListTaskDefinitionsRequest(input *ecs.ListTaskDefinitionsInput) (*request.Request, *ecs.ListTaskDefinitionsOutput) {
	return c.ecs.ListTaskDefinitionsRequest(input)
}

func (c *Client) ListTaskDefinitionsPages(input *ecs.ListTaskDefinitionsInput, fn func(*ecs.ListTaskDefinitionsOutput, bool) bool) error {
	return c.ecs.ListTaskDefinitionsPages(input, fn)
}

func (c *Client) ListTaskDefinitionsPagesWithContext(ctx aws.Context, input *ecs.ListTaskDefinitionsInput, fn func(*ecs.ListTaskDefinitionsOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListTaskDefinitionsPagesWithContext(ctx, input, fn, options...)
}

// List Tasks

func (c *Client) ListTasks(input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	return c.ecs.ListTasks(input)
}

func (c *Client) ListTasksWithContext(ctx aws.Context, input *ecs.ListTasksInput, options ...request.Option) (*ecs.ListTasksOutput, error) {
	return c.ecs.ListTasksWithContext(ctx, input, options...)
}

func (c *Client) ListTasksRequest(input *ecs.ListTasksInput) (*request.Request, *ecs.ListTasksOutput) {
	return c.ecs.ListTasksRequest(input)
}

func (c *Client) ListTasksPages(input *ecs.ListTasksInput, fn func(*ecs.ListTasksOutput, bool) bool) error {
	return c.ecs.ListTasksPages(input, fn)
}

func (c *Client) ListTasksPagesWithContext(ctx aws.Context, input *ecs.ListTasksInput, fn func(*ecs.ListTasksOutput, bool) bool, options ...request.Option) error {
	return c.ecs.ListTasksPagesWithContext(ctx, input, fn, options...)
}

// Put Attributes

func (c *Client) PutAttributes(input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
	return c.ecs.PutAttributes(input)
}

func (c *Client) PutAttributesWithContext(ctx aws.Context, input *ecs.PutAttributesInput, options ...request.Option) (*ecs.PutAttributesOutput, error) {
	return c.ecs.PutAttributesWithContext(ctx, input, options...)
}

func (c *Client) PutAttributesRequest(input *ecs.PutAttributesInput) (*request.Request, *ecs.PutAttributesOutput) {
	return c.ecs.PutAttributesRequest(input)
}

// Register Container Instance

func (c *Client) RegisterContainerInstance(input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	return c.ecs.RegisterContainerInstance(input)
}

func (c *Client) RegisterContainerInstanceWithContext(ctx aws.Context, input *ecs.RegisterContainerInstanceInput, options ...request.Option) (*ecs.RegisterContainerInstanceOutput, error) {
	return c.ecs.RegisterContainerInstanceWithContext(ctx, input, options...)
}

func (c *Client) RegisterContainerInstanceRequest(input *ecs.RegisterContainerInstanceInput) (*request.Request, *ecs.RegisterContainerInstanceOutput) {
	return c.ecs.RegisterContainerInstanceRequest(input)
}

// Register Task Definition

func (c *Client) RegisterTaskDefinition(input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	return c.ecs.RegisterTaskDefinition(input)
}

func (c *Client) RegisterTaskDefinitionWithContext(ctx aws.Context, input *ecs.RegisterTaskDefinitionInput, options ...request.Option) (*ecs.RegisterTaskDefinitionOutput, error) {
	return c.ecs.RegisterTaskDefinitionWithContext(ctx, input, options...)
}

func (c *Client) RegisterTaskDefinitionRequest(input *ecs.RegisterTaskDefinitionInput) (*request.Request, *ecs.RegisterTaskDefinitionOutput) {
	return c.ecs.RegisterTaskDefinitionRequest(input)
}

// Run Task

func (c *Client) RunTask(input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	return c.ecs.RunTask(input)
}

func (c *Client) RunTaskWithContext(ctx aws.Context, input *ecs.RunTaskInput, options ...request.Option) (*ecs.RunTaskOutput, error) {
	return c.ecs.RunTaskWithContext(ctx, input, options...)
}

func (c *Client) RunTaskRequest(input *ecs.RunTaskInput) (*request.Request, *ecs.RunTaskOutput) {
	return c.ecs.RunTaskRequest(input)
}

// Start Task

func (c *Client) StartTask(input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
	return c.ecs.StartTask(input)
}

func (c *Client) StartTaskWithContext(ctx aws.Context, input *ecs.StartTaskInput, options ...request.Option) (*ecs.StartTaskOutput, error) {
	return c.ecs.StartTaskWithContext(ctx, input, options...)
}

func (c *Client) StartTaskRequest(input *ecs.StartTaskInput) (*request.Request, *ecs.StartTaskOutput) {
	return c.ecs.StartTaskRequest(input)
}

// Stop Task

func (c *Client) StopTask(input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	return c.ecs.StopTask(input)
}

func (c *Client) StopTaskWithContext(ctx aws.Context, input *ecs.StopTaskInput, options ...request.Option) (*ecs.StopTaskOutput, error) {
	return c.ecs.StopTaskWithContext(ctx, input, options...)
}

func (c *Client) StopTaskRequest(input *ecs.StopTaskInput) (*request.Request, *ecs.StopTaskOutput) {
	return c.ecs.StopTaskRequest(input)
}

// Submit Container State Change

func (c *Client) SubmitContainerStateChange(input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	return c.ecs.SubmitContainerStateChange(input)
}

func (c *Client) SubmitContainerStateChangeWithContext(ctx aws.Context, input *ecs.SubmitContainerStateChangeInput, options ...request.Option) (*ecs.SubmitContainerStateChangeOutput, error) {
	return c.ecs.SubmitContainerStateChangeWithContext(ctx, input, options...)
}

func (c *Client) SubmitContainerStateChangeRequest(input *ecs.SubmitContainerStateChangeInput) (*request.Request, *ecs.SubmitContainerStateChangeOutput) {
	return c.ecs.SubmitContainerStateChangeRequest(input)
}

// Submit Task State Change

func (c *Client) SubmitTaskStateChange(input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	return c.ecs.SubmitTaskStateChange(input)
}

func (c *Client) SubmitTaskStateChangeWithContext(ctx aws.Context, input *ecs.SubmitTaskStateChangeInput, options ...request.Option) (*ecs.SubmitTaskStateChangeOutput, error) {
	return c.ecs.SubmitTaskStateChangeWithContext(ctx, input, options...)
}

func (c *Client) SubmitTaskStateChangeRequest(input *ecs.SubmitTaskStateChangeInput) (*request.Request, *ecs.SubmitTaskStateChangeOutput) {
	return c.ecs.SubmitTaskStateChangeRequest(input)
}

// Update Container Agent

func (c *Client) UpdateContainerAgent(input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
	return c.ecs.UpdateContainerAgent(input)
}

func (c *Client) UpdateContainerAgentWithContext(ctx aws.Context, input *ecs.UpdateContainerAgentInput, options ...request.Option) (*ecs.UpdateContainerAgentOutput, error) {
	return c.ecs.UpdateContainerAgentWithContext(ctx, input, options...)
}

func (c *Client) UpdateContainerAgentRequest(input *ecs.UpdateContainerAgentInput) (*request.Request, *ecs.UpdateContainerAgentOutput) {
	return c.ecs.UpdateContainerAgentRequest(input)
}

// Update Container Instances State

func (c *Client) UpdateContainerInstancesState(input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
	return c.ecs.UpdateContainerInstancesState(input)
}

func (c *Client) UpdateContainerInstancesStateWithContext(ctx aws.Context, input *ecs.UpdateContainerInstancesStateInput, options ...request.Option) (*ecs.UpdateContainerInstancesStateOutput, error) {
	return c.ecs.UpdateContainerInstancesStateWithContext(ctx, input, options...)
}

func (c *Client) UpdateContainerInstancesStateRequest(input *ecs.UpdateContainerInstancesStateInput) (*request.Request, *ecs.UpdateContainerInstancesStateOutput) {
	return c.ecs.UpdateContainerInstancesStateRequest(input)
}

// Update Service

func (c *Client) UpdateService(input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	return c.ecs.UpdateService(input)
}

func (c *Client) UpdateServiceWithContext(ctx aws.Context, input *ecs.UpdateServiceInput, options ...request.Option) (*ecs.UpdateServiceOutput, error) {
	return c.ecs.UpdateServiceWithContext(ctx, input, options...)
}

func (c *Client) UpdateServiceRequest(input *ecs.UpdateServiceInput) (*request.Request, *ecs.UpdateServiceOutput) {
	return c.ecs.UpdateServiceRequest(input)
}

// Wait Until Services Inactive

func (c *Client) WaitUntilServicesInactive(input *ecs.DescribeServicesInput) error {
	return c.ecs.WaitUntilServicesInactive(input)
}

func (c *Client) WaitUntilServicesInactiveWithContext(ctx aws.Context, input *ecs.DescribeServicesInput, options ...request.WaiterOption) error {
	return c.ecs.WaitUntilServicesInactiveWithContext(ctx, input, options...)
}

// Wait Until Services Stable

func (c *Client) WaitUntilServicesStable(input *ecs.DescribeServicesInput) error {
	return c.ecs.WaitUntilServicesStable(input)
}

func (c *Client) WaitUntilServicesStableWithContext(ctx aws.Context, input *ecs.DescribeServicesInput, options ...request.WaiterOption) error {
	return c.ecs.WaitUntilServicesStableWithContext(ctx, input, options...)
}

// Wait Until Tasks Running

func (c *Client) WaitUntilTasksRunning(input *ecs.DescribeTasksInput) error {
	return c.ecs.WaitUntilTasksRunning(input)
}

func (c *Client) WaitUntilTasksRunningWithContext(ctx aws.Context, input *ecs.DescribeTasksInput, options ...request.WaiterOption) error {
	return c.ecs.WaitUntilTasksRunningWithContext(ctx, input, options...)
}

// Wait Until Tasks Stopped

func (c *Client) WaitUntilTasksStopped(input *ecs.DescribeTasksInput) error {
	return c.ecs.WaitUntilTasksStopped(input)
}

func (c *Client) WaitUntilTasksStoppedWithContext(ctx aws.Context, input *ecs.DescribeTasksInput, options ...request.WaiterOption) error {
	return c.ecs.WaitUntilTasksStoppedWithContext(ctx, input, options...)
}

// blox helpers
//
// TODO

// client helpers

func ecsClient(config *Config) *ecs.ECS {
	return ecs.New(config.Session)
}

func bloxClient(config *Config) *blox.BloxCSS {
	t := blox.DefaultTransportConfig().WithHost(config.BloxURL)
	return blox.NewHTTPClientWithConfig(nil, t)
}
