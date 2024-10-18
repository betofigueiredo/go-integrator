import os

from dotenv import load_dotenv
from sqlalchemy.orm import DeclarativeBase

load_dotenv()


class Base(DeclarativeBase):
    pass


class Settings:
    ENV: str | None = os.environ.get("ENV")
    PORT: int = int(os.environ.get("PORT") or 8000)
    DB_URL: str = f"postgresql+asyncpg://{os.environ.get('POSTGRES_USER')}:{os.environ.get('POSTGRES_PASSWORD')}@{os.environ.get('POSTGRES_DB')}/{os.environ.get('POSTGRES_DB')}"
    API_V1_PREFIX = "/v1"

    class Config:
        case_sensitive = True


settings = Settings()
