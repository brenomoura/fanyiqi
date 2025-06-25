import os
import secrets
from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    API_V1_STR: str = "/api/v1"
    SECRET_KEY: str = secrets.token_urlsafe(32)
    PROJECT_NAME: str = "fanyiqi"
    DEBUG: bool = True
    BASE_MODEL_PATH: str = os.environ.get("BASE_MODEL_PATH", "data/models")

settings = Settings()