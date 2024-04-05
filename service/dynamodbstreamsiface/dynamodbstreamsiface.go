package dynamodbstreamsiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

type (
	Options                = dynamodbstreams.Options
	DescribeStreamInput    = dynamodbstreams.DescribeStreamInput
	DescribeStreamOutput   = dynamodbstreams.DescribeStreamOutput
	GetRecordsInput        = dynamodbstreams.GetRecordsInput
	GetRecordsOutput       = dynamodbstreams.GetRecordsOutput
	GetShardIteratorInput  = dynamodbstreams.GetShardIteratorInput
	GetShardIteratorOutput = dynamodbstreams.GetShardIteratorOutput
	ListStreamsInput       = dynamodbstreams.ListStreamsInput
	ListStreamsOutput      = dynamodbstreams.ListStreamsOutput
)

var _ Client = &dynamodbstreams.Client{}

type Client interface {
	DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error)
	GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error)
	GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error)
	ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error)
}
