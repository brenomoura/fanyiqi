import hashlib
import secrets
from datetime import datetime, timedelta, timezone

from sqlalchemy import Column, DateTime, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import Session
from werkzeug.security import generate_password_hash, check_password_hash

Base = declarative_base()


class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, autoincrement=True)
    username = Column(String(50), unique=True, nullable=False)
    email = Column(String(120), unique=True, nullable=False)
    password_hash = Column(String(128), nullable=False)


class Token(Base):
    __tablename__ = "tokens"

    id = Column(Integer, primary_key=True, autoincrement=True)
    user_id = Column(Integer, nullable=False)
    token_hash = Column(String(64), nullable=False, unique=True)
    expires_at = Column(DateTime, nullable=True)


def generate_token() -> str:
    return secrets.token_hex(32)


def hash_token(token: str) -> str:
    return hashlib.sha256(token.encode()).hexdigest()


def create_user(db: Session, username: str, email: str, password: str):
    password_hash = generate_password_hash(password)

    user = User(
        username=username,
        email=email,
        password_hash=password_hash,
    )

    db.add(user)
    db.commit()
    db.refresh(user)

    return user


def authenticate_user(db: Session, username: str, plain_password: str):
    user = db.query(User).filter(User.username == username).first()
    if not user:
        return None

    if not check_password_hash(user.password_hash, plain_password):
        return None

    token = generate_token()
    token_hash = hash_token(token)
    expires_at = datetime.now(timezone.utc) + timedelta(days=7)

    token_obj = Token(user_id=user.id, token_hash=token_hash, expires_at=expires_at)
    db.add(token_obj)
    db.commit()

    return user, token


def get_user_from_token(db: Session, token: str) -> User | None:
    token_hash = hash_token(token)
    token_obj = db.query(Token).filter(Token.token_hash == token_hash).first()

    if token_obj.expires_at.tzinfo is None:
        expires_at = token_obj.expires_at.replace(tzinfo=timezone.utc)
    else:
        expires_at = token_obj.expires_at

    if not token_obj or expires_at < datetime.now(timezone.utc):
        return None

    return db.query(User).filter(User.id == token_obj.user_id).first()
