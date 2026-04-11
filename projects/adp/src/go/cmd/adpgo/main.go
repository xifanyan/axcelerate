package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	cli "github.com/urfave/cli/v3"
	adp "github.com/xifanyan/axcelerate/adp"
)

const pollInterval = 0

type cliConfigFile struct {
	Host     *string `json:"host"`
	Port     *int    `json:"port"`
	Path     *string `json:"path"`
	User     *string `json:"user"`
	Password *string `json:"password"`
	Insecure *bool   `json:"insecure"`
	Debug    *bool   `json:"debug"`
}

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Args))
}

func run(stdout io.Writer, stderr io.Writer, args []string) int {
	err := newApp(stdout, stderr).Run(context.Background(), args)
	if err == nil {
		return 0
	}
	var execErr *adp.TaskExecutionError
	if errors.As(err, &execErr) {
		return 1
	}
	if _, writeErr := fmt.Fprintln(stderr, err); writeErr != nil {
		return 1
	}
	return 1
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
			&cli.StringFlag{Name: "host", Usage: "ADP server host"},
			&cli.IntFlag{Name: "port", Usage: "ADP server port", Value: 8443},
			&cli.StringFlag{Name: "path", Usage: "ADP task API path", Value: "/adp/rest/api/task"},
			&cli.StringFlag{Name: "user", Usage: "ADP username"},
			&cli.StringFlag{Name: "password", Usage: "ADP password"},
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
					&cli.StringFlag{Name: "engineName", Required: true},
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
					&cli.StringFlag{Name: "engineName", Required: true},
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
					&cli.StringFlag{Name: "csvFile", Required: true},
					&cli.StringFlag{Name: "csvIdFieldKey"},
					&cli.StringFlag{Name: "mergeType"},
					&cli.StringFlag{Name: "csvMode"},
					&cli.StringFlag{Name: "engineName"},
					&cli.StringFlag{Name: "engineUser"},
					&cli.StringFlag{Name: "enginePassword"},
					&cli.StringFlag{Name: "engineIdFieldKey"},
					&cli.StringFlag{Name: "applicationIdentifier"},
					&cli.StringFlag{Name: "fieldMappings"},
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
					if cmd.IsSet("fieldMappings") {
						var fieldMappings []map[string]any
						if err := json.Unmarshal([]byte(cmd.String("fieldMappings")), &fieldMappings); err != nil {
							return err
						}
						builder.FieldMappings(fieldMappings)
					}
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

					if cmd.Bool("wait") {
						result, err := builder.Wait(ctx, pollInterval)
						if err != nil {
							return printCommandError(stderr, err)
						}
						return writeJSON(stdout, result)
					}

					result, err := builder.Start(ctx)
					if err != nil {
						return printCommandError(stderr, err)
					}
					return writeJSON(stdout, result)
				},
			},
			{
				Name: "cli",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "batchScriptPath", Required: true},
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
	cfg, err := loadCLIConfigFile()
	if err != nil {
		return nil, err
	}

	host := strings.TrimSpace(resolvedString(cmd, "host", cfg.Host))
	user := resolvedString(cmd, "user", cfg.User)
	password := resolvedString(cmd, "password", cfg.Password)
	missing := make([]string, 0, 3)
	if host == "" {
		missing = append(missing, "host")
	}
	if strings.TrimSpace(user) == "" {
		missing = append(missing, "user")
	}
	if strings.TrimSpace(password) == "" {
		missing = append(missing, "password")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required global setting(s): %s", strings.Join(missing, ", "))
	}

	baseURL, err := validatedBaseURL(host, resolvedInt(cmd, "port", cfg.Port), resolvedString(cmd, "path", cfg.Path))
	if err != nil {
		return nil, fmt.Errorf("invalid host: %w", err)
	}

	return adp.NewClient(adp.ClientConfig{
		BaseURL:  baseURL,
		Username: user,
		Password: password,
		Insecure: resolvedBool(cmd, "insecure", cfg.Insecure),
		Debug:    resolvedBool(cmd, "debug", cfg.Debug),
		DebugOut: cmd.ErrWriter,
	})
}

func validatedBaseURL(host string, port int, path string) (_ string, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("%v", recovered)
		}
	}()

	return adp.MustBaseURL(host, port, path), nil
}

func loadCLIConfigFile() (cliConfigFile, error) {
	body, err := os.ReadFile("adp_config.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cliConfigFile{}, nil
		}
		return cliConfigFile{}, fmt.Errorf("read adp_config.json: %w", err)
	}

	var cfg cliConfigFile
	if err := json.Unmarshal(body, &cfg); err != nil {
		return cliConfigFile{}, fmt.Errorf("invalid adp_config.json: %w", err)
	}
	return cfg, nil
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

func resolvedString(cmd *cli.Command, name string, configValue *string) string {
	if cmd.IsSet(name) {
		return cmd.String(name)
	}
	if configValue != nil {
		return *configValue
	}
	return cmd.String(name)
}

func resolvedInt(cmd *cli.Command, name string, configValue *int) int {
	if cmd.IsSet(name) {
		return cmd.Int(name)
	}
	if configValue != nil {
		return *configValue
	}
	return cmd.Int(name)
}

func resolvedBool(cmd *cli.Command, name string, configValue *bool) bool {
	if cmd.IsSet(name) {
		return cmd.Bool(name)
	}
	if configValue != nil {
		return *configValue
	}
	return cmd.Bool(name)
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
