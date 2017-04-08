package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// BindGetInfrastructureAccount validates and binds request parameters to the gin
// context.
func BindGetInfrastructureAccount(ctx *gin.Context) {
	params := &GetInfrastructureAccountParams{}
	err := params.bindRequest(ctx)
	if err != nil {
		errObj := err.(*errors.CompositeError)
		ctx.JSON(int(errObj.Code()), errObj)
		return
	}
	ctx.Set("params", params)
	ctx.Next()
}

// GetInfrastructureAccountParams contains all the bound params for the get infrastructure account operation
// typically these are obtained from a http.Request
//
// swagger:parameters getInfrastructureAccount
type GetInfrastructureAccountParams struct {

	/*ID of the infrastructure account.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-:]*[a-z0-9]$
	  In: path
	*/
	AccountID string
}

// GetInfrastructureAccountParamsFromCtx gets the params struct from the gin context.
func GetInfrastructureAccountParamsFromCtx(ctx *gin.Context) *GetInfrastructureAccountParams {
	params, _ := ctx.Get("params")
	return params.(*GetInfrastructureAccountParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetInfrastructureAccountParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rAccountID := []string{ctx.Param("account_id")}
	if err := o.bindAccountID(rAccountID, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInfrastructureAccountParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.AccountID = raw

	if err := o.validateAccountID(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetInfrastructureAccountParams) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Pattern("account_id", "path", o.AccountID, `^[a-z][a-z0-9-:]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}