import asyncio

from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy import asc, func
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.future import select

from src.core.database import get_session
from src.core.exception_handlers import unhandled_exception_handler
from src.custom_types import ErrorResponse, ListMetadata, UserResponse, UsersResponse
from src.models import User
from src.schemas import CreateUserSchema, FullUserSchema, UserSchema
from src.utils import utils

app = FastAPI()


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.add_exception_handler(Exception, unhandled_exception_handler)


@app.get("/users", response_model=UsersResponse[UserSchema])
async def get_users(
    page: int,
    per_page: int,
    session: AsyncSession = Depends(get_session),
) -> UsersResponse[User]:
    list_query = (
        select(User)
        .order_by(asc(User.id))
        .limit(per_page)
        .offset((page - 1) * per_page)
    )
    count_query = select(func.count(User.id))
    users = await session.scalars(list_query)
    total_count = await session.scalar(count_query) or 0
    metadata = ListMetadata(
        {"page": page, "per_page": per_page, "total_count": total_count}
    )
    await asyncio.sleep(0.4)
    return {"users": list(users), "metadata": metadata}


@app.get("/users/{user_public_id}", response_model=UserResponse[FullUserSchema])
async def get_user(
    user_public_id: str, session: AsyncSession = Depends(get_session)
) -> UserResponse[User] | ErrorResponse:
    # TODO: simulate delay
    # TODO: random error
    query = select(User).where(User.public_id == user_public_id)
    user = await session.scalar(query)
    if not user:
        raise HTTPException(
            detail={"code": "USER_NOT_FOUND", "message": "User not found"},
            status_code=status.HTTP_404_NOT_FOUND,
        )
    return {"user": user}


@app.post("/users", response_model=UserResponse[FullUserSchema])
async def create_user(
    data: CreateUserSchema, session: AsyncSession = Depends(get_session)
) -> UserResponse[User] | ErrorResponse:
    public_id = utils.ids.generateNano()
    user = User(
        public_id=public_id,
        name=data.name,
        email=data.email,
        phone=data.phone,
        sex=data.sex,
        birth_date=data.birth_date,
        role=data.role,
        is_active=data.is_active,
    )
    session.add(user)
    await session.flush()
    created_user = user
    await session.commit()
    return {"user": created_user}
