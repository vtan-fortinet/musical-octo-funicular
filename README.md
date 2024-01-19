Log
```
INFO[0000]log.go:216 gosnowflake.(*defaultLogger).Info Open                                         
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Query String: application=<REDACTED>&database=<REDACTED>&ocspFailOpen=true&schema=<REDACTED>&tmpDirPath=%2Ftmp%2F&validateDefaultParameters=true&warehouse=<REDACTED>
INFO[0000]log.go:216 gosnowflake.(*defaultLogger).Info OpenWithConfig                               
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Authenticating via SNOWFLAKE                 
INFO[0000]auth.go:353 gosnowflake.authenticate PARAMS for Auth: &map[databaseName:[<REDACTED>] schemaName:[<REDACTED>] warehouse:[<REDACTED>]], https, ytb20147.snowflakecomputing.com, 443, 5m0s, SNOWFLAKE 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof full URL: https://ytb20147.snowflakecomputing.com:443/session/v1/login-request?databaseName=<REDACTED>&requestId=5e279685-a583-4709-73ef-25b32de3ce28&request_guid=16ca554c-ad61-4933-6c5f-e9620269be26&schemaName=<REDACTED>&warehouse=<REDACTED>
INFO[0000]retry.go:300 gosnowflake.(*retryHTTP).execute retryHTTP.totalTimeout: 5m0s                 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf retry count: 0                               
INFO[0000]log.go:216 gosnowflake.(*defaultLogger).Info Username and password                        
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=*.prod3.us-west-2.snowflakecomputing.com, Issuer: CN=Amazon RSA 2048 M02,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Amazon RSA 2048 M02,O=Amazon,C=US, Issuer: CN=Amazon Root CA 1,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 iH2kRF5n6nyUd05DGJw+zuTIcxI= nF8A36oB1zArOIiiuG1KnPIRkYM= 144918191876577076464031512351042010504348870} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=*.prod3.us-west-2.snowflakecomputing.com, Issuer: CN=Amazon RSA 2048 M02,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Amazon RSA 2048 M02,O=Amazon,C=US, Issuer: CN=Amazon Root CA 1,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Amazon Root CA 1,O=Amazon,C=US, Issuer: CN=Starfield Services Root Certificate Authority - G2,O=Starfield Technologies\, Inc.,L=Scottsdale,ST=Arizona,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 iH2kRF5n6nyUd05DGJw+zuTIcxI= nF8A36oB1zArOIiiuG1KnPIRkYM= 144918191876577076464031512351042010504348870} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 iH2kRF5n6nyUd05DGJw+zuTIcxI= nF8A36oB1zArOIiiuG1KnPIRkYM= 144918191876577076464031512351042010504348870} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 i8GehFuYHWHPVGkhGmi44xEzbZA= v1+30c7dH4b0W1Ws3NcQwg6piOc= 12037640545166866303} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=*.prod3.us-west-2.snowflakecomputing.com, Issuer: CN=Amazon RSA 2048 M02,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 Zm0LcJyJ5L5L6+wTRUfktpU2D4w= wDFSzVpQw4J8dHHOy+mc+XrrguI= 16707542143754441499341827299770005449} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Amazon RSA 2048 M02,O=Amazon,C=US, Issuer: CN=Amazon Root CA 1,O=Amazon,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 T1mjlFPPuVWef2vYxU2lPaZCtxQ= hBjMhTTsvAyUlC4IWZzHshBOCgg= 166129353110899469622597955040406457904926625} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Amazon Root CA 1,O=Amazon,C=US, Issuer: CN=Starfield Services Root Certificate Authority - G2,O=Starfield Technologies\, Inc.,L=Scottsdale,ST=Arizona,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 iH2kRF5n6nyUd05DGJw+zuTIcxI= nF8A36oB1zArOIiiuG1KnPIRkYM= 144918191876577076464031512351042010504348870} 
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof Subject: CN=Starfield Services Root Certificate Authority - G2,O=Starfield Technologies\, Inc.,L=Scottsdale,ST=Arizona,C=US, Issuer: OU=Starfield Class 2 Certification Authority,O=Starfield Technologies\, Inc.,C=US 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf OCSP status found in cache: &{-1 <nil>}; certIdKey: &{SHA-1 i8GehFuYHWHPVGkhGmi44xEzbZA= v1+30c7dH4b0W1Ws3NcQwg6piOc= 12037640545166866303} 
INFO[0000]log.go:216 gosnowflake.(*defaultLogger).Info Authentication SUCCESS                       
INFO[0000]connection_util.go:141 gosnowflake.(*snowflakeConn).populateSessionParameters params: []gosnowflake.nameValueParameter{gosnowflake.nameValueParameter{Name:"CLIENT_PREFETCH_THREADS", Value:4}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_OUTPUT_FORMAT", Value:"YYYY-MM-DD HH24:MI:SS.FF3 TZHTZM"}, gosnowflake.nameValueParameter{Name:"TIME_OUTPUT_FORMAT", Value:"HH24:MI:SS"}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_CHUNK_SIZE", Value:160}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_TZ_OUTPUT_FORMAT", Value:""}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_KEEP_ALIVE", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_OUT_OF_BAND_TELEMETRY_ENABLED", Value:false}, gosnowflake.nameValueParameter{Name:"QUERY_CONTEXT_CACHE_SIZE", Value:5}, gosnowflake.nameValueParameter{Name:"CLIENT_METADATA_USE_SESSION_DATABASE", Value:false}, gosnowflake.nameValueParameter{Name:"ENABLE_STAGE_S3_PRIVATELINK_FOR_US_EAST_1", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_PREFETCH_THREADS", Value:1}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_NTZ_OUTPUT_FORMAT", Value:"YYYY-MM-DD HH24:MI:SS.FF3"}, gosnowflake.nameValueParameter{Name:"CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_HONOR_CLIENT_TZ_FOR_TIMESTAMP_NTZ", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_MEMORY_LIMIT", Value:1536}, gosnowflake.nameValueParameter{Name:"CLIENT_TIMESTAMP_TYPE_MAPPING", Value:"TIMESTAMP_LTZ"}, gosnowflake.nameValueParameter{Name:"TIMEZONE", Value:"America/Los_Angeles"}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_PREFETCH_SLOTS", Value:2}, gosnowflake.nameValueParameter{Name:"CLIENT_TELEMETRY_ENABLED", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_USE_V1_QUERY_API", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_DISABLE_INCIDENTS", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_COLUMN_CASE_INSENSITIVE", Value:false}, gosnowflake.nameValueParameter{Name:"CSV_TIMESTAMP_FORMAT", Value:""}, gosnowflake.nameValueParameter{Name:"BINARY_OUTPUT_FORMAT", Value:"HEX"}, gosnowflake.nameValueParameter{Name:"CLIENT_ENABLE_LOG_INFO_STATEMENT_PARAMETERS", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_TELEMETRY_SESSIONLESS_ENABLED", Value:true}, gosnowflake.nameValueParameter{Name:"DATE_OUTPUT_FORMAT", Value:"YYYY-MM-DD"}, gosnowflake.nameValueParameter{Name:"CLIENT_FORCE_PROTECT_ID_TOKEN", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_CONSENT_CACHE_ID_TOKEN", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_STAGE_ARRAY_BINDING_THRESHOLD", Value:65280}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY", Value:3600}, gosnowflake.nameValueParameter{Name:"AUTOCOMMIT", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_CLONE", Value:false}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_LTZ_OUTPUT_FORMAT", Value:""}} 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_PREFETCH_THREADS, value: 4 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_OUTPUT_FORMAT, value: YYYY-MM-DD HH24:MI:SS.FF3 TZHTZM 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIME_OUTPUT_FORMAT, value: HH24:MI:SS 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_CHUNK_SIZE, value: 160 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_TZ_OUTPUT_FORMAT, value:  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_KEEP_ALIVE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_OUT_OF_BAND_TELEMETRY_ENABLED, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: QUERY_CONTEXT_CACHE_SIZE, value: 5 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_METADATA_USE_SESSION_DATABASE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: ENABLE_STAGE_S3_PRIVATELINK_FOR_US_EAST_1, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_PREFETCH_THREADS, value: 1 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_NTZ_OUTPUT_FORMAT, value: YYYY-MM-DD HH24:MI:SS.FF3 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_HONOR_CLIENT_TZ_FOR_TIMESTAMP_NTZ, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_MEMORY_LIMIT, value: 1536 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TIMESTAMP_TYPE_MAPPING, value: TIMESTAMP_LTZ 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMEZONE, value: America/Los_Angeles 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_PREFETCH_SLOTS, value: 2 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TELEMETRY_ENABLED, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_USE_V1_QUERY_API, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_DISABLE_INCIDENTS, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_COLUMN_CASE_INSENSITIVE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CSV_TIMESTAMP_FORMAT, value:  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: BINARY_OUTPUT_FORMAT, value: HEX 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_ENABLE_LOG_INFO_STATEMENT_PARAMETERS, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TELEMETRY_SESSIONLESS_ENABLED, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: DATE_OUTPUT_FORMAT, value: YYYY-MM-DD 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_FORCE_PROTECT_ID_TOKEN, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_CONSENT_CACHE_ID_TOKEN, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_STAGE_ARRAY_BINDING_THRESHOLD, value: 65280 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY, value: 3600 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: AUTOCOMMIT, value: true     
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_CLONE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_LTZ_OUTPUT_FORMAT, value:  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf sending 1 logs to telemetry. inband telemetry payload being sent: {"logs":[{"timestamp":1705702592329,"message":{"DriverType":"Go","DriverVersion":"1.7.2","autocommit":"true","binary_output_format":"HEX","client_consent_cache_id_token":"false","client_disable_incidents":"true","client_enable_log_info_statement_parameters":"false","client_force_protect_id_token":"true","client_honor_client_tz_for_timestamp_ntz":"true","client_memory_limit":"1536","client_metadata_request_use_connection_ctx":"false","client_metadata_use_session_database":"false","client_out_of_band_telemetry_enabled":"false","client_prefetch_threads":"4","client_result_chunk_size":"160","client_result_column_case_insensitive":"false","client_result_prefetch_slots":"2","client_result_prefetch_threads":"1","client_session_clone":"false","client_session_keep_alive":"false","client_session_keep_alive_heartbeat_frequency":"3600","client_stage_array_binding_threshold":"65280","client_telemetry_enabled":"true","client_telemetry_sessionless_enabled":"true","client_timestamp_type_mapping":"TIMESTAMP_LTZ","client_use_v1_query_api":"true","csv_timestamp_format":"","date_output_format":"YYYY-MM-DD","enable_stage_s3_privatelink_for_us_east_1":"false","query_context_cache_size":"5","source":"golang_driver","time_output_format":"HH24:MI:SS","timestamp_ltz_output_format":"","timestamp_ntz_output_format":"YYYY-MM-DD HH24:MI:SS.FF3","timestamp_output_format":"YYYY-MM-DD HH24:MI:SS.FF3 TZHTZM","timestamp_tz_output_format":"","timezone":"America/Los_Angeles","type":"client_connection_parameters"}}]} 
INFO[0000]retry.go:300 gosnowflake.(*retryHTTP).execute retryHTTP.totalTimeout: 10s                  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf retry count: 0                               
DEBU[0000]log.go:210 gosnowflake.(*defaultLogger).Debug successfully uploaded metrics to telemetry   
INFO[0000]connection.go:388 gosnowflake.(*snowflakeConn).queryContextInternal Query: "SELECT $1 FROM (VALUES (1)) WHERE $1 = 2;", [] 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf empty qcc                                    
INFO[0000]connection.go:118 gosnowflake.(*snowflakeConn).exec parameters: map[]                            
INFO[0000]connection.go:127 gosnowflake.(*snowflakeConn).exec bindings: map[]                              
INFO[0000]log.go:154 gosnowflake.(*defaultLogger).Infof params: &map[]                               
INFO[0000]retry.go:300 gosnowflake.(*retryHTTP).execute retryHTTP.totalTimeout: 0s                   
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf retry count: 0                               
INFO[0000]restful.go:252 gosnowflake.postRestfulQueryHelper postQuery: resp: &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[no-cache, no-store] Connection:[keep-alive] Content-Type:[application/json] Date:[Fri, 19 Jan 2024 22:16:32 GMT] Expect-Ct:[enforce, max-age=3600] Strict-Transport-Security:[max-age=31536000] Vary:[Accept-Encoding, User-Agent] X-Content-Type-Options:[nosniff] X-Country:[Canada] X-Frame-Options:[deny] X-Xss-Protection:[1; mode=block]] 0xc000036fc0 -1 [] false true map[] 0xc0008ca200 0xc000342c60} 
INFO[0000]connection.go:157 gosnowflake.(*snowflakeConn).exec Success: true, Code: -1                      
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf adding query context: {0 45328376437008 0 }  
INFO[0000]connection.go:180 gosnowflake.(*snowflakeConn).exec Exec/Query SUCCESS                           
INFO[0000]connection_util.go:141 gosnowflake.(*snowflakeConn).populateSessionParameters params: []gosnowflake.nameValueParameter{gosnowflake.nameValueParameter{Name:"CLIENT_PREFETCH_THREADS", Value:4}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_OUTPUT_FORMAT", Value:"YYYY-MM-DD HH24:MI:SS.FF3 TZHTZM"}, gosnowflake.nameValueParameter{Name:"TIME_OUTPUT_FORMAT", Value:"HH24:MI:SS"}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_CHUNK_SIZE", Value:160}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_TZ_OUTPUT_FORMAT", Value:""}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_KEEP_ALIVE", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_OUT_OF_BAND_TELEMETRY_ENABLED", Value:false}, gosnowflake.nameValueParameter{Name:"QUERY_CONTEXT_CACHE_SIZE", Value:5}, gosnowflake.nameValueParameter{Name:"CLIENT_METADATA_USE_SESSION_DATABASE", Value:false}, gosnowflake.nameValueParameter{Name:"ENABLE_STAGE_S3_PRIVATELINK_FOR_US_EAST_1", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_PREFETCH_THREADS", Value:1}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_NTZ_OUTPUT_FORMAT", Value:"YYYY-MM-DD HH24:MI:SS.FF3"}, gosnowflake.nameValueParameter{Name:"CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_HONOR_CLIENT_TZ_FOR_TIMESTAMP_NTZ", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_MEMORY_LIMIT", Value:1536}, gosnowflake.nameValueParameter{Name:"CLIENT_TIMESTAMP_TYPE_MAPPING", Value:"TIMESTAMP_LTZ"}, gosnowflake.nameValueParameter{Name:"TIMEZONE", Value:"America/Los_Angeles"}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_PREFETCH_SLOTS", Value:2}, gosnowflake.nameValueParameter{Name:"CLIENT_TELEMETRY_ENABLED", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_USE_V1_QUERY_API", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_DISABLE_INCIDENTS", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_RESULT_COLUMN_CASE_INSENSITIVE", Value:false}, gosnowflake.nameValueParameter{Name:"CSV_TIMESTAMP_FORMAT", Value:""}, gosnowflake.nameValueParameter{Name:"BINARY_OUTPUT_FORMAT", Value:"HEX"}, gosnowflake.nameValueParameter{Name:"CLIENT_ENABLE_LOG_INFO_STATEMENT_PARAMETERS", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_TELEMETRY_SESSIONLESS_ENABLED", Value:true}, gosnowflake.nameValueParameter{Name:"DATE_OUTPUT_FORMAT", Value:"YYYY-MM-DD"}, gosnowflake.nameValueParameter{Name:"CLIENT_FORCE_PROTECT_ID_TOKEN", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_CONSENT_CACHE_ID_TOKEN", Value:false}, gosnowflake.nameValueParameter{Name:"CLIENT_STAGE_ARRAY_BINDING_THRESHOLD", Value:65280}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY", Value:3600}, gosnowflake.nameValueParameter{Name:"AUTOCOMMIT", Value:true}, gosnowflake.nameValueParameter{Name:"CLIENT_SESSION_CLONE", Value:false}, gosnowflake.nameValueParameter{Name:"TIMESTAMP_LTZ_OUTPUT_FORMAT", Value:""}}  LOG_SESSION_ID=14657069831119114
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_PREFETCH_THREADS, value: 4 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_OUTPUT_FORMAT, value: YYYY-MM-DD HH24:MI:SS.FF3 TZHTZM 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIME_OUTPUT_FORMAT, value: HH24:MI:SS 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_CHUNK_SIZE, value: 160 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_TZ_OUTPUT_FORMAT, value:  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_KEEP_ALIVE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_OUT_OF_BAND_TELEMETRY_ENABLED, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: QUERY_CONTEXT_CACHE_SIZE, value: 5 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_METADATA_USE_SESSION_DATABASE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: ENABLE_STAGE_S3_PRIVATELINK_FOR_US_EAST_1, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_PREFETCH_THREADS, value: 1 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_NTZ_OUTPUT_FORMAT, value: YYYY-MM-DD HH24:MI:SS.FF3 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_METADATA_REQUEST_USE_CONNECTION_CTX, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_HONOR_CLIENT_TZ_FOR_TIMESTAMP_NTZ, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_MEMORY_LIMIT, value: 1536 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TIMESTAMP_TYPE_MAPPING, value: TIMESTAMP_LTZ 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMEZONE, value: America/Los_Angeles 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_PREFETCH_SLOTS, value: 2 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TELEMETRY_ENABLED, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_USE_V1_QUERY_API, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_DISABLE_INCIDENTS, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_RESULT_COLUMN_CASE_INSENSITIVE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CSV_TIMESTAMP_FORMAT, value:  
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: BINARY_OUTPUT_FORMAT, value: HEX 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_ENABLE_LOG_INFO_STATEMENT_PARAMETERS, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_TELEMETRY_SESSIONLESS_ENABLED, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: DATE_OUTPUT_FORMAT, value: YYYY-MM-DD 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_FORCE_PROTECT_ID_TOKEN, value: true 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_CONSENT_CACHE_ID_TOKEN, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_STAGE_ARRAY_BINDING_THRESHOLD, value: 65280 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_KEEP_ALIVE_HEARTBEAT_FREQUENCY, value: 3600 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: AUTOCOMMIT, value: true     
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: CLIENT_SESSION_CLONE, value: false 
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf parameter. name: TIMESTAMP_LTZ_OUTPUT_FORMAT, value:  
DEBU[0000]rows.go:78 gosnowflake.(*snowflakeRows).Close Rows.Close                                    LOG_SESSION_ID=14657069831119114
INFO[0000]connection.go:268 gosnowflake.(*snowflakeConn).Close Close                                         LOG_SESSION_ID=14657069831119114
DEBU[0000]log.go:210 gosnowflake.(*defaultLogger).Debug nothing to send to telemetry                 
INFO[0000]restful.go:326 gosnowflake.closeSession close session                                 LOG_SESSION_ID=14657069831119114
INFO[0000]retry.go:300 gosnowflake.(*retryHTTP).execute retryHTTP.totalTimeout: 5s                    LOG_SESSION_ID=14657069831119114
DEBU[0000]log.go:148 gosnowflake.(*defaultLogger).Debugf retry count: 0                               
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1026139]

goroutine 1 [running]:
github.com/snowflakedb/gosnowflake.(*snowflakeChunkDownloader).getArrowBatches(0xc00083be10?)
        /home/vtan/go/pkg/mod/github.com/snowflakedb/gosnowflake@v1.7.2/chunk_downloader.go:252 +0x19
github.com/snowflakedb/gosnowflake.(*snowflakeRows).GetArrowBatches(0xc0001200d0?)
        /home/vtan/go/pkg/mod/github.com/snowflakedb/gosnowflake@v1.7.2/rows.go:171 +0x162
main.main()
        /home/vtan/repos/musical-octo-funicular/main.go:66 +0x306
exit status 2
```