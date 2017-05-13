package ecs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/segmentio/ecs-blox/blox/client/operations"
	"github.com/segmentio/kit/log"
	"github.com/segmentio/kit/stats"
)

func (c *Client) bloxDescribeContainerInstances(input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	instances := make([]*ecs.ContainerInstance, 0, len(input.ContainerInstances))
	failures := make([]*ecs.Failure, 0, len(input.ContainerInstances))

	for _, arn := range input.ContainerInstances {
		params := operations.GetInstanceParams{Cluster: *input.Cluster, Arn: *arn}
		if instance, err := c.blox.Operations.GetInstance(&params); err != nil {
			return nil, err
		} else if instance.Error() != "" {
			failures = append(failures, ECSFailure(*arn, instance))
		} else {
			instances = append(instances, BloxToECSContainerInstance(instance.Payload))
		}
	}

	output := new(ecs.DescribeContainerInstancesOutput)
	output.SetFailures(failures)
	output.SetContainerInstances(instances)
	return output, nil
}

func (c *Client) bloxDescribeTasks(input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	tasks := make([]*ecs.Task, 0, len(input.Tasks))
	failures := make([]*ecs.Failure, 0, len(input.Tasks))

	for _, arn := range input.Tasks {
		params := operations.GetTaskParams{Cluster: *input.Cluster, Arn: *arn}
		if task, err := c.blox.Operations.GetTask(&params); err != nil {
			return nil, err
		} else if task.Error() != "" {
			failures = append(failures, ECSFailure(*arn, task))
		} else {
			tasks = append(tasks, BloxToECSTask(task.Payload))
		}
	}

	output := new(ecs.DescribeTasksOutput)
	if len(failures) > 0 {
		output.SetFailures(failures)
	}
	if len(tasks) > 0 {
		output.SetTasks(tasks)
	}
	return output, nil
}

func (c *Client) bloxListContainerInstances(ctx aws.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	c.bloxRecordUnsupportedParams("ListContainerInstances", input, map[string]interface{}{
		"NextToken":  input.NextToken,
		"Filter":     input.Filter,
		"MaxResults": input.MaxResults,
	})

	params := operations.ListInstancesParams{
		Cluster: input.Cluster,
		Status:  input.Status,
	}

	results, err := c.blox.Operations.ListInstances(&params)
	if err != nil {
		return nil, err
	}

	instances := make([]*string, len(results.Payload.Items))
	for x, item := range results.Payload.Items {
		instances[x] = item.Entity.ContainerInstanceARN
	}
	output := new(ecs.ListContainerInstancesOutput)
	if len(instances) > 0 {
		output.SetContainerInstanceArns(instances)
	}
	return output, nil
}

func (c *Client) bloxRecordFailure(method string, input interface{}, err error) {
	stats.Incr("blox.failures", fmt.Sprintf("method:%s", method))
	log.With(log.M{
		"method": method,
		"input":  input,
	}).WithError(err).Warn("blox failure detected, falling back to ECS")
}

func (c *Client) bloxRecordUnsupportedMethod(method string, input interface{}) {
	stats.Incr("blox.unsupported", fmt.Sprintf("method:%s", method))
	log.With(log.M{
		"method": method,
		"input":  input,
	}).Warn("unable to support blox for method, using ECS directly")
}

func (c *Client) bloxRecordUnsupportedParams(method string, input interface{}, params map[string]interface{}) {
	for param, val := range params {
		if val != nil {
			stats.Incr("blox.unsupported", fmt.Sprintf("method:%s", method), fmt.Sprintf("param:%s", param))
			log.With(log.M{
				"method": method,
				"input":  input,
				"param":  param,
			}).Warn("unable to support blox for param, it will be ignored")
		}
	}
}
