package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/mikkeloscar/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/mikkeloscar/gin-swagger/example/models"
)

func BusinessLogicCreateInfrastructureAccount(f func(ctx *gin.Context, params *CreateInfrastructureAccountParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &CreateInfrastructureAccountParams{}
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

// CreateInfrastructureAccountParams contains all the bound params for the create infrastructure account operation
// typically these are obtained from a http.Request
//
// swagger:parameters createInfrastructureAccount
type CreateInfrastructureAccountParams struct {

	/*Account that will be created.
	  Required: true
	  In: body
	*/
	InfrastructureAccount *models.InfrastructureAccount
}

// CreateInfrastructureAccountParamsFromCtx gets the params struct from the gin context.
func CreateInfrastructureAccountParamsFromCtx(ctx *gin.Context) *CreateInfrastructureAccountParams {
	params, _ := ctx.Get("params")
	return params.(*CreateInfrastructureAccountParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateInfrastructureAccountParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body models.InfrastructureAccount
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("infrastructureAccount", "body"))
			} else {
				res = append(res, errors.NewParseError("infrastructureAccount", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.InfrastructureAccount = &body
			}
		}

	} else {
		res = append(res, errors.Required("infrastructureAccount", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
