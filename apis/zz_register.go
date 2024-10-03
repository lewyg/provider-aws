// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/provider-aws/apis/accessanalyzer/v1beta1"
	v1beta2 "github.com/upbound/provider-aws/apis/accessanalyzer/v1beta2"
	v1beta1account "github.com/upbound/provider-aws/apis/account/v1beta1"
	v1beta1acm "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta2acm "github.com/upbound/provider-aws/apis/acm/v1beta2"
	v1beta1acmpca "github.com/upbound/provider-aws/apis/acmpca/v1beta1"
	v1beta2acmpca "github.com/upbound/provider-aws/apis/acmpca/v1beta2"
	v1beta1amp "github.com/upbound/provider-aws/apis/amp/v1beta1"
	v1beta2amp "github.com/upbound/provider-aws/apis/amp/v1beta2"
	v1beta1amplify "github.com/upbound/provider-aws/apis/amplify/v1beta1"
	v1beta2amplify "github.com/upbound/provider-aws/apis/amplify/v1beta2"
	v1beta1apigateway "github.com/upbound/provider-aws/apis/apigateway/v1beta1"
	v1beta2apigateway "github.com/upbound/provider-aws/apis/apigateway/v1beta2"
	v1beta1apigatewayv2 "github.com/upbound/provider-aws/apis/apigatewayv2/v1beta1"
	v1beta2apigatewayv2 "github.com/upbound/provider-aws/apis/apigatewayv2/v1beta2"
	v1beta1appautoscaling "github.com/upbound/provider-aws/apis/appautoscaling/v1beta1"
	v1beta2appautoscaling "github.com/upbound/provider-aws/apis/appautoscaling/v1beta2"
	v1beta1appconfig "github.com/upbound/provider-aws/apis/appconfig/v1beta1"
	v1beta1appflow "github.com/upbound/provider-aws/apis/appflow/v1beta1"
	v1beta2appflow "github.com/upbound/provider-aws/apis/appflow/v1beta2"
	v1beta1appintegrations "github.com/upbound/provider-aws/apis/appintegrations/v1beta1"
	v1beta2appintegrations "github.com/upbound/provider-aws/apis/appintegrations/v1beta2"
	v1beta1applicationinsights "github.com/upbound/provider-aws/apis/applicationinsights/v1beta1"
	v1beta1appmesh "github.com/upbound/provider-aws/apis/appmesh/v1beta1"
	v1beta2appmesh "github.com/upbound/provider-aws/apis/appmesh/v1beta2"
	v1beta1apprunner "github.com/upbound/provider-aws/apis/apprunner/v1beta1"
	v1beta2apprunner "github.com/upbound/provider-aws/apis/apprunner/v1beta2"
	v1beta1appstream "github.com/upbound/provider-aws/apis/appstream/v1beta1"
	v1beta2appstream "github.com/upbound/provider-aws/apis/appstream/v1beta2"
	v1beta1appsync "github.com/upbound/provider-aws/apis/appsync/v1beta1"
	v1beta2appsync "github.com/upbound/provider-aws/apis/appsync/v1beta2"
	v1beta1athena "github.com/upbound/provider-aws/apis/athena/v1beta1"
	v1beta2athena "github.com/upbound/provider-aws/apis/athena/v1beta2"
	v1beta1autoscaling "github.com/upbound/provider-aws/apis/autoscaling/v1beta1"
	v1beta2autoscaling "github.com/upbound/provider-aws/apis/autoscaling/v1beta2"
	v1beta3 "github.com/upbound/provider-aws/apis/autoscaling/v1beta3"
	v1beta1autoscalingplans "github.com/upbound/provider-aws/apis/autoscalingplans/v1beta1"
	v1beta2autoscalingplans "github.com/upbound/provider-aws/apis/autoscalingplans/v1beta2"
	v1beta1backup "github.com/upbound/provider-aws/apis/backup/v1beta1"
	v1beta2backup "github.com/upbound/provider-aws/apis/backup/v1beta2"
	v1beta1batch "github.com/upbound/provider-aws/apis/batch/v1beta1"
	v1beta2batch "github.com/upbound/provider-aws/apis/batch/v1beta2"
	v1beta1budgets "github.com/upbound/provider-aws/apis/budgets/v1beta1"
	v1beta2budgets "github.com/upbound/provider-aws/apis/budgets/v1beta2"
	v1beta1ce "github.com/upbound/provider-aws/apis/ce/v1beta1"
	v1beta1chime "github.com/upbound/provider-aws/apis/chime/v1beta1"
	v1beta2chime "github.com/upbound/provider-aws/apis/chime/v1beta2"
	v1beta1cloud9 "github.com/upbound/provider-aws/apis/cloud9/v1beta1"
	v1beta1cloudcontrol "github.com/upbound/provider-aws/apis/cloudcontrol/v1beta1"
	v1beta1cloudformation "github.com/upbound/provider-aws/apis/cloudformation/v1beta1"
	v1beta2cloudformation "github.com/upbound/provider-aws/apis/cloudformation/v1beta2"
	v1beta1cloudfront "github.com/upbound/provider-aws/apis/cloudfront/v1beta1"
	v1beta2cloudfront "github.com/upbound/provider-aws/apis/cloudfront/v1beta2"
	v1beta1cloudsearch "github.com/upbound/provider-aws/apis/cloudsearch/v1beta1"
	v1beta2cloudsearch "github.com/upbound/provider-aws/apis/cloudsearch/v1beta2"
	v1beta1cloudtrail "github.com/upbound/provider-aws/apis/cloudtrail/v1beta1"
	v1beta1cloudwatch "github.com/upbound/provider-aws/apis/cloudwatch/v1beta1"
	v1beta2cloudwatch "github.com/upbound/provider-aws/apis/cloudwatch/v1beta2"
	v1beta1cloudwatchevents "github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1"
	v1beta2cloudwatchevents "github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta2"
	v1beta1cloudwatchlogs "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta2cloudwatchlogs "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta2"
	v1beta1codeartifact "github.com/upbound/provider-aws/apis/codeartifact/v1beta1"
	v1beta1codecommit "github.com/upbound/provider-aws/apis/codecommit/v1beta1"
	v1beta1codeguruprofiler "github.com/upbound/provider-aws/apis/codeguruprofiler/v1beta1"
	v1beta1codepipeline "github.com/upbound/provider-aws/apis/codepipeline/v1beta1"
	v1beta2codepipeline "github.com/upbound/provider-aws/apis/codepipeline/v1beta2"
	v1beta1codestarconnections "github.com/upbound/provider-aws/apis/codestarconnections/v1beta1"
	v1beta2codestarconnections "github.com/upbound/provider-aws/apis/codestarconnections/v1beta2"
	v1beta1codestarnotifications "github.com/upbound/provider-aws/apis/codestarnotifications/v1beta1"
	v1beta1cognitoidentity "github.com/upbound/provider-aws/apis/cognitoidentity/v1beta1"
	v1beta1cognitoidp "github.com/upbound/provider-aws/apis/cognitoidp/v1beta1"
	v1beta2cognitoidp "github.com/upbound/provider-aws/apis/cognitoidp/v1beta2"
	v1beta1configservice "github.com/upbound/provider-aws/apis/configservice/v1beta1"
	v1beta2configservice "github.com/upbound/provider-aws/apis/configservice/v1beta2"
	v1beta1connect "github.com/upbound/provider-aws/apis/connect/v1beta1"
	v1beta2connect "github.com/upbound/provider-aws/apis/connect/v1beta2"
	v1beta3connect "github.com/upbound/provider-aws/apis/connect/v1beta3"
	v1beta1cur "github.com/upbound/provider-aws/apis/cur/v1beta1"
	v1beta1dataexchange "github.com/upbound/provider-aws/apis/dataexchange/v1beta1"
	v1beta1datapipeline "github.com/upbound/provider-aws/apis/datapipeline/v1beta1"
	v1beta1datasync "github.com/upbound/provider-aws/apis/datasync/v1beta1"
	v1beta2datasync "github.com/upbound/provider-aws/apis/datasync/v1beta2"
	v1beta1dax "github.com/upbound/provider-aws/apis/dax/v1beta1"
	v1beta2dax "github.com/upbound/provider-aws/apis/dax/v1beta2"
	v1beta1deploy "github.com/upbound/provider-aws/apis/deploy/v1beta1"
	v1beta2deploy "github.com/upbound/provider-aws/apis/deploy/v1beta2"
	v1beta1detective "github.com/upbound/provider-aws/apis/detective/v1beta1"
	v1beta1devicefarm "github.com/upbound/provider-aws/apis/devicefarm/v1beta1"
	v1beta2devicefarm "github.com/upbound/provider-aws/apis/devicefarm/v1beta2"
	v1beta1directconnect "github.com/upbound/provider-aws/apis/directconnect/v1beta1"
	v1beta1dlm "github.com/upbound/provider-aws/apis/dlm/v1beta1"
	v1beta2dlm "github.com/upbound/provider-aws/apis/dlm/v1beta2"
	v1beta1dms "github.com/upbound/provider-aws/apis/dms/v1beta1"
	v1beta2dms "github.com/upbound/provider-aws/apis/dms/v1beta2"
	v1beta1docdb "github.com/upbound/provider-aws/apis/docdb/v1beta1"
	v1beta1ds "github.com/upbound/provider-aws/apis/ds/v1beta1"
	v1beta2ds "github.com/upbound/provider-aws/apis/ds/v1beta2"
	v1beta1dynamodb "github.com/upbound/provider-aws/apis/dynamodb/v1beta1"
	v1beta2dynamodb "github.com/upbound/provider-aws/apis/dynamodb/v1beta2"
	v1beta1ec2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta2ec2 "github.com/upbound/provider-aws/apis/ec2/v1beta2"
	v1beta1ecr "github.com/upbound/provider-aws/apis/ecr/v1beta1"
	v1beta2ecr "github.com/upbound/provider-aws/apis/ecr/v1beta2"
	v1beta1ecrpublic "github.com/upbound/provider-aws/apis/ecrpublic/v1beta1"
	v1beta2ecrpublic "github.com/upbound/provider-aws/apis/ecrpublic/v1beta2"
	v1beta1ecs "github.com/upbound/provider-aws/apis/ecs/v1beta1"
	v1beta2ecs "github.com/upbound/provider-aws/apis/ecs/v1beta2"
	v1beta1efs "github.com/upbound/provider-aws/apis/efs/v1beta1"
	v1beta2efs "github.com/upbound/provider-aws/apis/efs/v1beta2"
	v1beta1eks "github.com/upbound/provider-aws/apis/eks/v1beta1"
	v1beta2eks "github.com/upbound/provider-aws/apis/eks/v1beta2"
	v1beta1elasticache "github.com/upbound/provider-aws/apis/elasticache/v1beta1"
	v1beta2elasticache "github.com/upbound/provider-aws/apis/elasticache/v1beta2"
	v1beta1elasticbeanstalk "github.com/upbound/provider-aws/apis/elasticbeanstalk/v1beta1"
	v1beta2elasticbeanstalk "github.com/upbound/provider-aws/apis/elasticbeanstalk/v1beta2"
	v1beta1elasticsearch "github.com/upbound/provider-aws/apis/elasticsearch/v1beta1"
	v1beta2elasticsearch "github.com/upbound/provider-aws/apis/elasticsearch/v1beta2"
	v1beta1elastictranscoder "github.com/upbound/provider-aws/apis/elastictranscoder/v1beta1"
	v1beta2elastictranscoder "github.com/upbound/provider-aws/apis/elastictranscoder/v1beta2"
	v1beta1elb "github.com/upbound/provider-aws/apis/elb/v1beta1"
	v1beta2elb "github.com/upbound/provider-aws/apis/elb/v1beta2"
	v1beta1elbv2 "github.com/upbound/provider-aws/apis/elbv2/v1beta1"
	v1beta2elbv2 "github.com/upbound/provider-aws/apis/elbv2/v1beta2"
	v1beta1emr "github.com/upbound/provider-aws/apis/emr/v1beta1"
	v1beta1emrserverless "github.com/upbound/provider-aws/apis/emrserverless/v1beta1"
	v1beta2emrserverless "github.com/upbound/provider-aws/apis/emrserverless/v1beta2"
	v1beta1evidently "github.com/upbound/provider-aws/apis/evidently/v1beta1"
	v1beta2evidently "github.com/upbound/provider-aws/apis/evidently/v1beta2"
	v1beta1firehose "github.com/upbound/provider-aws/apis/firehose/v1beta1"
	v1beta2firehose "github.com/upbound/provider-aws/apis/firehose/v1beta2"
	v1beta1fis "github.com/upbound/provider-aws/apis/fis/v1beta1"
	v1beta2fis "github.com/upbound/provider-aws/apis/fis/v1beta2"
	v1beta1fsx "github.com/upbound/provider-aws/apis/fsx/v1beta1"
	v1beta2fsx "github.com/upbound/provider-aws/apis/fsx/v1beta2"
	v1beta1gamelift "github.com/upbound/provider-aws/apis/gamelift/v1beta1"
	v1beta2gamelift "github.com/upbound/provider-aws/apis/gamelift/v1beta2"
	v1beta1glacier "github.com/upbound/provider-aws/apis/glacier/v1beta1"
	v1beta2glacier "github.com/upbound/provider-aws/apis/glacier/v1beta2"
	v1beta1globalaccelerator "github.com/upbound/provider-aws/apis/globalaccelerator/v1beta1"
	v1beta2globalaccelerator "github.com/upbound/provider-aws/apis/globalaccelerator/v1beta2"
	v1beta1glue "github.com/upbound/provider-aws/apis/glue/v1beta1"
	v1beta2glue "github.com/upbound/provider-aws/apis/glue/v1beta2"
	v1beta1grafana "github.com/upbound/provider-aws/apis/grafana/v1beta1"
	v1beta2grafana "github.com/upbound/provider-aws/apis/grafana/v1beta2"
	v1beta1guardduty "github.com/upbound/provider-aws/apis/guardduty/v1beta1"
	v1beta2guardduty "github.com/upbound/provider-aws/apis/guardduty/v1beta2"
	v1beta1iam "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1identitystore "github.com/upbound/provider-aws/apis/identitystore/v1beta1"
	v1beta2identitystore "github.com/upbound/provider-aws/apis/identitystore/v1beta2"
	v1beta1imagebuilder "github.com/upbound/provider-aws/apis/imagebuilder/v1beta1"
	v1beta2imagebuilder "github.com/upbound/provider-aws/apis/imagebuilder/v1beta2"
	v1beta1inspector "github.com/upbound/provider-aws/apis/inspector/v1beta1"
	v1beta1inspector2 "github.com/upbound/provider-aws/apis/inspector2/v1beta1"
	v1beta1iot "github.com/upbound/provider-aws/apis/iot/v1beta1"
	v1beta2iot "github.com/upbound/provider-aws/apis/iot/v1beta2"
	v1beta1ivs "github.com/upbound/provider-aws/apis/ivs/v1beta1"
	v1beta2ivs "github.com/upbound/provider-aws/apis/ivs/v1beta2"
	v1beta1kafka "github.com/upbound/provider-aws/apis/kafka/v1beta1"
	v1beta2kafka "github.com/upbound/provider-aws/apis/kafka/v1beta2"
	v1beta3kafka "github.com/upbound/provider-aws/apis/kafka/v1beta3"
	v1beta1kafkaconnect "github.com/upbound/provider-aws/apis/kafkaconnect/v1beta1"
	v1beta2kafkaconnect "github.com/upbound/provider-aws/apis/kafkaconnect/v1beta2"
	v1beta1kendra "github.com/upbound/provider-aws/apis/kendra/v1beta1"
	v1beta2kendra "github.com/upbound/provider-aws/apis/kendra/v1beta2"
	v1beta1keyspaces "github.com/upbound/provider-aws/apis/keyspaces/v1beta1"
	v1beta2keyspaces "github.com/upbound/provider-aws/apis/keyspaces/v1beta2"
	v1beta1kinesis "github.com/upbound/provider-aws/apis/kinesis/v1beta1"
	v1beta2kinesis "github.com/upbound/provider-aws/apis/kinesis/v1beta2"
	v1beta1kinesisanalytics "github.com/upbound/provider-aws/apis/kinesisanalytics/v1beta1"
	v1beta2kinesisanalytics "github.com/upbound/provider-aws/apis/kinesisanalytics/v1beta2"
	v1beta1kinesisanalyticsv2 "github.com/upbound/provider-aws/apis/kinesisanalyticsv2/v1beta1"
	v1beta2kinesisanalyticsv2 "github.com/upbound/provider-aws/apis/kinesisanalyticsv2/v1beta2"
	v1beta1kinesisvideo "github.com/upbound/provider-aws/apis/kinesisvideo/v1beta1"
	v1beta1kms "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta1lakeformation "github.com/upbound/provider-aws/apis/lakeformation/v1beta1"
	v1beta2lakeformation "github.com/upbound/provider-aws/apis/lakeformation/v1beta2"
	v1beta1lambda "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	v1beta2lambda "github.com/upbound/provider-aws/apis/lambda/v1beta2"
	v1beta1lexmodels "github.com/upbound/provider-aws/apis/lexmodels/v1beta1"
	v1beta2lexmodels "github.com/upbound/provider-aws/apis/lexmodels/v1beta2"
	v1beta1licensemanager "github.com/upbound/provider-aws/apis/licensemanager/v1beta1"
	v1beta1lightsail "github.com/upbound/provider-aws/apis/lightsail/v1beta1"
	v1beta2lightsail "github.com/upbound/provider-aws/apis/lightsail/v1beta2"
	v1beta1location "github.com/upbound/provider-aws/apis/location/v1beta1"
	v1beta2location "github.com/upbound/provider-aws/apis/location/v1beta2"
	v1beta1macie2 "github.com/upbound/provider-aws/apis/macie2/v1beta1"
	v1beta2macie2 "github.com/upbound/provider-aws/apis/macie2/v1beta2"
	v1beta1mediaconvert "github.com/upbound/provider-aws/apis/mediaconvert/v1beta1"
	v1beta2mediaconvert "github.com/upbound/provider-aws/apis/mediaconvert/v1beta2"
	v1beta1medialive "github.com/upbound/provider-aws/apis/medialive/v1beta1"
	v1beta2medialive "github.com/upbound/provider-aws/apis/medialive/v1beta2"
	v1beta1mediapackage "github.com/upbound/provider-aws/apis/mediapackage/v1beta1"
	v1beta1mediastore "github.com/upbound/provider-aws/apis/mediastore/v1beta1"
	v1beta1memorydb "github.com/upbound/provider-aws/apis/memorydb/v1beta1"
	v1beta2memorydb "github.com/upbound/provider-aws/apis/memorydb/v1beta2"
	v1alpha1 "github.com/upbound/provider-aws/apis/mq/v1alpha1"
	v1beta1mq "github.com/upbound/provider-aws/apis/mq/v1beta1"
	v1beta2mq "github.com/upbound/provider-aws/apis/mq/v1beta2"
	v1beta1mwaa "github.com/upbound/provider-aws/apis/mwaa/v1beta1"
	v1beta1neptune "github.com/upbound/provider-aws/apis/neptune/v1beta1"
	v1beta2neptune "github.com/upbound/provider-aws/apis/neptune/v1beta2"
	v1beta1networkfirewall "github.com/upbound/provider-aws/apis/networkfirewall/v1beta1"
	v1beta2networkfirewall "github.com/upbound/provider-aws/apis/networkfirewall/v1beta2"
	v1beta1networkmanager "github.com/upbound/provider-aws/apis/networkmanager/v1beta1"
	v1beta2networkmanager "github.com/upbound/provider-aws/apis/networkmanager/v1beta2"
	v1beta1opensearch "github.com/upbound/provider-aws/apis/opensearch/v1beta1"
	v1beta2opensearch "github.com/upbound/provider-aws/apis/opensearch/v1beta2"
	v1beta1opensearchserverless "github.com/upbound/provider-aws/apis/opensearchserverless/v1beta1"
	v1beta1opsworks "github.com/upbound/provider-aws/apis/opsworks/v1beta1"
	v1beta2opsworks "github.com/upbound/provider-aws/apis/opsworks/v1beta2"
	v1beta1organizations "github.com/upbound/provider-aws/apis/organizations/v1beta1"
	v1beta1osis "github.com/upbound/provider-aws/apis/osis/v1beta1"
	v1beta1pinpoint "github.com/upbound/provider-aws/apis/pinpoint/v1beta1"
	v1beta2pinpoint "github.com/upbound/provider-aws/apis/pinpoint/v1beta2"
	v1beta1pipes "github.com/upbound/provider-aws/apis/pipes/v1beta1"
	v1beta1qldb "github.com/upbound/provider-aws/apis/qldb/v1beta1"
	v1beta2qldb "github.com/upbound/provider-aws/apis/qldb/v1beta2"
	v1beta1quicksight "github.com/upbound/provider-aws/apis/quicksight/v1beta1"
	v1beta1ram "github.com/upbound/provider-aws/apis/ram/v1beta1"
	v1beta1rds "github.com/upbound/provider-aws/apis/rds/v1beta1"
	v1beta2rds "github.com/upbound/provider-aws/apis/rds/v1beta2"
	v1beta3rds "github.com/upbound/provider-aws/apis/rds/v1beta3"
	v1beta1redshift "github.com/upbound/provider-aws/apis/redshift/v1beta1"
	v1beta2redshift "github.com/upbound/provider-aws/apis/redshift/v1beta2"
	v1beta1redshiftserverless "github.com/upbound/provider-aws/apis/redshiftserverless/v1beta1"
	v1beta1resourcegroups "github.com/upbound/provider-aws/apis/resourcegroups/v1beta1"
	v1beta2resourcegroups "github.com/upbound/provider-aws/apis/resourcegroups/v1beta2"
	v1beta1rolesanywhere "github.com/upbound/provider-aws/apis/rolesanywhere/v1beta1"
	v1beta1route53 "github.com/upbound/provider-aws/apis/route53/v1beta1"
	v1beta2route53 "github.com/upbound/provider-aws/apis/route53/v1beta2"
	v1beta1route53recoverycontrolconfig "github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta1"
	v1beta2route53recoverycontrolconfig "github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta2"
	v1beta1route53recoveryreadiness "github.com/upbound/provider-aws/apis/route53recoveryreadiness/v1beta1"
	v1beta2route53recoveryreadiness "github.com/upbound/provider-aws/apis/route53recoveryreadiness/v1beta2"
	v1beta1route53resolver "github.com/upbound/provider-aws/apis/route53resolver/v1beta1"
	v1beta1rum "github.com/upbound/provider-aws/apis/rum/v1beta1"
	v1beta2rum "github.com/upbound/provider-aws/apis/rum/v1beta2"
	v1beta1s3 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta2s3 "github.com/upbound/provider-aws/apis/s3/v1beta2"
	v1beta1s3control "github.com/upbound/provider-aws/apis/s3control/v1beta1"
	v1beta2s3control "github.com/upbound/provider-aws/apis/s3control/v1beta2"
	v1beta1sagemaker "github.com/upbound/provider-aws/apis/sagemaker/v1beta1"
	v1beta2sagemaker "github.com/upbound/provider-aws/apis/sagemaker/v1beta2"
	v1beta1scheduler "github.com/upbound/provider-aws/apis/scheduler/v1beta1"
	v1beta2scheduler "github.com/upbound/provider-aws/apis/scheduler/v1beta2"
	v1beta1schemas "github.com/upbound/provider-aws/apis/schemas/v1beta1"
	v1beta1secretsmanager "github.com/upbound/provider-aws/apis/secretsmanager/v1beta1"
	v1beta2secretsmanager "github.com/upbound/provider-aws/apis/secretsmanager/v1beta2"
	v1beta1securityhub "github.com/upbound/provider-aws/apis/securityhub/v1beta1"
	v1beta2securityhub "github.com/upbound/provider-aws/apis/securityhub/v1beta2"
	v1beta1serverlessrepo "github.com/upbound/provider-aws/apis/serverlessrepo/v1beta1"
	v1beta1servicecatalog "github.com/upbound/provider-aws/apis/servicecatalog/v1beta1"
	v1beta2servicecatalog "github.com/upbound/provider-aws/apis/servicecatalog/v1beta2"
	v1beta1servicediscovery "github.com/upbound/provider-aws/apis/servicediscovery/v1beta1"
	v1beta2servicediscovery "github.com/upbound/provider-aws/apis/servicediscovery/v1beta2"
	v1beta1servicequotas "github.com/upbound/provider-aws/apis/servicequotas/v1beta1"
	v1beta1ses "github.com/upbound/provider-aws/apis/ses/v1beta1"
	v1beta2ses "github.com/upbound/provider-aws/apis/ses/v1beta2"
	v1beta1sesv2 "github.com/upbound/provider-aws/apis/sesv2/v1beta1"
	v1beta2sesv2 "github.com/upbound/provider-aws/apis/sesv2/v1beta2"
	v1beta1sfn "github.com/upbound/provider-aws/apis/sfn/v1beta1"
	v1beta2sfn "github.com/upbound/provider-aws/apis/sfn/v1beta2"
	v1beta1signer "github.com/upbound/provider-aws/apis/signer/v1beta1"
	v1beta2signer "github.com/upbound/provider-aws/apis/signer/v1beta2"
	v1beta1simpledb "github.com/upbound/provider-aws/apis/simpledb/v1beta1"
	v1beta1sns "github.com/upbound/provider-aws/apis/sns/v1beta1"
	v1beta1sqs "github.com/upbound/provider-aws/apis/sqs/v1beta1"
	v1beta1ssm "github.com/upbound/provider-aws/apis/ssm/v1beta1"
	v1beta2ssm "github.com/upbound/provider-aws/apis/ssm/v1beta2"
	v1beta1ssoadmin "github.com/upbound/provider-aws/apis/ssoadmin/v1beta1"
	v1beta2ssoadmin "github.com/upbound/provider-aws/apis/ssoadmin/v1beta2"
	v1beta1swf "github.com/upbound/provider-aws/apis/swf/v1beta1"
	v1beta1timestreamwrite "github.com/upbound/provider-aws/apis/timestreamwrite/v1beta1"
	v1beta2timestreamwrite "github.com/upbound/provider-aws/apis/timestreamwrite/v1beta2"
	v1beta1transcribe "github.com/upbound/provider-aws/apis/transcribe/v1beta1"
	v1beta2transcribe "github.com/upbound/provider-aws/apis/transcribe/v1beta2"
	v1beta1transfer "github.com/upbound/provider-aws/apis/transfer/v1beta1"
	v1beta2transfer "github.com/upbound/provider-aws/apis/transfer/v1beta2"
	v1alpha1apis "github.com/upbound/provider-aws/apis/v1alpha1"
	v1beta1apis "github.com/upbound/provider-aws/apis/v1beta1"
	v1beta1vpc "github.com/upbound/provider-aws/apis/vpc/v1beta1"
	v1beta1waf "github.com/upbound/provider-aws/apis/waf/v1beta1"
	v1beta2waf "github.com/upbound/provider-aws/apis/waf/v1beta2"
	v1beta1wafregional "github.com/upbound/provider-aws/apis/wafregional/v1beta1"
	v1beta2wafregional "github.com/upbound/provider-aws/apis/wafregional/v1beta2"
	v1beta1wafv2 "github.com/upbound/provider-aws/apis/wafv2/v1beta1"
	v1beta1workspaces "github.com/upbound/provider-aws/apis/workspaces/v1beta1"
	v1beta2workspaces "github.com/upbound/provider-aws/apis/workspaces/v1beta2"
	v1beta1xray "github.com/upbound/provider-aws/apis/xray/v1beta1"
	v1beta2xray "github.com/upbound/provider-aws/apis/xray/v1beta2"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta2.SchemeBuilder.AddToScheme,
		v1beta1account.SchemeBuilder.AddToScheme,
		v1beta1acm.SchemeBuilder.AddToScheme,
		v1beta2acm.SchemeBuilder.AddToScheme,
		v1beta1acmpca.SchemeBuilder.AddToScheme,
		v1beta2acmpca.SchemeBuilder.AddToScheme,
		v1beta1amp.SchemeBuilder.AddToScheme,
		v1beta2amp.SchemeBuilder.AddToScheme,
		v1beta1amplify.SchemeBuilder.AddToScheme,
		v1beta2amplify.SchemeBuilder.AddToScheme,
		v1beta1apigateway.SchemeBuilder.AddToScheme,
		v1beta2apigateway.SchemeBuilder.AddToScheme,
		v1beta1apigatewayv2.SchemeBuilder.AddToScheme,
		v1beta2apigatewayv2.SchemeBuilder.AddToScheme,
		v1beta1appautoscaling.SchemeBuilder.AddToScheme,
		v1beta2appautoscaling.SchemeBuilder.AddToScheme,
		v1beta1appconfig.SchemeBuilder.AddToScheme,
		v1beta1appflow.SchemeBuilder.AddToScheme,
		v1beta2appflow.SchemeBuilder.AddToScheme,
		v1beta1appintegrations.SchemeBuilder.AddToScheme,
		v1beta2appintegrations.SchemeBuilder.AddToScheme,
		v1beta1applicationinsights.SchemeBuilder.AddToScheme,
		v1beta1appmesh.SchemeBuilder.AddToScheme,
		v1beta2appmesh.SchemeBuilder.AddToScheme,
		v1beta1apprunner.SchemeBuilder.AddToScheme,
		v1beta2apprunner.SchemeBuilder.AddToScheme,
		v1beta1appstream.SchemeBuilder.AddToScheme,
		v1beta2appstream.SchemeBuilder.AddToScheme,
		v1beta1appsync.SchemeBuilder.AddToScheme,
		v1beta2appsync.SchemeBuilder.AddToScheme,
		v1beta1athena.SchemeBuilder.AddToScheme,
		v1beta2athena.SchemeBuilder.AddToScheme,
		v1beta1autoscaling.SchemeBuilder.AddToScheme,
		v1beta2autoscaling.SchemeBuilder.AddToScheme,
		v1beta3.SchemeBuilder.AddToScheme,
		v1beta1autoscalingplans.SchemeBuilder.AddToScheme,
		v1beta2autoscalingplans.SchemeBuilder.AddToScheme,
		v1beta1backup.SchemeBuilder.AddToScheme,
		v1beta2backup.SchemeBuilder.AddToScheme,
		v1beta1batch.SchemeBuilder.AddToScheme,
		v1beta2batch.SchemeBuilder.AddToScheme,
		v1beta1budgets.SchemeBuilder.AddToScheme,
		v1beta2budgets.SchemeBuilder.AddToScheme,
		v1beta1ce.SchemeBuilder.AddToScheme,
		v1beta1chime.SchemeBuilder.AddToScheme,
		v1beta2chime.SchemeBuilder.AddToScheme,
		v1beta1cloud9.SchemeBuilder.AddToScheme,
		v1beta1cloudcontrol.SchemeBuilder.AddToScheme,
		v1beta1cloudformation.SchemeBuilder.AddToScheme,
		v1beta2cloudformation.SchemeBuilder.AddToScheme,
		v1beta1cloudfront.SchemeBuilder.AddToScheme,
		v1beta2cloudfront.SchemeBuilder.AddToScheme,
		v1beta1cloudsearch.SchemeBuilder.AddToScheme,
		v1beta2cloudsearch.SchemeBuilder.AddToScheme,
		v1beta1cloudtrail.SchemeBuilder.AddToScheme,
		v1beta1cloudwatch.SchemeBuilder.AddToScheme,
		v1beta2cloudwatch.SchemeBuilder.AddToScheme,
		v1beta1cloudwatchevents.SchemeBuilder.AddToScheme,
		v1beta2cloudwatchevents.SchemeBuilder.AddToScheme,
		v1beta1cloudwatchlogs.SchemeBuilder.AddToScheme,
		v1beta2cloudwatchlogs.SchemeBuilder.AddToScheme,
		v1beta1codeartifact.SchemeBuilder.AddToScheme,
		v1beta1codecommit.SchemeBuilder.AddToScheme,
		v1beta1codeguruprofiler.SchemeBuilder.AddToScheme,
		v1beta1codepipeline.SchemeBuilder.AddToScheme,
		v1beta2codepipeline.SchemeBuilder.AddToScheme,
		v1beta1codestarconnections.SchemeBuilder.AddToScheme,
		v1beta2codestarconnections.SchemeBuilder.AddToScheme,
		v1beta1codestarnotifications.SchemeBuilder.AddToScheme,
		v1beta1cognitoidentity.SchemeBuilder.AddToScheme,
		v1beta1cognitoidp.SchemeBuilder.AddToScheme,
		v1beta2cognitoidp.SchemeBuilder.AddToScheme,
		v1beta1configservice.SchemeBuilder.AddToScheme,
		v1beta2configservice.SchemeBuilder.AddToScheme,
		v1beta1connect.SchemeBuilder.AddToScheme,
		v1beta2connect.SchemeBuilder.AddToScheme,
		v1beta3connect.SchemeBuilder.AddToScheme,
		v1beta1cur.SchemeBuilder.AddToScheme,
		v1beta1dataexchange.SchemeBuilder.AddToScheme,
		v1beta1datapipeline.SchemeBuilder.AddToScheme,
		v1beta1datasync.SchemeBuilder.AddToScheme,
		v1beta2datasync.SchemeBuilder.AddToScheme,
		v1beta1dax.SchemeBuilder.AddToScheme,
		v1beta2dax.SchemeBuilder.AddToScheme,
		v1beta1deploy.SchemeBuilder.AddToScheme,
		v1beta2deploy.SchemeBuilder.AddToScheme,
		v1beta1detective.SchemeBuilder.AddToScheme,
		v1beta1devicefarm.SchemeBuilder.AddToScheme,
		v1beta2devicefarm.SchemeBuilder.AddToScheme,
		v1beta1directconnect.SchemeBuilder.AddToScheme,
		v1beta1dlm.SchemeBuilder.AddToScheme,
		v1beta2dlm.SchemeBuilder.AddToScheme,
		v1beta1dms.SchemeBuilder.AddToScheme,
		v1beta2dms.SchemeBuilder.AddToScheme,
		v1beta1docdb.SchemeBuilder.AddToScheme,
		v1beta1ds.SchemeBuilder.AddToScheme,
		v1beta2ds.SchemeBuilder.AddToScheme,
		v1beta1dynamodb.SchemeBuilder.AddToScheme,
		v1beta2dynamodb.SchemeBuilder.AddToScheme,
		v1beta1ec2.SchemeBuilder.AddToScheme,
		v1beta2ec2.SchemeBuilder.AddToScheme,
		v1beta1ecr.SchemeBuilder.AddToScheme,
		v1beta2ecr.SchemeBuilder.AddToScheme,
		v1beta1ecrpublic.SchemeBuilder.AddToScheme,
		v1beta2ecrpublic.SchemeBuilder.AddToScheme,
		v1beta1ecs.SchemeBuilder.AddToScheme,
		v1beta2ecs.SchemeBuilder.AddToScheme,
		v1beta1efs.SchemeBuilder.AddToScheme,
		v1beta2efs.SchemeBuilder.AddToScheme,
		v1beta1eks.SchemeBuilder.AddToScheme,
		v1beta2eks.SchemeBuilder.AddToScheme,
		v1beta1elasticache.SchemeBuilder.AddToScheme,
		v1beta2elasticache.SchemeBuilder.AddToScheme,
		v1beta1elasticbeanstalk.SchemeBuilder.AddToScheme,
		v1beta2elasticbeanstalk.SchemeBuilder.AddToScheme,
		v1beta1elasticsearch.SchemeBuilder.AddToScheme,
		v1beta2elasticsearch.SchemeBuilder.AddToScheme,
		v1beta1elastictranscoder.SchemeBuilder.AddToScheme,
		v1beta2elastictranscoder.SchemeBuilder.AddToScheme,
		v1beta1elb.SchemeBuilder.AddToScheme,
		v1beta2elb.SchemeBuilder.AddToScheme,
		v1beta1elbv2.SchemeBuilder.AddToScheme,
		v1beta2elbv2.SchemeBuilder.AddToScheme,
		v1beta1emr.SchemeBuilder.AddToScheme,
		v1beta1emrserverless.SchemeBuilder.AddToScheme,
		v1beta2emrserverless.SchemeBuilder.AddToScheme,
		v1beta1evidently.SchemeBuilder.AddToScheme,
		v1beta2evidently.SchemeBuilder.AddToScheme,
		v1beta1firehose.SchemeBuilder.AddToScheme,
		v1beta2firehose.SchemeBuilder.AddToScheme,
		v1beta1fis.SchemeBuilder.AddToScheme,
		v1beta2fis.SchemeBuilder.AddToScheme,
		v1beta1fsx.SchemeBuilder.AddToScheme,
		v1beta2fsx.SchemeBuilder.AddToScheme,
		v1beta1gamelift.SchemeBuilder.AddToScheme,
		v1beta2gamelift.SchemeBuilder.AddToScheme,
		v1beta1glacier.SchemeBuilder.AddToScheme,
		v1beta2glacier.SchemeBuilder.AddToScheme,
		v1beta1globalaccelerator.SchemeBuilder.AddToScheme,
		v1beta2globalaccelerator.SchemeBuilder.AddToScheme,
		v1beta1glue.SchemeBuilder.AddToScheme,
		v1beta2glue.SchemeBuilder.AddToScheme,
		v1beta1grafana.SchemeBuilder.AddToScheme,
		v1beta2grafana.SchemeBuilder.AddToScheme,
		v1beta1guardduty.SchemeBuilder.AddToScheme,
		v1beta2guardduty.SchemeBuilder.AddToScheme,
		v1beta1iam.SchemeBuilder.AddToScheme,
		v1beta1identitystore.SchemeBuilder.AddToScheme,
		v1beta2identitystore.SchemeBuilder.AddToScheme,
		v1beta1imagebuilder.SchemeBuilder.AddToScheme,
		v1beta2imagebuilder.SchemeBuilder.AddToScheme,
		v1beta1inspector.SchemeBuilder.AddToScheme,
		v1beta1inspector2.SchemeBuilder.AddToScheme,
		v1beta1iot.SchemeBuilder.AddToScheme,
		v1beta2iot.SchemeBuilder.AddToScheme,
		v1beta1ivs.SchemeBuilder.AddToScheme,
		v1beta2ivs.SchemeBuilder.AddToScheme,
		v1beta1kafka.SchemeBuilder.AddToScheme,
		v1beta2kafka.SchemeBuilder.AddToScheme,
		v1beta3kafka.SchemeBuilder.AddToScheme,
		v1beta1kafkaconnect.SchemeBuilder.AddToScheme,
		v1beta2kafkaconnect.SchemeBuilder.AddToScheme,
		v1beta1kendra.SchemeBuilder.AddToScheme,
		v1beta2kendra.SchemeBuilder.AddToScheme,
		v1beta1keyspaces.SchemeBuilder.AddToScheme,
		v1beta2keyspaces.SchemeBuilder.AddToScheme,
		v1beta1kinesis.SchemeBuilder.AddToScheme,
		v1beta2kinesis.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalytics.SchemeBuilder.AddToScheme,
		v1beta2kinesisanalytics.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalyticsv2.SchemeBuilder.AddToScheme,
		v1beta2kinesisanalyticsv2.SchemeBuilder.AddToScheme,
		v1beta1kinesisvideo.SchemeBuilder.AddToScheme,
		v1beta1kms.SchemeBuilder.AddToScheme,
		v1beta1lakeformation.SchemeBuilder.AddToScheme,
		v1beta2lakeformation.SchemeBuilder.AddToScheme,
		v1beta1lambda.SchemeBuilder.AddToScheme,
		v1beta2lambda.SchemeBuilder.AddToScheme,
		v1beta1lexmodels.SchemeBuilder.AddToScheme,
		v1beta2lexmodels.SchemeBuilder.AddToScheme,
		v1beta1licensemanager.SchemeBuilder.AddToScheme,
		v1beta1lightsail.SchemeBuilder.AddToScheme,
		v1beta2lightsail.SchemeBuilder.AddToScheme,
		v1beta1location.SchemeBuilder.AddToScheme,
		v1beta2location.SchemeBuilder.AddToScheme,
		v1beta1macie2.SchemeBuilder.AddToScheme,
		v1beta2macie2.SchemeBuilder.AddToScheme,
		v1beta1mediaconvert.SchemeBuilder.AddToScheme,
		v1beta2mediaconvert.SchemeBuilder.AddToScheme,
		v1beta1medialive.SchemeBuilder.AddToScheme,
		v1beta2medialive.SchemeBuilder.AddToScheme,
		v1beta1mediapackage.SchemeBuilder.AddToScheme,
		v1beta1mediastore.SchemeBuilder.AddToScheme,
		v1beta1memorydb.SchemeBuilder.AddToScheme,
		v1beta2memorydb.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1mq.SchemeBuilder.AddToScheme,
		v1beta2mq.SchemeBuilder.AddToScheme,
		v1beta1mwaa.SchemeBuilder.AddToScheme,
		v1beta1neptune.SchemeBuilder.AddToScheme,
		v1beta2neptune.SchemeBuilder.AddToScheme,
		v1beta1networkfirewall.SchemeBuilder.AddToScheme,
		v1beta2networkfirewall.SchemeBuilder.AddToScheme,
		v1beta1networkmanager.SchemeBuilder.AddToScheme,
		v1beta2networkmanager.SchemeBuilder.AddToScheme,
		v1beta1opensearch.SchemeBuilder.AddToScheme,
		v1beta2opensearch.SchemeBuilder.AddToScheme,
		v1beta1opensearchserverless.SchemeBuilder.AddToScheme,
		v1beta1opsworks.SchemeBuilder.AddToScheme,
		v1beta2opsworks.SchemeBuilder.AddToScheme,
		v1beta1organizations.SchemeBuilder.AddToScheme,
		v1beta1osis.SchemeBuilder.AddToScheme,
		v1beta1pinpoint.SchemeBuilder.AddToScheme,
		v1beta2pinpoint.SchemeBuilder.AddToScheme,
		v1beta1pipes.SchemeBuilder.AddToScheme,
		v1beta1qldb.SchemeBuilder.AddToScheme,
		v1beta2qldb.SchemeBuilder.AddToScheme,
		v1beta1quicksight.SchemeBuilder.AddToScheme,
		v1beta1ram.SchemeBuilder.AddToScheme,
		v1beta1rds.SchemeBuilder.AddToScheme,
		v1beta2rds.SchemeBuilder.AddToScheme,
		v1beta3rds.SchemeBuilder.AddToScheme,
		v1beta1redshift.SchemeBuilder.AddToScheme,
		v1beta2redshift.SchemeBuilder.AddToScheme,
		v1beta1redshiftserverless.SchemeBuilder.AddToScheme,
		v1beta1resourcegroups.SchemeBuilder.AddToScheme,
		v1beta2resourcegroups.SchemeBuilder.AddToScheme,
		v1beta1rolesanywhere.SchemeBuilder.AddToScheme,
		v1beta1route53.SchemeBuilder.AddToScheme,
		v1beta2route53.SchemeBuilder.AddToScheme,
		v1beta1route53recoverycontrolconfig.SchemeBuilder.AddToScheme,
		v1beta2route53recoverycontrolconfig.SchemeBuilder.AddToScheme,
		v1beta1route53recoveryreadiness.SchemeBuilder.AddToScheme,
		v1beta2route53recoveryreadiness.SchemeBuilder.AddToScheme,
		v1beta1route53resolver.SchemeBuilder.AddToScheme,
		v1beta1rum.SchemeBuilder.AddToScheme,
		v1beta2rum.SchemeBuilder.AddToScheme,
		v1beta1s3.SchemeBuilder.AddToScheme,
		v1beta2s3.SchemeBuilder.AddToScheme,
		v1beta1s3control.SchemeBuilder.AddToScheme,
		v1beta2s3control.SchemeBuilder.AddToScheme,
		v1beta1sagemaker.SchemeBuilder.AddToScheme,
		v1beta2sagemaker.SchemeBuilder.AddToScheme,
		v1beta1scheduler.SchemeBuilder.AddToScheme,
		v1beta2scheduler.SchemeBuilder.AddToScheme,
		v1beta1schemas.SchemeBuilder.AddToScheme,
		v1beta1secretsmanager.SchemeBuilder.AddToScheme,
		v1beta2secretsmanager.SchemeBuilder.AddToScheme,
		v1beta1securityhub.SchemeBuilder.AddToScheme,
		v1beta2securityhub.SchemeBuilder.AddToScheme,
		v1beta1serverlessrepo.SchemeBuilder.AddToScheme,
		v1beta1servicecatalog.SchemeBuilder.AddToScheme,
		v1beta2servicecatalog.SchemeBuilder.AddToScheme,
		v1beta1servicediscovery.SchemeBuilder.AddToScheme,
		v1beta2servicediscovery.SchemeBuilder.AddToScheme,
		v1beta1servicequotas.SchemeBuilder.AddToScheme,
		v1beta1ses.SchemeBuilder.AddToScheme,
		v1beta2ses.SchemeBuilder.AddToScheme,
		v1beta1sesv2.SchemeBuilder.AddToScheme,
		v1beta2sesv2.SchemeBuilder.AddToScheme,
		v1beta1sfn.SchemeBuilder.AddToScheme,
		v1beta2sfn.SchemeBuilder.AddToScheme,
		v1beta1signer.SchemeBuilder.AddToScheme,
		v1beta2signer.SchemeBuilder.AddToScheme,
		v1beta1simpledb.SchemeBuilder.AddToScheme,
		v1beta1sns.SchemeBuilder.AddToScheme,
		v1beta1sqs.SchemeBuilder.AddToScheme,
		v1beta1ssm.SchemeBuilder.AddToScheme,
		v1beta2ssm.SchemeBuilder.AddToScheme,
		v1beta1ssoadmin.SchemeBuilder.AddToScheme,
		v1beta2ssoadmin.SchemeBuilder.AddToScheme,
		v1beta1swf.SchemeBuilder.AddToScheme,
		v1beta1timestreamwrite.SchemeBuilder.AddToScheme,
		v1beta2timestreamwrite.SchemeBuilder.AddToScheme,
		v1beta1transcribe.SchemeBuilder.AddToScheme,
		v1beta2transcribe.SchemeBuilder.AddToScheme,
		v1beta1transfer.SchemeBuilder.AddToScheme,
		v1beta2transfer.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
		v1beta1vpc.SchemeBuilder.AddToScheme,
		v1beta1waf.SchemeBuilder.AddToScheme,
		v1beta2waf.SchemeBuilder.AddToScheme,
		v1beta1wafregional.SchemeBuilder.AddToScheme,
		v1beta2wafregional.SchemeBuilder.AddToScheme,
		v1beta1wafv2.SchemeBuilder.AddToScheme,
		v1beta1workspaces.SchemeBuilder.AddToScheme,
		v1beta2workspaces.SchemeBuilder.AddToScheme,
		v1beta1xray.SchemeBuilder.AddToScheme,
		v1beta2xray.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
