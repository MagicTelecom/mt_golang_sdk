/*
 * magictelecomapi
 *
 * This file was automatically generated by APIMATIC v2.0 on 06/23/2016
 */
package didsproducts

import(
    "github.com/apimatic/unirest-go"
    "magictelecomapi"
    "magictelecomapi/apihelper"

)
/*
 * Client structure as interface implementation
 */
type DIDSPRODUCTS_IMPL struct { }

/**
 * Allow clients to get the list of available phone_numbers
 * @param    *int           page       parameter: Optional
 * @param    *int           limit      parameter: Optional
 * @param    *string        filter     parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetDids (
            page *int,
            limit *int,
            filter *string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/dids"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "page" : apihelper.ToString(*page, "1"),
        "limit" : apihelper.ToString(*limit, "10"),
        "filter" : filter,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get a specific phone_number
 * @param    string        phoneNumber      parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetDidsByPhoneNumber (
            phoneNumber string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/dids/{phone_number}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "phone_number" : phoneNumber,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get the list of available locations.
 * @param    *int           page       parameter: Optional
 * @param    *int           limit      parameter: Optional
 * @param    *string        filter     parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetLocations (
            page *int,
            limit *int,
            filter *string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/locations"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "page" : apihelper.ToString(*page, "1"),
        "limit" : apihelper.ToString(*limit, "10"),
        "filter" : filter,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get a specific location.
 * @param    string        locationHandle      parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetLocationByHandle (
            locationHandle string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/locations/{location_handle}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "location_handle" : locationHandle,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get the list of available trunks
 * @param    *int        page      parameter: Optional
 * @param    *int        limit     parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetTrunks (
            page *int,
            limit *int) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/trunks"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "page" : apihelper.ToString(*page, "1"),
        "limit" : apihelper.ToString(*limit, "10"),
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get a specific trunk
 * @param    string        trunkHandle      parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetTrunkByHandle (
            trunkHandle string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/trunks/{trunk_handle}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithTemplateParameters(queryBuilder, map[string]interface{} {
        "trunk_handle" : trunkHandle,
    }) 
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Allow clients to get trunk zones.
 * @param    int        page      parameter: Required
 * @param    int        limit     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *DIDSPRODUCTS_IMPL) GetCountriesByTrunk (
            page int,
            limit int) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := magictelecomapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/dids/products/trunks/countries"

    //variable to hold errors
    var err error = nil
    
    //process optional query parameters
    queryBuilder, err = apihelper.AppendUrlWithQueryParameters(queryBuilder, map[string]interface{} {
        "page" : page,
        "limit" : limit,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : magictelecomapi.Config.XAuthToken,
    }

    //prepare API request
    request := unirest.Get(queryBuilder, headers)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("You are not authenticated", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("This action needs a valid WSSE header", response.Code, response.RawBody)
    } else if (response.Code == 404) {
        err = apihelper.NewAPIError("Resource not found", response.Code, response.RawBody)
    } else if (response.Code == 400) {
        err = apihelper.NewAPIError("Http bad request", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

