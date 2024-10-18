from datetime import datetime
from typing import Optional

from sqlalchemy import BigInteger, String, func
from sqlalchemy.dialects.postgresql import JSONB
from sqlalchemy.orm import (
    DeclarativeBase,
    Mapped,
    MappedAsDataclass,
    mapped_column,
    registry,
)

table_registry = registry()


class Base(MappedAsDataclass, DeclarativeBase):
    pass


class User(Base):
    __tablename__ = "user"

    id: Mapped[int] = mapped_column(BigInteger, init=False, primary_key=True)
    public_id: Mapped[str] = mapped_column(String(12), index=True, unique=True)
    name: Mapped[str]
    email: Mapped[str]
    details: Mapped[Optional[dict[str, str | int | bool]]] = mapped_column(JSONB)
    created_at: Mapped[datetime] = mapped_column(init=False, insert_default=func.now())
    updated_at: Mapped[datetime] = mapped_column(
        init=False, insert_default=func.now(), onupdate=func.now()
    )