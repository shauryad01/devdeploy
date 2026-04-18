from pydantic import BaseModel, EmailStr

class UserCreate(BaseModel):
    email: EmailStr
    password : str

class UserOut(BaseModel):
    id: int
    email: str

    class Config:
        from_attributes = True