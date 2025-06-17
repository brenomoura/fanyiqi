from app.core.security import authenticate_user, create_user, get_user_from_token


class Auth:
    def __init__(self, db):
        self.db = db

    def create_user(self, username: str, email: str, password: str):
        user = create_user(self.db, username, email, password)
        return user

    def authenticate_user(self, username: str, password: str) -> tuple:
        user, token = authenticate_user(self.db, username, password)
        if not user:
            raise ValueError("Invalid username or password")
        return user, token
    
    def get_user_from_token(self, token: str):
        user = get_user_from_token(self.db, token)
        if not user:
            raise ValueError("Invalid or expired token")
        return user
