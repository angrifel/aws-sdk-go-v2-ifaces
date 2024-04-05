package s3iface

import (
	"context"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	Options                                           = s3.Options
	PresignOptions                                    = s3.PresignOptions
	AbortMultipartUploadInput                         = s3.AbortMultipartUploadInput
	AbortMultipartUploadOutput                        = s3.AbortMultipartUploadOutput
	CompleteMultipartUploadInput                      = s3.CompleteMultipartUploadInput
	CompleteMultipartUploadOutput                     = s3.CompleteMultipartUploadOutput
	CopyObjectInput                                   = s3.CopyObjectInput
	CopyObjectOutput                                  = s3.CopyObjectOutput
	CreateBucketInput                                 = s3.CreateBucketInput
	CreateBucketOutput                                = s3.CreateBucketOutput
	CreateMultipartUploadInput                        = s3.CreateMultipartUploadInput
	CreateMultipartUploadOutput                       = s3.CreateMultipartUploadOutput
	CreateSessionInput                                = s3.CreateSessionInput
	CreateSessionOutput                               = s3.CreateSessionOutput
	DeleteBucketInput                                 = s3.DeleteBucketInput
	DeleteBucketOutput                                = s3.DeleteBucketOutput
	DeleteBucketAnalyticsConfigurationInput           = s3.DeleteBucketAnalyticsConfigurationInput
	DeleteBucketAnalyticsConfigurationOutput          = s3.DeleteBucketAnalyticsConfigurationOutput
	DeleteBucketCorsInput                             = s3.DeleteBucketCorsInput
	DeleteBucketCorsOutput                            = s3.DeleteBucketCorsOutput
	DeleteBucketEncryptionInput                       = s3.DeleteBucketEncryptionInput
	DeleteBucketEncryptionOutput                      = s3.DeleteBucketEncryptionOutput
	DeleteBucketIntelligentTieringConfigurationInput  = s3.DeleteBucketIntelligentTieringConfigurationInput
	DeleteBucketIntelligentTieringConfigurationOutput = s3.DeleteBucketIntelligentTieringConfigurationOutput
	DeleteBucketInventoryConfigurationInput           = s3.DeleteBucketInventoryConfigurationInput
	DeleteBucketInventoryConfigurationOutput          = s3.DeleteBucketInventoryConfigurationOutput
	DeleteBucketLifecycleInput                        = s3.DeleteBucketLifecycleInput
	DeleteBucketLifecycleOutput                       = s3.DeleteBucketLifecycleOutput
	DeleteBucketMetricsConfigurationInput             = s3.DeleteBucketMetricsConfigurationInput
	DeleteBucketMetricsConfigurationOutput            = s3.DeleteBucketMetricsConfigurationOutput
	DeleteBucketOwnershipControlsInput                = s3.DeleteBucketOwnershipControlsInput
	DeleteBucketOwnershipControlsOutput               = s3.DeleteBucketOwnershipControlsOutput
	DeleteBucketPolicyInput                           = s3.DeleteBucketPolicyInput
	DeleteBucketPolicyOutput                          = s3.DeleteBucketPolicyOutput
	DeleteBucketReplicationInput                      = s3.DeleteBucketReplicationInput
	DeleteBucketReplicationOutput                     = s3.DeleteBucketReplicationOutput
	DeleteBucketTaggingInput                          = s3.DeleteBucketTaggingInput
	DeleteBucketTaggingOutput                         = s3.DeleteBucketTaggingOutput
	DeleteBucketWebsiteInput                          = s3.DeleteBucketWebsiteInput
	DeleteBucketWebsiteOutput                         = s3.DeleteBucketWebsiteOutput
	DeleteObjectInput                                 = s3.DeleteObjectInput
	DeleteObjectOutput                                = s3.DeleteObjectOutput
	DeleteObjectTaggingInput                          = s3.DeleteObjectTaggingInput
	DeleteObjectTaggingOutput                         = s3.DeleteObjectTaggingOutput
	DeleteObjectsInput                                = s3.DeleteObjectsInput
	DeleteObjectsOutput                               = s3.DeleteObjectsOutput
	DeletePublicAccessBlockInput                      = s3.DeletePublicAccessBlockInput
	DeletePublicAccessBlockOutput                     = s3.DeletePublicAccessBlockOutput
	GetBucketAccelerateConfigurationInput             = s3.GetBucketAccelerateConfigurationInput
	GetBucketAccelerateConfigurationOutput            = s3.GetBucketAccelerateConfigurationOutput
	GetBucketAclInput                                 = s3.GetBucketAclInput
	GetBucketAclOutput                                = s3.GetBucketAclOutput
	GetBucketAnalyticsConfigurationInput              = s3.GetBucketAnalyticsConfigurationInput
	GetBucketAnalyticsConfigurationOutput             = s3.GetBucketAnalyticsConfigurationOutput
	GetBucketCorsInput                                = s3.GetBucketCorsInput
	GetBucketCorsOutput                               = s3.GetBucketCorsOutput
	GetBucketEncryptionInput                          = s3.GetBucketEncryptionInput
	GetBucketEncryptionOutput                         = s3.GetBucketEncryptionOutput
	GetBucketIntelligentTieringConfigurationInput     = s3.GetBucketIntelligentTieringConfigurationInput
	GetBucketIntelligentTieringConfigurationOutput    = s3.GetBucketIntelligentTieringConfigurationOutput
	GetBucketInventoryConfigurationInput              = s3.GetBucketInventoryConfigurationInput
	GetBucketInventoryConfigurationOutput             = s3.GetBucketInventoryConfigurationOutput
	GetBucketLifecycleConfigurationInput              = s3.GetBucketLifecycleConfigurationInput
	GetBucketLifecycleConfigurationOutput             = s3.GetBucketLifecycleConfigurationOutput
	GetBucketLocationInput                            = s3.GetBucketLocationInput
	GetBucketLocationOutput                           = s3.GetBucketLocationOutput
	GetBucketLoggingInput                             = s3.GetBucketLoggingInput
	GetBucketLoggingOutput                            = s3.GetBucketLoggingOutput
	GetBucketMetricsConfigurationInput                = s3.GetBucketMetricsConfigurationInput
	GetBucketMetricsConfigurationOutput               = s3.GetBucketMetricsConfigurationOutput
	GetBucketNotificationConfigurationInput           = s3.GetBucketNotificationConfigurationInput
	GetBucketNotificationConfigurationOutput          = s3.GetBucketNotificationConfigurationOutput
	GetBucketOwnershipControlsInput                   = s3.GetBucketOwnershipControlsInput
	GetBucketOwnershipControlsOutput                  = s3.GetBucketOwnershipControlsOutput
	GetBucketPolicyInput                              = s3.GetBucketPolicyInput
	GetBucketPolicyOutput                             = s3.GetBucketPolicyOutput
	GetBucketPolicyStatusInput                        = s3.GetBucketPolicyStatusInput
	GetBucketPolicyStatusOutput                       = s3.GetBucketPolicyStatusOutput
	GetBucketReplicationInput                         = s3.GetBucketReplicationInput
	GetBucketReplicationOutput                        = s3.GetBucketReplicationOutput
	GetBucketRequestPaymentInput                      = s3.GetBucketRequestPaymentInput
	GetBucketRequestPaymentOutput                     = s3.GetBucketRequestPaymentOutput
	GetBucketTaggingInput                             = s3.GetBucketTaggingInput
	GetBucketTaggingOutput                            = s3.GetBucketTaggingOutput
	GetBucketVersioningInput                          = s3.GetBucketVersioningInput
	GetBucketVersioningOutput                         = s3.GetBucketVersioningOutput
	GetBucketWebsiteInput                             = s3.GetBucketWebsiteInput
	GetBucketWebsiteOutput                            = s3.GetBucketWebsiteOutput
	GetObjectInput                                    = s3.GetObjectInput
	GetObjectOutput                                   = s3.GetObjectOutput
	GetObjectAclInput                                 = s3.GetObjectAclInput
	GetObjectAclOutput                                = s3.GetObjectAclOutput
	GetObjectAttributesInput                          = s3.GetObjectAttributesInput
	GetObjectAttributesOutput                         = s3.GetObjectAttributesOutput
	GetObjectLegalHoldInput                           = s3.GetObjectLegalHoldInput
	GetObjectLegalHoldOutput                          = s3.GetObjectLegalHoldOutput
	GetObjectLockConfigurationInput                   = s3.GetObjectLockConfigurationInput
	GetObjectLockConfigurationOutput                  = s3.GetObjectLockConfigurationOutput
	GetObjectRetentionInput                           = s3.GetObjectRetentionInput
	GetObjectRetentionOutput                          = s3.GetObjectRetentionOutput
	GetObjectTaggingInput                             = s3.GetObjectTaggingInput
	GetObjectTaggingOutput                            = s3.GetObjectTaggingOutput
	GetObjectTorrentInput                             = s3.GetObjectTorrentInput
	GetObjectTorrentOutput                            = s3.GetObjectTorrentOutput
	GetPublicAccessBlockInput                         = s3.GetPublicAccessBlockInput
	GetPublicAccessBlockOutput                        = s3.GetPublicAccessBlockOutput
	HeadBucketInput                                   = s3.HeadBucketInput
	HeadBucketOutput                                  = s3.HeadBucketOutput
	HeadObjectInput                                   = s3.HeadObjectInput
	HeadObjectOutput                                  = s3.HeadObjectOutput
	ListBucketAnalyticsConfigurationsInput            = s3.ListBucketAnalyticsConfigurationsInput
	ListBucketAnalyticsConfigurationsOutput           = s3.ListBucketAnalyticsConfigurationsOutput
	ListBucketIntelligentTieringConfigurationsInput   = s3.ListBucketIntelligentTieringConfigurationsInput
	ListBucketIntelligentTieringConfigurationsOutput  = s3.ListBucketIntelligentTieringConfigurationsOutput
	ListBucketInventoryConfigurationsInput            = s3.ListBucketInventoryConfigurationsInput
	ListBucketInventoryConfigurationsOutput           = s3.ListBucketInventoryConfigurationsOutput
	ListBucketMetricsConfigurationsInput              = s3.ListBucketMetricsConfigurationsInput
	ListBucketMetricsConfigurationsOutput             = s3.ListBucketMetricsConfigurationsOutput
	ListBucketsInput                                  = s3.ListBucketsInput
	ListBucketsOutput                                 = s3.ListBucketsOutput
	ListDirectoryBucketsInput                         = s3.ListDirectoryBucketsInput
	ListDirectoryBucketsOutput                        = s3.ListDirectoryBucketsOutput
	ListMultipartUploadsInput                         = s3.ListMultipartUploadsInput
	ListMultipartUploadsOutput                        = s3.ListMultipartUploadsOutput
	ListObjectVersionsInput                           = s3.ListObjectVersionsInput
	ListObjectVersionsOutput                          = s3.ListObjectVersionsOutput
	ListObjectsInput                                  = s3.ListObjectsInput
	ListObjectsOutput                                 = s3.ListObjectsOutput
	ListObjectsV2Input                                = s3.ListObjectsV2Input
	ListObjectsV2Output                               = s3.ListObjectsV2Output
	ListPartsInput                                    = s3.ListPartsInput
	ListPartsOutput                                   = s3.ListPartsOutput
	PutBucketAccelerateConfigurationInput             = s3.PutBucketAccelerateConfigurationInput
	PutBucketAccelerateConfigurationOutput            = s3.PutBucketAccelerateConfigurationOutput
	PutBucketAclInput                                 = s3.PutBucketAclInput
	PutBucketAclOutput                                = s3.PutBucketAclOutput
	PutBucketAnalyticsConfigurationInput              = s3.PutBucketAnalyticsConfigurationInput
	PutBucketAnalyticsConfigurationOutput             = s3.PutBucketAnalyticsConfigurationOutput
	PutBucketCorsInput                                = s3.PutBucketCorsInput
	PutBucketCorsOutput                               = s3.PutBucketCorsOutput
	PutBucketEncryptionInput                          = s3.PutBucketEncryptionInput
	PutBucketEncryptionOutput                         = s3.PutBucketEncryptionOutput
	PutBucketIntelligentTieringConfigurationInput     = s3.PutBucketIntelligentTieringConfigurationInput
	PutBucketIntelligentTieringConfigurationOutput    = s3.PutBucketIntelligentTieringConfigurationOutput
	PutBucketInventoryConfigurationInput              = s3.PutBucketInventoryConfigurationInput
	PutBucketInventoryConfigurationOutput             = s3.PutBucketInventoryConfigurationOutput
	PutBucketLifecycleConfigurationInput              = s3.PutBucketLifecycleConfigurationInput
	PutBucketLifecycleConfigurationOutput             = s3.PutBucketLifecycleConfigurationOutput
	PutBucketLoggingInput                             = s3.PutBucketLoggingInput
	PutBucketLoggingOutput                            = s3.PutBucketLoggingOutput
	PutBucketMetricsConfigurationInput                = s3.PutBucketMetricsConfigurationInput
	PutBucketMetricsConfigurationOutput               = s3.PutBucketMetricsConfigurationOutput
	PutBucketNotificationConfigurationInput           = s3.PutBucketNotificationConfigurationInput
	PutBucketNotificationConfigurationOutput          = s3.PutBucketNotificationConfigurationOutput
	PutBucketOwnershipControlsInput                   = s3.PutBucketOwnershipControlsInput
	PutBucketOwnershipControlsOutput                  = s3.PutBucketOwnershipControlsOutput
	PutBucketPolicyInput                              = s3.PutBucketPolicyInput
	PutBucketPolicyOutput                             = s3.PutBucketPolicyOutput
	PutBucketReplicationInput                         = s3.PutBucketReplicationInput
	PutBucketReplicationOutput                        = s3.PutBucketReplicationOutput
	PutBucketRequestPaymentInput                      = s3.PutBucketRequestPaymentInput
	PutBucketRequestPaymentOutput                     = s3.PutBucketRequestPaymentOutput
	PutBucketTaggingInput                             = s3.PutBucketTaggingInput
	PutBucketTaggingOutput                            = s3.PutBucketTaggingOutput
	PutBucketVersioningInput                          = s3.PutBucketVersioningInput
	PutBucketVersioningOutput                         = s3.PutBucketVersioningOutput
	PutBucketWebsiteInput                             = s3.PutBucketWebsiteInput
	PutBucketWebsiteOutput                            = s3.PutBucketWebsiteOutput
	PutObjectInput                                    = s3.PutObjectInput
	PutObjectOutput                                   = s3.PutObjectOutput
	PutObjectAclInput                                 = s3.PutObjectAclInput
	PutObjectAclOutput                                = s3.PutObjectAclOutput
	PutObjectLegalHoldInput                           = s3.PutObjectLegalHoldInput
	PutObjectLegalHoldOutput                          = s3.PutObjectLegalHoldOutput
	PutObjectLockConfigurationInput                   = s3.PutObjectLockConfigurationInput
	PutObjectLockConfigurationOutput                  = s3.PutObjectLockConfigurationOutput
	PutObjectRetentionInput                           = s3.PutObjectRetentionInput
	PutObjectRetentionOutput                          = s3.PutObjectRetentionOutput
	PutObjectTaggingInput                             = s3.PutObjectTaggingInput
	PutObjectTaggingOutput                            = s3.PutObjectTaggingOutput
	PutPublicAccessBlockInput                         = s3.PutPublicAccessBlockInput
	PutPublicAccessBlockOutput                        = s3.PutPublicAccessBlockOutput
	RestoreObjectInput                                = s3.RestoreObjectInput
	RestoreObjectOutput                               = s3.RestoreObjectOutput
	SelectObjectContentInput                          = s3.SelectObjectContentInput
	SelectObjectContentOutput                         = s3.SelectObjectContentOutput
	UploadPartInput                                   = s3.UploadPartInput
	UploadPartOutput                                  = s3.UploadPartOutput
	UploadPartCopyInput                               = s3.UploadPartCopyInput
	UploadPartCopyOutput                              = s3.UploadPartCopyOutput
	WriteGetObjectResponseInput                       = s3.WriteGetObjectResponseInput
	WriteGetObjectResponseOutput                      = s3.WriteGetObjectResponseOutput
)

var _ Client = &s3.Client{}
var _ PresignClient = &s3.PresignClient{}

type Client interface {
	AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error)
	CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error)
	CopyObject(ctx context.Context, params *CopyObjectInput, optFns ...func(*Options)) (*CopyObjectOutput, error)
	CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error)
	CreateMultipartUpload(ctx context.Context, params *CreateMultipartUploadInput, optFns ...func(*Options)) (*CreateMultipartUploadOutput, error)
	CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error)
	DeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*Options)) (*DeleteBucketOutput, error)
	DeleteBucketAnalyticsConfiguration(ctx context.Context, params *DeleteBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketCors(ctx context.Context, params *DeleteBucketCorsInput, optFns ...func(*Options)) (*DeleteBucketCorsOutput, error)
	DeleteBucketEncryption(ctx context.Context, params *DeleteBucketEncryptionInput, optFns ...func(*Options)) (*DeleteBucketEncryptionOutput, error)
	DeleteBucketIntelligentTieringConfiguration(ctx context.Context, params *DeleteBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*DeleteBucketIntelligentTieringConfigurationOutput, error)
	DeleteBucketInventoryConfiguration(ctx context.Context, params *DeleteBucketInventoryConfigurationInput, optFns ...func(*Options)) (*DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketLifecycle(ctx context.Context, params *DeleteBucketLifecycleInput, optFns ...func(*Options)) (*DeleteBucketLifecycleOutput, error)
	DeleteBucketMetricsConfiguration(ctx context.Context, params *DeleteBucketMetricsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketOwnershipControls(ctx context.Context, params *DeleteBucketOwnershipControlsInput, optFns ...func(*Options)) (*DeleteBucketOwnershipControlsOutput, error)
	DeleteBucketPolicy(ctx context.Context, params *DeleteBucketPolicyInput, optFns ...func(*Options)) (*DeleteBucketPolicyOutput, error)
	DeleteBucketReplication(ctx context.Context, params *DeleteBucketReplicationInput, optFns ...func(*Options)) (*DeleteBucketReplicationOutput, error)
	DeleteBucketTagging(ctx context.Context, params *DeleteBucketTaggingInput, optFns ...func(*Options)) (*DeleteBucketTaggingOutput, error)
	DeleteBucketWebsite(ctx context.Context, params *DeleteBucketWebsiteInput, optFns ...func(*Options)) (*DeleteBucketWebsiteOutput, error)
	DeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*Options)) (*DeleteObjectOutput, error)
	DeleteObjectTagging(ctx context.Context, params *DeleteObjectTaggingInput, optFns ...func(*Options)) (*DeleteObjectTaggingOutput, error)
	DeleteObjects(ctx context.Context, params *DeleteObjectsInput, optFns ...func(*Options)) (*DeleteObjectsOutput, error)
	DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error)
	GetBucketAccelerateConfiguration(ctx context.Context, params *GetBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*GetBucketAccelerateConfigurationOutput, error)
	GetBucketAcl(ctx context.Context, params *GetBucketAclInput, optFns ...func(*Options)) (*GetBucketAclOutput, error)
	GetBucketAnalyticsConfiguration(ctx context.Context, params *GetBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*GetBucketAnalyticsConfigurationOutput, error)
	GetBucketCors(ctx context.Context, params *GetBucketCorsInput, optFns ...func(*Options)) (*GetBucketCorsOutput, error)
	GetBucketEncryption(ctx context.Context, params *GetBucketEncryptionInput, optFns ...func(*Options)) (*GetBucketEncryptionOutput, error)
	GetBucketIntelligentTieringConfiguration(ctx context.Context, params *GetBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*GetBucketIntelligentTieringConfigurationOutput, error)
	GetBucketInventoryConfiguration(ctx context.Context, params *GetBucketInventoryConfigurationInput, optFns ...func(*Options)) (*GetBucketInventoryConfigurationOutput, error)
	GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error)
	GetBucketLocation(ctx context.Context, params *GetBucketLocationInput, optFns ...func(*Options)) (*GetBucketLocationOutput, error)
	GetBucketLogging(ctx context.Context, params *GetBucketLoggingInput, optFns ...func(*Options)) (*GetBucketLoggingOutput, error)
	GetBucketMetricsConfiguration(ctx context.Context, params *GetBucketMetricsConfigurationInput, optFns ...func(*Options)) (*GetBucketMetricsConfigurationOutput, error)
	GetBucketNotificationConfiguration(ctx context.Context, params *GetBucketNotificationConfigurationInput, optFns ...func(*Options)) (*GetBucketNotificationConfigurationOutput, error)
	GetBucketOwnershipControls(ctx context.Context, params *GetBucketOwnershipControlsInput, optFns ...func(*Options)) (*GetBucketOwnershipControlsOutput, error)
	GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error)
	GetBucketPolicyStatus(ctx context.Context, params *GetBucketPolicyStatusInput, optFns ...func(*Options)) (*GetBucketPolicyStatusOutput, error)
	GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error)
	GetBucketRequestPayment(ctx context.Context, params *GetBucketRequestPaymentInput, optFns ...func(*Options)) (*GetBucketRequestPaymentOutput, error)
	GetBucketTagging(ctx context.Context, params *GetBucketTaggingInput, optFns ...func(*Options)) (*GetBucketTaggingOutput, error)
	GetBucketVersioning(ctx context.Context, params *GetBucketVersioningInput, optFns ...func(*Options)) (*GetBucketVersioningOutput, error)
	GetBucketWebsite(ctx context.Context, params *GetBucketWebsiteInput, optFns ...func(*Options)) (*GetBucketWebsiteOutput, error)
	GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error)
	GetObjectAcl(ctx context.Context, params *GetObjectAclInput, optFns ...func(*Options)) (*GetObjectAclOutput, error)
	GetObjectAttributes(ctx context.Context, params *GetObjectAttributesInput, optFns ...func(*Options)) (*GetObjectAttributesOutput, error)
	GetObjectLegalHold(ctx context.Context, params *GetObjectLegalHoldInput, optFns ...func(*Options)) (*GetObjectLegalHoldOutput, error)
	GetObjectLockConfiguration(ctx context.Context, params *GetObjectLockConfigurationInput, optFns ...func(*Options)) (*GetObjectLockConfigurationOutput, error)
	GetObjectRetention(ctx context.Context, params *GetObjectRetentionInput, optFns ...func(*Options)) (*GetObjectRetentionOutput, error)
	GetObjectTagging(ctx context.Context, params *GetObjectTaggingInput, optFns ...func(*Options)) (*GetObjectTaggingOutput, error)
	GetObjectTorrent(ctx context.Context, params *GetObjectTorrentInput, optFns ...func(*Options)) (*GetObjectTorrentOutput, error)
	GetPublicAccessBlock(ctx context.Context, params *GetPublicAccessBlockInput, optFns ...func(*Options)) (*GetPublicAccessBlockOutput, error)
	HeadBucket(ctx context.Context, params *HeadBucketInput, optFns ...func(*Options)) (*HeadBucketOutput, error)
	HeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*Options)) (*HeadObjectOutput, error)
	ListBucketAnalyticsConfigurations(ctx context.Context, params *ListBucketAnalyticsConfigurationsInput, optFns ...func(*Options)) (*ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketIntelligentTieringConfigurations(ctx context.Context, params *ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*Options)) (*ListBucketIntelligentTieringConfigurationsOutput, error)
	ListBucketInventoryConfigurations(ctx context.Context, params *ListBucketInventoryConfigurationsInput, optFns ...func(*Options)) (*ListBucketInventoryConfigurationsOutput, error)
	ListBucketMetricsConfigurations(ctx context.Context, params *ListBucketMetricsConfigurationsInput, optFns ...func(*Options)) (*ListBucketMetricsConfigurationsOutput, error)
	ListBuckets(ctx context.Context, params *ListBucketsInput, optFns ...func(*Options)) (*ListBucketsOutput, error)
	ListDirectoryBuckets(ctx context.Context, params *ListDirectoryBucketsInput, optFns ...func(*Options)) (*ListDirectoryBucketsOutput, error)
	ListMultipartUploads(ctx context.Context, params *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error)
	ListObjectVersions(ctx context.Context, params *ListObjectVersionsInput, optFns ...func(*Options)) (*ListObjectVersionsOutput, error)
	ListObjects(ctx context.Context, params *ListObjectsInput, optFns ...func(*Options)) (*ListObjectsOutput, error)
	ListObjectsV2(ctx context.Context, params *ListObjectsV2Input, optFns ...func(*Options)) (*ListObjectsV2Output, error)
	ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error)
	PutBucketAccelerateConfiguration(ctx context.Context, params *PutBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*PutBucketAccelerateConfigurationOutput, error)
	PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error)
	PutBucketAnalyticsConfiguration(ctx context.Context, params *PutBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*PutBucketAnalyticsConfigurationOutput, error)
	PutBucketCors(ctx context.Context, params *PutBucketCorsInput, optFns ...func(*Options)) (*PutBucketCorsOutput, error)
	PutBucketEncryption(ctx context.Context, params *PutBucketEncryptionInput, optFns ...func(*Options)) (*PutBucketEncryptionOutput, error)
	PutBucketIntelligentTieringConfiguration(ctx context.Context, params *PutBucketIntelligentTieringConfigurationInput, optFns ...func(*Options)) (*PutBucketIntelligentTieringConfigurationOutput, error)
	PutBucketInventoryConfiguration(ctx context.Context, params *PutBucketInventoryConfigurationInput, optFns ...func(*Options)) (*PutBucketInventoryConfigurationOutput, error)
	PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error)
	PutBucketLogging(ctx context.Context, params *PutBucketLoggingInput, optFns ...func(*Options)) (*PutBucketLoggingOutput, error)
	PutBucketMetricsConfiguration(ctx context.Context, params *PutBucketMetricsConfigurationInput, optFns ...func(*Options)) (*PutBucketMetricsConfigurationOutput, error)
	PutBucketNotificationConfiguration(ctx context.Context, params *PutBucketNotificationConfigurationInput, optFns ...func(*Options)) (*PutBucketNotificationConfigurationOutput, error)
	PutBucketOwnershipControls(ctx context.Context, params *PutBucketOwnershipControlsInput, optFns ...func(*Options)) (*PutBucketOwnershipControlsOutput, error)
	PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error)
	PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error)
	PutBucketRequestPayment(ctx context.Context, params *PutBucketRequestPaymentInput, optFns ...func(*Options)) (*PutBucketRequestPaymentOutput, error)
	PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error)
	PutBucketVersioning(ctx context.Context, params *PutBucketVersioningInput, optFns ...func(*Options)) (*PutBucketVersioningOutput, error)
	PutBucketWebsite(ctx context.Context, params *PutBucketWebsiteInput, optFns ...func(*Options)) (*PutBucketWebsiteOutput, error)
	PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error)
	PutObjectAcl(ctx context.Context, params *PutObjectAclInput, optFns ...func(*Options)) (*PutObjectAclOutput, error)
	PutObjectLegalHold(ctx context.Context, params *PutObjectLegalHoldInput, optFns ...func(*Options)) (*PutObjectLegalHoldOutput, error)
	PutObjectLockConfiguration(ctx context.Context, params *PutObjectLockConfigurationInput, optFns ...func(*Options)) (*PutObjectLockConfigurationOutput, error)
	PutObjectRetention(ctx context.Context, params *PutObjectRetentionInput, optFns ...func(*Options)) (*PutObjectRetentionOutput, error)
	PutObjectTagging(ctx context.Context, params *PutObjectTaggingInput, optFns ...func(*Options)) (*PutObjectTaggingOutput, error)
	PutPublicAccessBlock(ctx context.Context, params *PutPublicAccessBlockInput, optFns ...func(*Options)) (*PutPublicAccessBlockOutput, error)
	RestoreObject(ctx context.Context, params *RestoreObjectInput, optFns ...func(*Options)) (*RestoreObjectOutput, error)
	SelectObjectContent(ctx context.Context, params *SelectObjectContentInput, optFns ...func(*Options)) (*SelectObjectContentOutput, error)
	UploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*Options)) (*UploadPartOutput, error)
	UploadPartCopy(ctx context.Context, params *UploadPartCopyInput, optFns ...func(*Options)) (*UploadPartCopyOutput, error)
	WriteGetObjectResponse(ctx context.Context, params *WriteGetObjectResponseInput, optFns ...func(*Options)) (*WriteGetObjectResponseOutput, error)
}

type PresignClient interface {
	PresignGetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignHeadBucket(ctx context.Context, params *HeadBucketInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignHeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignUploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error)
}
