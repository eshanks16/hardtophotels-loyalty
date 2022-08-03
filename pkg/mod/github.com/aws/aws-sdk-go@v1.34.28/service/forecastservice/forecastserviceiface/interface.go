// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package forecastserviceiface provides an interface to enable mocking the Amazon Forecast Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package forecastserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/forecastservice"
)

// ForecastServiceAPI provides an interface to enable mocking the
// forecastservice.ForecastService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Forecast Service.
//    func myFunc(svc forecastserviceiface.ForecastServiceAPI) bool {
//        // Make svc.CreateDataset request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := forecastservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockForecastServiceClient struct {
//        forecastserviceiface.ForecastServiceAPI
//    }
//    func (m *mockForecastServiceClient) CreateDataset(input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockForecastServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ForecastServiceAPI interface {
	CreateDataset(*forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *forecastservice.CreateDatasetInput, ...request.Option) (*forecastservice.CreateDatasetOutput, error)
	CreateDatasetRequest(*forecastservice.CreateDatasetInput) (*request.Request, *forecastservice.CreateDatasetOutput)

	CreateDatasetGroup(*forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error)
	CreateDatasetGroupWithContext(aws.Context, *forecastservice.CreateDatasetGroupInput, ...request.Option) (*forecastservice.CreateDatasetGroupOutput, error)
	CreateDatasetGroupRequest(*forecastservice.CreateDatasetGroupInput) (*request.Request, *forecastservice.CreateDatasetGroupOutput)

	CreateDatasetImportJob(*forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobWithContext(aws.Context, *forecastservice.CreateDatasetImportJobInput, ...request.Option) (*forecastservice.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobRequest(*forecastservice.CreateDatasetImportJobInput) (*request.Request, *forecastservice.CreateDatasetImportJobOutput)

	CreateForecast(*forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error)
	CreateForecastWithContext(aws.Context, *forecastservice.CreateForecastInput, ...request.Option) (*forecastservice.CreateForecastOutput, error)
	CreateForecastRequest(*forecastservice.CreateForecastInput) (*request.Request, *forecastservice.CreateForecastOutput)

	CreateForecastExportJob(*forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error)
	CreateForecastExportJobWithContext(aws.Context, *forecastservice.CreateForecastExportJobInput, ...request.Option) (*forecastservice.CreateForecastExportJobOutput, error)
	CreateForecastExportJobRequest(*forecastservice.CreateForecastExportJobInput) (*request.Request, *forecastservice.CreateForecastExportJobOutput)

	CreatePredictor(*forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error)
	CreatePredictorWithContext(aws.Context, *forecastservice.CreatePredictorInput, ...request.Option) (*forecastservice.CreatePredictorOutput, error)
	CreatePredictorRequest(*forecastservice.CreatePredictorInput) (*request.Request, *forecastservice.CreatePredictorOutput)

	DeleteDataset(*forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *forecastservice.DeleteDatasetInput, ...request.Option) (*forecastservice.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*forecastservice.DeleteDatasetInput) (*request.Request, *forecastservice.DeleteDatasetOutput)

	DeleteDatasetGroup(*forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupWithContext(aws.Context, *forecastservice.DeleteDatasetGroupInput, ...request.Option) (*forecastservice.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupRequest(*forecastservice.DeleteDatasetGroupInput) (*request.Request, *forecastservice.DeleteDatasetGroupOutput)

	DeleteDatasetImportJob(*forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error)
	DeleteDatasetImportJobWithContext(aws.Context, *forecastservice.DeleteDatasetImportJobInput, ...request.Option) (*forecastservice.DeleteDatasetImportJobOutput, error)
	DeleteDatasetImportJobRequest(*forecastservice.DeleteDatasetImportJobInput) (*request.Request, *forecastservice.DeleteDatasetImportJobOutput)

	DeleteForecast(*forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error)
	DeleteForecastWithContext(aws.Context, *forecastservice.DeleteForecastInput, ...request.Option) (*forecastservice.DeleteForecastOutput, error)
	DeleteForecastRequest(*forecastservice.DeleteForecastInput) (*request.Request, *forecastservice.DeleteForecastOutput)

	DeleteForecastExportJob(*forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error)
	DeleteForecastExportJobWithContext(aws.Context, *forecastservice.DeleteForecastExportJobInput, ...request.Option) (*forecastservice.DeleteForecastExportJobOutput, error)
	DeleteForecastExportJobRequest(*forecastservice.DeleteForecastExportJobInput) (*request.Request, *forecastservice.DeleteForecastExportJobOutput)

	DeletePredictor(*forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error)
	DeletePredictorWithContext(aws.Context, *forecastservice.DeletePredictorInput, ...request.Option) (*forecastservice.DeletePredictorOutput, error)
	DeletePredictorRequest(*forecastservice.DeletePredictorInput) (*request.Request, *forecastservice.DeletePredictorOutput)

	DescribeDataset(*forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *forecastservice.DescribeDatasetInput, ...request.Option) (*forecastservice.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*forecastservice.DescribeDatasetInput) (*request.Request, *forecastservice.DescribeDatasetOutput)

	DescribeDatasetGroup(*forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupWithContext(aws.Context, *forecastservice.DescribeDatasetGroupInput, ...request.Option) (*forecastservice.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupRequest(*forecastservice.DescribeDatasetGroupInput) (*request.Request, *forecastservice.DescribeDatasetGroupOutput)

	DescribeDatasetImportJob(*forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobWithContext(aws.Context, *forecastservice.DescribeDatasetImportJobInput, ...request.Option) (*forecastservice.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobRequest(*forecastservice.DescribeDatasetImportJobInput) (*request.Request, *forecastservice.DescribeDatasetImportJobOutput)

	DescribeForecast(*forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error)
	DescribeForecastWithContext(aws.Context, *forecastservice.DescribeForecastInput, ...request.Option) (*forecastservice.DescribeForecastOutput, error)
	DescribeForecastRequest(*forecastservice.DescribeForecastInput) (*request.Request, *forecastservice.DescribeForecastOutput)

	DescribeForecastExportJob(*forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error)
	DescribeForecastExportJobWithContext(aws.Context, *forecastservice.DescribeForecastExportJobInput, ...request.Option) (*forecastservice.DescribeForecastExportJobOutput, error)
	DescribeForecastExportJobRequest(*forecastservice.DescribeForecastExportJobInput) (*request.Request, *forecastservice.DescribeForecastExportJobOutput)

	DescribePredictor(*forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error)
	DescribePredictorWithContext(aws.Context, *forecastservice.DescribePredictorInput, ...request.Option) (*forecastservice.DescribePredictorOutput, error)
	DescribePredictorRequest(*forecastservice.DescribePredictorInput) (*request.Request, *forecastservice.DescribePredictorOutput)

	GetAccuracyMetrics(*forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error)
	GetAccuracyMetricsWithContext(aws.Context, *forecastservice.GetAccuracyMetricsInput, ...request.Option) (*forecastservice.GetAccuracyMetricsOutput, error)
	GetAccuracyMetricsRequest(*forecastservice.GetAccuracyMetricsInput) (*request.Request, *forecastservice.GetAccuracyMetricsOutput)

	ListDatasetGroups(*forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error)
	ListDatasetGroupsWithContext(aws.Context, *forecastservice.ListDatasetGroupsInput, ...request.Option) (*forecastservice.ListDatasetGroupsOutput, error)
	ListDatasetGroupsRequest(*forecastservice.ListDatasetGroupsInput) (*request.Request, *forecastservice.ListDatasetGroupsOutput)

	ListDatasetGroupsPages(*forecastservice.ListDatasetGroupsInput, func(*forecastservice.ListDatasetGroupsOutput, bool) bool) error
	ListDatasetGroupsPagesWithContext(aws.Context, *forecastservice.ListDatasetGroupsInput, func(*forecastservice.ListDatasetGroupsOutput, bool) bool, ...request.Option) error

	ListDatasetImportJobs(*forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsWithContext(aws.Context, *forecastservice.ListDatasetImportJobsInput, ...request.Option) (*forecastservice.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsRequest(*forecastservice.ListDatasetImportJobsInput) (*request.Request, *forecastservice.ListDatasetImportJobsOutput)

	ListDatasetImportJobsPages(*forecastservice.ListDatasetImportJobsInput, func(*forecastservice.ListDatasetImportJobsOutput, bool) bool) error
	ListDatasetImportJobsPagesWithContext(aws.Context, *forecastservice.ListDatasetImportJobsInput, func(*forecastservice.ListDatasetImportJobsOutput, bool) bool, ...request.Option) error

	ListDatasets(*forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *forecastservice.ListDatasetsInput, ...request.Option) (*forecastservice.ListDatasetsOutput, error)
	ListDatasetsRequest(*forecastservice.ListDatasetsInput) (*request.Request, *forecastservice.ListDatasetsOutput)

	ListDatasetsPages(*forecastservice.ListDatasetsInput, func(*forecastservice.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *forecastservice.ListDatasetsInput, func(*forecastservice.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListForecastExportJobs(*forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error)
	ListForecastExportJobsWithContext(aws.Context, *forecastservice.ListForecastExportJobsInput, ...request.Option) (*forecastservice.ListForecastExportJobsOutput, error)
	ListForecastExportJobsRequest(*forecastservice.ListForecastExportJobsInput) (*request.Request, *forecastservice.ListForecastExportJobsOutput)

	ListForecastExportJobsPages(*forecastservice.ListForecastExportJobsInput, func(*forecastservice.ListForecastExportJobsOutput, bool) bool) error
	ListForecastExportJobsPagesWithContext(aws.Context, *forecastservice.ListForecastExportJobsInput, func(*forecastservice.ListForecastExportJobsOutput, bool) bool, ...request.Option) error

	ListForecasts(*forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error)
	ListForecastsWithContext(aws.Context, *forecastservice.ListForecastsInput, ...request.Option) (*forecastservice.ListForecastsOutput, error)
	ListForecastsRequest(*forecastservice.ListForecastsInput) (*request.Request, *forecastservice.ListForecastsOutput)

	ListForecastsPages(*forecastservice.ListForecastsInput, func(*forecastservice.ListForecastsOutput, bool) bool) error
	ListForecastsPagesWithContext(aws.Context, *forecastservice.ListForecastsInput, func(*forecastservice.ListForecastsOutput, bool) bool, ...request.Option) error

	ListPredictors(*forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error)
	ListPredictorsWithContext(aws.Context, *forecastservice.ListPredictorsInput, ...request.Option) (*forecastservice.ListPredictorsOutput, error)
	ListPredictorsRequest(*forecastservice.ListPredictorsInput) (*request.Request, *forecastservice.ListPredictorsOutput)

	ListPredictorsPages(*forecastservice.ListPredictorsInput, func(*forecastservice.ListPredictorsOutput, bool) bool) error
	ListPredictorsPagesWithContext(aws.Context, *forecastservice.ListPredictorsInput, func(*forecastservice.ListPredictorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *forecastservice.ListTagsForResourceInput, ...request.Option) (*forecastservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*forecastservice.ListTagsForResourceInput) (*request.Request, *forecastservice.ListTagsForResourceOutput)

	TagResource(*forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *forecastservice.TagResourceInput, ...request.Option) (*forecastservice.TagResourceOutput, error)
	TagResourceRequest(*forecastservice.TagResourceInput) (*request.Request, *forecastservice.TagResourceOutput)

	UntagResource(*forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *forecastservice.UntagResourceInput, ...request.Option) (*forecastservice.UntagResourceOutput, error)
	UntagResourceRequest(*forecastservice.UntagResourceInput) (*request.Request, *forecastservice.UntagResourceOutput)

	UpdateDatasetGroup(*forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error)
	UpdateDatasetGroupWithContext(aws.Context, *forecastservice.UpdateDatasetGroupInput, ...request.Option) (*forecastservice.UpdateDatasetGroupOutput, error)
	UpdateDatasetGroupRequest(*forecastservice.UpdateDatasetGroupInput) (*request.Request, *forecastservice.UpdateDatasetGroupOutput)
}

var _ ForecastServiceAPI = (*forecastservice.ForecastService)(nil)
