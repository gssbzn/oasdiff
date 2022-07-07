## Examples of breaking changes
- adding a new required property in request body is breaking
- adding a required request body is breaking
- changing an existing header param from optional to required is breaking
- changing an existing property in request body to required is breaking
- changing an existing property in request header to required is breaking
- changing an existing property in response body to optional is breaking
- changing an existing request body from optional to required is breaking
- changing an existing response header from required to optional is breaking
- changing max length from nil to any value is breaking
- deleting a media-type from response is breaking
- deleting a path is breaking
- deleting a required property in response body is breaking
- deleting a required property under AllOf in response body is breaking
- deleting an enum value is breaking
- deleting an operation is breaking
- increasing min items is breaking
- new required header param is breaking
- new required path param is breaking
- new required property in request header is breaking
- reducing max in request is breaking
- reducing max in response is breaking
- reducing max length is breaking

## Examples of non-breaking changes
- adding a media-type to response isn't breaking
- adding a new required property in response body isn't breaking
- adding a new required property under AllOf in response body isn't breaking
- adding a new required property under AllOf in response body isn't breaking but when multiple inline (without $ref) schemas under AllOf are modified simultaneously, we detect is as breaking
- adding an enum value isn't breaking
- adding an optional request body isn't breaking
- both max lengths are nil isn't breaking
- changing a link to operation ID isn't breaking
- changing an existing header param to optional isn't breaking
- changing an existing property in request body to optional isn't breaking
- changing an existing property in request header to optional isn't breaking
- changing an existing property in response body to required isn't breaking
- changing an existing request body from required to optional isn't breaking
- changing comments isn't breaking
- changing extensions isn't breaking
- changing max length from any value to nil isn't breaking
- changing operation ID isn't breaking
- deleting a required property in request isn't breaking
- deleting a tag isn't breaking
- deprecating a header isn't breaking
- deprecating a parameter isn't breaking
- deprecating a schema isn't breaking
- deprecating an operation isn't breaking
- increasing max length isn't breaking
- new optional header param isn't breaking
- new optional property in request header isn't breaking
- new required response header param isn't breaking
- no change isn't breaking
- reducing min items isn't breaking
- reducing min length isn't breaking