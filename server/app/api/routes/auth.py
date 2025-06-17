from fastapi import APIRouter, Depends, HTTPException
from pydantic import BaseModel
from sqlalchemy.orm import Session

from app.core.database import get_db
from app.services.auth import Auth

router = APIRouter(prefix="/auth", tags=["auth"])


class AuthRequest(BaseModel):
    username: str
    password: str


class RegisterRequest(BaseModel):
    username: str
    email: str
    password: str


class AuthResponse(BaseModel):
    token: str


@router.post("/generate-token", response_model=AuthResponse)
async def generate_token(payload: AuthRequest, db: Session = Depends(get_db)):
    try:
        _, token = Auth(db).authenticate_user(payload.username, payload.password)
        return AuthResponse(token=token)
    except ValueError as e:
        raise HTTPException(status_code=401, detail=str(e))


@router.post("/register", response_model=str)
async def register(payload: RegisterRequest, db: Session = Depends(get_db)):
    try:
        _ = Auth(db).create_user(
            payload.username, payload.email, payload.password
        )
        return "User registered successfully"
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
