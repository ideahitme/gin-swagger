package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/mikkeloscar/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"
)

func BusinessLogicListClusters(f func(ctx *gin.Context, params *ListClustersParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &ListClustersParams{}
		err := params.bindRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			ctx.JSON(int(errObj.Code()), errObj)
			return
		}

		resp := f(ctx, params)
		ctx.JSON(resp.Code, resp.Body)
	}
}

// ListClustersParams contains all the bound params for the list clusters operation
// typically these are obtained from a http.Request
//
// swagger:parameters listClusters
type ListClustersParams struct {

	/*Filter on cluster alias.
	  In: query
	*/
	Alias *string
	/*Filter on API server URL.
	  In: query
	*/
	APIServerURL *string
	/*Filter on channel.
	  In: query
	*/
	Channel *string
	/*Filter on criticality level.
	  In: query
	*/
	CriticalityLevel *int32
	/*Filter on environment.
	  In: query
	*/
	Environment *string
	/*Filter on infrastructure account.
	  In: query
	*/
	InfrastructureAccount *string
	/*Filter on cluster lifecycle status.
	  In: query
	*/
	LifecycleStatus *string
	/*Filter on local id.
	  In: query
	*/
	LocalID *string
	/*Filter on provider.
	  In: query
	*/
	Provider *string
	/*Filter on region.
	  In: query
	*/
	Region *string
}

// ListClustersParamsFromCtx gets the params struct from the gin context.
func ListClustersParamsFromCtx(ctx *gin.Context) *ListClustersParams {
	params, _ := ctx.Get("params")
	return params.(*ListClustersParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListClustersParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	qs := runtime.Values(ctx.Request.URL.Query())

	qAlias, qhkAlias, _ := qs.GetOK("alias")
	if err := o.bindAlias(qAlias, qhkAlias, formats); err != nil {
		res = append(res, err)
	}

	qAPIServerURL, qhkAPIServerURL, _ := qs.GetOK("api_server_url")
	if err := o.bindAPIServerURL(qAPIServerURL, qhkAPIServerURL, formats); err != nil {
		res = append(res, err)
	}

	qChannel, qhkChannel, _ := qs.GetOK("channel")
	if err := o.bindChannel(qChannel, qhkChannel, formats); err != nil {
		res = append(res, err)
	}

	qCriticalityLevel, qhkCriticalityLevel, _ := qs.GetOK("criticality_level")
	if err := o.bindCriticalityLevel(qCriticalityLevel, qhkCriticalityLevel, formats); err != nil {
		res = append(res, err)
	}

	qEnvironment, qhkEnvironment, _ := qs.GetOK("environment")
	if err := o.bindEnvironment(qEnvironment, qhkEnvironment, formats); err != nil {
		res = append(res, err)
	}

	qInfrastructureAccount, qhkInfrastructureAccount, _ := qs.GetOK("infrastructure_account")
	if err := o.bindInfrastructureAccount(qInfrastructureAccount, qhkInfrastructureAccount, formats); err != nil {
		res = append(res, err)
	}

	qLifecycleStatus, qhkLifecycleStatus, _ := qs.GetOK("lifecycle_status")
	if err := o.bindLifecycleStatus(qLifecycleStatus, qhkLifecycleStatus, formats); err != nil {
		res = append(res, err)
	}

	qLocalID, qhkLocalID, _ := qs.GetOK("local_id")
	if err := o.bindLocalID(qLocalID, qhkLocalID, formats); err != nil {
		res = append(res, err)
	}

	qProvider, qhkProvider, _ := qs.GetOK("provider")
	if err := o.bindProvider(qProvider, qhkProvider, formats); err != nil {
		res = append(res, err)
	}

	qRegion, qhkRegion, _ := qs.GetOK("region")
	if err := o.bindRegion(qRegion, qhkRegion, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListClustersParams) bindAlias(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Alias = &raw

	return nil
}

func (o *ListClustersParams) bindAPIServerURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.APIServerURL = &raw

	return nil
}

func (o *ListClustersParams) bindChannel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Channel = &raw

	return nil
}

func (o *ListClustersParams) bindCriticalityLevel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("criticality_level", "query", "int32", raw)
	}
	o.CriticalityLevel = &value

	return nil
}

func (o *ListClustersParams) bindEnvironment(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Environment = &raw

	return nil
}

func (o *ListClustersParams) bindInfrastructureAccount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.InfrastructureAccount = &raw

	return nil
}

func (o *ListClustersParams) bindLifecycleStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LifecycleStatus = &raw

	if err := o.validateLifecycleStatus(formats); err != nil {
		return err
	}

	return nil
}

func (o *ListClustersParams) validateLifecycleStatus(formats strfmt.Registry) error {

	if err := validate.Enum("lifecycle_status", "query", *o.LifecycleStatus, []interface{}{"requested", "creating", "ready", "decommission-requested", "decommissioned"}); err != nil {
		return err
	}

	return nil
}

func (o *ListClustersParams) bindLocalID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LocalID = &raw

	return nil
}

func (o *ListClustersParams) bindProvider(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Provider = &raw

	return nil
}

func (o *ListClustersParams) bindRegion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Region = &raw

	return nil
}
