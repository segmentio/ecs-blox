package ecs

import (
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	blox "github.com/segmentio/ecs-blox/blox/models"
)

const bloxTimeFormat = "2006-01-02T15:04:05Z"

func ECSFailure(arn string, err error) *ecs.Failure {
	return &ecs.Failure{
		Arn:    aws.String(arn),
		Reason: aws.String(err.Error()),
	}
}

func BloxToECSContainerInstance(input *blox.ContainerInstance) *ecs.ContainerInstance {
	if input == nil {
		return nil
	}

	e := *input.Entity
	return &ecs.ContainerInstance{
		AgentConnected:       e.AgentConnected,
		AgentUpdateStatus:    aws.String(e.AgentUpdateStatus),
		Attributes:           BloxToECSContainerInstanceAttributes(e.Attributes),
		ContainerInstanceArn: e.ContainerInstanceARN,
		Ec2InstanceId:        aws.String(e.EC2InstanceID),
		RegisteredResources:  BloxToECSContainerInstanceResources(e.RegisteredResources),
		RemainingResources:   BloxToECSContainerInstanceResources(e.RemainingResources),
		Status:               e.Status,
		VersionInfo:          BloxToECSVersionInfo(e.VersionInfo),
		// TODO: PendingTasksCount does not appear in blox struct
	}
}

func BloxToECSContainerInstanceAttributes(input []*blox.ContainerInstanceAttribute) []*ecs.Attribute {
	output := make([]*ecs.Attribute, len(input))
	for x, attr := range input {
		output[x] = &ecs.Attribute{
			Name:  attr.Name,
			Value: attr.Value,
		}
	}
	return output
}

func BloxToECSContainerInstanceResources(input []*blox.ContainerInstanceResource) []*ecs.Resource {
	output := make([]*ecs.Resource, len(input))
	for x, resource := range input {
		output[x] = &ecs.Resource{
			Name: resource.Name,
			Type: resource.Type,
		}
		switch *resource.Type {
		case "INTEGER":
			if v, err := strconv.ParseInt(*resource.Value, 10, 64); err == nil {
				output[x].SetIntegerValue(v)
			}
		case "DOUBLE":
			if v, err := strconv.ParseFloat(*resource.Value, 64); err == nil {
				output[x].SetDoubleValue(v)
			}
		case "LONG":
			if v, err := strconv.ParseInt(*resource.Value, 10, 64); err == nil {
				output[x].SetLongValue(v)
			}
		case "STRINGSET":
			output[x].SetStringSetValue(BloxToECSStringSet(*resource.Value))
		}
	}
	return output
}

func BloxToECSStringSet(input string) []*string {
	set := strings.Split(input, ",")
	output := make([]*string, len(set))
	for x, str := range set {
		output[x] = aws.String(str)
	}
	return output
}

func BloxToECSVersionInfo(input *blox.ContainerInstanceVersionInfo) *ecs.VersionInfo {
	return &ecs.VersionInfo{
		AgentHash:     aws.String(input.AgentHash),
		AgentVersion:  aws.String(input.AgentVersion),
		DockerVersion: aws.String(input.DockerVersion),
	}
}

func BloxToECSTask(input *blox.Task) *ecs.Task {
	if input == nil {
		return nil
	}

	e := *input.Entity
	return &ecs.Task{
		ClusterArn:           e.ClusterARN,
		ContainerInstanceArn: e.ContainerInstanceARN,
		Containers:           BloxToECSTaskContainer(e.Containers, e.TaskARN),
		CreatedAt:            BloxToECSTime(e.CreatedAt),
		DesiredStatus:        e.DesiredStatus,
		LastStatus:           e.LastStatus,
		Overrides:            BloxToECSTaskOverrides(e.Overrides),
		StartedAt:            BloxToECSTime(&e.StartedAt),
		StartedBy:            aws.String(e.StartedBy),
		StoppedAt:            BloxToECSTime(&e.StoppedAt),
		StoppedReason:        aws.String(e.StoppedReason),
		TaskArn:              e.TaskARN,
		TaskDefinitionArn:    e.TaskDefinitionARN,

		// TODO: Group, Version do not appear in blox struct
	}
}

func BloxToECSTaskContainer(input []*blox.TaskContainer, taskArn *string) []*ecs.Container {
	output := make([]*ecs.Container, len(input))
	for x, container := range input {
		output[x] = &ecs.Container{
			ContainerArn:    container.ContainerARN,
			ExitCode:        aws.Int64(container.ExitCode),
			LastStatus:      container.LastStatus,
			Name:            container.Name,
			NetworkBindings: BloxToECSNetworkBindings(container.NetworkBindings),
			Reason:          aws.String(container.Reason),
			TaskArn:         taskArn,
		}
	}
	return output
}

func BloxToECSNetworkBindings(input []*blox.TaskNetworkBinding) []*ecs.NetworkBinding {
	output := make([]*ecs.NetworkBinding, len(input))
	for x, binding := range input {
		output[x] = &ecs.NetworkBinding{
			BindIP:        binding.BindIP,
			ContainerPort: binding.ContainerPort,
			HostPort:      binding.HostPort,
			Protocol:      aws.String(binding.Protocol),
		}
	}
	return output
}

func BloxToECSTime(input *string) *time.Time {
	if input == nil {
		return nil
	} else if *input == "" {
		return nil
	}

	t, err := time.Parse(*input, bloxTimeFormat)
	if err != nil {
		return nil
	}

	return aws.Time(t)
}

func BloxToECSTaskOverrides(input *blox.TaskOverride) *ecs.TaskOverride {
	containers := make([]*ecs.ContainerOverride, len(input.ContainerOverrides))
	for x, override := range input.ContainerOverrides {
		containers[x] = &ecs.ContainerOverride{
			Command:     aws.StringSlice(override.Command),
			Environment: BloxToECSEnvironment(override.Environment),
			Name:        override.Name,
		}
	}

	output := new(ecs.TaskOverride)
	output.SetContainerOverrides(containers)
	output.SetTaskRoleArn(input.TaskRoleArn)
	return output
}

func BloxToECSEnvironment(input []*blox.TaskEnvironment) []*ecs.KeyValuePair {
	output := make([]*ecs.KeyValuePair, len(input))
	for x, env := range input {
		output[x] = &ecs.KeyValuePair{
			Name:  env.Name,
			Value: env.Value,
		}
	}
	return output
}
