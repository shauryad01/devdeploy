from pydantic import BaseModel

class ProductCreate(BaseModel):
    name:str
    price:int

class ProductOut(BaseModel):
    id:int
    name:str
    price:int

    class Config:
        from_attributes = True