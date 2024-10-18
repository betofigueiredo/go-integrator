from datetime import datetime

from pydantic import BaseModel


class CreateUserSchema(BaseModel):
    name: str
    email: str


class UserSchema(BaseModel):
    public_id: str
    name: str
    email: str

    class ConfigDict:
        from_attributes = True


class FullUserSchema(BaseModel):
    public_id: str
    name: str
    email: str
    details: dict[str, str | int | float | bool | None]
    created_at: datetime
    updated_at: datetime

    class ConfigDict:
        from_attributes = True
