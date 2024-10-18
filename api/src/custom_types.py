from typing import Generic, List, Optional, TypedDict, TypeVar

from fastapi import HTTPException

T = TypeVar("T")


class ListMetadata(TypedDict):
    page: Optional[int]
    per_page: Optional[int]
    total_count: int


class ErrorDetail(TypedDict):
    code: str
    message: str


class ErrorResponse(HTTPException):
    details: ErrorDetail
    status_code: int


class SuccessResponse(TypedDict):
    message: str


class UserResponse(TypedDict, Generic[T]):
    user: T


class UsersResponse(TypedDict, Generic[T]):
    users: List[T]
    metadata: ListMetadata
