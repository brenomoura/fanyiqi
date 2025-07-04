from fastapi import APIRouter, Depends, HTTPException, Query
from pydantic import BaseModel

from app.api.routes.shared import get_current_user
from app.core.providers import PROVIDERS
from app.core.security import User
from app.services.translation import TranslationService

DEFAULT_TRANSLATION_MODEL = "michaelfeil/ct2fast-m2m100_418M"


router = APIRouter(prefix="", tags=["translation"])


class TranslationRequest(BaseModel):
    text: str
    source_language: str
    target_language: str
    model: str = DEFAULT_TRANSLATION_MODEL


class TranslationResponse(BaseModel):
    translated_text: str


@router.post("/translate", response_model=TranslationResponse)
async def translate(
    request: TranslationRequest,
    current_user: User = Depends(get_current_user),
):
    try:
        service = TranslationService(request.model)
        translated_text = service.translate(
            text=request.text,
            source_language=request.source_language,
            target_language=request.target_language,
        )
        return TranslationResponse(translated_text=translated_text)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/languages", response_model=list[tuple[str, str]])
async def get_available_languages(
    model: str = Query(default=DEFAULT_TRANSLATION_MODEL),
    current_user: User = Depends(get_current_user),
):
    try:
        service = TranslationService(model)
        languages = service.get_languages()
        return languages
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.get("/models", response_model=list[str])
async def get_available_models(current_user: User = Depends(get_current_user)):
    try:
        return list(PROVIDERS.keys())
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
