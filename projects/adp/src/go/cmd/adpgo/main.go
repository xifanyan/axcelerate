package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	cli "github.com/urfave/cli/v3"
	adp "github.com/xifanyan/axcelerate/adp"
)

const pollInterval = 0

func main() {
	if err := newApp(os.Stdout, os.Stderr).Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}

func newApp(stdout io.Writer, stderr io.Writer) *cli.Command {
	return &cli.Command{
		Name:      "adpgo",
		Usage:     "ADP Go CLI",
		Writer:    stdout,
		ErrWriter: stderr,
		ExitErrHandler: func(_ context.Context, _ *cli.Command, _ error) {
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "host", Usage: "ADP server host", Required: true},
			&cli.IntFlag{Name: "port", Usage: "ADP server port", Value: 8443},
			&cli.StringFlag{Name: "path", Usage: "ADP task API path", Value: "/adp/rest/api/task"},
			&cli.StringFlag{Name: "user", Usage: "ADP username", Required: true},
			&cli.StringFlag{Name: "password", Usage: "ADP password", Required: true},
			&cli.BoolFlag{Name: "insecure", Usage: "Skip TLS certificate verification"},
			&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Usage: "Enable debug logging"},
		},
		Commands: []*cli.Command{
			{
				Name: "list-entities",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "type"},
					&cli.StringFlag{Name: "id"},
					&cli.StringFlag{Name: "relatedEntity"},
					&cli.StringFlag{Name: "whiteList"},
					&cli.StringFlag{Name: "workspace"},
					&cli.StringFlag{Name: "status"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewListEntitiesBuilder(client)
					applyString(cmd, "type", builder.Type)
					applyString(cmd, "id", builder.ID)
					applyString(cmd, "relatedEntity", builder.RelatedEntity)
					applyString(cmd, "whiteList", builder.WhiteList)
					applyString(cmd, "workspace", builder.Workspace)
					applyString(cmd, "status", builder.Status)

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result.Entities)
				},
			},
			{
				Name: "query-engine",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "engineName"},
					&cli.StringFlag{Name: "engineQuery"},
					&cli.StringFlag{Name: "engineUserName"},
					&cli.StringFlag{Name: "engineUserPassword"},
					&cli.StringSliceFlag{Name: "engineTaxonomies"},
					&cli.StringFlag{Name: "applicationIdentifier"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewQueryEngineBuilder(client)
					applyString(cmd, "engineName", builder.EngineName)
					applyString(cmd, "engineQuery", builder.EngineQuery)
					applyString(cmd, "engineUserName", builder.EngineUserName)
					applyString(cmd, "engineUserPassword", builder.EngineUserPassword)
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)

					if cmd.IsSet("engineTaxonomies") {
						taxonomies, err := parseEngineTaxonomies(cmd.StringSlice("engineTaxonomies"))
						if err != nil {
							return err
						}
						builder.EngineTaxonomies(taxonomies)
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
			{
				Name: "taxonomy-statistic",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "engineName"},
					&cli.StringFlag{Name: "engineQuery"},
					&cli.BoolFlag{Name: "computeCounts", Value: true},
					&cli.BoolFlag{Name: "listCategoryProperties"},
					&cli.StringSliceFlag{Name: "engineTaxonomies"},
					&cli.StringFlag{Name: "outputTaxonomies"},
					&cli.StringFlag{Name: "applicationIdentifier"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewTaxonomyStatisticBuilder(client)
					applyString(cmd, "engineName", builder.EngineName)
					applyString(cmd, "engineQuery", builder.EngineQuery)
					if cmd.IsSet("computeCounts") {
						builder.ComputeCounts(cmd.Bool("computeCounts"))
					}
					if cmd.IsSet("listCategoryProperties") {
						builder.ListCategoryProperties(cmd.Bool("listCategoryProperties"))
					}
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)

					if cmd.IsSet("engineTaxonomies") {
						taxonomies, err := parseEngineTaxonomies(cmd.StringSlice("engineTaxonomies"))
						if err != nil {
							return err
						}
						builder.EngineTaxonomies(taxonomies)
					}
					if cmd.IsSet("outputTaxonomies") {
						outputTaxonomies, err := adp.ParseOutputTaxonomies(cmd.String("outputTaxonomies"))
						if err != nil {
							return err
						}
						builder.OutputTaxonomies(outputTaxonomies)
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result.Statistics)
				},
			},
			{
				Name: "start-application",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "applicationIdentifier"},
					&cli.BoolFlag{Name: "useHttps"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewStartApplicationBuilder(client)
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)
					if cmd.IsSet("useHttps") {
						builder.UseHTTPS(cmd.Bool("useHttps"))
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					_, err = fmt.Fprintln(stdout, result.ApplicationURL)
					return err
				},
			},
			{
				Name: "csv-merge",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "csvFile"},
					&cli.StringFlag{Name: "csvIdFieldKey"},
					&cli.StringFlag{Name: "mergeType"},
					&cli.StringFlag{Name: "csvMode"},
					&cli.StringFlag{Name: "engineName"},
					&cli.StringFlag{Name: "engineUser"},
					&cli.StringFlag{Name: "enginePassword"},
					&cli.StringFlag{Name: "engineIdFieldKey"},
					&cli.StringFlag{Name: "applicationIdentifier"},
					&cli.StringFlag{Name: "fieldSeparator"},
					&cli.StringFlag{Name: "imageBasePath"},
					&cli.StringFlag{Name: "nativeBasePath"},
					&cli.StringFlag{Name: "csvFieldImageLocation"},
					&cli.StringFlag{Name: "csvFieldNativeLocation"},
					&cli.StringFlag{Name: "multiValueDelimiter"},
					&cli.StringFlag{Name: "textIndicator"},
					&cli.BoolFlag{Name: "doNotChangeProtectedDocuments"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewCSVMergeBuilder(client)
					applyString(cmd, "csvFile", builder.CSVFile)
					applyString(cmd, "csvIdFieldKey", builder.CSVIDFieldKey)
					applyString(cmd, "mergeType", builder.MergeType)
					applyString(cmd, "csvMode", builder.CSVMode)
					applyString(cmd, "engineName", builder.EngineName)
					applyString(cmd, "engineUser", builder.EngineUser)
					applyString(cmd, "enginePassword", builder.EnginePassword)
					applyString(cmd, "engineIdFieldKey", builder.EngineIDFieldKey)
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)
					applyString(cmd, "fieldSeparator", builder.FieldSeparator)
					applyString(cmd, "imageBasePath", builder.ImageBasePath)
					applyString(cmd, "nativeBasePath", builder.NativeBasePath)
					applyString(cmd, "csvFieldImageLocation", builder.CSVFieldImageLocation)
					applyString(cmd, "csvFieldNativeLocation", builder.CSVFieldNativeLocation)
					applyString(cmd, "multiValueDelimiter", builder.MultiValueDelimiter)
					applyString(cmd, "textIndicator", builder.TextIndicator)
					if cmd.IsSet("doNotChangeProtectedDocuments") {
						builder.DoNotChangeProtectedDocuments(cmd.Bool("doNotChangeProtectedDocuments"))
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
			{
				Name: "export-documents",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "fieldSeparator"},
					&cli.BoolFlag{Name: "waitForExport"},
					&cli.StringFlag{Name: "query"},
					&cli.StringFlag{Name: "applicationIdentifier"},
					&cli.StringFlag{Name: "applicationType"},
					&cli.StringFlag{Name: "engineIdentifier"},
					&cli.StringFlag{Name: "engineUser"},
					&cli.StringFlag{Name: "enginePassword"},
					&cli.StringFlag{Name: "exportName"},
					&cli.StringFlag{Name: "exportFields"},
					&cli.StringFlag{Name: "exportDirectory"},
					&cli.StringFlag{Name: "fileEnding"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewExportDocumentsBuilder(client)
					applyString(cmd, "fieldSeparator", builder.FieldSeparator)
					if cmd.IsSet("waitForExport") {
						builder.WaitForExport(cmd.Bool("waitForExport"))
					}
					applyString(cmd, "query", builder.Query)
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)
					applyString(cmd, "applicationType", builder.ApplicationType)
					applyString(cmd, "engineIdentifier", builder.EngineIdentifier)
					applyString(cmd, "engineUser", builder.EngineUser)
					applyString(cmd, "enginePassword", builder.EnginePassword)
					applyString(cmd, "exportName", builder.ExportName)
					applyString(cmd, "exportFields", builder.ExportFields)
					applyString(cmd, "exportDirectory", builder.ExportDirectory)
					applyString(cmd, "fileEnding", builder.FileEnding)

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
			{
				Name: "read-configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "entityIdToRead"},
					&cli.StringFlag{Name: "configsToRead"},
					&cli.StringFlag{Name: "fileFormat"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewReadConfigurationBuilder(client)
					applyString(cmd, "entityIdToRead", builder.EntityIDToRead)
					applyString(cmd, "fileFormat", builder.FileFormat)
					if cmd.IsSet("configsToRead") {
						configs, err := adp.ParseConfigArgs(cmd.String("configsToRead"))
						if err != nil {
							return err
						}
						builder.ConfigsToRead(configs)
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result.Configuration)
				},
			},
			{
				Name: "create-ocr-job",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "engineName"},
					&cli.StringFlag{Name: "query"},
					&cli.StringFlag{Name: "engineUserName"},
					&cli.StringFlag{Name: "engineUserPassword"},
					&cli.StringFlag{Name: "jobName"},
					&cli.StringFlag{Name: "jobDescription"},
					&cli.IntFlag{Name: "jobPriority"},
					&cli.StringFlag{Name: "applicationIdentifier"},
					&cli.StringFlag{Name: "applicationType"},
					&cli.BoolFlag{Name: "wait"},
					&cli.StringFlag{Name: "engineType"},
					&cli.StringFlag{Name: "listOfJobProperties"},
					&cli.StringFlag{Name: "globalSearchJson"},
					&cli.StringFlag{Name: "globalSearchId"},
					&cli.StringSliceFlag{Name: "restrictions"},
					&cli.StringSliceFlag{Name: "advancedRestrictions"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewCreateOcrJobBuilder(client)
					applyString(cmd, "engineName", builder.EngineName)
					applyString(cmd, "query", builder.Query)
					applyString(cmd, "engineUserName", builder.EngineUserName)
					applyString(cmd, "engineUserPassword", builder.EngineUserPassword)
					applyString(cmd, "jobName", builder.JobName)
					applyString(cmd, "jobDescription", builder.JobDescription)
					if cmd.IsSet("jobPriority") {
						builder.JobPriority(cmd.Int("jobPriority"))
					}
					applyString(cmd, "applicationIdentifier", builder.ApplicationIdentifier)
					applyString(cmd, "applicationType", builder.ApplicationType)
					if cmd.IsSet("wait") {
						builder.WaitFlag(cmd.Bool("wait"))
					}
					applyString(cmd, "engineType", builder.EngineType)
					applyString(cmd, "listOfJobProperties", builder.ListOfJobProperties)
					applyString(cmd, "globalSearchJson", builder.GlobalSearchJSON)
					applyString(cmd, "globalSearchId", builder.GlobalSearchID)
					if cmd.IsSet("restrictions") {
						items, err := parseEngineTaxonomies(cmd.StringSlice("restrictions"))
						if err != nil {
							return err
						}
						builder.Restrictions(items)
					}
					if cmd.IsSet("advancedRestrictions") {
						items, err := parseEngineTaxonomies(cmd.StringSlice("advancedRestrictions"))
						if err != nil {
							return err
						}
						builder.AdvancedRestrictions(items)
					}

					result, err := builder.Wait(ctx, pollInterval)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
			{
				Name: "cli",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "batchScriptPath"},
					&cli.StringFlag{Name: "batchScriptParameters"},
					&cli.StringFlag{Name: "workingDirectory"},
					&cli.StringFlag{Name: "batchScriptJsonLogOutput"},
					&cli.BoolFlag{Name: "batchScriptRedirectLogging"},
					&cli.StringFlag{Name: "batchScriptPositiveExecutionCodes"},
					&cli.BoolFlag{Name: "batchScriptFilterPasswords", Value: true},
					&cli.StringFlag{Name: "batchScriptLoggingDirectory"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					client, err := newClient(cmd)
					if err != nil {
						return err
					}

					builder := adp.NewCLITaskBuilder(client)
					applyString(cmd, "batchScriptPath", builder.BatchScriptPath)
					applyString(cmd, "workingDirectory", builder.WorkingDirectory)
					applyString(cmd, "batchScriptJsonLogOutput", builder.BatchScriptJSONLogOutput)
					if cmd.IsSet("batchScriptRedirectLogging") {
						builder.BatchScriptRedirectLogging(cmd.Bool("batchScriptRedirectLogging"))
					}
					applyString(cmd, "batchScriptPositiveExecutionCodes", builder.BatchScriptPositiveExecutionCodes)
					if cmd.IsSet("batchScriptFilterPasswords") {
						builder.BatchScriptFilterPasswords(cmd.Bool("batchScriptFilterPasswords"))
					}
					applyString(cmd, "batchScriptLoggingDirectory", builder.BatchScriptLoggingDirectory)
					if cmd.IsSet("batchScriptParameters") {
						params, err := adp.ParseBatchScriptParameters(cmd.String("batchScriptParameters"))
						if err != nil {
							return err
						}
						builder.BatchScriptParameters(params)
					}

					result, err := builder.Execute(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
		},
	}
}

func newClient(cmd *cli.Command) (*adp.Client, error) {
	return adp.NewClient(adp.ClientConfig{
		BaseURL:  adp.MustBaseURL(cmd.String("host"), cmd.Int("port"), cmd.String("path")),
		Username: cmd.String("user"),
		Password: cmd.String("password"),
		Insecure: cmd.Bool("insecure"),
		Debug:    cmd.Bool("debug"),
		DebugOut: cmd.ErrWriter,
	})
}

func writeJSON(w io.Writer, value any) error {
	body, err := adp.PrettyJSON(value)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(body))
	return err
}

func printCommandError(w io.Writer, err error) error {
	var execErr *adp.TaskExecutionError
	if errors.As(err, &execErr) {
		if _, writeErr := fmt.Fprintf(w, "Error: %s\nExecutionID: %s\nTaskType: %s\n", execErr.ErrorMessage, execErr.ExecutionID, execErr.TaskType); writeErr != nil {
			return writeErr
		}
	}
	return err
}

func applyString[T any](cmd *cli.Command, name string, apply func(string) T) {
	if cmd.IsSet(name) {
		apply(cmd.String(name))
	}
}

func parseEngineTaxonomies(values []string) ([]adp.EngineTaxonomyArg, error) {
	items := make([]adp.EngineTaxonomyArg, 0, len(values))
	for _, value := range values {
		item, err := adp.ParseEngineTaxonomyArg(value)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
