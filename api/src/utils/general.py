from typing import Any, Dict, Generic, NamedTuple, Type, TypeVar

from pydantic import BaseModel, ValidationError

R = TypeVar("R", bound=BaseModel)


class ErrorSchema(NamedTuple):
    message: str
    field: str


class SuccessResponse(NamedTuple, Generic[R]):
    fields: R
    error: None


class ErrorResponse(NamedTuple):
    fields: None
    error: ErrorSchema


class General:
    def validate_schema(
        self,
        schema: Type[R],
        params: dict[str, str | int | float | bool | None],
    ) -> SuccessResponse[R] | ErrorResponse:
        try:
            parsed_fields = schema(**params)
            return SuccessResponse(fields=parsed_fields, error=None)
        except ValidationError as exc:
            error = exc.errors()[0]
            error_msg = error.get("msg", "")
            error_field = str(error.get("loc", ())[0]) if error.get("loc", None) else ""
            error_response = ErrorSchema(message=error_msg, field=error_field)
            return ErrorResponse(fields=None, error=error_response)

    def as_dict(self, values: Any) -> Dict[str, Any]:
        return {c.name: str(getattr(values, c.name)) for c in values.__table__.columns}
