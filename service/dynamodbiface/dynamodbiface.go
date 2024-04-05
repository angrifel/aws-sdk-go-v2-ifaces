package dynamodbiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type (
	Options                                   = dynamodb.Options
	BatchExecuteStatementInput                = dynamodb.BatchExecuteStatementInput
	BatchExecuteStatementOutput               = dynamodb.BatchExecuteStatementOutput
	BatchGetItemInput                         = dynamodb.BatchGetItemInput
	BatchGetItemOutput                        = dynamodb.BatchGetItemOutput
	BatchWriteItemInput                       = dynamodb.BatchWriteItemInput
	BatchWriteItemOutput                      = dynamodb.BatchWriteItemOutput
	CreateBackupInput                         = dynamodb.CreateBackupInput
	CreateBackupOutput                        = dynamodb.CreateBackupOutput
	CreateGlobalTableInput                    = dynamodb.CreateGlobalTableInput
	CreateGlobalTableOutput                   = dynamodb.CreateGlobalTableOutput
	CreateTableInput                          = dynamodb.CreateTableInput
	CreateTableOutput                         = dynamodb.CreateTableOutput
	DeleteBackupInput                         = dynamodb.DeleteBackupInput
	DeleteBackupOutput                        = dynamodb.DeleteBackupOutput
	DeleteItemInput                           = dynamodb.DeleteItemInput
	DeleteItemOutput                          = dynamodb.DeleteItemOutput
	DeleteResourcePolicyInput                 = dynamodb.DeleteResourcePolicyInput
	DeleteResourcePolicyOutput                = dynamodb.DeleteResourcePolicyOutput
	DeleteTableInput                          = dynamodb.DeleteTableInput
	DeleteTableOutput                         = dynamodb.DeleteTableOutput
	DescribeBackupInput                       = dynamodb.DescribeBackupInput
	DescribeBackupOutput                      = dynamodb.DescribeBackupOutput
	DescribeContinuousBackupsInput            = dynamodb.DescribeContinuousBackupsInput
	DescribeContinuousBackupsOutput           = dynamodb.DescribeContinuousBackupsOutput
	DescribeContributorInsightsInput          = dynamodb.DescribeContributorInsightsInput
	DescribeContributorInsightsOutput         = dynamodb.DescribeContributorInsightsOutput
	DescribeEndpointsInput                    = dynamodb.DescribeEndpointsInput
	DescribeEndpointsOutput                   = dynamodb.DescribeEndpointsOutput
	DescribeExportInput                       = dynamodb.DescribeExportInput
	DescribeExportOutput                      = dynamodb.DescribeExportOutput
	DescribeGlobalTableInput                  = dynamodb.DescribeGlobalTableInput
	DescribeGlobalTableOutput                 = dynamodb.DescribeGlobalTableOutput
	DescribeGlobalTableSettingsInput          = dynamodb.DescribeGlobalTableSettingsInput
	DescribeGlobalTableSettingsOutput         = dynamodb.DescribeGlobalTableSettingsOutput
	DescribeImportInput                       = dynamodb.DescribeImportInput
	DescribeImportOutput                      = dynamodb.DescribeImportOutput
	DescribeKinesisStreamingDestinationInput  = dynamodb.DescribeKinesisStreamingDestinationInput
	DescribeKinesisStreamingDestinationOutput = dynamodb.DescribeKinesisStreamingDestinationOutput
	DescribeLimitsInput                       = dynamodb.DescribeLimitsInput
	DescribeLimitsOutput                      = dynamodb.DescribeLimitsOutput
	DescribeTableInput                        = dynamodb.DescribeTableInput
	DescribeTableOutput                       = dynamodb.DescribeTableOutput
	DescribeTableReplicaAutoScalingInput      = dynamodb.DescribeTableReplicaAutoScalingInput
	DescribeTableReplicaAutoScalingOutput     = dynamodb.DescribeTableReplicaAutoScalingOutput
	DescribeTimeToLiveInput                   = dynamodb.DescribeTimeToLiveInput
	DescribeTimeToLiveOutput                  = dynamodb.DescribeTimeToLiveOutput
	DisableKinesisStreamingDestinationInput   = dynamodb.DisableKinesisStreamingDestinationInput
	DisableKinesisStreamingDestinationOutput  = dynamodb.DisableKinesisStreamingDestinationOutput
	EnableKinesisStreamingDestinationInput    = dynamodb.EnableKinesisStreamingDestinationInput
	EnableKinesisStreamingDestinationOutput   = dynamodb.EnableKinesisStreamingDestinationOutput
	ExecuteStatementInput                     = dynamodb.ExecuteStatementInput
	ExecuteStatementOutput                    = dynamodb.ExecuteStatementOutput
	ExecuteTransactionInput                   = dynamodb.ExecuteTransactionInput
	ExecuteTransactionOutput                  = dynamodb.ExecuteTransactionOutput
	ExportTableToPointInTimeInput             = dynamodb.ExportTableToPointInTimeInput
	ExportTableToPointInTimeOutput            = dynamodb.ExportTableToPointInTimeOutput
	GetItemInput                              = dynamodb.GetItemInput
	GetItemOutput                             = dynamodb.GetItemOutput
	GetResourcePolicyInput                    = dynamodb.GetResourcePolicyInput
	GetResourcePolicyOutput                   = dynamodb.GetResourcePolicyOutput
	ImportTableInput                          = dynamodb.ImportTableInput
	ImportTableOutput                         = dynamodb.ImportTableOutput
	ListBackupsInput                          = dynamodb.ListBackupsInput
	ListBackupsOutput                         = dynamodb.ListBackupsOutput
	ListContributorInsightsInput              = dynamodb.ListContributorInsightsInput
	ListContributorInsightsOutput             = dynamodb.ListContributorInsightsOutput
	ListExportsInput                          = dynamodb.ListExportsInput
	ListExportsOutput                         = dynamodb.ListExportsOutput
	ListGlobalTablesInput                     = dynamodb.ListGlobalTablesInput
	ListGlobalTablesOutput                    = dynamodb.ListGlobalTablesOutput
	ListImportsInput                          = dynamodb.ListImportsInput
	ListImportsOutput                         = dynamodb.ListImportsOutput
	ListTablesInput                           = dynamodb.ListTablesInput
	ListTablesOutput                          = dynamodb.ListTablesOutput
	ListTagsOfResourceInput                   = dynamodb.ListTagsOfResourceInput
	ListTagsOfResourceOutput                  = dynamodb.ListTagsOfResourceOutput
	PutItemInput                              = dynamodb.PutItemInput
	PutItemOutput                             = dynamodb.PutItemOutput
	PutResourcePolicyInput                    = dynamodb.PutResourcePolicyInput
	PutResourcePolicyOutput                   = dynamodb.PutResourcePolicyOutput
	QueryInput                                = dynamodb.QueryInput
	QueryOutput                               = dynamodb.QueryOutput
	RestoreTableFromBackupInput               = dynamodb.RestoreTableFromBackupInput
	RestoreTableFromBackupOutput              = dynamodb.RestoreTableFromBackupOutput
	RestoreTableToPointInTimeInput            = dynamodb.RestoreTableToPointInTimeInput
	RestoreTableToPointInTimeOutput           = dynamodb.RestoreTableToPointInTimeOutput
	ScanInput                                 = dynamodb.ScanInput
	ScanOutput                                = dynamodb.ScanOutput
	TagResourceInput                          = dynamodb.TagResourceInput
	TagResourceOutput                         = dynamodb.TagResourceOutput
	TransactGetItemsInput                     = dynamodb.TransactGetItemsInput
	TransactGetItemsOutput                    = dynamodb.TransactGetItemsOutput
	TransactWriteItemsInput                   = dynamodb.TransactWriteItemsInput
	TransactWriteItemsOutput                  = dynamodb.TransactWriteItemsOutput
	UntagResourceInput                        = dynamodb.UntagResourceInput
	UntagResourceOutput                       = dynamodb.UntagResourceOutput
	UpdateContinuousBackupsInput              = dynamodb.UpdateContinuousBackupsInput
	UpdateContinuousBackupsOutput             = dynamodb.UpdateContinuousBackupsOutput
	UpdateContributorInsightsInput            = dynamodb.UpdateContributorInsightsInput
	UpdateContributorInsightsOutput           = dynamodb.UpdateContributorInsightsOutput
	UpdateGlobalTableInput                    = dynamodb.UpdateGlobalTableInput
	UpdateGlobalTableOutput                   = dynamodb.UpdateGlobalTableOutput
	UpdateGlobalTableSettingsInput            = dynamodb.UpdateGlobalTableSettingsInput
	UpdateGlobalTableSettingsOutput           = dynamodb.UpdateGlobalTableSettingsOutput
	UpdateItemInput                           = dynamodb.UpdateItemInput
	UpdateItemOutput                          = dynamodb.UpdateItemOutput
	UpdateKinesisStreamingDestinationInput    = dynamodb.UpdateKinesisStreamingDestinationInput
	UpdateKinesisStreamingDestinationOutput   = dynamodb.UpdateKinesisStreamingDestinationOutput
	UpdateTableInput                          = dynamodb.UpdateTableInput
	UpdateTableOutput                         = dynamodb.UpdateTableOutput
	UpdateTableReplicaAutoScalingInput        = dynamodb.UpdateTableReplicaAutoScalingInput
	UpdateTableReplicaAutoScalingOutput       = dynamodb.UpdateTableReplicaAutoScalingOutput
	UpdateTimeToLiveInput                     = dynamodb.UpdateTimeToLiveInput
	UpdateTimeToLiveOutput                    = dynamodb.UpdateTimeToLiveOutput
)

var _ Client = &dynamodb.Client{}

type Client interface {
	BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error)
	BatchGetItem(ctx context.Context, params *BatchGetItemInput, optFns ...func(*Options)) (*BatchGetItemOutput, error)
	BatchWriteItem(ctx context.Context, params *BatchWriteItemInput, optFns ...func(*Options)) (*BatchWriteItemOutput, error)
	CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error)
	CreateGlobalTable(ctx context.Context, params *CreateGlobalTableInput, optFns ...func(*Options)) (*CreateGlobalTableOutput, error)
	CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error)
	DeleteBackup(ctx context.Context, params *DeleteBackupInput, optFns ...func(*Options)) (*DeleteBackupOutput, error)
	DeleteItem(ctx context.Context, params *DeleteItemInput, optFns ...func(*Options)) (*DeleteItemOutput, error)
	DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error)
	DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error)
	DescribeBackup(ctx context.Context, params *DescribeBackupInput, optFns ...func(*Options)) (*DescribeBackupOutput, error)
	DescribeContinuousBackups(ctx context.Context, params *DescribeContinuousBackupsInput, optFns ...func(*Options)) (*DescribeContinuousBackupsOutput, error)
	DescribeContributorInsights(ctx context.Context, params *DescribeContributorInsightsInput, optFns ...func(*Options)) (*DescribeContributorInsightsOutput, error)
	DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error)
	DescribeExport(ctx context.Context, params *DescribeExportInput, optFns ...func(*Options)) (*DescribeExportOutput, error)
	DescribeGlobalTable(ctx context.Context, params *DescribeGlobalTableInput, optFns ...func(*Options)) (*DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettings(ctx context.Context, params *DescribeGlobalTableSettingsInput, optFns ...func(*Options)) (*DescribeGlobalTableSettingsOutput, error)
	DescribeImport(ctx context.Context, params *DescribeImportInput, optFns ...func(*Options)) (*DescribeImportOutput, error)
	DescribeKinesisStreamingDestination(ctx context.Context, params *DescribeKinesisStreamingDestinationInput, optFns ...func(*Options)) (*DescribeKinesisStreamingDestinationOutput, error)
	DescribeLimits(ctx context.Context, params *DescribeLimitsInput, optFns ...func(*Options)) (*DescribeLimitsOutput, error)
	DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(ctx context.Context, params *DescribeTableReplicaAutoScalingInput, optFns ...func(*Options)) (*DescribeTableReplicaAutoScalingOutput, error)
	DescribeTimeToLive(ctx context.Context, params *DescribeTimeToLiveInput, optFns ...func(*Options)) (*DescribeTimeToLiveOutput, error)
	DisableKinesisStreamingDestination(ctx context.Context, params *DisableKinesisStreamingDestinationInput, optFns ...func(*Options)) (*DisableKinesisStreamingDestinationOutput, error)
	EnableKinesisStreamingDestination(ctx context.Context, params *EnableKinesisStreamingDestinationInput, optFns ...func(*Options)) (*EnableKinesisStreamingDestinationOutput, error)
	ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error)
	ExecuteTransaction(ctx context.Context, params *ExecuteTransactionInput, optFns ...func(*Options)) (*ExecuteTransactionOutput, error)
	ExportTableToPointInTime(ctx context.Context, params *ExportTableToPointInTimeInput, optFns ...func(*Options)) (*ExportTableToPointInTimeOutput, error)
	GetItem(ctx context.Context, params *GetItemInput, optFns ...func(*Options)) (*GetItemOutput, error)
	GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error)
	ImportTable(ctx context.Context, params *ImportTableInput, optFns ...func(*Options)) (*ImportTableOutput, error)
	ListBackups(ctx context.Context, params *ListBackupsInput, optFns ...func(*Options)) (*ListBackupsOutput, error)
	ListContributorInsights(ctx context.Context, params *ListContributorInsightsInput, optFns ...func(*Options)) (*ListContributorInsightsOutput, error)
	ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error)
	ListGlobalTables(ctx context.Context, params *ListGlobalTablesInput, optFns ...func(*Options)) (*ListGlobalTablesOutput, error)
	ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error)
	ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error)
	ListTagsOfResource(ctx context.Context, params *ListTagsOfResourceInput, optFns ...func(*Options)) (*ListTagsOfResourceOutput, error)
	PutItem(ctx context.Context, params *PutItemInput, optFns ...func(*Options)) (*PutItemOutput, error)
	PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error)
	Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error)
	RestoreTableFromBackup(ctx context.Context, params *RestoreTableFromBackupInput, optFns ...func(*Options)) (*RestoreTableFromBackupOutput, error)
	RestoreTableToPointInTime(ctx context.Context, params *RestoreTableToPointInTimeInput, optFns ...func(*Options)) (*RestoreTableToPointInTimeOutput, error)
	Scan(ctx context.Context, params *ScanInput, optFns ...func(*Options)) (*ScanOutput, error)
	TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error)
	TransactGetItems(ctx context.Context, params *TransactGetItemsInput, optFns ...func(*Options)) (*TransactGetItemsOutput, error)
	TransactWriteItems(ctx context.Context, params *TransactWriteItemsInput, optFns ...func(*Options)) (*TransactWriteItemsOutput, error)
	UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error)
	UpdateContinuousBackups(ctx context.Context, params *UpdateContinuousBackupsInput, optFns ...func(*Options)) (*UpdateContinuousBackupsOutput, error)
	UpdateContributorInsights(ctx context.Context, params *UpdateContributorInsightsInput, optFns ...func(*Options)) (*UpdateContributorInsightsOutput, error)
	UpdateGlobalTable(ctx context.Context, params *UpdateGlobalTableInput, optFns ...func(*Options)) (*UpdateGlobalTableOutput, error)
	UpdateGlobalTableSettings(ctx context.Context, params *UpdateGlobalTableSettingsInput, optFns ...func(*Options)) (*UpdateGlobalTableSettingsOutput, error)
	UpdateItem(ctx context.Context, params *UpdateItemInput, optFns ...func(*Options)) (*UpdateItemOutput, error)
	UpdateKinesisStreamingDestination(ctx context.Context, params *UpdateKinesisStreamingDestinationInput, optFns ...func(*Options)) (*UpdateKinesisStreamingDestinationOutput, error)
	UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error)
	UpdateTableReplicaAutoScaling(ctx context.Context, params *UpdateTableReplicaAutoScalingInput, optFns ...func(*Options)) (*UpdateTableReplicaAutoScalingOutput, error)
	UpdateTimeToLive(ctx context.Context, params *UpdateTimeToLiveInput, optFns ...func(*Options)) (*UpdateTimeToLiveOutput, error)
}
