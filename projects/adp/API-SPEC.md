```
[
  {
    "taskType": "Write Configuration",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_writeConfiguration_parameter": "{adp_entities_json_output}",
      "adp_taskActive": true,
      "adp_writeConfiguration_dynAction": "update",
      "adp_writeConfiguration_fileFormat": "JSON",
      "adp_writeConfiguration_entityIdToWrite": "",
      "adp_executionPersistent": true,
      "adp_writeConfiguration_JSON": "",
      "adp_abortWfOnFailure": true,
      "adp_writeConfiguration_fromFile": false,
      "adp_cleanUpHistory": false,
      "adp_writeConfiguration_file": "input.json",
      "adp_loggingEnabled": true,
      "adp_writeConfiguration_convertUNC2Linux": "false",
      "adp_taskTimeout": 0
    },
    "taskDescription": "A Task to write configurations from JSON.",
    "taskDisplayName": "Write Configuration"
  },
  {
    "taskType": "Restart Processes",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_restartProcess_ignoreSuspended": "false",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_restartProcess_processIdentifiers": []
    },
    "taskDescription": "Restarts processes",
    "taskDisplayName": "Restart processes"
  },
  {
    "taskType": "PSS Create FieldTypes",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "pss_createFieldTypes_targetFieldTypesFile": "",
      "pss_createFieldTypes_inputModelDataJson": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "pss_createFieldTypes_allViewId": "",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "pss_createFieldTypes_outputFilename": "pss_fieldtypes_file",
      "pss_createFieldTypes_outputFieldTypesJson": null
    },
    "taskDescription": "Creates the Perceptiv FieldTypes-File.",
    "taskDisplayName": "PSS Create FieldTypes"
  },
  {
    "taskType": "PSS Create MindserverConfig",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "pss_createMindserverConfig_outputDataModelFilename": "pss_datamodel_file",
      "adp_executionPersistent": true,
      "pss_createMindserverConfig_inputCsvBaseDirectory": "",
      "pss_createMindserverConfig_inputDataModelFilename": "",
      "adp_abortWfOnFailure": true,
      "pss_createMindserverConfig_inputFieldTypesBaseDirectory": "",
      "pss_createMindserverConfig_outputApplicationFilename": "pss_application_file",
      "adp_cleanUpHistory": false,
      "pss_createMindserverConfig_inputFieldTypesFileName": "",
      "adp_loggingEnabled": true,
      "pss_createMindserverConfig_outputDataSourceFilename": "pss_datasource_file",
      "adp_taskTimeout": 0,
      "pss_createMindserverConfig_inputReviewPanelFile": "",
      "pss_createMindserverConfig_inputModelDataJson": ""
    },
    "taskDescription": "Creates the Perceptiv MindserverConfig.",
    "taskDisplayName": "PSS Create Perceptiv MindserverConfig"
  },
  {
    "taskType": "Configure Engine Security",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_configureEngineSecurity_documentLevelConceptSearch": "",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_configureEngineSecurity_documentLevelName": null,
      "adp_loggingEnabled": true,
      "adp_configureEngineSecurity_documentLevelEnableSiblingExpansion": "false",
      "adp_configureEngineSecurity_documentLevelRestriction": null,
      "adp_configureEngineSecurity_fieldLevelSecurityTable": [],
      "adp_configureEngineSecurity_documentLevelGlobalSearchId": "",
      "adp_configureEngineSecurity_applicationType": "",
      "adp_configureEngineSecurity_appendConfig": "true",
      "adp_configureEngineSecurity_engineUserPassword": "",
      "adp_configureEngineSecurity_documentLevelDescription": null,
      "adp_configureEngineSecurity_documentLevelExactMatch": "",
      "adp_configureEngineSecurity_documentLevelSearchString": "*",
      "adp_configureEngineSecurity_documentLevelFuzzySearch": "",
      "adp_configureEngineSecurity_documentLevelSiblingFields": "rm_attachmentroot",
      "adp_configureEngineSecurity_documentLevelGroups": null,
      "adp_configureEngineSecurity_applicationIdentifier": "",
      "adp_configureEngineSecurity_engineType": "true",
      "adp_cleanUpHistory": false,
      "adp_configureEngineSecurity_documentLevelSecurityTable": [],
      "adp_configureEngineSecurity_engineName": "{adp_engineId}",
      "adp_taskTimeout": 0,
      "adp_configureEngineSecurity_folderLevelSecurityTable": [],
      "adp_configureEngineSecurity_engineUserName": "{adp_user}",
      "adp_configureEngineSecurity_documentLevelGlobalSearchJson": "",
      "adp_configureEngineSecurity_documentLevelMainQueryType": "exact match search"
    },
    "taskDescription": "Task to define field, folder and document level security.",
    "taskDisplayName": "Configure engine security"
  },
  {
    "taskType": "Axcelerate Import",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_axcelerate_import_user": "",
      "adp_taskActive": true,
      "adp_axcelerate_import_runId": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_axcelerate_import_projectId": "",
      "adp_cleanUpHistory": false,
      "adp_axcelerate_import_archiveLocation": "",
      "adp_loggingEnabled": true,
      "adp_axcelerate_import_httpsKeystoreFile": null,
      "adp_axcelerate_import_archiveLocationLinux": "",
      "adp_axcelerate_import_httpsTrustCertificate": "",
      "adp_axcelerate_import_httpsPassword": "",
      "adp_axcelerate_import_password": "",
      "adp_taskTimeout": 0,
      "adp_axcelerate_import_requestTimeoutSeconds": 900,
      "adp_axcelerate_import_httpsAllowUntrustedHosts": "true",
      "adp_axcelerate_import_connectTimeoutSeconds": 300
    },
    "taskDescription": "Import Axclerate NG data for a particular project.",
    "taskDisplayName": "Axcelerate Import"
  },
  {
    "taskType": "APWT OCR",
    "taskConfiguration": {
      "adp_poemOutputXml": false,
      "adp_progressTaskTimeout": 0,
      "adp_poemOutputPdf": true,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_ocrType": "POEM",
      "adp_poemLogFolder": null,
      "adp_cleanUpHistory": false,
      "adp_poemColumnMode": "AUTO",
      "adp_ocr_tempDirectory": null,
      "adp_poemOutputDocx": true,
      "adp_poemDocxFormat": "RFP",
      "adp_loggingEnabled": true,
      "adp_ocrBaseFolder": null,
      "adp_existingOcrResults": null,
      "adp_poemTableDetection": false,
      "adp_taskTimeout": 0,
      "adp_poemDpi": 300,
      "adp_poemPort": -2147483648,
      "adp_ocrBatchManagerNumberChannels": -2147483648,
      "adp_poemServer": null,
      "adp_poemTableMode": "AUTO",
      "adp_poemOutputStat": false
    },
    "taskDescription": "Executes OCR on chain files",
    "taskDisplayName": "APWT OCR Task"
  },
  {
    "taskType": "Create Data Source",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_createDataSource_abortOnExistingDataSource": false,
      "adp_createDataSource_applicationIdentifier": null,
      "adp_taskActive": true,
      "adp_createDataSource_choosenHostNameParameter": "adp_hostname",
      "adp_executionPersistent": true,
      "adp_createDataSource_choosenHostMemoryRatio": "adp_chosen_host_memory_ratio",
      "adp_abortWfOnFailure": true,
      "adp_createDataSource_chosenHostCpuLoad": "adp_chosen_host_cpu_load",
      "adp_loggingEnabled": true,
      "adp_createDataSource_dataSourceSystemTemplateDisplayName": "Server - file share",
      "adp_createDataSource_usedTemplate": "adp_used_data_source_template",
      "adp_createDataSource_hostCpuLoadThreshold": "50",
      "adp_createDataSource_createdDataSourceNameParameter": "adp_created_data_source_name",
      "adp_createDataSource_retryMaxNumberRunningCrawlers": "30",
      "adp_createDataSource_choosenHostMemory": "adp_chosen_host_memory",
      "adp_createDataSource_workspaceIdentifier": null,
      "adp_createDataSource_hostIdentifier": null,
      "adp_createDataSource_hostMemoryLimit": "0",
      "adp_createDataSource_maxNumberRunningCrawlers": "0",
      "adp_cleanUpHistory": false,
      "adp_createDataSource_engineIdentifier": null,
      "adp_createDataSource_engineBoxDocThreshold": "1000000",
      "adp_createDataSource_hostMemoryLimitRatio": "0",
      "adp_createDataSource_choosenEngineNameParameter": "adp_chosen_engine",
      "adp_createDataSource_hostRolesBlackList": null,
      "adp_createDataSource_dataSourceIdentifier": "{datasource_id}",
      "adp_taskTimeout": 0,
      "adp_createDataSource_createdDataSourceDisplaynameParameter": "adp_created_data_source_displayname",
      "adp_createDataSource_dataSourceTemplate": "",
      "adp_createDataSource_dataSourceName": "{datasource_name}"
    },
    "taskDescription": "Creates a new data source",
    "taskDisplayName": "Create data source"
  },
  {
    "taskType": "XPath",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_xPath_inputFileMaxSize": 1,
      "adp_xPath_inputFilePath": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_xPath_queries": [],
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_xPath_outputMetadataMaxSize": 256
    },
    "taskDescription": "Evaluates a set of XPath expressions on an XML document.",
    "taskDisplayName": "XPath"
  },
  {
    "taskType": "Create Review Interface",
    "taskConfiguration": {
      "adp_cri_appHostMemoryLimitRatio": "0",
      "adp_cri_matterSpecificWorkspace": null,
      "adp_progressTaskTimeout": 0,
      "adp_cri_matterId": "default",
      "adp_cri_matterSpecificEngineHostId": "localhost",
      "adp_cri_httpsPassword": "",
      "adp_cri_matterSpecificTemplate": null,
      "adp_taskActive": true,
      "adp_cri_reduceMemoryFootPrint": false,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cri_engineHostMemoryLimit": "0",
      "adp_loggingEnabled": true,
      "adp_cri_chosenEngineHostMemoryRatio": "adp_create_review_engine_host_memory_ratio",
      "adp_cri_webserviceRequestTimeoutSeconds": 900,
      "adp_cri_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_cri_publishEngineId": "adp_created_publish_engine_id",
      "adp_cri_appHostMemoryLimit": "0",
      "adp_cri_publishApplicationId": "adp_created_publish_application_id",
      "adp_cri_chosenApplicationHostMemory": "adp_create_review_application_host_memory",
      "adp_cri_chosenEngineHostMemory": "adp_create_review_engine_host_memory",
      "adp_cri_chosenEngineHostNameParameter": "adp_create_review_engine_host",
      "adp_cri_httpsKeystoreFile": null,
      "adp_cri_matterSpecificApplicationHostId": "",
      "adp_cri_matterSpecificApplication": null,
      "adp_cri_usedWebserviceUrl": "adp_url_used_for_create_review_interface",
      "adp_cri_appHostDetection": "true",
      "adp_cri_matterSpecificSmallProject": false,
      "adp_cri_engineHostMemoryLimitRatio": "0",
      "adp_cri_expectedNumberOfDocuments": 1000000,
      "adp_cri_predictiveCodingEnabled": false,
      "adp_cri_chosenApplicationHostNameParameter": "adp_create_review_application_host",
      "adp_cleanUpHistory": false,
      "adp_cri_httpsTrustCertificate": "",
      "adp_cri_httpsAllowUntrustedHosts": "false",
      "adp_cri_matterSpecificUrlRegEx": null,
      "adp_cri_chosenApplicationHostMemoryRatio": "adp_create_review_application_host_memory_ratio",
      "adp_taskTimeout": 0,
      "adp_cri_webserviceUser": "{service_user}",
      "adp_cri_webserviceConnectTimeoutSeconds": 300,
      "adp_cri_webservicePassword": null
    },
    "taskDescription": "",
    "taskDisplayName": "Task to create a review interface"
  },
  {
    "taskType": "Matter Entities",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_matterEntities_httpsPassword": "",
      "adp_matterEntities_matterId": null,
      "adp_matterEntities_webserviceUser": "{service_user}",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_matterEntities_webserviceRequestTimeoutSeconds": 900,
      "adp_matterEntities_usedWebserviceUrl": "adp_url_used_for_matter_reporting",
      "adp_cleanUpHistory": false,
      "adp_matterEntities_webservicePassword": null,
      "adp_matterEntities_httpsTrustCertificate": "",
      "adp_loggingEnabled": true,
      "adp_matterEntities_httpsKeystoreFile": null,
      "adp_matterEntities_webserviceConnectTimeoutSeconds": 300,
      "adp_matterEntities_httpsAllowUntrustedHosts": "false",
      "adp_matterEntities_processedMatterId": "adp_processed_matter_id",
      "adp_taskTimeout": 0,
      "adp_matterEntities_matterDisplayName": null,
      "adp_matterEntities_matterSpecificUrlRegEx": null,
      "adp_matterEntities_matterSpecificApplication": null,
      "adp_matterEntities_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_matterEntities_configuration": ""
    },
    "taskDescription": "Task to manage matter specific entities.",
    "taskDisplayName": "Matter Entities"
  },
  {
    "taskType": "Skip on value",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_skipOnValue_sourceFile": "script.js",
      "adp_skipOnValue_scriptCode": "returnValue = '{adp_current_task_name}' == adp_current_task_name;",
      "adp_skipOnValue_useSourceFile": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_skipOnValue_scriptEngine": "JavaScript",
      "adp_abortWfOnFailure": true,
      "adp_skipOnValue_targetTask": "myTask",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Task that will execute a java script condition and will jump to another task in the case the script returns true.",
    "taskDisplayName": "Skip on value"
  },
  {
    "taskType": "Create Workspace",
    "taskConfiguration": {
      "adp_createWorkspace_applicationWorkspace": null,
      "adp_createWorkspace_chosenWorkspaceNameParameter": "adp_create_workspace",
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_createWorkspace_applicationWorkspaceId": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_createWorkspace_createWorkspace": "false",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Creates a workspace",
    "taskDisplayName": "Create workspace"
  },
  {
    "taskType": "Export Ingestion Report",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_exportIngestionReport_applicationId": "{adp_applicationId}",
      "adp_taskActive": true,
      "adp_exportIngestionReport_engineUserName": "{adp_user}",
      "adp_exportIngestionReport_reportJsonConfiguration": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_exportIngestionReport_engineId": "{adp_engineId}",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_exportIngestionReport_reportFilePath": "adp_exportIngestionReport_reportFilePath",
      "adp_exportIngestionReport_engineUserPassword": ""
    },
    "taskDescription": "Task to export ingestion report.",
    "taskDisplayName": "Export Ingestion Report"
  },
  {
    "taskType": "Taxonomy Statistic",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taxonomyStatistic_outputJsonAbsFilePath": "adp_taxonomy_statistics_json_file_path",
      "adp_taxonomyStatistic_applicationIdentifier": "",
      "adp_taskActive": true,
      "adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType": null,
      "adp_executionPersistent": true,
      "adp_taxonomyStatistic_engineUserName": "{adp_user}",
      "adp_abortWfOnFailure": true,
      "adp_taxonomyStatistic_applicationType": "",
      "adp_taxonomyStatistic_computeCounts": "true",
      "adp_loggingEnabled": true,
      "adp_taxonomyStatistic_outputJsonFilePath": null,
      "adp_taxonomyStatistic_engineTaxonomies": [],
      "adp_taxonomyStatistic_engineUserPassword": "",
      "adp_taxonomyStatistic_outputXmlAbsFilePath": "adp_taxonomy_statistics_xml_file_path",
      "adp_taxonomyStatistic_engineQuery": "*",
      "adp_taxonomyStatistic_listCategoryProperties": "false",
      "adp_taxonomyStatistic_outputTaxonomies": [],
      "adp_taxonomyStatistic_outputJson": "adp_taxonomy_statistics_json_output",
      "adp_taxonomyStatistic_engineType": "true",
      "adp_cleanUpHistory": false,
      "adp_taxonomyStatistic_outputXmlFilePath": null,
      "adp_taxonomyStatistic_outputFields": [],
      "adp_taxonomyStatistic_engineGlobalSearch": "",
      "adp_taxonomyStatistic_listDocuments": "false",
      "adp_taskTimeout": 0,
      "adp_taxonomyStatistic_engineName": "{adp_engine_name}"
    },
    "taskDescription": "Retrieves category counts for a taxonomy",
    "taskDisplayName": "Taxonomy statistic"
  },
  {
    "taskType": "List Bucket",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_listBucket_rootPath": "",
      "adp_listBucket_accessKey": "",
      "adp_listBucket_entriesJson": "adp_list_bucket_entries_json",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_listBucket_cloudServiceName": "Amazon S3",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_listBucket_maxNumRetries": "",
      "adp_loggingEnabled": true,
      "adp_listBucket_pathRegex": ".*",
      "adp_taskTimeout": 0,
      "adp_listBucket_bucketId": "",
      "adp_listBucket_chunkSize": 5000,
      "adp_listBucket_maxNumberResults": 10000,
      "adp_listBucket_secretKey": "",
      "adp_listBucket_numberThreads": ""
    },
    "taskDescription": "Returns a JSON list of the entries of the specified bucket.",
    "taskDisplayName": "List bucket task"
  },
  {
    "taskType": "Start Workflow",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_startWorkflow_inputMetaDataJson": null,
      "adp_startWorkflow_startedWorkflowMetaDataJson": "adp_startWorkflow_output_workflowDirectory",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_startWorkflow_workflowExecutionId": "adp_startWorkflow_output_workflowExecutionId",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_startWorkflow_adpPassword": null,
      "adp_startWorkflow_synchronous": true,
      "adp_taskTimeout": 0,
      "adp_startWorkflow_adpUser": null,
      "adp_startWorkflow_adpWorkflowName": null
    },
    "taskDescription": "Task that will start an ADP workflow.",
    "taskDisplayName": "Start Workflow"
  },
  {
    "taskType": "APWT Standalone Amendment Post Processor",
    "taskConfiguration": {
      "adp_sap_amendedFields": [],
      "adp_progressTaskTimeout": 0,
      "adp_sap_fieldTypes_filename": null,
      "adp_taskActive": true,
      "adp_sap_tempDirectory": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_sap_amendmentStartTagName": "tx.p",
      "adp_sap_fieldTypes_path": null,
      "adp_apwt_checkReviewState": false,
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_sap_textTagName": "text",
      "adp_sap_useDocumentProxy": false,
      "adp_sap_index_transaction_id": null
    },
    "taskDescription": "Post processing for standalone procsessing request",
    "taskDisplayName": "APWT Standalone Amendment Post Processor Task"
  },
  {
    "taskType": "Read from LDAP",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_readFromLdap_outputToFile": true,
      "adp_readFromLdap_csvOutputFilter": "(objectclass=person)",
      "adp_readFromLdap_csvOutputFile": "output.csv",
      "adp_readFromLdap_csvOutputBaseDN": "dc=com",
      "adp_readFromLdap_anonymous": false,
      "adp_taskActive": true,
      "adp_readFromLdap_userPassword": null,
      "adp_readFromLdap_csvOutputAtributes": "givenName, sn, l, mail",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_readFromLdap_url": "ldap://localhost:389",
      "adp_loggingEnabled": true,
      "adp_readFromLdap_dnToMetaData": [],
      "adp_taskTimeout": 0,
      "adp_readFromLdap_userName": "cn=Manager",
      "adp_readFromLdap_pageSize": 50,
      "adp_readFromLdap_csvOutputFileMetaDataKey": "ldap_task_output"
    },
    "taskDescription": "Reads meta data and many more from an LDAP source.",
    "taskDisplayName": "Read from LDAP"
  },
  {
    "taskType": "Change Display Name",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_changeDisplayName_displayName": null,
      "adp_changeDisplayName_configIdentifier": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Changes the display name of an application.",
    "taskDisplayName": "Change display name"
  },
  {
    "taskType": "PSS Create CSV Files",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "pss_createCsvFiles_inputModelDataJson": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "pss_createCsvFiles_outputDirectory": "pss_csv_files_location",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "pss_createCsvFiles_inputCsvDirectory": "",
      "adp_cleanUpHistory": false,
      "pss_createCsvFiles_allViewId": ""
    },
    "taskDescription": "Creates the Perceptiv CSV Files.",
    "taskDisplayName": "PSS Create CSV Files"
  },
  {
    "taskType": "Query Engine",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_queryEngine_fieldName": "virtual_filesize",
      "adp_queryEngine_enableSiblingExpansion": "false",
      "adp_queryEngine_engineName": "{adp_engineName}",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_queryEngine_engineUserPassword": "",
      "adp_abortWfOnFailure": true,
      "adp_loggingEnabled": true,
      "adp_queryEngine_engineTaxonomies": [],
      "adp_queryEngine_engineUserName": "{adp_user}",
      "adp_queryEngine_engineType": "true",
      "adp_queryEngine_saveVariable": "{engine_save}",
      "adp_queryEngine_categoryToDelete": "",
      "adp_queryEngine_activateCategoryDeletion": false,
      "adp_queryEngine_applicationIdentifier": "",
      "adp_queryEngine_taxonomyToDelete": "",
      "adp_queryEngine_successIfCountIs": "{adp_expectedDsDoccount}",
      "adp_queryEngine_category": "",
      "adp_queryEngine_activateTagging": false,
      "adp_queryEngine_globalSearchId": "",
      "adp_queryEngine_aggregatedValue": "adp_query_engine_aggregated_value",
      "adp_queryEngine_AdvancedRestrictions": [],
      "adp_queryEngine_taxonomy": "",
      "adp_queryEngine_globalSearchJson": "",
      "adp_queryEngine_saveCompareString": "true",
      "adp_cleanUpHistory": false,
      "adp_queryEngine_numberOfDocuments": "adp_query_engine_documents_count",
      "adp_queryEngine_siblingFields": "rm_attachmentroot",
      "adp_queryEngine_engineQuery": "*",
      "adp_queryEngine_mainQueryType": null,
      "adp_queryEngine_waitForResult": false,
      "adp_queryEngine_categoryDisplayName": "",
      "adp_queryEngine_waitWhileCountIs": "{adp_oldDsDoccount}",
      "adp_taskTimeout": 0,
      "adp_queryEngine_applicationType": "",
      "adp_queryEngine_exitOnValueChanged": true
    },
    "taskDescription": "Queries an engine",
    "taskDisplayName": "Query engine"
  },
  {
    "taskType": "CSV Validation",
    "taskConfiguration": {
      "adp_csvv_numberOfTotalReferencedNativeFiles": "validated_number_of_total_referenced_nativefiles",
      "adp_csvv_csvFile": null,
      "adp_executionPersistent": true,
      "adp_csvv_validationReportLevel": "Fulfilled",
      "adp_csvv_headerColumns": "validated_header_columns",
      "adp_csvv_usedLoadFileFormatIdentifier": "used_load_file_format_id",
      "adp_csvv_validatedOpticonFileFieldSeparator": "validated_opticon_file_fieldseparator",
      "adp_csvv_fieldSeparator": "{detected_field_separator}",
      "file_checkups": [
        {
          "Checkup id": "loadfile_checkup_fileexistsvalidation",
          "Checkup name": "Checks if exists",
          "Description": "Checks for the existence of files referenced by file fields types defined in the Fields table.",
          "Enabled": true,
          "Error on failure": true
        },
        {
          "Checkup id": "loadfile_checkup_filesizevalidation",
          "Checkup name": "Checks not empty",
          "Description": "Checks the size of of files referenced by file field types defined in the Fields table.",
          "Enabled": true,
          "Error on failure": true
        },
        {
          "Checkup id": "loadfile_checkup_fileexvalidation",
          "Checkup name": "Checks for extension",
          "Description": "Checks for the existence of file extensions referenced by file field types in the Fields table.",
          "Enabled": true,
          "Error on failure": true
        },
        {
          "Checkup id": "loadfile_checkup_filepathlengthvalidation",
          "Checkup name": "Checks maximum path length",
          "Description": "Checks the size of file paths referenced by file field types in the Fields table.",
          "Enabled": true,
          "Error on failure": true
        }
      ],
      "loadfile_checkup_mandatoryfields_enabled": true,
      "adp_csvv_validatedOpticonFileCharset": "validated_opticon_file_encoding",
      "loadfile_fields": [
        {
          "Field name": "id",
          "Mandatory": true,
          "Type": "ID",
          "Multi-value separator": "",
          "Mapped to": "",
          "x": ""
        }
      ],
      "adp_csvv_overwriteSchemaParameter": false,
      "loadfile_detection_error_on_additional_fields": false,
      "adp_csvv_projectId": "",
      "adp_csvv_numberOfTotalReferencedTextFiles": "validated_number_of_total_referenced_textfiles",
      "loadfile_prepend_ticket_to_id": false,
      "adp_csvv_numberOfValidReferencedNativeFiles": "validated_number_of_valid_referenced_nativefiles",
      "adp_csvv_validatedCsvFile": "validated_csv_file",
      "loadfile_companion_auto_correct_image_ref_enabled": false,
      "adp_csvv_numberOfValidReferencedImageFiles": "validated_number_of_valid_referenced_imagefiles",
      "inputtimeformats": [
        {
          "Time format": "h:m:s a z",
          "Locale": "US"
        },
        {
          "Time format": "h:m:s a Z",
          "Locale": "US"
        },
        {
          "Time format": "h:m:s a",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s z",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s z",
          "Locale": ""
        },
        {
          "Time format": "H:m:s Z",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s",
          "Locale": "US"
        },
        {
          "Time format": "H:m a z",
          "Locale": "US"
        },
        {
          "Time format": "H:m a Z",
          "Locale": "US"
        },
        {
          "Time format": "H:m a",
          "Locale": "US"
        },
        {
          "Time format": "h:m a z",
          "Locale": "US"
        },
        {
          "Time format": "h:m a Z",
          "Locale": "US"
        },
        {
          "Time format": "h:m a",
          "Locale": "US"
        },
        {
          "Time format": "h:mma z",
          "Locale": "GB"
        },
        {
          "Time format": "h:mma Z",
          "Locale": "GB"
        },
        {
          "Time format": "h:mma",
          "Locale": "GB"
        },
        {
          "Time format": "H:m z",
          "Locale": "US"
        },
        {
          "Time format": "H:m Z",
          "Locale": "US"
        },
        {
          "Time format": "H:m",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s-z",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s-Z",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s y",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s '+'",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s z y",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s Z y",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s.z",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s.Z",
          "Locale": "US"
        },
        {
          "Time format": "H'' m.s z",
          "Locale": "US"
        },
        {
          "Time format": "H'' m.s Z",
          "Locale": "US"
        },
        {
          "Time format": "H'' m.s",
          "Locale": "US"
        },
        {
          "Time format": "H:m:ssz",
          "Locale": "US"
        },
        {
          "Time format": "H:m:ssZ",
          "Locale": "US"
        },
        {
          "Time format": "H:m:ss",
          "Locale": "US"
        },
        {
          "Time format": "h:mma z",
          "Locale": "US"
        },
        {
          "Time format": "h:mma Z",
          "Locale": "US"
        },
        {
          "Time format": "h:mma",
          "Locale": "US"
        },
        {
          "Time format": "hh:mm a",
          "Locale": "US"
        },
        {
          "Time format": "H:m:s.S",
          "Locale": "US"
        },
        {
          "Time format": "H:mmZ",
          "Locale": "US"
        },
        {
          "Time format": "HHmmss,SSZZZZZ",
          "Locale": ""
        }
      ],
      "adp_csvv_csvErrorFindingFile": null,
      "loadfile_checkup_regexpattern_enabled": true,
      "loadfile_merge_companion": false,
      "adp_progressTaskTimeout": 0,
      "loadfile_checkup_filepathmaxlength": 255,
      "adp_csvv_reports": "validation_reports",
      "adp_taskActive": true,
      "adp_csvv_successfullyMappedHeaderColumns": "validated_successfully_mapped_header_columns",
      "adp_csvv_pathToCsvErrorFindingFile": "csv_error_finding_file",
      "loadfile_auto_fix_file_references": [],
      "loadfile_no_header": false,
      "loadfile_keep_temporary_csv_merge_file": false,
      "adp_csvv_numberOfWarnings": "validated_number_of_warnings",
      "loadfile_checkup_count_number_of_records": true,
      "adp_csvv_numberOfErrors": "validated_number_of_errors",
      "adp_csvv_validatedOpticonFileLineSeparator": "validated_opticon_file_lineseparator",
      "loadfile_field_auto_correction": [
        {
          "Field name": "id",
          "Ignore case": true,
          "Normalize": false,
          "Aliases": ""
        }
      ],
      "loadfile_companion_normalization_enabled": false,
      "loadfile_checkup_filepathvalidation_sizes": [
        {
          "File suffix": "xls",
          "Max path size": 217
        },
        {
          "File suffix": "doc",
          "Max path size": 245
        }
      ],
      "adp_csvv_csvWarningFindingFile": null,
      "adp_csvv_successfullyMappedFieldsMissingInDataModel": "validated_successfully_mapped_fields_missing_in_data_model",
      "adp_csvv_dataSourceConfigurationDiff": "data_source_configuration_diff",
      "loadfile_field_image_key": null,
      "lengthLimitForCategories": 1024,
      "adp_csvv_errorOnMissingOpticonIds": "true",
      "loadfile_preprocessing": true,
      "loadfile_checkup_datetimeformat_enabled": true,
      "loadfile_perl_substitutions": [],
      "adp_csvv_loadFormatName": "",
      "loadfile_check_endbates_in_opticon": false,
      "adp_csvv_validatedOpticonFile": "validated_opticon_file",
      "adp_csvv_opticonFile": null,
      "adp_csvv_emptyColumns": "validated_unmapped_empty_columns",
      "adp_csvv_validatedOpticonQuoteCharacter": "validated_opticon_file_quotecharacter",
      "adp_csvv_numberOfTotalReferencedImageFiles": "validated_number_of_total_referenced_imagefiles",
      "loadfile_checkup_detect_mandatory_fields": true,
      "adp_csvv_backupOpticonFile": "backup_opticon_file",
      "adp_loggingEnabled": true,
      "loadfile_checkup_regexpatterns": [
        {
          "Field name": "id",
          "Pattern": ""
        }
      ],
      "orphane_checkups": [
        {
          "Checkup id": "loadfile_checkup_image_orphans",
          "Checkup name": "Check for images that are not referenced",
          "Description": "Check for images that are not referenced",
          "Enabled": true,
          "Error on failure": false
        },
        {
          "Checkup id": "loadfile_checkup_native_orphans",
          "Checkup name": "Check for natives that are not referenced",
          "Description": "Check for natives that are not referenced",
          "Enabled": true,
          "Error on failure": false
        },
        {
          "Checkup id": "loadfile_checkup_textfile_orphans",
          "Checkup name": "Check for text files that are not referenced",
          "Description": "Check for text-files that are not referenced",
          "Enabled": true,
          "Error on failure": false
        },
        {
          "Checkup id": "loadfile_checkup_external_document_orphans",
          "Checkup name": "Check for external documents that are not referenced",
          "Description": "Check for external documents that are not referenced",
          "Enabled": true,
          "Error on failure": false
        }
      ],
      "loadfile_checkup_unique_file_references": false,
      "loadfile_checkup_detect_additional_fields": true,
      "adp_csvv_numberOfRecords": "number_of_records",
      "loadfile_field_locations": [
        {
          "Field name": "id",
          "Location": ""
        }
      ],
      "image_checkups": [
        {
          "Checkup id": "loadfile_checkup_imageformatvalidation",
          "Checkup name": "Checks for supported image formats",
          "Description": "Checks for supported image formats referenced by image-fields defined in the fields-table",
          "Enabled": true,
          "Error on failure": false
        }
      ],
      "adp_csvv_useCsvValidationFormat": "false",
      "loadfile_perl_substitutions_enabled": false,
      "loadfile_checkup_count_file_references": true,
      "adp_csvv_unsuccessfullyMappedHeaderColumns": "validated_unsuccessfully_mapped_header_columns",
      "loadfile_checkup_iduniqueness_enabled": true,
      "adp_csvv_lineSeparator": "{detected_line_separator}",
      "adp_csvv_validatedCsvFileCharset": "validated_csv_file_encoding",
      "adp_csvv_ignoreOrphanedChecks": "false",
      "adp_csvv_doNotAbortOnAdditionalEmptyColumns": "false",
      "adp_csvv_dataModelName": "",
      "adp_csvv_numberOfValidReferencedTextFiles": "validated_number_of_valid_referenced_textfiles",
      "loadfile_additional_fieldmappings": [],
      "loadfile_prepend_ticket_to_id_length": 5,
      "adp_csvv_csvMergeConfiguration": "csv_merge_configuration",
      "adp_csvv_missingOptionalHeaderColumns": "validated_missing_optional_header_columns",
      "adp_csvv_encoding": "{detected_encoding}",
      "adp_csvv_pathToCsvWarningFindingFile": "csv_warning_finding_file",
      "adp_taskTimeout": 0,
      "adp_csvv_checkOnlyHeader": false,
      "adp_csvv_csvFilePattern": "",
      "adp_csvv_validatedCsvFileFieldSeparator": "validated_csv_file_fieldseparator",
      "inputdateformats": [
        {
          "Date format": "M/d/y",
          "Locale": "US"
        },
        {
          "Date format": "M/d/y",
          "Locale": ""
        },
        {
          "Date format": "yyyy-MM-d",
          "Locale": "US"
        },
        {
          "Date format": "E M/d/y",
          "Locale": "US"
        },
        {
          "Date format": "E M/d/y",
          "Locale": ""
        },
        {
          "Date format": "d/M/y",
          "Locale": "GB"
        },
        {
          "Date format": "d/M/y",
          "Locale": ""
        },
        {
          "Date format": "d MMM y",
          "Locale": "US"
        },
        {
          "Date format": ", d MMM y",
          "Locale": "US"
        },
        {
          "Date format": "y-M-d",
          "Locale": "US"
        },
        {
          "Date format": "E, d MMM y",
          "Locale": "US"
        },
        {
          "Date format": "E MMM d y",
          "Locale": "US"
        },
        {
          "Date format": "E,d MMM y",
          "Locale": "US"
        },
        {
          "Date format": "E, d, MMM y '+'",
          "Locale": "US"
        },
        {
          "Date format": "E, d, MMM y",
          "Locale": "US"
        },
        {
          "Date format": "E, MMM d y",
          "Locale": "US"
        },
        {
          "Date format": "E. d MMM y",
          "Locale": "US"
        },
        {
          "Date format": "E. d MMMyyyy",
          "Locale": "US"
        },
        {
          "Date format": "E, MMM d, y",
          "Locale": "US"
        },
        {
          "Date format": "EEEE, MMMM d, yyyy",
          "Locale": "US"
        },
        {
          "Date format": "MMM d, y",
          "Locale": "US"
        },
        {
          "Date format": "MMM. d y",
          "Locale": "US"
        },
        {
          "Date format": "y/M/d E",
          "Locale": "US"
        },
        {
          "Date format": "'xxx', d MMM y",
          "Locale": "US"
        },
        {
          "Date format": "y/M/d",
          "Locale": "US"
        },
        {
          "Date format": "MMM d,y",
          "Locale": "US"
        },
        {
          "Date format": "MMM d.y",
          "Locale": "US"
        },
        {
          "Date format": "MMMMdd, y",
          "Locale": "US"
        },
        {
          "Date format": "MMMMdd. y",
          "Locale": "US"
        },
        {
          "Date format": "d-M-y",
          "Locale": "US"
        },
        {
          "Date format": "d.M.y",
          "Locale": "US"
        },
        {
          "Date format": "yyyyMMdd",
          "Locale": ""
        }
      ],
      "adp_csvv_validatedCsvFileQuoteCharacter": "validated_csv_file_quotecharacter",
      "inputdatetimeformats": [
        {
          "Datetime format": "M/d/y h:m:s a z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y h:m:s a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y h:m:s a",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m:s z",
          "Locale": ""
        },
        {
          "Datetime format": "M/d/y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m a z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m a",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y h:m a z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y h:m a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y h:m a",
          "Locale": "US"
        },
        {
          "Datetime format": "E M/d/y h:m a z",
          "Locale": "US"
        },
        {
          "Datetime format": "E M/d/y h:m a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E M/d/y h:m a",
          "Locale": "US"
        },
        {
          "Datetime format": "d/M/y h:mma z",
          "Locale": "GB"
        },
        {
          "Datetime format": "d/M/y h:mma Z",
          "Locale": "GB"
        },
        {
          "Datetime format": "d/M/y h:mma",
          "Locale": "GB"
        },
        {
          "Datetime format": "M/d/y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y h:m:s a z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y h:m:s a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y h:m:s a",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": ", d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": ", d MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": ", d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m:s-z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m:s-Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E d MMM y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m:s y",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "E,d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "E,d MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E,d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d, MMM y H:m:s '+'",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d, MMM y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y, h:m a z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y, h:m a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y, h:m a",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d H:m:s z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d H:m:s Z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d H:m:s y",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m:s z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m:s Z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m Z y",
          "Locale": "US"
        },
        {
          "Datetime format": "E MMM d H:m y",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H:m:s.z",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H:m:s.Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H'' m.s z",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H'' m.s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMM y H'' m.s",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMMyyyy H:m:ssz",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMMyyyy H:m:ssZ",
          "Locale": "US"
        },
        {
          "Datetime format": "E. d MMMyyyy H:m:ss",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y h:m:s a z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y h:m:s a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y h:m:s a",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y h:mma z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y h:mma Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y h:mma",
          "Locale": "US"
        },
        {
          "Datetime format": "EEEE, MMMM d, yyyy hh:mm a",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y H:m z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y H:m Z",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y H:m",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m:s a z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m:s a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m:s a",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m a z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m a Z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y h:m a",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM. d y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM. d y H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM. d y H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "y/M/d E H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "y/M/d E H:m:s Z",
          "Locale": "US"
        },
        {
          "Datetime format": "y/M/d E H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "'xxx', d MMM y H:m:s z",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d'T'H:m:ssZ",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d'T'H:m:s.S",
          "Locale": "US"
        },
        {
          "Datetime format": "y-M-d'T'H:m:s",
          "Locale": "US"
        },
        {
          "Datetime format": "M/d/y'T'H:mmZ",
          "Locale": "US"
        },
        {
          "Datetime format": "E, MMM d, y",
          "Locale": "US"
        },
        {
          "Datetime format": "E, d MMM y",
          "Locale": "US"
        },
        {
          "Datetime format": "y/M/d",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d,y",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d.y",
          "Locale": "US"
        },
        {
          "Datetime format": "MMM d, y",
          "Locale": "US"
        },
        {
          "Datetime format": "MMMMdd, y",
          "Locale": "US"
        },
        {
          "Datetime format": "MMMMdd. y",
          "Locale": "US"
        },
        {
          "Datetime format": "d MMM y",
          "Locale": "US"
        },
        {
          "Datetime format": "d-M-y",
          "Locale": "US"
        },
        {
          "Datetime format": "d.M.y",
          "Locale": "US"
        },
        {
          "Datetime format": "yyyyMMdd'T'HHmmss,SSZZZZZ",
          "Locale": ""
        }
      ],
      "adp_csvv_validatedCsvFileLineSeparator": "validated_csv_file_lineseparator",
      "loadfile_prepend_ticket_to_id_pattern": "{ticket_id}-",
      "adp_csvv_valueDelimiter": "{detected_value_delimiter}",
      "adp_abortWfOnFailure": true,
      "loadfile_companion_error_on_missing_ids": true,
      "adp_csvv_csvMergeConfigurationAppendMode": false,
      "adp_csvv_aliasesFile": null,
      "adp_csvv_backupCsvFile": "backup_csv_file",
      "loadfile_companion_image_auto_correction_as_absolute_path": false,
      "adp_csvv_opticonFilePattern": "",
      "adp_cleanUpHistory": false,
      "loadfile_companion": false,
      "loadfile_auto_fix_file_references_enabled": false,
      "adp_csvv_aliasesFileEncoding": "UTF-8",
      "loadfile_prepend_ticket_to_id_check_pattern": "",
      "loadfile_checkup_detect_optional_fields": true,
      "loadfile_checkup_emptycols_enabled": true
    },
    "taskDescription": "Validates a csv file",
    "taskDisplayName": "Csv file validation task"
  },
  {
    "taskType": "Manage Users and Groups",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_manageUsersAndGroups_userDefinition": [],
      "adp_taskActive": true,
      "adp_manageUsersAndGroups_GroupUserIdsToFilterFor": "",
      "adp_manageUsersAndGroups_groupDefinition": [],
      "adp_manageUsersAndGroups_AppIdsToFilterFor": "",
      "adp_executionPersistent": true,
      "adp_manageUsersAndGroups_assignmentUserToGroup": [],
      "adp_manageUsersAndGroups_addApplicationRoles": [],
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_manageUsersAndGroups_file": "output.json",
      "adp_loggingEnabled": true,
      "adp_manageUsersAndGroups_outputFilename": "adp_manageUsersAndGroups_output_file_name",
      "adp_manageUsersAndGroups_outputJson": "adp_manageUsersAndGroups_json_output",
      "adp_taskTimeout": 0,
      "adp_manageUsersAndGroups_ReturnAllUsersUnderGroup": ""
    },
    "taskDescription": "Creates users, creates groups, assign users to group and set roles of users and groups in applications.",
    "taskDisplayName": "Manage users and groups"
  },
  {
    "taskType": "Manage Taxonomy",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_manageTaxonomy_enginePassword": "",
      "adp_manageTaxonomy_engineUser": "",
      "adp_manageTaxonomy_applicationName": "",
      "adp_manageTaxonomy_engineName": "",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_manageTaxonomy_taxonomiesJson": null,
      "adp_manageTaxonomy_engineType": "",
      "adp_taskTimeout": 0,
      "adp_manageTaxonomy_applicationType": "",
      "adp_manageTaxonomy_outputJson": "adp_manage_taxonomy_json_output"
    },
    "taskDescription": "Performs actions on a taxonomy like creating or updating categories, removing categories and quering categories",
    "taskDisplayName": "Manage taxonomy"
  },
  {
    "taskType": "CSV Multivalue Separator Detection",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvMultivalueSeparatorDetection_fieldSeparator": "{detected_field_separator}",
      "adp_csvMultivalueSeparatorDetection_outputFilePath": "multivalue_detection_json_file_path",
      "adp_taskActive": true,
      "adp_csvMultivalueSeparatorDetection_criteriaPerColumn": [
        {
          "Enable": true,
          "Column names": "",
          "Multivalue separators": "u+007C,;,/,,^",
          "Detected value pattern": ".*",
          "Min. value length": "0",
          "Max. value length": "128",
          "Error threshold": "3"
        }
      ],
      "adp_csvMultivalueSeparatorDetection_encoding": "{detected_encoding}",
      "adp_executionPersistent": true,
      "adp_csvMultivalueSeparatorDetection_detectedMultivalueSeparator": "detected_multivalue_separator",
      "adp_abortWfOnFailure": true,
      "adp_csvMultivalueSeparatorDetection_csvFile": "{current_csv_file}",
      "adp_cleanUpHistory": false,
      "adp_csvMultivalueSeparatorDetection_valueDelimiter": "{detected_value_delimiter}",
      "adp_loggingEnabled": true,
      "adp_csvMultivalueSeparatorDetection_outputFile": null,
      "adp_taskTimeout": 0,
      "adp_csvMultivalueSeparatorDetection_lineSeparator": "{detected_line_separator}",
      "adp_csvMultivalueSeparatorDetection_amountOfLinesToLoad": "10000"
    },
    "taskDescription": "Auto detects csv mulivalue separators",
    "taskDisplayName": "Csv file multivalue separator detection task"
  },
  {
    "taskType": "Move Files",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_moveFiles_allowBucketCreation": "false",
      "adp_taskActive": true,
      "adp_moveFiles_secretKey": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_moveFiles_moveFilesConditionsTable": [],
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_moveFiles_numberThreads": "",
      "adp_moveFiles_maxNumRetries": "",
      "adp_taskTimeout": 0,
      "adp_moveFiles_accessKey": "",
      "adp_moveFiles_cloudServiceName": "No Cloud Service",
      "adp_moveFiles_newBucketS3Location": "eu-central-1"
    },
    "taskDescription": "Move files to defined folders depending on results of workflow and tasks.",
    "taskDisplayName": "Move files task"
  },
  {
    "taskType": "Axcelerate Export",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_axcelerate_export_httpsAllowUntrustedHosts": "true",
      "adp_executionPersistent": true,
      "adp_axcelerate_export_httpsPassword": "",
      "adp_axcelerate_export_runId": "",
      "adp_axcelerate_export_httpsTrustCertificate": "",
      "adp_abortWfOnFailure": true,
      "adp_axcelerate_export_archiveLocationLinux": "",
      "adp_axcelerate_export_password": "",
      "adp_cleanUpHistory": false,
      "adp_axcelerate_export_httpsKeystoreFile": null,
      "adp_loggingEnabled": true,
      "adp_axcelerate_export_requestTimeoutSeconds": 900,
      "adp_axcelerate_export_dryRun": "false",
      "adp_axcelerate_export_archiveLocation": "",
      "adp_axcelerate_export_connectTimeoutSeconds": 300,
      "adp_taskTimeout": 0,
      "adp_axcelerate_export_projectId": "",
      "adp_axcelerate_export_compressionOn": "true",
      "adp_axcelerate_export_user": ""
    },
    "taskDescription": "Export Axclerate NG data for a particular project.",
    "taskDisplayName": "Axcelerate Export"
  },
  {
    "taskType": "Global Searches",
    "taskConfiguration": {
      "adp_globalSearches_outputFilename": "adp_globalSearches_output_file_name",
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_globalSearches_createUpdateGlobalSearches": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_globalSearches_regex": "",
      "adp_globalSearches_outputJson": "adp_globalSearches_json_output",
      "adp_taskTimeout": 0,
      "adp_globalSearches_jsonFile": "output.json",
      "adp_globalSearches_globalSearchesToDelete": "",
      "adp_globalSearches_updateCollisionResolutionMode": "error"
    },
    "taskDescription": "Manage Global Searches.",
    "taskDisplayName": "Global Searches"
  },
  {
    "taskType": "Configure Data Source",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_configuredDataSource_configuredDataSourceNameParameter": "adp_configured_data_source_name",
      "adp_taskActive": true,
      "adp_configureDataSource_configurationDifferences": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_configureDataSource_metaDataMappingToConfigTables": [],
      "adp_taskTimeout": 0,
      "adp_configureDataSource_javaHeapSize": "768",
      "adp_configureDataSource_dataSourceNames": "{adp_created_data_source_name}",
      "adp_configureDataSource_metaDataMappingToSimpleConfigParams": []
    },
    "taskDescription": "Configures a data source",
    "taskDisplayName": "Configure data source"
  },
  {
    "taskType": "Remote Process Status",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_remoteProcessStatusTask_isCustomStatus": false,
      "adp_taskActive": true,
      "adp_remoteProcessStatusTask_predefinedStatusValue": null,
      "adp_remoteProcessStatusTask_remoteFileStatus": "adp_remote_process_file_status",
      "adp_remoteProcessStatusTask_bucketId": "{adp_remote_process_file_bucket}",
      "adp_executionPersistent": true,
      "adp_remoteProcessStatusTask_cloudServiceName": "Amazon S3",
      "adp_remoteProcessStatusTask_remoteFileNameOuputParameter": "adp_remote_process_file",
      "adp_abortWfOnFailure": true,
      "adp_remoteProcessStatusTask_remoteFileName": "{adp_processed_chunk_file_name}",
      "adp_cleanUpHistory": false,
      "adp_remoteProcessStatusTask_numberThreads": null,
      "adp_remoteProcessStatusTask_secretKey": "",
      "adp_remoteProcessStatusTask_customStatusValue": "",
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_remoteProcessStatusTask_maxNumRetries": null,
      "adp_remoteProcessStatusTask_accessKey": ""
    },
    "taskDescription": "Task to upload a status information file to the cloud.",
    "taskDisplayName": "Remote process status"
  },
  {
    "taskType": "CSV Loadfile Normalizer",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvLoadfileNormalizer_csvFormatAndTargetDefinition": [
        {
          "Normalization rules": "{\"name\":\"Pattern 2 [begbates>begattach,endbates>endattach]\", \"normalizations\":[{\"type\":\"First\",\"target\":\"begattach\",\"source\":\"begbates\",\"alternativeSources\":null},{\"type\":\"Last\",\"target\":\"endattach\",\"source\":\"endbates\",\"alternativeSources\":null}]}"
        },
        {
          "Normalization rules": "{\"name\":\"Pattern 4 [begbates>begattach,attachid>endattach,begbates>endbates]\", \"normalizations\":[{\"type\":\"First\",\"target\":\"begattach\",\"source\":\"begbates\",\"alternativeSources\":null},{\"type\":\"Largest\",\"target\":\"endattach\",\"source\":\"attachid\",\"alternativeSources\":[\"endbates\",\"batesrange\",\"begbates\"]}]},{\"type\":\"PreNextDoc\",\"target\":\"endbates\",\"source\":\"begbates\",\"alternativeSources\":[\"begbates\"]}]}"
        },
        {
          "Normalization rules": "{\"name\":\"Pattern 5 [begbates>begattach,attachmentrange>endattach,batesrange>endbates]\", \"normalizations\":[{\"type\":\"First\",\"target\":\"begattach\",\"source\":\"begbates\",\"alternativeSources\":null},{\"type\":\"Largest\",\"target\":\"endattach\",\"source\":\"attachmentrange\",\"alternativeSources\":[\"endbates\",\"batesrange\",\"begbates\"]}]},{\"type\":\"Copy\",\"target\":\"endbates\",\"source\":\"batesrange\",\"alternativeSources\":null}]}"
        },
        {
          "Normalization rules": "{\"name\":\"Pattern 6 [begbates>begattach,attachmentrange>endattach,begbates>endbates]\", \"normalizations\":[{\"type\":\"First\",\"target\":\"begattach\",\"source\":\"begbates\",\"alternativeSources\":null},{\"type\":\"Largest\",\"target\":\"endattach\",\"source\":\"attachmentrange\",\"alternativeSources\":[\"endbates\",\"batesrange\",\"begbates\"]}]},{\"type\":\"PreNextDoc\",\"target\":\"endbates\",\"source\":\"begbates\",\"alternativeSources\":[\"begbates\"]}]}"
        },
        {
          "Normalization rules": "{\"name\":\"Pattern 7 [batesrange>endbates]\", \"normalizations\":[{\"type\":\"Copy\",\"target\":\"endbates\",\"source\":\"batesrange\",\"alternativeSources\":null}]}"
        },
        {
          "Normalization rules": "{\"name\":\"Pattern 1 [begbates>endbates]\", \"normalizations\":[{\"type\":\"PreNextDoc\",\"target\":\"endbates\",\"source\":\"begbates\",\"alternativeSources\":[\"endattach\",\"begattach\"]}]}"
        }
      ],
      "adp_csvLoadfileNormalizer_lineSeparator": "{detected_line_separator}",
      "adp_taskActive": true,
      "adp_csvLoadfileNormalizer_fieldSeparator": "{detected_field_separator}",
      "adp_executionPersistent": true,
      "adp_csvLoadfileNormalizer_targetCsvFile": "new.csv",
      "adp_abortWfOnFailure": true,
      "adp_csvLoadfileNormalizer_indicateNormalization": "adp_normalization",
      "adp_csvLoadfileNormalizer_fileEncoding": "{detected_encoding}",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_csvLoadfileNormalizer_sourceCsvFile": "existing.csv",
      "adp_taskTimeout": 0,
      "adp_csvLoadfileNormalizer_valueDelimiter": "{detected_value_delimiter}",
      "adp_csvLoadfileNormalizer_outputVariable": "adp_new_csv_file"
    },
    "taskDescription": "Normalizes given csv to a new csv file.",
    "taskDisplayName": "CSV Loadfile Normalizer"
  },
  {
    "taskType": "Ping Project",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_pingProject_Password": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_pingProject_User": "",
      "adp_abortWfOnFailure": true,
      "adp_pingProject_IdentifierType": "",
      "adp_pingProject_JsonOutput": "ping_project_result",
      "adp_cleanUpHistory": false,
      "adp_pingProject_Identifiers": ""
    },
    "taskDescription": "Pings applications or engines",
    "taskDisplayName": "Ping Project Task"
  },
  {
    "taskType": "APWT Clear Fields",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_clearFields_fieldsToClear": []
    },
    "taskDescription": "Clears fields as marked by the clearMaker",
    "taskDisplayName": "APWT Clear Fields"
  },
  {
    "taskType": "Read Configuration",
    "taskConfiguration": {
      "adp_readConfiguration_outputJson": "adp_entities_json_output",
      "adp_progressTaskTimeout": 0,
      "adp_readConfiguration_configsToRead": [],
      "adp_taskActive": true,
      "adp_readConfiguration_outputFilename": "adp_entities_output_file_name",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_readConfiguration_entityIdToRead": "",
      "adp_loggingEnabled": true,
      "adp_readConfiguration_file": "output.json",
      "adp_taskTimeout": 0,
      "adp_readConfiguration_fileFormat": "JSON"
    },
    "taskDescription": "A Task to read configurations into JSON or XML.",
    "taskDisplayName": "Read Configuration"
  },
  {
    "taskType": "Axcelerate Delete",
    "taskConfiguration": {
      "adp_axcelerate_delete_password": "",
      "adp_progressTaskTimeout": 0,
      "adp_axcelerate_delete_connectTimeoutSeconds": 300,
      "adp_axcelerate_delete_user": "",
      "adp_taskActive": true,
      "adp_axcelerate_delete_runId": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_axcelerate_delete_httpsKeystoreFile": null,
      "adp_axcelerate_delete_projectId": "",
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_axcelerate_delete_httpsAllowUntrustedHosts": "true",
      "adp_axcelerate_delete_httpsTrustCertificate": "",
      "adp_axcelerate_delete_requestTimeoutSeconds": 900,
      "adp_axcelerate_delete_httpsPassword": ""
    },
    "taskDescription": "Delete Axclerate NG data for a particular project.",
    "taskDisplayName": "Axcelerate Delete"
  },
  {
    "taskType": "Create Application",
    "taskConfiguration": {
      "adp_createApplication_chosenEngineHostMemory": "adp_create_application_engine_host_memory",
      "adp_progressTaskTimeout": 0,
      "adp_createApplication_abortOnExistingApplication": false,
      "adp_createApplication_applicationEngineHost": null,
      "adp_taskActive": true,
      "adp_createApplication_applicationWorkspaceId": null,
      "adp_createApplication_applicationHostMemoryLimitRatio": "0",
      "adp_createApplication_applicationDescription": null,
      "adp_createApplication_applicationHost": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_createApplication_chosenApplicationHostMemoryRatio": "adp_create_application_application_host_memory_ratio",
      "adp_createApplication_createWorkspace": "false",
      "adp_createApplication_metaEngineIdentifier": "adp_create_application_metaengine_identifier",
      "adp_loggingEnabled": true,
      "adp_createApplication_engineHostMemoryLimitRatio": "0",
      "adp_createApplication_engineIdentifier_2": "adp_create_application_engine_identifier_2",
      "adp_createApplication_applicationWorkspace": null,
      "adp_createApplication_engineIdentifier_1": "adp_create_application_engine_identifier_1",
      "adp_createApplication_applicationHostDetection": "true",
      "adp_createApplication_chosenApplicationHostMemory": "adp_create_application_application_host_memory",
      "adp_createApplication_applicationContext": null,
      "adp_createApplication_applicationTemplate": null,
      "adp_createApplication_applicationId": null,
      "adp_createApplication_namespacePrefix": null,
      "adp_createApplication_applicationDisplayname": "adp_create_application_application_displayname",
      "adp_createApplication_engineHostMemoryLimit": "0",
      "adp_createApplication_chosenEngineHostMemoryRatio": "adp_create_application_engine_host_memory_ratio",
      "adp_createApplication_applicationHostMemoryLimit": "0",
      "adp_createApplication_applicationIdentifier": "adp_create_application_application_identifier",
      "adp_createApplication_applicationName": null,
      "adp_createApplication_chosenWorkspaceNameParameter": "adp_create_application_workspace",
      "adp_createApplication_applicationTypeEnum": "ECA",
      "adp_cleanUpHistory": false,
      "adp_taskTimeout": 0,
      "adp_createApplication_chosenApplicationHostNameParameter": "adp_create_application_application_host",
      "adp_createApplication_storagePaths": [],
      "adp_createApplication_applicationType": null,
      "adp_createApplication_chosenEngineHostNameParameter": "adp_create_application_engine_host"
    },
    "taskDescription": "Creates an application",
    "taskDisplayName": "Create application"
  },
  {
    "taskType": "List Entities",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_listEntities_file": "output.json",
      "adp_listEntities_numberOfEntities": "-1",
      "adp_listEntities_axcRequestTimeoutSeconds": 900,
      "adp_taskActive": true,
      "adp_listEntities_userHasAccess": "",
      "adp_listEntities_whiteList": "id,displayName",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_listEntities_relatedEntity": "",
      "adp_listEntities_workspace": "",
      "adp_loggingEnabled": true,
      "adp_listEntities_status": "",
      "adp_listEntities_axcServiceCoreAddress": "",
      "adp_listEntities_relatedEntityType": "",
      "adp_listEntities_type": "",
      "adp_listEntities_httpsKeystoreFile": null,
      "adp_listEntities_httpsPassword": "",
      "adp_listEntities_axcConnectTimeoutSeconds": 300,
      "adp_listEntities_axcServicePassword": "",
      "adp_listEntities_startingEntity": "1",
      "adp_listEntities_outputJson": "adp_entities_json_output",
      "adp_cleanUpHistory": false,
      "adp_listEntities_descriptionSettingFilterValueDateFormat": "yyyy-MM-dd",
      "adp_listEntities_descriptionFilters": [],
      "adp_listEntities_axcServiceUser": "",
      "adp_listEntities_axcFields": "",
      "adp_taskTimeout": 0,
      "adp_listEntities_httpsTrustCertificate": "",
      "adp_listEntities_host": "",
      "adp_listEntities_outputFilename": "adp_entities_output_file_name",
      "adp_listEntities_id": "",
      "adp_listEntities_httpsAllowUntrustedHosts": "true"
    },
    "taskDescription": "Writes a list of entities ot an output variable",
    "taskDisplayName": "List entities"
  },
  {
    "taskType": "CSV Column Type Detection",
    "taskConfiguration": {
      "adp_csvctd_columnStatistics": "column_statistics",
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_csvctd_csvFile": "{validated_csv_file}",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_csvctd_columnStatisticsOutputFilePath": "column_statistics_file_path",
      "adp_loggingEnabled": true,
      "adp_csvctd_numericOnlyColumns": "numeric_only_columns",
      "adp_csvctd_dateFormatsJson": "",
      "adp_csvctd_csvValueDelimiter": "{validated_csv_file_quotecharacter}",
      "adp_csvctd_csvFileEncoding": "{validated_csv_file_encoding}",
      "adp_csvctd_dateColumns": "date_columns",
      "adp_csvctd_csvFieldSeparator": "{validated_csv_file_fieldseparator}",
      "adp_csvctd_csvLineSeparator": "{validated_csv_file_lineseparator}",
      "adp_csvctd_extendedDataSourceConfigurationDiff": "data_source_configuration_diff",
      "adp_csvctd_textOnlyColumns": "text_only_columns",
      "adp_csvctd_includeColumns": "{validated_header_columns}",
      "adp_csvctd_dataSourceIdentifier": null,
      "adp_csvctd_dateFormatsFile": null,
      "adp_csvctd_csvMergeConfiguration": null,
      "adp_cleanUpHistory": false,
      "adp_csvctd_columnStatisticsOutputFile": null,
      "adp_csvctd_extendedCsvMergeConfiguration": "csv_merge_configuration",
      "adp_csvctd_referencedTextColumns": "referenced_text_columns",
      "adp_csvctd_maxPercentageDifferentValues": 80,
      "adp_csvctd_maxNumberDifferentValues": 50000,
      "adp_csvctd_optFile": null,
      "adp_taskTimeout": 0,
      "adp_csvctd_atMostMaxNumberDifferentValuesColumns": "at_most_max_number_different_values_columns",
      "adp_csvctd_dataSourceConfigurationDiff": null,
      "adp_csvctd_excludeColumns": "{validated_unmapped_empty_columns}",
      "adp_csvctd_inlineTextColumns": "inline_text_columns",
      "adp_csvctd_textFromOpt": "text_from_opt"
    },
    "taskDescription": "Detects column types of a specified csv-file.",
    "taskDisplayName": "CSV column type detection task"
  },
  {
    "taskType": "Remove Processes",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_removeProcess_removeAssociatedStorages": "false",
      "adp_removeProcess_processIdentifiers": [],
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_removeProcess_storagesRemoved": "adp_storages_included_in_removal",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Removes processes",
    "taskDisplayName": "Remove processes"
  },
  {
    "taskType": "Matter Report",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_matterReport_matterSpecificUrlRegEx": null,
      "adp_matterReport_file": "output.json",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_matterReport_webserviceUser": "{service_user}",
      "adp_loggingEnabled": true,
      "adp_matterReport_processedMatterId": "adp_processed_matter_id",
      "adp_matterReport_httpsKeystoreFile": null,
      "adp_matterReport_configuration": "",
      "adp_matterReport_httpsAllowUntrustedHosts": "false",
      "adp_matterReport_usedWebserviceUrl": "adp_url_used_for_matter_reporting",
      "adp_matterReport_matterId": null,
      "adp_matterReport_httpsPassword": "",
      "adp_matterReport_webserviceRequestTimeoutSeconds": 900,
      "adp_matterReport_matterDisplayName": null,
      "adp_matterReport_matterSpecificApplication": null,
      "adp_cleanUpHistory": false,
      "adp_matterReport_webserviceConnectTimeoutSeconds": 300,
      "adp_matterReport_outputJson": "adp_matterReport_json_output",
      "adp_matterReport_reportFilePath": "adp_matterReport_output_file_name",
      "adp_matterReport_httpsTrustCertificate": "",
      "adp_matterReport_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_taskTimeout": 0,
      "adp_matterReport_webservicePassword": null
    },
    "taskDescription": "Task to export matter reports.",
    "taskDisplayName": "Matter Report"
  },
  {
    "taskType": "Manage Host Roles",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_manageHostRoles_outputFilename": "adp_manageHostRoles_output_file_name",
      "adp_manageHostRoles_hostRoles": [],
      "adp_taskActive": true,
      "adp_manageHostRoles_filterForAutomatedCreation": "false",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_manageHostRoles_extJson": "false",
      "adp_manageHostRoles_file": "output.json",
      "adp_manageHostRoles_outputJson": "adp_manageHostRoles_json_output",
      "adp_taskTimeout": 0,
      "adp_manageHostRoles_hostIdsToFilterFor": "",
      "adp_manageHostRoles_inputJson": ""
    },
    "taskDescription": "Mangage roles for a host.",
    "taskDisplayName": "Manage host roles"
  },
  {
    "taskType": "Stop Processes",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_stopProcess_processIdentifiers": [],
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Stops processes",
    "taskDisplayName": "Stop processes"
  },
  {
    "taskType": "Delete Storages",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_deleteStorages_engineUser": "{adp_user}",
      "adp_deleteStorages_enginePassword": "",
      "adp_deleteStorages_waitForDeletion": "false",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_deleteStorages_deletionJobIds": "adp_delete_storages_job_ids",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_deleteStorages_engineId": ""
    },
    "taskDescription": "Deletes engine storages.",
    "taskDisplayName": "Delete Storages"
  },
  {
    "taskType": "PSS Trigger DB Back-up",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "pss_TriggerDBBackupRestPassword": null,
      "pss_TriggerDBBackupDbBackupBaseDirectory": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "pss_TriggerDBBackupAuthURL": null,
      "adp_abortWfOnFailure": true,
      "pss_TriggerDBBackupPssURL": null,
      "adp_cleanUpHistory": false,
      "pss_TriggerDBBackupRestUserName": null
    },
    "taskDescription": "Triggers PSS DB back-up.",
    "taskDisplayName": "PSS Trigger DB Back-up"
  },
  {
    "taskType": "Configure Hosts in Workspaces",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_configureHostsInWorkspaces_addRemoveHostsInWorkspaces": "",
      "adp_configureHostsInWorkspaces_outputFilename": "adp_configureHostsInWorkspaces_output_file_name",
      "adp_configureHostsInWorkspaces_jsonFile": "output.json",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_configureHostsInWorkspaces_outputJson": "adp_configureHostsInWorkspaces_json_output",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_configureHostsInWorkspaces_cfgAutomaticallyAddNewHostsToWorkspace": "",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Configure 'add hosts automatically to workspace' and add/remove hosts to/from workspaces.",
    "taskDisplayName": "Configure Hosts in Workspaces"
  },
  {
    "taskType": "Manage Taggers",
    "taskConfiguration": {
      "adp_mantags_jsonInstall": "[]",
      "adp_progressTaskTimeout": 0,
      "adp_mantags_engineUser": null,
      "adp_mantags_jsonFile": null,
      "adp_mantags_jsonUninstall": "",
      "adp_mantags_outputJson": "",
      "adp_taskActive": true,
      "adp_mantags_applicationIdentifier": "",
      "adp_mantags_applicationType": "",
      "adp_executionPersistent": true,
      "adp_mantags_engineIdentifier": "",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_mantags_enginePassword": null,
      "adp_mantags_engineType": "true",
      "adp_taskTimeout": 0,
      "adp_mantags_outputFilename": "adp_mantags_output_file_name",
      "adp_mantags_allowOverwrite": "true",
      "adp_mantags_wait4Completion": "false"
    },
    "taskDescription": "Install and uninstall taggers.",
    "taskDisplayName": "Manage Taggers"
  },
  {
    "taskType": "Export",
    "taskConfiguration": {
      "adp_exportTask_savedSearchesFile": "",
      "adp_progressTaskTimeout": 0,
      "adp_exportTask_mode": "COMPLETE",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_exportTask_skipStorageLocationExport": "false",
      "adp_loggingEnabled": true,
      "adp_exportTask_reenableProjects": "false",
      "adp_taskTimeout": 0,
      "adp_exportTask_dryRun": "false",
      "adp_exportTask_applicationsJson": "",
      "adp_exportTask_compressionMethod": "",
      "adp_exportTask_storageExportMode": "LOCATION",
      "adp_exportTask_disableSanityChecks": "true"
    },
    "taskDescription": "Interface to exportApplicationTree",
    "taskDisplayName": "Export"
  },
  {
    "taskType": "Suspend Data Source",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_suspendDataSource_wait": "true",
      "adp_loggingEnabled": true,
      "adp_suspendDataSource_dataSources": "",
      "adp_suspendDataSource_allDataSources": "ds_all",
      "adp_suspendDataSource_mode": "suspend",
      "adp_taskTimeout": 0,
      "adp_suspendDataSource_dataSourcesInUnknownState": "ds_in_unknown_state",
      "adp_suspendDataSource_hosts": "",
      "adp_suspendDataSource_dataSourcesWithError": "ds_with_suspend_resume_error"
    },
    "taskDescription": "Suspends and resumes data sources",
    "taskDisplayName": "Suspend data source"
  },
  {
    "taskType": "APWT Compute Chains",
    "taskConfiguration": {
      "adp_chainFilePrefix": null,
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_computeChains_tempDirectory": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Computes chain files (concatenation of agreements and amendments)",
    "taskDisplayName": "APWT Compute Chains Task"
  },
  {
    "taskType": "CSV Detection",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvDetection_detectedFieldSeparator": "detected_field_separator",
      "adp_csvDetection_readBufferSize": 8192,
      "adp_taskActive": true,
      "adp_csvDetection_setEncoding": false,
      "adp_csvDetection_valueDelimiter": [
        {
          "Unicode character": "u+0022"
        },
        {
          "Unicode character": "u+00FE"
        },
        {
          "Unicode character": "u+005E"
        }
      ],
      "adp_csvDetection_fieldSeparator": [
        {
          "Unicode character": "u+002C"
        },
        {
          "Unicode character": "u+0014"
        },
        {
          "Unicode character": "u+007C"
        },
        {
          "Unicode character": "u+003B"
        },
        {
          "Unicode character": "u+0009"
        }
      ],
      "adp_executionPersistent": true,
      "adp_csvDetection_csvFile": "new.csv",
      "adp_csvDetection_encoding": "UTF-8",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_csvDetection_maxCharactersLoad": 100000,
      "adp_loggingEnabled": true,
      "adp_csvDetection_amountOfLinesToLoad": "3",
      "adp_csvDetection_detectedValueDelimiter": "detected_value_delimiter",
      "adp_taskTimeout": 0,
      "adp_csvDetection_detectedLineSeparator": "detected_line_separator",
      "adp_csvDetection_detectedEncoding": "detected_encoding",
      "adp_csvDetection_maxLineLength": -1
    },
    "taskDescription": "Auto detects csv file criterias",
    "taskDisplayName": "Csv file detection task"
  },
  {
    "taskType": "Change Field Contents",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_changeFieldContents_searchParents": false,
      "adp_changeFieldContents_applicationName": null,
      "adp_changeFieldContents_documentSelectionMode": "Select by document ID",
      "adp_changeFieldContents_fieldChanges": [],
      "adp_changeFieldContents_engineName": null,
      "adp_changeFieldContents_enginePassword": null,
      "adp_changeFieldContents_NumberOfChangedDocuments": "adp_changeFieldContents_output_numberOfDocuments",
      "adp_taskActive": true,
      "adp_changeFieldContents_includeFamily": false,
      "adp_executionPersistent": true,
      "adp_changeFieldContents_engineType": "true",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_changeFieldContents_docID": "",
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_changeFieldContents_engineUser": null,
      "adp_changeFieldContents_query": "",
      "adp_changeFieldContents_adp_changeFieldContents_mainQueryType": null
    },
    "taskDescription": "Change meta data or text type contents in a configured engine",
    "taskDisplayName": "Change field contents task"
  },
  {
    "taskType": "Axcelerate Remove Bootstrap Information",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_axcelerate_removeBootstrapInformation_requestTimeoutSeconds": 900,
      "adp_axcelerate_removeBootstrapInformation_password": "",
      "adp_taskActive": true,
      "adp_axcelerate_removeBootstrapInformation_httpsAllowUntrustedHosts": "true",
      "adp_axcelerate_removeBootstrapInformation_connectTimeoutSeconds": 300,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_axcelerate_removeBootstrapInformation_projectId": "",
      "adp_loggingEnabled": true,
      "adp_axcelerate_removeBootstrapInformation_httpsKeystoreFile": null,
      "adp_taskTimeout": 0,
      "adp_axcelerate_removeBootstrapInformation_httpsPassword": "",
      "adp_axcelerate_removeBootstrapInformation_httpsTrustCertificate": "",
      "adp_axcelerate_removeBootstrapInformation_user": ""
    },
    "taskDescription": "Sets bootstrap information for Axclerate NG project.",
    "taskDisplayName": "Axcelerate Remove Bootstrap Information"
  },
  {
    "taskType": "APWT ZIP",
    "taskConfiguration": {
      "adp_zip_resultFile": "adp_result_zip_file",
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_zip_numberOfFilesInZip": "adp_zip_number_files_in_zip",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_zip_inputFolder": null,
      "adp_zip_batchId": null,
      "adp_zip_filePostfix": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Compress PDF files into a ZIP archive",
    "taskDisplayName": "APWT ZIP Task"
  },
  {
    "taskType": "Axcelerate Set Bootstrap Information",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_axcelerate_setBootstrapInformation_archiveLocation": "",
      "adp_axcelerate_setBootstrapInformation_projectId": "",
      "adp_axcelerate_setBootstrapInformation_httpsKeystoreFile": null,
      "adp_axcelerate_setBootstrapInformation_requestTimeoutSeconds": 900,
      "adp_taskActive": true,
      "adp_axcelerate_setBootstrapInformation_httpsTrustCertificate": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_axcelerate_setBootstrapInformation_httpsPassword": "",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_axcelerate_setBootstrapInformation_user": "",
      "adp_axcelerate_setBootstrapInformation_httpsAllowUntrustedHosts": "true",
      "adp_taskTimeout": 0,
      "adp_axcelerate_setBootstrapInformation_connectTimeoutSeconds": 300,
      "adp_axcelerate_setBootstrapInformation_password": ""
    },
    "taskDescription": "Sets bootstrap information for Axclerate NG project.",
    "taskDisplayName": "Axcelerate Set Bootstrap Information"
  },
  {
    "taskType": "Create OCR Job",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_createOcrJob_engineUserPassword": "",
      "adp_createOcrJob_query": "*",
      "adp_taskActive": true,
      "adp_createOcrJob_listOfJobProperties": "",
      "adp_executionPersistent": true,
      "adp_createOcrJob_engineType": "true",
      "adp_abortWfOnFailure": true,
      "adp_loggingEnabled": true,
      "adp_createOcrJob_AdvancedRestrictions": [],
      "adp_createOcrJob_globalSearchJson": "",
      "adp_createOcrJob_wait": "false",
      "adp_createOcrJob_engineName": "{adp_engineName}",
      "adp_createOcrJob_jobDescription": "",
      "adp_createOcrJob_applicationIdentifier": "",
      "adp_createOcrJob_jobPriority": "10",
      "adp_createOcrJob_jobName": "",
      "adp_createOcrJob_restrictions": [],
      "adp_cleanUpHistory": false,
      "adp_createOcrJob_engineUserName": "{adp_user}",
      "adp_createOcrJob_mainQueryType": null,
      "adp_createOcrJob_applicationType": "",
      "adp_createOcrJob_globalSearchId": "",
      "adp_taskTimeout": 0,
      "adp_createOcrJob_jsonOutputVariable": "adp_createOcrJob_json_output"
    },
    "taskDescription": "Changes metaData by using regEx replacement.",
    "taskDisplayName": "Create OCR Job"
  },
  {
    "taskType": "Save Engines",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_saveEngine_savesPerHost": -2147483648,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_saveEngine_engineIdentifiers": [],
      "adp_saveEngine_appIdentifiers": [],
      "adp_cleanUpHistory": false,
      "adp_saveEngine_unsavedEngines": "enginesNotSaved",
      "adp_loggingEnabled": true,
      "adp_saveEngine_engineUser": null,
      "adp_taskTimeout": 0,
      "adp_saveEngine_unstoppedEngines": "enginesNotStopped",
      "adp_saveEngine_enginePassword": null
    },
    "taskDescription": "Saves engines",
    "taskDisplayName": "Save engines"
  },
  {
    "taskType": "Create Custodian",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_createCustodian_engineUser": "{adp_user}",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_createCustodian_applicationType": "",
      "adp_createCustodian_engineIdentifier": "",
      "adp_loggingEnabled": true,
      "adp_createCustodian_custodians": [],
      "adp_createCustodian_custodianTaxonomy": "rm_custodian",
      "adp_taskTimeout": 0,
      "adp_createCustodian_applicationIdentifier": "",
      "adp_createCustodian_engineType": "true",
      "adp_createCustodian_engineUserPassword": "",
      "adp_createCustodian_updateExistingCustodians": true
    },
    "taskDescription": "Creates a custodian from the given input parameters.",
    "taskDisplayName": "Create Custodian"
  },
  {
    "taskType": "Create Investigation",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_cinv_webserviceConnectTimeoutSeconds": 300,
      "adp_cinv_appHostMemoryLimit": "0",
      "adp_cinv_matterSpecificTemplate": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_cinv_matterId": "default",
      "adp_cinv_matterSpecificApplicationHostId": "",
      "adp_cinv_predictiveCodingEnabled": false,
      "adp_cinv_chosenEngineHostMemoryRatio": "adp_create_investigation_engine_host_memory_ratio",
      "adp_abortWfOnFailure": true,
      "adp_cinv_publishApplicationId": "adp_created_publish_application_id",
      "adp_cinv_webservicePassword": null,
      "adp_cinv_matterSpecificSmallProject": false,
      "adp_loggingEnabled": true,
      "adp_cinv_matterSpecificWorkspace": null,
      "adp_cinv_matterSpecificApplication": null,
      "adp_cinv_matterSpecificEngineHostId": "localhost",
      "adp_cinv_publishEngineId": "adp_created_publish_engine_id",
      "adp_cinv_webserviceUser": "{service_user}",
      "adp_cinv_httpsAllowUntrustedHosts": "false",
      "adp_cinv_httpsPassword": "",
      "adp_cinv_chosenEngineHostMemory": "adp_create_investigation_engine_host_memory",
      "adp_cinv_httpsTrustCertificate": "",
      "adp_cinv_webserviceRequestTimeoutSeconds": 900,
      "adp_cinv_chosenApplicationHostMemory": "adp_create_investigation_application_host_memory",
      "adp_cinv_chosenEngineHostNameParameter": "adp_create_investigation_engine_host",
      "adp_cinv_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_cleanUpHistory": false,
      "adp_cinv_expectedNumberOfDocuments": 1000000,
      "adp_cinv_reduceMemoryFootPrint": false,
      "adp_cinv_appHostDetection": "true",
      "adp_cinv_engineHostMemoryLimit": "0",
      "adp_cinv_appHostMemoryLimitRatio": "0",
      "adp_cinv_engineHostMemoryLimitRatio": "0",
      "adp_cinv_matterSpecificUrlRegEx": null,
      "adp_taskTimeout": 0,
      "adp_cinv_httpsKeystoreFile": null,
      "adp_cinv_chosenApplicationHostNameParameter": "adp_create_investigation_application_host",
      "adp_cinv_usedWebserviceUrl": "adp_url_used_for_create_investigation",
      "adp_cinv_chosenApplicationHostMemoryRatio": "adp_create_investigation_application_host_memory_ratio"
    },
    "taskDescription": "This task creates Investigation applications using MatterSpecific web-service.",
    "taskDisplayName": "Task to create Investigation application"
  },
  {
    "taskType": "Export Storage",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_exportStorage_mode": "LOCATION",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_exportStorage_engineId": "",
      "adp_executionPersistent": true,
      "adp_exportStorage_targetPathLinux": null,
      "adp_abortWfOnFailure": true,
      "adp_exportStorage_targetPathWindows": null,
      "adp_exportStorage_applicationId": "",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Exports the storage of a project",
    "taskDisplayName": "Export Storage"
  },
  {
    "taskType": "Configure Service Properties",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_configureServiceProperties_ignoreSuspended": "false",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_configureServiceProperties_actionsOnEntity": []
    },
    "taskDescription": "Configures service properties",
    "taskDisplayName": "Configure service properties"
  },
  {
    "taskType": "Configure Entity",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_configureEntity_entityIdentifier": "{adp_create_application_application_identifier}",
      "adp_taskActive": true,
      "adp_configureEntity_metaDataMappingToDynamicComponentsSimpleConfigParams": [],
      "adp_configureEntity_metaDataMappingToDynamicComponentsConfigStringLists": [],
      "adp_configureEntity_metaDataMappingToSimpleConfigParams": [],
      "adp_executionPersistent": true,
      "adp_configureEntity_metaDataMappingToConfigStringLists": [],
      "adp_configureEntity_configuredEntityIdentifier": "adp_configured_entity_identifier",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_configureEntity_configurationDifferences": "",
      "adp_configureEntity_javaHeapSize": "768",
      "adp_configureEntity_metaDataMappingToDynamicComponentsConfigTables": [],
      "adp_loggingEnabled": true,
      "adp_configureEntity_metaDataMappingToConfigTables": [],
      "adp_configureEntity_metaDataMappingToDynamicComponents": [],
      "adp_taskTimeout": 0
    },
    "taskDescription": "Configures an entity",
    "taskDisplayName": "Configure entity"
  },
  {
    "taskType": "APWT Self Service Processor",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_apwt_pss_view_identifier": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_apwt_pss_per_model_identifier": null
    },
    "taskDescription": "Perceptiv Self Service specifiv data manipulations",
    "taskDisplayName": "APWT Perceptiv Self Service Processor Task"
  },
  {
    "taskType": "Receive meta data from file",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_receiveMetaDataFromFile_pollingInterval": 1,
      "adp_receiveMetaDataFromFile_triggerFileExists": true,
      "adp_receiveMetaDataFromFile_sourceFile": "metadata.properties",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_receiveMetaDataFromFile_timeFileNotModified": 5,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Creates or overrides new meta data key/value pairs from a properties file for the workflow",
    "taskDisplayName": "Receive meta data from file"
  },
  {
    "taskType": "Start Application",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_startApplication_useHttps": false,
      "adp_startApplication_applicationUrl": "adp_started_application_url",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_startApplication_applicationIdentifier": "{adp_create_application_application_identifier}"
    },
    "taskDescription": "Starts an application",
    "taskDisplayName": "Start application"
  },
  {
    "taskType": "CSV Column Filter",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvColumnFilter_notWrittenColumns": "column_filtered_not_written",
      "adp_csvColumnFilter_originalCsvFile": "original.csv",
      "adp_taskActive": true,
      "adp_csvColumnFilter_encoding": "{detected_encoding}",
      "adp_executionPersistent": true,
      "adp_csvColumnFilter_valueDelimiter": "{detected_value_delimiter}",
      "adp_abortWfOnFailure": true,
      "adp_csvColumnFilter_writtenColumns": "column_filtered_written",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_csvColumnFilter_newCsvFile": "new.csv",
      "adp_csvColumnFilter_fieldSeparator": "{detected_field_separator}",
      "adp_taskTimeout": 0,
      "adp_csvColumnFilter_newCsvFileVariable": "column_filtered_csv_file",
      "adp_csvColumnFilter_lineSeparator": "{detected_line_separator}",
      "adp_csvColumnFilter_whiteList": "",
      "adp_csvColumnFilter_blackList": ""
    },
    "taskDescription": "Creates a new csv file with a column subset of the original csv file.",
    "taskDisplayName": "CSV Column Filter"
  },
  {
    "taskType": "Purge History",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_ph_historyDescriptionJson": null,
      "adp_ph_purgeHistoryStatus": "purge_history_status",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Purges ADP task data base history.",
    "taskDisplayName": "Purge ADP Task History"
  },
  {
    "taskType": "Manage Core Host Roles",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_manageCoreHostRoles_file": "output.json",
      "adp_manageCoreHostRoles_outputJson": "adp_manageCoreHostRoles_json_output",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_manageCoreHostRoles_filterForAutomatedCreation": "false",
      "adp_cleanUpHistory": false,
      "adp_manageCoreHostRoles_inputJson": "",
      "adp_loggingEnabled": true,
      "adp_manageCoreHostRoles_hostIdsToFilterFor": "",
      "adp_taskTimeout": 0,
      "adp_manageCoreHostRoles_outputFilename": "adp_manageCoreHostRoles_output_file_name",
      "adp_manageCoreHostRoles_extJson": "true"
    },
    "taskDescription": "Mangage core roles for a host.",
    "taskDisplayName": "Manage core host roles"
  },
  {
    "taskType": "Matter Management",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_mm_webservicePassword": null,
      "adp_mm_matterSpecificUrlRegEx": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_mm_moveNativesToMatterStorage": "false",
      "adp_mm_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_mm_usedWebserviceUrl": "adp_url_used_for_matter_management",
      "adp_loggingEnabled": true,
      "adp_mm_savedSearchId": null,
      "adp_mm_webserviceRequestTimeoutSeconds": 900,
      "adp_mm_matterId": null,
      "adp_mm_searchString": "*",
      "adp_mm_retryCopyNativesOnError": "false",
      "adp_mm_matterDisplayName": null,
      "adp_mm_processedSavedSearchId": "adp_processed_saved_search_id",
      "adp_mm_webserviceUser": "{service_user}",
      "adp_mm_httpsAllowUntrustedHosts": "false",
      "adp_mm_httpsPassword": "",
      "adp_cleanUpHistory": false,
      "adp_mm_httpsKeystoreFile": null,
      "adp_mm_processedMatterId": "adp_processed_matter_id",
      "adp_mm_webserviceConnectTimeoutSeconds": 300,
      "adp_mm_httpsTrustCertificate": "",
      "adp_mm_enableSearch": "true",
      "adp_taskTimeout": 0,
      "adp_mm_copyNotCopiedNatives": "false",
      "adp_mm_searchFields": [],
      "adp_mm_matterSpecificApplication": null,
      "adp_mm_savedSearchDisplayName": null
    },
    "taskDescription": "",
    "taskDisplayName": "Task to manage matters and saved searches"
  },
  {
    "taskType": "Enable Bucket",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_enableBucket_bucketEnabled": "adp_bucket_enabled",
      "adp_enableBucket_numberThreads": "",
      "adp_taskActive": true,
      "adp_enableBucket_bucketId": "",
      "adp_executionPersistent": true,
      "adp_enableBucket_enableBucket": "true",
      "adp_enableBucket_maxNumRetries": "",
      "adp_abortWfOnFailure": true,
      "adp_enableBucket_accessKey": "",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_enableBucket_cloudServiceName": "Amazon S3",
      "adp_enableBucket_secretKey": ""
    },
    "taskDescription": "Enables or disables the specified bucket.",
    "taskDisplayName": "Enable bucket task"
  },
  {
    "taskType": "Export Tagger Global Searches",
    "taskConfiguration": {
      "adp_extags_enginePassword": null,
      "adp_progressTaskTimeout": 0,
      "adp_extags_engineUser": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_extags_outputJson": "",
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_extags_jsonFile": null,
      "adp_extags_applicationType": "",
      "adp_extags_applicationIdentifier": "",
      "adp_extags_outputFilename": "adp_extags_output_file_name"
    },
    "taskDescription": "Exports global searches associated with taggers of a project.",
    "taskDisplayName": "Export Tagger Global Searches"
  },
  {
    "taskType": "Export Storage Locations",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_exportStorageLocations_engineId": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_exportStorageLocations_applicationId": ""
    },
    "taskDescription": "Interface to exportStorageLocations",
    "taskDisplayName": "Export Storage Locations"
  },
  {
    "taskType": "List Entity Ids",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_listEntityIds_outputFilename": "adp_entity_ids_output_file_name",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_listEntityIds_outputOptions": "",
      "adp_listEntityIds_outputJson": "adp_entity_ids_json_output",
      "adp_listEntityIds_includeRelatedTypes": "",
      "adp_taskTimeout": 0,
      "adp_listEntityIds_file": null,
      "adp_listEntityIds_entities": "",
      "adp_listEntityIds_types": ""
    },
    "taskDescription": "Writes a list of entity ids and optionally their related entity ids into an output variable",
    "taskDisplayName": "List entity ids"
  },
  {
    "taskType": "APWT Workflow Synchronization",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_workflowSynchronization_engine_id": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_workflowSynchronization_checkDelay": 30,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Synchronizes APWT workflows based on agreement id and target engine",
    "taskDisplayName": "APWT Workflow Synchronization Task"
  },
  {
    "taskType": "Export Documents",
    "taskConfiguration": {
      "adp_exportDocuments_field_separator": ";",
      "adp_progressTaskTimeout": 0,
      "adp_exportDocuments_waitForExport": false,
      "adp_exportDocuments_image_field": null,
      "adp_exportDocuments_searchResultSize": "adp_exportDocuments_searchResultSize",
      "adp_taskActive": true,
      "adp_exportDocuments_File_Ending": "csv",
      "adp_exportDocuments_applicationType": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_exportDocuments_query": "*",
      "adp_loggingEnabled": true,
      "adp_exportDocuments_exportName": null,
      "adp_exportDocuments_text_indicator": "\"",
      "adp_exportDocuments_natives_field": null,
      "adp_exportDocuments_multivalue_separator": "|",
      "adp_exportDocuments_line_break": "",
      "adp_exportDocuments_applicationIdentifier": "",
      "adp_exportDocuments_engineIdentifier": null,
      "adp_exportDocuments_exportFileName": "adp_exportDocuments_exportFileName",
      "adp_exportDocuments_engineUser": null,
      "adp_exportDocuments_image_volume": "Volume",
      "adp_exportDocuments_exportFields": null,
      "adp_exportDocuments_fullExportPath": "adp_exportDocuments_exportPath",
      "adp_exportDocuments_text_field": null,
      "adp_cleanUpHistory": false,
      "adp_exportDocuments_exportDirectory": null,
      "adp_taskTimeout": 0,
      "adp_exportDocuments_enginePassword": null,
      "adp_exportDocuments_adp_exportDocuments_mainQueryType": null
    },
    "taskDescription": "Export documents in CSV format.",
    "taskDisplayName": "Export documents task"
  },
  {
    "taskType": "Split Csv Data Source",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_splitCsvDataSource_chunkSize": 100000,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_splitCsvDataSource_dataSourceName": "{adp_configured_data_source_name}",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_splitCsvDataSource_splitDataSources": "adp_split_data_sources"
    },
    "taskDescription": "Splits a CSV data source into chunks producing multiple data sources if necessary",
    "taskDisplayName": "Split CSV data source"
  },
  {
    "taskType": "PSS Validate Model",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "pss_validateModel_inputModelDataJson": "",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Validates the Model for Consistency.",
    "taskDisplayName": "PSS Validate Model"
  },
  {
    "taskType": "Clear Artifacts",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_ca_artifactsDescriptionJson": null,
      "adp_ca_artifactsRemovalResultJson": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_ca_artifactsRemovalStatus": "removal_status",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Clears ADP task artifacts like log-files, meta-data files, validation reports etc.",
    "taskDisplayName": "Clear ADP Task Artifacts"
  },
  {
    "taskType": "Axcelerate Application Lock",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_axcelerate_applicationLock_httpsTrustCertificate": "",
      "adp_axcelerate_applicationLock_mode": "",
      "adp_axcelerate_applicationLock_password": "",
      "adp_taskActive": true,
      "adp_axcelerate_applicationLock_projectId": "",
      "adp_executionPersistent": true,
      "adp_axcelerate_applicationLock_requestTimeoutSeconds": 900,
      "adp_abortWfOnFailure": true,
      "adp_axcelerate_applicationLock_user": "",
      "adp_cleanUpHistory": false,
      "adp_axcelerate_applicationLock_connectTimeoutSeconds": 300,
      "adp_loggingEnabled": true,
      "adp_axcelerate_applicationLock_httpsPassword": "",
      "adp_taskTimeout": 0,
      "adp_axcelerate_applicationLock_httpsAllowUntrustedHosts": "true",
      "adp_axcelerate_applicationLock_httpsKeystoreFile": null
    },
    "taskDescription": "(Un)Lock an Axclerate NG project.",
    "taskDisplayName": "Axcelerate Application Lock"
  },
  {
    "taskType": "CSV Column Alias",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvColumnAlias_valueDelimiter": "{detected_value_delimiter}",
      "adp_csvColumnAlias_aliasFile": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_csvColumnAlias_aliasJson": "",
      "adp_csvColumnAlias_lineSeparator": "{detected_line_separator}",
      "adp_csvColumnAlias_originalColumns": "column_aliased_original",
      "adp_cleanUpHistory": false,
      "adp_csvColumnAlias_newCsvFile": "new.csv",
      "adp_loggingEnabled": true,
      "adp_csvColumnAlias_encoding": "{detected_encoding}",
      "adp_csvColumnAlias_originalCsvFile": "original.csv",
      "adp_csvColumnAlias_newCsvFileVariable": "column_aliased_csv_file",
      "adp_csvColumnAlias_resultingColumns": "column_aliased_result",
      "adp_csvColumnAlias_fieldSeparator": "{detected_field_separator}",
      "adp_taskTimeout": 0,
      "adp_csvColumnAlias_aliasFileEncoding": "UTF-8",
      "adp_csvColumnAlias_autoAliasing": false
    },
    "taskDescription": "Transforms CSV columns based on an alias file.",
    "taskDisplayName": "CSV Column Alias"
  },
  {
    "taskType": "APWT Image to PDF A Converter",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_pitpac_baseFolder": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_pitpac_poemPort": -2147483648,
      "adp_abortWfOnFailure": true,
      "adp_pitpac_poemServer": null,
      "adp_cleanUpHistory": false,
      "adp_pitpac_tempDirectory": null
    },
    "taskDescription": "Convert images to PDF-A files",
    "taskDisplayName": "APWT Image to PDF A Converter Task"
  },
  {
    "taskType": "Query Postgresql DB",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_qpgdb_dbTrustCerts": "",
      "adp_qpgdb_outputJson": "",
      "adp_qpgdb_dbUser": "",
      "adp_qpgdb_jsonFile": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_qpgdb_dbPassword": "",
      "adp_abortWfOnFailure": true,
      "adp_qpgdb_sqlQuery": "",
      "adp_qpgdb_dbConnectionUrl": "",
      "adp_cleanUpHistory": false,
      "adp_qpgdb_dbConnectionPoolClientKeyPath": "ProcessControl/security/postgres/client.pgbouncer.storage.pkcs.key",
      "adp_loggingEnabled": true,
      "adp_qpgdb_outputFilename": "adp_qpgdb_output_file_name",
      "adp_qpgdb_dbConnectionPoolRootCertPath": "ProcessControl/security/postgres/reco_root.pem.crt",
      "adp_taskTimeout": 0,
      "adp_qpgdb_dbConnectionPoolClientCertPath": "ProcessControl/security/postgres/client.pgbouncer.storage.crt",
      "adp_qpgdb_dbConnectionPoolName": "",
      "adp_qpgdb_jsonResultSizeLimitMB": "2"
    },
    "taskDescription": "Queries postgresql databases.",
    "taskDisplayName": "Query Postgresql DB"
  },
  {
    "taskType": "Start Elastic Crawler Hosts",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_startElasticHost_hosts": ""
    },
    "taskDescription": "Starts elastic crawler hosts",
    "taskDisplayName": "Start elastic crawler hosts"
  },
  {
    "taskType": "APWT Generate Csv",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_generateCsv_csvFilesDirPath": null,
      "adp_generateCsv_CsvSubFundsFile": "adp_generateCsv_CsvSubFundsFile",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_generateCsv_CsvFile": "adp_generateCsv_CsvFile"
    },
    "taskDescription": "Generates CSV files (normal agreements, umbrella funds) from the given agreements",
    "taskDisplayName": "APWT Generate Csv Task"
  },
  {
    "taskType": "APWT Parse and Validate",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_descriptorAgreementDateIdentifier": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_descriptorUmbrellaCounterPartyIdsIdentifier": null,
      "adp_loggingEnabled": true,
      "adp_batchZipFile": null,
      "adp_csvDescriptorFundMetaDataColumnPrefix": null,
      "adp_descriptorCounterPartyNameIdentifier": null,
      "adp_loadFileFormatName": null,
      "adp_descriptorAgreementIdIdentifier": null,
      "adp_descriptorRelativePathIdentifier": null,
      "adp_csvColumnClearMarker": null,
      "adp_descriptorUmbrellaCounterPartyNamesIdentifier": null,
      "adp_descriptorAgreementTypeIdentifier": null,
      "adp_descriptorAddedUmbrellaCounterPartyIdsIdentifier": null,
      "adp_parseAndValidate_tempDirectory": null,
      "adp_cleanUpHistory": false,
      "adp_descriptorChangedUmbrellaCounterPartyIdsIdentifier": null,
      "adp_descriptorNativeFilePathIdentifier": null,
      "adp_batchId": null,
      "adp_taskTimeout": 0,
      "adp_descriptorChangedUmbrellaCounterPartyNamesIdentifier": null,
      "adp_descriptorAddedUmbrellaCounterPartyNamesIdentifier": null,
      "adp_descriptorRemovedUmbrellaCounterPartyIdsIdentifier": null,
      "adp_descriptorAmendmentIdIdentifier": null
    },
    "taskDescription": "Runs a native task in its onw process",
    "taskDisplayName": "APWT Parse and Validate Task"
  },
  {
    "taskType": "Receive meta data from database",
    "taskConfiguration": {
      "adp_receiveMetaDataFromDatabase_triggerDatabaseValue": true,
      "adp_progressTaskTimeout": 0,
      "adp_receiveMetaDataFromDatabase_sqlTriggerCommand": "select {db_field} from {trigger_table} where {trigger_condition};",
      "adp_receiveMetaDataFromDatabase_writeHeader": true,
      "adp_taskActive": true,
      "adp_receiveMetaDataFromDatabase_databaseUrl": "default",
      "adp_executionPersistent": true,
      "adp_receiveMetaDataFromDatabase_customLineSeparator": "U+000DU+000A",
      "adp_receiveMetaDataFromDatabase_userName": null,
      "adp_abortWfOnFailure": true,
      "adp_receiveMetaDataFromDatabase_pollingInterval": 30,
      "adp_receiveMetaDataFromDatabase_fetchSize": 50,
      "adp_cleanUpHistory": false,
      "adp_receiveMetaDataFromDatabase_quoteCharacter": "\"",
      "adp_receiveMetaDataFromDatabase_outputParametersToCsvFileTable": [],
      "adp_loggingEnabled": true,
      "adp_receiveMetaDataFromDatabase_expectedDatabaseValue": "true",
      "adp_receiveMetaDataFromDatabase_userPassword": null,
      "adp_receiveMetaDataFromDatabase_charset": "UTF-8",
      "adp_taskTimeout": 0,
      "adp_receiveMetaDataFromDatabase_defaultValueInCaseOfNull": null,
      "adp_receiveMetaDataFromDatabase_outputParameters": [],
      "adp_receiveMetaDataFromDatabase_fieldSeperator": ";",
      "adp_receiveMetaDataFromDatabase_driverClassPath": "default"
    },
    "taskDescription": "Creates or overrides new meta data key/value pairs from database values for the workflow",
    "taskDisplayName": "Receive meta data from database"
  },
  {
    "taskType": "Insert Documents",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_insertDocuments_enginePassword": null,
      "adp_insertDocuments_parser": null,
      "adp_insertDocuments_fieldNameMapping": [],
      "adp_taskActive": true,
      "adp_insertDocuments_preconfiguredData": [],
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_insertDocuments_engineUser": null,
      "adp_cleanUpHistory": false,
      "adp_insertDocuments_overwriteExistingDocuments": null,
      "adp_insertDocuments_engineIdentifier": null,
      "adp_loggingEnabled": true,
      "adp_insertDocuments_descriptorFilePath": null,
      "adp_taskTimeout": 0
    },
    "taskDescription": "Insert documents into an engine.",
    "taskDisplayName": "Insert documents task"
  },
  {
    "taskType": "Read Documents",
    "taskConfiguration": {
      "adp_readDocuments_applicationType": "",
      "adp_progressTaskTimeout": 0,
      "adp_readDocuments_startIndex": "0",
      "adp_readDocuments_engineGlobalSearch": "",
      "adp_taskActive": true,
      "adp_readDocuments_resultSize": "20",
      "adp_readDocuments_engineTaxonomies": [],
      "adp_readDocuments_engineQuery": "*",
      "adp_executionPersistent": true,
      "adp_readDocuments_engineUserName": "",
      "adp_abortWfOnFailure": true,
      "adp_readDocuments_mainQueryType": null,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_readDocuments_engineName": "",
      "adp_readDocuments_engineFieldLimit": "1000",
      "adp_readDocuments_engineField": "",
      "adp_readDocuments_resultName": "adp_readDocuments_result",
      "adp_readDocuments_applicationIdentifier": "",
      "adp_taskTimeout": 0,
      "adp_readDocuments_engineType": "true",
      "adp_readDocuments_engineUserPassword": ""
    },
    "taskDescription": "Read documents.",
    "taskDisplayName": "Read documents task"
  },
  {
    "taskType": "Data Model Synchronization",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_dms_synchronizationInfoOutputFilePath": "synchronization_info_file_path",
      "adp_dms_excludeFields": "rm_application_level_security,rm_document_hold,rm_stored_searches,rm_exports,rm_resolutions",
      "adp_dms_targetEngineUser": null,
      "adp_dms_synchronizationInfo": "synchronization_info",
      "adp_taskActive": true,
      "adp_dms_synchronizedFields": "synchronized_fields",
      "adp_executionPersistent": true,
      "adp_dms_engineIdentifiersToFix": "engines_to_fix",
      "adp_abortWfOnFailure": true,
      "adp_dms_sourceDataModelIdentifier": null,
      "adp_dms_engineIdentifiersToRestart": "engines_to_restart",
      "adp_cleanUpHistory": false,
      "adp_dms_synchronizationInfoOutputFile": null,
      "adp_dms_targetDataModelIdentifier": null,
      "adp_loggingEnabled": true,
      "adp_dms_targetApplicationIdentifier": "",
      "adp_dms_engineIdentifiersUpdated": "engines_udpated",
      "adp_taskTimeout": 0,
      "adp_dms_sourceApplicationIdentifier": "",
      "adp_dms_targetEnginePassword": null,
      "adp_dms_sourceECAApplicationIdentifier": "{adp_create_application_application_identifier}"
    },
    "taskDescription": "Syncronizeds field tables of source and target data models",
    "taskDisplayName": "Data model synchronization task"
  },
  {
    "taskType": "Publish To Investigation",
    "taskConfiguration": {
      "adp_pti_publishApplicationId": "adp_publish_application_id",
      "adp_progressTaskTimeout": 0,
      "adp_pti_webservicePassword": "",
      "adp_pti_matterSpecificApplication": null,
      "adp_pti_httpsTrustCertificate": "",
      "adp_taskActive": true,
      "adp_pti_matterSpecificUrlRegEx": null,
      "adp_pti_matterId": "default",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_pti_startLearner": "false",
      "adp_loggingEnabled": true,
      "adp_pti_mode": "all",
      "adp_pti_httpsPassword": "",
      "adp_pti_enforceDeduplication": false,
      "adp_pti_publishEngineId": "adp_publish_engine_id",
      "adp_pti_usedWebserviceUrl": "adp_url_used_for_publish_to_review",
      "adp_pti_secondsBetweenNextTryToWaitForMatterCompletion": "100",
      "adp_pti_httpsKeystoreFile": null,
      "adp_pti_publishApplicationUrl": "adp_publish_application_url",
      "adp_pti_webserviceUser": "{service_user}",
      "adp_pti_httpsAllowUntrustedHosts": "false",
      "adp_pti_deduplicate": false,
      "adp_pti_webserviceRequestTimeoutSeconds": 900,
      "adp_pti_waitForMatterCompletion": true,
      "adp_cleanUpHistory": false,
      "adp_pti_searchDetails": [],
      "adp_pti_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_pti_webserviceConnectTimeoutSeconds": 300,
      "adp_taskTimeout": 0,
      "adp_pti_searchString": null
    },
    "taskDescription": "This task publishes documents from a matter into the Investigation application using MatterSpecific web-service",
    "taskDisplayName": "Task to publish into Investigation application"
  },
  {
    "taskType": "Import Storage",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_importStorage_renamePerlSubst": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_importStorage_importLocationsJson": "",
      "adp_importStorage_shutdownEngineAfterImport": "true"
    },
    "taskDescription": "Imports storages of a project",
    "taskDisplayName": "Import Storage"
  },
  {
    "taskType": "APWT Failed Documents Checker",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_checkFailed_folder": null,
      "adp_checkFailed_tempDirectory": null,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Checks for documents that failed during processing",
    "taskDisplayName": "APWT Failed Documents Checker Task"
  },
  {
    "taskType": "CLI",
    "taskConfiguration": {
      "adp_batchScriptRedirectLogging": "false",
      "adp_progressTaskTimeout": 0,
      "adp_envVariables": null,
      "adp_taskActive": true,
      "adp_pathEntries": null,
      "adp_batchScriptPositiveExecutionCodes": "0",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_batchScriptResultCode": "cli_result",
      "adp_cleanUpHistory": false,
      "adp_batchScriptResultLogPath": "cli_result_path",
      "adp_loggingEnabled": true,
      "adp_batchScriptFilterPasswords": true,
      "adp_batchScriptLoggingDirectory": null,
      "adp_batchScriptErrorLogPath": "cli_error_path",
      "adp_taskTimeout": 0,
      "adp_batchScriptPath": null,
      "adp_workingDirectory": null,
      "adp_batchScriptParameters": [],
      "adp_batchScriptJsonLogOutput": ""
    },
    "taskDescription": "Runs a native task in its own process",
    "taskDisplayName": "Command Line Task"
  },
  {
    "taskType": "Compute Counts",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_computeCounts_advancedField": null,
      "adp_computeCounts_engineUserPassword": "",
      "adp_taskActive": true,
      "adp_computeCounts_engineGlobalSearch": "",
      "adp_executionPersistent": true,
      "adp_computeCounts_engineType": "true",
      "adp_computeCounts_engineUserName": "",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_computeCounts_queryParts": "",
      "adp_loggingEnabled": true,
      "adp_computeCounts_types": "[\"standard\"]",
      "adp_computeCounts_resultName": "adpComputeCountResult",
      "adp_computeCounts_engineName": "",
      "adp_taskTimeout": 0,
      "adp_computeCounts_applicationIdentifier": "",
      "adp_computeCounts_engineQuery": "*",
      "adp_computeCounts_applicationType": "",
      "adp_computeCounts_engineTaxonomies": [],
      "adp_computeCounts_mainQueryType": null
    },
    "taskDescription": "Compute Counts.",
    "taskDisplayName": "Compute counts task"
  },
  {
    "taskType": "Search Based Enrichment",
    "taskConfiguration": {
      "adp_enrichment_applicationType": "",
      "adp_progressTaskTimeout": 0,
      "adp_enrichment_parameters": "",
      "adp_enrichment_engineUserPassword": "",
      "adp_taskActive": true,
      "adp_enrichment_engineName": "",
      "adp_enrichment_type": "",
      "adp_executionPersistent": true,
      "adp_enrichment_name": "",
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_enrichment_applicationIdentifier": "",
      "adp_enrichment_engineUserName": "",
      "adp_loggingEnabled": true,
      "adp_enrichment_engineQuery": "*",
      "adp_enrichment_engineGlobalSearch": "",
      "adp_enrichment_engineTaxonomies": [],
      "adp_enrichment_engineType": "true",
      "adp_taskTimeout": 0,
      "adp_enrichment_mainQueryType": null,
      "adp_enrichment_priority": "10"
    },
    "taskDescription": "Search Based Enrichment",
    "taskDisplayName": "Search based enrichment task"
  },
  {
    "taskType": "APWT Check Config",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_crawlerUmbrellaParentAgreementIdIdentifier": null,
      "adp_crawlerBatchIdIdentifier": null,
      "adp_crawlerAttachementParentIdentifier": null,
      "adp_taskActive": true,
      "adp_crawlerWorkflowBatchIdentifier": null,
      "adp_crawlerAmendmentsIdentifier": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_dsName": "",
      "adp_cleanUpHistory": false,
      "adp_crawlerFundsNameIdentifier": null,
      "adp_crawlerAttachementPrimaryIdentifier": null,
      "adp_loggingEnabled": true,
      "adp_crawlerFilePathIdentifier": null,
      "adp_taskTimeout": 0,
      "adp_crawlerTitleIdentifier": null,
      "adp_dsSubfundName": "",
      "adp_crawlerTechnicalAmendmentsIdentifier": null,
      "adp_crawlerAttachementRootIdentifier": null
    },
    "taskDescription": "Runs sanity checks on crawlers and features",
    "taskDisplayName": "APWT Check Config Task"
  },
  {
    "taskType": "CSV Converter",
    "taskConfiguration": {
      "adp_csvConverter_outputVariable": "adp_new_csv_file",
      "adp_progressTaskTimeout": 0,
      "adp_csvConverter_csvFileAppend": false,
      "adp_csvConverter_fieldSeparator": ";",
      "adp_taskActive": true,
      "adp_csvConverter_csvFile": "new.csv",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_csvConverter_valueDelimiter": "\"",
      "adp_loggingEnabled": true,
      "adp_csvConverter_fileEncoding": "UTF-8",
      "adp_taskTimeout": 0,
      "adp_csvConverter_csvMappingAndTargetDefinition": [],
      "adp_csvConverter_csvSourceData": []
    },
    "taskDescription": "Converts data from the dynamic table data service to a new csv file.",
    "taskDisplayName": "CSV Converter"
  },
  {
    "taskType": "Extract Zip",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_extractZip_extractedFileLocation": "adp_extracted_zip_location",
      "adp_extractZip_success": "adp_extracted_zip_success",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_extractZip_deleteAfterUnzip": "false",
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_extractZip_sourceZipFile": null,
      "adp_extractZip_targetExtractionLocation": null,
      "adp_extractZip_passwordsJson": ""
    },
    "taskDescription": "Extracts a Zip to a target location.",
    "taskDisplayName": "Extract Zip"
  },
  {
    "taskType": "Add Subengine",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_addSubengine_jvmOptions": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_addSubengine_hostIdentifier": null,
      "adp_addSubengine_chosenEngineHostMemory": "adp_created_engine_host_memory",
      "adp_addSubengine_hostMemoryLimitRatio": "0",
      "adp_addSubengine_hostMemoryLimit": "0",
      "adp_loggingEnabled": true,
      "adp_addSubengine_heapSetting": null,
      "adp_addSubengine_engineName": "{engine_name}",
      "adp_addSubengine_abortOnExistingEngine": false,
      "adp_taskTimeout": 0,
      "adp_addSubengine_singleEngineTemplate": null,
      "adp_addSubengine_chosenEngineHostMemoryRatio": "adp_created_engine_host_memory_ratio",
      "adp_addSubengine_createdEngineIdentifier": "adp_created_engine_name",
      "adp_addSubengine_mergingMetaEngineTemplate": "mergingMeta.documentHoldMetaSystemTemplate",
      "adp_addSubengine_applicationIdentifier": "{application_id}"
    },
    "taskDescription": "Adds sub-engine to an existing application.",
    "taskDisplayName": "Add subengine task"
  },
  {
    "taskType": "Add Fields",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_addFields_engineType": "true",
      "adp_addFields_dataModelName": "",
      "adp_taskActive": true,
      "adp_addFields_applicationName": null,
      "adp_addFields_engineName": null,
      "adp_addFields_cloneTemplates": [
        {
          "Type": "Smart Filter",
          "Field": "meta_bcc"
        },
        {
          "Type": "Text field",
          "Field": "attachment"
        },
        {
          "Type": "Unique Dates",
          "Field": "meta_docdate"
        },
        {
          "Type": "Grouped Dates",
          "Field": "meta_bcc"
        }
      ],
      "adp_executionPersistent": true,
      "adp_addFields_enginePassword": null,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_addFields_engineUser": null,
      "adp_taskTimeout": 0,
      "adp_addFields_fields2Clone": []
    },
    "taskDescription": "Adds fields to an index.",
    "taskDisplayName": "Add fields task"
  },
  {
    "taskType": "Engine Job Monitoring",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_engineJobMonitoring_jobType": "",
      "adp_engineJobMonitoring_jsonFile": "output.json",
      "adp_taskActive": true,
      "adp_engineJobMonitoring_outputFilename": "adp_engineJobMonitoring_output_file_name",
      "adp_executionPersistent": true,
      "adp_engineJobMonitoring_engineUserName": "{adp_user}",
      "adp_abortWfOnFailure": true,
      "adp_engineJobMonitoring_startingJob": "1",
      "adp_engineJobMonitoring_engineIds": "",
      "adp_cleanUpHistory": false,
      "adp_engineJobMonitoring_waitForJobs": "",
      "adp_loggingEnabled": true,
      "adp_engineJobMonitoring_engineUserPassword": "",
      "adp_engineJobMonitoring_outputJson": "adp_engineJobMonitoring_json_output",
      "adp_taskTimeout": 0,
      "adp_engineJobMonitoring_jobStates": "RUNNING,STOPPING,PAUSED,PAUSING,FINISHING,WAITING,ATOMIC",
      "adp_engineJobMonitoring_startedAfter": "2015-01-01T12:00:00",
      "adp_engineJobMonitoring_listOfJobProperties": "",
      "adp_engineJobMonitoring_workspaceId": "",
      "adp_engineJobMonitoring_numberOfJobs": "-1",
      "adp_engineJobMonitoring_appIds": ""
    },
    "taskDescription": "Monitor/list engine jobs.",
    "taskDisplayName": "Engine Job Monitoring"
  },
  {
    "taskType": "Start Data Source",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_startDataSource_engineUser": null,
      "adp_startDataSource_ignoreSuspended": "false",
      "adp_startDataSource_indexedDsDoccount": "adp_indexedDsDoccount",
      "adp_startDataSource_expectedDsDoccount": "adp_expectedDsDoccount",
      "adp_startDataSource_numberDocuments": "adp_numberDocuments",
      "adp_taskActive": true,
      "adp_startDataSource_numberDocumentsToCrawl": "{number_of_records}",
      "adp_startDataSource_selectedHost": "adp_start_data_source_host",
      "adp_startDataSource_oldDsDoccount": "adp_oldDsDoccount",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_startDataSource_hostIdentifiers": null,
      "adp_cleanUpHistory": false,
      "adp_startDataSource_dataSourceName": "{adp_configured_data_source_name}",
      "adp_loggingEnabled": true,
      "adp_startDataSource_synchronous": true,
      "adp_startDataSource_loadBalancingEnabled": "false",
      "adp_startDataSource_enginePassword": null,
      "adp_taskTimeout": 0,
      "adp_startDataSource_maxNumberRunningCrawlers": "0",
      "adp_startDataSource_maxUnknownStatusPollingDurationSeconds": 300,
      "adp_startDataSource_hostRolesBlackList": null
    },
    "taskDescription": "Starts a data source",
    "taskDisplayName": "Start data source"
  },
  {
    "taskType": "Publish To Review",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_ptr_matterSpecificUrlRegEx": null,
      "adp_taskActive": true,
      "adp_ptr_matterId": "default",
      "adp_executionPersistent": true,
      "adp_ptr_webserviceRequestTimeoutSeconds": 900,
      "adp_abortWfOnFailure": true,
      "adp_ptr_searchDetails": [],
      "adp_loggingEnabled": true,
      "adp_ptr_publishEngineId": "adp_publish_engine_id",
      "adp_ptr_startLearner": "false",
      "adp_ptr_searchString": null,
      "adp_ptr_waitForMatterCompletion": true,
      "adp_ptr_usedWebserviceUrl": "adp_url_used_for_publish_to_review",
      "adp_ptr_secondsBetweenNextTryToWaitForMatterCompletion": "100",
      "adp_ptr_ecaEngine": "{project_engine}",
      "adp_ptr_ecaMasterPort": "",
      "adp_ptr_httpsKeystoreFile": null,
      "adp_ptr_mode": "all",
      "adp_ptr_webserviceConnectTimeoutSeconds": 300,
      "adp_ptr_publishApplicationId": "adp_publish_application_id",
      "adp_ptr_httpsPassword": "",
      "adp_ptr_webservicePassword": "",
      "adp_ptr_webserviceUrl": "http://{host_name}/{project_name}",
      "adp_ptr_publishApplicationUrl": "adp_publish_application_url",
      "adp_ptr_enforceDeduplication": false,
      "adp_cleanUpHistory": false,
      "adp_ptr_matterSpecificApplication": null,
      "adp_ptr_webserviceUser": "{service_user}",
      "adp_taskTimeout": 0,
      "adp_ptr_deduplicate": false,
      "adp_ptr_ecaApplication": "{project_name}",
      "adp_ptr_httpsAllowUntrustedHosts": "false",
      "adp_ptr_ecaPublish": false,
      "adp_ptr_ecaMasterHost": "",
      "adp_ptr_httpsTrustCertificate": ""
    },
    "taskDescription": "",
    "taskDisplayName": "Task to publish stuff"
  },
  {
    "taskType": "Import",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_importTask_importDataJson": "",
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_importTask_renamePerlSubst": "",
      "adp_cleanUpHistory": false,
      "adp_importTask_inheritFromTemplate": "",
      "adp_loggingEnabled": true,
      "adp_importTask_dryRun": "false",
      "adp_importTask_skipStorageLocationImport": "false",
      "adp_taskTimeout": 0,
      "adp_importTask_debug": "false",
      "adp_importTask_importProjectData": "true",
      "adp_importTask_savedSearchesFile": null,
      "adp_importTask_mode": "IMPORT_NEW_SECURE",
      "adp_importTask_shutdownEngineAfterImport": "true"
    },
    "taskDescription": "Interface to importApplicationTree",
    "taskDisplayName": "Import"
  },
  {
    "taskType": "CSV Merge",
    "taskConfiguration": {
      "adp_csvMerge_noUniqueMatch": "false",
      "adp_csvMerge_noFlushAfterMerge": "false",
      "adp_csvMerge_forceChange": "false",
      "adp_csvMerge_engineName": null,
      "adp_csvMerge_applicationIdentifier": "",
      "adp_csvMerge_mergeType": "Merge content",
      "adp_csvMerge_displayNameMappingErrorFile": null,
      "adp_csvMerge_lockDocumentChanges": "false",
      "adp_csvMerge_csvFieldNativeLocation": null,
      "adp_executionPersistent": true,
      "adp_csvMerge_doNotChangeProtectedDocuments": "false",
      "adp_csvMerge_imageBasePath": null,
      "adp_loggingEnabled": true,
      "adp_csvMerge_nullValue": "",
      "adp_csvMerge_deduplicateWhenAppending": "true",
      "adp_csvMerge_lineSeparatorForFulltext": "",
      "adp_csvMerge_allowUpdateMultipleDocs": "false",
      "adp_csvMerge_csvIdPostfix": null,
      "adp_csvMerge_origFile": null,
      "adp_csvMerge_maxCategoryLength": "128",
      "adp_csvMerge_fieldSeperator": ";",
      "adp_csvMerge_applicationType": "",
      "adp_csvMerge_enginePassword": null,
      "adp_taskTimeout": 0,
      "adp_csvMerge_charset": "UTF-8",
      "adp_progressTaskTimeout": 0,
      "adp_csvMerge_textIndicator": "",
      "adp_csvMerge_noFlushBeforeMerge": "false",
      "adp_taskActive": true,
      "adp_csvMerge_engineIdFieldKey": null,
      "adp_abortWfOnFailure": true,
      "adp_csvMerge_customLineSeparator": "U+000DU+000A",
      "adp_csvMerge_multiValueDelimiter": null,
      "adp_csvMerge_csvFile": null,
      "adp_csvMerge_csvIdFieldKey": null,
      "adp_csvMerge_textFileRefIndicator": "",
      "adp_csvMerge_engineUser": null,
      "adp_csvMerge_textFileCharset": "",
      "adp_csvMerge_doNotAddTimestampToOutputFiles": "false",
      "adp_csvMerge_nativeBasePath": null,
      "adp_csvMerge_csvMode": "append",
      "adp_csvMerge_fieldMappings": [],
      "adp_cleanUpHistory": false,
      "adp_csvMerge_csvMergeConfiguration": "",
      "adp_csvMerge_errorLogFile": null,
      "adp_csvMerge_dryRun": "false",
      "adp_csvMerge_csvFieldImageLocation": null,
      "adp_csvMerge_matchesLogFile": null,
      "adp_csvMerge_csvIdPrefix": null,
      "adp_csvMerge_expandToFamily": "false"
    },
    "taskDescription": "Merges metaData or natives/images by using a csv file.",
    "taskDisplayName": "Csv merge task"
  },
  {
    "taskType": "Crawler Monitoring",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_CrawlerMonitoring_waitOnUnknown": "true",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_CrawlerMonitoring_crawlerList": []
    },
    "taskDescription": "Monitors a set of crawlers and waits for completion/failure of all crawlers",
    "taskDisplayName": "Crawler Monitoring Task"
  },
  {
    "taskType": "CSV Copy Header",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_csvCopyHeader_targetFile": "",
      "adp_csvCopyHeader_origFile": null,
      "adp_csvCopyHeader_lineSeparator": null,
      "adp_csvCopyHeader_targetFilePath": "adp_new_csv_file",
      "adp_taskActive": true,
      "adp_csvCopyHeader_aliasedFile": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_taskTimeout": 0,
      "adp_csvCopyHeader_encoding": "{detected_encoding}"
    },
    "taskDescription": "Task to copy the header of a CSV file.",
    "taskDisplayName": "Copy CSV header task"
  },
  {
    "taskType": "Consumer Trigger",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Empty Task to trigger message consumers.",
    "taskDisplayName": "Consumer Trigger"
  },
  {
    "taskType": "Check for entity status",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_checkForEntityStatus_processStatusConfigTable": []
    },
    "taskDescription": "Retrieves a status of a given entity (data source, application or engine)",
    "taskDisplayName": "Check for entity status"
  },
  {
    "taskType": "Meta Data Changer",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_metaDataChanger_changeMetaDataConfig": [],
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Changes metaData by using regEx replacement.",
    "taskDisplayName": "Meta data changer"
  },
  {
    "taskType": "PSS Read Self-Service Configuration",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "pss_ReadConfigModelDataJSON": "pss_model_data_json",
      "pss_ReadConfigRestPassword": null,
      "adp_taskActive": true,
      "pss_ReadConfigModelChangesetId": "pss_changeset_id",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "pss_ReadConfigVersionId": null,
      "pss_ReadConfigModelDisplayName": "pss_model_display_name",
      "adp_taskTimeout": 0,
      "pss_ReadConfigRestUserName": null,
      "pss_ReadConfigPssURL": null,
      "pss_ReadConfigModelId": null,
      "pss_ReadConfigAuthURL": null,
      "pss_ReadConfigPageSize": 5000
    },
    "taskDescription": "Reads configuration from self-service back-end.",
    "taskDisplayName": "PSS Read Self-Service Configuration"
  },
  {
    "taskType": "CSV File Reduction",
    "taskConfiguration": {
      "adp_fileReduction_enableOriginalLineCountDetection": false,
      "adp_progressTaskTimeout": 0,
      "adp_fileReduction_originalFileName": "new.csv",
      "adp_fileReduction_lineCountReducedFile": "reduced_file_line_count",
      "adp_fileReduction_lineCountOriginalFile": "original_file_line_count",
      "adp_taskActive": true,
      "adp_fileReduction_charset": "UTF-8",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_fileReduction_lastNumberOfLines": 0,
      "adp_fileReduction_numberOfLines": 100,
      "adp_fileReduction_reducedFile": "reduced_file",
      "adp_taskTimeout": 0,
      "adp_fileReduction_startLine": 1,
      "adp_fileReduction_reducedFileName": "reduced.csv"
    },
    "taskDescription": "Reduces a source file (csv) to a smaller excerpt.",
    "taskDisplayName": "CSV File Reduction"
  },
  {
    "taskType": "Convert Opticon To Csv",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_convertOptToCsv_storeAbsolutePaths": true,
      "adp_taskActive": true,
      "adp_convertOptToCsv_csvMultiValueSeparator": "#",
      "adp_convertOptToCsv_opticonFile": null,
      "adp_convertOptToCsv_csvFile": null,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_convertOptToCsv_csvValueDelimiter": "\"",
      "adp_convertOptToCsv_csvLineSeparator": "u+000Du+000A",
      "adp_cleanUpHistory": false,
      "adp_convertOptToCsv_invalidImages": "invalid_images",
      "adp_loggingEnabled": true,
      "adp_convertOptToCsv_checkImageFormat": false,
      "adp_convertOptToCsv_numberOfInvalidImages": "number_of_invalid_images",
      "adp_convertOptToCsv_csvIdColumnName": "ID",
      "adp_convertOptToCsv_csvImageReferencesColumnName": "IMAGES",
      "adp_convertOptToCsv_csvFieldSeparator": ",",
      "adp_taskTimeout": 0,
      "adp_convertOptToCsv_csvFileEncoding": "UTF-8"
    },
    "taskDescription": "Converts an Opticon file to CSV format.",
    "taskDisplayName": "Convert Opticon To Csv"
  },
  {
    "taskType": "Import Storage Locations",
    "taskConfiguration": {
      "adp_importStorageLocations_shutdownEngineAfterImport": "true",
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_importStorageLocations_renamePerlSubst": "",
      "adp_importStorageLocations_importLocationsJson": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Interface to importStorageLocations",
    "taskDisplayName": "Import Storage Locations"
  },
  {
    "taskType": "No Operation",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false
    },
    "taskDescription": "No operation task.",
    "taskDisplayName": "No operation task"
  },
  {
    "taskType": "Export Matter",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_exportMatter_enginePassword": null,
      "adp_exportMatter_matterId": null,
      "adp_exportMatter_webserviceUrl": null,
      "adp_exportMatter_matterDisplayName": null,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_exportMatter_webserviceUser": "",
      "adp_loggingEnabled": true,
      "adp_exportMatter_processedMatterId": "adp_processed_matter_id",
      "adp_exportMatter_httpsKeystoreFile": null,
      "adp_exportMatter_globalSearchId": null,
      "adp_exportMatter_name": null,
      "adp_exportMatter_httpsPassword": "",
      "adp_exportMatter_usedWebserviceUrl": "eca_url_used_for_matter_export",
      "adp_exportMatter_webservicePassword": null,
      "adp_exportMatter_searchQuery": null,
      "adp_exportMatter_waitForExport": "false",
      "adp_exportMatter_type": "Export",
      "adp_exportMatter_savedSearchId": null,
      "adp_exportMatter_engineUser": null,
      "adp_exportMatter_webserviceRequestTimeoutSeconds": 900,
      "adp_exportMatter_matterSpecificApplication": null,
      "adp_exportMatter_httpsAllowUntrustedHosts": "false",
      "adp_cleanUpHistory": false,
      "adp_exportMatter_matterSpecificUrlRegEx": null,
      "adp_exportMatter_webserviceConnectTimeoutSeconds": 300,
      "adp_exportMatter_httpsTrustCertificate": "",
      "adp_taskTimeout": 0,
      "adp_exportMatter_globalSearchJson": "",
      "adp_exportMatter_exportConfig": null,
      "adp_exportMatter_exportDestinations": "adp_matter_exportDestinations",
      "adp_exportMatter_exportName": "adp_matter_exportName",
      "adp_exportMatter_mainQueryType": null
    },
    "taskDescription": "Export matter.",
    "taskDisplayName": "Export Matter"
  },
  {
    "taskType": "Health",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_cleanUpHistory": false,
      "adp_health_outputJson": "adp_health_json_output"
    },
    "taskDescription": "Retrieves ADP end-point system health attributes",
    "taskDisplayName": "ADP end-point health"
  },
  {
    "taskType": "PSS Create XHTML Files",
    "taskConfiguration": {
      "pss_createXhtmlFiles_inputXhtmlDirectory": "",
      "adp_progressTaskTimeout": 0,
      "adp_loggingEnabled": true,
      "pss_createXhtmlFiles_inputModelDataJson": "",
      "pss_createXhtmlFiles_relativeXhtmlBaseFilePath": "",
      "adp_taskActive": true,
      "adp_taskTimeout": 0,
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "pss_createXhtmlFiles_outputDirectory": "pss_xhtml_files_location",
      "adp_cleanUpHistory": false
    },
    "taskDescription": "Creates the Perceptiv XHTML Files.",
    "taskDisplayName": "PSS Create XHTML Files"
  },
  {
    "taskType": "Move Files",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_moveFiles_allowBucketCreation": "false",
      "adp_taskActive": true,
      "adp_moveFiles_secretKey": "",
      "adp_executionPersistent": true,
      "adp_abortWfOnFailure": true,
      "adp_moveFiles_moveFilesConditionsTable": [],
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_moveFiles_numberThreads": "",
      "adp_moveFiles_maxNumRetries": "",
      "adp_taskTimeout": 0,
      "adp_moveFiles_accessKey": "",
      "adp_moveFiles_cloudServiceName": "No Cloud Service",
      "adp_moveFiles_newBucketS3Location": "eu-central-1"
    },
    "taskDescription": "Move files to defined folders depending on results of workflow and tasks.",
    "taskDisplayName": "Move files task"
  },
  {
    "taskType": "Read Service Alerts",
    "taskConfiguration": {
      "adp_progressTaskTimeout": 0,
      "adp_taskActive": true,
      "adp_executionPersistent": true,
      "adp_readServiceAlerts_date": null,
      "adp_abortWfOnFailure": true,
      "adp_readServiceAlerts_blacklist": "",
      "adp_cleanUpHistory": false,
      "adp_loggingEnabled": true,
      "adp_readServiceAlerts_lastDate": "adp_readServiceAlerts_first_skipped_date",
      "adp_taskTimeout": 0,
      "adp_readServiceAlerts_listOfProperties": "",
      "adp_readServiceAlerts_maximum": "1000",
      "adp_readServiceAlerts_outputJson": "adp_readServiceAlerts_json_output"
    },
    "taskDescription": "Changes the display name of an application.",
    "taskDisplayName": "Read Service Alerts"
  }
]
```
 