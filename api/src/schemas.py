from datetime import datetime

from pydantic import BaseModel


class CreateUserSchema(BaseModel):
    name: str
    email: str
    phone: str
    sex: str
    birth_date: datetime
    role: str
    is_active: bool


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
    phone: str
    sex: str
    birth_date: datetime
    role: str
    is_active: bool
    created_at: datetime
    updated_at: datetime

    class ConfigDict:
        from_attributes = True
