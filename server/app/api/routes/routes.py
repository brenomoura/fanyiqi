from fastapi import APIRouter

from app.api.routes import translation

api_router = APIRouter()
api_router.include_router(translation.router)
